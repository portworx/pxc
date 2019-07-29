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

type volumeFormatter struct {
	util.BaseFormatOutput
	wide        bool
	showLabels  bool
	showK8s     bool
	pxVolumeOps portworx.PxVolumeOps
}

func (p *volumeFormatter) close() {
	p.pxVolumeOps.GetPxVolumeOpsInfo().Close()
}

func newVolumeFormatter(cmd *cobra.Command, args []string) (*volumeFormatter, error) {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return nil, err
	}

	var (
		cc clientcmd.ClientConfig
		cs *kclikube.Clientset
	)

	showK8s, _ := cmd.Flags().GetBool("show-k8s-info")
	if showK8s {
		cc, cs, err = KubeConnectDefault()
		if err != nil {
			return nil, err
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
		VolNames: args,
	}

	output, _ := cmd.Flags().GetString("output")
	showLabels, _ := cmd.Flags().GetBool("show-labels")
	pxVolumeOps, _ := portworx.NewPxVolumeOps(volOpsInfo)

	return &volumeFormatter{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		pxVolumeOps: pxVolumeOps,
		showLabels:  showLabels,
		showK8s:     showK8s,
	}, nil
}
