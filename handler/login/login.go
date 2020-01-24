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
	"fmt"
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

			save := false
			kconfig, err := cc.RawConfig()
			currentContext := kconfig.Contexts[kconfig.CurrentContext]
			configFlags := config.CM().GetFlags()

			// Initialize authInfo object
			authInfo := &config.AuthInfo{
				Name: currentContext.AuthInfo,
			}

			// Check for token
			if len(configFlags.Token) != 0 {
				save = true
				authInfo.Token = configFlags.Token
				// TODO: Validate if the token is expired
			}

			// Check for Kubernetes secret and secret namespace
			if len(configFlags.SecretNamespace) != 0 && len(configFlags.SecretName) != 0 {
				save = true
				authInfo.KubernetesAuthInfo = &config.KubernetesAuthInfo{
					SecretName:      configFlags.SecretName,
					SecretNamespace: configFlags.SecretNamespace,
				}
			} else if len(configFlags.SecretNamespace) == 0 && len(configFlags.SecretName) != 0 {
				return fmt.Errorf("Must supply secret namespace with secret name")
			} else if len(configFlags.SecretNamespace) != 0 && len(configFlags.SecretName) == 0 {
				return fmt.Errorf("Must supply secret name with secret namespace")
			}

			if !save {
				return fmt.Errorf("Must supply authentication information")
			}

			config.CM().Config.AuthInfos[currentContext.AuthInfo] = authInfo

			clusterInfo := config.CM().GetCurrentCluster()
			clusterInfo.Name = currentContext.Cluster
			clusterInfo.Kubeconfig = kconfig.Clusters[currentContext.Cluster].LocationOfOrigin

			err = config.CM().Write()
			if err != nil {
				return fmt.Errorf("Failed to save login information to %s: %v\n",
					config.CM().GetConfigFile(),
					err)
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
