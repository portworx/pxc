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
	_ "github.com/portworx/pxc/pkg/util"
	"github.com/stretchr/testify/assert"
)

//To test successful case.
func TestGetVolumeWithLabels(t *testing.T) {
	//creating volumes with lables
	volName := test.GenVolName("labelvol")
	selector := "type=labelvol"

	// Create Volume with label
	test.PxTestCreateVolumeWithLabel(t, volName, 1, selector)
	assert.True(t, test.PxTestHasVolume(volName))

	// Delete the created volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

// Test to error out when --selector is provided along with volume name
func TestGetVolumeWithNameSelector(t *testing.T) {
	//creating volumes with lables
	volName := test.GenVolName("labelvol")
	selector := "type1=labelvol"

	// Create Volume with label
	test.PxTestCreateVolumeWithLabel(t, volName, 1, selector)
	test.PxTestGetVolumeWithNameSelector(t, volName, selector)

	// Delete the created volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

//Test passing k,v pair which is not present
func TestGetVolumeWithDummySelector(t *testing.T) {
	//creating volumes with lables
	volName := test.GenVolName("labelvol")
	selector := "type1=labelvol"
	dummySelector := "invalid=label"

	// Create Volume with label
	test.PxTestCreateVolumeWithLabel(t, volName, 1, selector)
	so, _ := test.PxTestGetVolumeWithLabels(t, dummySelector)
	assert.Contains(t, so.String(), "No resources found")

	// Delete the created volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}

//Test to error is inavlid (k,v) label pair is provided.
func TestGetVolumeInvalidLabels(t *testing.T) {
	//creating volumes with lables
	volName := test.GenVolName("labelvol")
	selector := "type1=labelvol"
	invalidSelector := "type1,labelvol"

	// Create Volume
	test.PxTestCreateVolumeWithLabel(t, volName, 1, selector)
	_, err := test.PxTestGetVolumeWithLabels(t, invalidSelector)
	assert.Error(t, err)

	// Delete the created volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
}
