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
	req            *api.SdkVolumeCreateRequest
	labelsAsString string
	sizeInGi       int
}

var (
	cvOpts = &createVolumeOpts{
		req: &api.SdkVolumeCreateRequest{
			Spec: &api.VolumeSpec{},
		},
	}
)

// createVolumeCmd represents the createVolume command
var createVolumeCmd = &cobra.Command{
	Use:   "volume [NAME]",
	Short: "Create a volume in Portworx",

	// TODO:
	Example: `$ px create volume myvolume --size=3
This creates a volume called 'myvolume' of 3Gi.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Must supply a name for volume")
		}
		return nil
	},
	RunE: createVolumeExec,
}

func init() {
	createCmd.AddCommand(createVolumeCmd)

	createVolumeCmd.Flags().IntVar(&cvOpts.sizeInGi, "size", 0, "Size in GiB")
	createVolumeCmd.Flags().Int64Var(&cvOpts.req.Spec.HaLevel, "replicas", 0, "Number of replicas also called HA level [1-3]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Shared, "shared", false, "Shared volume")
	createVolumeCmd.Flags().StringVar(&cvOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
	createVolumeCmd.Flags().SortFlags = false

	// TODO bring the flags from rootCmd

	// TODO add more flags here
}

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

	// Set a default value
	if cvOpts.req.Spec.HaLevel == 0 {
		cvOpts.req.Spec.HaLevel = 1
	}

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Create(ctx, cvOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create volume")
	}

	// Show user information
	msg := fmt.Sprintf("Volume %s created with id %s",
		cvOpts.req.GetName(),
		resp.GetVolumeId())

	output, _ := cmd.Flags().GetString("output")
	formattedOut := &util.DefaultFormatOutput{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		Cmd:  "create volume",
		Desc: msg,
		Id:   []string{resp.GetVolumeId()},
	}
	util.PrintFormatted(formattedOut)
	return nil
}
