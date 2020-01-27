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
		// Important to save the actual user here so that we do not have to decode
		// the name from `pxcName`
		Name: user,

		// Change the pxc AuthInfo to a map
		Config: a.toMap(),
	}

	// Save the information in the kubeconfig
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
