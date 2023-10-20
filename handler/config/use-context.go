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
package configcli

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type useContextFlagsTypes struct {
	Name string
}

// useContextCmd represents the config command
var (
	useContextFlags *useContextFlagsTypes
	useContextCmd   *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	useContextFlags = &useContextFlagsTypes{}
	useContextCmd = &cobra.Command{
		Use:   "use",
		Short: "Sets the context to use",
		RunE:  useContextExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if !util.InKubectlPluginMode() {
		ContextAddCommand(useContextCmd)

		useContextCmd.Flags().StringVar(&useContextFlags.Name,
			"name", "", "Name of context")
	}
})

func useContextExec(cmd *cobra.Command, args []string) error {
	if err := config.CM().ConfigUseContext(useContextFlags.Name); err != nil {
		return err
	}
	util.Printf("Switched to context %s\n", useContextFlags.Name)
	return nil
}
