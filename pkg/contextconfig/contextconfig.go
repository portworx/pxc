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
package contextconfig

/*
import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	yaml "gopkg.in/yaml.v2"
)

// ClientTLSInfo provide client TLS configuration information
type ClientTLSInfo struct {
	Cacert string `json:"cacert" yaml:"cacert"`
}

// Identity contains token unqiue ID information
type Identity struct {
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Email   string `json:"email,omitempty" yaml:"email,omitempty"`
}

// ClientContext provides information about the client context
type ClientContext struct {
	Name       string        `json:"context" yaml:"context"`
	Token      string        `json:"token" yaml:"token"`
	Identity   Identity      `json:"identity,omitempty" yaml:"identity,omitempty"`
	Error      string        `json:"error,omitempty" yaml:"error,omitempty"`
	TlsData    ClientTLSInfo `json:"tlsdata,omitempty" yaml:"tlsdata,omitempty"`
	Endpoint   string        `json:"endpoint" yaml:"endpoint"`
	Secure     bool          `json:"secure" yaml:"secure"`
	Kubeconfig string        `json:"kubeconfig" yaml:"kubeconfig"`
}

// ContextConfig provides information about the pxc context information
type ContextConfig struct {
	Current        string          `json:"current" yaml:"current"`
	Configurations []ClientContext `json:"configurations" yaml:"configurations"`
}

// ContextManager is a reference to a ContextConfig and the path associated with it
type ContextManager struct {
	path string
	cfg  *ContextConfig
}

// New returns an empty, unloaded context manager
func New(configFile string) *ContextManager {
	return &ContextManager{
		path: configFile,
	}
}

// GetContextManager loads an in memory reference of the Context Configuration file from disk.
// This reference is the primary object to use when managing the user's context configuration.
func NewContextManager(configFile string) (*ContextManager, error) {
	if configFile == "" {
		return nil, fmt.Errorf("Invalid configuration file path. Path must be non-empty.")
	}

	cm := &ContextManager{
		path: configFile,
	}

	err := cm.loadContext()
	if err != nil {
		return cm, err
	}

	return cm, nil
}

// Add inserts a given clientContext into cm.cfg.Configurations, and
// saves the context.
func (cm *ContextManager) Add(clientContext *ClientContext) error {
	if cm.cfg == nil {
		cm.cfg = new(ContextConfig)
		cm.cfg.Configurations = make([]ClientContext, 0)
	}

	if cm.cfg.Current == "" {
		cm.cfg.Current = clientContext.Name
	}

	// Check for existing configuration. If one exists,
	// update it. Otherwise, we will append a new configuration.
	for i := range cm.cfg.Configurations {
		if cm.cfg.Configurations[i].Name == clientContext.Name {
			cm.cfg.Configurations[i] = *clientContext
			return cm.saveContext()
		}
	}

	cm.cfg.Configurations = append(cm.cfg.Configurations, *clientContext)
	return cm.saveContext()
}

func (cm *ContextManager) getContextDetails(name string) (*ClientContext, error) {
	if err := cm.loadContext(); err != nil {
		return nil, err
	}

	if len(cm.cfg.Configurations) == 0 {
		return nil, fmt.Errorf("No configurations found in %s", cm.path)
	}

	if len(name) == 0 {
		if len(cm.cfg.Current) == 0 {
			return &cm.cfg.Configurations[0], nil
		}
		name = cm.cfg.Current
	}

	for _, cctx := range cm.cfg.Configurations {
		if cctx.Name == name {
			return &cctx, nil
		}
	}

	return nil, fmt.Errorf("Context %s not found in %s",
		cm.cfg.Current,
		cm.path)
}

func (cm *ContextManager) GetCurrent() (*ClientContext, error) {
	ctx, err := cm.getContextDetails("")
	return ctx, err
}

func (cm *ContextManager) GetNamedContext(name string) (*ClientContext, error) {
	ctx, err := cm.getContextDetails(name)
	return ctx, err
}

// GetAll simply returns all configurations
func (cm *ContextManager) GetAll() *ContextConfig {
	return cm.cfg
}

// contains checks if a context name exists in a given config reference.
func (cm *ContextManager) contains(contextName string) bool {
	for _, ccfg := range cm.cfg.Configurations {
		if ccfg.Name == contextName {
			return true
		}
	}

	return false
}

// Remove deletes a context from the configurations list
func (cm *ContextManager) Remove(nameToDelete string) error {
	removed := false
	for i, acfg := range cm.cfg.Configurations {
		if nameToDelete != "" && acfg.Name == nameToDelete {
			cm.cfg.Configurations = append(cm.cfg.Configurations[:i],
				cm.cfg.Configurations[i+1:]...)
			if acfg.Name == cm.cfg.Current {
				cm.cfg.Current = ""
			}
			removed = true
			break
		}
	}

	if !removed {
		return fmt.Errorf("Context does not exist")
	}

	return cm.saveContext()
}

func (cm *ContextManager) SetCurrent(name string) error {
	// Unset is performed when no name is passed in.
	if name == "" {
		cm.cfg.Current = ""
		return cm.saveContext()
	}

	// Otherwise, we set the context to the given name if it exist.
	if cm.contains(name) {
		cm.cfg.Current = name
	} else {
		return fmt.Errorf("Context %s does not exist", name)
	}

	return cm.saveContext()
}

// GetContext loads the context by name (if provided), or the current context set.
func (cm *ContextManager) GetContext(contextName string) (*ClientContext, error) {
	// Here we read context/auth input based on the following order
	//
	// Precedence for context/auth input, from highest to lowest
	// 1. env var for token
	// 2. env var for context
	// 3. --context flag
	// 4. currentcontext in cm.cfg

	// Second, we handle context flag or current in contextconfig
	currentContext := cm.cfg.Current
	if contextName != "" {
		currentContext = contextName
	}

	if currentContext != "" {
		for _, ccfg := range cm.cfg.Configurations {
			if ccfg.Name == currentContext {
				return &ccfg, nil
			}
		}

		return nil, fmt.Errorf("Context %s does not exist", currentContext)
	}

	return nil, fmt.Errorf("Context does not exist, please use 'px context create' to create one")
}

// UpdateCurrentContext loads the context file and sets a new
// currentcontext name.
func (cm *ContextManager) UpdateCurrentContext(contextName string) error {

	// if contextName is empty, clear current name
	if contextName == "" {
		cm.cfg.Current = ""
		if err := cm.saveContext(); err != nil {
			return err
		}
		return nil
	}

	// only set if context name exists
	for _, clientConfig := range cm.cfg.Configurations {
		if clientConfig.Name == contextName {
			cm.cfg.Current = contextName
			if err := cm.saveContext(); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("context %s does not exist.", contextName)
}

// saveContext saves the currently set configuration to the given file reference location.
func (cm *ContextManager) saveContext() error {
	if cm.cfg == nil || cm.path == "" {
		return fmt.Errorf("Failed to save context config data. Invalid data...")
	}

	contextYaml, err := yaml.Marshal(cm.cfg)
	if err != nil {
		return fmt.Errorf("Failed to create yaml parse: %v", err)
	}

	// Create the contextconfig location
	err = os.MkdirAll(path.Dir(cm.path), 0700)
	if err != nil {
		return fmt.Errorf("Failed to create context config dir: %v", err)
	}

	return ioutil.WriteFile(cm.path, contextYaml, 0600)
}

// loadContext loads configuration data located at cm.path into cm.cfg
func (cm *ContextManager) loadContext() error {
	if _, err := os.Stat(cm.path); err != nil {
		return fmt.Errorf("Context does not exist, please use 'px context create' to create one")
	}

	data, err := ioutil.ReadFile(cm.path)
	if err != nil {
		return fmt.Errorf("Failed to load context config file, %v", err)
	}
	if len(data) == 0 {
		return fmt.Errorf("The config file is empty")
	}

	if err := yaml.Unmarshal(data, &cm.cfg); err != nil {
		return fmt.Errorf("Failed to process context config data, %v", err)
	}

	return nil
}
*/
