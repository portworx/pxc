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
	"fmt"

	"github.com/portworx/px/pkg/contextconfig"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// KubeConnect will return a Kubernetes client using the kubeconfig file
// set in the default context.
func KubeConnect(cfgFile string) (*kubernetes.Clientset, error) {
	pxctx, err := contextconfig.NewConfigReference(cfgFile).GetCurrent()
	if err != nil {
		return nil, err
	}
	if len(pxctx.Kubeconfig) == 0 {
		return nil, fmt.Errorf("No kubeconfig found in context %s\n", pxctx.Name)
	}

	r, err := clientcmd.BuildConfigFromFlags("", pxctx.Kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to configure kubernetes client: %v\n", err)
	}
	clientset, err := kubernetes.NewForConfig(r)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to Kubernetes: %v\n", err)
	}

	return clientset, nil
}
