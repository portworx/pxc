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
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var enableGuestAccessCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	enableGuestAccessCmd = &cobra.Command{
		Use:   "enable",
		Short: "enable guest access role",
		RunE:  enableGuestAccessExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	GuestAccessAddCommand(enableGuestAccessCmd)
})

func enableGuestAccessExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
	_ = ctx
	if err != nil {
		return err
	}
	defer conn.Close()
	// Parse out all of the common cli volume flags
	cai := cliops.GetCliAuthInputs(cmd, args)

	// Create a cliVolumeOps object
	authOps := cliops.NewCliAuthOps(cai)

	// initialize alertOP interface
	authOps.AuthOps = portworx.NewAuthOps()

	err = authOps.AuthOps.UpdateRole(&portworx.RoleGuestEnabled)
	if err == nil {
		util.Printf("Guest access enabled successfully\n")
	}

	return err
}
