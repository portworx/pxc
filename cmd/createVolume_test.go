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
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing creation of volume with "sticky" flag set.
func TestCreateStickyVolume(t *testing.T) {
	volName := genVolName("stickyVol")

	// Create volume with "sticky" flag set
	testCreateStickyVolume(t, volName, 1)
	assert.True(t, testHasVolume(volName))

	// Delete Volume
	// TODO: Place holder for unsetting sticky flag and then deleting it.
	// Will be done once update patch is in
	// testDeleteVolume(t, volName)
	// assert.False(t, testHasVolume(volName))
}

// Testing creation of volume with "encryption" flag set.
func TestCreateEncrypVolume(t *testing.T) {
	volName := genVolName("encrypVol")

	// Create volume with "sticky" flag set
	testCreateEncrypVolume(t, volName, 1)
	assert.True(t, testHasVolume(volName))

	// Delete Volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

// Testing creation of volume with "journal" flag set.
func TestCreateJournalVolume(t *testing.T) {
	volName := genVolName("journalVol")

	// Create volume with "journal" flag set
	testCreateJournalVolume(t, volName, 1)
	assert.True(t, testHasVolume(volName))

	// Delete Volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

// Testing creation of volume with "aggregation" flag set.
func TestCreateAggrVolume(t *testing.T) {
	volName := genVolName("aggrVol")
	aggrLevel := 2

	// Create volume with "journal" flag set
	testCreateAggrVolume(t, volName, 1, uint32(aggrLevel))
	assert.True(t, testHasVolume(volName))

	// Delete Volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

// Testing creation of volume with different IO profile.
func TestCreateIoProfVolume(t *testing.T) {
	ioProfile := []string{"cms", "db_remote", "sync_shared"}

	for _, profile := range ioProfile {
		volName := genVolName("ioprofVol")
		testCreateIoProfVolume(t, volName, 1, profile)
		assert.True(t, testHasVolume(volName))

		// Delete Volume
		testDeleteVolume(t, volName)
		assert.False(t, testHasVolume(volName))
	}
}
