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

type rootFlags struct {
	showOptions bool
}

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd     *cobra.Command
	rootOptions *rootFlags

	// This template allow pxc to override the way it prints out the help. This template
	// allows pxc to not print out global options unless requested.
	rootTmpl = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalNonPersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
Use "pxc --options" for a list of global command-line options (applies to all commands)
`
	globalsOnlyTmpl = `Global Flags:
{{.PersistentFlags.FlagUsages | trimTrailingWhitespaces}}
`
)

func RootAddCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)
}

var _ = commander.RegisterCommandVar(func() {
	rootOptions = &rootFlags{}
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "pxc",
		Short: "Portworx client",
		Long: `pxc is a Portworx client which allows users to communicate with their storage cluster
from their client machine combining information from both Kubernetes and Portworx.

Please see https://docs.portworx.com/reference/ for more information.`,
		SilenceUsage:       true,
		SilenceErrors:      true,
		PersistentPreRunE:  rootPersistentPreRunE,
		PersistentPostRunE: rootPersistentPostRunE,
		RunE:               rootCmdExec,
	}
})

var _ = commander.RegisterCommandInit(func() {

	// Add persistent flags
	if util.InKubectlPluginMode() {
		// As kubectl plugin mode
		config.KM().ConfigFlags().AddFlags(rootCmd.PersistentFlags())
		config.CM().GetFlags().AddFlagsPluginMode(rootCmd.PersistentFlags())
	} else {
		// Not in plugin mode
		config.CM().GetFlags().AddFlags(rootCmd.PersistentFlags())
	}
	rootCmd.Flags().BoolVar(&rootOptions.showOptions, "options", false, "Show global options for all commands")
	rootCmd.SetUsageTemplate(rootTmpl)
})

func rootCmdExec(cmd *cobra.Command, args []string) error {
	if rootOptions.showOptions {
		rootCmd.SetUsageTemplate(globalsOnlyTmpl)
	}
	rootCmd.Usage()
	return nil
}

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

	// Set version
	logrus.Infof("pxc version: %s", PxVersion)

	// The following commands do not need to load configuration
	switch cmd.Name() {
	case "version":
		return nil
	}

	// Load configuration information from disk if any
	if err := config.CM().Load(); err != nil {
		return err
	}

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	// Close the global tunnel if any
	kubernetes.StopTunnel()

	return nil
}
