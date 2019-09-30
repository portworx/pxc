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
	"errors"
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type createVolumeOpts struct {
	req                *api.SdkVolumeCreateRequest
	labelsAsString     string
	sizeInGi           int
	filesystemAsString string
	replicaSet         []string
	IoProfile          string
	groups             string
	collaborators      string
	earlyAck           bool
	asyncIo            bool
	passPhrase         string
}

var (
	cvOpts          *createVolumeOpts
	createVolumeCmd *cobra.Command
)

func CreateAddCommand(c *cobra.Command) {
	createVolumeCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	// createVolumeCmd represents the createVolume command
	cvOpts = &createVolumeOpts{
		req: &api.SdkVolumeCreateRequest{
			Spec: &api.VolumeSpec{},
		},
	}
	createVolumeCmd = &cobra.Command{
		Use:   "volume [NAME]",
		Short: "Create a volume in Portworx",

		Example: `
  # To create volume called "myvolume" with size as 3GiB:
  pxc create volume myvolume --size=3
  # To create volume called "myvolume" with size as 3GiB and replicas set to 3:
  pxc create volume myvolume --size=3 --replicas=3
  # To create shared volume called "myvolume" with size as 3GiB:
  pxc create volume myvolume --size=3 --shared
  # To create shared volume called "myvolume" with size as 2GiB and replicas set to 3:
  pxc create volume myvolume --size=3 --shared --replicas=3
  # To create volume called "myvolume" with label as "access=slow" and size as 3 GiB:
  pxc create volume myvolume --size=3 --labels 'access=slow'
  # To create volume called 'myvolume" with volume access (collaborators and groups) option flag.
  # r - read, w - write, a -admin:
  pxc create volume myvolume --size=3 --groups group1:r,group2:w,group3:a --collaborators user1:r,user2:a,user3:w`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for volume")
			}
			return nil
		},
		RunE: createVolumeExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.CreateAddCommand(createVolumeCmd)

	createVolumeCmd.Flags().IntVar(&cvOpts.sizeInGi, "size", 0, "Size in GiB")
	createVolumeCmd.Flags().Int64Var(&cvOpts.req.Spec.HaLevel, "replicas", 1, "Number of replicas also called HA level [1-3]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Shared, "shared", false, "Shared volume")
	createVolumeCmd.Flags().StringVar(&cvOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
	createVolumeCmd.Flags().StringVar(&cvOpts.filesystemAsString, "fs", "ext4", "Filesystem type for the volume [none, ext4]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Sticky, "sticky", false, "Sitcky volume")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Journal, "journal", false, "Journal data for this volume")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Encrypted, "encryption", false, "encrypt this volume")
	createVolumeCmd.Flags().Uint32Var(&cvOpts.req.Spec.AggregationLevel, "aggregation-level", 0, "aggregation level (Valid Values: [1, 2, 3] (default 1)")
	createVolumeCmd.Flags().StringSliceVar(&cvOpts.replicaSet, "nodes", []string{}, "Replicat set nodes for this volume")
	createVolumeCmd.Flags().StringVar(&cvOpts.IoProfile, "ioprofile", "", "IO Profile (Valid Values: [sequential cms db db_remote sync_shared]) (default sequential)")
	createVolumeCmd.Flags().StringVar(&cvOpts.groups, "groups", "", "list of group with volume access details, 'group1:r, group2:w'")
	createVolumeCmd.Flags().StringVar(&cvOpts.collaborators, "collaborators", "", "list of collaborators with volume access details, 'user1:r, user2:w'")
	createVolumeCmd.Flags().Uint32Var(&cvOpts.req.Spec.QueueDepth, "queue-depth", 128, "block device queue depth (Valid Range: [1 256]) (default 128)")
	createVolumeCmd.Flags().BoolVar(&cvOpts.earlyAck, "early-ack", false, "Reply to async write requests after it is copied to shared memory")
	createVolumeCmd.Flags().BoolVar(&cvOpts.asyncIo, "async-io", false, "Enable async IO to backing storage")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Nodiscard, "nodiscard", false, "Disable discard support for this volume")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.GroupEnforced, "group-enforced", false, "Enforce group during provision")
	createVolumeCmd.Flags().Uint32Var(&cvOpts.req.Spec.Scale, "scale", 0, "auto scale to max number (Valid Range: [1 1024]) (default 1)")
	createVolumeCmd.Flags().StringVar(&cvOpts.passPhrase, "passphrase", "", "Passphrase for an encrypted volume")
	createVolumeCmd.Flags().Uint32Var(&cvOpts.req.Spec.SnapshotInterval, "snapshot-interval", 0, "SnapshotInterval in minutes, set to 0 to disable snapshots")
	createVolumeCmd.Flags().SortFlags = false
})

func createVolumeExec(c *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get name
	cvOpts.req.Name = args[0]

	// Get labels
	if len(cvOpts.labelsAsString) != 0 {
		var err error
		cvOpts.req.Labels, err = util.CommaStringToStringMap(cvOpts.labelsAsString)
		if err != nil {
			return fmt.Errorf("Failed to parse labels: %v\n", err)
		}
	}

	if len(cvOpts.collaborators) != 0 || len(cvOpts.groups) != 0 {
		cvOpts.req.Spec.Ownership = &api.Ownership{}
		cvOpts.req.Spec.Ownership.Acls = &api.Ownership_AccessControl{}
	}
	// Get collaborators
	if len(cvOpts.collaborators) != 0 {
		collaborators, err := util.GetAclMapFromString(cvOpts.collaborators)
		if err != nil {
			return err
		}
		cvOpts.req.Spec.Ownership.Acls.Collaborators = collaborators
	}

	// Get groups
	if len(cvOpts.groups) != 0 {
		groups, err := util.GetAclMapFromString(cvOpts.groups)
		if err != nil {
			return err
		}
		cvOpts.req.Spec.Ownership.Acls.Groups = groups
	}

	// Convert size to bytes in uint64
	cvOpts.req.Spec.Size = uint64(cvOpts.sizeInGi) * uint64(cmd.Gi)

	// Add fs to request
	switch {
	case cvOpts.filesystemAsString == "ext4":
		cvOpts.req.Spec.Format = api.FSType_FS_TYPE_EXT4
	case cvOpts.filesystemAsString == "none":
		cvOpts.req.Spec.Format = api.FSType_FS_TYPE_NONE
	default:
		return fmt.Errorf("Error: --fs valid values are [none, ext4]\n")
	}

	// Update default EXT4, if it fs is 'none' and shared volume
	if cvOpts.req.Spec.Format == api.FSType_FS_TYPE_NONE && cvOpts.req.Spec.Shared {
		cvOpts.req.Spec.Format = api.FSType_FS_TYPE_EXT4
	}

	// setting replica set nodes if provided
	if len(cvOpts.replicaSet) != 0 {
		cvOpts.req.Spec.ReplicaSet = &api.ReplicaSet{
			Nodes: cvOpts.replicaSet,
		}
	}

	// setting IO profile if provided.
	if len(cvOpts.IoProfile) > 0 {
		switch cvOpts.IoProfile {
		case "db":
			cvOpts.req.Spec.IoProfile = api.IoProfile_IO_PROFILE_DB
		case "cms":
			cvOpts.req.Spec.IoProfile = api.IoProfile_IO_PROFILE_CMS
		case "db_remote":
			cvOpts.req.Spec.IoProfile = api.IoProfile_IO_PROFILE_DB_REMOTE
		case "sync_shared":
			cvOpts.req.Spec.IoProfile = api.IoProfile_IO_PROFILE_SYNC_SHARED
		default:
			flagError := errors.New("Invalid IO profile")
			return flagError
		}
	}

	if cvOpts.req.Spec.QueueDepth < 1 || cvOpts.req.Spec.QueueDepth > 256 {
		return fmt.Errorf("Queuedepth has to be in the range of 1 to 256.")
	}

	// Setting Iostrategy
	ioStrategy := &api.IoStrategy{
		AsyncIo:  cvOpts.asyncIo,
		EarlyAck: cvOpts.earlyAck,
	}
	cvOpts.req.Spec.IoStrategy = ioStrategy

	// Setting passphrase
	if len(cvOpts.passPhrase) != 0 {
		cvOpts.req.Spec.Passphrase = cvOpts.passPhrase
	}

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Create(ctx, cvOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create volume")
	}

	// Show user information
	msg := fmt.Sprintf("Volume %s created with id %s\n",
		cvOpts.req.GetName(),
		resp.GetVolumeId())

	formattedOut := &util.DefaultFormatOutput{
		Cmd:  "create volume",
		Desc: msg,
		Id:   []string{resp.GetVolumeId()},
	}
	util.PrintFormatted(formattedOut)
	return nil
}
