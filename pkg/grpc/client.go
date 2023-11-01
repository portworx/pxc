/*
Copyright 2017 The Kubernetes Authors.
Copyright 2019 Portworx

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
package grpcserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Connect address by grpc
func Connect(address string, dialOptions []grpc.DialOption) (*grpc.ClientConn, error) {
	u, err := url.Parse(address)
	if err == nil && (!u.IsAbs() || u.Scheme == "unix") {
		dialOptions = append(dialOptions,
			grpc.WithDialer(
				func(addr string, timeout time.Duration) (net.Conn, error) {
					return net.DialTimeout("unix", u.Path, timeout)
				}))
	}

	cadata, err := config.KM().GetCurrentCA()
	if err != nil {
		logrus.WithError(err).Warnf("Could not get kubernetes CA certificate from current conext")
	} else if len(cadata) == 0 {
		logrus.Debugf("Adding grpc.WithInsecure to dial context")
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		var isSecure bool
		if err = util.VerifyConnection(address, &isSecure, cadata); err != nil {
			return nil, err
		}

		if isSecure {
			logrus.Debugf("Setting up TLS for grpc connection to %s", address)
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(
				credentials.NewTLS(&tls.Config{
					RootCAs: util.AppendCertPool(cadata),
				})))
		}
	}

	dialOptions = append(dialOptions, grpc.WithBackoffMaxDelay(time.Second))
	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		return nil, err
	}

	// We wait for 1 minute until conn.GetState() is READY.
	// The interval for this check is 1 second.
	if err := util.WaitFor(5*time.Second, 10*time.Millisecond, func() (bool, error) {
		if conn.GetState() == connectivity.Ready {
			return false, nil
		}
		return true, nil
	}); err != nil {
		// Clean up the connection
		if err := conn.Close(); err != nil {
			return nil, fmt.Errorf("Connection timed out and failed to close connection: %v", err)
		}
		return nil, fmt.Errorf("Connection timed out to server %s: %s", address, err)
	}

	return conn, nil
}

func AddMetadataToContext(ctx context.Context, k, v string) context.Context {
	md, _ := metadata.FromOutgoingContext(ctx)
	md = metadata.Join(md, metadata.New(map[string]string{
		k: v,
	}))
	return metadata.NewOutgoingContext(ctx, md)
}

func GetMetadataValueFromKey(ctx context.Context, k string) string {
	return metautils.ExtractIncoming(ctx).Get(k)
}
