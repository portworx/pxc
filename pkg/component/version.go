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
package component

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"

	"github.com/spf13/cobra"
)

var versionCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			versionExec(cmd, args)
		},
	}
})

var _ = commander.RegisterCommandInit(func() {
	RootAddCommand(versionCmd)
})

func VersionAddCommand(cmd *cobra.Command) {
	versionCmd.AddCommand(cmd)
}

func versionExec(cmd *cobra.Command, args []string) {
	// Print client version
	fmt.Printf("%s Version: %s\n"+
		"Portworx SDK Version: %s\n",
		ComponentName,
		ComponentVersion,
		fmt.Sprintf("%d.%d.%d",
			api.SdkVersion_Major,
			api.SdkVersion_Minor,
			api.SdkVersion_Patch))
}
