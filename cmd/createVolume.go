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
	"os"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createVolumeExec(cmd, args)
	},
}

func init() {
	createCmd.AddCommand(createVolumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createVolumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createVolumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createVolumeCmd.Flags().StringVar(&cvOpts.req.Name, "name", "", "Name of volume (required)")
	createVolumeCmd.Flags().IntVar(&cvOpts.sizeInGi, "size", 0, "Size in GiB")
	createVolumeCmd.Flags().Int64Var(&cvOpts.req.Spec.HaLevel, "replicas", 0, "Number of replicas [1-3]")
	createVolumeCmd.Flags().BoolVar(&cvOpts.req.Spec.Shared, "shared", false, "Number of replicas [1-3]")
	createVolumeCmd.Flags().StringVar(&cvOpts.labelsAsString, "labels", "", "Comma separated list of labels as key-value pairs: 'k1=v1,k2=v2'")
}

func createVolumeExec(cmd *cobra.Command, args []string) {
	ctx, conn := pxConnect()
	defer conn.Close()

	if len(cvOpts.labelsAsString) != 0 {
		var err error
		cvOpts.req.Labels, err = commaKVStringToMap(cvOpts.labelsAsString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse labels: %v\n", err)
			return
		}
	}

	// Convert size to bytes in uint64
	cvOpts.req.Spec.Size = uint64(cvOpts.sizeInGi) * uint64(Gi)

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	resp, err := volumes.Create(ctx, cvOpts.req)
	if err != nil {
		pxPrintGrpcErrorWithMessage(err, "Failed to create volume")
		return
	}
	fmt.Printf("Volume %s created with id %s\n",
		cvOpts.req.GetName(),
		resp.GetVolumeId())
}
