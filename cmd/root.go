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
	"fmt"

	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd           *cobra.Command
	kubePortForwarder kubernetes.PortForwarder
)

func RootAddCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:                "pxc",
		Short:              "Portworx command line tool",
		SilenceUsage:       true,
		SilenceErrors:      true,
		PersistentPreRunE:  rootPersistentPreRunE,
		PersistentPostRunE: rootPersistentPostRunE,
	}

	// Initialize persistent flag objects
	kubernetes.KubeCliOpts = genericclioptions.NewConfigFlags(true)
	config.CliOpts = config.NewConfigFlags()
})

var _ = commander.RegisterCommandInit(func() {

	// Add persistent flags
	if kubernetes.InKubectlPluginMode() {
		// As kubectl plugin mode
		kubernetes.KubeCliOpts.AddFlags(rootCmd.PersistentFlags())
		config.CliOpts.AddFlagsPluginMode(rootCmd.PersistentFlags())
	} else {
		// Not in plugin mode
		config.CliOpts.AddFlags(rootCmd.PersistentFlags())
	}
})

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {

	// Setup verbosity
	switch config.CliOpts.Verbosity {
	case 0:
		logrus.SetLevel(logrus.PanicLevel)
	case 1:
		logrus.SetLevel(logrus.FatalLevel)
	case 2:
		logrus.SetLevel(logrus.WarnLevel)
	case 3:
		logrus.SetLevel(logrus.InfoLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Set verbosity
	logrus.Infof("pxc version: %s", PxVersion)

	// Setup port forwarding if running as a kubectl plugin
	if kubernetes.InKubectlPluginMode() {
		logrus.Info("Kubectl plugin mode detected")
		kubePortForwarder = kubernetes.NewKubectlPortForwarder(*kubernetes.KubeCliOpts.KubeConfig)
		if err := kubePortForwarder.Start(); err != nil {
			return fmt.Errorf("Failed to setup port forward: %v", err)
		}
		config.Set(config.PluginEndpoint, kubePortForwarder.Endpoint())
	}

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	if kubePortForwarder != nil {
		kubePortForwarder.Stop()
	}

	return nil
}
