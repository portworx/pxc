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
	"fmt"
	"math/big"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/kubernetes"
	"github.com/portworx/px/pkg/util"

	"google.golang.org/grpc"

	v1 "k8s.io/api/core/v1"
)

type PxConnectionData struct {
	// Context of connection
	Ctx context.Context
	// Connection
	Conn *grpc.ClientConn
}

type PxVolumeOpsInfo struct {
	// Px connection data
	PxConnectionData
	// K8S connection data
	KubeConnectionData
	// k8s namespace. Using "" means all namespaces
	Namespace string
	// volumeNames. Using
	VolNames []string
	// node ids
	NodeIds []string
	Labels  map[string]string
}

func (p *PxVolumeOpsInfo) Close() {
	p.Conn.Close()
}

type ReplicationSetInfo struct {
	Id         int
	NodeInfo   []string
	HaIncrease string
	ReAddOn    []string
}

type ReplicationInfo struct {
	Rsi    []*ReplicationSetInfo
	Status string
}

type PxVolumeOps interface {
	// GetPxVolumeOpsInfo returns the PxVolumeOpsInfo
	GetPxVolumeOpsInfo() *PxVolumeOpsInfo
	// GetCOps returns the COps inerface
	GetCOps() COps
	// GetVolumes returns the array of volume objects
	// filtered by the list of volume names specified
	GetVolumes() ([]*api.SdkVolumeInspectResponse, error)
	// PodsUsingVolume returns the list of pods uing the given volume
	PodsUsingVolume(v *api.Volume) ([]v1.Pod, error)
	// GetAttachedOn returns the node where the specified volume is attached
	GetAttachedOn(v *api.Volume) (*api.StorageNode, error)
	// GetAttachedState returns the attached state of the specified volume
	GetAttachedState(v *api.Volume) (string, error)
	// GetReplicationInfo returns the details of the replicas of the specified volume
	GetReplicationInfo(v *api.Volume) (*ReplicationInfo, error)
	// GetStats returns the stats for the specified volume
	GetStats(v *api.Volume) (*api.Stats, error)
	// GetPxPvcs returns the list of PxPvcs
	GetPxPvcs() ([]*kubernetes.PxPvc, error)
	// EnumerateNodes returns list of nodes  ids
	EnumerateNodes() ([]string, error)
	// GetNode returns details of given node
	GetNode(id string) (*api.StorageNode, error)
	// GetAllNodesForVolume will return all nodes that currently has anything to do with the volume
	// nodeNames are filled up and returned
	GetAllNodesForVolume(v *api.Volume, nodeNames map[string]bool) error
}

type pxVolumeOps struct {
	pxVolumeOpsInfo *PxVolumeOpsInfo
	// array of volume objects based on the list of volume names specified
	Vols []*api.SdkVolumeInspectResponse
	// arrays of PxPvcs in the specified namespace
	pxPvcs []*kubernetes.PxPvc
	// List of all pvcs in the spacified namespace
	Pvcs []v1.PersistentVolumeClaim
	// List of all pods in the specified namespace
	Pods []v1.Pod
	// A cache of node ids to StorageNodes. Build as and when we see new node ids
	// Just to ensure we don't repeatedly ask for the same storage node
	NodeMap map[string]*api.StorageNode
}

func NewPxVolumeOps(pxVolOpsInfo *PxVolumeOpsInfo) (PxVolumeOps, error) {
	if pxVolOpsInfo.Ctx == nil || pxVolOpsInfo.Conn == nil {
		return nil, fmt.Errorf("px connection objects cannot be nil")
	}
	return &pxVolumeOps{
		pxVolumeOpsInfo: pxVolOpsInfo,
		NodeMap:         make(map[string]*api.StorageNode),
	}, nil
}

// GetPxVolumeOpsInfo returns the connection data for px
func (p *pxVolumeOps) GetPxVolumeOpsInfo() *PxVolumeOpsInfo {
	return p.pxVolumeOpsInfo
}

func (p *pxVolumeOps) GetCOps() COps {
	return NewCOps(&p.pxVolumeOpsInfo.KubeConnectionData)
}

func (p *pxVolumeOps) GetVolumes() ([]*api.SdkVolumeInspectResponse, error) {
	if p.Vols != nil {
		return p.Vols, nil
	}
	// Get volume information
	volumes := api.NewOpenStorageVolumeClient(p.pxVolumeOpsInfo.Conn)

	// Determine if we should get all the volumes or specific ones
	if len(p.pxVolumeOpsInfo.VolNames) != 0 {
		p.Vols = make([]*api.SdkVolumeInspectResponse, 0, len(p.pxVolumeOpsInfo.VolNames))
		for _, v := range p.pxVolumeOpsInfo.VolNames {
			vol, err := volumes.Inspect(p.pxVolumeOpsInfo.Ctx,
				&api.SdkVolumeInspectRequest{VolumeId: v})
			if err != nil {
				return nil, util.PxErrorMessagef(err, "Failed to get volume %s", v)
			}
			p.Vols = append(p.Vols, vol)
		}
	} else {
		// If it is no volumes (all)
		var (
			err      error
			volsInfo *api.SdkVolumeInspectWithFiltersResponse
		)

		// if label is set, use them as filter
		volsInfo, err = volumes.InspectWithFilters(p.pxVolumeOpsInfo.Ctx,
			&api.SdkVolumeInspectWithFiltersRequest{
				Labels: p.GetPxVolumeOpsInfo().Labels,
			})

		if err != nil {
			return nil, util.PxErrorMessage(err, "Failed to get volumes")
		}
		p.Vols = volsInfo.GetVolumes()
	}
	return p.Vols, nil
}

func (p *pxVolumeOps) getPods() ([]v1.Pod, error) {
	if p.Pods != nil {
		return p.Pods, nil
	}

	co := p.GetCOps()

	pods, err := co.GetPodsByLabels(p.pxVolumeOpsInfo.Namespace,
		util.StringMapToCommaString(p.pxVolumeOpsInfo.Labels))
	if err != nil {
		return nil, err
	}
	p.Pods = pods
	return p.Pods, nil
}

func (p *pxVolumeOps) PodsUsingVolume(v *api.Volume) ([]v1.Pod, error) {
	pods, err := p.getPods()
	if err != nil {
		return nil, err
	}
	usedPods := make([]v1.Pod, 0)
	namespace := v.Locator.VolumeLabels["namespace"]
	pvc := v.Locator.VolumeLabels["pvc"]
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

func (p *pxVolumeOps) EnumerateNodes() ([]string, error) {
	if len(p.pxVolumeOpsInfo.NodeIds) != 0 {
		return p.pxVolumeOpsInfo.NodeIds, nil
	}

	nodes := api.NewOpenStorageNodeClient(p.pxVolumeOpsInfo.Conn)
	nodesInfo, err := nodes.Enumerate(p.pxVolumeOpsInfo.Ctx, &api.SdkNodeEnumerateRequest{})
	if err != nil {
		return make([]string, 0), err
	}
	return nodesInfo.GetNodeIds(), nil
}

func (p *pxVolumeOps) GetNode(nodeId string) (*api.StorageNode, error) {
	// Check if we have already queried for the node. If so just return from cachew
	if n, ok := p.NodeMap[nodeId]; ok {
		return n, nil
	}

	// Get the node details
	nodes := api.NewOpenStorageNodeClient(p.pxVolumeOpsInfo.Conn)
	nodeInfo, err := nodes.Inspect(
		p.pxVolumeOpsInfo.Ctx,
		&api.SdkNodeInspectRequest{NodeId: nodeId})
	if err != nil {
		return nil, err
	}

	// Cache the info
	n := nodeInfo.GetNode()
	p.NodeMap[nodeId] = n
	return n, nil
}

func (p *pxVolumeOps) GetAttachedOn(v *api.Volume) (*api.StorageNode, error) {
	if len(v.GetAttachedOn()) != 0 {
		node, err := p.GetNode(v.GetAttachedOn())
		if err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, nil
}

func (p *pxVolumeOps) GetAttachedState(v *api.Volume) (string, error) {
	n, err := p.GetAttachedOn(v)
	if err != nil {
		return "", err
	}
	return getState(v, n), nil
}

func (p *pxVolumeOps) GetStats(v *api.Volume) (*api.Stats, error) {
	volumes := api.NewOpenStorageVolumeClient(p.pxVolumeOpsInfo.Conn)
	volStats, err := volumes.Stats(
		p.pxVolumeOpsInfo.Ctx,
		&api.SdkVolumeStatsRequest{VolumeId: v.GetId()})
	if err != nil {
		return &api.Stats{}, err
	}
	return volStats.GetStats(), nil
}

const (
	PXReplCurrSetMid         = "ReplicaSetCurrMid"
	PXReplSetCreateMid       = "ReplicaSetCreateMid"
	PXReplNewNodeMid         = "ReplNewNodeMid"
	PXReplReAddPools         = "PXReplReAddPools"
	PXReplReAddNodeMid       = "PXReplReAddNodeMid"
	PXReplReAddUsedSize      = "PXReplReAddUsedSize"
	PXReplNodePools          = "ReplNodePools"
	PXReplNewNodePools       = "ReplNewNodePools"
	PXReplRemoveMids         = "ReplRemoveMids"
	PXReplRuntimeState       = "RuntimeState"
	RuntimeStateResync       = "resync"
	RuntimeStateResyncFailed = "resync_failed"
)

func (p *pxVolumeOps) GetReplicationInfo(v *api.Volume) (*ReplicationInfo, error) {
	numResync := 0
	nodesDown := false
	numResyncFailed := 0
	numReAdd := 0

	rinfo := &ReplicationInfo{
		Rsi: make([]*ReplicationSetInfo, 0),
	}
	runtimeStates := v.GetRuntimeState()

	// Parse each of the replica sets
	for i, rset := range v.GetReplicaSets() {
		var (
			currMidStr  string
			poolIds     []string
			removeMids  string
			createNodes []string
			ok          bool
		)

		replicationSetInfo := &ReplicationSetInfo{
			Id:       i,
			NodeInfo: make([]string, 0),
			ReAddOn:  make([]string, 0),
		}

		if i < len(runtimeStates) && runtimeStates[i].GetRuntimeState() != nil {
			irs := runtimeStates[i].GetRuntimeState()

			createNodes = strings.Split(irs[PXReplSetCreateMid], ",")
			// get all of the ids of nodes that were removed
			removeMids = irs[PXReplRemoveMids]

			// Get all of the pool ids for this replica set
			if poolIdsStr, ok := irs[PXReplNodePools]; ok {
				poolIds = strings.Split(poolIdsStr, ",")
			}

			// If current set of nodes don't match the set of nodes in the set
			// this means that one or more nodes are down
			if currMidStr, ok = irs[PXReplCurrSetMid]; ok {
				if len(strings.Split(currMidStr, ",")) != len(rset.Nodes) {
					nodesDown = true
				}
			}

			// If it is attached check if state is any of the nodes are in
			// resync or resync failed states
			if len(v.GetAttachedOn()) > 0 {
				if rstate, ok := irs[PXReplRuntimeState]; ok {
					if rstate == RuntimeStateResync {
						numResync++
					} else if rstate == RuntimeStateResyncFailed {
						numResyncFailed++
					}
				}
			}

			// Check if a HA Increase is in progress and if so get the details
			if newNodeMid, ok := irs[PXReplNewNodeMid]; ok {
				newNodePool := ""
				if newNodePool, ok =
					irs[PXReplNewNodePools]; ok {
					newNodePool = " (Pool " + newNodePool + ")"
				}
				if usedSize, ok := irs[PXReplReAddUsedSize]; ok {
					if bytesUsed, err := strconv.ParseUint(usedSize, 10, 64); err == nil {
						newNodePool += " (" + humanize.BigIBytes(big.NewInt(int64(bytesUsed))) + " transferred)"
					}
				}
				var hn string
				n, err := p.GetNode(newNodeMid)
				if err != nil {
					util.Eprintf("%v\n",
						util.PxErrorMessagef(err,
							"Failed to get node (%s) information for replica of volume %s ",
							newNodeMid, v.GetLocator().GetName()))
					// Use id instead
					hn = newNodeMid
				} else {
					hn = n.GetHostname()
				}
				replicationSetInfo.HaIncrease = fmt.Sprintf("%s%s", hn, newNodePool)
			}

			// Check if a readd of a node is in progress and if os get the details
			if reAddMids, ok := irs[PXReplReAddNodeMid]; ok {
				reAddPools := strings.Split(irs[PXReplReAddPools], ",")
				reAddNodes := strings.Split(reAddMids, ",")
				numReAdd = len(reAddMids)
				if numReAdd > 0 {
					for k := 0; k < len(reAddNodes); k++ {
						var hn string
						n, err := p.GetNode(reAddNodes[k])
						if err != nil {
							util.Eprintf("%v\n",
								util.PxErrorMessagef(err,
									"Failed to get node (%s) information for replica of volume %s ",
									reAddNodes[k], v.GetLocator().GetName()))
							// Use id instead
							hn = reAddNodes[k]
						} else {
							hn = n.GetHostname()
						}
						replicationSetInfo.ReAddOn = append(replicationSetInfo.ReAddOn,
							hn+" (Pool "+reAddPools[k]+")")
					}
				}
			}
		}

		// If for any reason the RuntimeState did not have list of nodes use the nodes in the rset
		if len(createNodes) == 0 {
			createNodes = rset.Nodes
		}

		// Get the list of nodes in the replica set and their states
		for j, id := range createNodes {
			notInCurrSet := ""
			removed := ""
			if !strings.Contains(currMidStr, id) {
				notInCurrSet = "* "
			}

			if strings.Contains(removeMids, id) {
				removed = "(removal in-progess)"
			}
			poolId := ""
			if j < len(poolIds) {
				poolId = " (Pool " + poolIds[j] + ")"
			}

			var hn string
			node, err := p.GetNode(id)
			if err != nil {
				util.Eprintf("%v\n",
					util.PxErrorMessagef(err,
						"Failed to get node (%s) information for replica of volume %s ",
						id, v.GetLocator().GetName()))
				// Use id instead
				hn = id
			} else {
				hn = node.GetHostname()
			}
			str := fmt.Sprintf("%s%s%s%s", hn, poolId, notInCurrSet, removed)
			replicationSetInfo.NodeInfo = append(replicationSetInfo.NodeInfo, str)
		}
		rinfo.Rsi = append(rinfo.Rsi, replicationSetInfo)
	}

	// Figure out the status of replication
	replStatus := "UP"
	if v.State == api.VolumeState_VOLUME_STATE_RESTORE {
		replStatus = "Restore"
	} else if len(v.GetAttachedOn()) == 0 {
		replStatus = "Detached"
	} else if numResyncFailed != 0 {
		replStatus = "Not in quorum"
	} else if nodesDown && numResync == 0 || numReAdd > 0 {
		replStatus = "Degraded"
	} else if numResync > 0 {
		replStatus = "Resync"
	}
	rinfo.Status = replStatus
	return rinfo, nil
}

func (p *pxVolumeOps) getPvcs() ([]v1.PersistentVolumeClaim, error) {
	if p.Pvcs != nil {
		return p.Pvcs, nil
	}

	co := p.GetCOps()

	pvcs, err := co.GetPvcsByLabels(p.pxVolumeOpsInfo.Namespace,
		util.StringMapToCommaString(p.pxVolumeOpsInfo.Labels))
	if err != nil {
		return nil, err
	}

	p.Pvcs = pvcs
	return p.Pvcs, nil
}

func (p *pxVolumeOps) GetPxPvcs() ([]*kubernetes.PxPvc, error) {
	if p.pxPvcs != nil {
		return p.pxPvcs, nil
	}

	k8sPvcs, err := p.getPvcs()
	if err != nil {
		return nil, err
	}

	vols, err := p.GetVolumes()
	if err != nil {
		return nil, err
	}

	pods, err := p.getPods()
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

func addToNodeIds(nodeIds map[string]bool, nodes []string) {
	for _, n := range nodes {
		nodeIds[n] = true
	}
}

func getMidsArray(irs map[string]string, key string, nodeIds map[string]bool) {
	if v, ok := irs[key]; ok {
		if len(v) > 0 {
			nn := strings.Split(v, ",")
			addToNodeIds(nodeIds, nn)
		}
	}
}

// This basically looks at the current runtime state of the volume and picks out all of the nodes referenced
func (p *pxVolumeOps) GetAllNodesForVolume(v *api.Volume, nodeNames map[string]bool) error {
	nodeIds := make(map[string]bool)
	runtimeStates := v.GetRuntimeState()
	for i, rset := range v.GetReplicaSets() {
		if i < len(runtimeStates) && runtimeStates[i].GetRuntimeState() != nil {
			irs := runtimeStates[i].GetRuntimeState()
			getMidsArray(irs, PXReplSetCreateMid, nodeIds)
			getMidsArray(irs, PXReplRemoveMids, nodeIds)
			getMidsArray(irs, PXReplCurrSetMid, nodeIds)
			getMidsArray(irs, PXReplReAddNodeMid, nodeIds)
			if len(v.GetAttachedOn()) > 0 {
				nodeIds[v.GetAttachedOn()] = true
			}
			if newNodeMid, ok := irs[PXReplNewNodeMid]; ok {
				if len(newNodeMid) > 0 {
					nodeIds[newNodeMid] = true
				}
			}
		}
		addToNodeIds(nodeIds, rset.Nodes)
	}
	for k, _ := range nodeIds {
		n, err := p.GetNode(k)
		if err != nil {
			return err
		}
		nodeNames[n.GetHostname()] = true
	}
	return nil
}
