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
	"fmt"
	"testing"

	"github.com/portworx/px/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestPxDeleteVolume(t *testing.T) {
	volName := fmt.Sprintf("%v-%v", "testVol", getRandom())

	// Create Volume
	volId := testCreateVolume(t, volName, 1)

	// Verify that the volume got created
	exists := testGetVolume(t, volId, volName)
	assert.Equal(t, exists, true, "Volume create failed")

	// Delete Volume
	testDeleteVolume(t, volId)

	// Verify that volume got deleted
	vols, _ := testGetAllVolumes(t)
	assert.Equal(t, util.ListContains(vols, volId), false, "Volume delete failed")

	// Delete it again to ensure we don't get an error
	testDeleteVolume(t, volId)
}
