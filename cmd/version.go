/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"

	"github.com/spf13/cobra"
)

var (
	PxVersion = "(DEV)"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show px version information",
	Run: func(cmd *cobra.Command, args []string) {
		versionExec(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionExec(cmd *cobra.Command, args []string) {
	// Print client version
	fmt.Printf("Client Version: %s\n"+
		"Client SDK Version: %s\n",
		PxVersion,
		fmt.Sprintf("%d.%d.%d",
			api.SdkVersion_Major,
			api.SdkVersion_Minor,
			api.SdkVersion_Patch))
}
