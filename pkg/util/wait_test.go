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
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	oneMillisecond  = 1 * time.Millisecond
	twoMillisecond  = 2 * time.Millisecond
	fiveMillisecond = 5 * time.Millisecond
	//Timeout error message
	timeoutErrMsg = "Timed out"
	// Dummy error message
	dummyErrMsg = "Error Message"
)

/* InputFunctionTimeout:
 * Test the case, where the input function takes more than the duration value.
 * Set the duration parameter value of WaitFor() to 2 Millisec.
 * Set the period parameter value of WaitFor() to 1 Millisec.
 * In the input function wait for 5 Millisec which greater than duration value.
 * Expected behaviour: WaitFor should return timeout as the input function
 *                     takes more than 2 Millisec.
 */
func InputFunctionTimeout(t *testing.T) {
	err := WaitFor(twoMillisecond, oneMillisecond, func() (bool, error) {
		time.Sleep(fiveMillisecond)
		return true, nil
	})
	assert.NotEqual(t, err, nil)
	assert.Equal(t, 0, strings.Compare(err.Error(), timeoutErrMsg))
}

/* InputFunctionReturnsFalseAndNil:
 * Test the case, when the input function return false and nil
 * Expected behaviour: WaitFor should return nil.
 */
func InputFunctionReturnsFalseAndNil(t *testing.T) {
	err := WaitFor(fiveMillisecond, oneMillisecond, func() (bool, error) {
		time.Sleep(twoMillisecond)
		return false, nil
	})
	assert.Equal(t, err, nil)
}

/* InputFunctionReturnsFalseAndError:
 * Test the case, when the input function return false and nil
 * Expected behaviour: WaitFor should return nil.
 */
func InputFunctionReturnsFalseAndError(t *testing.T) {
	err := WaitFor(fiveMillisecond, oneMillisecond, func() (bool, error) {
		time.Sleep(twoMillisecond)
		return false, fmt.Errorf(dummyErrMsg)
	})
	assert.NotEqual(t, err, nil)
	assert.Equal(t, 0, strings.Compare(err.Error(), dummyErrMsg))
}

/* InputFunctionReturnsTrueAndNil:
 * Test the case, when the input function return true and nil value
 * Expected behaviour: WaitFor should return timeoutErrMsg.
 */
func InputFunctionReturnsTrueAndNil(t *testing.T) {
	err := WaitFor(fiveMillisecond, oneMillisecond, func() (bool, error) {
		time.Sleep(twoMillisecond)
		return true, nil
	})
	assert.NotEqual(t, err, nil)
	assert.Equal(t, 0, strings.Compare(err.Error(), timeoutErrMsg))
}

/* InputFunctionReturnsTrueAndError:
 * Test the case, when the input function return true and non nil value
 * Expected behaviour: WaitFor should return nil.
 */
func InputFunctionReturnsTrueAndError(t *testing.T) {
	err := WaitFor(fiveMillisecond, oneMillisecond, func() (bool, error) {
		time.Sleep(twoMillisecond)
		return true, fmt.Errorf(dummyErrMsg)
	})
	assert.NotEqual(t, err, nil)
	assert.Equal(t, 0, strings.Compare(err.Error(), dummyErrMsg))
}

func TestWaitFor(t *testing.T) {

	// Testing timeout case
	InputFunctionTimeout(t)
	// Testing the case, where input function returns false and nil err message
	InputFunctionReturnsFalseAndNil(t)
	// Testing the case, where input function returns false and err message
	InputFunctionReturnsFalseAndError(t)
	// Testing the case, where input function returns true and nil err message
	InputFunctionReturnsTrueAndNil(t)
	// Tesing the case, where input function returns true and err message
	InputFunctionReturnsTrueAndError(t)
}
