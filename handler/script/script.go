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
package script

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	EvEndpoint        = "OPENSTORAGE_SDK_ENDPOINT"
	EvSecretName      = "OPENSTORAGE_SDK_SECRET_NAME"
	EvSecretNamespace = "OPENSTORAGE_SDK_SECRET_NAMESPACE"
	EvToken           = "OPENSTORAGE_SDK_TOKEN"
	EvSecure          = "OPENSTORAGE_SDK_SECURE"
	EvCAFile          = "OPENSTORAGE_SDK_CAFILE"
)

type scriptFlags struct {
	language string
}

// scriptCmd represents the script command
var (
	scriptCmd *cobra.Command
	flags     *scriptFlags
)

var _ = commander.RegisterCommandVar(func() {
	flags = &scriptFlags{}
	scriptCmd = &cobra.Command{
		Use:   "script [NAME]",
		Short: "Run a script against the current cluster",
		RunE:  runScriptExec,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a script")
			}
			return nil
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.RootAddCommand(scriptCmd)
	scriptCmd.Flags().StringVarP(&flags.language, "language", "l", "python", "Script language. Currently only python is supported.")
})

func runScriptExec(cmd *cobra.Command, args []string) error {
	switch flags.language {
	case "python":
		return pythonScriptExec(cmd, args)
	default:
		return fmt.Errorf("Unsupported script language type: %s", flags.language)
	}
}

func pythonScriptExec(cmd *cobra.Command, args []string) error {

	// Setup a connetion to Portworx
	clusterInfo := config.CM().GetCurrentCluster()
	authInfo := config.CM().GetCurrentAuthInfo()

	// Check if we need to tunnel
	if len(config.CM().GetEndpoint()) == 0 &&
		util.InKubectlPluginMode() {
		err := kubernetes.StartTunnel()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Must provide an endpoint to Portworx server")
	}

	logrus.Infof("args: %+v", args)
	scriptCmd := exec.Command("python3", args...)
	scriptCmd.Env = append(os.Environ(),
		EvEndpoint+"="+config.CM().GetEndpoint(),
		EvCAFile+"="+clusterInfo.CACert,
	)

	if clusterInfo.Secure {
		scriptCmd.Env = append(scriptCmd.Env, EvSecure+"=true")
	}

	if len(authInfo.KubernetesAuthInfo.SecretNamespace) != 0 &&
		len(authInfo.KubernetesAuthInfo.SecretName) != 0 {
		scriptCmd.Env = append(scriptCmd.Env,
			EvSecretNamespace+"="+authInfo.KubernetesAuthInfo.SecretNamespace,
			EvSecretName+"="+authInfo.KubernetesAuthInfo.SecretName)
	}

	if len(authInfo.Token) != 0 {
		scriptCmd.Env = append(scriptCmd.Env,
			EvToken+"="+authInfo.Token)
	}

	logrus.Debugf("env: %+v", scriptCmd.Env)
	output, err := scriptCmd.CombinedOutput()
	util.Printf("%s", string(output))
	return err
}
