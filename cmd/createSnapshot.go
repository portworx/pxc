// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

type createSnapshotOpts struct {
	req            *api.SdkVolumeSnapshotCreateRequest
	labelsAsString string
}

var (
	csOpts = &createSnapshotOpts{
		req: &api.SdkVolumeSnapshotCreateRequest{},
	}
)

// createSnapshotCmd represents the createSnapshot command
var createSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create a volume snapshot",
	Long: `Create a snapshot for the specified volume.
Example: px create snapshot --name mysnap --labels 'color=blue,fabric=wool --volume myvol'
This creates a snapshot named mysnap for the specified volume myvol.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return createSnapshotExec(cmd, args)
	},
}

func init() {
	createCmd.AddCommand(createSnapshotCmd)

	createSnapshotCmd.Flags().StringVar(&csOpts.req.Name, "name", "", "Name of snapshot to be created (required)")
	createSnapshotCmd.Flags().StringVar(&csOpts.req.VolumeId, "volume", "", "Name/id of volume (required)")
	createSnapshotCmd.Flags().StringVar(&csOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
	createSnapshotCmd.Flags().SortFlags = false

	// TODO bring the flags from rootCmd
}

func createSnapshotExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get labels
	if len(csOpts.labelsAsString) != 0 {
		var err error
		csOpts.req.Labels, err = util.CommaStringToStringMap(csOpts.labelsAsString)
		if err != nil {
			return fmt.Errorf("Failed to parse labels: %v\n", err)
		}
	}

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.SnapshotCreate(ctx, csOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create snapshot")
	}

	// TODO: Show output in appropriate format

	// Show user information
	util.Printf("Snapshot of %s created with id %s\n",
		csOpts.req.GetVolumeId(),
		resp.GetSnapshotId())

	return nil
}
