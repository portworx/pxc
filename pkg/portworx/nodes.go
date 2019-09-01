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
	"github.com/portworx/pxc/pkg/util"
)

type Nodes interface {
	Objs
	// EnumerateNodes lists all nodes in the cluster
	EnumerateNodes() ([]string, error)
	// Gets the nodes details for the named nodes
	GetNodes(names []string) ([]*api.StorageNode, error)
	// Gets a specific node
	GetNode(id string) (*api.StorageNode, error)
	// Get the node that volume is attached on
	GetAttachedOn(v *api.Volume) (*api.StorageNode, error)
	// Gets the state of the volume the node is attached on
	GetAttachedState(v *api.Volume) (string, error)
	// GetReplicationInfo returns the details of the replicas of the specified volume
	GetReplicationInfo(v *api.Volume) (*ReplicationInfo, error)
	// GetAllNodesForVolume will return all nodes that currently has anything to do with the volume
	// nodeNames are filled up and returned
	GetAllNodesForVolume(v *api.Volume, nodeNames map[string]bool) error
}

type nodes struct {
	pxops   PxOps
	nodeMap map[string]*api.StorageNode
}

func NewNodes(pxops PxOps) Nodes {
	return &nodes{
		pxops:   pxops,
		nodeMap: make(map[string]*api.StorageNode),
	}
}

func (p *nodes) Reset() {
	p.nodeMap = make(map[string]*api.StorageNode)
}

func (p *nodes) GetNode(id string) (*api.StorageNode, error) {
	if n, ok := p.nodeMap[id]; ok {
		return n, nil
	}
	n, err := p.pxops.GetNode(id)
	if err != nil {
		return nil, err
	}
	p.nodeMap[id] = n
	return n, nil
}

func (p *nodes) EnumerateNodes() ([]string, error) {
	return p.pxops.EnumerateNodes()
}

func (p *nodes) GetNodes(names []string) ([]*api.StorageNode, error) {
	var err error
	if len(names) == 0 {
		names, err = p.EnumerateNodes()
		if err != nil {
			return nil, err
		}
	}
	m := make([]*api.StorageNode, 0, len(names))
	for _, n := range names {
		node, err := p.GetNode(n)
		if err != nil {
			return nil, err
		}
		m = append(m, node)
	}
	return m, nil
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
func (p *nodes) GetAllNodesForVolume(
	v *api.Volume,
	nodeNames map[string]bool,
) error {
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
