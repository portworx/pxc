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
	Context    string        `json:"context" yaml:"context"`
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

type ContextConfigFile struct {
	ctxFile string
	ctx     *ContextConfig
}

func NewContextConfig(contextFile string) *ContextConfigFile {
	return &ContextConfigFile{
		ctxFile: contextFile,
	}
}

// TODO: Return error if exists already
// TODO: Add Update support
func (c *ContextConfigFile) Add(cctx *ClientContext) error {
	var ctx *ContextConfig

	ctx, _ = c.loadContext()
	if ctx == nil {
		ctx = new(ContextConfig)
		ctx.Configurations = []ClientContext{*cctx}
		ctx.Current = cctx.Context
	} else {
		ctx.Configurations = append(ctx.Configurations, *cctx)
	}

	return c.saveContext(ctx)
}

//TODO: Add GetWithContext to get non-default
func (c *ContextConfigFile) Get() (*ClientContext, error) {
	ctx, err := c.loadContext()
	if err != nil {
		return nil, err
	}

	if len(ctx.Configurations) == 0 {
		return nil, fmt.Errorf("No configurations found in %s", c.ctxFile)
	}

	if len(ctx.Current) == 0 {
		return &ctx.Configurations[0], nil
	}

	for _, cctx := range ctx.Configurations {
		if cctx.Context == ctx.Current {
			return &cctx, nil
		}
	}

	return nil, fmt.Errorf("Default context %s not found in %s",
		ctx.Current,
		c.ctxFile)
}

func (c *ContextConfigFile) GetAll() (*ContextConfig, error) {
	return c.loadContext()
}

// TODO:
func (c *ContextConfigFile) Remove(cctx *ClientContext) error {
	return nil
}

func (c *ContextConfigFile) Set(cctx *ClientContext) error {
	return nil
}

func (c *ContextConfigFile) UnSet(cctx *ClientContext) error {
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
				if ccfg.Context == currentContext {
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
		if ac.Context == contextName {
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

func (c *ContextConfigFile) saveContext(ctx *ContextConfig) error {
	if ctx == nil || c.ctxFile == "" {
		return fmt.Errorf("Failed to save context config data. Invalid data...")
	}

	contextYaml, err := yaml.Marshal(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create yaml parse: %v", err)
	}

	// Create the contextconfig location
	err = os.MkdirAll(path.Dir(c.ctxFile), 0700)
	if err != nil {
		return fmt.Errorf("Failed to create context config dir: %v", err)
	}

	return ioutil.WriteFile(c.ctxFile, contextYaml, 0600)
}

func (c *ContextConfigFile) loadContext() (*ContextConfig, error) {
	var contextCfg ContextConfig

	if _, err := os.Stat(c.ctxFile); err != nil {
		return nil, fmt.Errorf("Context does not exist, please use 'px context create' to create one")
	}

	data, err := ioutil.ReadFile(c.ctxFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to load context config file, %v", err)
	}

	if err := yaml.Unmarshal(data, &contextCfg); err != nil {
		return nil, fmt.Errorf("Failed to process context config data, %v", err)
	}

	return &contextCfg, err
}
