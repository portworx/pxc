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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPxStatus(t *testing.T) {

	// Setup to run: px status and return the Stdout, Stderr,
	// and a function to restore state
	so, _, r := pxTestSetupCli("px status")

	// Defer to cleanup state
	defer r()

	// Start the CLI
	Execute()

	lines := strings.Split(so.String(), "\n")
	assert.Contains(t, lines, "Cluster ID: mock")
	assert.Contains(t, lines, "Cluster Status: STATUS_OK")
	assert.Contains(t, lines, "Version: 1.0.0-fake")
	assert.Contains(t, lines, "Cluster UUID: ce6b309e289d620be08c964969aeb39f")
	assert.Contains(t, lines[9], "node-1")

	// Only one node and an empty line
	assert.Len(t, lines[9:len(lines)], 2)

	/*
		Code:
		for i, line := range lines {
			fmt.Printf("%d:%s:\n", i, line)
		}

		Output:
		0:Cluster ID: mock:
		1:Cluster UUID: ce6b309e289d620be08c964969aeb39f:
		2:Cluster Status: STATUS_OK:
		3:Version: 1.0.0-fake:
		4:  example: data:
		5:SDK Version 0.42.14:
		6::
		7:Hostname      IP          SchedulerNodeName  Used  Capacity  Status:
		8:--------      --          -----------------  ----  --------  ------:
		9:577831342bec  172.17.0.2  node-1             0 Gi  0 Gi      STATUS_OK:
		10::
	*/
}
