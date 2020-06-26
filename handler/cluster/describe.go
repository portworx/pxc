// Copyright © 2019 Portworx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"fmt"
	"math/big"

	"github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

var describeClusterCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	// describeClusterCmd represents the describeCluster command
	describeClusterCmd = &cobra.Command{
		Use:   "describe [NAME]",
		Short: "Describe a Portworx cluster",
		Long:  "Show detailed information of a Portworx cluster",
		RunE:  describeClusterExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ClusterAddCommand(describeClusterCmd)
})

func describeClusterExec(c *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
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
		versionDetails += fmt.Sprintf("%s:%s", k, v)
	}

	// Print cluster information
	cluster := api.NewOpenStorageClusterClient(conn)
	clusterInfo, err := cluster.InspectCurrent(ctx, &api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to inspect cluster")
	}

	logrus.Infof("Information about the node which was used: SDK Version[%s] MoreInfo[%s]",
		version.GetSdkVersion().GetVersion(),
		versionDetails)

	util.Printf("Name: %s\n"+
		"UUID: %s\n"+
		"Status: %s\n",
		clusterInfo.GetCluster().GetName(),
		clusterInfo.GetCluster().GetId(),
		util.SdkStatusToPrettyString(clusterInfo.GetCluster().GetStatus()))

	// Get all node Ids
	nodes := api.NewOpenStorageNodeClient(conn)
	nodesInfo, err := nodes.Enumerate(ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get nodes")
	}

	util.Printf("\nNodes:\n")
	t := util.NewTabby()
	t.AddHeader("Hostname", "Version", "Used", "Capacity", "Status", "Kernel Version", "OS")
	for _, nid := range nodesInfo.GetNodeIds() {
		node, err := nodes.Inspect(ctx, &api.SdkNodeInspectRequest{NodeId: nid})
		if err != nil {
			return util.PxErrorMessagef(err, "Failed to get information about node %s", nid)
		}
		n := node.GetNode()

		used, capacity := portworx.GetTotalCapacity(n)
		usedStr := humanize.BigIBytes(big.NewInt(int64(used)))
		capacityStr := humanize.BigIBytes(big.NewInt(int64(capacity)))
		kernelVersionStr := portworx.GetStorageNodeKernelVersion(n)
		osFlavorStr := portworx.GetStorageNodeOS(n)

		t.AddLine(n.GetHostname(), portworx.GetStorageNodeVersion(n),
			usedStr, capacityStr, util.SdkStatusToPrettyString(n.GetStatus()),
			kernelVersionStr, osFlavorStr)
	}
	t.Print()

	return nil
}
