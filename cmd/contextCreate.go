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

var contextCreateCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	// contextCreateCmd represents the contextCreate command
	contextCreateCmd = &cobra.Command{
		Use:     "create [NAME]",
		Short:   "Create a context",
		Example: "$ px context create mycluster --endpoint=123.456.1.10:9020",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for context")
			}
			return nil
		},
		Long: `A context is the information needed to connect to
Portworx and any other system. This information will be saved
to a file called config.yml in a directory called .px under
your home directory.`,
		RunE: contextCreateExec,
	}
})

var _ = RegisterCommandInit(func() {
	contextCmd.AddCommand(contextCreateCmd)

	contextCreateCmd.Flags().String("token", "", "Token for use in this context")
	contextCreateCmd.Flags().String("endpoint", "", "Portworx service endpoint. Ex. 127.0.0.1:9020")
	contextCreateCmd.Flags().Bool("secure", false, "Use secure connection")
	contextCreateCmd.Flags().String("cafile", "", "Path to client CA certificate if needed")
	contextCreateCmd.Flags().String("kubeconfig", "", "Path to Kubeconfig file if any")
})

func contextCreateExec(cmd *cobra.Command, args []string) error {

	c := new(contextconfig.ClientContext)
	c.Name = args[0]

	// Check if this an update
	update := false
	contextManager, err := contextconfig.NewContextManager(cfgFile)
	if err == nil {
		// Check if we need to update a context
		existingContext, err := contextManager.GetNamedContext(c.Name)
		if err == nil {
			update = true
			c = existingContext
		}
	} else {
		contextManager = contextconfig.New(cfgFile)
	}

	// Update endpoint
	if s, _ := cmd.Flags().GetString("endpoint"); len(s) != 0 {
		// TODO: If no port is provided, assume 9020
		c.Endpoint = s
	} else {
		return fmt.Errorf("Must supply an endpoint for the context")
	}

	// Optional
	if s, _ := cmd.Flags().GetString("kubeconfig"); len(s) != 0 {
		if !util.IsFileExists(s) {
			return fmt.Errorf("kubeconfig file: %s does not exists", s)
		}
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
		if !util.IsFileExists(s) {
			return fmt.Errorf("CA file: %s does not exists", s)
		}
		data, err := ioutil.ReadFile(s)
		if err != nil {
			return fmt.Errorf("Unable to read CA file %s", s)
		}
		c.TlsData.Cacert = string(data)
	}

	if err = contextManager.Add(c); err != nil {
		return err
	}

	if update {
		util.Printf("Updated context %s in %s\n", c.Name, cfgFile)
	} else {
		util.Printf("New context %s saved to %s\n", c.Name, cfgFile)
	}

	return nil
}
