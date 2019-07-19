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
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

var (
	ccReq = &api.SdkVolumeCloneRequest{}
)

// createCloneCmd represents the createClone command
var createCloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a volume clone",
	Long: `Create a clone for the specified volume.
Example: px create clone --name myclone --volume myvol'
This creates a clone named myclone for the specified volume myvol.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return createCloneExec(cmd, args)
	},
}

func init() {
	createCmd.AddCommand(createCloneCmd)

	createCloneCmd.Flags().StringVar(&ccReq.Name, "name", "", "Name of clone to be created (required)")
	createCloneCmd.Flags().StringVar(&ccReq.ParentId, "volume", "", "Name or id of volume (required)")
	createCloneCmd.Flags().SortFlags = false

	// TODO bring the flags from rootCmd
}

func createCloneExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectCurrent(GetConfigFile())
	if err != nil {
		return err
	}
	defer conn.Close()

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
		Cmd:  "create clone",
		Desc: msg,
		Id:   []string{resp.GetVolumeId()},
	}
	formattedOut.SetFormat(output)
	util.Printf("%v\n", formattedOut)
	return nil
}
