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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func checkOutput(t *testing.T, out *DefaultFormatOutput, formatStr string) {
	out.SetFormat(formatStr)
	v, err := GetFormattedOutput(out)
	assert.Equal(t, err, nil, "Got error in formatting")
	assert.Equalf(t, v, out.Desc, "%s output is not correct", formatStr)
}

func TestDefaultFormatOutput(t *testing.T) {
	out := &DefaultFormatOutput{
		Id:   []string{"new uuid"},
		Cmd:  "Create Test",
		Desc: "Create Test with new uuid successful",
	}

	checkOutput(t, out, "wide")
	checkOutput(t, out, "")

	out.SetFormat("yaml")
	v, err := GetFormattedOutput(out)
	assert.Equal(t, err, nil, "Got error in formatting")
	var newout DefaultFormatOutput
	err = yaml.Unmarshal([]byte(v), &newout)
	assert.NoError(t, err, "yaml unmarshal failed")
	assert.Equal(t, newout.Id[0], out.Id[0], "yaml output not correct for Id")
	assert.Equal(t, newout.Cmd, out.Cmd, "yaml output not correct for Cmd")
	assert.Equal(t, newout.Desc, out.Desc, "yaml output not correct for Desc")

	newout = DefaultFormatOutput{}
	out.SetFormat("json")
	v, err = GetFormattedOutput(out)
	assert.Equal(t, err, nil, "Got error in formatting")
	err = json.Unmarshal([]byte(v), &newout)
	assert.NoError(t, err, "yaml unmarshal failed")
	assert.Equal(t, newout.Id[0], out.Id[0], "yaml output not correct for Id")
	assert.Equal(t, newout.Cmd, out.Cmd, "yaml output not correct for Cmd")
	assert.Equal(t, newout.Desc, out.Desc, "yaml output not correct for Desc")
}
