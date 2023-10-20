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
)

type CliInputs struct {
	util.BaseFormatOutput
	Wide          bool
	ShowLabels    bool
	AllNamespaces bool
	Owner         string
	Labels        map[string]string
	Args          []string
}

type CliOps interface {
	// Connect creates connections to portworx and if needed to k8s
	Connect() error
	// Close connections to portworx and k8s
	Close()
	// CliInputs returns the CliInputs
	CliInputs() *CliInputs
	// PxOps returns the portwor connection object
	PxOps() portworx.PxOps
	// COps returns the k8s connection object
	COps() kubernetes.COps
}

type cliOps struct {
	cliInputs *CliInputs
	pxops     portworx.PxOps
	cops      kubernetes.COps
}

var (
	inst CliOps
)

func GetCliOps() CliOps {
	return inst
}

// NewCliVolumeInputs looks for all of the common flags and create a new cliVolumeInputs object
func NewCliInputs(cmd *cobra.Command, args []string) *CliInputs {
	output, _ := cmd.Flags().GetString("output")
	wide := false
	if output == "wide" {
		wide = true
	}
	showLabels, _ := cmd.Flags().GetBool("show-labels")
	labels, _ := cmd.Flags().GetString("selector")
	owner, _ := cmd.Flags().GetString("owner")

	//convert string to map
	mlabels, _ := util.CommaStringToStringMap(labels)

	allNamespaces, _ := cmd.Flags().GetBool("all-namespaces")
	// If valid label is present, we need to pass it.
	return &CliInputs{
		BaseFormatOutput: util.BaseFormatOutput{
			FormatType: output,
		},
		Wide:          wide,
		Owner:         owner,
		ShowLabels:    showLabels,
		AllNamespaces: allNamespaces,
		Args:          args,
		Labels:        mlabels,
	}
}

func NewCliOps(ci *CliInputs) CliOps {
	inst = &cliOps{
		cliInputs: ci,
	}
	return inst
}

func (co *cliOps) CliInputs() *CliInputs {
	return co.cliInputs
}

func (co *cliOps) PxOps() portworx.PxOps {
	return co.pxops
}

func (co *cliOps) COps() kubernetes.COps {
	return co.cops
}

// Connect will make connections to pxc and k8s (if needed).
func (p *cliOps) Connect() error {
	// Already connected, just return
	if p.pxops != nil {
		return nil
	}

	pxops, err := portworx.NewPxOps()
	if err != nil {
		return err
	}
	p.pxops = pxops

	cops, err := kubernetes.NewCOps()
	if err != nil {
		return err
	}
	p.cops = cops

	return nil
}

func (p *cliOps) Close() {
	if p.pxops != nil {
		p.pxops.Close()
		p.pxops = nil
	}
	if p.cops != nil {
		p.cops.Close()
		p.cops = nil
	}
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
