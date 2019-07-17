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

package util

import (
	"testing"
	"reflect"

	"github.com/stretchr/testify/assert"
)

/*
 To test postivive case of utils.ListContains function.
 Test if given element is present in the list.
 Assert if element is not found in the list.
 */
func TestListContainsUT1(t *testing.T) {
	// list containg elements
	elements := []string{"node", "drive", "volume"}
	matchString := "volume"

	ret := ListContains(elements, matchString)
	assert.Equal(t, ret, true, "%s entity not found in the list", matchString)
}

/*
To test negative case of utils.ListContains function.
Test if given element is not present in the list.
Asserts if element is found in the list.
*/
func TestListContainsUT2(t *testing.T) {
	// list containg elements
	elements := []string{"node", "drive", "volume"}
	matchString := "portworx"

	ret := ListContains(elements, matchString)
	assert.Equal(t, ret, false, "%s entity found in the list", matchString)
}

/*
To test positive case of utils.ListHaveMatch.
Test if given element present in the both the list.
Assert if none of the elements is not found in both the list.
*/
func TestListHaveMatchUT1(t *testing.T) {
	elements := []string{"node", "drive", "volume", "portworx"}
	match := []string{"portworx", "osd"}

	ret := ListHaveMatch(elements, match)
	assert.Equal(t, ret, true, "Elements %s not found in the list" ,match)
}

/*
Test for negative case of utils.ListHaveMatch.
Tests if the given entity is not present in the list.
Assert if any one of the elements is found in both the list.
*/
func TestListHaveMatchUT2(t *testing.T) {
	elements := []string{"node", "drive", "volume", "portworx"}
	match := []string{"oci", "osd"}

	ret := ListHaveMatch(elements, match)
	assert.Equal(t, ret, false, "One of the elements %s found in the list" ,match)
}

/*
Test for positive case of utils.StringMapToCommaString
Test if the given map can be converted to valid string.
Assert on conversion failure.
*/
func TestStringMapToCommaString(t *testing.T) {
	elements := map[string]string {
		"pod": "portworx",
		"cluster": "k8s",
	}
	expectedResult := "pod=portworx,cluster=k8s"
	ret := StringMapToCommaString(elements)
	assert.Equal(t, ret, expectedResult, "Failed to convert (k,v) to string")
}


/*
Test for positive case of utils.CommaStringToStringMap
Tests if the given valid string is converted to  (k,v) pair.
Asserts if the conversion fails
*/
func TestCommaStringToStringMapUT1(t *testing.T) {
	element := "pod=portworx,cluster=k8s"
	expectedResult := map[string]string {
		"pod": "portworx",
		"cluster": "k8s",
	}

	ret, _ := CommaStringToStringMap(element)
	state := reflect.DeepEqual(ret, expectedResult)
	assert.Equal(t, state, true, "Failed to convert string %s to (k,v) pair", element)
}

/*
Test function for negative case of utils.CommaStringToStringMap
Tests if the given valid string is not converted to  (k,v) pair.
Asserts if the conversion succeeds.
*/
func TestCommaStringToStringMapUT2(t *testing.T) {
	element := "pod+portworx,cluster/k8s"
	expectedResult := map[string]string {
		"pod": "portworx",
		"cluster": "k8s",
	}

	ret, _ := CommaStringToStringMap(element)
	state := reflect.DeepEqual(ret, expectedResult)

	assert.Equal(t, state, false, "Successfully converted string %s to (k,v) pair " +
				"which shouldn't have been!!!", element)
}

/*
Test function for negative case of utils.CommaStringToStringMap
Tests if the given valid string is not converted to  (k,v) pair.
Asserts if the conversion succeeds.
*/
func TestCommaStringToStringMapUT3(t *testing.T) {
	element := "pod=portworx,cluster/k8s"
	expectedResult := map[string]string {
		"pod": "portworx",
		"cluster": "k8s",
	}

	ret, _ := CommaStringToStringMap(element)
	state := reflect.DeepEqual(ret, expectedResult)

	assert.Equal(t, state, false, "Successfully converted string %s to (k,v) pair " +
				"which shouldn't have been!!!", element)
}