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

	"github.com/portworx/pxc/pkg/util"
)

const (
	KubeconfigUserPrefix = "pxc@"
)

type ConfigManager struct {
	Config         *Config
	Flags          *ConfigFlags
	configrw       ConfigReaderWriter
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

	if util.InKubectlPluginMode() {
		configManager.configrw = KM()
	} else {
		configManager.configrw = newPxcConfigReaderWriter()
	}

	return configManager
}

func (cm *ConfigManager) Load() error {
	// Load from config file if any
	var err error
	cm.Config, err = cm.ConfigLoad()
	if err != nil {
		return err
	}

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

func (cm *ConfigManager) override() {
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

// ConfigSaveCluster saves the cluster configuration to disk
func (cm *ConfigManager) ConfigSaveCluster(c *Cluster) error {
	return cm.configrw.ConfigSaveCluster(c)
}

// ConfigDeleteCluster deletes the cluster configuration from disk
func (cm *ConfigManager) ConfigDeleteCluster(name string) error {
	return cm.configrw.ConfigDeleteCluster(name)
}

// ConfigLoad loads the configuration from disk
func (cm *ConfigManager) ConfigLoad() (*Config, error) {
	return cm.configrw.ConfigLoad()
}

// ConfigSaveAuthInfo saves user configuration to disk
func (cm *ConfigManager) ConfigSaveAuthInfo(a *AuthInfo) error {
	return cm.configrw.ConfigSaveAuthInfo(a)
}

// ConfigSaveContext saves a context to disk
func (cm *ConfigManager) ConfigSaveContext(c *Context) error {
	configInfo, err := cm.configrw.ConfigLoad()
	if err != nil {
		return err
	}

	if len(c.AuthInfo) != 0 {
		if _, ok := configInfo.AuthInfos[c.AuthInfo]; !ok {
			return fmt.Errorf("Credentials %s do not exist", c.AuthInfo)
		}
	} else {
		c.AuthInfo = "default"
	}
	if _, ok := configInfo.Clusters[c.Cluster]; !ok {
		return fmt.Errorf("Cluster %s does not exist", c.Cluster)
	}
	return cm.configrw.ConfigSaveContext(c)
}

// ConfigDeleteAuthInfo deletes credentials from disk
func (cm *ConfigManager) ConfigDeleteAuthInfo(name string) error {
	return cm.configrw.ConfigDeleteAuthInfo(name)
}

// ConfigDeleteContext deletes context from disk
func (cm *ConfigManager) ConfigDeleteContext(name string) error {
	return cm.configrw.ConfigDeleteContext(name)
}

// ConfigUseContext is not supported by kubectl plugin
func (cm *ConfigManager) ConfigUseContext(name string) error {
	return cm.configrw.ConfigUseContext(name)
}

// ConfigGetCurrentContext returns the current context set by kubectl
func (cm *ConfigManager) ConfigGetCurrentContext() (string, error) {
	return cm.configrw.ConfigGetCurrentContext()
}
