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
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd *cobra.Command
var authInfo *config.AuthInfo

var _ = commander.RegisterCommandVar(func() {
	authInfo = config.NewAuthInfo()
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Set authentication information for Portworx cluster",
		Example: `
  # Login to portworx using a secret in Kubernetes
  pxc login --k8s-secret-name=abc --k8s-secret-namespace=ns

  # Login to portworx using a specified token
  pxc login --auth-token=eyJh...sb30ro`,
		Long: `Saves your Portworx authentication information for the current
user in the kubeconfig file for future access of the Portworx system.`,
		RunE: loginExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if util.InKubectlPluginMode() {
		cmd.RootAddCommand(loginCmd)

		loginCmd.Flags().StringVar(&authInfo.Token,
			"auth-token", "", "Auth token if any (optional)")
		loginCmd.Flags().StringVar(&authInfo.KubernetesAuthInfo.SecretNamespace,
			"k8s-secret-namespace", "", "Kubernetes namespace containing the secret with the auth token")
		loginCmd.Flags().StringVar(&authInfo.KubernetesAuthInfo.SecretName,
			"k8s-secret-name", "", "Kubernetes secret name with the auth token")
	}
})

func LoginAddCommand(cmd *cobra.Command) {
	loginCmd.AddCommand(cmd)
}

func loginExec(cmd *cobra.Command, args []string) error {
	err := config.CM().ConfigSaveAuthInfo(authInfo)
	if err != nil {
		return err
	}
	util.Printf("Successfully saved login information in Kubeconfig\n")
	return nil
}
