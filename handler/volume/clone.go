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

package volume

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var volumeCloneCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	volumeCloneCmd = &cobra.Command{
		Use:   "clone [VOLUME] [NAME]",
		Short: "Creates a new volume from a volume or snapshot",
		Long:  `Create a clone for the specified volume`,
		Example: `
  # Create a new volume called 'newvolume' from an existing volume called 'oldvolume'
  pxc create volumeclone oldvolume newvolume`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("Must supply the volume to clone and a new name for the clone")
			}
			return nil
		},
		RunE: volumeCloneExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	VolumeAddCommand(volumeCloneCmd)
})

func VolumeCloneAddCommand(cmd *cobra.Command) {
	volumeCloneCmd.AddCommand(cmd)
}

func volumeCloneExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
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

	formattedOut := &util.DefaultFormatOutput{
		Cmd:  "create clone",
		Desc: msg,
		Id:   []string{resp.GetVolumeId()},
	}
	return util.PrintFormatted(formattedOut)
}
