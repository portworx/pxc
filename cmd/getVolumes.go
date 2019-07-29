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
	vf, err := newVolumeFormatter(cmd, args)
	if err != nil {
		return err
	}
	defer vf.close()

	vcf := volumeGetFormatter{
		volumeFormatter: *vf,
	}
	vcf.Print()
	return nil
}

type volumeGetFormatter struct {
	volumeFormatter
}

// String returns the formatted output of the object as per the format set.
func (p *volumeGetFormatter) String() string {
	return util.GetFormattedOutput(p)
}

// Print writes the object to stdout
func (p *volumeGetFormatter) Print() {
	util.Printf("%v\n", p)
}

// YamlFormat returns the yaml representation of the object
func (p *volumeGetFormatter) YamlFormat() string {
	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		util.Eprintf("%v\n", err)
		return ""
	}
	return util.ToYaml(vols)
}

// JsonFormat returns the json representation of the object
func (p *volumeGetFormatter) JsonFormat() string {
	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		util.Eprintf("%v\n", err)
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

	// Start the columns
	t.AddHeader(p.getHeader()...)

	vols, err := p.pxVolumeOps.GetVolumes()
	if err != nil {
		util.Eprintf("%v\n", err)
		return ""
	}
	for _, n := range vols {
		t.AddLine(p.getLine(n)...)
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

func (p *volumeGetFormatter) getLine(resp *api.SdkVolumeInspectResponse) []interface{} {

	v := resp.GetVolume()
	spec := v.GetSpec()

	state := p.pxVolumeOps.GetAttachedState(v)

	// Size needs to be done better
	var line []interface{}
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
		line = append(line, p.podsUsingVolume(v))
	}
	if p.showLabels {
		line = append(line, util.StringMapToCommaString(v.GetLocator().GetVolumeLabels()))
	}
	return line
}

func (p *volumeGetFormatter) podsUsingVolume(v *api.Volume) string {
	usedPods := p.pxVolumeOps.PodsUsingVolume(v)
	usedPodNames := make([]string, 0)
	namespace := v.Locator.VolumeLabels["namespace"]
	for _, pod := range usedPods {
		usedPodNames = append(usedPodNames, namespace+"/"+pod.Name)
	}
	return strings.Join(usedPodNames, ",")
}
