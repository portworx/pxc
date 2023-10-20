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
package util

import (
	"strings"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

const (
	statusPrefix = "STATUS_"
)

// SdkStatusToPrettyString returns a human readable version of the Sdk Status
func SdkStatusToPrettyString(status api.Status) string {

	if status == api.Status_STATUS_OK {
		return "Ready"
	}

	s := strings.TrimLeft(status.String(), statusPrefix)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ToLower(s)
	s = strings.Title(s)

	return s
}
