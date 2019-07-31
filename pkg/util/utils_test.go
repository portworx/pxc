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
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListContainsSubString(t *testing.T) {
	tests := []struct {
		list  []string
		s     string
		found bool
	}{
		{
			list:  []string{"Hello", "World"},
			s:     "ell",
			found: true,
		},
		{
			list:  []string{"Hello", "World"},
			s:     "nothere",
			found: false,
		},
		{
			list:  []string{"Hello", "World"},
			s:     "Hellow",
			found: false,
		},
		{
			list:  []string{"Hello", "World"},
			s:     "Hello",
			found: true,
		},
		{
			list:  []string{},
			s:     "Hello",
			found: false,
		},
		{
			list:  []string{},
			s:     "",
			found: false,
		},
	}

	for _, test := range tests {
		assert.True(t, test.found == ListContainsSubString(test.list, test.s))
	}
}

/*
 To test postivive case of utils.ListContains function.
 Test if given element is present in the list.
 Assert if element is not found in the list.
*/
func TestListContainsElement(t *testing.T) {
	// list containg elements
	elements := []string{"node", "drive", "volume"}
	matchString := "volume"

	ret := ListContains(elements, matchString)
	assert.Equal(t, ret, true)
}

/*
To test negative case of utils.ListContains function.
Test if given element is not present in the list.
Asserts if element is found in the list.
*/
func TestListContainsNoElement(t *testing.T) {
	// list containg elements
	elements := []string{"node", "drive", "volume"}
	matchString := "portworx"

	ret := ListContains(elements, matchString)
	assert.Equal(t, ret, false)
}

/*
To test positive case of utils.ListHaveMatch.
Test if given element present in the both the list.
Assert if none of the elements is not found in both the list.
*/
func TestListHaveMatchPresent(t *testing.T) {
	elements := []string{"node", "drive", "volume", "portworx"}
	match := []string{"portworx", "osd"}

	m, ret := ListHaveMatch(elements, match)
	assert.Equal(t, ret, true)
	assert.Equal(t, m, "portworx")
}

/*
Test for negative case of utils.ListHaveMatch.
Tests if the given entity is not present in the list.
Assert if any one of the elements is found in both the list.
*/
func TestListHaveMatchNotPresent(t *testing.T) {
	elements := []string{"node", "drive", "volume", "portworx"}
	match := []string{"oci", "osd"}

	_, ret := ListHaveMatch(elements, match)
	assert.Equal(t, ret, false)
}

/*
Test for positive case of utils.StringMapToCommaString
Test if the given map can be converted to valid string.
Assert on conversion failure.
*/
func TestStringMapToCommaString(t *testing.T) {
	elements := map[string]string{
		"pod":     "portworx",
		"cluster": "k8s",
	}

	ret := StringMapToCommaString(elements)
	val, _ := CommaStringToStringMap(ret)
	state := reflect.DeepEqual(val, elements)
	assert.Equal(t, state, true)
}

/*
Test for positive case of utils.CommaStringToStringMap
Tests if the given valid string is converted to  (k,v) pair.
Asserts if the conversion fails
*/
func TestCommaStringToStringMapPositive(t *testing.T) {
	element := "pod=portworx,cluster=k8s"
	expectedResult := map[string]string{
		"pod":     "portworx",
		"cluster": "k8s",
	}

	ret, _ := CommaStringToStringMap(element)
	state := reflect.DeepEqual(ret, expectedResult)
	assert.Equal(t, state, true)
}

/*
Test function for negative cases of utils.CommaStringToStringMap
Tests if the given valid string is not converted to  (k,v) pair.
Asserts if the conversion succeeds.
*/
func TestCommaStringToStringMapNegative(t *testing.T) {
	// case 1
	element := "pod+portworx,cluster/k8s"
	expectedResult := map[string]string{
		"pod":     "portworx",
		"cluster": "k8s",
	}

	state := deepCompare(element, expectedResult)
	assert.Equal(t, state, false)

	// case 2
	element = "pod=portworx,cluster/k8s"
	expectedResult = map[string]string{
		"pod":     "portworx",
		"cluster": "k8s",
	}

	state = deepCompare(element, expectedResult)
	assert.Equal(t, state, false)
}

/*
Compares provided map with map generated as part of CommaStringToStringMap
Return true or false
*/
func deepCompare(element string, expectedResult map[string]string) (state bool) {
	//ret is a map
	ret, _ := CommaStringToStringMap(element)
	state = reflect.DeepEqual(ret, expectedResult)
	return
}

var isFileExistsTests = []struct {
	testCase string
	status   bool
}{
	{
		"existsFile",
		true,
	},
	{
		"nonExistsFile",
		false,
	},
}

/*
Test function to test various possible case of IsFileExists util function
*/
func TestIsFileExists(t *testing.T) {
	var filename = ""

	for _, test := range isFileExistsTests {

		filename = fmt.Sprintf("/tmp/%s", GetRandomName(test.testCase))
		if strings.Compare(test.testCase, "existsFile") == 0 {
			// Create the file with random name
			_, err := os.Create(filename)
			assert.Equal(t, err, nil)
		}
		status := IsFileExists(filename)
		assert.Equal(t, test.status, status)

		if strings.Compare(test.testCase, "existsFile") == 0 {
			// Remove the file
			err := os.Remove(filename)
			assert.Equal(t, err, nil)
		}
	}
}
