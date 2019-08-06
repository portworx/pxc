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

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func addCommonLogOptions(lc *cobra.Command) {
	lc.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	lc.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	lc.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	lc.Flags().Bool("ignore-errors", false, "If watching / following Portworx logs, allow for any errors that occur to be non-fatal")
	lc.Flags().Int("max-log-requests", 5, "Specify maximum number of concurrent logs to follow. Defaults to 5.")
	lc.Flags().Int64("limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	lc.Flags().Int64("tail", portworx.NO_TAIL_LINES, "Lines of recent log file to display. Defaults to -1, showing all log lines")
	lc.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	lc.Flags().Duration("since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	lc.Flags().StringP("px-namespace", "n", "kube-system", "Kubernetes namespace in which Portworx is installed")
}

func parseRFC3339(s string) (metav1.Time, error) {
	if t, timeErr := time.Parse(time.RFC3339Nano, s); timeErr == nil {
		return metav1.Time{Time: t}, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return metav1.Time{}, err
	}
	return metav1.Time{Time: t}, nil
}

func getCommonLogOptions(cmd *cobra.Command) (*portworx.COpsLogOptions, error) {
	lo := &portworx.COpsLogOptions{}

	lo.PortworxNamespace, _ = cmd.Flags().GetString("px-namespace")
	lo.IgnoreLogErrors, _ = cmd.Flags().GetBool("ignore-errors")
	lo.MaxFollowConcurency, _ = cmd.Flags().GetInt("max-log-requests")
	if lo.MaxFollowConcurency <= 0 {
		return nil, fmt.Errorf("--max-log-requests should be greater than 0")
	}
	f, _ := cmd.Flags().GetString("filter")
	if len(f) > 0 {
		lo.Filters = strings.Split(f, ",")
		lo.ApplyFilters = true
	}

	lo.PodLogOptions.Follow, _ = cmd.Flags().GetBool("follow")
	lo.PodLogOptions.Timestamps, _ = cmd.Flags().GetBool("timestamps")
	lo.PodLogOptions.Previous, _ = cmd.Flags().GetBool("previous")

	lm, _ := cmd.Flags().GetInt64("limit-bytes")
	if lm != 0 {
		if lm < 0 {
			return nil, fmt.Errorf("--limit-bytes must be greater than 0")
		}
		lo.PodLogOptions.LimitBytes = &lm
	}

	tail, _ := cmd.Flags().GetInt64("tail")
	if tail != portworx.NO_TAIL_LINES {
		if tail < 0 {
			return nil, fmt.Errorf("TailLines must be greater than or equal to 0")
		}
		lo.PodLogOptions.TailLines = &tail
	}
	if lo.PodLogOptions.Follow == true && lo.PodLogOptions.TailLines == nil {
		x := portworx.DEFAULT_TAIL_LINES
		lo.PodLogOptions.TailLines = &x
	}

	sec, err := cmd.Flags().GetDuration("since")
	if err != nil {
		return nil, err
	}
	sinceSeconds := int64(sec.Round(time.Second).Seconds())
	sinceTime, _ := cmd.Flags().GetString("since-time")
	if len(sinceTime) > 0 && sinceSeconds > 0 {
		return nil, fmt.Errorf("at most one of --since or -since-time may be specified")
	}
	if len(sinceTime) > 0 {
		t, err := parseRFC3339(sinceTime)
		if err != nil {
			return nil, err
		}
		lo.PodLogOptions.SinceTime = &t
	}
	if sinceSeconds != 0 {
		if sinceSeconds < 0 {
			return nil, fmt.Errorf("--since must be greater than or equal to 0")
		}
		lo.PodLogOptions.SinceSeconds = &sinceSeconds
	}
	return lo, nil
}

func getRequiredPortworxPods(
	cvOps *cliVolumeOps,
	nodeNames []string,
	portworxNamespace string,
) ([]v1.Pod, error) {
	co := cvOps.pxVolumeOps.GetCOps()
	allPods, err := co.GetPodsByLabels(portworxNamespace, "name=portworx")
	if err != nil {
		return nil, err
	}
	if len(nodeNames) > 0 {
		selPods := make([]v1.Pod, 0, len(allPods))
		selPodNames := make([]string, 0, len(allPods))
		for _, p := range allPods {
			if util.ListContains(nodeNames, p.Spec.NodeName) == true {
				selPods = append(selPods, p)
				selPodNames = append(selPodNames, p.Spec.NodeName)
			}
		}
		for _, n := range nodeNames {
			if !util.ListContains(selPodNames, n) {
				return nil, fmt.Errorf("Node %s not found", n)
			}
		}
		return selPods, nil
	}
	return allPods, nil
}
