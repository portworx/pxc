// Copyright © 2019 Portworx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
)

var logsVolumeCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	logsVolumeCmd = &cobra.Command{
		Use:   "volume",
		Short: "Print Portworx logs related to specified volume(s)",
		Example: `$ px logs volume abc
        Return Portworx logs related to volume abc

        $ px logs volume -f  abc
        Begin streaming the Portworx logs related to volume abc

        $ px logs volume --tail=20 abc
        Display only the most recent 20 lines of Portworx logs related to volume abc from each node

        $ px logs volume --since=1h volume
        Show all Portworx logs related to volume abc written in the last hour`,
		RunE: logsVolumesExec,
	}
})

// logsCmd represents the logs command
var _ = RegisterCommandInit(func() {
	logsCmd.AddCommand(logsVolumeCmd)
	addCommonLogOptions(logsVolumeCmd)
	logsVolumeCmd.Flags().Bool("all-logs", false, "If specified all logs from the pods related to the volume are displayed. Otherwise only log lines that reference the volume or its id is displayed ")
	logsVolumeCmd.Flags().StringP("selector", "l", "", "Selector (label query) comma-separated name=value pairs")
})

func getVolumeLogOptions(
	cmd *cobra.Command,
	args []string,
	cvOps *cliVolumeOps,
) (*portworx.COpsLogOptions, error) {
	err := validateCliInput(cmd, args)
	if err != nil {
		return nil, err
	}

	lo, err := getCommonLogOptions(cmd)
	if err != nil {
		return nil, err
	}

	vols, err := cvOps.pxVolumeOps.GetVolumes()
	if err != nil {
		return nil, err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return nil, nil
	}

	// Get All relevant pods.
	allLogs, _ := cmd.Flags().GetBool("all-logs")
	nodeNamesMap := make(map[string]bool)
	podList := make(map[string]v1.Pod)
	for _, resp := range vols {
		// Get all of the nodes associated with the volume
		// Get all of the pods using the volume
		v := resp.GetVolume()
		if allLogs != true {
			lo.Filters = append(lo.Filters, v.GetLocator().GetName())
			lo.Filters = append(lo.Filters, v.GetId())
			lo.ApplyFilters = true
		}

		err := cvOps.pxVolumeOps.GetAllNodesForVolume(v, nodeNamesMap)
		if err != nil {
			return nil, err
		}

		pods, err := cvOps.pxVolumeOps.PodsUsingVolume(v)
		if err != nil {
			return nil, err
		}

		for _, p := range pods {
			key := fmt.Sprintf("%s-%s", p.Namespace, p.Name)
			podList[key] = p
		}

	}

	nodeNames := make([]string, 0)
	for k, _ := range nodeNamesMap {
		nodeNames = append(nodeNames, k)
	}

	// Convert node names to pods
	pods, err := getRequiredPortworxPods(cvOps, nodeNames, lo.PortworxNamespace)
	if err != nil {
		return nil, err
	}

	// Remove duplicates between the list of pods that are attaching the volume and the portworx pods if any
	for _, p := range pods {
		key := fmt.Sprintf("%s-%s", p.Namespace, p.Name)
		podList[key] = p
	}

	// Covert the pod map to an array of pods
	lo.Pods = make([]v1.Pod, 0)
	for _, p := range podList {
		lo.Pods = append(lo.Pods, p)
	}
	return lo, nil
}

func logsVolumesExec(cmd *cobra.Command, args []string) error {
	validateCliInput(cmd, args)
	cvi := GetCliVolumeInputs(cmd, args)
	cvi.showK8s = true
	if len(cvi.labels) == 0 && len(args) == 0 {
		return fmt.Errorf("Please specify either --selector or volume name")
	}

	// Create a cliVolumeOps object
	cvOps := NewCliVolumeOps(cvi)

	// Connect to px and k8s (if needed)
	err := cvOps.Connect()
	if err != nil {
		return err
	}
	defer cvOps.Close()

	lo, err := getVolumeLogOptions(cmd, args, cvOps)
	if err != nil {
		return err
	}

	if lo == nil || len(lo.Pods) == 0 {
		return nil
	}
	return cvOps.pxVolumeOps.GetCOps().GetLogs(lo, util.Stdout)
}
