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

// deleteVolumeCmd represents the deleteVolume command
var deleteVolumeCmd = &cobra.Command{
	Use:     "volume [NAME]",
	Short:   "Delete a volume in Portworx",
	Example: "$ px delete volume myvolume",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Must supply a volume name")
		}
		return nil
	},
	RunE: deleteVolumeExec,
}

func init() {
	deleteCmd.AddCommand(deleteVolumeCmd)
}

func deleteVolumeExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	name := args[0]
	_, err = volumes.Delete(ctx, &api.SdkVolumeDeleteRequest{
		VolumeId: name,
	})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to delete volume")
	}

	msg := fmt.Sprintf("Volume %s deleted", name)

	output, _ := cmd.Flags().GetString("output")

	formattedOut := &util.DefaultFormatOutput{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		Cmd:  "delete volume",
		Desc: msg,
		Id:   []string{name},
	}
	formattedOut.Print()

	return nil
}
