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
package login

import (
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Set authentication information for Portworx cluster",
		Example: `
  # Login to portworx using a secret in Kubernetes
  pxc login --pxc.secret-name=abc --pxc.secret-namespace=ns

  # Login to portworx using a specified token
  pxc login --pxc.token=ey..`,
		Long: `Saves your Portworx authentication information in the pxc
config file. This will enable pxc to fetch the authentication information
from the config file without having the user provide it each time.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cc, _, err := kubernetes.KubeConnectDefault()
			if err != nil {
				return err
			}

			kconfig, err := cc.RawConfig()
			currentContext := kconfig.Contexts[kconfig.CurrentContext]

			config.CM().Config.AuthInfos[currentContext.AuthInfo] = &config.AuthInfo{
				Name:  currentContext.AuthInfo,
				Token: config.CM().GetFlags().Token,
				KubernetesAuthInfo: &config.KubernetesAuthInfo{
					SecretName:      config.CM().GetFlags().SecretName,
					SecretNamespace: config.CM().GetFlags().SecretNamespace,
				},
			}

			err = config.CM().Write()
			if err != nil {
				util.Eprintf("Failed to save login information to %s: %v\n",
					config.CM().GetConfigFile(),
					err)
				return err
			}

			util.Printf("Login information saved to %s\n", config.CM().GetConfigFile())
			return nil
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	if util.InKubectlPluginMode() {
		cmd.RootAddCommand(loginCmd)
	}
})
