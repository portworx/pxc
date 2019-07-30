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
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

// getVolumesCmd represents the getVolumes command
var getVolumesCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"volumes"},
	Short:   "Get information about Portworx volumes",
	RunE:    getVolumesExec,
}

func init() {
	getCmd.AddCommand(getVolumesCmd)
	getVolumesCmd.Flags().String("owner", "", "Owner of volume")
	getVolumesCmd.Flags().String("volumegroup", "", "Volume group id")
	getVolumesCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")
	getVolumesCmd.Flags().Bool("show-k8s-info", false, "Show kubernetes information")

	// TODO: Place here support for selectors and move the flags from the rootCmd
}

func getVolumesExec(cmd *cobra.Command, args []string) error {
	// Parse out all of the common cli volume flags
	cvi := GetCliVolumeInputs(cmd, args)

	// Create a cliVolumeOps object
	cvOps := NewCliVolumeOps(cvi)

	// Connect to px and k8s (if needed)
	err := cvOps.Connect()
	if err != nil {
		return err
	}
	defer cvOps.Close()

	// Create the parser object
	vgf := NewVolumeGetFormatter(cvOps)

	// Print the details
	vgf.Print()

	// Return any errors found during parsing
	return vgf.GetError()
}

type volumeGetFormatter struct {
	cliVolumeOps
}

func NewVolumeGetFormatter(cvOps *cliVolumeOps) *volumeGetFormatter {
	return &volumeGetFormatter{
		cliVolumeOps: *cvOps,
	}
}

// String returns the formatted output of the object as per the format set.
func (p *volumeGetFormatter) String() string {
	return util.GetFormattedOutput(p)
}

// Print writes the object to stdout
func (p *volumeGetFormatter) Print() {
	util.Printf("%v", p)
}

// YamlFormat returns the yaml representation of the object
func (p *volumeGetFormatter) YamlFormat() string {
	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		p.SetError(err)
		return ""
	}
	return util.ToYaml(vols)
}

// JsonFormat returns the json representation of the object
func (p *volumeGetFormatter) JsonFormat() string {
	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		p.SetError(err)
		return ""
	}
	return util.ToJson(vols)
}

// WideFormat returns the wide string representation of the object
func (p *volumeGetFormatter) WideFormat() string {
	p.wide = true
	return p.toTabbed()
}

// DefaultFormat returns the default string representation of the object
func (p *volumeGetFormatter) DefaultFormat() string {
	return p.toTabbed()
}

func (p *volumeGetFormatter) toTabbed() string {
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(writer)

	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		p.SetError(err)
		return ""
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return ""
	}

	// Start the columns
	t.AddHeader(p.getHeader()...)

	for _, n := range vols {
		l, err := p.getLine(n)
		if err != nil {
			p.SetError(err)
			return ""
		}
		t.AddLine(l...)
	}
	t.Print()

	return b.String()
}

func (p *volumeGetFormatter) getHeader() []interface{} {
	var header []interface{}
	if p.wide {
		header = []interface{}{"Id", "Name", "Size Gi", "HA", "Shared", "Encrypted", "Io Profile", "Status", "State", "Snap Enabled"}
	} else {
		header = []interface{}{"Name", "Size", "HA", "Shared", "Status", "State"}
	}
	if p.showK8s {
		header = append(header, "Pods")
	}
	if p.showLabels {
		header = append(header, "Labels")
	}

	return header
}

func (p *volumeGetFormatter) getLine(resp *api.SdkVolumeInspectResponse) ([]interface{}, error) {
	v := resp.GetVolume()
	spec := v.GetSpec()

	var line []interface{}

	state, err := p.pxVolumeOps.GetAttachedState(v)
	if err != nil {
		return line, err
	}

	// Size needs to be done better
	if p.wide {
		line = []interface{}{
			v.GetId(), v.GetLocator().GetName(), spec.GetSize() / Gi, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), spec.GetEncrypted(),
			spec.GetCos(), portworx.PrettyStatus(v), state, spec.GetSnapshotSchedule() != "",
		}
	} else {
		line = []interface{}{
			v.GetLocator().GetName(), spec.GetSize() / Gi, spec.GetHaLevel(),
			spec.GetShared() || spec.GetSharedv4(), portworx.PrettyStatus(v), state,
		}
	}
	if p.showK8s {
		pods, err := p.podsUsingVolume(v)
		if err != nil {
			return line, err
		}
		line = append(line, pods)
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line, nil
}

func (p *volumeGetFormatter) podsUsingVolume(v *api.Volume) (string, error) {
	usedPods, err := p.pxVolumeOps.PodsUsingVolume(v)
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
