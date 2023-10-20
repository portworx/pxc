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
package utilities

import (
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// utilitiesCmd represents the utilities command
var utilitiesCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	utilitiesCmd = &cobra.Command{
		Use:     "utilities",
		Aliases: []string{"utils", "util"},
		Short:   "pxc utility commands",
		Run: func(cmd *cobra.Command, args []string) {
			util.Printf("Please see pxc utilities --help for more commands\n")
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.RootAddCommand(utilitiesCmd)
})

func UtilitiesAddCommand(cmd *cobra.Command) {
	utilitiesCmd.AddCommand(cmd)
}
