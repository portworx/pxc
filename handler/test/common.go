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
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/pxc/cmd"

	_ "github.com/portworx/pxc/handler"
	"github.com/portworx/pxc/pkg/tests"
	"github.com/portworx/pxc/pkg/util"

	"github.com/stretchr/testify/assert"
)

// Returns a buffer for stdout, stderr, and a function.
// The function should be used as a defer to restore the state
// See status_test.go for an example
func PxTestSetupCli(args string) (*bytes.Buffer, *bytes.Buffer, tests.Restorer) {
	// Save
	oldargs := os.Args
	oldStdout := util.Stdout
	oldStderr := util.Stderr

	// Create new buffers
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	// Set buffers
	util.Stdout = stdout
	util.Stderr = stderr
	os.Args = strings.Split(args, " ")

	return stdout, stderr, func() {
		os.Args = oldargs
		util.Stdout = oldStdout
		util.Stderr = oldStderr
	}
}

// runPx runs a command saved in os.Args and returns the error if any
func RunPx() error {
	return cmd.Main()
}

// genVolName generates a unique name for a volume appended to a prefix
func GenVolName(prefix string) string {
	return util.GetRandomName(prefix)
}

// Execute cli and return the buffers of standard out, standard error,
// and error.
// Callers must managage global variables using Patch from pkg/tests
func executeCliRaw(cli string) (*bytes.Buffer, *bytes.Buffer, error) {
	so, se, r := PxTestSetupCli(cli)

	// Defer to cleanup state
	defer r()

	// Start the CLI
	err := RunPx()

	return so, se, err
}

// Execute cli and return string lists of standard out, standard error,
// and error if any.
// Callers must managage global variables using Patch from pkg/tests
func ExecuteCli(cli string) ([]string, []string, error) {
	so, se, err := executeCliRaw(cli)

	return strings.Split(so.String(), "\n"),
		strings.Split(se.String(), "\n"),
		err
}

// Takes a volume name and size. Returns the created volume id.
// For some reason our test container only recoganizes id and not name for some calls.
func PxTestCreateVolume(t *testing.T, volName string, size uint64) {
	cli := fmt.Sprintf("pxc create volume %s --size %s --groups group1:r --collaborators user1:w", volName, strconv.FormatUint(size, 10))
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Takes a volume name and size. Returns the created volume id.
func PxTestCreateVolumeWithLabel(t *testing.T, volName string, size uint64, labels string) {
	cli := fmt.Sprintf("pxc create volume %s --size %s --labels %s",
		volName, strconv.FormatUint(size, 10), labels)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Takes a volume id assert volume exists
func PxTestHasVolume(id string) bool {
	// Get volume information
	cli := "pxc get volume " + id
	_, _, err := ExecuteCli(cli)

	return err == nil
}

func PxTestGetVolumeWithLabels(t *testing.T, selector string) (*bytes.Buffer, error) {
	cli := fmt.Sprintf("pxc get volume --show-labels --selector %s", selector)
	so, _, err := executeCliRaw(cli)
	return so, err
}

func PxTestGetVolumeWithNameSelector(t *testing.T, volName string, selector string) {
	cli := fmt.Sprintf("pxc get volume %s --show-labels --selector %s", volName, selector)
	_, _, err := executeCliRaw(cli)
	assert.Error(t, err)
}

// Return volume information
// TODO: If necessary, we can do a `pxc get volume <id> -o json` then
//       unmarshal the JSON to appropriate object
// 		 then return the &api.Volume inside of it.
func PxTestVolumeInfo(t *testing.T, id string) *api.Volume {
	cli := fmt.Sprintf("pxc get volume %s -o json", id)
	so, _, err := executeCliRaw(cli)
	assert.NoError(t, err)

	var vols []api.SdkVolumeInspectResponse
	err = json.Unmarshal([]byte(so.String()), &vols)
	assert.NoError(t, err)
	assert.Len(t, vols, 1)

	return vols[0].GetVolume()
}

// Returns a list of all volume ids
// TODO: We may need to bring this back to the TestXXX functions
//       depending on what it does, because the output would be
//       parse, and as a library function, it may be easier to
//       get a specific volume.
func PxTestGetAllVolumes(t *testing.T) []string {
	cli := fmt.Sprintf("pxc get volume")
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	lines = lines[2:]
	volNames := make([]string, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		x := strings.Split(l, " ")
		volNames = append(volNames, x[0])
	}
	return volNames
}

// Deletes specified volume
func PxTestDeleteVolume(t *testing.T, volName string) {
	cli := fmt.Sprintf("pxc delete volume %s", volName)
	_, _, err := ExecuteCli(cli)
	assert.NoError(t, err)
}

// Takes a volume name and snapshot name
func PxTestCreateSnapshot(t *testing.T, volId string, snapName string) {
	cli := fmt.Sprintf("pxc create volumesnapshot %s %s", volId, snapName)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Snapshot of %s created with id", volId)))
}

// Takes a volume name and clone name
func PxTestCreateClone(t *testing.T, volId string, cloneName string) {
	cli := fmt.Sprintf("pxc create volumeclone %s %s", volId, cloneName)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Clone of %s created with id", volId)))
}

// Helper function to create volume with "sticky" flag set
func PxTestCreateStickyVolume(t *testing.T, volName string, size uint64) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --sticky",
		volName, size)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Helper function to create volume with "encryption" flag set
func PxTestCreateEncrypVolume(t *testing.T, volName string, size uint64) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --encryption",
		volName, size)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Helper function to create volume with "journal" flag set
func PxTestCreateJournalVolume(t *testing.T, volName string, size uint64) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --journal",
		volName, size)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Helper function to create volume with access (--groups and --collaborators) flag set
func PxTestCreateVolumeWithAccess(t *testing.T, volName string, size uint64, groups string, collaborators string) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --groups %s --collaborators %s",
		volName, size, groups, collaborators)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Helper function to create volume with "aggregation level" flag set
func PxTestCreateAggrVolume(t *testing.T, volName string, size uint64, aggrLevel uint32) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --aggregation-level %d",
		volName, size, aggrLevel)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Helper function to create volume with "io profile" flag set
func PxTestCreateIoProfVolume(t *testing.T, volName string, size uint64, IoProfile string) {
	cli := fmt.Sprintf("pxc create volume %s --size %d --ioprofile %s",
		volName, size, IoProfile)
	lines, _, err := ExecuteCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

func PxTestPatchVolumeHalevel(t *testing.T, volName string, haLevel int) {
	cli := fmt.Sprintf("pxc patch volume %s --halevel %d", volName, haLevel)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeHalevelWithNodes(t *testing.T, volName string, haLevel int64, node string) {
	cli := fmt.Sprintf("pxc patch volume %s --halevel %d --node %s", volName, haLevel, node)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines)
}

func PxTestPatchVolumeResize(t *testing.T, volName string, size uint64) {
	cli := fmt.Sprintf("pxc patch volume %s --size %d", volName, size)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeShared(t *testing.T, volName string, shared string) {
	cli := fmt.Sprintf("pxc patch volume %s --shared %s", volName, shared)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeAddCollaborators(t *testing.T, volName string, collaborators string) {
	cli := fmt.Sprintf("pxc patch volume %s  --add-collaborators %s", volName, collaborators)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeRemoveCollaborators(t *testing.T, volName string, collaborators string) {
	cli := fmt.Sprintf("pxc patch volume %s  --remove-collaborators %s", volName, collaborators)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeRemoveAllCollaborators(t *testing.T, volName string) {
	cli := fmt.Sprintf("pxc patch volume %s  --remove-all-collaborators=true", volName)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeAddGroups(t *testing.T, volName string, groups string) {
	cli := fmt.Sprintf("px patch volume %s  --add-groups %s", volName, groups)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeRemoveGroups(t *testing.T, volName string, groups string) {
	cli := fmt.Sprintf("px patch volume %s  --remove-groups %s", volName, groups)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}

func PxTestPatchVolumeRemoveAllGroups(t *testing.T, volName string) {
	cli := fmt.Sprintf("px patch volume %s  --remove-all-groups=true", volName)
	lines, _, _ := ExecuteCli(cli)
	assert.Equal(t, "Volume "+volName+" parameter updated successfully", lines[0])
}
