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
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	kubeCliOpts *genericclioptions.ConfigFlags
)

// KM returns the Kubernetes configuration flags and settings
func KM() *genericclioptions.ConfigFlags {
	if kubeCliOpts == nil {
		kubeCliOpts = genericclioptions.NewConfigFlags(true)
	}
	return kubeCliOpts
}

// GetStartingKubeconfig is used to adjust the current Kubernetes config. You can then
// call ModifyKubeconfig() with the modified configuration
func GetStartingKubeconfig() (*clientcmdapi.Config, error) {
	return KM().ToRawKubeConfigLoader().ConfigAccess().GetStartingConfig()
}

// ModifyKubeconfig takes a modified configuration and seves it to disk
func ModifyKubeconfig(newConfig *clientcmdapi.Config) error {
	return clientcmd.ModifyConfig(KM().ToRawKubeConfigLoader().ConfigAccess(), *newConfig, true)
}

// KubectlFlagsToCliArgs rebuilds the flags as cli args
func KubectlFlagsToCliArgs() string {
	var args string

	if len(*KM().KubeConfig) != 0 {
		args = "--kubeconfig=" + *KM().KubeConfig + " "
	}
	if len(*KM().Context) != 0 {
		args += "--context=" + *KM().Context + " "
	}
	if len(*KM().BearerToken) != 0 {
		args += "--token=" + *KM().BearerToken + " "
	}
	if len(*KM().APIServer) != 0 {
		args += "--server=" + *KM().APIServer + " "
	}
	if len(*KM().CAFile) != 0 {
		args += "--certificate-authority=" + *KM().CAFile + " "
	}
	if len(*KM().AuthInfoName) != 0 {
		args += "--user=" + *KM().AuthInfoName + " "
	}
	if len(*KM().CertFile) != 0 {
		args += "--client-certificate=" + *KM().CertFile + " "
	}
	if len(*KM().KeyFile) != 0 {
		args += "--client-key=" + *KM().KeyFile + " "
	}
	return args
}

// SaveAuthInfoForKubeUser saves the pxc configuration in the kubeconfig file as a new user entry.
// Supply locationOfOrigin so that the Kubernetes saves the object with the appropriate user. LocationOfOrigin
// is found in each of the user objects in the kubernetes Config object.
func SaveAuthInfoForKubeUser(user, locationOfOrigin string, a *AuthInfo) error {
	pxcName := KubeconfigUserPrefix + user
	oldConfig, err := GetStartingKubeconfig()
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
	return ModifyKubeconfig(oldConfig)
}

// SaveClusterInKubeconfig stores pxc cluster configuration information in Kubeconfig
func SaveClusterInKubeconfig(clusterName, location string, c *Cluster) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := GetStartingKubeconfig()
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

	return ModifyKubeconfig(oldConfig)
}

//
func DeleteClusterInKubeconfig(clusterName string) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.Clusters[pxcName]; v == nil {
		return nil
	}

	delete(oldConfig.Clusters, pxcName)
	return ModifyKubeconfig(oldConfig)
}

// GetKubernetesCurrentContext returns the context currently selected by either the config
// file or from the command line
func GetKubernetesCurrentContext() (string, error) {
	var contextName string

	kConfig, err := KM().ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return "", err
	}

	// Check if the was passed in the CLI flags
	if KM().Context != nil && len(*KM().Context) != 0 {
		contextName = *KM().Context
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
