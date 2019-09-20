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

var typeToSpec = map[AlertType]AlertSpec{

	DriveOperationFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_DRIVE,
		"Drive operation failure",
		"DriveOperationFailure",
		false,
	},

	DriveOperationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_DRIVE,
		"Drive operation success",
		"DriveOperationSuccess",
		false,
	},

	DriveStateChange: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_DRIVE,
		"Drive state change",
		"DriveStateChange",
		false,
	},

	VolumeOperationFailureAlarm: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume operation failure",
		"VolumeOperationFailureAlarm",
		true,
	},

	VolumeOperationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume operation success",
		"VolumeOperationSuccess",
		false,
	},

	VolumeStateChange: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume state change",
		"VolumeStateChange",
		true,
	},

	VolGroupOperationFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Volume group operation failure",
		"VolGroupOperationFailure",
		true,
	},

	VolGroupOperationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Volume group operation failure",
		"VolGroupOperationSuccess",
		true,
	},

	VolGroupStateChange: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Volume group state change",
		"VolGroupStateChange",
		true,
	},

	NodeStartFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node start failure",
		"NodeStartFailure",
		false,
	},

	NodeStartSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node start success",
		"NodeStartSuccess",
		false,
	},

	NodeStateChange: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node state change",
		"NodeStateChange",
		false,
	},

	NodeJournalHighUsage: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node journal high usage",
		"NodeJournalHighUsage",
		false,
	},

	IOOperation: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Io operation",
		"IOOperation",
		false,
	},

	ContainerOperationFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Container operation failure",
		"ContainerOperationFailure",
		true,
	},

	ContainerOperationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Container operation succes",
		"ContainerOperationSuccess",
		true,
	},

	ContainerStateChange: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Container state change",
		"ContainerStateChange",
		true,
	},

	PXInitFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Px init failure",
		"PXInitFailure",
		false,
	},

	PXInitSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Px init success",
		"PXInitSuccess",
		false,
	},

	PXStateChange: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Px state change",
		"PXStateChange",
		false,
	},

	VolumeOperationFailureWarn: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume operation failure",
		"VolumeOperationFailureWarn",
		true,
	},

	ClusterManagerFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Cluster manager failure",
		"ClusterManagerFailure",
		false,
	},

	KernelDriverFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Kernel driver error",
		"KernelDriverFailure",
		false,
	},

	NodeDecommissionSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node decommission success",
		"NodeDecommissionSuccess",
		false,
	},

	NodeDecommissionFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node decommission failure",
		"NodeDecommissionFailure",
		false,
	},

	NodeDecommissionPending: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node decommission pending",
		"NodeDecommissionPending",
		false,
	},

	NodeInitFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node init failure",
		"NodeInitFailure",
		false,
	},

	PXAlertMax: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Px Alert Max",
		"PXAlertMax",
		false,
	},

	NodeScanCompletion: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Node media scan completion",
		"NodeScanCompletion",
		false,
	},

	VolumeSpaceLow: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume space low",
		"VolumeSpaceLow",
		false,
	},

	VolumeSpaceLowCleared: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume space no longer low",
		"VolumeSpaceLowCleared",
		false,
	},

	ReplAddVersionMismatch: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume HA increase operation",
		"ReplAddVersionMismatch",
		false,
	},

	CloudsnapScheduleFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Cloudsnap schedule configuration failure",
		"CloudsnapScheduleFailure",
		false,
	},

	CloudsnapOperationUpdate: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Cloudsnap operation update",
		"CloudsnapOperationUpdate",
		true,
	},

	CloudsnapOperationFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Cloudsnap operation failure",
		"CloudsnapOperationFailure",
		true,
	},

	CloudsnapOperationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Cloudsnap opertion success",
		"CloudsnapOperationSuccess",
		true,
	},

	NodeMarkedDown: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Node marked down",
		"NodeMarkedDown",
		false,
	},

	VolumeCreateSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume create success",
		"VolumeCreateSuccess",
		true,
	},

	VolumeCreateFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume create failure",
		"VolumeCreateFailure",
		true,
	},

	VolumeDeleteSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume delete success",
		"VolumeDeleteSuccess",
		true,
	},

	VolumeDeleteFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume delete failure",
		"VolumeDeleteFailure",
		true,
	},

	VolumeMountSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume mount success",
		"VolumeMountSuccess",
		true,
	},

	VolumeMountFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume mount failure",
		"VolumeMountFailure",
		true,
	},

	VolumeUnmountSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume unmount success",
		"VolumeUnmountSuccess",
		true,
	},

	VolumeUnmountFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume unmount failure",
		"VolumeUnmountFailure",
		true,
	},

	VolumeHAUpdateSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume ha update success",
		"VolumeHAUpdateSuccess",
		true,
	},

	VolumeHAUpdateFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Volume ha update failure",
		"VolumeHAUpdateFailure",
		true,
	},

	SnapshotCreateSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot create success",
		"SnapshotCreateSuccess",
		true,
	},

	SnapshotCreateFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot create failure",
		"SnapshotCreateFailure",
		true,
	},

	SnapshotRestoreSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot restore success",
		"SnapshotRestoreSuccess",
		true,
	},

	SnapshotRestoreFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot restore failure",
		"SnapshotRestoreFailure",
		true,
	},

	SnapshotIntervalUpdateFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot interval update failure",
		"SnapshotIntervalUpdateFailure",
		true,
	},

	SnapshotIntervalUpdateSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot interval update success",
		"SnapshotIntervalUpdateSuccess",
		true,
	},

	PXReady: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"PX ready",
		"PXReady",
		false,
	},

	StorageFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Storage could not be mounted",
		"StorageFailure",
		true,
	},

	ObjectstoreStateChange: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Objectstore operation update",
		"ObjectstoreStateChange",
		false,
	},

	ObjectstoreFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Objectstore failure",
		"ObjectstoreFailure",
		false,
	},

	ObjectstoreSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Objectstore success",
		"ObjectstoreSuccess",
		false,
	},

	LicenseExpiring: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"License expiring",
		"LicenseExpiring",
		true,
	},

	VolumeExtentDiffSlow: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Extent diff is slow",
		"VolumeExtentDiffSlow",
		true,
	},

	VolumeExtentDiffOk: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Extent diff is ok",
		"VolumeExtentDiffOk",
		true,
	},

	SharedV4SetupFailure: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_NODE,
		"Sharedv4 service setup failure",
		"SharedV4SetupFailure",
		true,
	},

	SnapshotDeleteSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot delete success",
		"SnapshotDeleteSuccess",
		true,
	},

	SnapshotDeleteFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Snapshot delete failure",
		"SnapshotDeleteFailure",
		true,
	},

	DriveStateChangeClear: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_DRIVE,
		"Drive state change clear",
		"DriveStateChangeClear",
		false,
	},

	ClusterPairSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Cluster Pair created successfully",
		"ClusterPairSuccess",
		false,
	},
	ClusterPairFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Failed to create Cluster Pair",
		"ClusterPairFailure",
		false,
	},
	CloudMigrationUpdate: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"CloudMigration operation update",
		"CloudMigrationUpdate",
		false,
	},
	CloudMigrationSuccess: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"CloudMigration operation success",
		"CloudMigrationSuccess",
		false,
	},
	CloudMigrationFailure: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"CloudMigration operation failure",
		"CloudMigrationFailure",
		false,
	},
	ClusterDomainAdded: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Cluster domain added",
		"ClusterDomainAdded",
		false,
	},

	ClusterDomainRemoved: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Cluster domain removed",
		"ClusterDomainRemoved",
		false,
	},

	ClusterDomainActivated: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Cluster domain activated",
		"ClusterDomainActivated",
		false,
	},

	ClusterDomainDeactivated: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"Cluster domain deactivated",
		"ClusterDomainDeactivated",
		false,
	},
	MeteringAgentWarning: {
		api.SeverityType_SEVERITY_TYPE_WARNING,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"MeteringAgent operation warning",
		"MeteringAgentWarning",
		false,
	},
	MeteringAgentCritical: {
		api.SeverityType_SEVERITY_TYPE_ALARM,
		api.ResourceType_RESOURCE_TYPE_CLUSTER,
		"MeteringAgent operation critical",
		"MeteringAgentCritical",
		false,
	},
	PXMaxAlertNum: {
		api.SeverityType_SEVERITY_TYPE_NOTIFY,
		api.ResourceType_RESOURCE_TYPE_VOLUME,
		"Px max alert num",
		"PXMaxAlertNum",
		false,
	},
}
