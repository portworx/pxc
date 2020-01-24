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
package portworx

import (
	"context"
	"crypto/x509"
	"fmt"

	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/contextconfig"
	pxgrpc "github.com/portworx/pxc/pkg/grpc"
	"github.com/portworx/pxc/pkg/kubernetes"
	"github.com/portworx/pxc/pkg/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PxConnectDefault returns a Portworx client to the default or
// named context
func PxConnectDefault() (context.Context, *grpc.ClientConn, error) {

	if util.InKubectlPluginMode() {
		return PxConnectAsPlugin()
	}

	// Global information will be set here, like forced context
	file := config.Get(config.File)
	context := config.Get(config.SpecifiedContext)
	if len(context) == 0 {
		return PxConnectCurrent(file)
	} else {
		return PxConnectNamed(file, context)
	}
}

// PxConnectAsPlugin expects that the portforwarder has been setup and uses
// a local port to communicate with the gRPC port in Portworx through the
// Kubernetes API.
func PxConnectAsPlugin() (context.Context, *grpc.ClientConn, error) {

	var (
		dialOptions []grpc.DialOption
	)

	// If secure: true set in config.yaml file, use TLS
	dialOptions = append(dialOptions, grpc.WithInsecure())

	// Get config
	endpoint := config.CM().GetEndpoint()
	authInfo := config.CM().GetCurrentAuthInfo()

	// Connect to server
	conn, err := pxgrpc.Connect(endpoint, dialOptions)
	if err != nil {
		return nil, nil, err
	}

	// Check if the token is in a secret
	token := authInfo.Token
	if len(authInfo.KubernetesAuthInfo.SecretName) != 0 &&
		len(authInfo.KubernetesAuthInfo.SecretNamespace) != 0 {
		token, err = PxGetTokenFromSecret(authInfo.KubernetesAuthInfo.SecretName, authInfo.KubernetesAuthInfo.SecretNamespace)
		if err != nil {
			return nil, nil, err
		}
	}

	ctx := context.Background()
	if len(token) != 0 {
		ctx = pxgrpc.AddMetadataToContext(ctx, "authorization", "bearer "+token)
	}

	logrus.Infof("Connected through API server to %s\n", endpoint)
	return ctx, conn, nil
}

// TODO: Add Support to connect to a context name

// PxConnectCurrent will connect to the default context server using TLS if needed
// and returns the context setup with any security if any and the grpc client.
// The context will not have a timeout set, that should be setup by the caller
// of the gRPC call.
func PxConnectCurrent(cfgFile string) (context.Context, *grpc.ClientConn, error) {
	contextManager, err := contextconfig.NewContextManager(cfgFile)
	if err != nil {
		return nil, nil, err
	}
	pxctx, err := contextManager.GetCurrent()
	if err != nil {
		return nil, nil, err
	}

	var (
		dialOptions []grpc.DialOption
		caerr       error
	)

	// If secure: true set in config.yaml file, use TLS
	if pxctx.Secure {
		// cannot set Insecure with TLS.
		if len(pxctx.TlsData.Cacert) != 0 {
			// If user has provided valid CA cert, append to the existing system CA pool.
			// Parameter "true" signifies user provided CA.
			dialOptions, caerr = PxAppendCaCertcontext(pxctx, true)
		} else {
			// Parameter "false" signifies load available CA from the system.
			dialOptions, caerr = PxAppendCaCertcontext(pxctx, false)
		}
		if caerr != nil {
			return nil, nil, caerr
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	conn, err := pxgrpc.Connect(pxctx.Endpoint, dialOptions)
	if err != nil {
		return nil, nil, err
	}

	// Add authentication metadata
	ctx := context.Background()
	if len(pxctx.Token) != 0 {
		ctx = pxgrpc.AddMetadataToContext(ctx, "authorization", "bearer "+pxctx.Token)
	}

	return ctx, conn, nil
}

// PxConnectNamed will connect to a specified context server using TLS if needed
// and returns the context setup with any security if any and the grpc client.
// The context will not have a timeout set, that should be setup by the caller
// of the gRPC call.
func PxConnectNamed(cfgFile string, name string) (context.Context, *grpc.ClientConn, error) {
	contextManager, err := contextconfig.NewContextManager(cfgFile)
	if err != nil {
		return nil, nil, err
	}
	pxctx, err := contextManager.GetNamedContext(name)
	if err != nil {
		return nil, nil, err
	}
	var (
		dialOptions []grpc.DialOption
		caerr       error
	)

	// If secure: true set in config.yaml file, use TLS
	if pxctx.Secure {
		// cannot set Insecure with TLS.
		if len(pxctx.TlsData.Cacert) != 0 {
			// If user has provided valid CA cert, append to the existing system CA pool.
			// Parameter "true" signifies user provided CA.
			dialOptions, caerr = PxAppendCaCertcontext(pxctx, true)
		} else {
			// Parameter "false" signifies load available CA from the system.
			dialOptions, caerr = PxAppendCaCertcontext(pxctx, false)
		}
		if caerr != nil {
			return nil, nil, caerr
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	conn, err := pxgrpc.Connect(pxctx.Endpoint, dialOptions)
	if err != nil {
		return nil, nil, err
	}

	// Add authentication metadata
	ctx := context.Background()
	if len(pxctx.Token) != 0 {
		ctx = pxgrpc.AddMetadataToContext(ctx, "Authorization", "bearer "+pxctx.Token)
	}
	return ctx, conn, nil
}

// PxAppendCaCertcontext appends the provided valid CA from the user to the existing systemPool or
// load the default CA certs used for authentication with the sdk server.
func PxAppendCaCertcontext(pxctx *contextconfig.ClientContext, userCa bool) ([]grpc.DialOption, error) {
	// Read the provided CA cert from the user
	capool, err := x509.SystemCertPool()
	// If user provided CA cert, then append it to systemCertPool.
	if userCa {
		if !capool.AppendCertsFromPEM([]byte(pxctx.TlsData.Cacert)) {
			return nil, err
		}
	}

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(
		credentials.NewClientTLSFromCert(capool, ""))}
	return dialOptions, nil
}

func PxGetTokenFromSecret(secretName, secretNamespace string) (string, error) {
	_, clientSet, err := kubernetes.KubeConnectDefault()
	if err != nil {
		logrus.Errorf("Failed to get kube client: %v", err)
		return "", err
	}

	secretsClient := clientSet.CoreV1().Secrets(secretNamespace)
	secret, err := secretsClient.Get(secretName, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Failed to fetch secret: %v", err)
		return "", err
	}

	var (
		tokenRaw []byte
		ok       bool
	)
	if tokenRaw, ok = secret.Data["auth-token"]; !ok {
		return "", fmt.Errorf("Token not found in secret. Token is expected to be under 'auth-token' in the secret")
	}
	logrus.Infof("TokenRaw retrieved from secret %s/%s", secretNamespace, secretName)

	return string(tokenRaw), nil
}
