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
	"os"

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

	kubernetes.KubeCliOpts = genericclioptions.NewConfigFlags(true)
})

var _ = commander.RegisterCommandInit(func() {

	rootCmd.PersistentFlags().Int32Var(&verbosity, "v", 0, "[0-4] Log level verbosity")

	// Add the kubernetes flags
	if kubernetes.InKubectlPluginMode() {
		kubernetes.KubeCliOpts.AddFlags(rootCmd.PersistentFlags())
	} else {
		rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/"+pxDefaultDir+"/"+pxDefaultConfigName+")")
		rootCmd.PersistentFlags().StringVar(&cfgContext, "context", "", "Force context name for the command")
	}

	// Global cobra configurations
	rootCmd.Flags().SortFlags = false
})

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {

	// Setup verbosity
	switch verbosity {
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

	// Setup port forwarding if running as a kubectl plugin
	if kubernetes.InKubectlPluginMode() {
		logrus.Info("Kubectl plugin mode detected")
		kubeConfig := *kubernetes.KubeCliOpts.KubeConfig
		if len(kubeConfig) == 0 {
			kubeConfig = os.Getenv("KUBECONFIG")
		}
		if len(kubeConfig) == 0 {
			return fmt.Errorf("KUBECONFIG or --kubeconfig must be defined (for now until we add support to do this automically)")
		}
		logrus.Infof("Using Kubeconfig: %s", kubeConfig)
		kubePortForwarder = kubernetes.NewKubectlPortForwarder(kubeConfig)
		if err := kubePortForwarder.Start(); err != nil {
			return fmt.Errorf("Failed to setup port forward: %v", err)
		}
		config.Set(config.PluginEndpoint, kubePortForwarder.Endpoint())
	}

	// Set verbosity
	logrus.Infof("pxc version: %s", PxVersion)

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	if kubePortForwarder != nil {
		kubePortForwarder.Stop()
	}

	return nil
}
