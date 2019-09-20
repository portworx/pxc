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

	"github.com/portworx/pxc/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	NODE_KEY = []byte("node=")
)

type kubeConnection struct {
	clientConfig clientcmd.ClientConfig
	clientSet    *kclikube.Clientset
}

func NewCOps(connect bool) (COps, error) {
	if connect == true {
		cc, cs, err := KubeConnectDefault()
		if err != nil {
			return nil, err
		}
		return &kubeConnection{
			clientConfig: cc,
			clientSet:    cs,
		}, nil
	}
	return &kubeConnection{}, nil
}

func (p *kubeConnection) Close() {
	// Nothing to do
}

func (p *kubeConnection) GetNamespace(s *string) (string, error) {
	if s != nil {
		return *s, nil
	}
	return p.GetDefaultNamespace()
}

func (p *kubeConnection) GetDefaultNamespace() (string, error) {
	ns, _, err := p.clientConfig.Namespace()
	if err != nil {
		return "", err
	}
	return ns, nil
}

func (p *kubeConnection) GetPodsByLabels(
	namespace string,
	labels string,
) ([]v1.Pod, error) {
	if p.clientSet == nil {
		return make([]v1.Pod, 0), nil
	}
	podClient := p.clientSet.CoreV1().Pods(namespace)
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

func (p *kubeConnection) GetPvcsByLabels(
	namespace string,
	labels string,
) ([]v1.PersistentVolumeClaim, error) {
	if p.clientSet == nil {
		return make([]v1.PersistentVolumeClaim, 0), nil
	}
	p.clientSet.CoreV1().Pods(namespace)
	pvcClient := p.clientSet.CoreV1().PersistentVolumeClaims(namespace)

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

func (p *kubeConnection) GetLogs(
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
		ret := p.clientSet.CoreV1().Pods(ci.Pod.Namespace).GetLogs(ci.Pod.Name,
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
	prefix := make([]byte, 0)
	if lo.ShowPodInfo == true {
		prefix = []byte(fmt.Sprintf("pod=%s namespace=%s ", lp.podName, lp.podNamespace))
	}
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
		if len(prefix) > 0 {
			_, err := out.Write(prefix)
			if err != nil {
				return err
			}
		}
		if bytes[0] == '@' {
			_, err := out.Write(NODE_KEY)
			if err != nil {
				return err
			}
			bytes = bytes[1:]
		}
	}
	_, err := out.Write(bytes)
	return err
}
