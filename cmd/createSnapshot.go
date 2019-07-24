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
	Use:   "volumesnapshot [VOLUME] [NAME]",
	Short: "Create a volume snapshot",
	Long:  `Create a snapshot for the specified volume`,
	Example: `$ px create volumesnapshot mysnap --labels color=blue,fabric=wool --volume myvol
This creates a snapshot named mysnap for the specified volume myvol.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("Must supply the volume to snap and a new name for the snapshot")
		}
		return nil
	},
	RunE: createSnapshotExec,
}

func init() {
	createCmd.AddCommand(createSnapshotCmd)
	createSnapshotCmd.Flags().StringVar(&csOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
}

func createSnapshotExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get name
	csOpts.req.VolumeId = args[0]
	csOpts.req.Name = args[1]

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

	// Show user information
	msg := fmt.Sprintf("Snapshot of %s created with id %s",
		csOpts.req.GetVolumeId(),
		resp.GetSnapshotId())

	output, _ := cmd.Flags().GetString("output")
	formattedOut := &util.DefaultFormatOutput{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},

		Cmd:  "create snapshot",
		Desc: msg,
		Id:   []string{resp.GetSnapshotId()},
	}
	formattedOut.Print()

	return nil
}
