// Copyright © 2019 Portworx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package role

import (
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var describeRoleCmd *cobra.Command
var name string

var describeRoleOpts struct {
	req  *api.SdkRoleInspectRequest
	name string
	all  bool
}

var _ = commander.RegisterCommandVar(func() {
	// describeRoleCmd represents the describeRole command
	describeRoleCmd = &cobra.Command{
		Use:     "role [NAME]",
		Aliases: []string{"roles"},
		Short:   "Describe Role",
		Long:    "Display permission rules for a specific role or for all the roles.",
		Example: `
  # To describe testrole.view
  pxc describe role testrole.view

  # To describe all roles
  pxc describe role`,

		RunE: describeRoleExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.DescribeAddCommand(describeRoleCmd)
})

func DescribeAddCommand(cmd *cobra.Command) {
	describeRoleCmd.AddCommand(cmd)
}

func describeRoleExec(c *cobra.Command, args []string) error {
	var roleName string

	ctx, conn, err := portworx.PxConnectDefault()

	if err != nil {
		return err
	}

	defer conn.Close()

	if len(args) != 0 {
		roleName = args[0]
	}

	s := api.NewOpenStorageRoleClient(conn)

	if len(args) == 0 {
		enumRoles, err := s.Enumerate(ctx, &api.SdkRoleEnumerateRequest{})
		if err != nil {
			return util.PxErrorMessage(err, "Role enumeration failed")
		}

		for _, name := range enumRoles.GetNames() {
			roleData, err := s.Inspect(ctx, &api.SdkRoleInspectRequest{
				Name: name,
			})
			if err != nil {
				return util.PxErrorMessage(err, "Describe role failed")
			}
			y, _ := util.ConvertJsonOutputToYaml(roleData)
			util.Printf(string(y))
			util.Printf("\n\n")
		}
	} else {
		roleData, err := s.Inspect(ctx, &api.SdkRoleInspectRequest{
			Name: roleName,
		})
		if err != nil {
			return util.PxErrorMessage(err, "Describe role failed")
		}
		y, _ := util.ConvertJsonOutputToYaml(roleData)
		util.Printf(string(y))
	}
	return nil
}
