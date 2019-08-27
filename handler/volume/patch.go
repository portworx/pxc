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
package volume

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type volumeUpdateOpts struct {
	req        *api.SdkVolumeUpdateRequest
	halevel    int64
	replicaSet []string
	size       uint64
	shared     string
	sticky     string
}

var (
	updateReq      *volumeUpdateOpts
	patchVolumeCmd *cobra.Command
)

// updateVolumeCmd represents the updateVolume command
var _ = commander.RegisterCommandVar(func() {
	updateReq = &volumeUpdateOpts{
		req: &api.SdkVolumeUpdateRequest{},
	}

	patchVolumeCmd = &cobra.Command{
		Use:   "volume [NAME]",
		Short: "Update field(s) of a portworx volume",
		Example: `$ px patch  volume test --halevel 3
		This set halevel of volume to 3.

		$px patch volume test --size 2
		Updating volume size to 2GiB

		$px patch volume test --sticky
		Updating volume to be sticky

		$px patch volume test --shared
		Updating volume to shared`,

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for the volume")
			}
			return nil
		},
		RunE: updateVolume,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.PatchAddCommand(patchVolumeCmd)

	patchVolumeCmd.Flags().Int64Var(&updateReq.halevel, "halevel", 0, "New replication factor (Valid Range: [1, 3]) (default 1)")
	patchVolumeCmd.Flags().StringSliceVarP(&updateReq.replicaSet, "nodes", "n", []string{}, "Desired set of nodes for the volume data")
	patchVolumeCmd.Flags().Uint64VarP(&updateReq.size, "size", "s", 0, "New size for the volume (GiB) (default 1)")
	patchVolumeCmd.Flags().StringVarP(&updateReq.shared, "shared", "r", "", "set shared setting (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVarP(&updateReq.sticky, "sticky", "t", "", "set sticky setting (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().SortFlags = false
})

func PatchAddCommand(cmd *cobra.Command) {
	patchVolumeCmd.AddCommand(cmd)
}

func updateVolume(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	// To track if any value of updateReq is set.
	updateReqSet := false

	defer conn.Close()
	// fetch the volume name from args
	updateReq.req.VolumeId = args[0]
	updateReq.req.Spec = &api.VolumeSpecUpdate{}

	// check if halevel providied is valid one
	if updateReq.halevel > 0 {
		updateReq.req.Spec.HaLevelOpt = &api.VolumeSpecUpdate_HaLevel{
			HaLevel: int64(updateReq.halevel),
		}
		// Replicaset needs to be passed due to know volume driver issue (pwx-9500)
		// If user provides one it will be overriden and need to have additional checks.
		updateReq.req.Spec.ReplicaSet = &api.ReplicaSet{
			Nodes: updateReq.replicaSet,
		}
		updateReqSet = true
	}

	// check prvoide size is valid
	if updateReq.size > 0 {
		//Provided size has to be converted to bytes
		updateReq.req.Spec.SizeOpt = &api.VolumeSpecUpdate_Size{
			Size: (updateReq.size * 1024 * 1024 * 1024),
		}
		updateReqSet = true
	}

	// For setting volume as shared or not
	switch updateReq.shared {
	case "on":
		updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
			Shared: true,
		}
		updateReqSet = true
	case "off":
		updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
			Shared: false,
		}
		updateReqSet = true
	}

	// for setting volume to be sticky
	switch updateReq.sticky {
	case "on":
		updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
			Sticky: true,
		}
		updateReqSet = true
	case "off":
		updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
			Sticky: false,
		}
		updateReqSet = true
	}

	if !updateReqSet {
		return util.PxErrorMessage(err, "Must supply any one of the flags with valid parameters."+
			"Please see help for more info")
	}

	// Before submitting the request, need to make sure volumespec is checked for valid flag combinations.
	err = portworx.ValidateVolumeSpec(updateReq.req.Spec)
	if err != nil {
		// validation of VolumeSpec has failed, hence return failure
		return err
	}

	volumes := api.NewOpenStorageVolumeClient(conn)
	_, err = volumes.Update(ctx, updateReq.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to patch volume")
	}
	util.Printf("Volume %s parameter updated successfully\n", updateReq.req.VolumeId)
	return nil
}
