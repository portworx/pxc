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

package cmd

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

// createCloneCmd represents the createClone command
var createCloneCmd = &cobra.Command{
	Use:   "volumeclone [VOLUME] [NAME]",
	Short: "Creates a new volume from a volume or snapshot",
	Long:  `Create a clone for the specified volume`,
	Example: `$ px create volumeclone oldvolume newvolume
This creates a new volume called 'newvolume' from an existing volume called 'oldvolume'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("Must supply the volume to clone and a new name for the clone")
		}
		return nil
	},
	RunE: createCloneExec,
}

func init() {
	createCmd.AddCommand(createCloneCmd)
}

func createCloneExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get args
	volume := args[0]
	name := args[1]
	ccReq := &api.SdkVolumeCloneRequest{
		Name:     name,
		ParentId: volume,
	}

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Clone(ctx, ccReq)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create clone")
	}

	msg := fmt.Sprintf("Clone of %s created with id %s",
		ccReq.GetParentId(),
		resp.GetVolumeId())

	output, _ := cmd.Flags().GetString("output")
	formattedOut := &util.DefaultFormatOutput{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		Cmd:  "create clone",
		Desc: msg,
		Id:   []string{resp.GetVolumeId()},
	}
	return util.PrintFormatted(formattedOut)
}
