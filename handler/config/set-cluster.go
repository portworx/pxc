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

type setClusterFlagsTypes struct {
	TunnelServiceNamespace string
	TunnelServiceName      string
	TunnelServicePort      string
}

// setClusterCmd represents the config command
var (
	setClusterCmd   *cobra.Command
	setClusterFlags *setClusterFlagsTypes
)

var _ = commander.RegisterCommandVar(func() {
	setClusterFlags = &setClusterFlagsTypes{}
	setClusterCmd = &cobra.Command{
		Use:   "set-cluster",
		Short: "Setup pxc cluster configuration",
		RunE:  setClusterExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ConfigAddCommand(setClusterCmd)

	setClusterCmd.Flags().StringVar(&setClusterFlags.TunnelServiceNamespace,
		"portworx-service-namespace", "kube-system", "Kubernetes namespace for the Portworx service")
	setClusterCmd.Flags().StringVar(&setClusterFlags.TunnelServiceName,
		"portworx-service-name", "portworx-api", "Kubernetes name for the Portworx service")
	setClusterCmd.Flags().StringVar(&setClusterFlags.TunnelServicePort,
		"portworx-service-port", "9020", "Port for the Portworx SDK endpoint in the Kubernetes service")
})

func setClusterExec(cmd *cobra.Command, args []string) error {
	cc := config.KM().ToRawKubeConfigLoader()

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

	// Initialize cluster object
	clusterInfo := config.NewCluster()
	clusterInfo.Name = currentContext.Cluster
	clusterInfo.TunnelServiceName = setClusterFlags.TunnelServiceName
	clusterInfo.TunnelServiceNamespace = setClusterFlags.TunnelServiceNamespace
	clusterInfo.TunnelServicePort = setClusterFlags.TunnelServicePort

	// Get the location of the kubeconfig for this specific object. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.Clusters[currentContext.Cluster].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := config.SaveClusterInKubeconfig(currentContext.Cluster, location, clusterInfo); err != nil {
		return err
	}

	util.Printf("Portworx server information saved in %s for Kubernetes cluster %s\n",
		location,
		currentContext.Cluster)
	return nil
}
