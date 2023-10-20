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
package configcli

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

// contextSetCmd represents the credentials command
var contextSetCmd *cobra.Command
var contextInfo *config.Context

var _ = commander.RegisterCommandVar(func() {
	contextInfo = config.NewContext()
	contextSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set context information for which cluster and credentials to uset",
		Example: `
  # Login to portworx using a secret in Kubernetes
  pxc config context set --credentials=mycreds --cluster=mycluster`,
		RunE: contextExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if !util.InKubectlPluginMode() {
		ContextAddCommand(contextSetCmd)

		contextSetCmd.Flags().StringVar(&contextInfo.Name,
			"name", "", "Name of context")
		contextSetCmd.Flags().StringVar(&contextInfo.AuthInfo,
			"credentials", "", "Name of credentials")
		contextSetCmd.Flags().StringVar(&contextInfo.Cluster,
			"cluster", "", "Name of cluster information")
	}
})

func SetContextAddCommand(cmd *cobra.Command) {
	contextSetCmd.AddCommand(cmd)
}

func contextExec(cmd *cobra.Command, args []string) error {
	if err := config.CM().ConfigSaveContext(contextInfo); err != nil {
		return err
	}
	util.Printf("Context %s set\n", contextInfo.Name)
	return nil
}
