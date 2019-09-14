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
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

type AlertType int64

type AlertSpec struct {
	Severity     api.SeverityType
	ResourceType api.ResourceType
	Description  string
	Name         string
	Uniq         bool
}

var resourceToString = map[api.ResourceType]string{
	api.ResourceType_RESOURCE_TYPE_NONE:    "UNKNOWN RESOURCE",
	api.ResourceType_RESOURCE_TYPE_DRIVE:   "DRIVE",
	api.ResourceType_RESOURCE_TYPE_NODE:    "NODE",
	api.ResourceType_RESOURCE_TYPE_CLUSTER: "CLUSTER",
	api.ResourceType_RESOURCE_TYPE_VOLUME:  "VOLUME",
}

// TypeToSpec fetches info about an alert. In PX a specific alert is always linked with
// its corresponding resource type, while, from OSD perspective, an alert type and
// resource type are two separate entities.
func TypeToSpec() map[AlertType]AlertSpec {
	return typeToSpec
}

func GetResourceTypeString(resourceType api.ResourceType) string {
	return resourceToString[resourceType]
}

func SeverityString(severity api.SeverityType) string {
	switch severity {
	case api.SeverityType_SEVERITY_TYPE_ALARM:
		return "ALARM"
	case api.SeverityType_SEVERITY_TYPE_WARNING:
		return "WARN"
	default:
		return "NOTIFY"
	}
}
