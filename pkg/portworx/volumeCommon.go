// Copyright © 2019 Portworx
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package portworx

import (
	"strings"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/openstorage/sched"
)

const (
	timeLayout = "Jan 2 15:04:05 UTC 2006"
)

// SchedSummary returns the formatted string version of the schedule
func SchedSummary(v *api.Volume) ([]string, error) {
	schedule := v.GetSpec().GetSnapshotSchedule()
	sspec, policies, err := sched.ParseScheduleAndPolicies(schedule)
	if err != nil {
		return make([]string, 0), err
	}
	return scheduleSummary(sspec, policies), nil
}

func scheduleSummary(items []sched.RetainInterval, policyTags *sched.PolicyTags) []string {
	summary := make([]string, 0)
	if policyTags != nil {
		for _, t := range policyTags.Names {
			summary = append(summary, "policy="+t)
		}
	}
	if len(items) == 0 {
		return summary
	}
	for _, iv := range items {
		summary = append(summary, iv.String())
	}
	return summary
}

// SharedString returns the string representation of the shared flag of a volume
func SharedString(v *api.Volume) string {
	if v.Spec.Sharedv4 {
		return "v4"
	}
	return YesOrNo(v.Spec.Shared)
}

// YesOrNo returns the string representation of bool
func YesOrNo(b bool) string {
	if b {
		return "yes"
	} else {
		return "no"
	}
}

// BoolenaAttributes returns the string representations of all of the boolian attribute flags
func BooleanAttributes(v *api.Volume) []string {
	attrs := make([]string, 0, 3)
	if v.Readonly {
		attrs = append(attrs, "read-only")
	}
	if v.Spec.Encrypted {
		attrs = append(attrs, "encrypted")
	}
	if v.Spec.Sticky {
		attrs = append(attrs, "sticky")
	}
	if v.Spec.Compressed {
		attrs = append(attrs, "compressed")
	}
	return attrs
}

// PrettyStatus trims out the VOLUME_STATUS_ prefix
func PrettyStatus(v *api.Volume) string {
	return strings.TrimPrefix(v.GetStatus().String(), "VOLUME_STATUS_")
}

func getState(v *api.Volume, node *api.StorageNode) string {
	state := "Detached"
	if v.State == api.VolumeState_VOLUME_STATE_ATTACHED {
		if node != nil {
			if v.AttachedState == api.AttachState_ATTACH_STATE_EXTERNAL {
				state = "on " + node.GetHostname()
			}
		} else {
			state = "Attached"
		}
	} else if v.State == api.VolumeState_VOLUME_STATE_DETATCHING {
		if node != nil {
			state = "Was on " + node.GetHostname()
		} else {
			state = "Detaching"
		}
	}
	return state
}
