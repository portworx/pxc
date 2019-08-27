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
package cliops

import (
	"fmt"

	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type CliVolumeInputs struct {
	util.BaseFormatOutput
	Wide        bool
	ShowLabels  bool
	VolumeNames []string
	ShowK8s     bool
	// If namespace is nil, use default namespace
	// if namespace is "", use all-namespaces
	// else use specified namespace
	Namespace *string
	Labels    map[string]string
}

type CliVolumeOps struct {
	CliVolumeInputs
	PxVolumeOps portworx.PxVolumeOps
}

// GetCliVolumeInputs looks for all of the common flags and create a new cliVolumeInputs object
func GetCliVolumeInputs(cmd *cobra.Command, args []string) *CliVolumeInputs {
	showK8s, _ := cmd.Flags().GetBool("show-k8s-info")
	output, _ := cmd.Flags().GetString("output")
	showLabels, _ := cmd.Flags().GetBool("show-labels")
	namespace := string("")
	labels, _ := cmd.Flags().GetString("selector")
	//convert string to map
	mlabels, _ := util.CommaStringToStringMap(labels)
	// If valid label is present, we need to pass it.
	return &CliVolumeInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		ShowK8s:     showK8s,
		ShowLabels:  showLabels,
		VolumeNames: args,
		Namespace:   &namespace, // In most places use all namespaces
		Labels:      mlabels,
	}
}

// Checks if namespace is specified and if so set it
func (p *CliVolumeInputs) GetNamespace(cmd *cobra.Command) {
	flagNamespace, _ := cmd.Flags().GetString("namespace")
	if len(flagNamespace) != 0 {
		p.Namespace = &flagNamespace
		return
	}

	allNamespaces, _ := cmd.Flags().GetBool("all-namespaces")
	if allNamespaces {
		str := string("")
		p.Namespace = &str
		return
	}

	// No default namespace was specified so use default namespace
	p.Namespace = nil
}

// Create a new cliVolumeOps object
func NewCliVolumeOps(
	cvi *CliVolumeInputs,
) *CliVolumeOps {
	return &CliVolumeOps{
		CliVolumeInputs: *cvi,
	}
}

// Connect will make connections to px and k8s (if needed).
func (p *CliVolumeOps) Connect() error {
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}

	var (
		cc clientcmd.ClientConfig
		cs *kclikube.Clientset
	)

	if p.ShowK8s {
		cc, cs, err = kubernetes.KubeConnectDefault()
		if err != nil {
			return err
		}
		// Determine default namespace if nothing specified
		if p.Namespace == nil {
			ns, _, err := cc.Namespace()
			if err != nil {
				return err
			}
			p.Namespace = &ns
		}
	}

	volOpsInfo := &portworx.PxVolumeOpsInfo{
		PxConnectionData: portworx.PxConnectionData{
			Ctx:  ctx,
			Conn: conn,
		},
		KubeConnectionData: kubernetes.KubeConnectionData{
			ClientConfig: cc,
			ClientSet:    cs,
		},
		Namespace: *p.Namespace,
		VolNames:  p.VolumeNames,
		Labels:    p.Labels,
	}

	p.PxVolumeOps, err = portworx.NewPxVolumeOps(volOpsInfo)
	return err
}

func (p *CliVolumeOps) Close() {
	p.PxVolumeOps.GetPxVolumeOpsInfo().Close()
}

// Validating the user provided options
func ValidateCliInput(cmd *cobra.Command, args []string) error {
	selector, _ := cmd.Flags().GetString("selector")
	// A case in which user mentions args like <volname> along with label flag but
	// not mentioning labels (k,v) pair. Need to error out in this case.
	if len(args) > 0 && len(selector) > 0 {
		return fmt.Errorf("name cannot be provided when a selector is specified")
	}

	if len(selector) > 0 {
		_, err := util.CommaStringToStringMap(selector)
		if err != nil {
			return err
		}
	}

	return nil
}
