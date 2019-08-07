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
package kubernetes

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/portworx/px/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeConnectionData struct {
	ClientConfig clientcmd.ClientConfig
	ClientSet    *kclikube.Clientset
}

func NewCOps(kc *KubeConnectionData) COps {
	return kc
}

func (p *KubeConnectionData) GetPodsByLabels(
	namespace string,
	labels string,
) ([]v1.Pod, error) {
	if p.ClientSet == nil {
		return make([]v1.Pod, 0), nil
	}
	podClient := p.ClientSet.CoreV1().Pods(namespace)
	lo := metav1.ListOptions{}
	if len(labels) > 0 {
		lo.LabelSelector = labels
	}
	podList, err := podClient.List(lo)
	if err != nil {
		return nil, err
	}
	return podList.Items, nil
}

func (p *KubeConnectionData) GetPvcsByLabels(
	namespace string,
	labels string,
) ([]v1.PersistentVolumeClaim, error) {
	if p.ClientSet == nil {
		return make([]v1.PersistentVolumeClaim, 0), nil
	}
	p.ClientSet.CoreV1().Pods(namespace)
	pvcClient := p.ClientSet.CoreV1().PersistentVolumeClaims(namespace)

	lo := metav1.ListOptions{}
	if len(labels) > 0 {
		lo.LabelSelector = labels
	}
	pvcList, err := pvcClient.List(lo)
	if err != nil {
		return nil, err
	}

	return pvcList.Items, nil
}

type logPayload struct {
	rw           rest.ResponseWrapper
	podName      string
	podNamespace string
}

func (p *KubeConnectionData) GetLogs(
	lo *COpsLogOptions,
	out io.Writer,
) error {
	if len(lo.CInfo) == 0 {
		util.Printf("No resources found\n")
		return nil
	}

	lps := make([]*logPayload, 0, len(lo.CInfo))
	for _, ci := range lo.CInfo {
		// We don't need a deep copy since we are only modifying the container name
		optionsCopy := lo.PodLogOptions
		optionsCopy.Container = ci.Container
		ret := p.ClientSet.CoreV1().Pods(ci.Pod.Namespace).GetLogs(ci.Pod.Name,
			&optionsCopy)

		lp := &logPayload{
			rw:           ret,
			podName:      ci.Pod.Name,
			podNamespace: ci.Pod.Namespace,
		}
		lps = append(lps, lp)
	}

	if lo.PodLogOptions.Follow && len(lps) > 1 {
		if len(lps) > lo.MaxFollowConcurency {
			return fmt.Errorf(
				"you are attempting to follow %d log streams, but maximum allowed concurency is %d, use --max-log-requests to increase the limit",
				len(lps), lo.MaxFollowConcurency,
			)
		}

		return writeLogsParallel(lps, lo, out)
	}
	return writeLogs(lps, lo, out)
}

func writeLogsParallel(
	lps []*logPayload,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	reader, writer := io.Pipe()

	wg := &sync.WaitGroup{}
	wg.Add(len(lps))

	for _, lp := range lps {
		go func(lp *logPayload) {
			if err := doWrite(lp, lo, writer); err != nil {
				if !lo.IgnoreLogErrors {
					writer.CloseWithError(err)
					return
				}
				fmt.Fprintf(writer, "error: %v\n", err)
			}
			wg.Done()
		}(lp)
	}

	go func() {
		wg.Wait()
		writer.Close()
	}()

	_, err := io.Copy(out, reader)
	return err
}

func writeLogs(
	lps []*logPayload,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	for _, lp := range lps {
		if err := doWrite(lp, lo, out); err != nil {
			return err
		}
	}
	return nil
}

func doWrite(
	lp *logPayload,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	prefix := []byte(fmt.Sprintf("%s: %s: ", lp.podName, lp.podNamespace))
	rc, err := lp.rw.Stream()
	if err != nil {
		return err
	}
	defer rc.Close()

	r := bufio.NewReader(rc)
	for {
		bytes, err := r.ReadBytes('\n')
		if e := writeLine(prefix, bytes, lo, out); e != nil {
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

func writeLine(
	prefix []byte,
	bytes []byte,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	if lo.ApplyFilters == true {
		if !util.StringContainsAnyFromList(string(bytes), lo.Filters) {
			return nil
		}
	}
	if len(bytes) > 0 {
		_, err := out.Write(prefix)
		if err != nil {
			return err
		}
	}
	_, err := out.Write(bytes)
	return err
}
