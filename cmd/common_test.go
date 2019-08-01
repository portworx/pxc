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
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/tests"
	"github.com/portworx/px/pkg/util"

	"github.com/stretchr/testify/assert"
)

// Returns a buffer for stdout, stderr, and a function.
// The function should be used as a defer to restore the state
// See status_test.go for an example
func pxTestSetupCli(args string) (*bytes.Buffer, *bytes.Buffer, tests.Restorer) {
	// Save
	oldargs := os.Args
	oldStdout := util.Stdout
	oldStderr := util.Stderr
	oldcfgFile := cfgFile

	// Create new buffers
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	// Set buffers
	util.Stdout = stdout
	util.Stderr = stderr
	os.Args = strings.Split(args, " ")
	cfgFile = os.Getenv("PXTESTCONFIG")

	return stdout, stderr, func() {
		cfgFile = oldcfgFile
		os.Args = oldargs
		util.Stdout = oldStdout
		util.Stderr = oldStderr
	}
}

// runPx runs a command saved in os.Args and returns the error if any
func runPx() error {
	return rootCmd.Execute()
}

// genVolName generates a unique name for a volume appended to a prefix
func genVolName(prefix string) string {
	return fmt.Sprintf("%s-%v", prefix, time.Now().Unix())
}

// Execute cli and return the buffers of standard out, standard error,
// and error.
// Callers must managage global variables using Patch from pkg/tests
func executeCliRaw(cli string) (*bytes.Buffer, *bytes.Buffer, error) {
	so, se, r := pxTestSetupCli(cli)

	// Defer to cleanup state
	defer r()

	// Start the CLI
	err := runPx()

	return so, se, err
}

// Execute cli and return string lists of standard out, standard error,
// and error if any.
// Callers must managage global variables using Patch from pkg/tests
func executeCli(cli string) ([]string, []string, error) {
	so, se, err := executeCliRaw(cli)

	return strings.Split(so.String(), "\n"),
		strings.Split(se.String(), "\n"),
		err
}

// Takes a volume name and size. Returns the created volume id.
// For some reason our test container only recoganizes id and not name for some calls.
func testCreateVolume(t *testing.T, volName string, size uint64) {
	cli := "px create volume " + volName + " --size " + strconv.FormatUint(size, 10)
	lines, _, err := executeCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Volume %s created with id", volName)))
}

// Takes a volume id assert volume exists
func testHasVolume(id string) bool {
	// Get volume information
	cli := "px get volume " + id
	_, _, err := executeCli(cli)

	return err == nil
}

// Return volume information
// TODO: If necessary, we can do a `px get volume <id> -o json` then
//       unmarshal the JSON to appropriate object
// 		 then return the &api.Volume inside of it.
func testVolumeInfo(t *testing.T, id string) *api.Volume {
	cli := fmt.Sprintf("px get volume %s -o json", id)
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
func testGetAllVolumes(t *testing.T) []string {
	cli := "px get volume"
	lines, _, err := executeCli(cli)
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
func testDeleteVolume(t *testing.T, volName string) {
	cli := "px delete volume " + volName
	_, _, err := executeCli(cli)
	assert.NoError(t, err)
}

// Takes a volume name and snapshot name
func testCreateSnapshot(t *testing.T, volId string, snapName string) {
	cli := fmt.Sprintf("px create volumesnapshot %s %s", volId, snapName)
	lines, _, err := executeCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Snapshot of %s created with id", volId)))
}

// Takes a volume name and clone name
func testCreateClone(t *testing.T, volId string, cloneName string) {
	cli := fmt.Sprintf("px create volumeclone %s %s", volId, cloneName)
	lines, _, err := executeCli(cli)
	assert.NoError(t, err)

	assert.True(t, util.ListContainsSubString(lines, fmt.Sprintf("Clone of %s created with id", volId)))
}
