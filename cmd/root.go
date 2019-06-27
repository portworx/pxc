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
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/portworx/px/pkg/contextconfig"
	pxgrpc "github.com/portworx/px/pkg/grpc"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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
	optEndpoint string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "px",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+pxDefaultDir+"/"+pxDefaultConfigName+".yaml)")
	rootCmd.PersistentFlags().StringVar(&optEndpoint, "endpoint", "", "Portworx service endpoint")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output in yaml|json|wide")
	rootCmd.PersistentFlags().Bool("show-labels", false, "Show labels in the last column of the output")
	rootCmd.PersistentFlags().StringP("selector", "l", "", "Comma separated label selector of the form 'key=value,key=value'")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".px" (without extension).
		viper.AddConfigPath(path.Join(home, pxDefaultDir))
		viper.SetConfigName(pxDefaultConfigName)
		cfgFile = path.Join(home, pxDefaultDir, pxDefaultConfigName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// pxConnect will connect to the server using TLS if needed and returns
// the context setup with any security if any and the grpc client. The
// context will not have a timeout set, that should be setup by each caller.
func pxConnect() (context.Context, *grpc.ClientConn) {
	pxctx, err := contextconfig.NewContextConfig(cfgFile).Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	conn, err := pxgrpc.Connect(pxctx.Endpoint, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return context.Background(), conn
}

func pxPrintGrpcErrorWithMessagef(err error, format string, args ...string) {
	pxPrintGrpcErrorWithMessage(err, fmt.Sprintf(format, args))
}

func pxPrintGrpcErrorWithMessage(err error, msg string) {
	gerr, _ := status.FromError(err)
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, gerr.Message())
}

func pxPrintGrpcError(err error) {
	gerr, _ := status.FromError(err)
	fmt.Fprintf(os.Stderr, "%s\n", gerr.Message())
}

func listContains(list []string, s string) bool {
	for _, value := range list {
		if value == s {
			return true
		}
	}
	return false
}

func listHaveMatch(list, match []string) bool {
	for _, s := range match {
		if listContains(list, s) {
			return true
		}
	}
	return false
}

func labelsToString(labels map[string]string) string {
	s := make([]string, 0, len(labels))
	for k, v := range labels {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, ",")
}
