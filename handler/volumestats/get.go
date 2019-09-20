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
package volumestats

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cheynewallace/tabby"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/cliops"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/tui"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var getVolumeStatsCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	getVolumeStatsCmd = &cobra.Command{
		Use:   "volumestats",
		Short: "Get stats of Portworx volumes",
		RunE:  getVolumeStatsExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.GetAddCommand(getVolumeStatsCmd)
	headers := strings.Join(allHeaders, "|")
	sortMsg := fmt.Sprintf("Specify one of '%s' to sort on", headers)
	getVolumeStatsCmd.Flags().StringP("selector", "l", "", "Selector (label query) comma-separated name=value pairs")
	getVolumeStatsCmd.Flags().StringP("output", "o", "", "Output in yaml|json|wide")
	getVolumeStatsCmd.Flags().String("sort-on", allHeaders[int(WRITE_TPUT)], sortMsg)
	getVolumeStatsCmd.Flags().String("sort-order", "desc", "Sort in ascending or descending order. Specify one of asc|desc")
	getVolumeStatsCmd.Flags().BoolP("watch", "w", false, "Monitor stats at a periodic interval")
	getVolumeStatsCmd.Flags().DurationP("interval", "i", time.Second*2, "Specify refresh interval")
	getVolumeStatsCmd.Flags().Bool("no-graphs", false, "Don't show graphs")
	getVolumeStatsCmd.Flags().MarkHidden("no-graphs")
})

func GetAddCommand(cmd *cobra.Command) {
	getVolumeStatsCmd.AddCommand(cmd)
}

func getVolumeStatsExec(cmd *cobra.Command, args []string) error {
	// Check if any flag value is not provided, error out
	err := cliops.ValidateCliInput(cmd, args)
	if err != nil {
		return err
	}

	sortOn, _ := cmd.Flags().GetString("sort-on")
	if err != nil {
		return err
	}

	if util.ListContains(allHeaders, sortOn) == false {
		return fmt.Errorf("Unknown column %s to sort on", sortOn)
	}

	sortOrder, _ := cmd.Flags().GetString("sort-order")
	if err != nil {
		return err
	}

	so := false
	switch sortOrder {
	case "asc":
		so = true
	case "desc":
		so = false
	default:
		return fmt.Errorf("sort-order should be one of asc or desc")
	}

	// Parse out all of the common cli volume flags
	cvi := cliops.NewCliInputs(cmd, args)

	// Create a cliOps object
	cliOps := cliops.NewCliOps(cvi)

	// Connect to pxc and k8s (if needed)
	err = cliOps.Connect()
	if err != nil {
		return err
	}
	defer cliOps.Close()

	volSpec := &portworx.VolumeSpec{
		VolNames: cliOps.CliInputs().Args,
		Labels:   cliOps.CliInputs().Labels,
	}

	vs := portworx.NewVolumes(cliOps.PxOps(), volSpec)

	vols, err := vs.GetVolumes()
	if err != nil {
		return err
	}

	if len(vols) == 0 {
		util.Printf("No resources found\n")
		return nil
	}

	vsd := NewVolumeStats(cliOps, vs)
	vsd.SetSortInfo(sortOn, so)
	watch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return err
	}

	if watch {
		vsd.ShowSortMarker(true)
		return doWatch(cmd, cliOps, vsd)
	} else {
		vsd.ShowSortMarker(false)
		volStatsFormatter := NewVolumeStatsGetFormatter(cliOps, vsd)
		return util.PrintFormatted(volStatsFormatter)
	}
}

func NewVolumeStatsGetFormatter(
	cliOps cliops.CliOps,
	vsd VolumeStats,
) *volumeStatsGetFormatter {
	v := &volumeStatsGetFormatter{
		cliOps: cliOps,
		vsd:    vsd,
	}
	v.FormatType = cliOps.CliInputs().FormatType
	return v
}

type volumeStatsGetFormatter struct {
	util.BaseFormatOutput
	cliOps cliops.CliOps
	vsd    VolumeStats
}

func (p *volumeStatsGetFormatter) getStats() (map[string]*api.Stats, error) {
	stats := make(map[string]*api.Stats)
	vols, err := p.vsd.GetVolumes()
	if err != nil {
		return nil, err
	}
	for _, v := range vols {
		s, err := p.cliOps.PxOps().GetStats(v, true)
		if err != nil {
			return nil, err
		}
		stats[v.GetLocator().GetName()] = s
	}
	return stats, nil
}

// YamlFormat returns the yaml representation of the object
func (p *volumeStatsGetFormatter) YamlFormat() (string, error) {
	stats, err := p.getStats()
	if err != nil {
		return "", err
	}
	return util.ToYaml(stats)
}

// JsonFormat returns the json representation of the object
func (p *volumeStatsGetFormatter) JsonFormat() (string, error) {
	stats, err := p.getStats()
	if err != nil {
		return "", err
	}
	return util.ToJson(stats)
}

func (p *volumeStatsGetFormatter) WideFormat() (string, error) {
	return p.toTabbed()
}

// DefaultFormat returns the default string representation of the object
func (p *volumeStatsGetFormatter) DefaultFormat() (string, error) {
	return p.toTabbed()
}

func (p *volumeStatsGetFormatter) toTabbed() (string, error) {
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(writer)

	err := p.vsd.Refresh()
	if err != nil {
		return "", nil
	}

	h := p.vsd.GetHeaders()
	hi := make([]interface{}, len(h))
	for i, _ := range h {
		hi[i] = h[i]
	}
	t.AddHeader(hi...)
	for {
		line, err := p.vsd.NextRow()
		if err != nil {
			return "", nil
		}
		if len(line) == 0 {
			break
		}
		l := make([]interface{}, len(h))
		for i, _ := range line {
			l[i] = line[i]
		}
		t.AddLine(l...)
	}
	t.Print()
	return b.String(), nil
}

func doWatch(
	cmd *cobra.Command,
	cliOps cliops.CliOps,
	vsd VolumeStats,
) error {
	// Get all the watch specific flags
	interval, err := cmd.Flags().GetDuration("interval")
	if err != nil {
		return err
	}
	if interval < 2*time.Second {
		return fmt.Errorf("--interval should not be less than 2s")
	}

	noGraphs, err := cmd.Flags().GetBool("no-graphs")
	if err != nil {
		return err
	}

	numGraphs := 5
	if noGraphs == true {
		numGraphs = 0
	}
	tv := tui.NewStatsView(numGraphs)
	return tv.Display(vsd, interval)
}
