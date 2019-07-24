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
	"fmt"

	"github.com/portworx/px/pkg/contextconfig"
	"github.com/spf13/cobra"
)

// contextDeleteCmd represents the contextDelete command
var contextDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the given context",
	Long: `Usage:
px context delete --name context1
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return contextDeleteExec(cmd, args)
	},
}

func init() {
	contextCmd.AddCommand(contextDeleteCmd)
	contextDeleteCmd.Flags().String("name", "", "Name of context to delete")
}

func contextDeleteExec(cmd *cobra.Command, args []string) error {
	var nameToDelete string
	if s, _ := cmd.Flags().GetString("name"); len(s) != 0 {
		nameToDelete = s
	} else {
		return fmt.Errorf("Must supply a context name to delete")
	}

	contextManager, err := contextconfig.NewContextManager(cfgFile)
	if err != nil {
		return err
	}

	if err := contextManager.Remove(nameToDelete); err != nil {
		return err
	}
	return nil
}
