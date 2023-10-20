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
package config

import (
	"os"
	"path"

	"github.com/spf13/pflag"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

const (
	SpecifiedContext    = "cfgcontext"
	File                = "cfgfile"
	PxDefaultDir        = ".pxc"
	PxDefaultConfigName = "config.yml"

	flagPrefix          = "pxc."
	flagConfigFile      = flagPrefix + "config"
	flagConfigDir       = flagPrefix + "config-dir"
	flagContext         = flagPrefix + "context"
	flagSecretNamespace = flagPrefix + "secret-namespace"
	flagSecretName      = flagPrefix + "secret-name"
	flagToken           = flagPrefix + "token"
	flagVerbosity       = flagPrefix + "v"
)

type ConfigFlags struct {
	ConfigDir       string
	ConfigFile      string
	Context         string
	SecretNamespace string
	SecretName      string
	Token           string
	Verbosity       int32
}

func newConfigFlags() *ConfigFlags {
	// If the cfgFile has not been setup in the arguments, then
	// read it from the HOME directory

	// Check env
	configDir := os.Getenv("PXCONFIGDIR")
	if len(configDir) == 0 {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatalf("unable to determine home directory: %v\n", err)
		}
		configDir = path.Join(home, PxDefaultDir)
	}

	return &ConfigFlags{
		ConfigDir: configDir,
	}
}

func (c *ConfigFlags) GetConfigFile() string {
	// Check env
	envCF := os.Getenv("PXCONFIG")
	if len(envCF) != 0 {
		c.ConfigFile = envCF
		return c.ConfigFile
	}

	if len(c.ConfigFile) == 0 {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatalf("unable to determine home directory: %v\n", err)
		}
		c.ConfigFile = path.Join(home, PxDefaultDir, PxDefaultConfigName)
	}

	return c.ConfigFile
}

// AddFlags adds the appropriate global flags for non-kubectl plugin mode
func (c *ConfigFlags) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.Context, flagContext, c.Context, "Force context name for the command")
	c.addFlagsCommon(flags)
}

// AddFlagsPluginMode adds the appropriate global flags when running as a kubectl plugin
func (c *ConfigFlags) AddFlagsPluginMode(flags *pflag.FlagSet) {
	flags.StringVar(&c.SecretNamespace, flagSecretNamespace, c.SecretNamespace, "Kubernetes namespace where secret contains token")
	flags.StringVar(&c.SecretName, flagSecretName, c.SecretName, "Kubernetes secret name containing authentication token")
	c.addFlagsCommon(flags)
}

func (c *ConfigFlags) addFlagsCommon(flags *pflag.FlagSet) {
	flags.StringVar(&c.Token, flagToken, c.Token, "Portworx authentication token")
	flags.StringVar(&c.ConfigFile, flagConfigFile, c.ConfigFile, "Config file (default is $HOME/"+PxDefaultDir+"/"+PxDefaultConfigName+")")
	flags.StringVar(&c.ConfigDir, flagConfigDir, c.ConfigDir, "Config directory")
	flags.Int32Var(&c.Verbosity, flagVerbosity, c.Verbosity, "[0-3] Log level verbosity")
}
