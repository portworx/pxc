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
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type volumeUpdateOpts struct {
	req                    *api.SdkVolumeUpdateRequest
	halevel                int64
	replicaSet             []string
	size                   uint64
	shared                 string
	sharedv4               string
	sticky                 string
	addCollaborators       string
	addGroups              string
	removeCollaborators    string
	removeGroups           string
	removeAllCollaborators bool
	removeAllGroups        bool
	earlyAck               string
	asyncIo                string
	ioProfile              string
	noDiscard              string
}

var (
	updateReq      *volumeUpdateOpts
	patchVolumeCmd *cobra.Command
	cliOps         cliops.CliOps
)

// updateVolumeCmd represents the updateVolume command
var _ = commander.RegisterCommandVar(func() {
	updateReq = &volumeUpdateOpts{
		req: &api.SdkVolumeUpdateRequest{},
	}

	patchVolumeCmd = &cobra.Command{
		Use:     "update [NAME]",
		Aliases: []string{"patch"},
		Short:   "Update field(s) of a portworx volume",
		Example: `
  #### Update Volume Spec ####

  # Update the size of the volume to 2GiB
  pxc volume update xyz --size 2

  # Set the shared flag of the volume xyz
  pxc volume update xyz --shared=on

  #### Update Volume Access Controls ####

  # Update collaborators and groups of the volume access list
  pxc volume update xyz --add-collaborators user1:r,user2:w,user3:a --add-groups group1:r,group2:w,group3:a

  # Update collaborators and remove few groups from the volume access list
  pxc volume update xyz --add-collaborators user4:r,user5:w, --remove-groups group1:r

  # Update the access type of the existing collaborators and groups
  pxc volume update xyz --add-collaborators user1:a --add-groups group1:a`,

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
	VolumeAddCommand(patchVolumeCmd)

	patchVolumeCmd.Flags().Int64Var(&updateReq.halevel, "halevel", 0, "New replication factor (Valid Range: [1, 3]) (default 1)")
	patchVolumeCmd.Flags().StringSliceVar(&updateReq.replicaSet, "nodes", []string{}, "Desired set of nodes for the volume data")
	patchVolumeCmd.Flags().Uint64Var(&updateReq.size, "size", 0, "New size for the volume (GiB) (default 1)")
	patchVolumeCmd.Flags().StringVar(&updateReq.shared, "shared", "", "set shared setting (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.sharedv4, "sharedv4", "", "set sharedv4 setting (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.sticky, "sticky", "", "set sticky setting (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.addCollaborators, "add-collaborators", "", "Add list of collaborators to the existing list")
	patchVolumeCmd.Flags().StringVar(&updateReq.addGroups, "add-groups", "", "Add list of groups to the existing list")
	patchVolumeCmd.Flags().StringVar(&updateReq.removeCollaborators, "remove-collaborators", "", "Remove the given users from the collaborators list")
	patchVolumeCmd.Flags().StringVar(&updateReq.removeGroups, "remove-groups", "", "Remove the given groups from the group list")
	patchVolumeCmd.Flags().BoolVar(&updateReq.removeAllCollaborators, "remove-all-collaborators", false, "Remove all the user from the collaborators list")
	patchVolumeCmd.Flags().BoolVar(&updateReq.removeAllGroups, "remove-all-groups", false, "Remove all the groups from the group list")
	patchVolumeCmd.Flags().StringVar(&updateReq.earlyAck, "early-ack", "", "Reply to async write requests after it is copied to shared memory (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.asyncIo, "async-io", "", "Enable async IO to backing storage (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.ioProfile, "io-profile", "", "IO Profile (Valid Values: [sequential cms db db_remote sync_shared]) (default \"sequential\")")
	patchVolumeCmd.Flags().StringVar(&updateReq.noDiscard, "nodiscard", "", "Disable discard support for this volume (Valid Values: [on off]) (default \"off\")")
	patchVolumeCmd.Flags().SortFlags = false
})

// validateVolumeUpdateOptions will check for the valid combination of the flag option.
//func validateVolumeUpdateOptions(status volumeUpdateOptionStatus) error {
func validateVolumeUpdateOptions() error {
	haLevelSet, _ := patchVolumeCmd.Flags().GetInt64("halevel")
	sizeSet, _ := patchVolumeCmd.Flags().GetUint64("size")
	sharedSet, _ := patchVolumeCmd.Flags().GetString("shared")
	stickySet, _ := patchVolumeCmd.Flags().GetString("sticky")
	if haLevelSet > 0 {
		if sizeSet > 0 || len(sharedSet) != 0 || len(stickySet) != 0 {
			return fmt.Errorf("Error: --halevel is not a valid combination with --size or --shared or --sticky")
		}
	}

	addCollaborators, _ := patchVolumeCmd.Flags().GetString("add-collaborators")
	removeAllCollaborators, _ := patchVolumeCmd.Flags().GetBool("remove-all-collaborators")
	if removeAllCollaborators {
		if len(addCollaborators) != 0 {
			return fmt.Errorf("Error: remove-all-collaborators is not a valid combination with --add-collaborators")
		}
	}

	removeAllGroupsSet, _ := patchVolumeCmd.Flags().GetBool("remove-all-groups")
	addGroupsSet, _ := patchVolumeCmd.Flags().GetString("add-groups")
	if removeAllGroupsSet {
		if len(addGroupsSet) != 0 {
			return fmt.Errorf("Error: remove-all-groups is not a valid combination with --add-groups ")
		}
	}

	return nil
}

func PatchAddCommand(cmd *cobra.Command) {
	patchVolumeCmd.AddCommand(cmd)
}

func updateVolume(cmd *cobra.Command, args []string) error {

	changed := false
	// Parse out all of the common cli volume flags
	cvi := cliops.NewCliInputs(cmd, args)
	// Create a CliVolumeOps object
	cliOps = cliops.NewCliOps(cvi)
	// Connect to px and k8s (if needed)
	err := cliOps.Connect()

	defer cliOps.Close()

	// fetch the volume name from args
	updateReq.req.VolumeId = args[0]

	updateReq.req.Spec = &api.VolumeSpecUpdate{}
	currentGroups := make(map[string]api.Ownership_AccessType)
	currentCollaborators := make(map[string]api.Ownership_AccessType)

	// Fetch the current acls for the volume
	acls, err := readCurrentAcls()
	currentCollaborators = acls.GetCollaborators()
	currentGroups = acls.GetGroups()

	if err != nil {
		return err
	}

	if len(updateReq.addCollaborators) != 0 || len(updateReq.addGroups) != 0 ||
		len(updateReq.removeCollaborators) != 0 || len(updateReq.removeGroups) != 0 ||
		updateReq.removeAllCollaborators || updateReq.removeAllGroups {
		// For removeAllCollaborators and removeAllGroups, Initialize a empty Ownership
		// For addCollaborators, addGroups, removeCollaborators and removeGroups,
		// Intialize a empty Ownership and update the values later.
		if acls == nil {
			updateReq.req.Spec.Ownership = &api.Ownership{
				Acls: &api.Ownership_AccessControl{},
			}
		} else {
			updateReq.req.Spec.Ownership = &api.Ownership{
				Acls: acls,
			}
		}

		changed = true
	}

	// Removing all collaborators
	if updateReq.removeAllCollaborators {
		updateReq.req.Spec.Ownership.Acls.Collaborators = map[string]api.Ownership_AccessType{}
		changed = true
	}

	// Removing all groups
	if updateReq.removeAllGroups {
		updateReq.req.Spec.Ownership.Acls.Groups = map[string]api.Ownership_AccessType{}
		changed = true
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
		changed = true
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
		changed = true
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
		changed = true
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
		changed = true
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
		changed = true
	}

	// check prvoide size is valid
	if updateReq.size >= 0 {
		//Provided size has to be converted to bytes
		updateReq.req.Spec.SizeOpt = &api.VolumeSpecUpdate_Size{
			Size: (updateReq.size * 1024 * 1024 * 1024),
		}
		changed = true
	}

	// For setting volume as shared or not
	if len(updateReq.shared) != 0 {
		switch updateReq.shared {
		case "on":
			updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
				Shared: true,
			}
		case "off":
			updateReq.req.Spec.SharedOpt = &api.VolumeSpecUpdate_Shared{
				Shared: false,
			}
		default:
			return fmt.Errorf("Invalid input given for shared flag. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	// For setting volume as sharedv4 or not
	if len(updateReq.sharedv4) != 0 {
		switch updateReq.sharedv4 {
		case "on":
			updateReq.req.Spec.Sharedv4Opt = &api.VolumeSpecUpdate_Sharedv4{
				Sharedv4: true,
			}
		case "off":
			updateReq.req.Spec.Sharedv4Opt = &api.VolumeSpecUpdate_Sharedv4{
				Sharedv4: false,
			}
		default:
			return fmt.Errorf("Invalid input given for shared flag. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	// For setting volume to be sticky
	if len(updateReq.sticky) != 0 {
		switch updateReq.sticky {
		case "on":
			updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
				Sticky: true,
			}
		case "off":
			updateReq.req.Spec.StickyOpt = &api.VolumeSpecUpdate_Sticky{
				Sticky: false,
			}
		default:
			return fmt.Errorf("Invalid input given for sticky flag. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	// Setting IoStrategy
	updateReq.req.Spec.IoStrategy = &api.IoStrategy{}
	if len(updateReq.earlyAck) != 0 {
		switch updateReq.earlyAck {
		case "on":
			updateReq.req.Spec.IoStrategy.EarlyAck = true
		case "off":
			updateReq.req.Spec.IoStrategy.EarlyAck = false
		default:
			return fmt.Errorf("Invalid input given for early-ack. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	if len(updateReq.asyncIo) != 0 {
		switch updateReq.asyncIo {
		case "on":
			updateReq.req.Spec.IoStrategy.AsyncIo = true
		case "off":
			updateReq.req.Spec.IoStrategy.AsyncIo = false
		default:
			return fmt.Errorf("Invalid input given for async-io. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	// Setting IoProfile
	if len(updateReq.ioProfile) > 0 {
		switch updateReq.ioProfile {
		case "db":
			updateReq.req.Spec.IoProfileOpt = &api.VolumeSpecUpdate_IoProfile{
				IoProfile: api.IoProfile_IO_PROFILE_DB,
			}
		case "cms":
			updateReq.req.Spec.IoProfileOpt = &api.VolumeSpecUpdate_IoProfile{
				IoProfile: api.IoProfile_IO_PROFILE_CMS,
			}
		case "db_remote":
			updateReq.req.Spec.IoProfileOpt = &api.VolumeSpecUpdate_IoProfile{
				IoProfile: api.IoProfile_IO_PROFILE_DB_REMOTE,
			}
		case "sync_shared":
			updateReq.req.Spec.IoProfileOpt = &api.VolumeSpecUpdate_IoProfile{
				IoProfile: api.IoProfile_IO_PROFILE_SYNC_SHARED,
			}
		case "sequential":
			updateReq.req.Spec.IoProfileOpt = &api.VolumeSpecUpdate_IoProfile{
				IoProfile: api.IoProfile_IO_PROFILE_SEQUENTIAL,
			}
		default:
			return fmt.Errorf("Invalid input given for IoProfile. Please see help.")
		}
		changed = true
	}

	// Setting discard
	if len(updateReq.noDiscard) != 0 {
		switch updateReq.noDiscard {
		case "on":
			updateReq.req.Spec.NodiscardOpt = &api.VolumeSpecUpdate_Nodiscard{
				Nodiscard: true,
			}
		case "off":
			updateReq.req.Spec.NodiscardOpt = &api.VolumeSpecUpdate_Nodiscard{
				Nodiscard: false,
			}
		default:
			return fmt.Errorf("Invalid input given for nodiscard. Valid values are \"on\" or \"off\"")
		}
		changed = true
	}

	if !changed {
		return fmt.Errorf("Error: Must supply any one of the flags with valid parameters. " +
			"Please see help for more info")
	}

	// Check whether the flag options are in valid combination.
	err = validateVolumeUpdateOptions()
	if err != nil {
		return err
	}

	volumes := api.NewOpenStorageVolumeClient(cliOps.PxOps().GetConn())
	_, err = volumes.Update(cliOps.PxOps().GetCtx(), updateReq.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to patch volume")
	}
	util.Printf("Volume %s parameter updated successfully\n", updateReq.req.VolumeId)
	return nil
}

// Reads current acls for the given volume
func readCurrentAcls() (*api.Ownership_AccessControl, error) {
	volNames := make([]string, 1, 1)
	// Assign the user given volume Name
	volNames[0] = updateReq.req.VolumeId
	volSpec := &portworx.VolumeSpec{
		VolNames: volNames,
	}
	vo := portworx.NewVolumes(cliOps.PxOps(), volSpec)

	// Get the current copy of the volume spec
	vols, err := vo.GetVolumes()
	if err != nil {
		return nil, err
	}
	if len(vols) == 0 {
		return nil, fmt.Errorf("Error: Volume: %s not found\n", updateReq.req.VolumeId)
	}
	spec := vols[0].GetSpec()
	acls := spec.GetOwnership().GetAcls()

	return acls, nil
}
