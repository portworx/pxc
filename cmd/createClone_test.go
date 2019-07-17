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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPxCreateclone(t *testing.T) {
	volName := "testVol"
	cloneName := "cloneVol"

	// Create Volume
	volId := testCreateVolume(t, volName, 1)

	// Verify that the volume got created
	exists := testGetVolume(t, volId, volName)
	assert.Equal(t, exists, true, "Volume create failed")

	// Create clone
	cloneId := testCreateClone(t, volId, cloneName)

	//Verify that the clone got created
	exists = testGetVolume(t, cloneId, cloneName)

	// Delete volume. Only clone must exist
	testDeleteVolume(t, volId)
	vols, _ := testGetAllVolumes(t)
	assert.Equal(t, len(vols), 1, "Volume delete failed")

	// Delete clone, No volumes must remain
	testDeleteVolume(t, cloneId)
	vols, _ = testGetAllVolumes(t)
	assert.Equal(t, len(vols), 0, "Volume delete failed")
}
