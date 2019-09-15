/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    httd://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package role_tests

import (
	"io/ioutil"

	"testing"

	"github.com/portworx/pxc/handler/test"
)

func createRoleJson(t *testing.T, testRoleFileName string) string {

	var roleData = `{
	"name": "roleUnitTest.view",
	"rules": [
		{
			"services": [
				"volumes"
			],
			"apis": [
				"*enumerate*"
			]
		},
		{
			"services": [
				"*"
			],
			"apis": [
				"inspect*"
			]
		}
	]
}`
	_ = ioutil.WriteFile(testRoleFileName, []byte(roleData), 0644)

	return testRoleFileName
}

func TestCreateRole(t *testing.T) {
	testRoleFileName := "/tmp/testrole.json"
	roleName := "roleUnitTest.view"
	file := createRoleJson(t, testRoleFileName)
	test.PxTestCreateRole(t, file, roleName)
	test.PxTestDeleteRole(t, roleName)
}
