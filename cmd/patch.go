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

package cmd

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var patchCmd *cobra.Command

// patchCmd represents the patch command
var _ = commander.RegisterCommandVar(func() {
	patchCmd = &cobra.Command{
		Use:   "patch",
		Short: "Update field(s) of a Portworx resource",
		Run: func(cmd *cobra.Command, args []string) {
			util.Printf("Please see pxc patch --help for more information\n")
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	RootAddCommand(patchCmd)
})

func PatchAddCommand(cmd *cobra.Command) {
	patchCmd.AddCommand(cmd)
}
