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

	_ "github.com/portworx/px/pkg/util"
	"github.com/stretchr/testify/assert"
)

/*
To test successful case.
*/
func TestGetVolumeWithLabels(t *testing.T) {
	//creating volumes with lables
	volName := genVolName("labelvol")
	labels := "type=labelvol"

	// Create Volume
	testCreateVolumeWithLabel(t, volName, 1, labels)
	assert.True(t, testHasVolume(volName))

	// Delete the created volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

/*
To test not providing labels along withj --labels flag.
*/
func TestGetVolumeWithoutLabels(t *testing.T) {
	//creating volumes with lables
	volName := genVolName("labelvol")
	labels := "type1=labelvol"

	// Create Volume
	testCreateVolumeWithLabel(t, volName, 1, labels)
	testHasVolumeWithEmptyLabels(t, volName)

	// Delete the created volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

/*
To test if provided lable is not present.
*/
func TestGetVolumeDummyLabels(t *testing.T) {
	//creating volumes with lables
	volName := genVolName("labelvol")
	labels := "type1=" + volName
	dummyLabels := "name=dummy"

	// Create Volume
	testCreateVolumeWithLabel(t, volName, 1, labels)
	so, _ := testHasVolumeWithLabels(t, dummyLabels)
	assert.Contains(t, so.String(), "No resources found")

	// Delete the created volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}

/*
Test to error is inavlid (k,v) label pair is provided.
*/
func TestGetVolumeInvalidLabels(t *testing.T) {
	//creating volumes with lables
	volName := genVolName("labelvol")
	labels := "type1=" + volName
	invalidLabel := "type1," + volName

	// Create Volume
	testCreateVolumeWithLabel(t, volName, 1, labels)
	_, err := testHasVolumeWithLabels(t, invalidLabel)
	assert.Contains(t, err.Error(), "invalid pair")

	// Delete the created volume
	testDeleteVolume(t, volName)
	assert.False(t, testHasVolume(volName))
}
