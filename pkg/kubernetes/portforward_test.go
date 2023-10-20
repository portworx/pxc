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
package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Manual test. You will need a running kubernetes system for this
// test to run
func TestKubectlPortForwarder(t *testing.T) {
	t.Skip()
	kubeConfig := "path/to/your/kubeconfig.conf"
	p := newKubectlPortForwarder(kubeConfig)
	err := p.Start()
	assert.NoError(t, err)
	assert.NotEmpty(t, p.Endpoint())
	err = p.Stop()
	assert.NoError(t, err)
}

func TestGetEndpointFromKubectlOutput(t *testing.T) {

	tests := []struct {
		expectFailure  bool
		expectedOutput string
		buffer         string
	}{
		{
			expectFailure:  false,
			expectedOutput: "localhost:12345",
			buffer:         " Forwarding from 127.0.0.1:12345 --> 9020",
		},
		{
			expectFailure:  false,
			expectedOutput: "[::1]:35572",
			buffer:         " Forwarding from [::1]:35572 --> 9020",
		},
		{
			expectFailure: true,
			buffer:        " Forwarding from ]:35572 --> 9020",
		},
		{
			expectFailure: true,
			buffer:        " Bad",
		},
		{
			expectFailure:  false,
			buffer:         "Forwarding from 127.0.0.1:41240 -> 9020\n Forwarding from [::1]:41240 -> 9020",
			expectedOutput: "localhost:41240",
		},
	}

	p := &KubectlPortForwarder{}
	for _, test := range tests {
		s, err := p.getEndpointFromKubectlOutput(test.buffer)
		if test.expectFailure {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, s, test.expectedOutput)
		}
	}
}
