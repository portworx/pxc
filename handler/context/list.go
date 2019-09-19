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
package context

import (
	pxcmd "github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/contextconfig"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var contextListCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	contextListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"contexts", "ctx"},
		Short:   "List all context configurations",
		Long:    `List all available context configurations from config.yaml file`,
		Example: `
  # List all context configurations:
  pxc context list`,
		RunE: contextListExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	pxcmd.ContextAddCommand(contextListCmd)
})

func ListAddCommand(cmd *cobra.Command) {
	contextListCmd.AddCommand(cmd)
}

func contextListExec(cmd *cobra.Command, args []string) error {
	contextManager, err := contextconfig.NewContextManager(pxcmd.GetConfigFile())
	if err != nil {
		return util.PxErrorMessagef(err, "Failed to get context configuration at location %s",
			pxcmd.GetConfigFile())
	}
	cfg := contextManager.GetAll()

	// add extra information
	cfg = contextconfig.AddClaimsInfo(cfg)
	cfg = contextconfig.MarkInvalidTokens(cfg)

	// Print out config
	util.PrintYaml(cfg)

	return nil
}
