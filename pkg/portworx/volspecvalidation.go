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
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

// ValidateVolumeSpec checks if a volume spec if valid. Currently due to volume driver limitation, only few
// combination of fields can be set as paramaters in VolumeSpecUpdate. This function will check for the same.
func ValidateVolumeSpec(volspec *api.VolumeSpecUpdate) error {
	// case of checking possible halevel flag combination
	if volspec.GetHaLevel() > 0 {
		if volspec.GetSize() > 0 || volspec.GetShared() || volspec.GetSticky() {
			// Please have unique msgs for each case so it's easy for use to identity the
			// flags mismatch combination.
			return fmt.Errorf("Invalid halevel flag combination. Size, Shared or Sticky flag not supported " +
				"with halevel flag")
		}
	}
	return nil
}
