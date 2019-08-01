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
	"bytes"
	"math/big"
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	humanize "github.com/dustin/go-humanize"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	v1 "k8s.io/api/core/v1"

	"github.com/spf13/cobra"
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
	// Parse out all of the common cli volume flags
	cvi := GetCliVolumeInputs(cmd, make([]string, 0))
	cvi.showK8s = true
	cvi.GetNamespace(cmd)

	// Create a cliVolumeOps object
	cvOps := NewCliVolumeOps(cvi)

	// Connect to px and k8s (if needed)
	err := cvOps.Connect()
	if err != nil {
		return err
	}
	defer cvOps.Close()

	// Create the parser object
	pgf := NewPvcGetFormatter(cvOps)

	// Print the details and return errors if any
	return util.PrintFormatted(pgf)
}

type pvcGetFormatter struct {
	cliVolumeOps
}

func NewPvcGetFormatter(cvOps *cliVolumeOps) *pvcGetFormatter {
	return &pvcGetFormatter{
		cliVolumeOps: *cvOps,
	}
}

func (p *pvcGetFormatter) getPvcs() ([]*v1.PersistentVolumeClaim, error) {
	pxpvcs, err := p.pxVolumeOps.GetPxPvcs()
	if err != nil {
		return make([]*v1.PersistentVolumeClaim, 0), err
	}
	pvcs := make([]*v1.PersistentVolumeClaim, len(pxpvcs))
	for i, _ := range pvcs {
		pvcs[i] = pxpvcs[i].Pvc
	}
	return pvcs, nil
}

// YamlFormat returns the yaml representation of the pvc
func (p *pvcGetFormatter) YamlFormat() (string, error) {
	pvcs, err := p.getPvcs()
	if err != nil {
		return "", err
	}
	return util.ToYaml(pvcs)
}

// JsonFormat returns the json representation of the pvc
func (p *pvcGetFormatter) JsonFormat() (string, error) {
	pvcs, err := p.getPvcs()
	if err != nil {
		return "", err
	}
	return util.ToJson(pvcs)
}

// WideFormat returns the wide string representation of the object
func (p *pvcGetFormatter) WideFormat() (string, error) {
	p.wide = true
	return p.toTabbed()
}

// DefaultFormat returns the default string representation of the object
func (p *pvcGetFormatter) DefaultFormat() (string, error) {
	return p.toTabbed()
}

func (p *pvcGetFormatter) toTabbed() (string, error) {
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(writer)

	pvcs, err := p.pxVolumeOps.GetPxPvcs()
	if err != nil {
		return "", err
	}

	if len(pvcs) == 0 {
		util.Printf("No resources found\n")
		return "", nil
	}

	// Start the columns
	t.AddHeader(p.getHeader()...)
	for _, n := range pvcs {
		l, err := p.getLine(n)
		if err != nil {
			return "", err
		}
		t.AddLine(l...)
	}
	t.Print()

	return b.String(), err
}

func (p *pvcGetFormatter) getHeader() []interface{} {
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

func (p *pvcGetFormatter) getLine(pxpvc *kubernetes.PxPvc) ([]interface{}, error) {
	v := pxpvc.GetVolume()
	if v == nil {
		return []interface{}{pxpvc.Name}, nil
	}

	var line []interface{}

	spec := v.GetSpec()
	state, err := p.pxVolumeOps.GetAttachedState(v)
	if err != nil {
		return line, err
	}
	size := humanize.BigIBytes(big.NewInt(int64(spec.GetSize())))
	pods := strings.Join(pxpvc.PodNames, ",")

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

	if p.wide {
		line = []interface{}{
			pxpvc.Name,
			v.GetLocator().GetName(),
			v.GetId(),
			spec.GetHaLevel(),
			size,
			portworx.SharedString(v),
			portworx.PrettyStatus(v),
			state,
			spec.GetSnapshotSchedule() != "",
			spec.GetEncrypted(),
			pods,
		}
	} else {
		line = []interface{}{
			pxpvc.Name,
			v.GetLocator().GetName(),
			size,
			portworx.SharedString(v),
			state,
			pods,
		}
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line, nil
}
