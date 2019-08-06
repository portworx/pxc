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

func (p *KubeConnectionData) GetLogs(
	lo *COpsLogOptions,
	out io.Writer,
) error {
	if len(lo.Pods) == 0 {
		util.Printf("No resources found\n")
		return nil
	}

	rws := make([]rest.ResponseWrapper, 0, len(lo.Pods))
	for _, pod := range lo.Pods {
		ret := p.ClientSet.CoreV1().Pods(lo.PortworxNamespace).GetLogs(pod.Name,
			&lo.PodLogOptions)
		rws = append(rws, ret)
	}

	if lo.PodLogOptions.Follow && len(rws) > 1 {
		if len(rws) > lo.MaxFollowConcurency {
			return fmt.Errorf(
				"you are attempting to follow %d log streams, but maximum allowed concurency is %d, use --max-log-requests to increase the limit",
				len(rws), lo.MaxFollowConcurency,
			)
		}

		return writeLogsParallel(rws, lo, out)
	}
	return writeLogs(rws, lo, out)
}

func writeLogsParallel(
	rws []rest.ResponseWrapper,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	reader, writer := io.Pipe()

	wg := &sync.WaitGroup{}
	wg.Add(len(rws))

	for _, rw := range rws {
		go func(rw rest.ResponseWrapper) {
			if err := doWrite(rw, lo, writer); err != nil {
				if !lo.IgnoreLogErrors {
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

	_, err := io.Copy(out, reader)
	return err
}

func writeLogs(
	rws []rest.ResponseWrapper,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	for _, rw := range rws {
		if err := doWrite(rw, lo, out); err != nil {
			return err
		}
	}
	return nil
}

func doWrite(
	rw rest.ResponseWrapper,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	rc, err := rw.Stream()
	if err != nil {
		return err
	}
	defer rc.Close()

	r := bufio.NewReader(rc)
	for {
		bytes, err := r.ReadBytes('\n')
		if e := writeLine(bytes, lo, out); e != nil {
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
	bytes []byte,
	lo *COpsLogOptions,
	out io.Writer,
) error {
	if lo.ApplyFilters == true {
		if !util.StringContainsAnyFromList(string(bytes), lo.Filters) {
			return nil
		}
	}
	_, err := out.Write(bytes)
	return err
}
