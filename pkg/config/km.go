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

package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// KubernetesConfigManager contains all the Kubernetes configuration
type KubernetesConfigManager struct {
	kubeCliOpts *genericclioptions.ConfigFlags
}

var (
	km *KubernetesConfigManager
)

// KM returns the Kubernetes configuration flags and settings
func KM() *KubernetesConfigManager {
	if km == nil {
		km = newKubernetesConfigManager()
	}
	return km
}

func newKubernetesConfigManager() *KubernetesConfigManager {
	return &KubernetesConfigManager{
		kubeCliOpts: genericclioptions.NewConfigFlags(true),
	}
}

// ConfigFlags returns the kubernetes raw configuration object
func (k *KubernetesConfigManager) ConfigFlags() *genericclioptions.ConfigFlags {
	return k.kubeCliOpts
}

// ToRawKubeConfigLoader binds config flag values to config overrides
// Returns an interactive clientConfig if the password flag is enabled,
// or a non-interactive clientConfig otherwise.
// comment from k8s.io/cli-runtime
func (k *KubernetesConfigManager) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	return k.ConfigFlags().ToRawKubeConfigLoader()
}

// ToRESTConfig implements RESTClientGetter.
// Returns a REST client configuration based on a provided path
// to a .kubeconfig file, loading rules, and config flag overrides.
// Expects the AddFlags method to have been called.
// comment from k8s.io/cli-runtime
func (k *KubernetesConfigManager) ToRESTConfig() (*rest.Config, error) {
	return k.ConfigFlags().ToRESTConfig()
}

// GetStartingKubeconfig is used to adjust the current Kubernetes config. You can then
// call ModifyKubeconfig() with the modified configuration
func (k *KubernetesConfigManager) GetStartingKubeconfig() (*clientcmdapi.Config, error) {
	return k.ToRawKubeConfigLoader().ConfigAccess().GetStartingConfig()
}

// ModifyKubeconfig takes a modified configuration and seves it to disk
func (k *KubernetesConfigManager) ModifyKubeconfig(newConfig *clientcmdapi.Config) error {
	return clientcmd.ModifyConfig(k.ToRawKubeConfigLoader().ConfigAccess(), *newConfig, true)
}

// KubectlFlagsToCliArgs rebuilds the flags as cli args
func (k *KubernetesConfigManager) KubectlFlagsToCliArgs() string {
	var args string

	if len(*k.kubeCliOpts.KubeConfig) != 0 {
		args = "--kubeconfig=" + *k.kubeCliOpts.KubeConfig + " "
	}
	if len(*k.kubeCliOpts.Context) != 0 {
		args += "--context=" + *k.kubeCliOpts.Context + " "
	}
	if len(*k.kubeCliOpts.BearerToken) != 0 {
		args += "--token=" + *k.kubeCliOpts.BearerToken + " "
	}
	if len(*k.kubeCliOpts.APIServer) != 0 {
		args += "--server=" + *k.kubeCliOpts.APIServer + " "
	}
	if len(*k.kubeCliOpts.CAFile) != 0 {
		args += "--certificate-authority=" + *k.kubeCliOpts.CAFile + " "
	}
	if len(*k.kubeCliOpts.AuthInfoName) != 0 {
		args += "--user=" + *k.kubeCliOpts.AuthInfoName + " "
	}
	if len(*k.kubeCliOpts.CertFile) != 0 {
		args += "--client-certificate=" + *k.kubeCliOpts.CertFile + " "
	}
	if len(*k.kubeCliOpts.KeyFile) != 0 {
		args += "--client-key=" + *k.kubeCliOpts.KeyFile + " "
	}
	return args
}

// SaveAuthInfoForKubeUser saves the pxc configuration in the kubeconfig file as a new user entry.
// Supply locationOfOrigin so that the Kubernetes saves the object with the appropriate user. LocationOfOrigin
// is found in each of the user objects in the kubernetes Config object.
func (k *KubernetesConfigManager) SaveAuthInfoForKubeUser(user, locationOfOrigin string, a *AuthInfo) error {
	pxcName := KubeconfigUserPrefix + user
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	// If one already exists it will be overwritten, if not create a new object
	if v := oldConfig.AuthInfos[pxcName]; v == nil {
		oldConfig.AuthInfos[pxcName] = clientcmdapi.NewAuthInfo()
	}

	// Store the pxc auth
	oldConfig.AuthInfos[pxcName].LocationOfOrigin = locationOfOrigin
	oldConfig.AuthInfos[pxcName].AuthProvider = &clientcmdapi.AuthProviderConfig{
		Name: "portworx",

		// Change the pxc AuthInfo to a map
		Config: a.toMap(),
	}

	// Save the information in the kubeconfig
	return k.ModifyKubeconfig(oldConfig)
}

// SaveClusterInKubeconfig stores pxc cluster configuration information in Kubeconfig
func (k *KubernetesConfigManager) SaveClusterInKubeconfig(clusterName, location string, c *Cluster) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.Clusters[pxcName]; v == nil {
		oldConfig.Clusters[pxcName] = clientcmdapi.NewCluster()
	}

	encodedString, err := c.toEncodedString()
	if err != nil {
		return err
	}

	oldConfig.Clusters[pxcName].LocationOfOrigin = location
	oldConfig.Clusters[pxcName].Server = "portworx-server"
	oldConfig.Clusters[pxcName].CertificateAuthorityData = []byte(encodedString)

	return k.ModifyKubeconfig(oldConfig)
}

// DeleteClusterInKubeconfig deletes the saved Portworx configuration in the kubeconfig
func (k *KubernetesConfigManager) DeleteClusterInKubeconfig(clusterName string) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.Clusters[pxcName]; v == nil {
		return nil
	}

	delete(oldConfig.Clusters, pxcName)
	return k.ModifyKubeconfig(oldConfig)
}

// GetKubernetesCurrentContext returns the context currently selected by either the config
// file or from the command line
func (k *KubernetesConfigManager) GetKubernetesCurrentContext() (string, error) {
	var contextName string

	kConfig, err := k.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return "", err
	}

	// Check if the was passed in the CLI flags
	if k.kubeCliOpts.Context != nil && len(*k.kubeCliOpts.Context) != 0 {
		contextName = *k.kubeCliOpts.Context
	} else {
		// Read it from the kubeconfig file
		contextName = kConfig.CurrentContext
	}
	logrus.Infof("CurrentContext = %s\n", contextName)

	// Check that it is actually on the kubeconfig file
	if _, ok := kConfig.Contexts[contextName]; !ok {
		return "", fmt.Errorf("context %q does not exist", contextName)
	}
	return contextName, nil
}

// Namespace returns the namespace resulting from the merged
// result of all overrides and a boolean indicating if it was
// overridden
func (k *KubernetesConfigManager) Namespace() (string, bool, error) {
	return k.ToRawKubeConfigLoader().Namespace()
}
