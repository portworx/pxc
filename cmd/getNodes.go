/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"

	"github.com/cheynewallace/tabby"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// getNodesCmd represents the getNodes command
var getNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getNodesExec(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getNodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//getNodesCmd.Flags().BoolP("output", "y", false, "Output to yaml")
}

func getNodesExec(cmd *cobra.Command, args []string) {
	ctx, conn := pxConnect()
	defer conn.Close()

	// Get all node Ids
	nodes := api.NewOpenStorageNodeClient(conn)
	nodesInfo, err := nodes.Enumerate(ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		pxPrintGrpcErrorWithMessage(err, "Failed to get nodes")
		return
	}

	// Get all node info
	storageNodes := make([]*api.StorageNode, 0, len(nodesInfo.GetNodeIds()))
	for _, nid := range nodesInfo.GetNodeIds() {
		node, err := nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: nid})
		if err != nil {
			pxPrintGrpcErrorWithMessagef(err, "Failed to get information about node %s", nid)
			continue
		}
		n := node.GetNode()

		// Check if we have been asked for specific node
		if len(args) != 0 && !listHaveMatch(args, []string{n.GetId(), n.GetHostname(), n.GetMgmtIp(), n.GetSchedulerNodeName()}) {
			continue
		}

		storageNodes = append(storageNodes, node.GetNode())
	}

	// Get output
	output, _ := cmd.Flags().GetString("output")
	switch output {
	case "yaml":
		getNodesYamlPrinter(cmd, args, storageNodes)
	case "json":
		getNodesJsonPrinter(cmd, args, storageNodes)
	case "wide":
		getNodesWidePrinter(cmd, args, storageNodes)
	default:
		getNodesDefaultPrinter(cmd, args, storageNodes)
	}
}

func getNodesYamlPrinter(cmd *cobra.Command, args []string, storageNodes []*api.StorageNode) {
	bytes, err := yaml.Marshal(storageNodes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create yaml output")
		return
	}
	fmt.Println(string(bytes))
}

func getNodesJsonPrinter(cmd *cobra.Command, args []string, storageNodes []*api.StorageNode) {
	bytes, err := json.MarshalIndent(storageNodes, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create json output")
		return
	}
	fmt.Println(string(bytes))
}

func getNodesWidePrinter(cmd *cobra.Command, args []string, storageNodes []*api.StorageNode) {
	t := tabby.New()
	t.AddHeader("Id", "Hostname", "IP", "SchedulerNodeName", "Used", "Capacity", "Status", "Labels")
	for _, n := range storageNodes {
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

		labels := make([]string, 0, len(n.GetNodeLabels()))
		for k, v := range n.GetNodeLabels() {
			labels = append(labels, k+"="+v)
		}

		t.AddLine(n.GetId(),
			n.GetHostname(),
			n.GetMgmtIp(),
			n.GetSchedulerNodeName(),
			usedStr,
			capacityStr,
			n.GetStatus(),
			strings.Join(labels, ","))
	}
	t.Print()

}

func getNodesDefaultPrinter(cmd *cobra.Command, args []string, storageNodes []*api.StorageNode) {
	t := tabby.New()
	t.AddHeader("Hostname", "IP", "SchedulerNodeName", "Used", "Capacity", "Status")
	for _, n := range storageNodes {
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

		t.AddLine(n.GetHostname(),
			n.GetMgmtIp(),
			n.GetSchedulerNodeName(),
			usedStr,
			capacityStr,
			n.GetStatus())
	}
	t.Print()
}
