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
package guestaccess

import (
	"github.com/portworx/pxc/handler/auth"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// guestAccessCmd represents the alert command
var guestAccessCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	guestAccessCmd = &cobra.Command{
		Use:     "guest-access",
		Aliases: []string{"ga"},
		Short:   "Manage guest access on a Portworx cluster",
		Run: func(cmd *cobra.Command, args []string) {
			util.Printf("Please see pxc auth guest-access --help for more commands\n")
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	auth.AuthAddCommand(guestAccessCmd)
})

// GuestAccessAddCommand adds a guest access command
func GuestAccessAddCommand(cmd *cobra.Command) {
	guestAccessCmd.AddCommand(cmd)
}
