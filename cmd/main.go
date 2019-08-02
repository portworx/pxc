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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
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
	// TODO Redo plugin model: pm          *plugin.PluginManager

	// The $HOME/.px/plugins dir will be added at runtime
	pxPluginDefaultDirs = []string{
		"/var/lib/px/plugins",
		"/etc/pwx/plugins",
		"/opt/pwx/plugins",
		"/var/lib/porx/plugins",
	}

	varInitFncs []func()
	cmdInitFncs []func()
)

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file
func initConfig() {
	// If the cfgFile has not been setup in the arguments, then
	// read it from the HOME directory
	cfgFile = os.Getenv("PXCONFIG")
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

// GetConfigFile returns the current config file
func GetConfigFile() string {
	return cfgFile
}

// PxConnectDefault returns a Portworx client to the default or
// named context
func PxConnectDefault() (context.Context, *grpc.ClientConn, error) {
	// Global information will be set here, like forced context
	if len(cfgContext) == 0 {
		return portworx.PxConnectCurrent(cfgFile)
	} else {
		return portworx.PxConnectNamed(cfgFile, cfgContext)
	}
}

// KubeConnectDefault returns a Kubernetes client to the default
// or named context.
func KubeConnectDefault() (clientcmd.ClientConfig, *kclikube.Clientset, error) {
	return kubernetes.KubeConnect(cfgFile, cfgContext)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Main(); err != nil {
		util.Eprintf("%v\n", err)
		os.Exit(1)
	}
}

// RegisterCommandVar is used to register with px the initialization function
// for the command variable.
// Something must be returned to use the `var _ = ` trick.
func RegisterCommandVar(c func()) bool {
	varInitFncs = append(varInitFncs, c)

	return true
}

// RegisterCommandInit is used to register with px the initialization function
// for the command flags.
// Something must be returned to use the `var _ = ` trick.
func RegisterCommandInit(c func()) bool {
	cmdInitFncs = append(cmdInitFncs, c)
	return true
}

// Main starts the px cli
// Stupid simple initialization
func Main() error {
	// Setup all variables.
	// Setting up all the variables first will allow px
	// to initialize the init functions in any order
	for _, v := range varInitFncs {
		v()
	}

	// Call all plugin inits
	for _, f := range cmdInitFncs {
		f()
	}

	// Execute px
	return rootCmd.Execute()
}
