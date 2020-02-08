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
package cmd

import (
	"os"

	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/plugin"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

const (
	Ki = 1024
	Mi = 1024 * Ki
	Gi = 1024 * Mi
	Ti = 1024 * Gi
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
	// Save configurations
	config.Set(config.SpecifiedContext, config.CM().GetFlags().Context)
	// TODO -- change to use CM()
	config.Set(config.File, config.CM().GetConfigFile())
}

// GetConfigFile returns the current config file
func GetConfigFile() string {
	return config.Get(config.File)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Main(); err != nil {
		util.Eprintf("%v\n", err)

		// Cobra will quit immediately and not run the PostRunE function on error
		// so we need to run it here.
		if rootCmd.PersistentPostRunE != nil {
			rootCmd.PersistentPostRunE(nil, []string{})
		}
		os.Exit(1)
	}
}

// Main starts the pxc cli
// Any initialization to pxc should be added to root.PersistentPreRunE
func Main() error {
	// Setup flags
	commander.Setup()

	// Search for plugins and only execute if a command is not found
	if len(os.Args) > 1 {
		cmdPathPieces := os.Args[1:]
		if _, _, err := rootCmd.Find(cmdPathPieces); err != nil {
			pluginHandler := NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes)
			if err = HandlePluginCommand(pluginHandler, cmdPathPieces); err != nil {
				return err
			}
		}
	}

	// Execute pxc
	return rootCmd.Execute()
}
