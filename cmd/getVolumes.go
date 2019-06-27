/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"

	"github.com/cheynewallace/tabby"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// getVolumesCmd represents the getVolumes command
var getVolumesCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"volumes", "vol"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getVolumesExec(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getVolumesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getVolumesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getVolumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getVolumesCmd.Flags().String("owner", "", "Owner of volume")
	getVolumesCmd.Flags().String("volumegroup", "", "Volume group id")
	getVolumesCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")

}

func getVolumesExec(cmd *cobra.Command, args []string) {
	ctx, conn := pxConnect()
	defer conn.Close()

	// Get volume information
	volumes := api.NewOpenStorageVolumeClient(conn)
	var vols []*api.SdkVolumeInspectResponse
	if len(args) != 0 {
		vols = make([]*api.SdkVolumeInspectResponse, 0, len(args))
		for _, v := range args {
			// If it is just one volume, just do an inspect
			vol, err := volumes.Inspect(ctx, &api.SdkVolumeInspectRequest{VolumeId: v})
			if err != nil {
				pxPrintGrpcErrorWithMessage(err, "Failed to get volume")
				return
			}
			vols = append(vols, vol)
		}
	} else {
		// If it is no volumes (all)
		volsInfo, err := volumes.InspectWithFilters(ctx, &api.SdkVolumeInspectWithFiltersRequest{})
		if err != nil {
			pxPrintGrpcErrorWithMessage(err, "Failed to get volumes")
			return
		}
		vols = volsInfo.GetVolumes()
	}

	// Get output
	output, _ := cmd.Flags().GetString("output")
	switch output {
	case "yaml":
		getVolumesYamlPrinter(cmd, args, vols)
	case "json":
		getVolumesJsonPrinter(cmd, args, vols)
	case "wide":
		// We can have a special one here, but for simplicity, we will use the
		// default printer
		fallthrough
	default:
		getVolumesDefaultPrinter(cmd, args, ctx, conn, vols)
	}
}

func getVolumesYamlPrinter(cmd *cobra.Command, args []string, vols []*api.SdkVolumeInspectResponse) {
	bytes, err := yaml.Marshal(vols)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create yaml output")
		return
	}
	fmt.Println(string(bytes))
}

func getVolumesJsonPrinter(cmd *cobra.Command, args []string, vols []*api.SdkVolumeInspectResponse) {
	bytes, err := json.MarshalIndent(vols, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create json output")
		return
	}
	fmt.Println(string(bytes))
}

func getVolumesDefaultPrinter(cmd *cobra.Command, args []string, ctx context.Context, conn *grpc.ClientConn, vols []*api.SdkVolumeInspectResponse) {

	// Determine if it is a wide output
	output, _ := cmd.Flags().GetString("output")
	wide := output == "wide"

	// Determine if we need to show labels
	showLabels, _ := cmd.Flags().GetBool("show-labels")

	// Start the columns
	t := tabby.New()
	np := &volumeColumnFormatter{
		wide:       wide,
		showLabels: showLabels,
		ctx:        ctx,
		conn:       conn,
	}
	t.AddHeader(np.getHeader()...)

	for _, n := range vols {
		t.AddLine(np.getLine(n)...)
	}
	t.Print()
}

type volumeColumnFormatter struct {
	wide       bool
	showLabels bool
	ctx        context.Context
	conn       *grpc.ClientConn
}

func (p *volumeColumnFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.wide {
		header = []interface{}{"Id", "Name", "Size Gi", "HA", "Shared", "Encrypted", "Io Profile", "Status", "State", "Snap Enabled"}
	} else {
		header = []interface{}{"Name", "Size", "HA", "Shared", "Status", "State"}
	}
	if p.showLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *volumeColumnFormatter) getLine(resp *api.SdkVolumeInspectResponse) []interface{} {

	v := resp.GetVolume()
	spec := v.GetSpec()

	var node *api.StorageNode
	if len(v.GetAttachedOn()) != 0 {
		nodes := api.NewOpenStorageNodeClient(p.conn)
		nodeInfo, err := nodes.Inspect(p.ctx, &api.SdkNodeInspectRequest{NodeId: v.GetAttachedOn()})
		if err != nil {
			pxPrintGrpcErrorWithMessage(err, "Failed to get node information where volume attached")
			return nil
		}
		node = nodeInfo.GetNode()
	}

	// Determine the status of the volume
	state := "Detached"
	if v.State == api.VolumeState_VOLUME_STATE_ATTACHED {
		if node != nil {
			state = "on " + node.GetHostname()
		} else {
			state = "Attached"
		}
	} else if v.State == api.VolumeState_VOLUME_STATE_DETATCHING {
		if node != nil {
			state = "Was on " + node.GetHostname()
		} else {
			state = "Detaching"
		}
	}

	// Size needs to be done better
	var line []interface{}
	if p.wide {
		line = []interface{}{
			v.GetId(), v.GetLocator().GetName(), spec.GetSize() / Gi, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), spec.GetEncrypted(),
			spec.GetCos(), prettyStatus(v), state, spec.GetSnapshotSchedule() != "",
		}
	} else {
		line = []interface{}{
			v.GetLocator().GetName(), spec.GetSize() / Gi, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), prettyStatus(v), state,
		}
	}
	if p.showLabels {
		line = append(line, labelsToString(v.GetLocator().GetVolumeLabels()))
	}
	return line
}

func prettyStatus(v *api.Volume) string {
	return strings.TrimPrefix(v.GetStatus().String(), "VOLUME_STATUS_")
}
