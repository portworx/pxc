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
	"strings"

	"google.golang.org/grpc"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// getPvcCmd represents the getPvc command
var getPvcCmd = &cobra.Command{
	Use:     "pvc",
	Aliases: []string{"pvcs"},
	Short:   "Show Portworx volume information for Kuberntes PVCs",
	RunE:    getPvcExec,
}

func init() {
	getCmd.AddCommand(getPvcCmd)
	getPvcCmd.Flags().StringP("namespace", "n", "", "Kubernetes namespace")
	getPvcCmd.Flags().Bool("all-namespaces", false, "Kubernetes namespace")
	getPvcCmd.Flags().StringP("output", "o", "", "Output in yaml|json|wide")
}

func getPvcExec(cmd *cobra.Command, args []string) error {
	// Connect to Portworx
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Connect to kubernetes
	cc, kc, err := KubeConnectDefault()
	if err != nil {
		return err
	}

	// Determine namespace
	namespace, _, err := cc.Namespace()
	if err != nil {
		return err
	}
	flagNamespace, _ := cmd.Flags().GetString("namespace")
	allNamespaces, _ := cmd.Flags().GetBool("all-namespaces")
	if len(flagNamespace) != 0 {
		namespace = flagNamespace
	}
	if allNamespaces {
		namespace = ""
	}

	// Get all the PVCs according to the request
	pvcClient := kc.CoreV1().PersistentVolumeClaims(namespace)
	pvcList, err := pvcClient.List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	pvcs := pvcList.Items
	if len(pvcs) == 0 {
		util.Printf("No resources found\n")
		return nil
	}

	// Get all pods in the namespace
	podClient := kc.CoreV1().Pods(namespace)
	podList, err := podClient.List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	pods := podList.Items

	// Get all volumes
	volumes := api.NewOpenStorageVolumeClient(conn)
	volsInfo, err := volumes.InspectWithFilters(ctx, &api.SdkVolumeInspectWithFiltersRequest{})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to get volumes")
	}
	vols := volsInfo.GetVolumes()

	// Collect data
	pxpvcs := make([]*kubernetes.PxPvc, len(pvcs))
	for i, pvc := range pvcs {
		pxpvcs[i] = kubernetes.NewPxPvc(&pvc)
		pxpvcs[i].SetVolume(vols)
		pxpvcs[i].SetPods(pods)
	}

	// Get output
	output, _ := cmd.Flags().GetString("output")
	// Determine if it is a wide output
	wide := output == "wide"
	// Determine if we need to show labels
	showLabels, _ := cmd.Flags().GetBool("show-labels")

	// Print
	t := util.NewTabby()
	f := &getPvcColumnFormatter{
		wide:       wide,
		showLabels: showLabels,
		ctx:        ctx,
		conn:       conn,
	}
	t.AddHeader(f.getHeader()...)
	for _, pxpvc := range pxpvcs {
		t.AddLine(f.getLine(pxpvc)...)
	}
	t.Print()

	return nil
}

type getPvcColumnFormatter struct {
	wide       bool
	showLabels bool
	ctx        context.Context
	conn       *grpc.ClientConn
}

func (p *getPvcColumnFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.wide {
		header = []interface{}{"NAME", "VOLUME", "VOLUME ID", "HA", "CAPACITY", "SHARED", "STATUS", "STATE", "SNAP ENABLED", "ENCRYPTED", "PODS"}
	} else {
		header = []interface{}{"NAME", "VOLUME", "CAPACITY", "SHARED", "STATE", "PODS"}
	}
	if p.showLabels {
		header = append(header, "LABELS")
	}

	return header
}

func (p *getPvcColumnFormatter) getLine(pxpvc *kubernetes.PxPvc) []interface{} {

	v := pxpvc.GetVolume()

	if v == nil {
		return []interface{}{pxpvc.Name}
		/*
			if p.wide {
				return []interface{}{
					pxpvc.Name,
					"",
					"",
					"",
					"",
					"",
					"",
					"",
					"",
					"",
				}
			} else {
			}
		*/
	}
	spec := v.GetSpec()

	// Get node information if it is attached
	// TODO: Make this a pkg/volume call
	// This shares code with cmd/getVolume
	//  -- start--
	var node *api.StorageNode
	if len(v.GetAttachedOn()) != 0 {
		nodes := api.NewOpenStorageNodeClient(p.conn)
		nodeInfo, err := nodes.Inspect(
			p.ctx,
			&api.SdkNodeInspectRequest{NodeId: v.GetAttachedOn()})
		if err != nil {
			util.Eprintf("%v\n",
				util.PxErrorMessagef(err,
					"Failed to get node information where volume %s is attached",
					v.GetLocator().GetName()))
			return nil
		}
		node = nodeInfo.GetNode()
	}

	// Determine the status of the volume
	state := "Detached"
	if v.State == api.VolumeState_VOLUME_STATE_ATTACHED {
		if node != nil {
			state = "on " + node.GetHostname()
		} else {
			state = "Attached"
		}
	} else if v.State == api.VolumeState_VOLUME_STATE_DETATCHING {
		if node != nil {
			state = "Was on " + node.GetHostname()
		} else {
			state = "Detaching"
		}
	}
	// -- end --
	/*
	   $ px get pvc
	   NAME        VOLUME                                    CAPACITY  SHARED  STATE  PODS
	   ----        ------                                    --------  ------  -----  ----
	   mysql-data  pvc-d2a47415-1aef-428c-b998-5aee138d93a9  2         1       false  on lpabon-k8s-1-node2  default/mysql-59b76b98f9-grcvd

	   lpabon@PDC4-SM26-N8 : ~/git/golang/porx/src/github.com/portworx/px
	   $ px get pvc -o wide
	   NAME        VOLUME                                    VOLUME ID           HA  CAPACITY  SHARED  STATUS  STATE  SNAP ENABLED           ENCRYPTED  PODS
	   ----        ------                                    ---------           --  --------  ------  ------  -----  ------------           ---------  ----
	   mysql-data  pvc-d2a47415-1aef-428c-b998-5aee138d93a9  605625582897896102  1   2         1       false   UP     on lpabon-k8s-1-node2  false      false  default/mysql-59b76b98f9-grcvd
	*/

	// Size needs to be done better
	var line []interface{}
	if p.wide {
		line = []interface{}{
			pxpvc.Name,
			v.GetLocator().GetName(),
			v.GetId(),

			spec.GetHaLevel(),
			fmt.Sprintf("%d Gi", spec.GetSize()/Gi),
			spec.GetShared() || spec.GetSharedv4(),

			portworx.PrettyStatus(v),
			state,
			spec.GetSnapshotSchedule() != "",

			spec.GetEncrypted(),
			strings.Join(pxpvc.PodNames, ","),
		}
	} else {
		line = []interface{}{
			pxpvc.Name,
			v.GetLocator().GetName(),
			fmt.Sprintf("%d Gi", spec.GetSize()/Gi),

			spec.GetShared() || spec.GetSharedv4(),
			state,
			strings.Join(pxpvc.PodNames, ","),
		}
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line
}
