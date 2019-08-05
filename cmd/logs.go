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
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

const (
	DEFAULT_TAIL_LINES = int64(100)
	NO_TAIL_LINES      = int64(-1)
)

var logsCmd *cobra.Command

var _ = RegisterCommandVar(func() {
	logsCmd = &cobra.Command{
		Use:   "logs",
		Short: "Print Portworx logs",
		Long:  "Show detailed information of Portworx volume for Kubernetes PVCs",
		Example: `$ px logs 
        Return Portworx logs from all nodes (last 100 lines from each node i there is more than 1 node)

        $ px logs abc
        Return Portworx logs from  node abc

        $ px logs -f  abc
        Begin streaming the Portworx logs from  node abc

        $ px logs --tail=20 abc
        Display only the most recent 20 lines of Portworx logs in  node abc

        $ px logs abc --filter "error,warning"
        Display all log lines that has either error or warning on node abc

        $ px logs --since=1h node
        Show all Portworx logs from node abc written in the last hour`,
		RunE: logsExec,
	}
})

// logsCmd represents the logs command
var _ = RegisterCommandInit(func() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.")
	logsCmd.Flags().Bool("timestamps", false, "Include timestamps on each line in the log output")
	logsCmd.Flags().BoolP("previous", "p", false, "If true, print the logs for the previous instance of the container in a pod if it exists.")
	logsCmd.Flags().Bool("ignore-errors", false, "If watching / following Portworx logs, allow for any errors that occur to be non-fatal")
	logsCmd.Flags().Int("max-log-requests", 5, "Specify maximum number of concurrent logs to follow. Defaults to 5.")
	logsCmd.Flags().Int64("limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	logsCmd.Flags().Int64("tail", NO_TAIL_LINES, "Lines of recent log file to display. Defaults to -1, showing all log lines")
	logsCmd.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().Duration("since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().StringP("namespace", "n", "kube-system", "Kubernetes namespace")
	logsCmd.Flags().Bool("all-namespaces", false, "Kubernetes namespace")
	logsCmd.Flags().String("filter", "", "comma seperated list of strings to search for. Log line will be printed if any one of the strings match")
})

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

type logOptions struct {
	ignoreLogErrors     bool
	maxFollowConcurency int
	podLogOptions       v1.PodLogOptions
	filters             []string
}

func getLogOptions(cmd *cobra.Command) (*logOptions, error) {
	lo := &logOptions{}

	lo.ignoreLogErrors, _ = cmd.Flags().GetBool("ignore-errors")
	lo.maxFollowConcurency, _ = cmd.Flags().GetInt("max-log-requests")
	if lo.maxFollowConcurency <= 0 {
		return nil, fmt.Errorf("--max-log-requests should be greater than 0")
	}
	f, _ := cmd.Flags().GetString("filter")
	if len(f) > 0 {
		lo.filters = strings.Split(f, ",")
	}

	lo.podLogOptions.Follow, _ = cmd.Flags().GetBool("follow")
	lo.podLogOptions.Timestamps, _ = cmd.Flags().GetBool("timestamps")
	lo.podLogOptions.Previous, _ = cmd.Flags().GetBool("previous")

	lm, _ := cmd.Flags().GetInt64("limit-bytes")
	if lm != 0 {
		if lm < 0 {
			return nil, fmt.Errorf("--limit-bytes must be greater than 0")
		}
		lo.podLogOptions.LimitBytes = &lm
	}

	tail, _ := cmd.Flags().GetInt64("tail")
	if tail != NO_TAIL_LINES {
		if tail < 0 {
			return nil, fmt.Errorf("TailLines must be greater than or equal to 0")
		}
		lo.podLogOptions.TailLines = &tail
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
		lo.podLogOptions.SinceTime = &t
	}
	if sinceSeconds != 0 {
		if sinceSeconds < 0 {
			return nil, fmt.Errorf("--since must be greater than or equal to 0")
		}
		lo.podLogOptions.SinceSeconds = &sinceSeconds
	}
	return lo, nil
}

func logsExec(cmd *cobra.Command, args []string) error {
	// Parse out all of the common cli volume flags
	cvi := GetCliVolumeInputs(cmd, make([]string, 0))
	cvi.showK8s = true
	cvi.GetNamespace(cmd)

	lo, err := getLogOptions(cmd)
	if err != nil {
		return err
	}

	// Create a cliVolumeOps object
	cvOps := NewCliVolumeOps(cvi)

	// Connect to px and k8s (if needed)
	err = cvOps.Connect()
	if err != nil {
		return err
	}
	defer cvOps.Close()

	lf := NewLogFormatter(cvOps, lo)

	return lf.GetLogs(args)
}

type logFormatter struct {
	cliVolumeOps
	lo *logOptions
}

func NewLogFormatter(
	cvOps *cliVolumeOps,
	lo *logOptions,
) *logFormatter {
	return &logFormatter{
		cliVolumeOps: *cvOps,
		lo:           lo,
	}
}

func (lf *logFormatter) getNeededPortworxPods(nodeNames []string) ([]v1.Pod, error) {
	opsInfo := lf.pxVolumeOps.GetPxVolumeOpsInfo()
	podClient := opsInfo.ClientSet.CoreV1().Pods(opsInfo.Namespace)
	podList, err := podClient.List(metav1.ListOptions{LabelSelector: "name=portworx"})
	if err != nil {
		return nil, err
	}
	allPods := podList.Items
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

func (lf *logFormatter) GetLogs(nodeNames []string) error {
	pods, err := lf.getNeededPortworxPods(nodeNames)
	if err != nil {
		return err
	}

	if len(pods) == 0 {
		util.Printf("No resources found\n")
		return nil
	}

	opsInfo := lf.pxVolumeOps.GetPxVolumeOpsInfo()
	rws := make([]rest.ResponseWrapper, 0, len(pods))
	if len(pods) > 1 && lf.lo.podLogOptions.TailLines == nil {
		x := DEFAULT_TAIL_LINES
		lf.lo.podLogOptions.TailLines = &x
	}
	for _, pod := range pods {
		ret := opsInfo.ClientSet.CoreV1().Pods(*lf.namespace).GetLogs(pod.Name, &lf.lo.podLogOptions)
		if err != nil {
			return err
		}
		rws = append(rws, ret)
	}

	if lf.lo.podLogOptions.Follow && len(rws) > 1 {
		if len(rws) > lf.lo.maxFollowConcurency {
			return fmt.Errorf(
				"you are attempting to follow %d log streams, but maximum allowed concurency is %d, use --max-log-requests to increase the limit",
				len(rws), lf.lo.maxFollowConcurency,
			)
		}

		return lf.writeLogsParallel(rws)
	}
	return lf.writeLogs(rws)
}

func (lf *logFormatter) writeLogsParallel(rws []rest.ResponseWrapper) error {
	reader, writer := io.Pipe()

	wg := &sync.WaitGroup{}
	wg.Add(len(rws))

	for _, rw := range rws {
		go func(rw rest.ResponseWrapper) {
			if err := lf.doWrite(rw, writer); err != nil {
				if !lf.lo.ignoreLogErrors {
					writer.CloseWithError(err)
					return
				}
				fmt.Fprintf(writer, "error: %v\n", err)
			}
			wg.Done()
		}(rw)
	}

	go func() {
		wg.Wait()
		writer.Close()
	}()

	_, err := io.Copy(util.Stdout, reader)
	return err
}

func (lf *logFormatter) writeLogs(rws []rest.ResponseWrapper) error {
	for _, rw := range rws {
		if err := lf.doWrite(rw, util.Stdout); err != nil {
			return err
		}
	}

	return nil
}

func (lf *logFormatter) doWrite(rw rest.ResponseWrapper, out io.Writer) error {
	rc, err := rw.Stream()
	if err != nil {
		return err
	}
	defer rc.Close()

	applyFilter := false
	if len(lf.lo.filters) > 0 {
		applyFilter = true
	}

	r := bufio.NewReader(rc)
	for {
		bytes, err := r.ReadBytes('\n')
		if e := lf.writeLine(bytes, out, applyFilter); e != nil {
			return e
		}

		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
}

func (lf *logFormatter) writeLine(bytes []byte, out io.Writer, applyFilter bool) error {
	if applyFilter == true {
		if !util.StringContains(string(bytes), lf.lo.filters) {
			return nil
		}
	}
	_, err := out.Write(bytes)
	return err
}
