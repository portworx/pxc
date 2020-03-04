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

package cluster

import (
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	pxc "github.com/portworx/pxc/pkg/component"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"

	"github.com/spf13/cobra"
)

type clusterShowFlags struct {
	output string
}

var (
	clusterShowCmd     *cobra.Command
	clusterShowOptions *clusterShowFlags
)

// Initialize variables and create the command
var _ = commander.RegisterCommandVar(func() {
	// clusterShowCmd represents the clusterShow command
	clusterShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Describe a Portworx cluster",
		Long:  "Show detailed information of Portworx cluster",
		Example: fmt.Sprintf("\n"+
			"# Display detailed information about Portworx cluster\n"+
			"  %s show", "cm"),
		RunE: clusterShowExec,
	}
	clusterShowOptions = newClusterShowFlags()
})

// Register this command
var _ = commander.RegisterCommandInit(func() {
	pxc.RootAddCommand(clusterShowCmd)

	clusterShowCmd.Flags().StringVarP(&clusterShowOptions.output, "output", "o", "", "Output in yaml|json")
})

// ClusterShowAddCommand allows sub commands to be added
func ClusterShowAddCommand(c *cobra.Command) {
	clusterShowCmd.AddCommand(c)
}

func newClusterShowFlags() *clusterShowFlags {
	return &clusterShowFlags{}
}

func clusterShowExec(c *cobra.Command, args []string) error {

	// This will automatically setup the the connection to portworx
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get Identity information
	identity := api.NewOpenStorageIdentityClient(conn)
	version, err := identity.Version(ctx, &api.SdkIdentityVersionRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get server version")
	}
	var versionDetails string
	for k, v := range version.GetVersion().GetDetails() {
		versionDetails += fmt.Sprintf(" %s: %s\n", k, v)
	}

	// Print cluster information
	cluster := api.NewOpenStorageClusterClient(conn)
	clusterInfo, err := cluster.InspectCurrent(ctx, &api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to inspect cluster")
	}

	formatter := &showFormatter{
		clusterInfo:    clusterInfo,
		version:        version,
		versionDetails: versionDetails,
	}
	formatter.SetFormat(clusterShowOptions.output)

	return util.PrintFormatted(formatter)
}

type showFormatter struct {
	util.BaseFormatOutput

	clusterInfo    *api.SdkClusterInspectCurrentResponse
	version        *api.SdkIdentityVersionResponse
	versionDetails string
}

func (p *showFormatter) YamlFormat() (string, error) {
	return util.ToYaml(p.clusterInfo)
}

func (p *showFormatter) JsonFormat() (string, error) {
	return util.ToJson(p.clusterInfo)
}

func (p *showFormatter) WideFormat() (string, error) {
	return p.DefaultFormat()
}

func (p *showFormatter) DefaultFormat() (string, error) {
	return fmt.Sprintf("Cluster ID: %s\n"+
		"Cluster UUID: %s\n"+
		"Cluster Status: %s\n"+
		"Version: %s\n"+
		"%s"+
		"SDK Version %s\n",
		p.clusterInfo.GetCluster().GetName(),
		p.clusterInfo.GetCluster().GetId(),
		util.SdkStatusToPrettyString(p.clusterInfo.GetCluster().GetStatus()),
		p.version.GetVersion().GetVersion(),
		p.versionDetails,
		p.version.GetSdkVersion().GetVersion()), nil
}
