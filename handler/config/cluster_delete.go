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
package configcli

import (
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type clusterDeleteFlagsTypes struct {
	Name string
}

// clusterDeleteCmd represents the config command
var (
	clusterDeleteCmd   *cobra.Command
	clusterDeleteFlags *clusterDeleteFlagsTypes
)

var _ = commander.RegisterCommandVar(func() {
	clusterDeleteFlags = &clusterDeleteFlagsTypes{}
	clusterDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete pxc cluster configuration",
		RunE:  clusterDeleteExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ClusterAddCommand(clusterDeleteCmd)

	clusterDeleteCmd.Flags().StringVar(&clusterDeleteFlags.Name,
		"name", "", "Name for Portworx cluster (ignored when used as a kubectl plugin)")
})

func clusterDeleteExec(cmd *cobra.Command, args []string) error {
	if err := config.CM().ConfigDeleteCluster(clusterDeleteFlags.Name); err != nil {
		return err
	}
	util.Printf("Cluster information removed\n")
	return nil
}
