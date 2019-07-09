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
package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"plugin"

	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

// TODO: Add comments

type PluginManifest struct {
	Name        string
	Version     string
	Location    string
	Description string
}

type PluginManagerConfig struct {
	PluginDirs []string
	RootCmd    *cobra.Command
}

type PluginManager struct {
	config  PluginManagerConfig
	plugins []*PluginManifest
}

const (
	pluginExt = ".px"
)

func NewPluginManager(config *PluginManagerConfig) *PluginManager {
	return &PluginManager{
		config:  *config,
		plugins: make([]*PluginManifest, 0),
	}
}

func (pm *PluginManager) Load() {

	for _, pluginDir := range pm.config.PluginDirs {
		files, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) == pluginExt {
				manifest, err := pm.loadPlugin(path.Join(pluginDir, file.Name()))
				if err != nil {
					util.Eprintf("%v\n", err)
					continue
				}
				pm.plugins = append(pm.plugins, manifest)
			}
		}
	}
}

func (pm *PluginManager) List() []*PluginManifest {
	return pm.plugins
}

func (pm *PluginManager) loadPlugin(soPath string) (*PluginManifest, error) {
	// Open plugin library
	p, err := plugin.Open(soPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open plugin %s: %v\n", soPath, err)
	}

	manifest, err := pm.getManifest(p, soPath)
	if err != nil {
		return nil, err
	}

	err = pm.registerPlugin(p, soPath)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}

func (pm *PluginManager) getManifest(p *plugin.Plugin, soPath string) (*PluginManifest, error) {
	// Get Plugin Manifest to get name and info
	m, err := p.Lookup("PluginManifest")
	if err != nil {
		return nil, fmt.Errorf("Plugin %s does not have init function", soPath)
	}

	// Confirm the interface is correct
	mmap, ok := m.(*map[string]string)
	if !ok {
		return nil, fmt.Errorf("Plugin %s failed to get manifest", soPath)
	}
	manifestMap := *mmap

	// Populate the manifest
	manifest := &PluginManifest{}
	if name, ok := manifestMap["name"]; !ok {
		return nil, fmt.Errorf("%s plugin is missing a required name", soPath)
	} else {
		manifest.Name = name
	}
	if version, ok := manifestMap["version"]; !ok {
		return nil, fmt.Errorf("%s plugin is missing a required version", soPath)
	} else {
		manifest.Version = version
	}
	manifest.Description = manifestMap["description"]
	manifest.Location = soPath

	return manifest, nil
}

func (pm *PluginManager) registerPlugin(p *plugin.Plugin, soPath string) error {
	// Get access to plugin init function
	f, err := p.Lookup("PluginInit")
	if err != nil {
		return fmt.Errorf("Plugin %s does not have init function", soPath)
	}

	// Confirm the interface is correct
	pinit, ok := f.(func(*cobra.Command, *os.File, *os.File))
	if !ok {
		return fmt.Errorf("Plugin %s does not have the correct init function", soPath)
	}

	// Finally, register the handler
	pinit(pm.config.RootCmd, util.Stdout, util.Stderr)
	return nil
}
