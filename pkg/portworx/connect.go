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

	"github.com/portworx/px/pkg/config"
	"github.com/portworx/px/pkg/contextconfig"
	pxgrpc "github.com/portworx/px/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// PxConnectDefault returns a Portworx client to the default or
// named context
func PxConnectDefault() (context.Context, *grpc.ClientConn, error) {
	// Global information will be set here, like forced context
	file := config.Get(config.File)
	context := config.Get(config.SpecifiedContext)
	if len(context) == 0 {
		return PxConnectCurrent(file)
	} else {
		return PxConnectNamed(file, context)
	}
}

// TODO: Add Support to connect to a context name

// PxConnect will connect to the default context server using TLS if needed
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

	// If user has provided valid CA cert, append to the existing system CA pool
	if len(pxctx.TlsData.Cacert) != 0 {
		// cannot set Insecure with TLS
		dialOptions, caerr = PxAppendCaCertcontext(pxctx)
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
	ctx := pxgrpc.AddMetadataToContext(context.Background(), "Authorization", "bearer "+pxctx.Token)

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

	// If user has provided valid CA cert, append to the existing system CA pool
	if len(pxctx.TlsData.Cacert) != 0 {
		// cannot set Insecure with TLS
		dialOptions, caerr = PxAppendCaCertcontext(pxctx)
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
	if pxctx.Token != "" {
		ctx = pxgrpc.AddMetadataToContext(ctx, "Authorization", "bearer "+pxctx.Token)
	}
	return ctx, conn, nil
}

// Append the provided valid CA from the user to the existing systemPool, used
// for authentication with the sdk server.
func PxAppendCaCertcontext(pxctx *contextconfig.ClientContext) ([]grpc.DialOption, error) {
	// Read the provided CA cert from the user
	capool, err := x509.SystemCertPool()
	if !capool.AppendCertsFromPEM([]byte(pxctx.TlsData.Cacert)) {
		return nil, err
	}

	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(
		credentials.NewClientTLSFromCert(capool, ""))}
	return dialOptions, nil
}
