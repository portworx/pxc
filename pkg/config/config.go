/*
Copyright © 2020 Portworx

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
package config

import (
	"encoding/base64"
	"encoding/json"
)

const (
	AuthKeyToken                     = "token"
	AuthKeyName                      = "name"
	AuthKeyKubernetesSecret          = "kube-secret-name"
	AuthKeyKubernetesSecretNamespace = "kube-secret-namespace"

	DefaultClusterTunnelServiceNamespace = "kube-system"
	DefaultClusterTunnelServiceName      = "portworx-api"
	DefaultClusterTunnelServicePort      = "9020"
)

var (
	// TODO: This may be removed
	config = map[string]string{}
)

// Preferences provides any pxc specific configuration
type Preferences struct {
	// Add here any pxc specific options
}

// Context provides information on who is trying to connect to a specific cluster
type Context struct {
	AuthInfo string `json:"user,omitempty" yaml:"user,omitempty"`
	Cluster  string `json:"cluster,omitempty" yaml:"cluster,omitempty"`
}

// Cluster provides information on how to connect to Portworx
type Cluster struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	CACert     string `json:"cacert,omitempty" yaml:"cacert,omitempty"`
	CACertData []byte `json:"cacert-data,omitempty" yaml:"cacert-data,omitempty"`
	Endpoint   string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Secure     bool   `json:"secure,omitempty" yaml:"secure,omitempty"`

	TunnelServiceNamespace string `json:"tunnelServiceNamespace,omitempty" yaml:"tunnelServiceNamespace,omitempty"`
	TunnelServiceName      string `json:"tunnelServiceName,omitempty" yaml:"tunnelServiceName,omitempty"`
	TunnelServicePort      string `json:"tunnelServicePort,omitempty" yaml:"tunnelServicePort,omitempty"`
}

// KubernetesAuthInfo provides information on where to access the token in Kubernetes
type KubernetesAuthInfo struct {
	SecretName      string `json:"secretName,omitempty" yaml:"secretName,omitempty"`
	SecretNamespace string `json:"secretNamespace,omitempty" yaml:"secretNamespace,omitempty"`
}

// AuthInfo provides authentication information about the user
type AuthInfo struct {
	Name               string              `json:"name,omitempty" yaml:"name,omitempty"`
	Token              string              `json:"token,omitempty" yaml:"token,omitempty"`
	KubernetesAuthInfo *KubernetesAuthInfo `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`
}

// Config is a a model to store information about the authentication and connection
// to a Portworx system.
// The design is to enable an easy extension of a Kubernetes configuration.
type Config struct {
	Preferences    Preferences          `json:"global" yaml:"global"`
	Clusters       map[string]*Cluster  `json:"clusters,omitempty" yaml:"clusters,omitempty"`
	AuthInfos      map[string]*AuthInfo `json:"users,omitempty" yaml:"users,omitempty"`
	Contexts       map[string]*Context  `json:"contexts,omitempty" yaml:"contexts,omitempty"`
	CurrentContext string               `json:"current-context,omitempty" yaml:"current-context,omitempty"`
}

func Get(k string) string {
	return config[k]
}

func Set(k, v string) {
	config[k] = v
}

func NewCluster() *Cluster {
	return &Cluster{}
}

func NewDefaultCluster() *Cluster {
	return &Cluster{
		TunnelServiceNamespace: DefaultClusterTunnelServiceNamespace,
		TunnelServiceName:      DefaultClusterTunnelServiceName,
		TunnelServicePort:      DefaultClusterTunnelServicePort,
	}
}

func NewClusterFromEncodedString(encodedString string) (*Cluster, error) {
	// decode base64
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	// unmarshal the json
	var c Cluster
	err = json.Unmarshal(decoded, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Cluster) toEncodedString() (string, error) {
	// Encode to json
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	// base64 encode
	es := base64.StdEncoding.EncodeToString(b)
	return es, nil
}

// NewAuthInfo returns an empty pxc Authinfo
func NewAuthInfo() *AuthInfo {
	return &AuthInfo{
		KubernetesAuthInfo: &KubernetesAuthInfo{},
	}
}

// NewAuthInfoFromMap returns a new pxc AuthInfo from a map. Normally used to
// create an object from the information saved in Kubeconfig
func NewAuthInfoFromMap(m map[string]string) *AuthInfo {
	a := NewAuthInfo()
	a.fromMap(m)
	return a
}

func (a *AuthInfo) toMap() map[string]string {
	m := map[string]string{
		AuthKeyName: a.Name,
	}
	if len(a.Token) != 0 {
		m[AuthKeyToken] = a.Token
	}

	if a.KubernetesAuthInfo != nil {
		if len(a.KubernetesAuthInfo.SecretName) != 0 {
			m[AuthKeyKubernetesSecret] = a.KubernetesAuthInfo.SecretName
		}
		if len(a.KubernetesAuthInfo.SecretNamespace) != 0 {
			m[AuthKeyKubernetesSecretNamespace] = a.KubernetesAuthInfo.SecretNamespace
		}
	}
	return m
}

func (a *AuthInfo) fromMap(config map[string]string) {
	a.Token = config[AuthKeyToken]
	a.Name = config[AuthKeyName]
	a.KubernetesAuthInfo.SecretName = config[AuthKeyKubernetesSecret]
	a.KubernetesAuthInfo.SecretNamespace = config[AuthKeyKubernetesSecretNamespace]
}
