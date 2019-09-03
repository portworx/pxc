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
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	v1 "k8s.io/api/core/v1"
)

type PxPvc struct {
	Name      string
	Namespace string
	PodNames  []string
	Pvc       *v1.PersistentVolumeClaim
	PxVolume  *api.Volume
	Pods      []v1.Pod
}

func NewPxPvc(pvc *v1.PersistentVolumeClaim) *PxPvc {
	return &PxPvc{
		Name:      pvc.GetName(),
		Namespace: pvc.GetNamespace(),
		Pvc:       pvc,
		Pods:      make([]v1.Pod, 0),
		PodNames:  make([]string, 0),
	}
}

func (p *PxPvc) GetVolume() *api.Volume {
	return p.PxVolume
}

func (p *PxPvc) GetPodNames() []string {
	return p.PodNames
}

// SetVolume return true if found
func (p *PxPvc) SetVolume(vols []*api.Volume) bool {

	if p.Pvc == nil {
		return false
	}

	for _, volume := range vols {
		// Match by name
		if volume.GetLocator().GetName() == p.Pvc.Spec.VolumeName {
			p.PxVolume = volume
			return true
		}

		// Fallback to match by labels
		if labels := volume.GetLocator().GetVolumeLabels(); labels != nil {
			if p.Name == labels["pvc"] &&
				p.Namespace == labels["namespace"] {
				p.PxVolume = volume
				return true
			}
		}
	}

	return false
}

// SetPods return true if found
func (p *PxPvc) SetPods(pods []v1.Pod) bool {
	// get the pvc name
	for _, pod := range pods {
		if pod.GetNamespace() == p.Namespace {
			for _, volumeInfo := range pod.Spec.Volumes {
				if volumeInfo.PersistentVolumeClaim != nil {
					if volumeInfo.PersistentVolumeClaim.ClaimName == p.Name {
						p.PodNames = append(p.PodNames, p.Namespace+"/"+pod.Name)
						p.Pods = append(p.Pods, pod)
					}
				}
			}
		}
	}

	return len(p.PodNames) != 0
}
