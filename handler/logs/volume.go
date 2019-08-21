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

package logs

import (
	"fmt"

	pxcmd "github.com/portworx/px/cmd"
	"github.com/portworx/px/pkg/cliops"
	"github.com/portworx/px/pkg/commander"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

var logsVolumeCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	logsVolumeCmd = &cobra.Command{
		Use:   "volume",
		Short: "Print Portworx logs related to specified volume(s)",
		Example: `
        $ px logs volume abc
        Return Portworx logs related to volume abc

        $ px logs volume -f  abc
        Begin streaming the Portworx logs related to volume abc

        $ px logs volume --tail=20 abc
        Apply the volume filters  and the filters specified in --filters to the most recent 20 log lines of each relevant pod  and display only lines that match

        $ px logs node abc --filter "error,warning"
        Display all log lines that is related to volume abc or has either error or warning in the log lines

        $ px logs volume --since=1h volume
        Show all Portworx logs related to volume abc written in the last hour`,
		RunE: logsVolumesExec,
	}
})

// logsCmd represents the logs command
var _ = commander.RegisterCommandInit(func() {
	pxcmd.LogsAddCommand(logsVolumeCmd)
	cliops.AddCommonLogOptions(logsVolumeCmd)
	logsVolumeCmd.Flags().Bool("all-logs", false, "If specified all logs from the pods related to the volume are displayed. Otherwise only log lines that reference the volume or its id is displayed ")
	logsVolumeCmd.Flags().StringP("selector", "l", "", "Selector (label query) comma-separated name=value pairs")
})

func VolumeAddCommand(cmd *cobra.Command) {
	logsVolumeCmd.AddCommand(cmd)
}

func getVolumeLogOptions(
	cmd *cobra.Command,
	args []string,
	cvOps *cliops.CliVolumeOps,
) (*kubernetes.COpsLogOptions, error) {
	err := cliops.ValidateCliInput(cmd, args)
	if err != nil {
		return nil, err
	}

	lo, err := cliops.GetCommonLogOptions(cmd)
	if err != nil {
		return nil, err
	}

	vols, err := cvOps.PxVolumeOps.GetVolumes()
	if err != nil {
		return nil, err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return nil, nil
	}

	allLogs, _ := cmd.Flags().GetBool("all-logs")
	err = cliops.FillContainerInfo(vols, cvOps, lo, allLogs)
	if err != nil {
		return nil, err
	}
	return lo, err
}

func logsVolumesExec(cmd *cobra.Command, args []string) error {
	cvi := cliops.GetCliVolumeInputs(cmd, args)
	cvi.ShowK8s = true
	if len(cvi.Labels) == 0 && len(args) == 0 {
		return fmt.Errorf("Please specify either --selector or volume name")
	}

	// Create a cliVolumeOps object
	cvOps := cliops.NewCliVolumeOps(cvi)

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

	if lo == nil || len(lo.CInfo) == 0 {
		return nil
	}
	return cvOps.PxVolumeOps.GetCOps().GetLogs(lo, util.Stdout)
}
