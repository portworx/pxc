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
package configcli

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// deleteClusterCmd represents the config command
var (
	deleteClusterCmd *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	deleteClusterCmd = &cobra.Command{
		Use:   "delete-cluster",
		Short: "Delete pxc cluster configuration",
		RunE:  deleteClusterExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ConfigAddCommand(deleteClusterCmd)
})

func deleteClusterExec(cmd *cobra.Command, args []string) error {
	cc := config.KM().ToRawKubeConfigLoader()

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := config.KM().GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	currentContext := kconfig.Contexts[currentContextName]

	// Get the location of the kubeconfig for this specific object. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.Clusters[currentContext.Cluster].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := config.KM().DeleteClusterInKubeconfig(currentContext.Cluster); err != nil {
		return err
	}

	util.Printf("Portworx server information removed from %s for Kubernetes cluster %s\n",
		location,
		currentContext.Cluster)
	return nil
}
