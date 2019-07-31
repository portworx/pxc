/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    httd://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/portworx/px/pkg/util"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	volName   string
	volId     string
	cloneName string
	cloneId   string
	snapName  string
	snapId    string
}

func testCreateAll(t *testing.T, td *testData) {
	// Create Volume
	td.volId = testCreateVolume(t, td.volName, 1)

	// Verify that the volume got created
	exists := testGetVolume(t, td.volId, td.volName)
	assert.Equal(t, exists, true, "Volume create failed")

	// Create clone
	td.cloneId = testCreateClone(t, td.volId, td.cloneName)

	//Verify that the clone got created
	exists = testGetVolume(t, td.cloneId, td.cloneName)
	assert.Equal(t, exists, true, "Clone create failed")

	// Create Snapshot
	td.snapId = testCreateSnapshot(t, td.volId, td.snapName)

	//Verify that the clone got created
	exists = testGetVolume(t, td.snapId, td.snapName)
	assert.Equal(t, exists, true, "Clone create failed")
}

func testDeleteAll(t *testing.T, td *testData) {
	// Delete volume. Only clone and snapshot must exist
	testDeleteVolume(t, td.volId)
	vols, _ := testGetAllVolumes(t)
	assert.Equal(t, util.ListContains(vols, td.volId), false, "Volume delete failed")
	assert.Equal(t, util.ListContains(vols, td.snapId), true, "Volume delete failed")
	assert.Equal(t, util.ListContains(vols, td.cloneId), true, "Volume delete failed")

	// Delete clone, Only snapshot must exist
	testDeleteVolume(t, td.cloneId)
	vols, _ = testGetAllVolumes(t)
	assert.Equal(t, util.ListContains(vols, td.cloneId), false, "Volume delete failed")
	assert.Equal(t, util.ListContains(vols, td.snapId), true, "Volume delete failed")

	// Delete clone, Only snapshot must exist
	testDeleteVolume(t, td.snapId)
	vols, _ = testGetAllVolumes(t)
	assert.Equal(t, util.ListContains(vols, td.snapId), false, "Volume delete failed")
}

func getKeyValue(s string) (string, string) {
	x := strings.Split(s, ":")
	x[0] = strings.Trim(x[0], " ")
	x[1] = strings.Trim(x[1], " ")
	return x[0], x[1]
}

func verifyKeyValue(
	t *testing.T,
	key string,
	value string,
	expectedKey string,
	expectedValue string,
) {
	assert.Equal(t, key, expectedKey, "key not correct")
	assert.Equal(t, value, expectedValue, "value not correct")
}

func verifyVolumeDescription(
	t *testing.T,
	volName string,
	volId string,
	parent string,
	desc string,
) {
	d := strings.Split(desc, "\n")
	if d[0] == "" {
		d = d[1:]
	}
	index := 0
	k, v := getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Volume", volId)
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Name", volName)
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Size", "1.0 GiB")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Format", "XFS")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "HA", "1")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "IO Priority", "NONE")
	index++
	k, v = getKeyValue(d[index])
	// Dont check value of creation time
	verifyKeyValue(t, k, "", "Creation Time", "")
	index++
	if parent != "" {
		k, v = getKeyValue(d[index])
		verifyKeyValue(t, k, v, "Parent", parent)
		index++
	}
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Shared", "no")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Status", "UP")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "State", "Detached")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Stats", "")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Reads", "12345")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Reads MS", "1")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Bytes Read", "1234567")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Writes", "9876")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Writes MS", "2")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Bytes Written", "7654321")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "IOs in progress", "987")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Bytes used", "1.1 GiB")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Replication Status", "Detached")
}

func testDescribeListedVolumes(t *testing.T, td *testData) {
	v := make([]string, 3)
	v[0] = td.volId
	v[1] = td.snapId
	v[2] = td.cloneId
	desc := testDescribeVolumes(t, v)
	for _, d := range desc {
		switch d {
		case td.volId:
			verifyVolumeDescription(t, td.volName, td.volId, "", d)
		case td.snapId:
			verifyVolumeDescription(t, td.snapName, td.snapId, td.volId, d)
		case td.cloneId:
			verifyVolumeDescription(t, td.cloneName, td.cloneId, td.volId, desc[2])
		}
	}
}

func testDescribeAllVolumes(t *testing.T, td *testData) {
	desc := testDescribeVolumes(t, make([]string, 0))
	assert.Equal(t, len(desc) >= 3, true, "Got wrong number of volumes")
	for _, d := range desc {
		dd := strings.Split(d, "\n")
		if len(dd) == 1 {
			continue
		}
		if dd[0] == "" {
			dd = dd[1:]
		}
		_, v := getKeyValue(dd[0])
		switch v {
		case td.volId:
			verifyVolumeDescription(t, td.volName, td.volId, "", d)
		case td.snapId:
			verifyVolumeDescription(t, td.snapName, td.snapId, td.volId, d)
		case td.cloneId:
			verifyVolumeDescription(t, td.cloneName, td.cloneId, td.volId, d)
		}
	}
}

func TestDescribeVolume(t *testing.T) {
	r := getRandom()
	td := &testData{
		volName:   fmt.Sprintf("%v-%v", "testVol", r),
		cloneName: fmt.Sprintf("%v-%v", "cloneVol", r),
		snapName:  fmt.Sprintf("%v-%v", "snapVol", r),
	}

	testCreateAll(t, td)
	testDescribeListedVolumes(t, td)
	testDescribeAllVolumes(t, td)
	testDeleteAll(t, td)
}
