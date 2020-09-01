/*
Copyright © 2020 Portworx

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
	"github.com/portworx/pxc/pkg/util"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
)

// RoleGuestDisabled indicates that the system.guest is disabled
var RoleGuestDisabled = api.SdkRole{
	Name: "system.guest",
	Rules: []*api.SdkRule{
		{
			Services: []string{"!*"},
			Apis:     []string{"!*"},
		},
	},
}

// RoleGuestEnabled indicates that the system.guest is enabled
var RoleGuestEnabled = api.SdkRole{
	Name: "system.guest",
	Rules: []*api.SdkRule{
		{
			Services: []string{"mountattach", "volume", "cloudbackup", "migrate"},
			Apis:     []string{"*"},
		},
		{
			Services: []string{"identity"},
			Apis:     []string{"version"},
		},
		{
			Services: []string{
				"cluster",
				"node",
			},
			Apis: []string{
				"inspect*",
				"enumerate*",
			},
		},
	},
}

// AuthOps represents all auth related commands
type AuthOps interface {
	UpdateRole(r *api.SdkRole) error
	GetRole(name string) (*api.SdkRole, error)
}

// CliAuthInputs represents input for auth commands
type CliAuthInputs struct {
	util.BaseFormatOutput
	Wide bool
}

type authOps struct{}

// NewAuthOps creates a new auth ops
func NewAuthOps() AuthOps {
	return &authOps{}
}

func (p *authOps) GetRole(name string) (*api.SdkRole, error) {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	roles := api.NewOpenStorageRoleClient(conn)
	resp, err := roles.Inspect(ctx, &api.SdkRoleInspectRequest{
		Name: name,
	})
	if err != nil {
		return nil, util.PxErrorMessage(err, "Failed to get role")
	}

	return resp.GetRole(), nil
}

func (p *authOps) UpdateRole(r *api.SdkRole) error {
	ctx, conn, err := PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	roles := api.NewOpenStorageRoleClient(conn)
	_, err = roles.Update(ctx, &api.SdkRoleUpdateRequest{
		Role: r,
	})
	if err != nil {
		return util.PxErrorMessage(err, "Failed to update role")
	}

	return nil
}
