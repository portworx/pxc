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
package pvc

import (
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// pvcCmd represents the pvc command
var pvcCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	pvcCmd = &cobra.Command{
		Use:     "pvc",
		Aliases: []string{"pvcs"},
		Short:   "Manage Kubernetes pvcs on a Portworx cluster",
		Run: func(cmd *cobra.Command, args []string) {
			util.Printf("Please see pxc pvc --help for more commands\n")
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	if util.InKubectlPluginMode() {
		cmd.RootAddCommand(pvcCmd)
	}
})

func PvcAddCommand(cmd *cobra.Command) {
	pvcCmd.AddCommand(cmd)
}
