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

	yaml "gopkg.in/yaml.v2"
)

type pxcConfigReaderWriter struct {
}

func newPxcConfigReaderWriter() *pxcConfigReaderWriter {
	return &pxcConfigReaderWriter{}
}

// ConfigSaveCluster saves the cluster configuration to disk
func (p *pxcConfigReaderWriter) ConfigSaveCluster(c *Cluster) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("Must supply a name for the cluster")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	configInfo.Clusters[c.Name] = c

	return p.save(configInfo)
}

// ConfigDeleteCluster deletes the cluster configuration from disk
func (p *pxcConfigReaderWriter) ConfigDeleteCluster(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("Must supply a name for the cluster")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	if _, ok := configInfo.Clusters[name]; !ok {
		return fmt.Errorf("Cluster %s was not found in %s", name, CM().GetConfigFile())
	}
	delete(configInfo.Clusters, name)

	return p.save(configInfo)
}

// ConfigLoad loads the configuration from disk
func (p *pxcConfigReaderWriter) ConfigLoad() (*Config, error) {
	return p.load()
}

// ConfigSaveAuthInfo saves user configuration to disk
func (p *pxcConfigReaderWriter) ConfigSaveAuthInfo(a *AuthInfo) error {
	if len(a.Name) == 0 {
		return fmt.Errorf("Must supply a name for the credential")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	configInfo.AuthInfos[a.Name] = a

	return p.save(configInfo)
}

func (p *pxcConfigReaderWriter) ConfigSaveContext(c *Context) error {
	if len(c.Name) == 0 {
		return fmt.Errorf("Must supply a name for the context")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	configInfo.Contexts[c.Name] = c

	return p.save(configInfo)
}

func (p *pxcConfigReaderWriter) ConfigDeleteAuthInfo(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("Must supply a name for the credential to delete")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	if _, ok := configInfo.AuthInfos[name]; !ok {
		return fmt.Errorf("Credentials %s were not found in %s", name, CM().GetConfigFile())
	}
	delete(configInfo.AuthInfos, name)

	return p.save(configInfo)
}

func (p *pxcConfigReaderWriter) ConfigDeleteContext(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("Must supply a name for the context to delete")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	if _, ok := configInfo.Contexts[name]; !ok {
		return fmt.Errorf("Context %s was not found in %s", name, CM().GetConfigFile())
	}
	delete(configInfo.Contexts, name)

	return p.save(configInfo)
}

// ConfigUseContext sets the current context
func (p *pxcConfigReaderWriter) ConfigUseContext(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("Must supply a context name")
	}

	configInfo, err := p.load()
	if err != nil {
		return err
	}
	if _, ok := configInfo.Contexts[name]; !ok {
		return fmt.Errorf("Context %s was not found in %s", name, CM().GetConfigFile())
	}
	configInfo.CurrentContext = name

	return p.save(configInfo)
}

// ConfigGetCurrentContext returns the current context
func (p *pxcConfigReaderWriter) ConfigGetCurrentContext() (string, error) {
	configInfo, err := p.load()
	if err != nil {
		return "", err
	}
	if len(configInfo.CurrentContext) == 0 {
		return "", fmt.Errorf("Current context has not been set")
	}
	return configInfo.CurrentContext, nil
}

// Write saves the pxc config file
func (p *pxcConfigReaderWriter) save(configInfo *Config) error {
	if len(CM().GetConfigFile()) == 0 {
		panic("path is empty")
	}
	contextYaml, err := yaml.Marshal(configInfo)
	if err != nil {
		return fmt.Errorf("Failed to create yaml parse: %v", err)
	}

	// Create the contextconfig location
	err = os.MkdirAll(path.Dir(CM().GetConfigFile()), 0700)
	if err != nil {
		return fmt.Errorf("Failed to create context config dir: %v", err)
	}

	return ioutil.WriteFile(CM().GetConfigFile(), contextYaml, 0600)
}

func (p *pxcConfigReaderWriter) load() (*Config, error) {
	if _, err := os.Stat(CM().GetConfigFile()); err != nil {
		// Does not exist
		return p.newDefaultConfig(), nil
	}

	data, err := ioutil.ReadFile(CM().GetConfigFile())
	if err != nil {
		return nil, fmt.Errorf("Failed to load config file %s, %v", CM().GetConfigFile(), err)
	}
	if len(data) == 0 {
		// Empty
		return p.newDefaultConfig(), nil
	}

	var configInfo *Config
	if err := yaml.Unmarshal(data, &configInfo); err != nil {
		return nil, fmt.Errorf("Failed to process config file %s, %v", CM().GetConfigFile(), err)
	}

	// Validate the config file
	if err := p.validate(configInfo); err != nil {
		return nil, err
	}

	return configInfo, nil
}

func (p *pxcConfigReaderWriter) newDefaultConfig() *Config {
	c := newConfig()
	c.CurrentContext = "default"
	c.AuthInfos[c.CurrentContext] = NewAuthInfo()
	c.AuthInfos[c.CurrentContext].Name = "default"
	c.Clusters[c.CurrentContext] = NewCluster()
	c.Clusters[c.CurrentContext].Name = "default"
	c.Clusters[c.CurrentContext].Endpoint = "127.0.0.1:9020"
	c.Contexts[c.CurrentContext] = &Context{
		Name:     "default",
		AuthInfo: "default",
		Cluster:  "default",
	}

	return c
}

func (p *pxcConfigReaderWriter) validate(c *Config) error {
	// REMOVED THESE LINES TO allow View for debugging. May have to check if these affect the rest of the system
	/*
		if len(c.CurrentContext) == 0 {
			return fmt.Errorf("Current context missing from config file %s", CM().GetConfigFile())
		}
	*/
	if c.AuthInfos == nil {
		c.AuthInfos = make(map[string]*AuthInfo)
	}
	if c.Clusters == nil {
		c.Clusters = make(map[string]*Cluster)
	}
	if c.Contexts == nil {
		c.Contexts = make(map[string]*Context)
	}

	/*
		if _, ok := c.Contexts[c.CurrentContext]; !ok {
			return fmt.Errorf("Context %s missing from config file %s", c.CurrentContext, CM().GetConfigFile())
		}

		currentCreds := c.Contexts[c.CurrentContext].AuthInfo
		if _, ok := c.AuthInfos[currentCreds]; !ok {
			return fmt.Errorf("Credentials %s missing from config file %s", currentCreds, CM().GetConfigFile())
		}

		currentCluster := c.Contexts[c.CurrentContext].Cluster
		if _, ok := c.Clusters[currentCluster]; !ok {
			return fmt.Errorf("Cluster %s missing from config file %s", currentCluster, CM().GetConfigFile())
		}

	*/
	return nil
}
