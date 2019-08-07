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
package cmd

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/util"

	"github.com/spf13/cobra"
)

type createVolumeOpts struct {
	req                *api.SdkVolumeCreateRequest
	labelsAsString     string
	sizeInGi           int
	filesystemAsString string
}

var (
	cvOpts          *createVolumeOpts
	createVolumeCmd *cobra.Command
)

var _ = RegisterCommandVar(func() {
	// createVolumeCmd represents the createVolume command
	cvOpts = &createVolumeOpts{
		req: &api.SdkVolumeCreateRequest{
			Spec: &api.VolumeSpec{},
		},
	}
	createVolumeCmd = &cobra.Command{
		Use:   "volume [NAME]",
		Short: "Create a volume in Portworx",

		// TODO:
		Example: `1. Create volume called "myvolume" with size as 3GiB:
	$ px create volume myvolume --size=3
2. Create volume called "myvolume" with size as 3GiB and replica set to 3:
	$ px create volume myvolume --size=3 --replicas=3
3. Create shared volume called "myvolume" with size as 3GiB:
	$ px create volume myvolume --size=3 --shared
4. Create shared volume called "myvolume" with size as 2GiB and replicas set to 3:
	$ px create volume myvolume --size=3 --shared --replicas=3
5. Create volume called "myvolume" with label as "access=slow" and size as 3 GiB:
	$ px create volume myvolume --size=3 --labels 'access=slow'`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for volume")
			}
			return nil
		},
		RunE: createVolumeExec,
	}
})

var _ = RegisterCommandInit(func() {
	createCmd.AddCommand(createVolumeCmd)

	createVolumeCmd.Flags().IntVar(&cvOpts.sizeInGi, "size", 0, "Size in GiB")
	createVolumeCmd.Flags().Int64Var(&cvOpts.req.Spec.HaLevel, "replicas", 1, "Number of replicas also called HA level [1-3]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Shared, "shared", false, "Shared volume")
	createVolumeCmd.Flags().StringVar(&cvOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
	createVolumeCmd.Flags().StringVar(&cvOpts.filesystemAsString, "fs", "ext4", "Filesystem type for the volume [none, ext4]")
	createVolumeCmd.Flags().SortFlags = false

	// TODO bring the flags from rootCmd

	// TODO add more flags here
})

func createVolumeExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
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

	// Convert size to bytes in uint64
	cvOpts.req.Spec.Size = uint64(cvOpts.sizeInGi) * uint64(Gi)

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
