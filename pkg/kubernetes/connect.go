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
	"os"

	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/contextconfig"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/sirupsen/logrus"
)

var (

	// KubeCliOpts is setup by cmd/root.go
	KubeCliOpts *genericclioptions.ConfigFlags
)

// KubeConnectDefault returns a Kubernetes client to the default
// or named context.
func KubeConnectDefault() (clientcmd.ClientConfig, *kubernetes.Clientset, error) {

	if InKubectlPluginMode() {
		logrus.Info("Setting up Kubernetes access in kubectl plugin mode")
		clientConfig := KubeCliOpts.ToRawKubeConfigLoader()

		r, err := KubeCliOpts.ToRESTConfig()
		if err != nil {
			return nil, nil, fmt.Errorf("unable to configure kubernetes client: %v", err)
		}

		clientSet, err := kubernetes.NewForConfig(r)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to connect to Kubernetes: %v", err)
		}
		return clientConfig, clientSet, nil
	}

	// TODO: Need to read the information from the command line using the Kubernetes APIs
	// to create a client set when running in kubectl plugin mode
	return KubeConnect(config.Get(config.File), config.Get(config.SpecifiedContext))
}

// KubeConnect will return a Kubernetes client using the kubeconfig file
// set in the default context.
// clientcmd.ClientConfig will allow the caller to call ClientConfig.Namespace() to get the namespace
// set by the caller on their Kubeconfig.
func KubeConnect(cfgFile, context string) (clientcmd.ClientConfig, *kubernetes.Clientset, error) {
	var (
		kubeconfig string
		pxctx      *contextconfig.ClientContext
		err        error
	)

	contextManager, err := contextconfig.NewContextManager(cfgFile)
	if err != nil {
		return nil, nil, err
	}
	if len(context) == 0 {
		pxctx, err = contextManager.GetCurrent()
	} else {
		pxctx, err = contextManager.GetNamedContext(context)
	}
	if err != nil {
		return nil, nil, err
	}
	if len(pxctx.Kubeconfig) == 0 {
		kubeconfig = os.Getenv("KUBECONFIG")
	} else {
		kubeconfig = pxctx.Kubeconfig
	}
	if len(kubeconfig) == 0 {
		return nil, nil, fmt.Errorf("no kubeconfig found in context %s", pxctx.Name)
	}

	// Get the client config
	cc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})
	r, err := cc.ClientConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to configure kubernetes client: %v", err)
	}
	// Get a client to the Kuberntes server
	clientset, err := kubernetes.NewForConfig(r)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to Kubernetes: %v", err)
	}

	return cc, clientset, nil
}
