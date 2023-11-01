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
	"os/exec"
	"strings"
	"sync"

	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"

	"github.com/sirupsen/logrus"
)

// PortForwarder provides a way to forward a local port to Portworx endpoint
type PortForwarder interface {
	Endpoint() string
	Start() error
	Stop() error
}

// KubectlPortForwarder object
type KubectlPortForwarder struct {
	kubeconfig  string
	endpoint    string
	cmd         *exec.Cmd
	signhandler *util.SigIntManager
	lock        sync.Mutex
	running     bool
}

var (
	kubePortForwarder *KubectlPortForwarder
)

// StartTunnel starts the global tunnel to the Portworx endpoint through the Kubernetes service
func StartTunnel() error {
	if kubePortForwarder == nil {
		logrus.Info("Kubectl plugin mode detected")
		logrus.Infof("Port forwarder using kubeconfig %s", *config.KM().ConfigFlags().KubeConfig)
		kubePortForwarder = newKubectlPortForwarder(*config.KM().ConfigFlags().KubeConfig)
		if err := kubePortForwarder.Start(); err != nil {
			StopTunnel()
			return fmt.Errorf("Failed to setup port forward: %v", err)
		}
		config.CM().SetTunnelEndpoint(kubePortForwarder.Endpoint())
	}

	return nil
}

// StopTunnel stops the global tunnel to the Portworx endpoint through the Kubernetes service
func StopTunnel() {
	if kubePortForwarder != nil {
		kubePortForwarder.Stop()
		kubePortForwarder = nil
	}
}

// NewKubectlPortForwarder forwards a local port to the Portworx gRPC SDK endpoint
// through the Kubernetes API server using kubectl
// If kubeconfig is not provided, then kubectl will use the default kubeconfig
func NewKubectlPortForwarder(kubeconfig string) PortForwarder {
	return newKubectlPortForwarder(kubeconfig)
}

func newKubectlPortForwarder(kubeconfig string) *KubectlPortForwarder {
	return &KubectlPortForwarder{
		kubeconfig: kubeconfig,
	}
}

// Start creates the portforward using kubectl
func (p *KubectlPortForwarder) Start() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.running {
		return fmt.Errorf("Tunnel already running")
	}

	args := config.KM().KubectlFlagsToCliArgs()
	currentCluster := config.CM().GetCurrentCluster()
	logrus.Debugf("port-forward: CurrentCluster: %#v", *currentCluster)
	args = args + fmt.Sprintf("-n %s port-forward svc/%s :%s",
		currentCluster.TunnelServiceNamespace,
		currentCluster.TunnelServiceName,
		currentCluster.TunnelServicePort)
	logrus.Debugf("port-forward: running \"kubectl %s\"", args)

	cmd := exec.Command("kubectl", strings.Split(args, " ")...)

	// Setup to read port
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Errorf("Error while executing [%s]: %v", cmd.String(), err)
		return fmt.Errorf("Unable to setup kubectl: %v", err)
	}

	p.signhandler = util.NewSigIntManager(func() {
		p.Stop()
	})
	p.signhandler.Start()

	// Start the port forward process
	err = cmd.Start()
	if err != nil {
		logrus.Errorf("Error while executing [%s]: %v", cmd.String(), err)
		return fmt.Errorf("Unable to execute kubectl. Please make sure kubectl is in your path")
	}
	logrus.Debug("port-forward: started")
	p.cmd = cmd

	// Read the port
	buf := make([]byte, 1024, 1024)
	n, err := stdout.Read(buf[:])
	if err != nil || n < 0 {
		logrus.Warningf("Error: read[%d] from buffer: %v", n, err)
		return fmt.Errorf("Failed to setup connection to Portworx cluster to svc %s/%s port %s",
			currentCluster.TunnelServiceNamespace,
			currentCluster.TunnelServiceName,
			currentCluster.TunnelServicePort)
	}
	sbuf := string(buf[:n])

	// Set endpoint
	p.endpoint, err = p.getEndpointFromKubectlOutput(sbuf)
	if err != nil {
		p.Stop()
		return err
	}

	logrus.Infof("Connected to %s", p.endpoint)
	logrus.Debugf("Read %d bytes", n)
	logrus.Debugf("Output: %s", sbuf)

	p.running = true
	return nil
}

// Stop ends the session
func (p *KubectlPortForwarder) Stop() error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if !p.running {
		return nil
	}

	if p.cmd != nil {
		logrus.Debug("Port forwarding stopped")
		err := p.cmd.Process.Kill()
		p.cmd = nil
		return err
	}

	if p.signhandler != nil {
		p.signhandler.Stop()
		p.signhandler = nil
	}
	p.running = false
	return nil
}

// Endpoint returns the gRPC endpoint to the SDK
func (p *KubectlPortForwarder) Endpoint() string {
	return p.endpoint
}

func (p *KubectlPortForwarder) getEndpointFromKubectlOutput(sbuf string) (string, error) {
	index := strings.Index(sbuf, "127.0.0.1:")
	if index >= 0 {
		e := strings.Split(sbuf[index:], " ")[0]
		e = "localhost:" + strings.Split(e, ":")[1]
		return e, nil
	}

	index = strings.Index(sbuf, "[::1]:")
	if index >= 0 {
		return strings.Split(sbuf[index:], " ")[0], nil
	}

	logrus.Warningf("Unable to find 127.0.0.1 or [::1]: in [%s]", sbuf)
	return "", fmt.Errorf("Failed to determine endpoint information")
}
