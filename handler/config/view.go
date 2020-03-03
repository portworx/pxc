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
package configcli

import (
	"fmt"

	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

// viewCmd represents the config command
var viewCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	viewCmd = &cobra.Command{
		Use:   "view",
		Short: "Show pxc configuration",
		RunE:  viewExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	if !util.InKubectlPluginMode() {
		ConfigAddCommand(viewCmd)
	}
})

func viewExec(cmd *cobra.Command, args []string) error {
	configInfo, err := config.CM().ConfigLoad()
	if err != nil {
		return fmt.Errorf("Unable to read configuration file %s: %v",
			config.CM().GetConfigFile(), err)
	}

	// Do not show any private data
	for _, user := range configInfo.AuthInfos {
		if len(user.Token) != 0 {
			user.Token = "<REDACTED>"
		}
	}

	for _, cluster := range configInfo.Clusters {
		if len(cluster.CACertData) != 0 {
			cluster.CACertData = []byte("<REDACTED>")
		}
	}

	util.PrintYaml(configInfo)
	return nil
}
