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
package portworx

import (
	"reflect"
	"testing"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/util"
	"github.com/stretchr/testify/assert"
)

var (
	expectedSchedule = map[string][]string{
		"tp1": []string{
			"periodic 20m0s,keep last 7",
			"daily @10:00,keep last 10",
			"weekly Monday@01:25,keep last 3",
			"monthly 1@00:30,keep last 3",
		},
		"tp2":      []string{},
		"tp2-snap": []string{},
		"tp3":      []string{},
		"pvc-6fc1fe2d-25f4-40b0-a616-04c019572154": []string{},
		"pvc-34d0f15c-65b9-4229-8b3e-b7bb912e382f": []string{},
	}
	expectedAttributes = map[string][]string{
		"tp1":      []string{},
		"tp2":      []string{},
		"tp2-snap": []string{"read-only", "sticky"},
		"tp3":      []string{},
		"pvc-6fc1fe2d-25f4-40b0-a616-04c019572154": []string{},
		"pvc-34d0f15c-65b9-4229-8b3e-b7bb912e382f": []string{},
	}
	expectedShared = map[string]string{
		"tp1":      "false",
		"tp2":      "false",
		"tp2-snap": "false",
		"tp3":      "false",
		"pvc-6fc1fe2d-25f4-40b0-a616-04c019572154": "false",
		"pvc-34d0f15c-65b9-4229-8b3e-b7bb912e382f": "true",
	}
)

func testVolumeCommon(t *testing.T, volOps PxVolumeOps, v *api.Volume) {
	name := v.GetLocator().GetName()
	snapSchedule, err := SchedSummary(v)
	assert.Equal(t, err, nil, "Got error parsing snapshot schedule")
	eSnapSchedule := expectedSchedule[name]
	for _, s := range snapSchedule {
		assert.Equalf(t, util.ListContains(eSnapSchedule, s), true, "Wrong schedule summary for %s", name)
	}
	sharedString := SharedString(v)
	eSharedString := expectedShared[name]
	assert.Equalf(t, sharedString, eSharedString, "Wrong shared flag for %s", name)
	ba := BooleanAttributes(v)
	eba := expectedAttributes[name]
	assert.Equalf(t, reflect.DeepEqual(ba, eba), true, "Wrong attributes for %s", name)
	y := TrueOrFalse(true)
	assert.Equal(t, y, "true", "bool translation not correct")
	n := TrueOrFalse(false)
	assert.Equal(t, n, "false", "bool translation not correct")
}

func TestVolumeCommon(t *testing.T) {
	volOps := testGetPxVolumeOps(t)
	svols, err := volOps.GetVolumes()
	assert.Equal(t, err, nil, "Could not get volumes")
	for _, sv := range svols {
		v := sv.GetVolume()
		testVolumeCommon(t, volOps, v)
	}
}
