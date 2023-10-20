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
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"

	v1 "k8s.io/api/core/v1"
)

type PvcSpec struct {
	Namespace string
	Labels    map[string]string
}

type Pvcs interface {
	Objs
	// GetPvcs returns the Pvcs as per the spec provided
	GetPvcs() ([]v1.PersistentVolumeClaim, error)
	// GetPxPvcs returns the list of PxPvcs
	GetPxPvcs() ([]*kubernetes.PxPvc, error)
}

type pvcs struct {
	// pxops object
	pxops PxOps
	// K8S connection data
	cops    kubernetes.COps
	vols    Volumes
	pods    Pods
	pvcSpec *PvcSpec
	pvcs    []v1.PersistentVolumeClaim
	pxPvcs  []*kubernetes.PxPvc
}

func NewPvcs(pxops PxOps, cops kubernetes.COps, pvcSpec *PvcSpec) Pvcs {
	podSpec := &PodSpec{
		Namespace: pvcSpec.Namespace,
	}
	volLabels := make(map[string]string)
	// If a namespace is specified only look for volumes
	// that has the namespace=pvcSpec.Namespace
	if pvcSpec.Namespace != "" {
		volLabels["namespace"] = pvcSpec.Namespace
	}
	volSpec := &VolumeSpec{
		Labels: volLabels,
	}
	return &pvcs{
		pxops:   pxops,
		cops:    cops,
		pvcSpec: pvcSpec,
		vols:    NewVolumes(pxops, volSpec),
		pods:    NewPods(cops, podSpec),
	}
}

func (p *pvcs) Reset() {
	p.vols.Reset()
	p.pods.Reset()
	p.pvcs = make([]v1.PersistentVolumeClaim, 0)
	p.pxPvcs = make([]*kubernetes.PxPvc, 0)
}

func (p *pvcs) GetPvcs() ([]v1.PersistentVolumeClaim, error) {
	if p.pvcs != nil {
		return p.pvcs, nil
	}

	pvcs, err := p.cops.GetPvcsByLabels(p.pvcSpec.Namespace,
		util.StringMapToCommaString(p.pvcSpec.Labels))
	if err != nil {
		return nil, err
	}

	p.pvcs = pvcs
	return p.pvcs, nil
}

func (p *pvcs) GetPxPvcs() ([]*kubernetes.PxPvc, error) {
	if p.pxPvcs != nil {
		return p.pxPvcs, nil
	}

	k8sPvcs, err := p.GetPvcs()
	if err != nil {
		return nil, err
	}

	vols, err := p.vols.GetVolumes()
	if err != nil {
		return nil, err
	}

	pods, err := p.pods.GetPods()
	if err != nil {
		return nil, err
	}

	p.pxPvcs = make([]*kubernetes.PxPvc, 0, len(k8sPvcs))
	for i, _ := range k8sPvcs {
		pxpvc := kubernetes.NewPxPvc(&k8sPvcs[i])
		vExists := pxpvc.SetVolume(vols)
		if vExists == true {
			pxpvc.SetPods(pods)
			p.pxPvcs = append(p.pxPvcs, pxpvc)
		}
	}
	return p.pxPvcs, nil
}
