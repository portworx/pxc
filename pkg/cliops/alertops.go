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
package cliops

import (
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type CliAlertOps struct {
	portworx.CliAlertInputs
	PxAlertOps portworx.PxAlertOps
}

func GetCliAlertInputs(cmd *cobra.Command, args []string) *portworx.CliAlertInputs {
	output, _ := cmd.Flags().GetString("output")
	alertType, _ := cmd.Flags().GetString("type")
	alertId, _ := cmd.Flags().GetString("id")
	startTime, _ := cmd.Flags().GetString("start-time")
	endTime, _ := cmd.Flags().GetString("end-time")
	severity, _ := cmd.Flags().GetString("severity")
	resourceId, _ := cmd.Flags().GetString("resource-id")
	return &portworx.CliAlertInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		AlertType:  alertType,
		AlertId:    alertId,
		StartTime:  startTime,
		EndTime:    endTime,
		Severity:   severity,
		ResourceId: resourceId,
	}
}

// Create a new cliAlertOps object
func NewCliAlertOps(
	cvi *portworx.CliAlertInputs,
) *CliAlertOps {
	return &CliAlertOps{
		CliAlertInputs: *cvi,
	}
}
