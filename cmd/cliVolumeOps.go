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
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type cliVolumeInputs struct {
	util.BaseFormatOutput
	wide        bool
	showLabels  bool
	volumeNames []string
	showK8s     bool
	// If namespace is nil, use default namespace
	// if namespace is "", use all-namespaces
	// else use specified namespace
	namespace *string
}

type cliVolumeOps struct {
	cliVolumeInputs
	pxVolumeOps portworx.PxVolumeOps
}

// Look for all of the common flags and create a new cliVolumeInputs object
func GetCliVolumeInputs(cmd *cobra.Command, args []string) *cliVolumeInputs {
	showK8s, _ := cmd.Flags().GetBool("show-k8s-info")
	output, _ := cmd.Flags().GetString("output")
	showLabels, _ := cmd.Flags().GetBool("show-labels")
	namespace := string("")
	return &cliVolumeInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		showK8s:     showK8s,
		showLabels:  showLabels,
		volumeNames: args,
		namespace:   &namespace, // In most places use all namespaces
	}
}

// Checks if namespace is specified and if so set it
func (p *cliVolumeInputs) GetNamespace(cmd *cobra.Command) {
	flagNamespace, _ := cmd.Flags().GetString("namespace")
	if len(flagNamespace) != 0 {
		p.namespace = &flagNamespace
		return
	}

	allNamespaces, _ := cmd.Flags().GetBool("all-namespaces")
	if allNamespaces {
		str := string("")
		p.namespace = &str
		return
	}

	// No default namespace was specified so use default namespace
	p.namespace = nil

}

// Create a new cliVolumeOps object
func NewCliVolumeOps(
	cvi *cliVolumeInputs,
) *cliVolumeOps {
	return &cliVolumeOps{
		cliVolumeInputs: *cvi,
	}
}

// Connect will make connections to px and k8s (if needed).
func (p *cliVolumeOps) Connect() error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}

	var (
		cc clientcmd.ClientConfig
		cs *kclikube.Clientset
	)

	if p.showK8s {
		cc, cs, err = KubeConnectDefault()
		if err != nil {
			return err
		}
		// Determine default namespace if nothing specified
		if p.namespace == nil {
			ns, _, err := cc.Namespace()
			if err != nil {
				return err
			}
			p.namespace = &ns
		}
	}

	volOpsInfo := &portworx.PxVolumeOpsInfo{
		PxConnectionData: portworx.PxConnectionData{
			Ctx:  ctx,
			Conn: conn,
		},
		KubeConnectionData: portworx.KubeConnectionData{
			ClientConfig: cc,
			ClientSet:    cs,
		},
		Namespace: *p.namespace,
		VolNames:  p.volumeNames,
	}

	p.pxVolumeOps, err = portworx.NewPxVolumeOps(volOpsInfo)
	return err
}

func (p *cliVolumeOps) Close() {
	p.pxVolumeOps.GetPxVolumeOpsInfo().Close()
}
