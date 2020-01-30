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
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd *cobra.Command
	//kubePortForwarder kubernetes.PortForwarder
)

func RootAddCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:                "pxc",
		Short:              "Portworx client",
		SilenceUsage:       true,
		SilenceErrors:      true,
		PersistentPreRunE:  rootPersistentPreRunE,
		PersistentPostRunE: rootPersistentPostRunE,
	}
})

var _ = commander.RegisterCommandInit(func() {

	// Add persistent flags
	if util.InKubectlPluginMode() {
		// As kubectl plugin mode
		config.KM().AddFlags(rootCmd.PersistentFlags())
		config.CM().GetFlags().AddFlagsPluginMode(rootCmd.PersistentFlags())
	} else {
		// Not in plugin mode
		config.CM().GetFlags().AddFlags(rootCmd.PersistentFlags())
	}
})

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {

	// Setup verbosity
	switch config.CM().GetFlags().Verbosity {
	case 0:
		logrus.SetLevel(logrus.FatalLevel)
	case 1:
		logrus.SetLevel(logrus.WarnLevel)
	case 2:
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Load configuration information from disk if any
	if err := config.CM().Load(); err != nil {
		return err
	}

	// Set verbosity
	logrus.Infof("pxc version: %s", PxVersion)

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	// Close the global tunnel if any
	kubernetes.StopTunnel()

	return nil
}
