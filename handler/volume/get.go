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
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var getVolumesCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	getVolumesCmd = &cobra.Command{
		Use:     "volume",
		Aliases: []string{"volumes"},
		Short:   "Get information about Portworx volumes",
		RunE:    getVolumesExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.GetAddCommand(getVolumesCmd)
	getVolumesCmd.Flags().String("owner", "", "Owner of volume")
	getVolumesCmd.Flags().String("volumegroup", "", "Volume group id")
	getVolumesCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")
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
	cvi := cliops.GetCliVolumeInputs(cmd, args)

	// Create a cliVolumeOps object
	cvOps := cliops.NewCliVolumeOps(cvi)

	// Connect to px and k8s (if needed)
	err := cvOps.Connect()
	if err != nil {
		return err
	}
	defer cvOps.Close()
	// Create the parser object
	vgf := NewVolumeGetFormatter(cvOps)

	// Print the details and return errors if any
	return util.PrintFormatted(vgf)
}

type volumeGetFormatter struct {
	cliops.CliVolumeOps
}

func NewVolumeGetFormatter(cvOps *cliops.CliVolumeOps) *volumeGetFormatter {
	return &volumeGetFormatter{
		CliVolumeOps: *cvOps,
	}
}

// YamlFormat returns the yaml representation of the object
func (p *volumeGetFormatter) YamlFormat() (string, error) {
	vols, err := p.PxVolumeOps.GetVolumes()
	if err != nil {
		return "", err
	}
	return util.ToYaml(vols)
}

// JsonFormat returns the json representation of the object
func (p *volumeGetFormatter) JsonFormat() (string, error) {
	vols, err := p.PxVolumeOps.GetVolumes()
	if err != nil {
		return "", err
	}
	return util.ToJson(vols)
}

// WideFormat returns the wide string representation of the object
func (p *volumeGetFormatter) WideFormat() (string, error) {
	p.Wide = true
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

	vols, err := p.PxVolumeOps.GetVolumes()
	if err != nil {
		return "", err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return "", nil
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
	if p.Wide {
		header = []interface{}{"Id", "Name", "Size Gi", "HA", "Shared", "Encrypted", "Io Profile", "Status", "State", "Snap Enabled"}
	} else {
		header = []interface{}{"Name", "Size", "HA", "Shared", "Status", "State"}
	}
	if p.ShowK8s {
		header = append(header, "Pods")
	}
	if p.ShowLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *volumeGetFormatter) getLine(resp *api.SdkVolumeInspectResponse) ([]interface{}, error) {
	v := resp.GetVolume()
	spec := v.GetSpec()

	var line []interface{}

	state, err := p.PxVolumeOps.GetAttachedState(v)
	if err != nil {
		return line, err
	}

	size := humanize.BigIBytes(big.NewInt(int64(spec.GetSize())))
	if p.Wide {
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
	if p.ShowK8s {
		pods, err := p.podsUsingVolume(v)
		if err != nil {
			return line, err
		}
		line = append(line, pods)
	}
	if p.ShowLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line, nil
}

func (p *volumeGetFormatter) podsUsingVolume(v *api.Volume) (string, error) {
	usedPods, err := p.PxVolumeOps.PodsUsingVolume(v)
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
