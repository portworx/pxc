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
package cluster_test

import (
	"strings"
	"testing"

	"github.com/portworx/pxc/handler/test"
	"github.com/stretchr/testify/assert"
)

func getKey(s string, delim string) []string {
	x := strings.Split(s, delim)
	return x
}

func TestDescribeCluster(t *testing.T) {

	// Setup to run: pxc describe cluster and return the Stdout, Stderr,
	// and a function to restore state
	so, _, r := test.PxTestSetupCli("px describe cluster")

	// Defer to cleanup state
	defer r()

	// Start the CLI
	err := test.RunPx()
	assert.NoError(t, err)

	lines := strings.Split(so.String(), "\n")
	index := 0
	k := getKey(lines[index], ":")
	assert.Contains(t, k[0], "Cluster ID")
	index++
	k = getKey(lines[index], ":")
	assert.Contains(t, k, "Cluster UUID")
	index++
	k = getKey(lines[index], ":")
	assert.Contains(t, k, "Cluster Status")
	index++
	k = getKey(lines[index], ":")
	assert.Contains(t, k, "Version")
	index++
	k = getKey(lines[index], ":")
	/* need to fix mock server
	   mock server returns example: data instead of build
	*/
	//assert.Contains(t, k, "build")
	index++
	k = getKey(lines[index], " ")
	assert.Contains(t, k, "SDK")
	assert.Contains(t, k[1], "Version")
	index++
	index++
	k = getKey(lines[index], " ")
	assert.Contains(t, k[0], "Hostname")
	assert.Contains(t, k[6], "IP")
	assert.Contains(t, k[16], "SchedulerNodeName")
	assert.Contains(t, k[18], "Used")
	assert.Contains(t, k[20], "Capacity")
	assert.Contains(t, k[22], "Status")
}
