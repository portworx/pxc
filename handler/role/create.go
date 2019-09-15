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

type createRoleOpts struct {
	req      *api.SdkRoleCreateRequest
	roleconf string
}

var (
	crOpts        *createRoleOpts
	createRoleCmd *cobra.Command
)

func CreateAddCommand(c *cobra.Command) {
	createRoleCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	// createRoleCmd represents the createRole command
	crOpts = &createRoleOpts{
		req: &api.SdkRoleCreateRequest{},
	}
	createRoleCmd = &cobra.Command{
		Use:     "role",
		Aliases: []string{"roles"},
		Short:   "Create a role in Portworx",
		Example: `
  # Create a role using a json file which specifies the role and its rules.
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
  
  # To create a role from the given yaml file
  pxc create role --role-config test.yaml`,

		RunE: createRoleExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.CreateAddCommand(createRoleCmd)

	createRoleCmd.Flags().StringVarP(&crOpts.roleconf, "role-config", "f", "", "Required role config yaml or json file'")
	cobra.MarkFlagRequired(createRoleCmd.Flags(), "role-config")
})

func createRoleExec(c *cobra.Command, args []string) error {

	ctx, conn, err := portworx.PxConnectDefault()
	defer conn.Close()

	s := api.NewOpenStorageRoleClient(conn)

	if _, err := os.Stat(crOpts.roleconf); err != nil {
		return fmt.Errorf("Unable to read role file %s\n", crOpts.roleconf)
	}

	role, err := util.LoadRoleCfg(crOpts.roleconf)
	if err != nil {
		return fmt.Errorf("Error loading role file %v", err)
	}

	_, err = s.Create(
		ctx,
		&api.SdkRoleCreateRequest{
			Role: role,
		})

	if err != nil {
		return util.PxErrorMessage(err, "Role creation failed")
	}

	msg := fmt.Sprintf("Role %s created\n", role.Name)
	formattedOut := &util.DefaultFormatOutput{
		Cmd:  "create role",
		Desc: msg,
		Id:   []string{role.Name},
	}
	util.PrintFormatted(formattedOut)
	return nil
}
