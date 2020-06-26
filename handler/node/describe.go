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
package node

import (
	"fmt"
	"math/big"
	"strings"

	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/ptypes"
)

var describeNodeCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	// describeNodeCmd represents the describeNode command
	describeNodeCmd = &cobra.Command{
		Use:   "describe [NAME]",
		Short: "Describe a Portworx node",
		Long:  "Show detailed information about a Portworx node",
		Example: `
  # Display detailed information about Portworx cluster
  pxc node describe abc`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a node")
			}
			return nil
		},
		RunE: describeNodeExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	NodeAddCommand(describeNodeCmd)
})

func describeNodeExec(c *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	id := args[0]
	// Try to use the id of the user
	var n *api.StorageNode
	nodes := api.NewOpenStorageNodeClient(conn)
	node, err := nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: id})
	if err != nil {
		// Try by looking at each node
		nodesInfo, err := nodes.Enumerate(ctx, &api.SdkNodeEnumerateRequest{})
		if err != nil {
			return util.PxErrorMessage(err, "Failed to get node")
		}
		for _, nodeID := range nodesInfo.GetNodeIds() {
			node, err = nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: nodeID})
			if err != nil {
				continue
			}

			if id == node.GetNode().GetHostname() ||
				id == node.GetNode().GetMgmtIp() ||
				id == node.GetNode().GetSchedulerNodeName() {
				n = node.GetNode()
				break
			}
		}
	} else {
		n = node.GetNode()
	}

	if n != nil {
		return describeStorageNode(n)
	}

	return fmt.Errorf("Node %s not found", id)

}

func describeStorageNode(n *api.StorageNode) error {
	util.Printf("Status: %s\n"+
		"Host name: %s\n"+
		"Scheduler name: %s\n"+
		"Node ID: %s\n"+
		"Version: %s\n"+
		"Kernel Version: %s\n"+
		"Operating System: %s\n"+
		"Management IP: %s\n"+
		"Data IP: %s\n",
		util.SdkStatusToPrettyString(n.GetStatus()),
		n.GetHostname(),
		n.GetSchedulerNodeName(),
		n.GetId(),
		portworx.GetStorageNodeVersion(n),
		portworx.GetStorageNodeKernelVersion(n),
		portworx.GetStorageNodeOS(n),
		n.GetMgmtIp(),
		n.GetDataIp())

	// Print pools
	util.Printf("\n")
	util.Printf("Number of Pools: %d\n", len(n.GetPools()))
	t := util.NewTabby()
	t.AddHeader("POOL", "IO_PRIORITY", "RAID_LEVEL", "USABLE", "USED", "STATUS", "ZONE", "REGION")
	for _, pool := range n.GetPools() {
		id := pool.GetUuid()
		if len(id) == 0 {
			id = fmt.Sprintf("%d", pool.GetID())
		}
		t.AddLine(id,
			strings.TrimPrefix(pool.GetCos().String(), "CosType_"),
			pool.GetRaidLevel(),
			humanize.BigIBytes(big.NewInt(int64(pool.GetTotalSize()))),
			humanize.BigIBytes(big.NewInt(int64(pool.GetUsed()))),
			"Unknown",
			"Unknown",
			"Unknown")
	}
	t.Print()

	// Print devices
	util.Printf("\n")
	util.Printf("Number of Disks: %d\n", len(n.GetDisks()))
	t = util.NewTabby()
	t.AddHeader("DEVICE", "PATH", "MEDIA TYPE", "SIZE", "USED", "STATUS", "LAST-SCAN")
	for _, disk := range n.GetDisks() {
		status := "Offline"
		if disk.GetOnline() {
			status = "Online"
		}
		t.AddLine(disk.GetId(),
			disk.GetPath(),
			strings.TrimPrefix(disk.GetMedium().String(), "STORAGE_MEDIUM_"),
			humanize.BigIBytes(big.NewInt(int64(disk.GetSize()))),
			humanize.BigIBytes(big.NewInt(int64(disk.GetUsed()))),
			status,
			ptypes.TimestampString(disk.GetLastScan()))

	}
	t.Print()

	util.Printf("\nCache Devices\nUnknown\n")
	util.Printf("\nJournal Devices\nUnknown\n")
	util.Printf("\n")
	return nil
}
