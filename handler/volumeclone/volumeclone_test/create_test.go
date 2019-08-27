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
package volumeclone_test

import (
	"testing"

	"github.com/portworx/pxc/handler/test"
	"github.com/stretchr/testify/assert"
)

func TestPxCreateclone(t *testing.T) {
	volName := test.GenVolName("testCVol")
	cloneName := test.GenVolName("cloneCVol")

	// Create Volume
	test.PxTestCreateVolume(t, volName, 1)
	assert.True(t, test.PxTestHasVolume(volName))

	// Create clone
	test.PxTestCreateClone(t, volName, cloneName)
	assert.True(t, test.PxTestHasVolume(cloneName))

	// Delete volume
	test.PxTestDeleteVolume(t, volName)
	assert.False(t, test.PxTestHasVolume(volName))
	assert.True(t, test.PxTestHasVolume(cloneName))

	// Delete clone
	test.PxTestDeleteVolume(t, cloneName)
	assert.False(t, test.PxTestHasVolume(volName))
	assert.False(t, test.PxTestHasVolume(cloneName))
}
