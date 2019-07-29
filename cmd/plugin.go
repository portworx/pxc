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
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

// pluginsCmd represents the getPlugins command
var pluginsCmd = &cobra.Command{
	Use:     "plugin",
	Aliases: []string{"plugins"},
	Short:   "Display px plugin information",
	Run: func(cmd *cobra.Command, args []string) {
		pluginsExec(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(pluginsCmd)
}

func pluginsExec(cmd *cobra.Command, args []string) {

	// Get the list of the plugins from the PluginManager
	if len(pm.List()) == 0 {
		util.Printf("No plugins installed")
		return
	}

	t := util.NewTabby()
	t.AddHeader("Name", "Version", "Location", "Description")
	for _, p := range pm.List() {
		t.AddLine(p.Name, p.Version, p.Location, p.Description)
	}
	t.Print()
}
