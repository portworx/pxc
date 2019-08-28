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
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type volumeUpdateOpts struct {
	req                    *api.SdkVolumeUpdateRequest
	halevel                int64
	replicaSet             []string
	size                   uint64
	shared                 string
	sticky                 string
	addCollaborators       string
	addGroups              string
	removeCollaborators    string
	removeGroups           string
	removeAllCollaborators bool
	removeAllGroups        bool
}

// Struct that contains flag to track the options that set in the current context.
// This struct variables will be used to check the invalid combination of the flag options.
type volumeUpdateOptionStatus struct {
	haLevelSet, sizeSet, sharedSet, stickySet                              bool
	addGroupsSet, removeGroupsSet, removeAllGroupsSet                      bool
	addCollaboratorsSet, removeCollaboratorsSet, removeAllCollaboratorsSet bool
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
	patchVolumeCmd.Flags().StringVar(&updateReq.addCollaborators, "add-collaborators", "", "Add list of collaborators to the existing list")
	patchVolumeCmd.Flags().StringVar(&updateReq.addGroups, "add-groups", "", "Add list of groups to the existing list")
	patchVolumeCmd.Flags().StringVar(&updateReq.removeCollaborators, "remove-collaborators", "", "Remove the given users from the collaborators list")
	patchVolumeCmd.Flags().StringVar(&updateReq.removeGroups, "remove-groups", "", "Remove the given groups from the group list")
	patchVolumeCmd.Flags().BoolVarP(&updateReq.removeAllCollaborators, "remove-all-collaborators", "", false, "Remove all the user from the collaborators list")
	patchVolumeCmd.Flags().BoolVarP(&updateReq.removeAllGroups, "remove-all-groups", "", false, "Remove all the groups from the group list")
	patchVolumeCmd.Flags().SortFlags = false
})

// validateVolumeUpdateOptions will check for the valid combination of the flag option.
func validateVolumeUpdateOptions(status volumeUpdateOptionStatus) error {
	if status.haLevelSet {
		if status.sizeSet || status.sharedSet || status.stickySet {
			return fmt.Errorf("Error: --halevel is not a valid combination with --size or --shared or --sticky")
		}
	}
	if status.removeAllCollaboratorsSet {
		if status.addCollaboratorsSet || status.removeCollaboratorsSet {
			return fmt.Errorf("Error: remove-all-collaborators is not a valid combination with --add-collaborators or --remove-collaborators")
		}
	}
	if status.removeAllGroupsSet {
		if status.addGroupsSet || status.removeGroupsSet {
			return fmt.Errorf("Error: remove-all-groups is not a valid combination with --add-groups or --remove-groups")

		}
	}
	return nil
}

func PatchAddCommand(cmd *cobra.Command) {
	patchVolumeCmd.AddCommand(cmd)
}

func updateVolume(cmd *cobra.Command, args []string) error {

	var volumeFlagStatus volumeUpdateOptionStatus

	// Parse out all of the common cli volume flags
	cvi := cliops.GetCliVolumeInputs(cmd, args)
	// Create a CliVolumeOps object
	cvOps := cliops.NewCliVolumeOps(cvi)
	// Connect to px and k8s (if needed)
	err := cvOps.Connect()

	defer cvOps.Close()

	// fetch the volume name from args
	updateReq.req.VolumeId = args[0]

	updateReq.req.Spec = &api.VolumeSpecUpdate{}

	if len(updateReq.addCollaborators) != 0 || len(updateReq.addGroups) != 0 ||
		!updateReq.removeAllCollaborators || !updateReq.removeAllGroups ||
		len(updateReq.removeCollaborators) != 0 || len(updateReq.removeGroups) != 0 {
		// For removeAllCollaborators and removeAllGroups, Initialize a empty Ownership
		// For addCollaborators, addGroups, removeCollaborators and removeGroups,
		// Intialize a empty Ownership and update the values later.
		updateReq.req.Spec.Ownership = &api.Ownership{}
		updateReq.req.Spec.Ownership.Acls = &api.Ownership_AccessControl{}
	}

	var currentGroups map[string]api.Ownership_AccessType
	var currentCollaborators map[string]api.Ownership_AccessType

	if len(updateReq.addCollaborators) != 0 || len(updateReq.addGroups) != 0 ||
		len(updateReq.removeCollaborators) != 0 || len(updateReq.removeGroups) != 0 {
		cvOps.PxVolumeOps.GetPxVolumeOpsInfo().VolNames = make([]string, 1, 1)
		// Assign the user given volume Name
		cvOps.PxVolumeOps.GetPxVolumeOpsInfo().VolNames[0] = updateReq.req.VolumeId

		// Get the current copy of the volume spec
		vols, err := cvOps.PxVolumeOps.GetVolumes()
		if err != nil {
			return err
		}
		if len(vols) == 0 {
			return fmt.Errorf("Error: Volume: %s not found\n", updateReq.req.VolumeId)
		}
		v := vols[0].GetVolume()
		spec := v.GetSpec()

		// Read the current collaborators
		currentCollaborators = spec.GetOwnership().GetAcls().GetCollaborators()
		// Read the current Groups
		currentGroups = spec.GetOwnership().GetAcls().GetGroups()
	}

	// Update the  collaborators list
	if len(updateReq.addCollaborators) != 0 {
		newCollaborators, err := util.GetAclMapFromString(updateReq.addCollaborators)
		if err != nil {
			return err
		}
		if currentCollaborators == nil {
			// If the currentCollaborators list is empty, directly assign the new set of collaborators
			currentCollaborators = newCollaborators
		} else {
			// If the currentCollaborators list is not empty, merge the new collaborators list to the current collaborators list.
			for key, value := range newCollaborators {
				currentCollaborators[key] = value
			}
		}
		updateReq.req.Spec.Ownership.Acls.Collaborators = currentCollaborators
		updateReq.req.Spec.Ownership.Acls.Groups = currentGroups
		volumeFlagStatus.addCollaboratorsSet = true
	}

	// Update the  groups list
	if len(updateReq.addGroups) != 0 {
		newGroups, err := util.GetAclMapFromString(updateReq.addGroups)
		if err != nil {
			return err
		}
		if currentGroups == nil {
			// If the currentGroups list is empty, directly assign the new set of Groups
			currentGroups = newGroups
		} else {
			// If the currentGroups list is not empty, merge the new Groups list to the current Groups list.
			for key, value := range newGroups {
				currentGroups[key] = value
			}
		}
		updateReq.req.Spec.Ownership.Acls.Groups = currentGroups
		updateReq.req.Spec.Ownership.Acls.Collaborators = currentCollaborators
		volumeFlagStatus.addGroupsSet = true
	}

	// Remove the given list of collaborators
	if len(updateReq.removeCollaborators) != 0 {
		removeCollaborators, err := util.GetAclMapFromString(updateReq.removeCollaborators)
		if err != nil {
			return nil
		}
		for key, _ := range removeCollaborators {
			delete(currentCollaborators, key)
		}
		updateReq.req.Spec.Ownership.Acls.Collaborators = currentCollaborators
		updateReq.req.Spec.Ownership.Acls.Groups = currentGroups
		volumeFlagStatus.removeCollaboratorsSet = true
	}

	// Remove the given list of Groups
	if len(updateReq.removeGroups) != 0 {
		removeGroups, err := util.GetAclMapFromString(updateReq.removeGroups)
		if err != nil {
			return nil
		}
		for key, _ := range removeGroups {
			delete(currentGroups, key)
		}
		updateReq.req.Spec.Ownership.Acls.Groups = currentGroups
		updateReq.req.Spec.Ownership.Acls.Collaborators = currentCollaborators
		volumeFlagStatus.removeGroupsSet = true
	}

	//Remove All the groups
	if updateReq.removeAllCollaborators {
		volumeFlagStatus.removeAllCollaboratorsSet = true
	}

	// Remove all the groups
	if updateReq.removeAllGroups {
		volumeFlagStatus.removeAllGroupsSet = true
	}

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
		volumeFlagStatus.haLevelSet = true
	}

	// check prvoide size is valid
	if updateReq.size > 0 {
		//Provided size has to be converted to bytes
		updateReq.req.Spec.SizeOpt = &api.VolumeSpecUpdate_Size{
			Size: (updateReq.size * 1024 * 1024 * 1024),
		}
		volumeFlagStatus.sizeSet = true
	}

	// For setting volume as shared or not
	switch updateReq.shared {
	case "on":
		updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
			Shared: true,
		}
		volumeFlagStatus.sharedSet = true
	case "off":
		updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
			Shared: false,
		}
		volumeFlagStatus.sharedSet = true
	}

	// For setting volume to be sticky
	switch updateReq.sticky {
	case "on":
		updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
			Sticky: true,
		}
		volumeFlagStatus.stickySet = true
	case "off":
		updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
			Sticky: false,
		}
		volumeFlagStatus.stickySet = true
	}

	if !volumeFlagStatus.addCollaboratorsSet && !volumeFlagStatus.addGroupsSet &&
		!volumeFlagStatus.haLevelSet && !volumeFlagStatus.removeAllCollaboratorsSet &&
		!volumeFlagStatus.removeAllGroupsSet && !volumeFlagStatus.removeCollaboratorsSet &&
		!volumeFlagStatus.removeGroupsSet && !volumeFlagStatus.sharedSet &&
		!volumeFlagStatus.sizeSet && !volumeFlagStatus.stickySet {
		return fmt.Errorf("Error: Must supply any one of the flags with valid parameters. " +
			"Please see help for more info")
	}

	// Check whether the flag options are in valid combination.
	err = validateVolumeUpdateOptions(volumeFlagStatus)
	if err != nil {
		return err
	}

	volumes := api.NewOpenStorageVolumeClient(cvOps.PxVolumeOps.GetPxVolumeOpsInfo().PxConnectionData.Conn)
	_, err = volumes.Update(cvOps.PxVolumeOps.GetPxVolumeOpsInfo().PxConnectionData.Ctx, updateReq.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to patch volume")
	}
	util.Printf("Volume %s parameter updated successfully\n", updateReq.req.VolumeId)
	return nil
}
