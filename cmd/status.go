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
	"github.com/portworx/px/pkg/util"

	"github.com/spf13/cobra"
)

var statusCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	statusCmd = &cobra.Command{
		Use: "status",
		// TODO
		Short: "TODO: this will move to px describe cluster",
		RunE:  statusExec,
	}
})

var _ = RegisterCommandInit(func() {
	rootCmd.AddCommand(statusCmd)
})

func statusExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get Identity information
	identity := api.NewOpenStorageIdentityClient(conn)
	version, err := identity.Version(ctx, &api.SdkIdentityVersionRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get server version")
	}
	var versionDetails string
	for k, v := range version.GetVersion().GetDetails() {
		versionDetails += fmt.Sprintf("  %s: %s\n", k, v)
	}

	// Print the cluster information
	cluster := api.NewOpenStorageClusterClient(conn)
	clusterInfo, err := cluster.InspectCurrent(ctx, &api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to inspect cluster")
	}

	util.Printf("Cluster ID: %s\n"+
		"Cluster UUID: %s\n"+
		"Cluster Status: %s\n"+
		"Version: %s\n"+
		"%s"+
		"SDK Version %s\n"+
		"\n",
		clusterInfo.GetCluster().GetName(),
		clusterInfo.GetCluster().GetId(),
		clusterInfo.GetCluster().GetStatus(),
		version.GetVersion().GetVersion(),
		versionDetails,
		version.GetSdkVersion().GetVersion())

	// Get all node Ids
	nodes := api.NewOpenStorageNodeClient(conn)
	nodesInfo, err := nodes.Enumerate(ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get nodes")
	}

	t := util.NewTabby()
	t.AddHeader("Hostname", "IP", "SchedulerNodeName", "Used", "Capacity", "Status")
	for _, nid := range nodesInfo.GetNodeIds() {
		node, err := nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: nid})
		if err != nil {
			return util.PxErrorMessagef(err, "Failed to get information about node %s", nid)
		}
		n := node.GetNode()

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

		t.AddLine(n.GetHostname(), n.GetMgmtIp(), n.GetSchedulerNodeName(), usedStr, capacityStr, n.GetStatus())
	}
	t.Print()

	return nil
}
