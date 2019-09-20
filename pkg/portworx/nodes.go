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
	"math/big"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"
)

type NodeSpec struct {
	NodeNames []string
}

type Nodes interface {
	Objs
	// Gets the nodes details for the named nodes
	GetNodes() ([]*api.StorageNode, error)
	// Gets a specific node
	GetNode(id string) (*api.StorageNode, error)
	// Get the node that volume is attached on
	GetAttachedOn(v *api.Volume) (*api.StorageNode, error)
	// Gets the state of the volume the node is attached on
	GetAttachedState(v *api.Volume) (string, error)
	// GetReplicationInfo returns the details of the replicas of the specified volume
	GetReplicationInfo(v *api.Volume) (*ReplicationInfo, error)
}

type nodes struct {
	pxops    PxOps
	nodeSpec *NodeSpec
	nodeMap  map[string]*api.StorageNode
	nodes    []*api.StorageNode
}

func NewNodes(pxops PxOps, nodeSpec *NodeSpec) Nodes {
	return &nodes{
		pxops:    pxops,
		nodeSpec: nodeSpec,
		nodeMap:  make(map[string]*api.StorageNode),
	}
}

func NewNodesForVolumes(pxops PxOps, vols []*api.Volume) (Nodes, error) {
	ns := GetNodeSpec(vols)
	nodes := NewNodes(pxops, ns)
	_, err := nodes.GetNodes()
	if err != nil {
		return nil, err
	}
	return nodes, err
}

func NewNodesForPxPvcs(pxops PxOps, pxpvcs []*kubernetes.PxPvc) (Nodes, error) {
	vols := make([]*api.Volume, len(pxpvcs))
	for i, pvc := range pxpvcs {
		vols[i] = pvc.PxVolume
	}
	return NewNodesForVolumes(pxops, vols)
}

func (p *nodes) Reset() {
	p.nodeMap = make(map[string]*api.StorageNode)
	p.nodes = make([]*api.StorageNode, 0)
}

func (p *nodes) GetNode(id string) (*api.StorageNode, error) {
	if len(p.nodes) == 0 {
		return nil, fmt.Errorf("Please call GetNodes before calling GetNode")
	}
	if n, ok := p.nodeMap[id]; ok {
		return n, nil
	}
	return nil, fmt.Errorf("Node not found")
}

func (p *nodes) GetNodes() ([]*api.StorageNode, error) {
	if len(p.nodes) == 0 {
		err := p.getNodes()
		if err != nil {
			return make([]*api.StorageNode, 0), nil
		}
	}
	return p.nodes, nil
}

func (p *nodes) getNodes() error {
	var (
		err   error
		names []string
	)
	if len(p.nodeSpec.NodeNames) == 0 {
		names, err = p.pxops.EnumerateNodes()
		if err != nil {
			return err
		}
	} else {
		names = p.nodeSpec.NodeNames
	}
	p.nodes = make([]*api.StorageNode, 0)
	for _, name := range names {
		node, err := p.pxops.GetNode(name)
		if err != nil {
			return err
		}
		p.nodeMap[name] = node
		p.nodes = append(p.nodes, node)
	}
	return nil
}

func (p *nodes) GetAttachedOn(
	v *api.Volume,
) (*api.StorageNode, error) {
	if len(v.GetAttachedOn()) != 0 {
		node, err := p.GetNode(v.GetAttachedOn())
		if err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, nil
}

func (p *nodes) GetAttachedState(
	v *api.Volume,
) (string, error) {
	n, err := p.GetAttachedOn(v)
	if err != nil {
		return "", err
	}
	return getState(v, n), nil
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

func (p *nodes) GetReplicationInfo(
	v *api.Volume,
) (*ReplicationInfo, error) {
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

// This basically looks at the current runtime state of the volume
// and picks out all of the nodes referenced
func GetNodeSpec(
	vols []*api.Volume,
) *NodeSpec {
	nodeNames := make(map[string]bool)
	for _, v := range vols {
		getAllNodeNamesForVolume(v, nodeNames)
	}
	names := make([]string, 0, len(nodeNames))
	for k, _ := range nodeNames {
		names = append(names, k)
	}
	return &NodeSpec{
		NodeNames: names,
	}
}

func getAllNodeNamesForVolume(v *api.Volume, nodeIds map[string]bool) {
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
}
