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
package volume

import (
	"bytes"
	"math/big"
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var getVolumesCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	getVolumesCmd = &cobra.Command{
		Use:     "list [NAME]",
		Aliases: []string{"get"},
		Short:   "Get information about Portworx volumes",
		Example: `
  # Get informtation about the portworx volumes
  pxc volume list`,
		RunE: getVolumesExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	VolumeAddCommand(getVolumesCmd)
	getVolumesCmd.Flags().String("owner", "", "Owner of volume")
	getVolumesCmd.Flags().String("volumegroup", "", "Volume group id")
	//getVolumesCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")
	getVolumesCmd.Flags().Bool("show-k8s-info", false, "Show kubernetes information")
	getVolumesCmd.Flags().StringP("output", "o", "", "Output in yaml|json|wide")
	getVolumesCmd.Flags().Bool("show-labels", false, "Show labels in the last column of the output")
	getVolumesCmd.Flags().StringP("selector", "l", "", "Selector (label query) comma-separated name=value pairs")
	// TODO: Place here support for selectors and move the flags from the rootCmd
})

func GetAddCommand(cmd *cobra.Command) {
	getVolumesCmd.AddCommand(cmd)
}

func getVolumesExec(cmd *cobra.Command, args []string) error {
	// Check if any flag value is not provided, error out
	valErr := cliops.ValidateCliInput(cmd, args)
	if valErr != nil {
		return valErr
	}

	// Parse out all of the common cli volume flags
	cvi := cliops.NewCliInputs(cmd, args)

	// Create a cliOps object
	cliOps := cliops.NewCliOps(cvi)

	// Connect to pxc and k8s (if needed)
	err := cliOps.Connect()
	if err != nil {
		return err
	}
	defer cliOps.Close()
	// Create the parser object
	vgf := NewVolumeGetFormatter(cliOps)

	// Print the details and return errors if any
	return util.PrintFormatted(vgf)
}

type volumeGetFormatter struct {
	util.BaseFormatOutput
	cliOps  cliops.CliOps
	volumes portworx.Volumes
	nodes   portworx.Nodes
	pods    portworx.Pods
}

func NewVolumeGetFormatter(cliOps cliops.CliOps) *volumeGetFormatter {
	volSpec := &portworx.VolumeSpec{
		VolNames: cliOps.CliInputs().Args,
		Labels:   cliOps.CliInputs().Labels,
		Owner:    cliOps.CliInputs().Owner,
	}
	v := &volumeGetFormatter{
		cliOps:  cliOps,
		volumes: portworx.NewVolumes(cliOps.PxOps(), volSpec),
		pods:    portworx.NewPods(cliOps.COps(), &portworx.PodSpec{}),
	}
	v.FormatType = cliOps.CliInputs().FormatType
	return v
}

// YamlFormat returns the yaml representation of the object
func (p *volumeGetFormatter) YamlFormat() (string, error) {
	vols, err := p.volumes.GetVolumes()
	if err != nil {
		return "", err
	}
	return util.ToYaml(vols)
}

// JsonFormat returns the json representation of the object
func (p *volumeGetFormatter) JsonFormat() (string, error) {
	vols, err := p.volumes.GetVolumes()
	if err != nil {
		return "", err
	}
	return util.ToJson(vols)
}

// WideFormat returns the wide string representation of the object
func (p *volumeGetFormatter) WideFormat() (string, error) {
	return p.toTabbed()
}

// DefaultFormat returns the default string representation of the object
func (p *volumeGetFormatter) DefaultFormat() (string, error) {
	return p.toTabbed()
}

func (p *volumeGetFormatter) toTabbed() (string, error) {
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(writer)

	vols, err := p.volumes.GetVolumes()
	if err != nil {
		return "", err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return "", nil
	}

	p.nodes, err = portworx.NewNodesForVolumes(p.cliOps.PxOps(), vols)
	if err != nil {
		return "", err
	}
	// Start the columns
	t.AddHeader(p.getHeader()...)

	for _, n := range vols {
		l, err := p.getLine(n)
		if err != nil {
			return "", nil
		}
		t.AddLine(l...)
	}
	t.Print()

	return b.String(), nil
}

func (p *volumeGetFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.cliOps.CliInputs().Wide {
		header = []interface{}{"Id", "Name", "Size Gi", "HA", "Shared", "Encrypted", "Io Profile", "Status", "State", "Snap Enabled"}
	} else {
		header = []interface{}{"Name", "Size", "HA", "Shared", "Status", "State"}
	}
	if util.InKubectlPluginMode() {
		header = append(header, "Pods")
	}
	if p.cliOps.CliInputs().ShowLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *volumeGetFormatter) getLine(v *api.Volume) ([]interface{}, error) {
	spec := v.GetSpec()

	var line []interface{}

	state, err := p.nodes.GetAttachedState(v)
	if err != nil {
		return line, err
	}

	size := humanize.BigIBytes(big.NewInt(int64(spec.GetSize())))
	if p.cliOps.CliInputs().Wide {
		line = []interface{}{
			v.GetId(), v.GetLocator().GetName(), size, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), spec.GetEncrypted(),
			spec.GetCos(), portworx.PrettyStatus(v), state, spec.GetSnapshotSchedule() != "",
		}
	} else {
		line = []interface{}{
			v.GetLocator().GetName(), size, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), portworx.PrettyStatus(v), state,
		}
	}
	if util.InKubectlPluginMode() {
		pods, err := p.podsUsingVolume(v)
		if err != nil {
			return line, err
		}
		line = append(line, pods)
	}
	if p.cliOps.CliInputs().ShowLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line, nil
}

func (p *volumeGetFormatter) podsUsingVolume(v *api.Volume) (string, error) {
	usedPods, err := p.pods.PodsUsingVolume(v)
	if err != nil {
		return "", err
	}
	usedPodNames := make([]string, 0)
	namespace := v.Locator.VolumeLabels["namespace"]
	for _, pod := range usedPods {
		usedPodNames = append(usedPodNames, namespace+"/"+pod.Name)
	}
	return strings.Join(usedPodNames, ","), nil
}
