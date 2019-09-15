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
package role

import (
	"fmt"
	"os"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type updateRoleOpts struct {
	req      *api.SdkRoleUpdateRequest
	roleconf string
}

var (
	urOpts        *updateRoleOpts
	updateRoleCmd *cobra.Command
)

func PatchAddCommand(c *cobra.Command) {
	updateRoleCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	// patchRoleCmd represents the patchRole command
	urOpts = &updateRoleOpts{
		req: &api.SdkRoleUpdateRequest{},
	}
	updateRoleCmd = &cobra.Command{
		Use:     "role",
		Aliases: []string{"roles"},
		Short:   "Update a role in Portworx",

		Example: `
  # Update/Patch a role using a json file which specifies the role and its rules.
  # A role consist of a set of rules defining services and api's which are allowable.
  # e.g. Rule file(test.yaml) which allows inspection of any object and listings of only volumes:

  ---
  name: testrole.view
  rules:
  - services: ["volumes"]
    apis: ["*enumerate*"]
  - services: ["*"]
    apis: ["inspect*"]
  ---

  pxc patch role --role-config test.yaml`,

		RunE: updateRoleExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.PatchAddCommand(updateRoleCmd)

	updateRoleCmd.Flags().StringVarP(&urOpts.roleconf, "role-config", "f", "", "Required role config yaml or json file'")
	cobra.MarkFlagRequired(updateRoleCmd.Flags(), "role-config")
})

func updateRoleExec(c *cobra.Command, args []string) error {

	ctx, conn, err := portworx.PxConnectDefault()

	if err != nil {
		return err
	}

	defer conn.Close()

	s := api.NewOpenStorageRoleClient(conn)

	if _, err := os.Stat(urOpts.roleconf); err != nil {
		return fmt.Errorf("Unable to read role file %s\n", urOpts.roleconf)
	}

	role, err := util.LoadRoleCfg(urOpts.roleconf)
	if err != nil {
		return fmt.Errorf("Error reading config file %v", err)
	}

	_, err = s.Update(
		ctx,
		&api.SdkRoleUpdateRequest{
			Role: role,
		})
	if err != nil {
		return util.PxErrorMessage(err, "Role patching failed")
	}
	msg := fmt.Sprintf("Role %s patched\n", role.Name)
	formattedOut := &util.DefaultFormatOutput{
		Cmd:  "patch role",
		Desc: msg,
		Id:   []string{role.Name},
	}
	return util.PrintFormatted(formattedOut)
}
