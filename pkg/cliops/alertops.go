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

type CliAlertInputs struct {
	util.BaseFormatOutput
	Wide      bool
	AlertType string
	AlertId   string
}

type CliAlertOps struct {
	CliAlertInputs
	PxAlertOps portworx.PxAlertOps
}

func GetCliAlertInputs(cmd *cobra.Command, args []string) *CliAlertInputs {
	output, _ := cmd.Flags().GetString("output")
	alertType, _ := cmd.Flags().GetString("type")
	alertId, _ := cmd.Flags().GetString("id")
	return &CliAlertInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		AlertType: alertType,
		AlertId:   alertId,
	}
}

// Create a new cliAlertOps object
func NewCliAlertOps(
	cvi *CliAlertInputs,
) *CliAlertOps {
	return &CliAlertOps{
		CliAlertInputs: *cvi,
	}
}
