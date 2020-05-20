/*
Copyright 2017 The Kubernetes Authors.

Originally from:
https://raw.githubusercontent.com/kubernetes/pxc/release-1.17/pkg/cmd/plugin/plugin.go

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

package plugin

import (
	"github.com/spf13/cobra"

	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
)

var (
	pluginLong = `
		Provides utilities for interacting with plugins.

		Plugins provide extended functionality that is not part of the major command-line distribution.
		Please refer to the documentation and examples for more information about how write your own plugins.`
)

type pluginCmdFlags struct {
	nameOnly bool
}

var (
	pluginCmd   *cobra.Command
	pluginFlags *pluginCmdFlags
)

var _ = commander.RegisterCommandVar(func() {
	pluginCmd = &cobra.Command{
		Use:                   "component",
		DisableFlagsInUseLine: true,
		Short:                 "Provides utilities for interacting with plugins",
		Long:                  pluginLong,
		Run: func(cmd *cobra.Command, args []string) {
			util.Printf("Please see pxc plugin --help for more information\n")
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.RootAddCommand(pluginCmd)
})

func PluginAddCommand(cmd *cobra.Command) {
	pluginCmd.AddCommand(cmd)
}
