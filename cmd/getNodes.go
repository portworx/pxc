/*
Copyright © 2019 Portworx

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
package cmd

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"

	"github.com/spf13/cobra"
)

// getNodesCmd represents the getNodes command
var getNodesCmd = &cobra.Command{
	Use:     "node",
	Aliases: []string{"nodes"},
	Short:   "Get Portworx node information",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getNodesExec(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getNodesCmd)
}

func getNodesExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectCurrent(GetConfigFile())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get all node Ids
	nodes := api.NewOpenStorageNodeClient(conn)
	nodesInfo, err := nodes.Enumerate(ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get node information")
	}

	// Get all node info
	storageNodes := make([]*api.StorageNode, 0, len(nodesInfo.GetNodeIds()))
	for _, nid := range nodesInfo.GetNodeIds() {
		node, err := nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: nid})
		if err != nil {
			// Just print it and continue to other nodes
			util.PrintPxErrorMessagef(err, "Failed to get information about node %s", nid)
			continue
		}
		n := node.GetNode()

		// Check if we have been asked for specific node
		if len(args) != 0 &&
			!util.ListHaveMatch(args,
				[]string{n.GetId(),
					n.GetHostname(),
					n.GetMgmtIp(),
					n.GetSchedulerNodeName()}) {
			continue
		}

		storageNodes = append(storageNodes, node.GetNode())
	}

	// Get output
	output, _ := cmd.Flags().GetString("output")
	switch output {
	case "yaml":
		util.PrintYaml(storageNodes)
	case "json":
		util.PrintJson(storageNodes)
	case "wide":
		// We can have a special one here, but for simplicity, we will use the
		// default printer
		fallthrough
	default:
		getNodesDefaultPrinter(cmd, args, storageNodes)
	}

	return nil
}

func getNodesDefaultPrinter(cmd *cobra.Command, args []string, storageNodes []*api.StorageNode) {

	// Determine if it is a wide output
	output, _ := cmd.Flags().GetString("output")
	wide := output == "wide"

	// Determine if we need to show labels
	showLabels, _ := cmd.Flags().GetBool("show-labels")

	// Start the columns
	t := util.NewTabby()
	np := &nodeColumnFormatter{wide: wide, showLabels: showLabels}
	t.AddHeader(np.getHeader()...)

	for _, n := range storageNodes {
		t.AddLine(np.getLine(n)...)
	}
	t.Print()
}

type nodeColumnFormatter struct {
	wide       bool
	showLabels bool
}

func (p *nodeColumnFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.wide {
		header = []interface{}{"Id", "Hostname", "IP", "Data IP", "SchedulerNodeName", "Used", "Capacity", "# Disks", "# Pools", "Status"}
	} else {
		header = []interface{}{"Hostname", "IP", "SchedulerNodeName", "Used", "Capacity", "Status"}
	}
	if p.showLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *nodeColumnFormatter) getLine(n *api.StorageNode) []interface{} {

	// Calculate used
	var (
		used, capacity uint64
	)
	for _, pool := range n.GetPools() {
		used += pool.GetUsed()
		capacity += pool.GetTotalSize()
	}
	usedStr := fmt.Sprintf("%d Gi", used/Gi)
	capacityStr := fmt.Sprintf("%d Gi", capacity/Gi)

	// Return a line
	var line []interface{}
	if p.wide {
		line = []interface{}{
			n.GetId(), n.GetHostname(), n.GetMgmtIp(),
			n.GetDataIp(), n.GetSchedulerNodeName(), usedStr, capacityStr,
			len(n.GetDisks()), len(n.GetPools()), n.GetStatus(),
		}
	} else {
		line = []interface{}{
			n.GetHostname(), n.GetMgmtIp(),
			n.GetSchedulerNodeName(), usedStr, capacityStr,
			n.GetStatus(),
		}
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(n.GetNodeLabels()))
	}
	return line
}
