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
	"fmt"

	pxcmd "github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/contextconfig"
	"github.com/spf13/cobra"
)

// contextDeleteCmd represents the contextDelete command
var contextDeleteCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	contextDeleteCmd = &cobra.Command{
		Use:     "delete [NAME]",
		Short:   "Deletes the given context",
		Example: "$ pxc context delete mycontext",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for context")
			}
			return nil
		},
		Long: `Usage:
px context delete --name context1
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return contextDeleteExec(cmd, args)
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	pxcmd.ContextAddCommand(contextDeleteCmd)
	contextDeleteCmd.Flags().String("name", "", "Name of context to delete")
})

func DeleteAddCommand(cmd *cobra.Command) {
	contextDeleteCmd.AddCommand(cmd)
}

func contextDeleteExec(cmd *cobra.Command, args []string) error {
	nameToDelete := args[0]

	contextManager, err := contextconfig.NewContextManager(pxcmd.GetConfigFile())
	if err != nil {
		return err
	}

	if err := contextManager.Remove(nameToDelete); err != nil {
		return err
	}
	return nil
}
