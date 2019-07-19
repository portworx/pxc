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

type DefaultFormatOutput struct {
	BaseFormatOutput `json:",omitempty" yaml:",omitempty"`
	Cmd              string   `json:"cmd,omitempty" yaml:"cmd,omitempty"`
	Desc             string   `json:"desc,omitempty" yaml:"desc,omitempty"`
	Id               []string `json:"id,omitempty" yaml:"id,omitempty"`
}

// DefaultFormat just returns the Desc
func (dfo *DefaultFormatOutput) DefaultFormat() string {
	return dfo.Desc
}

// WideFormat just returns the DefaultFormat
func (dfo *DefaultFormatOutput) WideFormat() string {
	return dfo.DefaultFormat()
}

// String returns the formatted output of the object as per the format set.
func (dfo *DefaultFormatOutput) String() string {
	return GetFormattedOutput(dfo)
}
