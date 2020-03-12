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
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	pxc "github.com/portworx/pxc/pkg/component"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

var clusterUuidCmd *cobra.Command

// Initialize variables and create the command
var _ = commander.RegisterCommandVar(func() {
	// clusterUuidCmd represents the clusterUuid command
	clusterUuidCmd = &cobra.Command{
		Use:   "uuid",
		Short: "Show cluster uuid",
		RunE:  clusterUuidExec,
	}
})

// Register this command
var _ = commander.RegisterCommandInit(func() {
	pxc.RootAddCommand(clusterUuidCmd)
})

// UuidAddCommand is used to install sub commands
func UuidAddCommand(c *cobra.Command) {
	clusterUuidCmd.AddCommand(c)
}

func clusterUuidExec(c *cobra.Command, args []string) error {

	// This will automatically setup the the connection to portworx
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Print cluster information
	cluster := api.NewOpenStorageClusterClient(conn)
	clusterInfo, err := cluster.InspectCurrent(ctx, &api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to inspect cluster")
	}
	util.Printf("%s\n", clusterInfo.GetCluster().GetId())

	return nil
}
