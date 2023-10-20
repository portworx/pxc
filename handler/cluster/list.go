/*
Copyright © 2020 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cluster

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	"github.com/dustin/go-humanize"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"

	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type clusterListOptions struct {
	contextMatch string
}

// clusterCmd represents the cluster command
var (
	clusterListCmd  *cobra.Command
	clusterListArgs *clusterListOptions
)

var _ = commander.RegisterCommandVar(func() {

	clusterListArgs = &clusterListOptions{}
	clusterListCmd = &cobra.Command{
		Use: "list-by-context",
		Example: `
  # Scan through all the contexts in your kubeconfig and display cluster information
  pxc cluster list-by-context

  # Show cluster information for contexts that match contain 'east-coast' in the name
  pxc cluster list-by-context --context-match='*east-coast*,*south*'

  # Output all Kubernetes and Portworx cluster information as json for your contexts that match '*on-prem*'
  pxc cluster list-by-context -o json --context-match='*on-prem*'`,
		Aliases: []string{"get", "list", "show"},
		Short:   "Show Portworx and Kubernetes information for every context in your kubeconfig",
		RunE:    listClusterExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ClusterAddCommand(clusterListCmd)
	clusterListCmd.Flags().StringVar(&clusterListArgs.contextMatch,
		"context-match", "", "Comma separated list of expressions match the appropriate context")
	clusterListCmd.Flags().StringP("output", "o", "", "Output in yaml|json|wide")
})

func ClusterListAddCommand(cmd *cobra.Command) {
	clusterCmd.AddCommand(cmd)
}

func listClusterExec(cmd *cobra.Command, args []string) error {
	// Parse out all of the common cli volume flags
	cvi := cliops.NewCliInputs(cmd, args)

	// Create the parser object
	cgf := NewClustersGetFormatter(cvi, args)

	// Print the details and return errors if any
	return util.PrintFormatted(cgf)
}

type KubernetesClusterInfo struct {
	Context string    `json:"context" yaml:"context"`
	Cluster string    `json:"cluster" yaml:"cluster"`
	Nodes   []v1.Node `json:"nodes" yaml:"nodes"`
}

type PortworxClusterInfo struct {
	Cluster *api.StorageCluster `json:"cluster" yaml:"cluster"`
	Nodes   []*api.StorageNode  `json:"nodes" yaml:"nodes"`
}

type ClusterInfo struct {
	Kubernetes *KubernetesClusterInfo `json:"kubernetes" yaml:"kubernetes"`
	Portworx   *PortworxClusterInfo   `json:"portworx" yaml:"portworx"`
}

func NewClusterInfo(context, clusterName string) *ClusterInfo {
	return &ClusterInfo{
		Kubernetes: &KubernetesClusterInfo{
			Context: context,
			Cluster: clusterName,
			Nodes:   make([]v1.Node, 0),
		},
		Portworx: &PortworxClusterInfo{
			Cluster: &api.StorageCluster{},
			Nodes:   make([]*api.StorageNode, 0),
		},
	}
}

type clustersGetFormatter struct {
	util.BaseFormatOutput
	cli  *cliops.CliInputs
	args []string
}

func NewClustersGetFormatter(cli *cliops.CliInputs, args []string) *clustersGetFormatter {
	f := &clustersGetFormatter{
		cli:  cli,
		args: args,
	}
	f.FormatType = cli.FormatType
	return f
}

func (f *clustersGetFormatter) getClusters() ([]*ClusterInfo, error) {

	clusterInfos := make([]*ClusterInfo, 0)

	config.CM().ForEachContext(func(ctxName, clusterName string) error {
		contextMatches := strings.Split(clusterListArgs.contextMatch, ",")
		if len(clusterListArgs.contextMatch) > 0 &&
			!util.ListContains(contextMatches, ctxName) &&
			!util.ListMatchGlob(contextMatches, ctxName) {
			return nil
		}

		// Initialize objects
		clusterInfo := NewClusterInfo(ctxName, clusterName)
		clusterInfos = append(clusterInfos, clusterInfo)

		// Get K8S Nodes
		_, cs, err := kubernetes.KubeConnectDefault()
		if err != nil {
			return fmt.Errorf("Unable to create a kubernetes connection: %v", err)
		}

		// Get Kubernetes nodes
		knodes := cs.CoreV1().Nodes()
		knodesInfo, err := knodes.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("Unable to get kubernetes nodes: %v", err)
		}
		clusterInfo.Kubernetes.Nodes = knodesInfo.Items

		// Get StorageCluster and Nodes
		ctx, conn, err := portworx.PxConnectDefault()
		if err != nil {
			return err
		}
		defer conn.Close()
		defer kubernetes.StopTunnel()

		// cluster information
		cluster := api.NewOpenStorageClusterClient(conn)
		currentClusterInfo, err := cluster.InspectCurrent(ctx, &api.SdkClusterInspectCurrentRequest{})
		if err != nil {
			return util.PxErrorMessage(err, "Failed to inspect cluster")
		}
		clusterInfo.Portworx.Cluster = currentClusterInfo.GetCluster()

		// Get all nodes
		pxops, err := portworx.NewPxOps()
		if err != nil {
			return err
		}
		nodes := portworx.NewNodes(pxops, &portworx.NodeSpec{})
		nodesInfo, err := nodes.GetNodes()
		if err != nil {
			return err
		}
		clusterInfo.Portworx.Nodes = nodesInfo

		return nil
	})

	sort.Slice(clusterInfos, func(i, j int) bool {
		return clusterInfos[i].Kubernetes.Context < clusterInfos[j].Kubernetes.Context
	})
	return clusterInfos, nil
}

func (f *clustersGetFormatter) YamlFormat() (string, error) {
	infos, err := f.getClusters()
	if err != nil {
		return "", err
	}
	return util.ToYaml(infos)
}

func (f *clustersGetFormatter) JsonFormat() (string, error) {
	infos, err := f.getClusters()
	if err != nil {
		return "", err
	}
	return util.ToJson(infos)
}

func (f *clustersGetFormatter) WideFormat() (string, error) {
	return f.toTabbed()
}

func (f *clustersGetFormatter) DefaultFormat() (string, error) {
	return f.toTabbed()
}

func (f *clustersGetFormatter) toTabbed() (string, error) {
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(writer)

	clusters, _ := f.getClusters()

	currentContext, err := config.CM().ConfigGetCurrentContext()
	if err != nil {
		return "", fmt.Errorf("Unable to get current context: %v", err)
	}

	if len(clusters) == 0 {
		util.Printf("No resources found\n")
		return "", nil
	}

	// Start the columns
	t.AddHeader(f.getHeader()...)

	for _, c := range clusters {
		l, err := f.getLine(c, currentContext)
		if err != nil {
			return "", nil
		}
		t.AddLine(l...)
	}
	t.Print()

	return b.String(), nil
}

func (f *clustersGetFormatter) getHeader() []interface{} {
	var header []interface{}
	if f.cli.Wide {
		header = []interface{}{"K8S CONTEXT", "K8S CLUSTER", "K8S VERSION", "PX CLUSTER", "UUID", "VERSION", "USED", "CAPACITY", "# NODES", "STATUS"}
	} else {
		header = []interface{}{"K8S CONTEXT", "K8S CLUSTER", "K8S VERSION", "PX CLUSTER", "VERSION", "USED", "CAPACITY", "STATUS"}
	}

	return header
}

func (f *clustersGetFormatter) getLine(cluster *ClusterInfo, currentContext string) ([]interface{}, error) {

	k8sversion := "(Unknown)"
	if len(cluster.Kubernetes.Nodes) > 0 {
		k8sversion = cluster.Kubernetes.Nodes[0].Status.NodeInfo.KubeletVersion
	}

	// Display current context with a *
	context := cluster.Kubernetes.Context
	if context == currentContext {
		context = "*" + context
	}

	// Kubernetes Information
	line := []interface{}{
		context,
		cluster.Kubernetes.Cluster,
		k8sversion,
	}

	// If there is no portworx information, just return here
	if len(cluster.Portworx.Cluster.GetName()) == 0 {
		if f.cli.Wide {
			line = append(line, []interface{}{
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			}...)
		} else {
			line = append(line, []interface{}{
				"",
				"",
				"",
				"",
				"",
			}...)
		}
		return line, nil
	}

	// Portworx info
	pxversion := "(Unknown)"
	if len(cluster.Portworx.Nodes) > 0 {
		pxversion = portworx.GetStorageNodeVersion(cluster.Portworx.Nodes[0])
	}

	used, capacity := totalCapacity(cluster.Portworx.Nodes)
	usedStr := humanize.BigIBytes(big.NewInt(int64(used)))
	capacityStr := humanize.BigIBytes(big.NewInt(int64(capacity)))

	if f.cli.Wide {
		line = append(line, []interface{}{
			cluster.Portworx.Cluster.GetName(),
			cluster.Portworx.Cluster.GetId(),
			pxversion,
			usedStr,
			capacityStr,
			len(cluster.Portworx.Nodes),
			util.SdkStatusToPrettyString(cluster.Portworx.Cluster.GetStatus()),
		}...)
	} else {
		line = append(line, []interface{}{
			cluster.Portworx.Cluster.Name,
			pxversion,
			usedStr,
			capacityStr,
			util.SdkStatusToPrettyString(cluster.Portworx.Cluster.GetStatus()),
		}...)
	}

	return line, nil
}

func totalCapacity(nodes []*api.StorageNode) (used, capacity uint64) {
	for _, node := range nodes {
		node_used, node_capacity := portworx.GetTotalCapacity(node)
		used += node_used
		capacity += node_capacity
	}
	return used, capacity
}
