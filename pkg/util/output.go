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
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
	"gopkg.in/yaml.v2"
)

var (
	// Stdout points to the output buffer to send screen output
	Stdout = os.Stdout
	// Stderr points to the output buffer to send errors to the screen
	Stderr = os.Stderr
)

// Printf is just like fmt.Printf except that it send the output to Stdout. It
// is equal to fmt.Fprintf(util.Stdout, format, args)
func Printf(format string, args ...interface{}) {
	fmt.Fprintf(Stdout, format, args...)
}

// Eprintf prints the errors to the output buffer Stderr. It is equal to
// fmt.Fprintf(util.Stderr, format, args)
func Eprintf(format string, args ...interface{}) {
	fmt.Fprintf(Stderr, format, args...)
}

// PrintYaml prints the object to yaml to Stdout
func PrintYaml(obj interface{}) {
	bytes, err := yaml.Marshal(obj)
	if err != nil {
		Eprintf("Unable to create yaml output")
		return
	}
	Printf(string(bytes))
}

// PrintJson prints the object to json to Stdout
func PrintJson(obj interface{}) {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		Eprintf("Unable to create json output")
		return
	}
	Printf(string(bytes))
}

// NewTabby is used to return a tabbing object set to the
// value of Stdout in the util package
func NewTabby() *tabby.Tabby {
	writer := tabwriter.NewWriter(Stdout, 0, 0, 2, ' ', 0)
	return tabby.NewCustom(writer)
}
