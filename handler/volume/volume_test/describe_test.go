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
package volume_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/portworx/px/handler/test"
	"github.com/portworx/px/pkg/util"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	volName   string
	cloneName string
	snapName  string
}

func testCreateAll(t *testing.T, td *testData) {
	// Create Volume
	test.PxTestCreateVolume(t, td.volName, 1)
	assert.True(t, test.PxTestHasVolume(td.volName))

	// Create clone
	test.PxTestCreateClone(t, td.volName, td.cloneName)
	assert.True(t, test.PxTestHasVolume(td.cloneName))

	// Create Snapshot
	test.PxTestCreateSnapshot(t, td.volName, td.snapName)
	assert.True(t, test.PxTestHasVolume(td.snapName))
}

func testDeleteAll(t *testing.T, td *testData) {
	// Delete volume. Only clone and snapshot must exist
	test.PxTestDeleteVolume(t, td.volName)

	vols := test.PxTestGetAllVolumes(t)
	assert.False(t, util.ListContains(vols, td.volName), "Volume delete failed")
	assert.True(t, util.ListContains(vols, td.snapName), "Volume delete failed")
	assert.True(t, util.ListContains(vols, td.cloneName), "Volume delete failed")

	// Delete clone, Only snapshot must exist
	test.PxTestDeleteVolume(t, td.cloneName)
	vols = test.PxTestGetAllVolumes(t)
	assert.False(t, util.ListContains(vols, td.volName), "Volume delete failed")
	assert.False(t, util.ListContains(vols, td.cloneName), "Volume delete failed")
	assert.True(t, util.ListContains(vols, td.snapName), "Volume delete failed")

	// Delete clone, Only snapshot must exist
	test.PxTestDeleteVolume(t, td.snapName)
	vols = test.PxTestGetAllVolumes(t)
	assert.False(t, util.ListContains(vols, td.volName), "Volume delete failed")
	assert.False(t, util.ListContains(vols, td.cloneName), "Volume delete failed")
	assert.False(t, util.ListContains(vols, td.snapName), "Volume delete failed")
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
	parent string,
	desc string,
) {
	d := strings.Split(desc, "\n")
	if d[0] == "" {
		d = d[1:]
	}
	index := 0
	k, v := getKeyValue(d[index])
	vInfo := test.PxTestVolumeInfo(t, v)
	verifyKeyValue(t, k, v, "Volume", vInfo.GetId())
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Name", volName)
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Size", "1.0 GiB")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Format", "EXT4")
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
		parentInfo := test.PxTestVolumeInfo(t, parent)
		verifyKeyValue(t, k, v, "Parent", parentInfo.GetId())
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
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Ownership", "")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Groups", "")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "group1", "Read")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "Collaborators", "")
	index++
	k, v = getKeyValue(d[index])
	verifyKeyValue(t, k, v, "user1", "Write")
}

// Takes a list of volumes and returns a array of string, one volume description per string
func testDescribeVolumes(t *testing.T, volNames []string) ([]string, error) {
	cli := "px describe volume"
	for _, v := range volNames {
		cli = fmt.Sprintf("%v %v", cli, v)
	}
	lines, _, err := test.ExecuteCli(cli)
	if err != nil {
		return make([]string, 0), err
	}

	l := strings.Join(lines, "\n")
	vols := strings.Split(l, "\n\n")
	return vols, nil
}

func testDescribeListedVolumes(t *testing.T, td *testData) {
	v := make([]string, 3)
	v[0] = td.volName
	v[1] = td.snapName
	v[2] = td.cloneName
	desc, err := testDescribeVolumes(t, v)
	assert.NoError(t, err)
	for _, d := range desc {
		dd := strings.Split(d, "\n")
		if len(dd) == 1 {
			continue
		}
		if dd[0] == "" {
			dd = dd[1:]
		}
		_, x := getKeyValue(dd[1])
		switch x {
		case td.volName:
			verifyVolumeDescription(t, td.volName, "", d)
		case td.snapName:
			verifyVolumeDescription(t, td.snapName, td.volName, d)
		case td.cloneName:
			verifyVolumeDescription(t, td.cloneName, td.volName, desc[2])
		}
	}
}

func testDescribeAllVolumes(t *testing.T, td *testData) {
	desc, err := testDescribeVolumes(t, make([]string, 0))
	assert.NoError(t, err)
	assert.Equal(t, len(desc) >= 3, true, "Got wrong number of volumes")
	for _, d := range desc {
		dd := strings.Split(d, "\n")
		if len(dd) == 1 {
			continue
		}
		if dd[0] == "" {
			dd = dd[1:]
		}
		_, v := getKeyValue(dd[1])
		switch v {
		case td.volName:
			verifyVolumeDescription(t, td.volName, "", d)
		case td.snapName:
			verifyVolumeDescription(t, td.snapName, td.volName, d)
		case td.cloneName:
			verifyVolumeDescription(t, td.cloneName, td.volName, d)
		}
	}
}

func testDescribeNonExistantVolume(t *testing.T, td *testData) {
	v := make([]string, 2)
	v[0] = "nonexistent-1"
	v[1] = "nonexistent-2"
	_, err := testDescribeVolumes(t, v)
	assert.Error(t, err)
}

func TestDescribeVolume(t *testing.T) {
	td := &testData{
		volName:   test.GenVolName("testVol"),
		cloneName: test.GenVolName("cloneVol"),
		snapName:  test.GenVolName("snapVol"),
	}

	testCreateAll(t, td)
	testDescribeListedVolumes(t, td)
	testDescribeAllVolumes(t, td)
	testDescribeNonExistantVolume(t, td)
	testDeleteAll(t, td)
}
