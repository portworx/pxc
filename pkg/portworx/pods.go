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
	"fmt"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"
	v1 "k8s.io/api/core/v1"
)

type PodSpec struct {
	Namespace string
	Labels    map[string]string
}

type Pods interface {
	Objs
	// Get all pods for the namespace and labels specified
	GetPods() ([]v1.Pod, error)
	// PodsUsingVolume returns the list of pods using the given volume
	// The search is done on the array of pods given
	PodsUsingVolume(v *api.Volume) ([]v1.Pod, error)
	// GetContainerInfoForVolume will return the container info for pods using the volume
	GetContainerInfoForVolume(v *api.Volume) ([]kubernetes.ContainerInfo, error)
}

type pods struct {
	cops    kubernetes.COps
	podSpec *PodSpec
	pods    []v1.Pod
}

func NewPods(cops kubernetes.COps, podSpec *PodSpec) Pods {
	return &pods{
		cops:    cops,
		podSpec: podSpec,
	}
}

func (p *pods) Reset() {
	p.pods = make([]v1.Pod, 0)
}

func (p *pods) GetPods() ([]v1.Pod, error) {
	if len(p.pods) > 0 {
		return p.pods, nil
	}
	pods, err := p.cops.GetPodsByLabels(p.podSpec.Namespace,
		util.StringMapToCommaString(p.podSpec.Labels))
	if err != nil {
		return nil, err
	}
	p.pods = pods
	return pods, nil
}

func (p *pods) PodsUsingVolume(v *api.Volume) ([]v1.Pod, error) {
	pods, err := p.GetPods()
	if err != nil {
		return nil, err
	}
	usedPods := make([]v1.Pod, 0)
	namespace := v.Locator.VolumeLabels["namespace"]
	pvc := v.Locator.VolumeLabels["pvc"]
	if namespace == "" && pvc == "" {
		return usedPods, nil
	}
	for _, pod := range pods {
		if pod.Namespace == namespace {
			for _, volumeInfo := range pod.Spec.Volumes {
				if volumeInfo.PersistentVolumeClaim != nil {
					if volumeInfo.PersistentVolumeClaim.ClaimName == pvc {
						usedPods = append(usedPods, pod)
					}
				}
			}
		}
	}
	return usedPods, nil
}

func (p *pods) GetContainerInfoForVolume(
	v *api.Volume,
) ([]kubernetes.ContainerInfo, error) {
	cinfo := make([]kubernetes.ContainerInfo, 0)
	pods, err := p.PodsUsingVolume(v)
	if err != nil {
		return nil, err
	}

	if len(pods) == 0 {
		return cinfo, nil
	}

	labels := v.GetLocator().GetVolumeLabels()
	pvcName := labels["pvc"]
	// Should not happen since if the pvc label is not there we should not get any pods for it.
	if pvcName == "" {
		return nil, fmt.Errorf("Got a volume with no pvc lable")
	}
	// Namespace is already checked in PodsUsingVolume
	for _, p := range pods {
		volName := ""
		// Figure out the name of the volume referrenced in the pod
		for _, volumeInfo := range p.Spec.Volumes {
			if volumeInfo.PersistentVolumeClaim != nil {
				if volumeInfo.PersistentVolumeClaim.ClaimName == pvcName {
					volName = volumeInfo.Name
					break
				}
			}
		}
		for _, c := range p.Spec.Containers {
			// If the pvcs used as a VolumeMount in this container add the container
			// We use the name of the volume to figure this out
			for _, volMounts := range c.VolumeMounts {
				if volMounts.Name == volName {
					ci := kubernetes.ContainerInfo{
						Pod:       p,
						Container: c.Name,
						MountPath: volMounts.MountPath,
					}
					cinfo = append(cinfo, ci)
				}
			}
			// If the pvcs used as a VolumeDevice in this container add the container
			// We use the pvc name as used in the pod to figure this out
			for _, volDevices := range c.VolumeDevices {
				if volDevices.Name == pvcName {
					ci := kubernetes.ContainerInfo{
						Pod:       p,
						Container: c.Name,
					}
					cinfo = append(cinfo, ci)
				}
			}
		}
	}
	return cinfo, nil
}
