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
	"github.com/portworx/px/pkg/util"

	"google.golang.org/grpc"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclikube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type PxConnectionData struct {
	// Context of connection
	Ctx context.Context
	// Connection
	Conn *grpc.ClientConn
}

type KubeConnectionData struct {
	ClientConfig clientcmd.ClientConfig
	ClientSet    *kclikube.Clientset
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
	// GetVolumes returns the array of volume objects
	// filtered by the list of volume names specified
	GetVolumes() ([]*api.SdkVolumeInspectResponse, error)
	// PodsUsingVolume returns the list of pods uing the given volume
	PodsUsingVolume(v *api.Volume) []v1.Pod
	// GetAttachedOn returns the attached state of the specified volume
	GetAttachedState(v *api.Volume) string
	// GetReplicationInfo returns the details of the replicas of the specified volume
	GetReplicationInfo(v *api.Volume) *ReplicationInfo
	// GetStats returns the stats for the specified volume
	GetStats(v *api.Volume) *api.Stats
}

type pxVolumeOps struct {
	pxVolumeOpsInfo *PxVolumeOpsInfo
	// array of volume objects based on the list of volume names specified
	Vols []*api.SdkVolumeInspectResponse
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
		volsInfo, err := volumes.InspectWithFilters(p.pxVolumeOpsInfo.Ctx,
			&api.SdkVolumeInspectWithFiltersRequest{})
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

	if p.pxVolumeOpsInfo.ClientSet == nil {
		return make([]v1.Pod, 0), nil
	}
	podClient := p.pxVolumeOpsInfo.ClientSet.CoreV1().Pods(p.pxVolumeOpsInfo.Namespace)
	podList, err := podClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	p.Pods = podList.Items
	return p.Pods, nil
}

func (p *pxVolumeOps) PodsUsingVolume(v *api.Volume) []v1.Pod {
	pods, err := p.getPods()
	if err != nil {
		util.Eprintf("%v\n",
			util.PxErrorMessagef(err,
				"Failed to get pod information for volume %s",
				v.GetLocator().GetName()))
		return nil
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
	return usedPods
}

func (p *pxVolumeOps) getNode(nodeId string) (*api.StorageNode, error) {
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

func (p *pxVolumeOps) getAttachedOn(v *api.Volume) *api.StorageNode {
	if len(v.GetAttachedOn()) != 0 {
		node, err := p.getNode(v.GetAttachedOn())
		if err != nil {
			util.Eprintf("%v\n",
				util.PxErrorMessagef(err,
					"Failed to get node information where volume %s is attached",
					v.GetLocator().GetName()))
			return nil
		}
		return node
	}
	return nil
}

func (p *pxVolumeOps) GetAttachedState(v *api.Volume) string {
	n := p.getAttachedOn(v)
	return getState(v, n)
}

func (p *pxVolumeOps) GetStats(v *api.Volume) *api.Stats {
	volumes := api.NewOpenStorageVolumeClient(p.pxVolumeOpsInfo.Conn)
	volStats, err := volumes.Stats(
		p.pxVolumeOpsInfo.Ctx,
		&api.SdkVolumeStatsRequest{VolumeId: v.GetId()})
	if err != nil {
		util.Eprintf("%v\n",
			util.PxErrorMessagef(err,
				"Failed to get stats for volume %s",
				v.GetLocator().GetName()))
		return &api.Stats{}
	}
	return volStats.GetStats()
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

func (p *pxVolumeOps) GetReplicationInfo(v *api.Volume) *ReplicationInfo {
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
				n, err := p.getNode(newNodeMid)
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
						n, err := p.getNode(reAddNodes[k])
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
			node, err := p.getNode(id)
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
	return rinfo
}
