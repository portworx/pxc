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

const (
	FORMAT_WIDE = "wide"
	FORMAT_JSON = "json"
	FORMAT_YAML = "yaml"
)

// FormatOutput is the interface used to ensure proper formatting
// json and yaml does not need methods since the ToYaml() and ToJson()
// functions are used to format these objects.
type FormatOutput interface {
	// SetFormat takes in as input the type of Formatting needed.
	// Currently recoganized values are "wide", "json" and "yaml".
	// Any other string including "" will end up with DefaultFormat.
	SetFormat(typeOfFormatting string)

	// GetFormat returns the format type set
	GetFormat() string

	// DefaultFormat formats the output as a regular string
	// This is called when either "wide", "yaml" or "json" is not set
	DefaultFormat() string

	// WideFormat formats the object in the "wide" format
	WideFormat() string

	// Print writes out the object in appropriate format
	Print()
}

// GetFormattedOutput returns the formatted output
func GetFormattedOutput(in FormatOutput) string {
	switch in.GetFormat() {
	case "yaml":
		return ToYaml(in)
	case "json":
		return ToJson(in)
	case "wide":
		return in.WideFormat()
	default:
		return in.DefaultFormat()
	}
}

type BaseFormatOutput struct {
	FormatType string
}

// SetFormat takes in as input the type of Formatting needed.
// Currently recoganized values are "wide", "json" and "yaml".
// Any other string including "" will end up with DefaultFormat.
func (bfo *BaseFormatOutput) SetFormat(typeOfOutput string) {
	bfo.FormatType = typeOfOutput
}

// GetFormat returns the format type set
func (bfo *BaseFormatOutput) GetFormat() string {
	return bfo.FormatType
}

// DefaultFormat just returns the Desc
func (bfo *BaseFormatOutput) DefaultFormat() string {
	return ""
}

// WideFormat just returns the DefaultFormat
func (bfo *BaseFormatOutput) WideFormat() string {
	return ""
}

// String returns the formatted output of the object as per the format set.
func (bfo *BaseFormatOutput) String() string {
	return GetFormattedOutput(bfo)
}

// Print writes the object to stdout
func (bfo *BaseFormatOutput) Print() {
	Printf("%v", bfo)
}
