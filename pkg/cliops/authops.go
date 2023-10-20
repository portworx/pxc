/*
Copyright © 2020 Portworx

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

// CliAuthOps represents an interface for auth commands
type CliAuthOps struct {
	portworx.CliAuthInputs
	AuthOps portworx.AuthOps
}

// GetCliAuthInputs gets all CLI auth inputs
func GetCliAuthInputs(cmd *cobra.Command, args []string) *portworx.CliAuthInputs {
	output, _ := cmd.Flags().GetString("output")
	return &portworx.CliAuthInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
	}
}

// NewCliAuthOps creates a new cliAuthOps object
func NewCliAuthOps(
	cvi *portworx.CliAuthInputs,
) *CliAuthOps {
	return &CliAuthOps{
		CliAuthInputs: *cvi,
	}
}
