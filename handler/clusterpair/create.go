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
package clusterpair

import (
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/spf13/cobra"
)

type createClusterpairOpts struct {
	req             *api.ClusterPairCreateRequest
	source          string
	destination     string
	destinationPort uint32
	mode            string
}

var (
	ccpOpts              createClusterpairOpts
	createClusterpairCmd *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	ccpOpts = createClusterpairOpts{
		req: &api.ClusterPairCreateRequest{},
	}

	createClusterpairCmd = &cobra.Command{
		Use:     "clusterpair",
		Aliases: []string{"clusterpairs"},
		Short:   "Pair this cluster with another Portworx cluster",
		Example: `
  # Pair two clusters
  pxc create clusterpair -s pxcluster1 -d pxcluster2`,
		Long: `Creates a cluster pair between two Portworx clusters`,
		RunE: createClusterpairExec,
	}
})

var _ = commander.RegisterCommandVar(func() {
	//reateAddCommand(createClusterpairCmd)

	createClusterpairCmd.Flags().StringVar(&ccpOpts.source, "source", "", "Context for the source cluster (required)")
	createClusterpairCmd.Flags().StringVar(&ccpOpts.destination, "destination", "", "Context for the destination cluster (required)")
	createClusterpairCmd.Flags().Uint32Var(&ccpOpts.destinationPort, "destination-port", 9001,
		"Port for destination cluster (optional)")
	createClusterpairCmd.Flags().StringVar(&ccpOpts.mode, "mode", "", "Pairing mode to use (optional)")
	createClusterpairCmd.Flags().BoolVar(&ccpOpts.req.SetDefault, "set-default", false, "Set this as the default cluster pair (optional)")
	createClusterpairCmd.Flags().SortFlags = false
})

func CreateAddCommand(cmd *cobra.Command) {
	createClusterpairCmd.AddCommand(cmd)
}

func createClusterpairExec(c *cobra.Command, args []string) error {
	/* REMOVED UNTIL CONTEXT DONE
	contextManager, err := contextconfig.NewContextManager(cmd.GetConfigFile())
	if err != nil {
		return util.PxErrorMessagef(err, "Failed to load context configuration at path %s", cmd.GetConfigFile())
	}

	// Get connection info for destination cluster and remote cluster pair request
	destContext, err := contextManager.GetNamedContext(ccpOpts.destination)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get destination context")
	}
	destHost, _, err := net.SplitHostPort(destContext.Endpoint)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get cluster token")
	}
	ccpOpts.req.RemoteClusterIp = destHost
	ccpOpts.req.RemoteClusterPort = ccpOpts.destinationPort

	// Add mode to request
	ccpOpts.mode = strings.ToLower(ccpOpts.mode)
	switch {
	case ccpOpts.mode == "dr" || ccpOpts.mode == "disasterrecovery":
		ccpOpts.req.Mode = api.ClusterPairMode_DisasterRecovery
	default:
		ccpOpts.req.Mode = api.ClusterPairMode_Default
	}

	// Connect to source
	ctxSource, connSource, err := portworx.PxConnectNamed(cmd.GetConfigFile(), ccpOpts.source)
	if err != nil {
		return util.PxErrorMessagef(err, "Failed to connect to %s", ccpOpts.source)
	}
	defer connSource.Close()
	clusterpairsSource := api.NewOpenStorageClusterPairClient(connSource)

	// Connect to destination
	ctxDest, connDest, err := portworx.PxConnectNamed(cmd.GetConfigFile(), ccpOpts.destination)
	if err != nil {
		return util.PxErrorMessagef(err, "Failed to connect to %s", ccpOpts.destination)
	}
	defer connDest.Close()
	clusterpairsDest := api.NewOpenStorageClusterPairClient(connDest)

	// Get token from destination cluster
	tokenResp, err := clusterpairsDest.GetToken(ctxDest, &api.SdkClusterPairGetTokenRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get cluster token")
	}
	ccpOpts.req.RemoteClusterToken = tokenResp.Result.Token

	// Create pair from source to destination cluster
	_, err = clusterpairsSource.Create(ctxSource, &api.SdkClusterPairCreateRequest{
		Request: ccpOpts.req,
	})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to create cluster pair")
	}

	// Show user information
	util.Printf("Cluster pair created from %s to %s\n",
		ccpOpts.source,
		ccpOpts.destination,
	)
	*/
	return nil
}
