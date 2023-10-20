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

// credentialsSetCmd represents the credentials command
var credentialsSetCmd *cobra.Command
var authInfo *config.AuthInfo

var _ = commander.RegisterCommandVar(func() {
	authInfo = config.NewAuthInfo()
	credentialsSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set authentication information for Portworx cluster",
		Example: `
  # Login to portworx using a secret in Kubernetes
  pxc config credentials set --name=mycreds --k8s-secret-name=abc --k8s-secret-namespace=ns

  # Login to portworx using a specified token
  pxc config credentials set --name=mycreds --auth-token=eyJh...sb30pro`,
		Long: `Saves your Portworx authentication information for the current
user in the kubeconfig file. This will enable pxc to fetch the authentication
information from the config file without having the user provide it each time.`,
		RunE: credentialsExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	CredentialsAddCommand(credentialsSetCmd)

	credentialsSetCmd.Flags().StringVar(&authInfo.Name,
		"name", "", "Name of credential (ignored as a kubectl plugin)")
	credentialsSetCmd.Flags().StringVar(&authInfo.Token,
		"auth-token", "", "Authentication token")
	credentialsSetCmd.Flags().StringVar(&authInfo.KubernetesAuthInfo.SecretNamespace,
		"k8s-secret-namespace", "", "Kubernetes namespace containing the secret with the auth token")
	credentialsSetCmd.Flags().StringVar(&authInfo.KubernetesAuthInfo.SecretName,
		"k8s-secret-name", "", "Kubernetes secret name with the auth token")
})

func credentialsExec(cmd *cobra.Command, args []string) error {
	if err := config.CM().ConfigSaveAuthInfo(authInfo); err != nil {
		return err
	}
	util.Printf("Credentials set\n")
	return nil
}
