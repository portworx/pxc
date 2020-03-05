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

package volume

import (
	"fmt"

	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var logsVolumeCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	logsVolumeCmd = &cobra.Command{
		Use:   "logs [NAME]",
		Short: "(NOT WORKING WELL) Print Portworx logs related to specified volume(s)",
		Example: `
  # Return Portworx logs related to volume abc
  pxc volume logs abc

  # Begin streaming the Portworx logs related to volume abc
  pxc volume logs -f  abc

  # Apply the volume filters  and the filters specified in --filters to the most recent 20 log lines of each relevant pod  and display only lines that match
  pxc volume logs --tail=20 abc

  # Display all log lines that is related to volume abc or has either error or warning in the log lines
  pxc volume logs abc --filter "error,warning"

  # Show all Portworx logs related to volume abc written in the last hour
  pxc volume logs --since=1h volume`,
		RunE: logsVolumesExec,
	}
})

// logsCmd represents the logs command
var _ = commander.RegisterCommandInit(func() {
	VolumeAddCommand(logsVolumeCmd)
	cliops.AddCommonLogOptions(logsVolumeCmd)
	logsVolumeCmd.Flags().Bool("all-logs", false, "If specified all logs from the pods related to the volume are displayed. Otherwise only log lines that reference the volume or its id is displayed ")
	logsVolumeCmd.Flags().StringP("selector", "l", "", "Selector (label query) comma-separated name=value pairs")
})

func getVolumeLogOptions(
	cmd *cobra.Command,
	args []string,
	cliOps cliops.CliOps,
) (*kubernetes.COpsLogOptions, error) {
	err := cliops.ValidateCliInput(cmd, args)
	if err != nil {
		return nil, err
	}

	lo, err := cliops.GetCommonLogOptions(cmd)
	if err != nil {
		return nil, err
	}

	volSpec := &portworx.VolumeSpec{
		VolNames: cliOps.CliInputs().Args,
		Labels:   cliOps.CliInputs().Labels,
	}

	vo := portworx.NewVolumes(cliOps.PxOps(), volSpec)

	vols, err := vo.GetVolumes()
	if err != nil {
		return nil, err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return nil, nil
	}

	allLogs, _ := cmd.Flags().GetBool("all-logs")
	err = cliops.FillContainerInfo(vols, cliOps, lo, allLogs)
	if err != nil {
		return nil, err
	}
	return lo, err
}

func logsVolumesExec(cmd *cobra.Command, args []string) error {
	cvi := cliops.NewCliInputs(cmd, args)
	cvi.ShowK8s = true
	if len(cvi.Labels) == 0 && len(args) == 0 {
		return fmt.Errorf("Please specify either --selector or volume name")
	}

	// Create a cliVolumeOps object
	cliOps := cliops.NewCliOps(cvi)

	// Connect to pxc and k8s (if needed)
	err := cliOps.Connect()
	if err != nil {
		return err
	}
	defer cliOps.Close()

	lo, err := getVolumeLogOptions(cmd, args, cliOps)
	if err != nil {
		return err
	}

	if lo == nil || len(lo.CInfo) == 0 {
		return nil
	}
	return cliOps.COps().GetLogs(lo, util.Stdout)
}
