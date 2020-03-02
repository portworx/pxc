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

type statsSorterInfo struct {
	ascending bool
	column    ColNum
}

type statsSorter []*statsData

func (ss statsSorter) Len() int {
	return len(ss)
}

func (ss statsSorter) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss statsSorter) Less(i int, j int) bool {
	switch ss[i].sortInfo.column {
	case WRITE_LAT:
		return ss.writeLat(i, j)
	case READ_LAT:
		return ss.readLat(i, j)
	case WRITE_TPUT:
		return ss.writeTput(i, j)
	case READ_TPUT:
		return ss.readTput(i, j)
	case IO_PROGRESS:
		return ss.ioProgress(i, j)
	case IOPS:
		return ss.iops(i, j)
	case NUM_WRITES:
		return ss.numWrites(i, j)
	case BYTES_WRITTEN:
		return ss.bytesWritten(i, j)
	case NUM_READS:
		return ss.numReads(i, j)
	case BYTES_READ:
		return ss.bytesRead(i, j)
	case VOL_NAME:
		return ss.volName(i, j)
	}
	// Default to write throughput
	return ss.writeTput(i, j)
}

func (ss statsSorter) writeLat(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].writeLat < ss[j].writeLat
	} else {
		return ss[i].writeLat > ss[j].writeLat
	}
}

func (ss statsSorter) readLat(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].readLat < ss[j].readLat
	} else {
		return ss[i].readLat > ss[j].readLat
	}
}

func (ss statsSorter) writeTput(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].writeTput < ss[j].writeTput
	} else {
		return ss[i].writeTput > ss[j].writeTput
	}
}

func (ss statsSorter) readTput(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].readTput < ss[j].readTput
	} else {
		return ss[i].readTput > ss[j].readTput
	}
}

func (ss statsSorter) ioProgress(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].ioProgress < ss[j].ioProgress
	} else {
		return ss[i].ioProgress > ss[j].ioProgress
	}
}

func (ss statsSorter) iops(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].iops < ss[j].iops
	} else {
		return ss[i].iops > ss[j].iops
	}
}

func (ss statsSorter) numWrites(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].numWrites < ss[j].numWrites
	} else {
		return ss[i].numWrites > ss[j].numWrites
	}
}

func (ss statsSorter) bytesWritten(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].bytesWritten < ss[j].bytesWritten
	} else {
		return ss[i].bytesWritten > ss[j].bytesWritten
	}
}

func (ss statsSorter) numReads(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].numReads < ss[j].numReads
	} else {
		return ss[i].numReads > ss[j].numReads
	}
}

func (ss statsSorter) bytesRead(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].bytesRead < ss[j].bytesRead
	} else {
		return ss[i].bytesRead > ss[j].bytesRead
	}
}

func (ss statsSorter) volName(i int, j int) bool {
	if ss[i].sortInfo.ascending {
		return ss[i].volName < ss[j].volName
	} else {
		return ss[i].volName > ss[j].volName
	}
}
