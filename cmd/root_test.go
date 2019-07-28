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
	"os"
	"strings"
	"testing"

	"github.com/portworx/px/pkg/util"
	"github.com/stretchr/testify/assert"
)

// From https://gist.github.com/imosquera/6716490#sthash.O4z2aQQp.LUHz2Cbb.dpuf
type Restorer func()

func (r Restorer) Restore() {
	r()
}

// Returns a buffer for stdout, stderr, and a function.
// The function should be used as a defer to restore the state
// See status_test.go for an example
func pxTestSetupCli(args string) (*bytes.Buffer, *bytes.Buffer, Restorer) {
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

func runPx() error {
	return rootCmd.Execute()
}

func TestMultipleClustersContextsFromCli(t *testing.T) {
	// This test depends on the ./hack/config.yml to have
	// a 'source' and 'target' context and two running mock-sdk-servers

	// Create a volume on the source that is now in the target.
	// We will use this to differentiate them.
	vol := "mysourcevolume"
	so, _, r := pxTestSetupCli("px create volume " + vol + " --size=10")
	err := runPx()
	r()
	assert.NoError(t, err)

	// Get the volume info on the source
	so, _, r = pxTestSetupCli("px get volumes")
	err = runPx()
	r()
	assert.NoError(t, err)
	lines := strings.Split(so.String(), "\n")
	assert.True(t, util.ListContainsSubString(lines, vol))

	// Fail to get that information on the target
	so, _, r = pxTestSetupCli("px --context=target get volumes")
	err = runPx()
	r()
	assert.NoError(t, err)
	lines = strings.Split(so.String(), "\n")
	assert.False(t, util.ListContainsSubString(lines, vol))
}
