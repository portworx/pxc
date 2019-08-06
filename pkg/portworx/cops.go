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
	"io"

	v1 "k8s.io/api/core/v1"
)

const (
	DEFAULT_TAIL_LINES = int64(10)
	NO_TAIL_LINES      = int64(-1)
)

type COpsLogOptions struct {
	PodLogOptions       v1.PodLogOptions
	IgnoreLogErrors     bool
	MaxFollowConcurency int
	Filters             []string
	ApplyFilters        bool
	PortworxNamespace   string
	Pods                []v1.Pod
}

type COps interface {
	// GetPodsByLabels returns pods from specified namespace with the given labels
	// labels should be of the form "abc=def,xyz=mno"
	GetPodsByLabels(namespace string, labels string) ([]v1.Pod, error)
	// GetPvcsByLabels returns pvcs from spacified namespace with given labels
	// labels should be of the form "abc=def,xyz=mno"
	GetPvcsByLabels(namespace string, labels string) ([]v1.PersistentVolumeClaim, error)
	// GetLogs writes out logs to out based on the logOptions specified
	GetLogs(logOptions *COpsLogOptions, out io.Writer) error
}
