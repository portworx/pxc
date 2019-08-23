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

package cliops

import (
	"fmt"
	"strings"
	"time"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PORTWORX_CONTAINER_NAME = "portworx"
)

func AddCommonLogOptions(lc *cobra.Command) {
	lc.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	lc.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	lc.Flags().Bool("show-pod-info", false, "Include pod info on each line in the log output")
	lc.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	lc.Flags().Bool("ignore-errors", false, "If watching / following Portworx logs, allow for any errors that occur to be non-fatal")
	lc.Flags().Int("max-log-requests", 5, "Specify maximum number of concurrent logs to follow. Defaults to 5.")
	lc.Flags().Int64("limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	lc.Flags().Int64("tail", kubernetes.NO_TAIL_LINES, "Lines of recent log file to work on. Defaults to -1, showing all log lines. All filters will be applied on top of these lines")
	lc.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	lc.Flags().Duration("since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	lc.Flags().StringP("px-namespace", "n", "kube-system", "Kubernetes namespace in which Portworx is installed")
	lc.Flags().String("filter", "", "Comma seperated list of strings to search for. Log line will be printed if any one of the strings match. Note that if --tail is specified the filter is applied on only those many lines.")
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

func GetCommonLogOptions(cmd *cobra.Command) (*kubernetes.COpsLogOptions, error) {
	lo := &kubernetes.COpsLogOptions{}

	lo.PortworxNamespace, _ = cmd.Flags().GetString("px-namespace")
	lo.IgnoreLogErrors, _ = cmd.Flags().GetBool("ignore-errors")
	lo.ShowPodInfo, _ = cmd.Flags().GetBool("show-pod-info")
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
	if tail != kubernetes.NO_TAIL_LINES {
		if tail < 0 {
			return nil, fmt.Errorf("TailLines must be greater than or equal to 0")
		}
		lo.PodLogOptions.TailLines = &tail
	}
	if lo.PodLogOptions.Follow == true && lo.PodLogOptions.TailLines == nil {
		x := kubernetes.DEFAULT_TAIL_LINES
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

// From the given lost of nodeName, figures out the Portworx pods on those nodes
func GetRequiredPortworxPods(
	cvOps *CliVolumeOps,
	nodeNames []string,
	portworxNamespace string,
) ([]kubernetes.ContainerInfo, error) {
	co := cvOps.PxVolumeOps.GetCOps()
	allPods, err := co.GetPodsByLabels(portworxNamespace, "name=portworx")
	if err != nil {
		return nil, err
	}

	allCinfo := make([]kubernetes.ContainerInfo, 0, len(allPods))
	for _, p := range allPods {
		cinfo := kubernetes.ContainerInfo{
			Pod:       p,
			Container: PORTWORX_CONTAINER_NAME,
		}
		allCinfo = append(allCinfo, cinfo)
	}

	if len(nodeNames) > 0 {
		selCInfo := make([]kubernetes.ContainerInfo, 0, len(allCinfo))
		selPodNames := make([]string, 0, len(allCinfo))
		for _, p := range allCinfo {
			if util.ListContains(nodeNames, p.Pod.Spec.NodeName) == true {
				selCInfo = append(selCInfo, p)
				selPodNames = append(selPodNames, p.Pod.Spec.NodeName)
			}
		}
		for _, n := range nodeNames {
			if !util.ListContains(selPodNames, n) {
				return nil, fmt.Errorf("Node %s not found", n)
			}
		}
		return selCInfo, nil
	}
	return allCinfo, nil
}

// This method looks at each of the volumes and figures out
// a) Which nodes the volume has relevance to such as where its replicas
//    are and where the node is attached
// b) Converts the node names to appropriate portworx pods
// c) Figures out which pods are using this volume and which of the
//    containers inside those pods use the volume
// Figures out all the unique namespace, pod and container combinations and returns those
func FillContainerInfo(
	vols []*api.SdkVolumeInspectResponse,
	cvOps *CliVolumeOps,
	lo *kubernetes.COpsLogOptions,
	allLogs bool,
) error {
	// Get All relevant pods.
	nodeNamesMap := make(map[string]bool)
	ciInfoList := make(map[string]kubernetes.ContainerInfo)
	for _, resp := range vols {
		// Get all of the nodes associated with the volume
		// Get all of the pods using the volume
		v := resp.GetVolume()
		err := cvOps.PxVolumeOps.GetAllNodesForVolume(v, nodeNamesMap)
		if err != nil {
			return err
		}

		cinfo, err := cvOps.PxVolumeOps.GetContainerInfoForVolume(v)

		if allLogs != true {
			lo.Filters = append(lo.Filters, v.GetLocator().GetName())
			lo.Filters = append(lo.Filters, v.GetId())
			labels := v.GetLocator().GetVolumeLabels()
			if pvcName, ok := labels["pvc"]; ok {
				lo.Filters = append(lo.Filters, pvcName)
			}
			lo.ApplyFilters = true
		}

		for _, ci := range cinfo {
			key := fmt.Sprintf("%s-%s-%s", ci.Pod.Namespace, ci.Pod.Name, ci.Container)
			ciInfoList[key] = ci
		}
	}

	nodeNames := make([]string, 0)
	for k, _ := range nodeNamesMap {
		nodeNames = append(nodeNames, k)
	}

	// Convert Portworx node names to pods
	cinfo, err := GetRequiredPortworxPods(cvOps, nodeNames, lo.PortworxNamespace)
	if err != nil {
		return err
	}

	// Remove duplicates between the list of pods that are attaching the volume and the portworx pods if any
	for _, ci := range cinfo {
		key := fmt.Sprintf("%s-%s-%s", ci.Pod.Namespace, ci.Pod.Name, ci.Container)
		ciInfoList[key] = ci
	}

	// Covert the pod map to an array of pods
	lo.CInfo = make([]kubernetes.ContainerInfo, 0)
	for _, ci := range ciInfoList {
		lo.CInfo = append(lo.CInfo, ci)
		if ci.MountPath != "" {
			lo.Filters = append(lo.Filters, ci.MountPath)
		}
	}

	return nil
}
