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
		Long: `Saves your Portworx authentication information for the current
user in the kubeconfig file. This will enable pxc to fetch the authentication
information from the config file without having the user provide it each time.`,
		RunE: loginExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if util.InKubectlPluginMode() {
		cmd.RootAddCommand(loginCmd)
	}
})

func LoginAddCommand(cmd *cobra.Command) {
	loginCmd.AddCommand(cmd)
}

func loginExec(cmd *cobra.Command, args []string) error {
	cc := config.KM().ToRawKubeConfigLoader()
	save := false

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := config.GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	currentContext := kconfig.Contexts[currentContextName]
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

	// Check if any information necessary was passed
	if !save {
		return fmt.Errorf("Must supply authentication information")
	}

	// Get the location of the kubeconfig for this specific authInfo. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.AuthInfos[currentContext.AuthInfo].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := config.SaveAuthInfoForKubeUser(currentContext.AuthInfo, location, authInfo); err != nil {
		return err
	}

	util.Printf("Portworx login information saved in %s for Kubernetes user context %s\n",
		location,
		currentContext.AuthInfo)
	return nil
}
