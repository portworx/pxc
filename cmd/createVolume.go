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
	"github.com/portworx/px/pkg/portworx"
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
	Use:   "volume",
	Short: "Create a volume in Portworx",

	// TODO:
	Long: `TODO
ADD EXAMPLES HERE.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return createVolumeExec(cmd, args)
	},
}

func init() {
	createCmd.AddCommand(createVolumeCmd)

	createVolumeCmd.Flags().StringVar(&cvOpts.req.Name, "name", "", "Name of volume (required)")
	createVolumeCmd.Flags().IntVar(&cvOpts.sizeInGi, "size", 0, "Size in GiB")
	createVolumeCmd.Flags().Int64Var(&cvOpts.req.Spec.HaLevel, "replicas", 0, "Number of replicas [1-3]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Shared, "shared", false, "Number of replicas [1-3]")
	createVolumeCmd.Flags().StringVar(&cvOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")

	// TODO bring the flags from rootCmd

	// TODO add more flags here
}

func createVolumeExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnect(GetConfigFile())
	if err != nil {
		return err
	}
	defer conn.Close()

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

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Create(ctx, cvOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create volume")
	}

	// TODO: Show output in appropriate format

	// Show user information
	util.Printf("Volume %s created with id %s\n",
		cvOpts.req.GetName(),
		resp.GetVolumeId())

	return nil
}
