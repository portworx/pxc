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
	"os"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	sched "github.com/portworx/pxc/pkg/openstorage/sched"
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
	periodic           string
	daily              []string
	weekly             []string
	monthly            []string
	policy             string
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
		Use:   "create [NAME]",
		Short: "Create a volume in Portworx",
		Example: `
  # Create a volume called "myvolume" with size as 3GiB:
  pxc volume create myvolume --size=3

  # Create a volume called "myvolume" with size as 3GiB and replicas set to 3:
  pxc volume create myvolume --size=3 --replicas=3

  # Create a shared volume called "myvolume" with size as 3GiB:
  pxc volume create myvolume --size=3 --shared

  # Create a shared volume called "myvolume" with size as 2GiB and replicas set to 3:
  pxc volume create myvolume --size=3 --shared --replicas=3

  # Create a volume called "myvolume" with label as "access=slow" and size as 3 GiB:
  pxc volume create myvolume --size=3 --labels 'access=slow'

  # Create a volume called 'myvolume" with volume access (collaborators and groups) option flag.
  # r - read, w - write, a -admin:
  pxc volume create myvolume --size=3 --groups group1:r,group2:w,group3:a --collaborators user1:r,user2:a,user3:w

  # Create a volume with periodic snapshot policy for every 15 minutes with retain=2 (maintaing two snapshot copies at a given time):
  pxc volume create snapvol --periodic 15,2

  # Create a volume with daily snapshot policy at 00h:10m with retain=2 (maintaing two snapshot copies at a given time):
  pxc volume create snapvol --daily 00:10,2

  # Create a volume with weekly snapshot for every monday at 00h:12m with retain=2 (maintaing two snapshot copies at a given time):
  pxc volume create snapvol --weekly monday@00:12,2

  # Create a volume with monthly snapshot on 25th of every month at 10h:10m with retain=2 (maintaing two snapshot copies at a given time):
  pxc volume create snapvol --monthly 25@10:10,2`,

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
	VolumeAddCommand(createVolumeCmd)

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
	createVolumeCmd.Flags().StringVar(&cvOpts.periodic, "periodic", "", "periodic snapshot interval in mins,k (keeps 5 by default), 0 disables all schedule snapshots")
	createVolumeCmd.Flags().StringSliceVar(&cvOpts.daily, "daily", []string{}, "daily snapshot at specified hh:mm,k (keeps 7 by default)")
	createVolumeCmd.Flags().StringSliceVar(&cvOpts.weekly, "weekly", []string{}, "weekly snapshot at specified weekday@hh:mm,k (keeps 5 by default)")
	createVolumeCmd.Flags().StringSliceVar(&cvOpts.monthly, "monthly", []string{}, "monthly snapshot at specified day@hh:mm,k (keeps 12 by default)")
	createVolumeCmd.Flags().StringVar(&cvOpts.policy, "policy", "", "Schedule policy names separated by comma")

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
	// Checking if snapshot schedule is set
	schedule, snapErr := makeSnapSchedule(cvOpts)
	if snapErr != nil {
		return snapErr
	}
	cvOpts.req.Spec.SnapshotSchedule = schedule

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Create(ctx, cvOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed Create a volume")
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

func makeSnapSchedule(snapOpts *createVolumeOpts) (string, error) {
	updates := []sched.RetainIntervalSpec{}
	var err error
	if len(snapOpts.periodic) > 0 {
		s, err := sched.ParsePeriodic(snapOpts.periodic)
		if err != nil {
			return "", err
		}
		updates = append(updates, s)
		if s.Period == 0 {
			return sched.ScheduleString(updates, nil)
		}
	}

	for freq, parse := range sched.ParseCLI {
		var items []string
		switch freq {
		case sched.DailyType:
			items = snapOpts.daily
		case sched.WeeklyType:
			items = snapOpts.weekly
		case sched.MonthlyType:
			items = snapOpts.monthly
		default:
			return "", fmt.Errorf("unknown periodicity")
		}

		// fix items if they have been split during CLI parsing due to comma in the format string.
		items = util.FixCommaBasedStringSliceInput(items, os.Args)

		for _, item := range items {
			var s sched.RetainIntervalSpec
			s, err = parse(item)
			if err != nil {
				break
			}
			updates = append(updates, s)
		}
	}

	if err != nil {
		return "", err
	}
	p, err := sched.NewPolicyTags(snapOpts.policy)
	if err != nil {
		return "", err
	}
	return sched.ScheduleString(updates, p)
}
