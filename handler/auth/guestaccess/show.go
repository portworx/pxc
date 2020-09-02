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

var showGuestAccessCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	showGuestAccessCmd = &cobra.Command{
		Use:   "show",
		Short: "show guest access role",
		RunE:  showGuestAccessExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	GuestAccessAddCommand(showGuestAccessCmd)

	showGuestAccessCmd.Flags().StringP("output", "o", "", "Output in yaml|json|wide")
})

func showGuestAccessExec(cmd *cobra.Command, args []string) error {
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

	// Create the parser object
	authgf := NewGuestAccessShowFormatter(authOps)
	return util.PrintFormatted(authgf)
}

type guestAccessShowFormatter struct {
	cliops.CliAuthOps
}

// NewGuestAccessShowFormatter creates a new guest access formatter
func NewGuestAccessShowFormatter(cvOps *cliops.CliAuthOps) *guestAccessShowFormatter {
	return &guestAccessShowFormatter{
		CliAuthOps: *cvOps,
	}
}

// YamlFormat returns the yaml representation of the object
func (f *guestAccessShowFormatter) YamlFormat() (string, error) {
	role, err := f.AuthOps.GetRole("system.guest")
	if err != nil {
		return "", err
	}
	return util.ToYaml(role)
}

// JsonFormat returns the json representation of the object
func (f *guestAccessShowFormatter) JsonFormat() (string, error) {
	role, err := f.AuthOps.GetRole("system.guest")
	if err != nil {
		return "", err
	}
	return util.ToJson(role)
}

// WideFormat returns the wide string representation of the object
func (f *guestAccessShowFormatter) WideFormat() (string, error) {
	f.Wide = true
	return f.DefaultFormat()
}

// DefaultFormat returns the default string representation of the object
func (f *guestAccessShowFormatter) DefaultFormat() (string, error) {
	role, err := f.AuthOps.GetRole("system.guest")
	if err != nil {
		return "", err
	}

	if role.String() == portworx.RoleGuestDisabled.String() {
		util.Printf("Guest access disabled\n")

	} else {
		util.Printf("Guest access enabled\n")
	}

	return "", nil
}
