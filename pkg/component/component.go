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
	"os"

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

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd     *cobra.Command
	rootOptions *rootFlags
	rootTmpl    = `Usage:{{if .Runnable}}
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
	// Component information
	ComponentName    string
	ComponentVersion string
)

// Component is the object managing the component
type Component struct {
	config ComponentConfig
}

// ComponentConfig provides the information necessary to create a component
type ComponentConfig struct {
	// Name of the component as it appears in the Usage of the help screen
	Name string

	// Version of the component
	Version string

	// Short description of the component
	Short string

	// If any flags need to be added to root of the component add them here
	RootFlags func(*cobra.Command)
}

// NewComponent creates a new component. Once created use Execute() to start the component
func NewComponent(config *ComponentConfig) *Component {
	ComponentName = config.Name
	ComponentVersion = config.Version
	return &Component{
		config: *config,
	}
}

// Execute starts the component
func (c *Component) Execute() {

	// Register the root component
	commander.RegisterCommandVar(func() {
		// rootCmd represents the base command when called without any subcommands
		rootOptions = &rootFlags{}
		rootCmd = &cobra.Command{
			Use:                c.config.Name,
			Short:              c.config.Short,
			SilenceUsage:       true,
			SilenceErrors:      true,
			PersistentPreRunE:  c.rootPersistentPreRunE,
			PersistentPostRunE: c.rootPersistentPostRunE,
			RunE:               rootCmdExec,
		}
	})

	// Register flags and other configuration
	commander.RegisterCommandInit(func() {

		// COMPONENT: Add the persistent pxc and kubectl flags
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

		if c.config.RootFlags != nil {
			c.config.RootFlags(rootCmd)
		}
	})

	// Start
	if err := c.main(); err != nil {
		util.Eprintf("%v\n", err)

		// Cobra will quit immediately and not run the PostRunE function on error
		// so we need to run it here.
		if rootCmd.PersistentPostRunE != nil {
			rootCmd.PersistentPostRunE(nil, []string{})
		}
		os.Exit(1)
	}
}

// Main starts the pxc cli
// Any initialization to pxc should be added to root.PersistentPreRunE
func (c *Component) main() error {
	// Setup flags
	commander.Setup()

	// Execute pxc
	return rootCmd.Execute()
}

func rootCmdExec(cmd *cobra.Command, args []string) error {
	if rootOptions.showOptions {
		rootCmd.SetUsageTemplate(globalsOnlyTmpl)
	}
	rootCmd.Usage()
	return nil
}

// RootAddCommand is called by handlers to add the top level options to
// the program
func RootAddCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)
}

func (c *Component) rootPersistentPreRunE(cmd *cobra.Command, args []string) error {

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
	logrus.Infof("%s version: %s", c.config.Name, c.config.Version)

	return nil
}

func (c *Component) rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	// Close the global Portworx API tunnel if any
	kubernetes.StopTunnel()

	return nil
}
