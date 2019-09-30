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
	"errors"
	"github.com/portworx/pxc/pkg/util"
	"io"
	"sort"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

type PxAlertOps interface {
	GetPxAlerts(alert string, alertId string) (AlertResp, error)
	DeletePxAlerts(alert string) error
}

type pxAlertOps struct{}

type AlertResp struct {
	AlertResp     []*api.Alert
	AlertNameToId map[string]int64
	AlertIdToName map[int64]string
}

type getAlertsOpts struct {
	req *api.SdkAlertsEnumerateWithFiltersRequest
}

type delAlertsOpts struct {
	req *api.SdkAlertsDeleteRequest
}

type alertsList []*api.Alert

func (g alertsList) Len() int           { return len(g) }
func (g alertsList) Less(i, j int) bool { return g[i].Timestamp.Seconds < g[j].Timestamp.Seconds }
func (g alertsList) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }

func NewPxAlertOps() PxAlertOps {
	return &pxAlertOps{}
}

func (p *pxAlertOps) GetPxAlerts(alert string, alertId string) (AlertResp, error) {
	alertResp := AlertResp{}

	ctx, conn, err := PxConnectDefault()
	_ = ctx
	if err != nil {
		return alertResp, err
	}
	defer conn.Close()
	getAlertsGetReq := getAlertsOpts{
		req: &api.SdkAlertsEnumerateWithFiltersRequest{},
	}
	var myAlerts []*api.Alert

	alertResp.AlertNameToId = make(map[string]int64)
	alertResp.AlertIdToName = make(map[int64]string)
	for k, v := range TypeToSpec() {
		id := int64(k)
		name := v.Name
		alertResp.AlertNameToId[name] = id
		alertResp.AlertIdToName[id] = name
	}

	alterType := getAlertType(alert)
	for _, resourceType := range alterType {
		resourceType := resourceType
		if len(alertId) > 0 {
			id, ok := alertResp.AlertNameToId[alertId]
			if !ok {
				return alertResp, nil
			}

			getAlertsGetReq.req.Queries = []*api.SdkAlertsQuery{
				{
					Query: &api.SdkAlertsQuery_AlertTypeQuery{
						AlertTypeQuery: &api.SdkAlertsAlertTypeQuery{
							ResourceType: resourceType,
							AlertType:    id,
						},
					},
				},
			}

		} else {
			getAlertsGetReq.req.Queries = []*api.SdkAlertsQuery{
				{
					Query: &api.SdkAlertsQuery_ResourceTypeQuery{
						ResourceTypeQuery: &api.SdkAlertsResourceTypeQuery{
							ResourceType: resourceType,
						},
					},
				},
			}

		}

		// Send request
		client := api.NewOpenStorageAlertsClient(conn)
		resp, err := client.EnumerateWithFilters(ctx, getAlertsGetReq.req)
		if err != nil {
			util.Eprintf("Failed to fetch alerts ")
			return alertResp, nil
		}

		for {
			res, err := resp.Recv()
			if err == io.EOF {
				break
			}
			myAlerts = append(myAlerts, res.Alerts...)
		}
	}
	sort.Sort(alertsList(myAlerts))
	alertResp.AlertResp = myAlerts
	return alertResp, nil
}

func (p *pxAlertOps) DeletePxAlerts(alert string) error {
	alertResp := AlertResp{}

	ctx, conn, err := PxConnectDefault()
	_ = ctx
	if err != nil {
		return err
	}

	defer conn.Close()
	delAlertsGetReq := delAlertsOpts{
		req: &api.SdkAlertsDeleteRequest{},
	}

	alertResp.AlertNameToId = make(map[string]int64)
	alertResp.AlertIdToName = make(map[int64]string)
	for k, v := range TypeToSpec() {
		id := int64(k)
		name := v.Name
		alertResp.AlertNameToId[name] = id
		alertResp.AlertIdToName[id] = name
	}

	// TODO: For now making it all, will change once PR#69 gets merged
	alterType := getAlertType(alert)
	for _, resourceType := range alterType {
		delAlertsGetReq.req.Queries = []*api.SdkAlertsQuery{
			{
				Query: &api.SdkAlertsQuery_ResourceTypeQuery{
					ResourceTypeQuery: &api.SdkAlertsResourceTypeQuery{
						ResourceType: resourceType,
					},
				},
			},
		}
		// Send request
		client := api.NewOpenStorageAlertsClient(conn)
		_, err = client.Delete(ctx, delAlertsGetReq.req)
		if err != nil {
			return errors.New("Failed to delete alerts")
		}
	}
	return err
}

func getAlertType(alr string) []api.ResourceType {
	var resourceTypes []api.ResourceType

	switch alr {
	case "volume":
		resourceTypes = append(resourceTypes, api.ResourceType_RESOURCE_TYPE_VOLUME)
	case "node":
		resourceTypes = append(resourceTypes, api.ResourceType_RESOURCE_TYPE_NODE)
	case "cluster":
		resourceTypes = append(resourceTypes, api.ResourceType_RESOURCE_TYPE_CLUSTER)
	case "drive":
		resourceTypes = append(resourceTypes, api.ResourceType_RESOURCE_TYPE_DRIVE)
	case "all":
		resourceTypes = append(resourceTypes,
			api.ResourceType_RESOURCE_TYPE_VOLUME,
			api.ResourceType_RESOURCE_TYPE_NODE,
			api.ResourceType_RESOURCE_TYPE_CLUSTER,
			api.ResourceType_RESOURCE_TYPE_DRIVE)
	}

	return resourceTypes
}
