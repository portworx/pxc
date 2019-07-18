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
	"context"
	"os"
	"path"

	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/plugin"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	pxDefaultDir        = ".px"
	pxDefaultConfigName = "config.yml"

	Ki = 1024
	Mi = 1024 * Ki
	Gi = 1024 * Mi
	Ti = 1024 * Gi
)

var (
	cfgDir      string
	cfgFile     string
	cfgContext  string
	optEndpoint string
	pm          *plugin.PluginManager

	// The $HOME/.px/plugins dir will be added at runtime
	pxPluginDefaultDirs = []string{
		"/var/lib/px/plugins",
		"/etc/pwx/plugins",
		"/opt/pwx/plugins",
		"/var/lib/porx/plugins",
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "px",
	Short: "Portworx command line tool",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.Eprintf("%v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+pxDefaultDir+"/"+pxDefaultConfigName+")")
	rootCmd.PersistentFlags().StringVar(&cfgContext, "context", "", "Force context name for the command")

	// TODO: move these flags out of persistent
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output in yaml|json|wide")
	rootCmd.PersistentFlags().Bool("show-labels", false, "Show labels in the last column of the output")
	rootCmd.PersistentFlags().StringP("selector", "l", "", "Comma separated label selector of the form 'key=value,key=value'")

	// Global cobra configurations
	rootCmd.Flags().SortFlags = false

	// Load plugins
	home, _ := homedir.Dir()
	pxPluginDefaultDirs = append(pxPluginDefaultDirs,
		path.Join(home, pxDefaultDir, "plugins"))
	pm = plugin.NewPluginManager(&plugin.PluginManagerConfig{
		PluginDirs: pxPluginDefaultDirs,
		RootCmd:    rootCmd,
	})
	pm.Load()
}

// initConfig reads in config file
func initConfig() {
	// If the cfgFile has not been setup in the arguments, then
	// read it from the HOME directory
	if len(cfgFile) == 0 {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			util.Eprintf("Error: %v\n", err)
			os.Exit(1)
		}
		cfgFile = path.Join(home, pxDefaultDir, pxDefaultConfigName)
	}
}

func GetConfigFile() string {
	return cfgFile
}

func PxConnectDefault() (context.Context, *grpc.ClientConn, error) {
	// Global information will be set here, like forced context
	if len(cfgContext) == 0 {
		return portworx.PxConnectCurrent(cfgFile)
	} else {
		return portworx.PxConnectNamed(cfgFile, cfgContext)
	}
}

func KubeConnectDefault() (clientcmd.ClientConfig, *kclikube.Clientset, error) {
	return kubernetes.KubeConnect(cfgFile, cfgContext)
}
