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
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

type VolumeSpec struct {
	VolNames []string
	Labels   map[string]string
	Owner    string
}

type Volumes interface {
	Objs
	// GetVolumes returns the array of volume objects
	// filtered by the list of volume names specified
	GetVolumes() ([]*api.Volume, error)
	// GetStats returns the stats for the specified volume
	GetStats(v *api.Volume, notCumulative bool) (*api.Stats, error)
}

type volumes struct {
	// pxops object
	pxops PxOps
	//volspec specifies the volume specification to use
	volSpec *VolumeSpec
	// volume details
	vols []*api.Volume
}

func NewVolumes(pxops PxOps, volSpec *VolumeSpec) Volumes {
	return &volumes{
		pxops:   pxops,
		volSpec: volSpec,
	}
}

func (p *volumes) Reset() {
	p.vols = make([]*api.Volume, 0)
}

func (p *volumes) GetVolumes() ([]*api.Volume, error) {
	if len(p.vols) > 0 {
		return p.vols, nil
	}
	var (
		err  error
		vols []*api.Volume
	)
	if len(p.volSpec.VolNames) == 0 {
		vols, err = p.getVolsBySpec(p.volSpec)
	} else {
		vols, err = p.getVolsByName(p.volSpec.VolNames)
	}
	if err != nil {
		return nil, err
	}
	p.vols = vols
	return p.vols, nil
}

func (p *volumes) getVolsByName(names []string) ([]*api.Volume, error) {
	vols := make([]*api.Volume, len(names))
	for i, v := range names {
		vol, err := p.pxops.GetVolumeById(v)
		if err != nil {
			return nil, err
		}
		vols[i] = vol.GetVolume()
	}
	return vols, nil
}

func (p *volumes) getVolsBySpec(vs *VolumeSpec) ([]*api.Volume, error) {
	v, err := p.pxops.GetVolumesBySpec(vs)
	if err != nil {
		return nil, err
	}
	vols := make([]*api.Volume, len(v))
	for i, vol := range v {
		vols[i] = vol.GetVolume()
	}
	return vols, nil
}

func (p *volumes) GetStats(
	v *api.Volume,
	notCumulative bool,
) (*api.Stats, error) {
	return p.pxops.GetStats(v, notCumulative)
}
