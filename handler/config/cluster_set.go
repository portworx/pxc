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
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// clusterSetCmd represents the config command
var (
	clusterSetCmd *cobra.Command
	clusterSet    *config.Cluster
	opt           struct {
		cafile string
		force  bool
	}
)

var _ = commander.RegisterCommandVar(func() {
	clusterSet = config.NewCluster()
	clusterSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Configure pxc to communicate with your cluster",
		Example: `
  # Setup simple endpoint
  pxc config cluster set --name=mycluster --endpoint=127.0.0.1:9020

  # If in kubectl plugin mode
  pxc config cluster set --portworx-service-port=8900`,
		RunE: clusterSetExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	ClusterAddCommand(clusterSetCmd)

	clusterSetCmd.Flags().StringVar(&clusterSet.Name,
		"name", "", "Name for Portworx cluster (ignored when used as a kubectl plugin)")
	clusterSetCmd.Flags().BoolVar(&opt.force,
		"force", false, "Force the configuration setting.")
	clusterSetCmd.Flags().BoolVar(&clusterSet.Secure,
		"tls", false, "Enable if using TLS. Passing a CA will enable this automatically.")
	clusterSetCmd.Flags().StringVar(&opt.cafile,
		"cafile", "", "Path to CA certificate")
	clusterSetCmd.Flags().StringVar(&clusterSet.Endpoint,
		"endpoint", "", "Direct connection to a Portworx node gRPC endpoint. "+
			"This endpoint would be used instead of the Kubernetes Portworx API service. "+
			"Example: pxnode123:9020")

	if util.InKubectlPluginMode() {
		clusterSetCmd.Flags().StringVar(&clusterSet.TunnelServiceNamespace,
			"portworx-service-namespace", "kube-system", "Kubernetes namespace for the Portworx service")
		clusterSetCmd.Flags().StringVar(&clusterSet.TunnelServiceName,
			"portworx-service-name", "portworx-api", "Kubernetes name for the Portworx service")
		clusterSetCmd.Flags().StringVar(&clusterSet.TunnelServicePort,
			"portworx-service-port", "9020", "Port for the Portworx SDK endpoint in the Kubernetes service")
	}
})

func clusterSetExec(cmd *cobra.Command, args []string) error {
	var cabytes []byte
	var err error

	if opt.cafile != "" {
		cabytes, err = ioutil.ReadFile(opt.cafile)
		if err != nil {
			return fmt.Errorf("failed to read %s: %v", opt.cafile, err)
		}
		clusterSet.CACertData = base64.StdEncoding.EncodeToString(cabytes)
	}

	if opt.force {
		logrus.Debug("Skipping endpoint validation (--force option used)")
	} else if err = util.VerifyConnection(clusterSet.Endpoint, &clusterSet.Secure, cabytes); err != nil {
		return err
	}

	if err = config.CM().ConfigSaveCluster(clusterSet); err != nil {
		return err
	}

	util.Printf("Cluster information set\n")
	return nil
}
