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
	volName := genVolName("testVol")
	cloneName := genVolName("cloneVol")

	// Create Volume
	testCreateVolume(t, volName, 1)
	assert.True(t, testHasVolume(volName))

	// Create clone
	testCreateClone(t, volName, cloneName)
	assert.True(t, testHasVolume(cloneName))

	// Delete volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
	assert.True(t, testHasVolume(cloneName))

	// Delete clone
	testDeleteVolume(t, cloneName)
	assert.False(t, testHasVolume(volName))
	assert.False(t, testHasVolume(cloneName))
}
