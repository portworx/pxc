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
package context

import (
	pxcmd "github.com/portworx/px/cmd"
	"github.com/portworx/px/pkg/commander"
	"github.com/portworx/px/pkg/contextconfig"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

var contextCurrentCmd *cobra.Command

// contextCurrentCmd represents the contextCurrent command
var _ = commander.RegisterCommandVar(func() {
	contextCurrentCmd = &cobra.Command{
		Use:     "current",
		Aliases: []string{"show", "current-context"},
		Short:   "Show current context name",
		RunE:    contextCurrentExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	pxcmd.ContextAddCommand(contextCurrentCmd)
})

func CurrentAddContext(cmd *cobra.Command) {
	contextCurrentCmd.AddCommand(cmd)
}

func contextCurrentExec(cmd *cobra.Command, args []string) error {
	contextManager, err := contextconfig.NewContextManager(pxcmd.GetConfigFile())
	if err != nil {
		return util.PxErrorMessagef(err, "Failed to get context configuration at location %s",
			pxcmd.GetConfigFile())
	}
	pxctx, err := contextManager.GetCurrent()
	if err != nil {
		return err
	}
	util.Printf("%s\n", pxctx.Name)
	return nil
}
