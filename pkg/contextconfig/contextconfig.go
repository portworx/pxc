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

// ContextConfig provides information about the px context information
type ContextConfig struct {
	Current        string          `json:"current" yaml:"current"`
	Configurations []ClientContext `json:"configurations" yaml:"configurations"`
}

// ConfigReference is a reference to a ContextConfig and the path associated with it
type ConfigReference struct {
	path string
	cfg  *ContextConfig
}

// NewContextConfig
func NewConfigReference(configFile string) *ConfigReference {
	return &ConfigReference{
		path: configFile,
	}
}

// TODO: Return error if exists already
// TODO: Add Update support
func (cr *ConfigReference) Add(cctx *ClientContext) error {
	var ctx *ContextConfig

	ctx, _ = cr.loadContext()
	if ctx == nil {
		ctx = new(ContextConfig)
		ctx.Configurations = []ClientContext{*cctx}
		ctx.Current = cctx.Name
	} else {
		ctx.Configurations = append(ctx.Configurations, *cctx)
	}

	return cr.saveContext(ctx)
}

func (cr *ConfigReference) GetCurrent() (*ClientContext, error) {
	ctx, err := cr.loadContext()
	if err != nil {
		return nil, err
	}

	if len(ctx.Configurations) == 0 {
		return nil, fmt.Errorf("No configurations found in %s", cr.path)
	}

	if len(ctx.Current) == 0 {
		return &ctx.Configurations[0], nil
	}

	for _, cctx := range ctx.Configurations {
		if cctx.Name == ctx.Current {
			return &cctx, nil
		}
	}

	return nil, fmt.Errorf("Default context %s not found in %s",
		ctx.Current,
		cr.path)
}

// TODO have GetNamedContext and GetCurrent() call a common helper function
func (cr *ConfigReference) GetNamedContext(name string) (*ClientContext, error) {
	ctx, err := cr.loadContext()
	if err != nil {
		return nil, err
	}

	if len(ctx.Configurations) == 0 {
		return nil, fmt.Errorf("No configurations found in %s", cr.path)
	}

	for _, cctx := range ctx.Configurations {
		if cctx.Name == name {
			return &cctx, nil
		}
	}

	return nil, fmt.Errorf("Context %s not found in %s",
		ctx.Current,
		cr.path)

}

func (cr *ConfigReference) GetAll() (*ContextConfig, error) {
	return cr.loadContext()
}

// TODO:
func (cr *ConfigReference) Remove(cctx *ClientContext) error {
	return nil
}

func (cr *ConfigReference) Set(cctx *ClientContext) error {
	return nil
}

func (cr *ConfigReference) UnSet(cctx *ClientContext) error {
	return nil
}

// TODO: below
// GetContext loads the context
/*
func GetContext(contextFile string) (*ContextConfig, error) {
	//envToken := os.Getenv("PX_AUTH_TOKEN")
	//envContextName := os.Getenv("PX_CONTEXT_NAME")

	// Here we read context/auth input based on the following order
	//
	// Precedence for context/auth input, from highest to lowest
	// 1. env var for token
	// 2. env var for context
	// 3. --context flag
	// 4. currentcontext in PxContextCfg

	// Second, we handle context flag or current in contextconfig
	/*
		currentContext := PxContextCfg.Current
		if contextFlag != "" {
			currentContext = contextFlag
		}

		if currentContext != "" {
			for _, ccfg := range PxContextCfg.Configurations {
				if ccfg.Name == currentContext {
					return ccfg, nil
				}
			}

			return nil, fmt.Errorf("Context %s does not exist", currentContext)
		}
	if _, err := os.Stat(contextFile); err == nil {
		c, err := loadContext(contextFile)
		if err != nil {
			return nil, err
		}
		return c, nil
	}

	return nil, fmt.Errorf("Context does not exist, please use 'px context create' to create one")
}

*/

/*
// UpdateCurrentContext loads the context file and sets a new
// currentcontext name.
func UpdateCurrentContext(contextName string) error {
	if err := GetContext(contextName); err != nil {
		return err
	}

	// if contextName is empty, clear current name
	if contextName == "" {
		PxContextCfg.Current = ""
		if err := SaveContext(PxContextCfg, PxContextFile); err != nil {
			return err
		}
		return nil
	}

	// only set if context name exists
	for _, ac := range PxContextCfg.Configurations {
		if acr.Name == contextName {
			PxContextCfg.Current = contextName
			if err := SaveContext(PxContextCfg, PxContextFile); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("context %s does not exist.", contextName)
}

*/

func (cr *ConfigReference) saveContext(ctx *ContextConfig) error {
	if ctx == nil || cr.path == "" {
		return fmt.Errorf("Failed to save context config data. Invalid data...")
	}

	contextYaml, err := yaml.Marshal(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create yaml parse: %v", err)
	}

	// Create the contextconfig location
	err = os.MkdirAll(path.Dir(cr.path), 0700)
	if err != nil {
		return fmt.Errorf("Failed to create context config dir: %v", err)
	}

	return ioutil.WriteFile(cr.path, contextYaml, 0600)
}

func (cr *ConfigReference) loadContext() (*ContextConfig, error) {
	var contextCfg ContextConfig

	if _, err := os.Stat(cr.path); err != nil {
		return nil, fmt.Errorf("Context does not exist, please use 'px context create' to create one")
	}

	data, err := ioutil.ReadFile(cr.path)
	if err != nil {
		return nil, fmt.Errorf("Failed to load context config file, %v", err)
	}

	if err := yaml.Unmarshal(data, &contextCfg); err != nil {
		return nil, fmt.Errorf("Failed to process context config data, %v", err)
	}

	return &contextCfg, err
}
