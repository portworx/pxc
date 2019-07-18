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
	"strings"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/util"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getVolumesCmd represents the getVolumes command
var getVolumesCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"volumes"},
	Short:   "Get information about Portworx volumes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getVolumesExec(cmd, args)
	},
}

func init() {
	getCmd.AddCommand(getVolumesCmd)
	getVolumesCmd.Flags().String("owner", "", "Owner of volume")
	getVolumesCmd.Flags().String("volumegroup", "", "Volume group id")
	getVolumesCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")
	getVolumesCmd.Flags().Bool("show-k8s-info", false, "Show kubernetes information")

	// TODO: Place here support for selectors and move the flags from the rootCmd
}

func getVolumesExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Check if we need to get information from Kubernetes
	var pods []v1.Pod
	showK8s, _ := cmd.Flags().GetBool("show-k8s-info")
	if showK8s {
		_, kc, err := KubeConnectDefault()
		if err != nil {
			return err
		}
		podClient := kc.CoreV1().Pods("")
		podList, err := podClient.List(metav1.ListOptions{})
		if err != nil {
			return err
		}
		pods = podList.Items
	}

	// Get volume information
	volumes := api.NewOpenStorageVolumeClient(conn)
	var vols []*api.SdkVolumeInspectResponse

	// Determine if we should get all the volumes or specific ones
	if len(args) != 0 {
		vols = make([]*api.SdkVolumeInspectResponse, 0, len(args))
		for _, v := range args {
			vol, err := volumes.Inspect(ctx, &api.SdkVolumeInspectRequest{VolumeId: v})
			if err != nil {
				return util.PxErrorMessage(err, "Failed to get volume")
			}
			vols = append(vols, vol)
		}
	} else {
		// If it is no volumes (all)
		volsInfo, err := volumes.InspectWithFilters(ctx, &api.SdkVolumeInspectWithFiltersRequest{})
		if err != nil {
			return util.PxErrorMessage(err, "Failed to get volumes")
		}
		vols = volsInfo.GetVolumes()
	}

	// Get output
	output, _ := cmd.Flags().GetString("output")

	// Determine if we need to output the object
	switch output {
	case "yaml":
		util.PrintYaml(vols)
		return nil
	case "json":
		util.PrintJson(vols)
		return nil
	}

	// Determine if it is a wide output
	wide := output == "wide"

	// Determine if we need to show labels
	showLabels, _ := cmd.Flags().GetBool("show-labels")

	// Start the columns
	t := util.NewTabby()
	np := &volumeColumnFormatter{
		wide:       wide,
		showLabels: showLabels,
		showK8s:    showK8s,
		ctx:        ctx,
		conn:       conn,
		pods:       pods,
	}
	t.AddHeader(np.getHeader()...)

	for _, n := range vols {
		t.AddLine(np.getLine(n)...)
	}
	t.Print()

	return nil
}

type volumeColumnFormatter struct {
	wide       bool
	showLabels bool
	showK8s    bool
	ctx        context.Context
	conn       *grpc.ClientConn
	pods       []v1.Pod
}

func (p *volumeColumnFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.wide {
		header = []interface{}{"Id", "Name", "Size Gi", "HA", "Shared", "Encrypted", "Io Profile", "Status", "State", "Snap Enabled"}
	} else {
		header = []interface{}{"Name", "Size", "HA", "Shared", "Status", "State"}
	}
	if p.showK8s {
		header = append(header, "Pods")
	}
	if p.showLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *volumeColumnFormatter) getLine(resp *api.SdkVolumeInspectResponse) []interface{} {

	v := resp.GetVolume()
	spec := v.GetSpec()

	// Get node information if it is attached
	var node *api.StorageNode
	if len(v.GetAttachedOn()) != 0 {
		nodes := api.NewOpenStorageNodeClient(p.conn)
		nodeInfo, err := nodes.Inspect(
			p.ctx,
			&api.SdkNodeInspectRequest{NodeId: v.GetAttachedOn()})
		if err != nil {
			util.Eprintf("%v\n",
				util.PxErrorMessagef(err,
					"Failed to get node information where volume %s is attached",
					v.GetLocator().GetName()))
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
	if p.showK8s {
		line = append(line, p.podsUsingVolume(v))
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line
}

func prettyStatus(v *api.Volume) string {
	return strings.TrimPrefix(v.GetStatus().String(), "VOLUME_STATUS_")
}

func (p *volumeColumnFormatter) podsUsingVolume(v *api.Volume) string {
	usedPods := make([]string, 0)
	// get the pvc name
	pvc := v.Locator.VolumeLabels["pvc"]
	namespace := v.Locator.VolumeLabels["namespace"]
	for _, pod := range p.pods {
		if pod.Namespace == namespace {
			for _, volumeInfo := range pod.Spec.Volumes {
				if volumeInfo.PersistentVolumeClaim != nil {
					if volumeInfo.PersistentVolumeClaim.ClaimName == pvc {
						usedPods = append(usedPods, namespace+"/"+pod.Name)
					}
				}
			}
		}
	}

	return strings.Join(usedPods, ",")
}
