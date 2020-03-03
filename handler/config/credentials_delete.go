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
package configcli

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type credentialsDeleteFlagsTypes struct {
	Name string
}

// credentialsDeleteCmd represents the config command
var (
	credentialsDeleteCmd   *cobra.Command
	credentialsDeleteFlags *credentialsDeleteFlagsTypes
)

var _ = commander.RegisterCommandVar(func() {
	credentialsDeleteFlags = &credentialsDeleteFlagsTypes{}
	credentialsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete pxc authentication information",
		RunE:  credentialsDeleteExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	CredentialsAddCommand(credentialsDeleteCmd)

	credentialsDeleteCmd.Flags().StringVar(&credentialsDeleteFlags.Name,
		"name", "", "Name for Portworx cluster (ignored when used as a kubectl plugin)")
})

func credentialsDeleteExec(cmd *cobra.Command, args []string) error {
	if err := config.CM().ConfigDeleteAuthInfo(credentialsDeleteFlags.Name); err != nil {
		return nil
	}
	util.Printf("Credentials deleted\n")
	return nil
}
