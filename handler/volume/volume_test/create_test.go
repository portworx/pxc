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
	"testing"

	"github.com/portworx/pxc/handler/test"
	"github.com/stretchr/testify/assert"
)

// Testing creation of volume with "sticky" flag set.
func TestCreateStickyVolume(t *testing.T) {
	volName := test.GenVolName("stickyVol")

	// Create volume with "sticky" flag set
	test.PxTestCreateStickyVolume(t, volName, 1)
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete Volume
	// TODO: Place holder for unsetting sticky flag and then deleting it.
	// Will be done once update patch is in
	// test.PxTestDeleteVolume(t, volName)
	// assert.False(t, test.PxTestHasVolume(volName))
}

// Testing creation of volume with "encryption" flag set.
func TestCreateEncrypVolume(t *testing.T) {
	volName := test.GenVolName("encrypVol")

	// Create volume with "sticky" flag set
	test.PxTestCreateEncrypVolume(t, volName, 1)
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete Volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

// Testing creation of volume with "journal" flag set.
func TestCreateJournalVolume(t *testing.T) {
	volName := test.GenVolName("journalVol")

	// Create volume with "journal" flag set
	test.PxTestCreateJournalVolume(t, volName, 1)
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete Volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

// Testing creation of volume with "aggregation" flag set.
func TestCreateAggrVolume(t *testing.T) {
	volName := test.GenVolName("aggrVol")
	aggrLevel := 2

	// Create volume with "journal" flag set
	test.PxTestCreateAggrVolume(t, volName, 1, uint32(aggrLevel))
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete Volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

// Testing creation of volume with different IO profile.
func TestCreateIoProfVolume(t *testing.T) {
	ioProfile := []string{"cms", "db_remote", "sync_shared"}

	for _, profile := range ioProfile {
		volName := test.GenVolName("ioprofVol")
		test.PxTestCreateIoProfVolume(t, volName, 1, profile)
		assert.True(t, test.PxTestHasVolume(volName))

		// Delete Volume
		test.PxTestDeleteVolume(t, volName)
		assert.False(t, test.PxTestHasVolume(volName))
	}
}

// Testing creation of volume with access (--groups and --collaborators) flag set.
func TestCreateVolumeWithAccess(t *testing.T) {
	volName := test.GenVolName("accessVol")

	// Create volume with access (--groups and --collaborators) flag set
	test.PxTestCreateVolumeWithAccess(t, volName, 1, "group1:r,group2:a", "user1:r,user2:w")
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete Volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}
