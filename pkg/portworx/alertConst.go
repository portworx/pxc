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

const (
	DriveOperationFailure         AlertType = 0
	DriveOperationSuccess         AlertType = 1
	DriveStateChange              AlertType = 2
	VolumeOperationFailureAlarm   AlertType = 3
	VolumeOperationSuccess        AlertType = 4
	VolumeStateChange             AlertType = 5
	VolGroupOperationFailure      AlertType = 6
	VolGroupOperationSuccess      AlertType = 7
	VolGroupStateChange           AlertType = 8
	NodeStartFailure              AlertType = 9
	NodeStartSuccess              AlertType = 10
	NodeStateChange               AlertType = 11
	NodeJournalHighUsage          AlertType = 12
	IOOperation                   AlertType = 13
	ContainerOperationFailure     AlertType = 14
	ContainerOperationSuccess     AlertType = 15
	ContainerStateChange          AlertType = 16
	PXInitFailure                 AlertType = 17
	PXInitSuccess                 AlertType = 18
	PXStateChange                 AlertType = 19
	VolumeOperationFailureWarn    AlertType = 20
	StorageVolumeMountDegraded    AlertType = 21
	ClusterManagerFailure         AlertType = 22
	KernelDriverFailure           AlertType = 23
	NodeDecommissionSuccess       AlertType = 24
	NodeDecommissionFailure       AlertType = 25
	NodeDecommissionPending       AlertType = 26
	NodeInitFailure               AlertType = 27
	PXAlertMax                    AlertType = 28
	NodeScanCompletion            AlertType = 29
	VolumeSpaceLow                AlertType = 30
	ReplAddVersionMismatch        AlertType = 31
	CloudsnapScheduleFailure      AlertType = 32
	CloudsnapOperationUpdate      AlertType = 33
	CloudsnapOperationFailure     AlertType = 34
	CloudsnapOperationSuccess     AlertType = 35
	NodeMarkedDown                AlertType = 36
	VolumeCreateSuccess           AlertType = 37
	VolumeCreateFailure           AlertType = 38
	VolumeDeleteSuccess           AlertType = 39
	VolumeDeleteFailure           AlertType = 40
	VolumeMountSuccess            AlertType = 41
	VolumeMountFailure            AlertType = 42
	VolumeUnmountSuccess          AlertType = 43
	VolumeUnmountFailure          AlertType = 44
	VolumeHAUpdateSuccess         AlertType = 45
	VolumeHAUpdateFailure         AlertType = 46
	SnapshotCreateSuccess         AlertType = 47
	SnapshotCreateFailure         AlertType = 48
	SnapshotRestoreSuccess        AlertType = 49
	SnapshotRestoreFailure        AlertType = 50
	SnapshotIntervalUpdateFailure AlertType = 51
	SnapshotIntervalUpdateSuccess AlertType = 52
	PXReady                       AlertType = 53
	StorageFailure                AlertType = 54
	ObjectstoreFailure            AlertType = 55
	ObjectstoreSuccess            AlertType = 56
	ObjectstoreStateChange        AlertType = 57
	LicenseExpiring               AlertType = 58
	VolumeExtentDiffSlow          AlertType = 59
	VolumeExtentDiffOk            AlertType = 60
	SharedV4SetupFailure          AlertType = 61
	SnapshotDeleteSuccess         AlertType = 62
	SnapshotDeleteFailure         AlertType = 63
	DriveStateChangeClear         AlertType = 64
	VolumeSpaceLowCleared         AlertType = 65
	ClusterPairSuccess            AlertType = 66
	ClusterPairFailure            AlertType = 67
	CloudMigrationUpdate          AlertType = 68
	CloudMigrationSuccess         AlertType = 69
	CloudMigrationFailure         AlertType = 70
	ClusterDomainAdded            AlertType = 71
	ClusterDomainRemoved          AlertType = 72
	ClusterDomainActivated        AlertType = 73
	ClusterDomainDeactivated      AlertType = 74
	MeteringAgentWarning          AlertType = 75
	MeteringAgentCritical         AlertType = 76
	// **********************************
	// Add new alerts before the MaxAlert
	// **********************************
	PXMaxAlertNum AlertType = 77 // --- ADD NEW ALERT ABOVE THIS LINE --
)
