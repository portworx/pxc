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
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"net"
	"os"
	"strings"
	"time"
)

var (
	ErrInvalidEndpoint = errors.New("Invalid Endpoint")
)

const (
	DefaultPort = "9020"
)

// ListContains returns true when string s is found in the list
func ListContainsSubString(list []string, s string) bool {
	for _, value := range list {
		if strings.Contains(value, s) {
			return true
		}
	}
	return false
}

func StringContainsAnyFromList(s string, list []string) bool {
	for _, value := range list {
		if strings.Contains(s, value) {
			return true
		}
	}
	return false
}

// ListContains returns true when string s is found in the list
func ListContains(list []string, s string) bool {
	for _, value := range list {
		if value == s {
			return true
		}
	}
	return false
}

// ListsHaveMatch returns the string and true if any one string is found in both lists
func ListHaveMatch(list, match []string) (string, bool) {
	for _, s := range match {
		if ListContains(list, s) {
			return s, true
		}
	}
	return "", false
}

// StringMapToCommaString returns a comma separated k=v as a single string
func StringMapToCommaString(labels map[string]string) string {
	s := make([]string, 0, len(labels))
	for k, v := range labels {
		s = append(s, k+"="+v)
	}
	return strings.Join(s, ",")
}

// Generate random string with the given prefix, by appending random numbers.
func GetRandomName(prefix string) string {
	return fmt.Sprintf("%s-%v", prefix, time.Now().Unix())
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

// IsFileExists returns true if the file exists else false
func IsFileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// ValidateEndpoint will valid whether given endpoint is a valid.
// Following validation will be done on the endpoint:
// 1. Check whether the host part (IP or hostname) is valid
// 2. Check whether port number is present in the endpoint, If not add Default_port.
func ValidateEndpoint(endpoint string) (string, error) {
	if !govalidator.IsIP(endpoint) {
		host, port, _ := net.SplitHostPort(endpoint)
		if len(port) != 0 && len(host) != 0 {
			// Endpoint has both host and port. Is it valid?
			// For example:
			// Valid: "192.168.1.1:9020" or "localhost:9020"
			// InValid: "192.168.1.500:9020" or "local*host:9020"
			if govalidator.IsIP(host) || govalidator.IsHost(host) {
				return endpoint, nil
			} else {
				return "", ErrInvalidEndpoint
			}
		} else if len(port) == 0 && len(host) != 0 {
			// Missing port after ":", add the DefaultPort value
			// For example: "localhost:" or "192.168.110.0:"
			if govalidator.IsIP(host) || govalidator.IsHost(host) {
				updatedEndpoint := net.JoinHostPort(host, DefaultPort)
				return updatedEndpoint, nil
			} else {
				return "", ErrInvalidEndpoint
			}
		} else if len(port) == 0 && len(host) == 0 {
			// Missing port as : was missing in the endpoint(in the case of hostname)
			// If it is a valid hostname, add DefautPort.
			if govalidator.IsHost(endpoint) {
				updatedEndpoint := net.JoinHostPort(endpoint, DefaultPort)
				return updatedEndpoint, nil
			} else {
				return "", ErrInvalidEndpoint
			}

		} else {
			return "", ErrInvalidEndpoint
		}
	}
	// Missing port in the case of proper IP address, add the DefaultPort value
	// For example: "192.168.1.1"
	updatedEndpoint := net.JoinHostPort(endpoint, DefaultPort)
	return updatedEndpoint, nil
}
