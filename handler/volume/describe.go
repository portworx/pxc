// Copyright © 2019 Portworx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package volume

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/handler/alerts"
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	prototime "github.com/portworx/pxc/pkg/openstorage/proto/time"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
)

var describeVolumeCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	// describeVolumeCmd represents the describeVolume command
	describeVolumeCmd = &cobra.Command{
		Use:     "inspect [NAME]",
		Aliases: []string{"describe"},
		Short:   "Describe a Portworx volume",
		Long:    "Show detailed information of Portworx volumes",
		Example: `
  # Describe all the volumes:
  pxc volume inspect

  # Describe specific volume called "abc":
  pxc volume inspect abc

  # Describe list of volumes (abc, xyz)
  pxc volume inspect abc xyz`,
		RunE: describeVolumesExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	VolumeAddCommand(describeVolumeCmd)
	describeVolumeCmd.Flags().String("owner", "", "Owner of volume")
	describeVolumeCmd.Flags().String("volumegroup", "", "Volume group id")
	describeVolumeCmd.Flags().Bool("deep", false, "Collect more information, this may delay the request")
	describeVolumeCmd.Flags().MarkHidden("deep")
	describeVolumeCmd.Flags().Bool("show-k8s-info", false, "Show kubernetes information")
})

func DescribeAddCommand(cmd *cobra.Command) {
	describeVolumeCmd.AddCommand(cmd)
}

func describeVolumesExec(cmd *cobra.Command, args []string) error {
	// Parse out all of the common cli volume flags
	cvi := cliops.NewCliInputs(cmd, args)

	// Create a CliOps object
	cliOps := cliops.NewCliOps(cvi)

	// Connect to pxc and k8s (if needed)
	err := cliOps.Connect()
	if err != nil {
		return err
	}
	defer cliOps.Close()

	// Create the parser object
	vcf := NewVolumeDescribeFormatter(cliOps)

	// Print details and return any errors found during parsing
	return util.PrintFormatted(vcf)
}

type VolumeDescribeFormatter struct {
	util.BaseFormatOutput
	cliOps  cliops.CliOps
	volumes portworx.Volumes
	nodes   portworx.Nodes
	pods    portworx.Pods
}

func NewVolumeDescribeFormatter(cliOps cliops.CliOps) *VolumeDescribeFormatter {
	volSpec := &portworx.VolumeSpec{
		VolNames: cliOps.CliInputs().Args,
		Labels:   cliOps.CliInputs().Labels,
	}
	d := &VolumeDescribeFormatter{
		cliOps:  cliOps,
		volumes: portworx.NewVolumes(cliOps.PxOps(), volSpec),
		nodes:   portworx.NewNodes(cliOps.PxOps(), &portworx.NodeSpec{}),
		pods:    portworx.NewPods(cliOps.COps(), &portworx.PodSpec{}),
	}
	d.FormatType = cliOps.CliInputs().FormatType
	return d
}

// DefaultFormat returns the default string representation of the object
func (p *VolumeDescribeFormatter) DefaultFormat() (string, error) {
	return p.toTabbed()
}

func (p *VolumeDescribeFormatter) toTabbed() (string, error) {
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

	for i, v := range vols {
		usedPods, err := p.pods.PodsUsingVolume(v)
		if err != nil {
			return "", err
		}
		err = p.AddVolumeDetails(v, t, usedPods)
		if err != nil {
			return "", err
		}
		// Put two empty lines between volumes
		if len(vols) > 1 && i != len(vols)-1 {
			t.AddLine("")
			t.AddLine("")
		}
	}
	t.Print()

	return b.String(), nil
}

func (p *VolumeDescribeFormatter) AddVolumeDetails(
	v *api.Volume,
	t *tabby.Tabby,
	usedPods []v1.Pod,
) error {
	err := p.addVolumeBasicInfo(v, t)
	if err != nil {
		return err
	}
	err = p.addVolumeStatsInfo(v, t)
	if err != nil {
		return err
	}
	err = p.addVolumeReplicationInfo(v, t)
	if err != nil {
		return err
	}
	err = p.addVolumeK8sInfo(v, t, usedPods)
	if err != nil {
		return err
	}
	p.addVolumeAccessInfo(v, t)
	p.addVolumeAlerts(v, t)

	return nil
}

func (p *VolumeDescribeFormatter) addVolumeAlerts(
	v *api.Volume,
	t *tabby.Tabby,
) {
	alertInfo := &cliops.CliAlertOps{}
	alertInfo.AlertType = "volume"
	alertInfo.ResourceId = v.GetId()
	alertInfo.PxAlertOps = portworx.NewPxAlertOps()
	f := alerts.NewAlertGetFormatter(alertInfo)

	t.AddLine("Alerts:")
	lines, err := f.DefaultFormat()
	if err != nil {
		t.AddLine("Unable to get alerts: %v", err)
	} else {
		t.AddLine(lines)
	}
}

func (p *VolumeDescribeFormatter) addVolumeAccessInfo(
	v *api.Volume,
	t *tabby.Tabby,
) {
	spec := v.GetSpec()

	// Get Owner
	owner := spec.GetOwnership().GetOwner()
	// Get Groups
	groups := spec.GetOwnership().GetAcls().GetGroups()
	// Get Collaborators
	collaborators := spec.GetOwnership().GetAcls().GetCollaborators()

	// If there is no ownership, then just return
	if spec.GetOwnership() == nil {
		return
	}

	t.AddLine("Ownership:")
	if len(owner) != 0 {
		t.AddLine("  Owner:", owner)
	}

	if len(groups) != 0 {
		t.AddLine("  Groups:")
		for key, value := range groups {
			t.AddLine(fmt.Sprintf("    %s:", key), value)
		}
	}

	if len(collaborators) != 0 {
		t.AddLine("  Collaborators:")
		for key, value := range collaborators {
			t.AddLine(fmt.Sprintf("    %s:", key), value)
		}
	}
}

func (p *VolumeDescribeFormatter) addVolumeBasicInfo(
	v *api.Volume,
	t *tabby.Tabby,
) error {
	spec := v.GetSpec()

	// Determine the state of the volume
	state, err := p.nodes.GetAttachedState(v)
	if err != nil {
		return err
	}

	// Print basic info
	t.AddLine("Volume:", v.GetId())
	t.AddLine("Name:", v.GetLocator().GetName())
	if util.InKubectlPluginMode() {
		labels := v.GetLocator().GetVolumeLabels()
		pvc := labels["pvc"]
		ns := labels["namespace"]
		if pvc != "" {
			t.AddLine("Pvc Name:", pvc)
		}
		if ns != "" {
			t.AddLine("Namespace:", ns)
		}
	}
	if v.GetGroup() != nil && len(v.GetGroup().GetId()) != 0 {
		t.AddLine("Group:", v.GetGroup().GetId())
	}
	if v.GetFormat() == api.FSType_FS_TYPE_FUSE {
		t.AddLine("Type:", "Namespace Volume Group")
		return nil
	}
	t.AddLine("Size:", humanize.BigIBytes(big.NewInt(int64(spec.GetSize()))))
	t.AddLine("Format:",
		strings.TrimPrefix(v.GetFormat().String(), "FS_TYPE_"))
	t.AddLine("HA:", spec.GetHaLevel())
	t.AddLine("IO Priority:", spec.GetCos())
	t.AddLine("Creation Time:",
		prototime.TimestampToTime(v.GetCtime()).Format(util.TimeFormat))
	if v.GetSource() != nil && len(v.GetSource().GetParent()) != 0 {
		t.AddLine("Parent:", v.GetSource().GetParent())
	}
	snapSched, err := portworx.SchedSummary(v)
	if err != nil {
		return err
	}
	if len(snapSched) != 0 {
		util.AddArray(t, "Snapshot Schedule:", snapSched)
	}
	if spec.GetStoragePolicy() != "" {
		t.AddLine("StoragePolicy:", spec.GetStoragePolicy())
	}
	t.AddLine("Shared:", portworx.SharedString(v))
	t.AddLine("Status:", portworx.PrettyStatus(v))
	t.AddLine("State:", state)
	attrs := portworx.BooleanAttributes(v)
	if len(attrs) != 0 {
		util.AddArray(t, "Attributes:", attrs)
	}
	if spec.GetScale() > 1 {
		t.AddLine("Scale:", v.Spec.Scale)
	}
	if v.GetAttachedOn() != "" && v.GetAttachedState() != api.AttachState_ATTACH_STATE_INTERNAL {
		t.AddLine("Device Path:", v.GetDevicePath())
	}
	if len(v.GetLocator().GetVolumeLabels()) != 0 {
		util.AddMap(t, "Labels:", v.GetLocator().GetVolumeLabels())
	}
	return nil
}

func (p *VolumeDescribeFormatter) addVolumeStatsInfo(
	v *api.Volume,
	t *tabby.Tabby,
) error {

	// Ignore the error if it cannot get stats
	stats, _ := p.volumes.GetStats(v, false)

	t.AddLine("Stats:")
	t.AddLine("  Reads:", stats.GetReads())
	t.AddLine("  Reads MS:", stats.GetReadMs())
	t.AddLine("  Bytes Read:", stats.GetReadBytes())
	t.AddLine("  Writes:", stats.GetWrites())
	t.AddLine("  Writes MS:", stats.GetWriteMs())
	t.AddLine("  Bytes Written:", stats.GetWriteBytes())
	t.AddLine("  IOs in progress:", stats.GetIoProgress())
	t.AddLine("  Bytes used:", humanize.BigIBytes(big.NewInt(int64(stats.GetBytesUsed()))))
	return nil
}

func (p *VolumeDescribeFormatter) addVolumeReplicationInfo(
	v *api.Volume,
	t *tabby.Tabby,
) error {
	replInfo, err := p.nodes.GetReplicationInfo(v)
	if err != nil {
		return err
	}
	t.AddLine("Replication Status:", replInfo.Status)
	if len(replInfo.Rsi) > 0 {
		t.AddLine("Replica sets on nodes:")
	}
	for _, rsi := range replInfo.Rsi {
		t.AddLine("  Set:", rsi.Id)
		util.AddArray(t, "    Node:", rsi.NodeInfo)
		if len(rsi.HaIncrease) > 0 {
			t.AddLine("    HA-Increase on:", rsi.HaIncrease)
		}
		if len(rsi.ReAddOn) > 0 {
			util.AddArray(t, "    Re-add on:", rsi.ReAddOn)
		}
	}
	return nil
}

func (p *VolumeDescribeFormatter) addVolumeK8sInfo(
	v *api.Volume,
	t *tabby.Tabby,
	usedPods []v1.Pod,
) error {
	if len(usedPods) > 0 {
		t.AddLine("Pods:")
		for _, consumer := range usedPods {
			t.AddLine("  - Name:", fmt.Sprintf("%s (%s)",
				consumer.GetName(), consumer.GetUID()))
			t.AddLine("    Namespace:", consumer.GetNamespace())
			t.AddLine("    Running on:", consumer.Spec.NodeName)
			o := make([]string, 0)
			for _, owner := range consumer.OwnerReferences {
				s := fmt.Sprintf("%s (%s)", owner.Name, owner.Kind)
				o = append(o, s)
			}
			util.AddArray(t, "    Controlled by:", o)
		}
	}
	return nil
}
