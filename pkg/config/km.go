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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// KubernetesConfigManager contains all the Kubernetes configuration
type KubernetesConfigManager struct {
	kubeCliOpts *genericclioptions.ConfigFlags
}

var (
	km *KubernetesConfigManager
)

// KM returns the Kubernetes configuration flags and settings
func KM() *KubernetesConfigManager {
	if km == nil {
		km = newKubernetesConfigManager()
	}
	return km
}

func SetKM(k *KubernetesConfigManager) {
	km = k
}

func newKubernetesConfigManager() *KubernetesConfigManager {
	return &KubernetesConfigManager{
		kubeCliOpts: genericclioptions.NewConfigFlags(true),
	}
}

func NewKubernetesConfigManagerForContext(context string) *KubernetesConfigManager {
	r := newKubernetesConfigManager()
	*r.kubeCliOpts.Context = context
	return r
}

// ConfigFlags returns the kubernetes raw configuration object
func (k *KubernetesConfigManager) ConfigFlags() *genericclioptions.ConfigFlags {
	return k.kubeCliOpts
}

// ToRawKubeConfigLoader binds config flag values to config overrides
// Returns an interactive clientConfig if the password flag is enabled,
// or a non-interactive clientConfig otherwise.
// comment from k8s.io/cli-runtime
func (k *KubernetesConfigManager) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	return k.ConfigFlags().ToRawKubeConfigLoader()
}

// ToRESTConfig implements RESTClientGetter.
// Returns a REST client configuration based on a provided path
// to a .kubeconfig file, loading rules, and config flag overrides.
// Expects the AddFlags method to have been called.
// comment from k8s.io/cli-runtime
func (k *KubernetesConfigManager) ToRESTConfig() (*rest.Config, error) {
	return k.ConfigFlags().ToRESTConfig()
}

// GetStartingKubeconfig is used to adjust the current Kubernetes config. You can then
// call ModifyKubeconfig() with the modified configuration
func (k *KubernetesConfigManager) GetStartingKubeconfig() (*clientcmdapi.Config, error) {
	return k.ToRawKubeConfigLoader().ConfigAccess().GetStartingConfig()
}

// ModifyKubeconfig takes a modified configuration and seves it to disk
func (k *KubernetesConfigManager) ModifyKubeconfig(newConfig *clientcmdapi.Config) error {
	return clientcmd.ModifyConfig(k.ToRawKubeConfigLoader().ConfigAccess(), *newConfig, true)
}

// GetCurrentCluster returns configuration information about the current cluster
func (k *KubernetesConfigManager) GetCurrentCluster() (*clientcmdapi.Cluster, error) {
	kConfig, err := k.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}

	currentContext, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return nil, err
	}

	clusterInfoName := kConfig.Contexts[currentContext].Cluster
	if len(clusterInfoName) == 0 {
		return nil, fmt.Errorf("Current cluster is not set in Kubeconfig")
	}

	if clusterInfo, ok := kConfig.Clusters[clusterInfoName]; ok {
		return clusterInfo, nil
	}
	return nil, fmt.Errorf("Current user information not found in Kubeconfig")
}

// GetCurrentAuthInfo returns configuration information about the current user
func (k *KubernetesConfigManager) GetCurrentAuthInfo() (*clientcmdapi.AuthInfo, error) {
	kConfig, err := k.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}

	currentContext, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return nil, err
	}

	authInfoName := kConfig.Contexts[currentContext].AuthInfo
	if len(authInfoName) == 0 {
		return nil, fmt.Errorf("Current user is not set in Kubeconfig")
	}

	if authInfo, ok := kConfig.AuthInfos[authInfoName]; ok {
		return authInfo, nil
	}
	return nil, fmt.Errorf("Current user information not found in Kubeconfig")
}

// GetCurrentCA returns current kubernetes CA.  Note, it will return "nil, nil" if insecure-skip-tls-verify used
func (k *KubernetesConfigManager) GetCurrentCA() ([]byte, error) {
	kConfig, err := k.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}

	currentContext, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return nil, err
	}

	cluster := kConfig.Contexts[currentContext].Cluster
	if len(cluster) == 0 {
		return nil, fmt.Errorf("current cluster is not set in Kubeconfig")
	}

	ci, has := kConfig.Clusters[cluster]
	if !has {
		return nil, fmt.Errorf("current cluster information not found in Kubeconfig")
	}

	if len(ci.CertificateAuthorityData) > 0 {
		return ci.CertificateAuthorityData, nil
	} else if len(ci.CertificateAuthority) != 0 {
		return ioutil.ReadFile(ci.CertificateAuthority)
	} else if ci.InsecureSkipTLSVerify {
		return nil, nil
	}
	return nil, fmt.Errorf("could not find CA information in Kubeconfig")
}

// KubectlFlagsToCliArgs rebuilds the flags as cli args
func (k *KubernetesConfigManager) KubectlFlagsToCliArgs() string {
	var args string

	if len(*k.kubeCliOpts.KubeConfig) != 0 {
		args = "--kubeconfig=" + *k.kubeCliOpts.KubeConfig + " "
	}
	if len(*k.kubeCliOpts.Context) != 0 {
		args += "--context=" + *k.kubeCliOpts.Context + " "
	}
	if len(*k.kubeCliOpts.BearerToken) != 0 {
		args += "--token=" + *k.kubeCliOpts.BearerToken + " "
	}
	if len(*k.kubeCliOpts.APIServer) != 0 {
		args += "--server=" + *k.kubeCliOpts.APIServer + " "
	}
	if len(*k.kubeCliOpts.CAFile) != 0 {
		args += "--certificate-authority=" + *k.kubeCliOpts.CAFile + " "
	}
	if len(*k.kubeCliOpts.AuthInfoName) != 0 {
		args += "--user=" + *k.kubeCliOpts.AuthInfoName + " "
	}
	if len(*k.kubeCliOpts.CertFile) != 0 {
		args += "--client-certificate=" + *k.kubeCliOpts.CertFile + " "
	}
	if len(*k.kubeCliOpts.KeyFile) != 0 {
		args += "--client-key=" + *k.kubeCliOpts.KeyFile + " "
	}
	if len(*k.kubeCliOpts.Namespace) != 0 {
		args += "--namespace=" + *k.kubeCliOpts.Namespace + " "
	}
	return args
}

// SaveAuthInfoForKubeUser saves the pxc configuration in the kubeconfig file as a new user entry.
// Supply locationOfOrigin so that the Kubernetes saves the object with the appropriate user. LocationOfOrigin
// is found in each of the user objects in the kubernetes Config object.
func (k *KubernetesConfigManager) SaveAuthInfoForKubeUser(user, locationOfOrigin string, a *AuthInfo) error {
	pxcName := KubeconfigUserPrefix + user
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	// If one already exists it will be overwritten, if not create a new object
	if v := oldConfig.AuthInfos[pxcName]; v == nil {
		oldConfig.AuthInfos[pxcName] = clientcmdapi.NewAuthInfo()
	}

	// Store the pxc auth
	oldConfig.AuthInfos[pxcName].LocationOfOrigin = locationOfOrigin
	oldConfig.AuthInfos[pxcName].AuthProvider = &clientcmdapi.AuthProviderConfig{
		Name: "portworx",

		// Change the pxc AuthInfo to a map
		Config: a.toMap(),
	}

	// Save the information in the kubeconfig
	return k.ModifyKubeconfig(oldConfig)
}

// SaveClusterInKubeconfig stores pxc cluster configuration information in Kubeconfig
func (k *KubernetesConfigManager) SaveClusterInKubeconfig(clusterName, location string, c *Cluster) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.Clusters[pxcName]; v == nil {
		oldConfig.Clusters[pxcName] = clientcmdapi.NewCluster()
	}

	encodedString, err := c.toEncodedString()
	if err != nil {
		return err
	}

	oldConfig.Clusters[pxcName].LocationOfOrigin = location
	oldConfig.Clusters[pxcName].Server = "portworx-server"
	oldConfig.Clusters[pxcName].CertificateAuthorityData = []byte(encodedString)

	return k.ModifyKubeconfig(oldConfig)
}

// DeleteClusterInKubeconfig deletes the saved Portworx configuration in the kubeconfig
func (k *KubernetesConfigManager) DeleteClusterInKubeconfig(clusterName string) error {
	pxcName := KubeconfigUserPrefix + clusterName
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.Clusters[pxcName]; v == nil {
		return nil
	}

	delete(oldConfig.Clusters, pxcName)
	return k.ModifyKubeconfig(oldConfig)
}

// DeleteAuthInfoInKubeconfig deletes the saved Portworx configuration in the kubeconfig
func (k *KubernetesConfigManager) DeleteAuthInfoInKubeconfig(authInfoName string) error {
	pxcName := KubeconfigUserPrefix + authInfoName
	oldConfig, err := k.GetStartingKubeconfig()
	if err != nil {
		return err
	}

	if v := oldConfig.AuthInfos[pxcName]; v == nil {
		return nil
	}

	delete(oldConfig.AuthInfos, pxcName)
	return k.ModifyKubeconfig(oldConfig)
}

// GetKubernetesCurrentContext returns the context currently selected by either the config
// file or from the command line
func (k *KubernetesConfigManager) GetKubernetesCurrentContext() (string, error) {
	var contextName string

	kConfig, err := k.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return "", err
	}

	// Check if the was passed in the CLI flags
	if k.kubeCliOpts.Context != nil && len(*k.kubeCliOpts.Context) != 0 {
		contextName = *k.kubeCliOpts.Context
	} else {
		// Read it from the kubeconfig file
		contextName = kConfig.CurrentContext
	}
	if len(contextName) == 0 {
		return "", fmt.Errorf("Current context is not set or kubeconfig is missing")
	}
	logrus.Infof("CurrentContext = %s\n", contextName)

	// Check that it is actually on the kubeconfig file
	if _, ok := kConfig.Contexts[contextName]; !ok {
		return "", fmt.Errorf("context %q does not exist", contextName)
	}
	return contextName, nil
}

// Namespace returns the namespace resulting from the merged
// result of all overrides and a boolean indicating if it was
// overridden
func (k *KubernetesConfigManager) Namespace() (string, bool, error) {
	n, b, e := k.ToRawKubeConfigLoader().Namespace()
	logrus.Infof("Kubernetes namespace: ns=%s b=%v e=%v", n, b, e)
	return n, b, e
}

// ConfigSaveCluster saves the cluster configuration as part of an extension to the
// current context cluster in the Kubeconfig
func (k *KubernetesConfigManager) ConfigSaveCluster(clusterInfo *Cluster) error {

	cc := k.ToRawKubeConfigLoader()

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	// Get the current context object
	currentContext := kconfig.Contexts[currentContextName]

	// Override the name to the name of the current cluster
	clusterInfo.Name = currentContext.Cluster

	// Get the location of the kubeconfig for this specific object. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.Clusters[currentContext.Cluster].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := k.SaveClusterInKubeconfig(currentContext.Cluster, location, clusterInfo); err != nil {
		return err
	}

	logrus.Infof("Portworx server information saved in %s for Kubernetes cluster %s\n",
		location,
		currentContext.Cluster)

	return nil
}

func (k *KubernetesConfigManager) ConfigDeleteCluster(name string) error {
	cc := k.ToRawKubeConfigLoader()

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	currentContext := kconfig.Contexts[currentContextName]

	// Get the location of the kubeconfig for this specific object. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.Clusters[currentContext.Cluster].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := k.DeleteClusterInKubeconfig(currentContext.Cluster); err != nil {
		return err
	}

	logrus.Infof("Portworx server information removed from %s for Kubernetes cluster %s\n",
		location,
		currentContext.Cluster)
	return nil
}

func (k *KubernetesConfigManager) ConfigLoad() (*Config, error) {

	clusterConfig := newConfig()

	// Get the current context, either from the file or from the args to the CLI
	contextName, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return nil, fmt.Errorf("failed to get kubectl context: %v", err)
	}

	// If the current context is not set, just create an empty object.
	// This happens when there is no kubeconfig.
	if len(contextName) == 0 {
		clusterConfig.CurrentContext = contextName
		clusterConfig.Contexts[contextName] = &Context{
			AuthInfo: "",
			Cluster:  "",
		}
		return clusterConfig, nil
	}

	clientConfig := k.ConfigFlags().ToRawKubeConfigLoader()
	kConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to read kubernetes configuration: %v", err)
	}

	// Load all contexts
	for k, v := range kConfig.Contexts {
		clusterConfig.Contexts[k] = &Context{
			Name:     k,
			AuthInfo: v.AuthInfo,
			Cluster:  v.Cluster,
		}
	}

	// Initialize the context
	clusterConfig.CurrentContext = contextName
	clusterConfig.Contexts[contextName] = &Context{
		AuthInfo: kConfig.Contexts[contextName].AuthInfo,
		Cluster:  kConfig.Contexts[contextName].Cluster,
	}

	// Load all the pxc authentication information from the kubeconfig file
	for k, v := range kConfig.AuthInfos {
		if strings.HasPrefix(k, KubeconfigUserPrefix) && v.AuthProvider != nil {
			logrus.Debugf("Loading user %s from %s", k, v.LocationOfOrigin)
			pxcAuthInfo := NewAuthInfoFromMap(v.AuthProvider.Config)
			clusterConfig.AuthInfos[pxcAuthInfo.Name] = pxcAuthInfo
		} else if _, ok := clusterConfig.AuthInfos[k]; !ok {
			clusterConfig.AuthInfos[k] = NewAuthInfo()
			clusterConfig.AuthInfos[k].Name = k
		}
	}

	// Load all the pxc cluster information from the kubeconfig file
	for k, c := range kConfig.Clusters {
		if strings.HasPrefix(k, KubeconfigUserPrefix) {
			pxcClusterInfo, err := NewClusterFromEncodedString(string(c.CertificateAuthorityData))
			if err == nil {
				logrus.Debugf("Loading cluster %s from %s", k, c.LocationOfOrigin)
				clusterConfig.Clusters[pxcClusterInfo.Name] = pxcClusterInfo
			} else {
				logrus.Debugf("Unable to load cluster %s from %s", k, c.LocationOfOrigin)
			}
		} else if _, ok := clusterConfig.Clusters[k]; !ok {
			clusterConfig.Clusters[k] = NewDefaultCluster()
			clusterConfig.Clusters[k].Name = k
		}
	}

	return clusterConfig, nil
}

func (k *KubernetesConfigManager) ConfigSaveAuthInfo(authInfo *AuthInfo) error {
	if authInfo == nil {
		panic("authInfo required")
	}
	cc := k.ToRawKubeConfigLoader()
	save := false

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	currentContext := kconfig.Contexts[currentContextName]

	// Initialize authInfo object
	authInfo.Name = currentContext.AuthInfo

	// Check for token
	if len(authInfo.Token) != 0 {
		save = true
		// TODO: Validate if the token is expired
	}

	// Check for Kubernetes secret and secret namespace
	if len(authInfo.KubernetesAuthInfo.SecretNamespace) != 0 &&
		len(authInfo.KubernetesAuthInfo.SecretName) != 0 {
		save = true
	} else if len(authInfo.KubernetesAuthInfo.SecretNamespace) == 0 && len(authInfo.KubernetesAuthInfo.SecretName) != 0 {
		return fmt.Errorf("Must supply secret namespace with secret name")
	} else if len(authInfo.KubernetesAuthInfo.SecretNamespace) != 0 && len(authInfo.KubernetesAuthInfo.SecretName) == 0 {
		return fmt.Errorf("Must supply secret name with secret namespace")
	}

	// Check if any information necessary was passed
	if !save {
		return fmt.Errorf("Must supply authentication information")
	}

	// Get the location of the kubeconfig for this specific authInfo. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.AuthInfos[currentContext.AuthInfo].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := k.SaveAuthInfoForKubeUser(currentContext.AuthInfo, location, authInfo); err != nil {
		return err
	}

	logrus.Infof("Portworx login information saved in %s for Kubernetes user context %s\n",
		location,
		currentContext.AuthInfo)
	return nil
}

// ConfigSaveContext does not do anything in kubectl plugin mode because it is managed
// by kubectl
func (k *KubernetesConfigManager) ConfigSaveContext(c *Context) error {
	return fmt.Errorf("Use <kubectl config set-context> to set the context instead")
}

// ConfigDeleteAuthInfo deletes auth information for the current context
func (k *KubernetesConfigManager) ConfigDeleteAuthInfo(name string) error {
	cc := k.ToRawKubeConfigLoader()

	// This is the raw kubeconfig which may have been overridden by CLI args
	kconfig, err := cc.RawConfig()
	if err != nil {
		return err
	}

	// Get the current context
	currentContextName, err := k.GetKubernetesCurrentContext()
	if err != nil {
		return err
	}

	currentContext := kconfig.Contexts[currentContextName]

	// Get the location of the kubeconfig for this specific object. This is necessary
	// because KUBECONFIG can have many kubeconfigs, example: KUBECONFIG=kube1.conf:kube2.conf
	location := kconfig.AuthInfos[currentContext.AuthInfo].LocationOfOrigin

	// Storage the information to the appropriate kubeconfig
	if err := k.DeleteAuthInfoInKubeconfig(currentContext.AuthInfo); err != nil {
		return err
	}

	logrus.Infof("Portworx server information removed from %s for Kubernetes cluster %s\n",
		location,
		currentContext.Cluster)
	return nil
}

// ConfigDeleteContext deletes auth information for the current context
func (k *KubernetesConfigManager) ConfigDeleteContext(name string) error {
	return fmt.Errorf("Use kubectl config to manage context")
}

// ConfigUseContext is not supported by kubectl plugin
func (k *KubernetesConfigManager) ConfigUseContext(name string) error {
	return fmt.Errorf("Use kubectl to set the current context")
}

// ConfigGetCurrentContext returns the current context set by kubectl
func (k *KubernetesConfigManager) ConfigGetCurrentContext() (string, error) {
	return k.GetKubernetesCurrentContext()
}
