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
package cloudmigration

import (
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type cloudMigrationCreateOpts struct {
	req *api.SdkCloudMigrateStartRequest

	all      bool
	volumeId string
	groupId  string
}

var (
	ccmOpts                 cloudMigrationCreateOpts
	createCloudmigrationCmd *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	ccmOpts = cloudMigrationCreateOpts{
		req: &api.SdkCloudMigrateStartRequest{},
	}

	createCloudmigrationCmd = &cobra.Command{
		Use:   "cloudmigration",
		Short: "Start a cloud migration",
		Long:  `Migrates set of volumes or a particular volume or a group from cluster pair`,
		Example: `
  # Migrating all volumes from cluster 9548a6e0-053b-43fb-8d26-42b628d6c405:
  pxc create cloudmigration --all -c 9548a6e0-053b-43fb-8d26-42b628d6c405

  # Migrating particular volume from clusterpair
  # Here volid 134771906220836406 is being migrated:
  pxc create cloudmigration --volume-id 134771906220836406`,
		RunE: createCloudmigrationExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	cmd.CreateAddCommand(createCloudmigrationCmd)

	createCloudmigrationCmd.Flags().BoolVarP(&ccmOpts.all, "all", "a", false, "Migrate all volumes")
	createCloudmigrationCmd.Flags().StringVarP(&ccmOpts.volumeId, "volume-id", "v", "", "Volume ID to migrate")
	createCloudmigrationCmd.Flags().StringVarP(&ccmOpts.groupId, "group-id", "g", "", "Group ID to migrate")
	createCloudmigrationCmd.Flags().StringVarP(&ccmOpts.req.ClusterId, "cluster-id", "c", "", "ID of the cluster in which volumes are to be migrated")
	createCloudmigrationCmd.Flags().StringVarP(&ccmOpts.req.TaskId, "task-id", "t", "", "Unique name assocaiated with this migration for idempotency (optional).")
	createCloudmigrationCmd.Flags().SortFlags = false
})

func CreateAddCommand(cmd *cobra.Command) {
	createCloudmigrationCmd.AddCommand(cmd)
}

func createCloudmigrationExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectDefault()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Parse input options
	switch {
	case ccmOpts.all:
		ccmOpts.req.Opt = &api.SdkCloudMigrateStartRequest_AllVolumes{
			AllVolumes: &api.SdkCloudMigrateStartRequest_MigrateAllVolumes{},
		}

	case ccmOpts.groupId != "":
		ccmOpts.req.Opt = &api.SdkCloudMigrateStartRequest_VolumeGroup{
			VolumeGroup: &api.SdkCloudMigrateStartRequest_MigrateVolumeGroup{
				GroupId: ccmOpts.groupId,
			},
		}

	case ccmOpts.volumeId != "":
		ccmOpts.req.Opt = &api.SdkCloudMigrateStartRequest_Volume{
			Volume: &api.SdkCloudMigrateStartRequest_MigrateVolume{
				VolumeId: ccmOpts.volumeId,
			},
		}

	default:
		return util.PxErrorMessage(err, "Must supply a --group-id, --volume-id, or --all")
	}

	// Send request
	migration := api.NewOpenStorageMigrateClient(conn)
	resp, err := migration.Start(ctx, ccmOpts.req)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to start volume migration")
	}

	// Show user information
	util.Printf("Cloud migration started. TaskId: %s\n", resp.Result.TaskId)
	return nil
}
