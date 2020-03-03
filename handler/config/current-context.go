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

// currentContextCmd represents the config command
var currentContextCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	currentContextCmd = &cobra.Command{
		Use:   "current-context",
		Short: "Display the current context",
		RunE:  currentContextExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if !util.InKubectlPluginMode() {
		ConfigAddCommand(currentContextCmd)
	}
})

func currentContextExec(cmd *cobra.Command, args []string) error {
	current, err := config.CM().ConfigGetCurrentContext()
	if err != nil {
		return err
	}
	util.Printf("Current context is %s\n", current)
	return nil
}
