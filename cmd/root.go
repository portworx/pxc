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
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:           "px",
		Short:         "Portworx command line tool",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
})

var _ = RegisterCommandInit(func() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+pxDefaultDir+"/"+pxDefaultConfigName+")")
	rootCmd.PersistentFlags().StringVar(&cfgContext, "context", "", "Force context name for the command")

	// TODO: move these flags out of persistent
	rootCmd.PersistentFlags().StringP("selector", "l", "", "Comma separated label selector of the form 'key=value,key=value'")

	// Global cobra configurations
	rootCmd.Flags().SortFlags = false

	// Load plugins
	/* TODO: Redo Plugin model
	home, _ := homedir.Dir()
	pxPluginDefaultDirs = append(pxPluginDefaultDirs,
		path.Join(home, pxDefaultDir, "plugins"))
	pm = plugin.NewPluginManager(&plugin.PluginManagerConfig{
		PluginDirs: pxPluginDefaultDirs,
		RootCmd:    rootCmd,
	})
	pm.Load()
	*/
})
