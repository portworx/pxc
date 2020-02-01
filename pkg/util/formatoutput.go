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
type FormatOutput interface {
	// SetFormat takes in as input the type of Formatting needed.
	// Currently recoganized values are "wide", "json" and "yaml".
	// Any other string including "" will end up with DefaultFormat.
	SetFormat(typeOfFormatting string)

	// GetFormat returns the format type set
	GetFormat() string

	// DefaultFormat formats the output as a regular string
	// This is called when either "wide", "yaml" or "json" is not set
	DefaultFormat() (string, error)

	// WideFormat formats the object in the "wide" format
	WideFormat() (string, error)

	// YamlFormat formats the object in the "yaml" format
	YamlFormat() (string, error)

	// JsonFormat formats the object in the "json" format
	JsonFormat() (string, error)
}

// GetFormattedOutput returns the formatted output
func GetFormattedOutput(in FormatOutput) (string, error) {
	switch in.GetFormat() {
	case FORMAT_YAML:
		return in.YamlFormat()
	case FORMAT_JSON:
		return in.JsonFormat()
	case FORMAT_WIDE:
		return in.WideFormat()
	default:
		return in.DefaultFormat()
	}
}

// Print the formatted output to stdout
// In case of any  error, just return the error
func PrintFormatted(in FormatOutput) error {
	str, err := GetFormattedOutput(in)
	if err != nil {
		return err
	}
	Printf("%s\n", str)
	return nil
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
func (bfo *BaseFormatOutput) DefaultFormat() (string, error) {
	return "", nil
}

// WideFormat just returns the DefaultFormat
func (bfo *BaseFormatOutput) WideFormat() (string, error) {
	return "", nil
}

// JsonFormat just returns the DefaultFormat
func (bfo *BaseFormatOutput) JsonFormat() (string, error) {
	return "", nil
}

// YamlFormat just returns the YamlFormat
func (bfo *BaseFormatOutput) YamlFormat() (string, error) {
	return "", nil
}
