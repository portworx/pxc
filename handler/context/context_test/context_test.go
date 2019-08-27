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

package context_test

import (
	"testing"

	"github.com/portworx/pxc/handler/test"
	"github.com/portworx/pxc/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMultipleClustersContextsFromCli(t *testing.T) {
	// This test depends on the ./hack/config.yml to have
	// a 'source' and 'target' context and two running mock-sdk-servers

	// Create a volume on the source that is now in the target.
	// We will use this to differentiate them.
	vol := test.GenVolName("mysourcevolume")
	test.PxTestCreateVolume(t, vol, 1)

	// Get the volume info on the source
	assert.True(t, test.PxTestHasVolume(vol))

	// Fail to get that information on the target
	lines, _, err := test.ExecuteCli("px --context=target get volumes")
	assert.NoError(t, err)
	assert.False(t, util.ListContainsSubString(lines, vol))

	// Delete volume
	test.PxTestDeleteVolume(t, vol)
}
