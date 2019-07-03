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
)

// ListContains returns true when string s is found in the list
func ListContains(list []string, s string) bool {
	for _, value := range list {
		if value == s {
			return true
		}
	}
	return false
}

// ListsHaveMatch returns true when any one string is found in both lists
func ListHaveMatch(list, match []string) bool {
	for _, s := range match {
		if ListContains(list, s) {
			return true
		}
	}
	return false
}

// StringMapToCommaString returns a comma separated k=v as a single string
func StringMapToCommaString(labels map[string]string) string {
	s := make([]string, 0, len(labels))
	for k, v := range labels {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, ",")
}

// CommaStringToStringMap returns a string map composed of the k=v comma
// separated values in the string
func CommaStringToStringMap(s string) (map[string]string, error) {
	kvMap := make(map[string]string)

	for _, pair := range strings.Split(s, ",") {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid pair %s", kv)
		}
		k := kv[0]
		v := kv[1]
		if len(k) == 0 || len(v) == 0 {
			return nil, fmt.Errorf("'%s' is invalid", pair)
		}
		kvMap[k] = v
	}

	return kvMap, nil
}
