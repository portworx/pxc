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
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

const (
	pxDefaultDir        = ".pxc"
	pxDefaultConfigName = "config.yml"

	Ki = 1024
	Mi = 1024 * Ki
	Gi = 1024 * Mi
	Ti = 1024 * Gi
)

var (
	cfgDir      string
	cfgFile     string
	cfgContext  string
	optEndpoint string
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
	// If the cfgFile has not been setup in the arguments, then
	// read it from the HOME directory
	cfgFile = os.Getenv("PXCONFIG")
	if len(cfgFile) == 0 {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			util.Eprintf("Error: %v\n", err)
			os.Exit(1)
		}
		cfgFile = path.Join(home, pxDefaultDir, pxDefaultConfigName)
	}

	// Save configurations
	config.Set(config.SpecifiedContext, cfgContext)
	config.Set(config.File, cfgFile)
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
		os.Exit(1)
	}
}

// Main starts the pxc cli
// Stupid simple initialization
func Main() error {
	commander.Setup()

	// Execute px
	return rootCmd.Execute()
}
