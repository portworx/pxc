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
	"io/ioutil"

	"github.com/portworx/px/pkg/contextconfig"
	"github.com/portworx/px/pkg/util"

	"github.com/spf13/cobra"
)

// contextCreateCmd represents the contextCreate command
var contextCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a context",
	Long: `A context is the information needed to connect to
Portworx and any other system. This information will be saved
to a file called config.yml in a directory called .px under
your home directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return contextCreateExec(cmd, args)
	},
}

func init() {
	contextCmd.AddCommand(contextCreateCmd)

	contextCreateCmd.Flags().String("name", "", "Name of context")
	contextCreateCmd.Flags().String("token", "", "Token for use in this context")
	contextCreateCmd.Flags().String("endpoint", "", "Portworx service endpoint. Ex. 127.0.0.1:9020")
	contextCreateCmd.Flags().Bool("secure", false, "Use secure connection")
	contextCreateCmd.Flags().String("cafile", "", "Path to client CA certificate if needed")
	contextCreateCmd.Flags().String("kubeconfig", "", "Path to Kubeconfig file if any")
}

func contextCreateExec(cmd *cobra.Command, args []string) error {

	c := new(contextconfig.ClientContext)

	// Required
	if s, _ := cmd.Flags().GetString("name"); len(s) != 0 {
		c.Context = s
	} else {
		return fmt.Errorf("Must supply a name for the context")
	}
	if s, _ := cmd.Flags().GetString("endpoint"); len(s) != 0 {
		c.Endpoint = s
	} else {
		return fmt.Errorf("Must supply an endpoint for the context")
	}

	// Optional
	if s, _ := cmd.Flags().GetString("kubeconfig"); len(s) != 0 {
		// TODO: Should check if the file exists
		c.Kubeconfig = s
	}
	if s, _ := cmd.Flags().GetString("token"); len(s) != 0 {
		// TODO: Check the token is a JWT token
		c.Token = s
	}
	if s, _ := cmd.Flags().GetString("secure"); len(s) != 0 {
		c.Secure = true
	}
	if s, _ := cmd.Flags().GetString("cafile"); len(s) != 0 {
		// TODO: Should check if the file exists
		data, err := ioutil.ReadFile(s)
		if err != nil {
			return fmt.Errorf("Unable to read CA file %s", s)
		}
		c.TlsData.Cacert = string(data)
	}

	config := contextconfig.NewContextConfig(cfgFile)
	if err := config.Add(c); err != nil {
		return err
	}

	util.Printf("New context saved to %s\n", cfgFile)

	return nil
}
