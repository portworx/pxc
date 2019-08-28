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
package volume_test

import (
	"testing"

	"github.com/portworx/pxc/handler/test"
	"github.com/stretchr/testify/assert"
)

// TestPatchVolumeHalevel runs a series of volume patching test
func TestPatchVolumeHalevel(t *testing.T) {
	volName := test.GenVolName("testVol")
	haLevel := 2
	volCreate(t, volName)
	// Now update halevel to 2
	test.PxTestPatchVolumeHalevel(t, volName, haLevel)
	volCleanup(t, volName)
}

func TestPatchVolumeResize(t *testing.T) {
	volName := test.GenVolName("testVol")
	var size uint64
	// Setting size to 2GB
	size = 2
	volCreate(t, volName)

	// Now update halevel to 2
	test.PxTestPatchVolumeResize(t, volName, size)
	volCleanup(t, volName)
}

func TestPatchVolumeShared(t *testing.T) {
	volName := test.GenVolName("testVol")
	sharedOn := "on"

	volCreate(t, volName)
	// Now update shared to true
	test.PxTestPatchVolumeShared(t, volName, sharedOn)
	volCleanup(t, volName)
}

func TestPatchVolumeUnsetShared(t *testing.T) {
	volName := test.GenVolName("testVol")
	sharedOn := "off"
	sharedOff := "on"

	volCreate(t, volName)
	// Now update shared to true
	test.PxTestPatchVolumeShared(t, volName, sharedOn)
	//Now unset shared aka to false
	test.PxTestPatchVolumeShared(t, volName, sharedOff)
	volCleanup(t, volName)
}

func TestPatchVolumeAddCollaborators(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	collaborators := "user2:r,user3:w"
	//Update the collaborators list to the volume access list.
	test.PxTestPatchVolumeAddCollaborators(t, volName, collaborators)
	volCleanup(t, volName)
}

func TestPatchVolumeRemoveCollaborators(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	collaborators := "user1:w"
	//Remove the collaborators list from the volume access list.
	test.PxTestPatchVolumeRemoveCollaborators(t, volName, collaborators)
	volCleanup(t, volName)
}

func TestPatchVolumeRemoveAllCollaborators(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	//Remove all the collaborators from the volume access list.
	test.PxTestPatchVolumeRemoveAllCollaborators(t, volName)
	volCleanup(t, volName)
}

func TestPatchVolumeAddGroups(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	groups := "group2:r,group3:a"
	//Update the group list to the volume access list.
	test.PxTestPatchVolumeAddGroups(t, volName, groups)
	volCleanup(t, volName)
}

func TestPatchVolumeRemoveGroups(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	groups := "group1:r"
	//Remove the group list from the volume access list.
	test.PxTestPatchVolumeRemoveGroups(t, volName, groups)
	volCleanup(t, volName)
}

func TestPatchVolumeRemoveAllGroups(t *testing.T) {
	volName := test.GenVolName("testVol")
	volCreate(t, volName)
	//Remove All the groups from the volume access list.
	test.PxTestPatchVolumeRemoveAllGroups(t, volName)
	volCleanup(t, volName)
}

// Helper to create a volume
func volCreate(t *testing.T, volName string) {
	// Create a volume
	test.PxTestCreateVolume(t, volName, 1)
	// Verify that the volume got created
	assert.True(t, test.PxTestHasVolume(volName))
}

// Helper function to cleanup volume created
func volCleanup(t *testing.T, volName string) {
	// Delete Volume
	test.PxTestDeleteVolume(t, volName)
}
