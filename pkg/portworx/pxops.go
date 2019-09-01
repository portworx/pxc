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
	"context"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/util"

	"google.golang.org/grpc"
)

type VolumeSpec struct {
	VolNames []string
	Labels   map[string]string
}

type PxOps interface {
	// Close the connection to Portworx
	Close()
	// GetVolumes returns the array of volume objects
	// filtered by the specification  in volSpec
	GetVolumes(volSpec *VolumeSpec) ([]*api.SdkVolumeInspectResponse, error)
	// GetStats returns the stats for the specified volume
	GetStats(v *api.Volume, notCumulative bool) (*api.Stats, error)
	// EnumerateNodes returns list of nodes  ids
	EnumerateNodes() ([]string, error)
	// GetNode returns details of given node
	GetNode(id string) (*api.StorageNode, error)
	// GetCtx returns the context
	GetCtx() context.Context
	// GetConn returns the grpc client connection
	GetConn() *grpc.ClientConn
}

type pxOps struct {
	// Context of connection
	ctx context.Context
	// Connection
	conn *grpc.ClientConn
}

func NewPxOps() (PxOps, error) {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return nil, err
	}
	return &pxOps{
		ctx:  ctx,
		conn: conn,
	}, nil
}

func (p *pxOps) Close() {
	p.conn.Close()
}

func (p *pxOps) GetCtx() context.Context {
	return p.ctx
}

func (p *pxOps) GetConn() *grpc.ClientConn {
	return p.conn
}

func (p *pxOps) GetVolumes(volSpec *VolumeSpec) ([]*api.SdkVolumeInspectResponse, error) {
	// Determine if we should get all the volumes or specific ones
	if len(volSpec.VolNames) != 0 {
		return p.getVolumesByName(volSpec.VolNames)
	} else {
		return p.getVolumesByLabel(volSpec.Labels)
	}
}

func (p *pxOps) getVolumesByLabel(labels map[string]string) ([]*api.SdkVolumeInspectResponse, error) {
	volumes := api.NewOpenStorageVolumeClient(p.conn)
	volsInfo, err := volumes.InspectWithFilters(p.ctx,
		&api.SdkVolumeInspectWithFiltersRequest{
			Labels: labels,
		})

	if err != nil {
		return nil, util.PxErrorMessage(err, "Failed to get volumes")
	}
	return volsInfo.GetVolumes(), nil
}

func (p *pxOps) getVolumesByName(names []string) ([]*api.SdkVolumeInspectResponse, error) {
	volumes := api.NewOpenStorageVolumeClient(p.conn)
	vols := make([]*api.SdkVolumeInspectResponse, len(names))
	for i, v := range names {
		vol, err := volumes.Inspect(p.ctx,
			&api.SdkVolumeInspectRequest{VolumeId: v})
		if err != nil {
			return nil, util.PxErrorMessagef(err, "Failed to get volume %s", v)
		}
		vols[i] = vol
	}
	return vols, nil
}

func (p *pxOps) GetStats(v *api.Volume, notCumulative bool) (*api.Stats, error) {
	volumes := api.NewOpenStorageVolumeClient(p.conn)
	volStats, err := volumes.Stats(p.ctx,
		&api.SdkVolumeStatsRequest{
			VolumeId:      v.GetId(),
			NotCumulative: notCumulative,
		})
	if err != nil {
		return &api.Stats{}, err
	}
	return volStats.GetStats(), nil
}

func (p *pxOps) EnumerateNodes() ([]string, error) {
	nodes := api.NewOpenStorageNodeClient(p.conn)
	nodesInfo, err := nodes.Enumerate(p.ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		return make([]string, 0), err
	}
	return nodesInfo.GetNodeIds(), nil
}

func (p *pxOps) GetNode(nodeId string) (*api.StorageNode, error) {
	nodes := api.NewOpenStorageNodeClient(p.conn)
	nodeInfo, err := nodes.Inspect(p.ctx,
		&api.SdkNodeInspectRequest{NodeId: nodeId})
	if err != nil {
		return nil, err
	}

	n := nodeInfo.GetNode()
	return n, nil
}
