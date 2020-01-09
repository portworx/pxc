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
	"github.com/spf13/pflag"
)

const (
	SpecifiedContext = "cfgcontext"
	File             = "cfgfile"
	PluginEndpoint
	PxDefaultDir        = ".pxc"
	PxDefaultConfigName = "config.yml"

	flagPrefix          = "pxc."
	flagConfigFile      = flagPrefix + "config"
	flagContext         = flagPrefix + "context"
	flagSecretNamespace = flagPrefix + "secret-namespace"
	flagSecretName      = flagPrefix + "secret-name"
	flagToken           = flagPrefix + "token"
	flagVerbosity       = flagPrefix + "v"
)

type ConfigFlags struct {
	ConfigFile      string
	Context         string
	SecretNamespace string
	SecretName      string
	Token           string
	Verbosity       int32
}

var (
	config  = map[string]string{}
	CliOpts *ConfigFlags
)

func Get(k string) string {
	return config[k]
}

func Set(k, v string) {
	config[k] = v
}

func NewConfigFlags() *ConfigFlags {
	return &ConfigFlags{}
}

// AddFlags adds the appropriate global flags
func (c *ConfigFlags) AddFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.ConfigFile, flagConfigFile, c.ConfigFile, "Config file (default is $HOME/"+PxDefaultDir+"/"+PxDefaultConfigName+")")
	flags.StringVar(&c.Context, flagContext, c.Context, "Force context name for the command")
	c.addFlagsCommon(flags)
}

// AddFlagsPluginMode adds the appropriate global flags when running as a kubectl plugin
func (c *ConfigFlags) AddFlagsPluginMode(flags *pflag.FlagSet) {
	flags.StringVar(&c.SecretNamespace, flagSecretNamespace, c.SecretNamespace, "Kubernetes namespace where secret contains token")
	flags.StringVar(&c.SecretName, flagSecretName, c.SecretName, "Kubernetes secret name containing authentication token")
	flags.StringVar(&c.Token, flagToken, c.Token, "Portworx authentication token")
	c.addFlagsCommon(flags)
}

func (c *ConfigFlags) addFlagsCommon(flags *pflag.FlagSet) {
	flags.Int32Var(&c.Verbosity, flagVerbosity, c.Verbosity, "[0-4] Log level verbosity")
}
