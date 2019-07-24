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

// setCurrentContextCmd represents the setCurrentContext command
var contextSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the current context configuration",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return contextSetExec(cmd, args)
	},
}

func init() {
	contextCmd.AddCommand(contextSetCmd)
	contextSetCmd.Flags().String("name", "", "Name of current context to set")
}

func contextSetExec(cmd *cobra.Command, args []string) error {
	var name string

	if s, _ := cmd.Flags().GetString("name"); len(s) != 0 {
		name = s
	} else {
		return fmt.Errorf("Must supply a name for the context")
	}

	contextManager, err := contextconfig.NewContextManager(GetConfigFile())
	if err != nil {
		return err
	}

	if err := contextManager.SetCurrent(name); err != nil {
		return err
	}

	return nil
}
