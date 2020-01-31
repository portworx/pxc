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
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/portworx/pxc/pkg/util"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

const (
	KubeconfigUserPrefix = "pxc@"
)

type ConfigManager struct {
	Config         *Config
	Flags          *ConfigFlags
	tunnelEndpoint string
}

var (
	cm *ConfigManager
)

// CM returns the instance to the config manager
func CM() *ConfigManager {
	if cm == nil {
		cm = newConfigManager()
	}

	return cm
}

func newConfig() *Config {
	return &Config{
		Clusters:  make(map[string]*Cluster),
		AuthInfos: make(map[string]*AuthInfo),
		Contexts:  make(map[string]*Context),
	}
}

func newConfigManager() *ConfigManager {
	configManager := &ConfigManager{
		Config: newConfig(),
		Flags:  newConfigFlags(),
	}

	return configManager
}

func (cm *ConfigManager) Load() error {
	// Load from config file if any
	cm.load()

	// Override with flags
	cm.override()

	return nil
}

// GetFlags returns all the pxc persistent flags
func (cm *ConfigManager) GetFlags() *ConfigFlags {
	return cm.Flags
}

// GetConfigFile returns the current config file used
func (cm *ConfigManager) GetConfigFile() string {
	return cm.Flags.GetConfigFile()
}

// SetTunnelEndpoint sets the local endpoint when a tunnel is used
func (cm *ConfigManager) SetTunnelEndpoint(tunnelEndpoint string) {
	cm.tunnelEndpoint = tunnelEndpoint
}

// GetEndpoint returns either the saved endpoint in the config file or the
// tunneled local endpoint
func (cm *ConfigManager) GetEndpoint() string {
	if len(cm.tunnelEndpoint) != 0 {
		return cm.tunnelEndpoint
	}
	return cm.GetCurrentCluster().Endpoint
}

// GetCurrentCluster returns configuration information about the current cluster
func (cm *ConfigManager) GetCurrentCluster() *Cluster {
	return cm.Config.Clusters[cm.Config.Contexts[cm.Config.CurrentContext].Cluster]
}

// GetCurrentAuthInfo returns configuration information about the current user
func (cm *ConfigManager) GetCurrentAuthInfo() *AuthInfo {
	return cm.Config.AuthInfos[cm.Config.Contexts[cm.Config.CurrentContext].AuthInfo]
}

// Write saves the pxc config file
func (cm *ConfigManager) Write() error {
	if len(cm.GetConfigFile()) == 0 {
		panic("cm.GetConfigFile() is 0")
	}
	contextYaml, err := yaml.Marshal(cm.Config)
	if err != nil {
		return fmt.Errorf("Failed to create yaml parse: %v", err)
	}

	// Create the contextconfig location
	err = os.MkdirAll(path.Dir(cm.GetConfigFile()), 0700)
	if err != nil {
		return fmt.Errorf("Failed to create context config dir: %v", err)
	}

	return ioutil.WriteFile(cm.GetConfigFile(), contextYaml, 0600)
}

func (cm *ConfigManager) override() {

	// See if we need to set current context from Kubernetes
	if util.InKubectlPluginMode() {
		// Get the current context, either from the file or from the args to the CLI
		contextName, err := GetKubernetesCurrentContext()
		if err != nil {
			logrus.Fatal(err)
		}

		clientConfig := KM().ToRawKubeConfigLoader()
		kConfig, err := clientConfig.RawConfig()
		if err != nil {
			logrus.Fatalf("unable to read kubernetes configuration: %v", err)
		}

		// Initialize the context
		cm.Config.CurrentContext = contextName
		cm.Config.Contexts[contextName] = &Context{
			AuthInfo: kConfig.Contexts[contextName].AuthInfo,
			Cluster:  kConfig.Contexts[contextName].Cluster,
		}

		// Load all the pxc authentication information from the kubeconfig file
		for k, v := range kConfig.AuthInfos {
			if strings.HasPrefix(k, KubeconfigUserPrefix) && v.AuthProvider != nil {
				logrus.Debugf("Loading user %s from %s", k, v.LocationOfOrigin)
				pxcAuthInfo := NewAuthInfoFromMap(v.AuthProvider.Config)
				cm.Config.AuthInfos[pxcAuthInfo.Name] = pxcAuthInfo
			}
		}

		// Load all the pxc cluster information from the kubeconfig file
		for k, c := range kConfig.Clusters {
			if strings.HasPrefix(k, KubeconfigUserPrefix) {
				pxcClusterInfo, err := NewClusterFromEncodedString(string(c.CertificateAuthorityData))
				if err == nil {
					logrus.Debugf("Loading cluster %s from %s", k, c.LocationOfOrigin)
					cm.Config.Clusters[pxcClusterInfo.Name] = pxcClusterInfo
				} else {
					logrus.Debugf("Unable to load cluster %s from %s", k, c.LocationOfOrigin)
				}
			}
		}
	} else {
		// Not in plugin mode
		if len(cm.Config.CurrentContext) == 0 {
			cm.Config.CurrentContext = "default"
			cm.Config.Contexts[cm.Config.CurrentContext] = &Context{
				AuthInfo: "default",
				Cluster:  "default",
			}
		}
	}

	currentAuth := cm.Config.Contexts[cm.Config.CurrentContext].AuthInfo
	currentCluster := cm.Config.Contexts[cm.Config.CurrentContext].Cluster

	if cm.Config.AuthInfos[currentAuth] == nil {
		cm.Config.AuthInfos[currentAuth] = NewAuthInfo()
	}
	if cm.Config.Clusters[currentCluster] == nil {
		cm.Config.Clusters[currentCluster] = NewDefaultCluster()
	}
	if cm.Config.AuthInfos[currentAuth].KubernetesAuthInfo == nil {
		cm.Config.AuthInfos[currentAuth].KubernetesAuthInfo = &KubernetesAuthInfo{}
	}

	// Get access to the current auth information
	authInfo := cm.GetCurrentAuthInfo()

	// Override with any flags given
	if len(cm.Flags.Token) != 0 {
		authInfo.Token = cm.Flags.Token
	}

	if len(cm.Flags.SecretName) != 0 {
		if authInfo.KubernetesAuthInfo == nil {
			authInfo.KubernetesAuthInfo = &KubernetesAuthInfo{}
		}
		authInfo.KubernetesAuthInfo.SecretName = cm.Flags.SecretName
	}

	if len(cm.Flags.SecretNamespace) != 0 {
		if authInfo.KubernetesAuthInfo == nil {
			authInfo.KubernetesAuthInfo = &KubernetesAuthInfo{}
		}
		authInfo.KubernetesAuthInfo.SecretNamespace = cm.Flags.SecretNamespace
	}
}

func (cm *ConfigManager) load() {
	if _, err := os.Stat(cm.GetConfigFile()); err != nil {
		// Does not exist
		return
	}

	data, err := ioutil.ReadFile(cm.GetConfigFile())
	if err != nil {
		logrus.Fatalf("Failed to load config file %s, %v", cm.GetConfigFile(), err)
	}
	if len(data) == 0 {
		// Empty
		return
	}

	if err := yaml.Unmarshal(data, &cm.Config); err != nil {
		logrus.Fatalf("Failed to process config file %s, %v", cm.GetConfigFile(), err)
	}
}
