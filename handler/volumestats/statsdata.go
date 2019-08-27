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
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/cliops"
	"github.com/portworx/px/pkg/tui"
)

type VolumeStats interface {
	tui.StatsModel
	// Returns the volumes that this Stats are for
	GetVolumes() []*api.Volume
	// Set if we need to show the sort marker in the column header
	ShowSortMarker(fg bool)
}

type ColNum int

const (
	VOL_NAME ColNum = iota
	BYTES_READ
	NUM_READS
	BYTES_WRITTEN
	NUM_WRITES
	IOPS
	IO_PROGRESS
	READ_TPUT
	WRITE_TPUT
	READ_LAT
	WRITE_LAT
)

const (
	DEFAULT_SORT_COLUMN = WRITE_TPUT
)

var (
	allHeaders = []string{
		"Name",
		"Read Bytes",
		"Reads",
		"Write Bytes",
		"Writes",
		"IOPS",
		"IODepth",
		"Read Tput",
		"Write Tput",
		"Read Lat",
		"Write Lat",
	}

	graphTitles = []string{
		"IOPS",
		"Read Throughput",
		"Write Throughput",
		"Avg. Read Latency",
		"Avg. Write Latency",
	}
)

type volumeStatsData struct {
	cvOps      *cliops.CliVolumeOps
	vols       []*api.Volume
	curStats   []*statsData
	curIndex   int
	sortInfo   *statsSorterInfo
	curTotal   *statsTotal
	sortMarker bool
}

type statsData struct {
	sortInfo     *statsSorterInfo
	volName      string
	bytesRead    uint64
	numReads     uint64
	bytesWritten uint64
	numWrites    uint64
	iops         uint64
	ioProgress   uint64
	readTput     uint64
	writeTput    uint64
	readLat      uint64
	writeLat     uint64
}

type statsTotal struct {
	statsData
	readMs  uint64
	writeMs uint64
}

func NewVolumeStats(
	cvOps *cliops.CliVolumeOps,
	resp []*api.SdkVolumeInspectResponse,
) VolumeStats {
	vsd := &volumeStatsData{
		cvOps: cvOps,
		vols:  make([]*api.Volume, len(resp)),
		sortInfo: &statsSorterInfo{
			ascending: false,
			column:    DEFAULT_SORT_COLUMN,
		},
		curTotal:   &statsTotal{},
		sortMarker: true,
	}
	for i, r := range resp {
		vsd.vols[i] = r.GetVolume()
	}
	return vsd
}

func (vsd *volumeStatsData) GetVolumes() []*api.Volume {
	return vsd.vols
}

func (vsd *volumeStatsData) ShowSortMarker(fg bool) {
	vsd.sortMarker = fg
}

func (vsd *volumeStatsData) Refresh() error {
	vsd.curStats = make([]*statsData, len(vsd.vols))
	vsd.curTotal = &statsTotal{}
	vsd.curIndex = 0
	for i, v := range vsd.vols {
		sd, err := vsd.getStats(v)
		if err != nil {
			return err
		}
		vsd.curStats[i] = sd
	}
	if vsd.curTotal.numReads != 0 {
		vsd.curTotal.readLat =
			(uint64)((vsd.curTotal.readMs * 1000) / vsd.curTotal.numReads)
	}
	if vsd.curTotal.numWrites != 0 {
		vsd.curTotal.writeLat =
			(uint64)((vsd.curTotal.writeMs * 1000) / vsd.curTotal.numWrites)
	}
	sort.Sort(statsSorter(vsd.curStats))
	return nil
}

func (vsd *volumeStatsData) GetTitle() string {
	return "Volume Stats (Press: q to quit; s to toggle sorting order; h|l to shift sort column); r to refresh"
}

func (vsd *volumeStatsData) GetHeaders() []string {
	tmp := make([]string, len(allHeaders))
	copy(tmp, allHeaders)
	if vsd.sortMarker == true {
		i := int(vsd.sortInfo.column)
		tmp[i] = fmt.Sprintf("*%s", tmp[i])
	}
	return tmp
}

func (vsd *volumeStatsData) NextRow() ([]string, error) {
	if vsd.curIndex >= len(vsd.vols) {
		return make([]string, 0), nil
	}

	sd := vsd.curStats[vsd.curIndex]
	vsd.curIndex++
	cols := make([]string, len(vsd.GetHeaders()))
	cols[0] = sd.volName
	cols[1] = humanize.Bytes(sd.bytesRead)
	cols[2] = fmt.Sprintf("%v", sd.numReads)
	cols[3] = humanize.Bytes(sd.bytesWritten)
	cols[4] = fmt.Sprintf("%v", sd.numWrites)
	cols[5] = fmt.Sprintf("%v", sd.iops)
	cols[6] = fmt.Sprintf("%v", sd.ioProgress)
	cols[7] = fmt.Sprintf("%v/s", humanize.Bytes(sd.readTput))
	cols[8] = fmt.Sprintf("%v/s", humanize.Bytes(sd.writeTput))
	cols[9] = getDurationStringRounded(sd.readLat)
	cols[10] = getDurationStringRounded(sd.writeLat)
	return cols, nil
}

// Ensure we don't get decimal places when displaying duration
func getDurationStringRounded(micro uint64) string {
	rn := micro * uint64(time.Microsecond)
	switch {
	case rn == 0:
		return fmt.Sprintf("0s")
	case rn < 2*uint64(time.Millisecond):
		return fmt.Sprintf("%vµs", micro)
	case rn < 2*uint64(time.Second):
		micro = uint64(math.Round(float64(micro) / 1000))
		return fmt.Sprintf("%vms", micro)
	default:
		x := time.Duration(rn)
		d := x.Round(time.Second)
		return d.String()
	}

}

func (vsd *volumeStatsData) SetSortInfo(colName string, ascending bool) {
	vsd.sortInfo.column = vsd.colNameToNum(colName)
	vsd.sortInfo.ascending = ascending
}

func (vsd *volumeStatsData) GetSortInfo() (string, bool) {
	return allHeaders[int(vsd.sortInfo.column)], vsd.sortInfo.ascending
}

func (vsd *volumeStatsData) MoveSortColumnNext() {
	index := int(vsd.sortInfo.column)
	if index == len(allHeaders)-1 {
		index = 0
	} else {
		index++
	}
	vsd.sortInfo.column = ColNum(index)
}

func (vsd *volumeStatsData) MoveSortColumnPrev() {
	index := int(vsd.sortInfo.column)
	if index == 0 {
		index = len(allHeaders) - 1
	} else {
		index--
	}
	vsd.sortInfo.column = ColNum(index)
}

func (vsd *volumeStatsData) colNameToNum(str string) ColNum {
	s := strings.TrimPrefix(str, "*")
	for i, r := range allHeaders {
		if r == s {
			return ColNum(i)
		}
	}
	return DEFAULT_SORT_COLUMN
}

func (vsd *volumeStatsData) getStats(v *api.Volume) (*statsData, error) {
	stats, err := vsd.cvOps.PxVolumeOps.GetStats(v, true)
	if err != nil {
		return nil, err
	}

	sd := &statsData{
		sortInfo: vsd.sortInfo,
	}

	sd.volName = v.GetLocator().GetName()
	sd.bytesRead = stats.GetReadBytes()
	sd.numReads = stats.GetReads()
	sd.bytesWritten = stats.GetWriteBytes()
	sd.numWrites = stats.GetWrites()
	sd.iops = iops(stats)
	sd.ioProgress = stats.GetIoProgress()
	sd.readTput = readThroughput(stats)
	sd.writeTput = writeThroughput(stats)
	sd.readLat = readLatency(stats)
	sd.writeLat = writeLatency(stats)

	vsd.curTotal.bytesRead += sd.bytesRead
	vsd.curTotal.numReads += sd.numReads
	vsd.curTotal.bytesWritten += sd.bytesWritten
	vsd.curTotal.numWrites += sd.numWrites
	vsd.curTotal.iops += sd.iops
	vsd.curTotal.ioProgress += sd.ioProgress
	vsd.curTotal.readTput += sd.readTput
	vsd.curTotal.writeTput += sd.writeTput
	vsd.curTotal.readMs += stats.GetReadMs()
	vsd.curTotal.writeMs += stats.GetWriteMs()
	return sd, nil
}

func toSec(ms uint64) uint64 {
	return ms / 1000
}

func iops(s *api.Stats) uint64 {
	intv := toSec(s.GetIntervalMs())
	if intv == 0 {
		return 0
	}
	return (s.GetWrites() + s.GetReads()) / intv
}

func writeThroughput(s *api.Stats) uint64 {
	intv := toSec(s.GetIntervalMs())
	if intv == 0 {
		return 0
	}
	return (s.GetWriteBytes()) / intv
}

func readThroughput(s *api.Stats) uint64 {
	intv := toSec(s.GetIntervalMs())
	if intv == 0 {
		return 0
	}
	return (s.GetReadBytes()) / intv
}

func readLatency(s *api.Stats) uint64 {
	if s.GetReads() == 0 {
		return 0
	}
	return (uint64)((s.GetReadMs() * 1000) / s.GetReads())
}

func writeLatency(s *api.Stats) uint64 {
	if s.GetWrites() == 0 {
		return 0
	}
	return (uint64)((s.GetWriteMs() * 1000) / s.GetWrites())
}

func (vsd *volumeStatsData) GetGraphTitle(index int) (string, error) {
	if index >= len(graphTitles) {
		return "", fmt.Errorf("Unknown index %d for graph title", index)
	}
	return graphTitles[index], nil
}

func (vsd *volumeStatsData) GetGraphData(index int) (float64, error) {
	switch index {
	case 0:
		return float64(vsd.curTotal.iops), nil
	case 1:
		return float64(vsd.curTotal.readTput), nil
	case 2:
		return float64(vsd.curTotal.writeTput), nil
	case 3:
		return float64(vsd.curTotal.readLat), nil
	case 4:
		return float64(vsd.curTotal.writeLat), nil
	}
	return 0.0, fmt.Errorf("Unknown index %d for graph title", index)
}

func (vsd *volumeStatsData) Humanize(index int, val float64) (string, error) {
	if val == 0.0 {
		return " ", nil
	}
	switch index {
	case 0:
		return fmt.Sprintf("%0.0f", val), nil
	case 1, 2:
		return fmt.Sprintf("%v/s", humanize.Bytes(uint64(val))), nil
	case 3, 4:
		return getDurationStringRounded(uint64(val)), nil
	}
	return "", fmt.Errorf("Unknown index %d for graph title", index)
}
