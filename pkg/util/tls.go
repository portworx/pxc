package util

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

// AppendCertPool returns system CA pool, with optional addition of <cadata> CA certificate
func AppendCertPool(cadata []byte) *x509.CertPool {
	capool, err := x509.SystemCertPool()
	if err != nil {
		logrus.WithError(err).Warnf("Could not get the system CA pool")
		capool = x509.NewCertPool()
	}

	if len(cadata) != 0 {
		if !capool.AppendCertsFromPEM(cadata) {
			logrus.Warn("could not append custom certificate")
		}
	}

	return capool
}

// VerifyConnection will check the connection (host:port endpoint), including the SSL validation
func VerifyConnection(endpoint string, isSecure *bool, cabytes []byte) error {
	ep := endpoint
	if ep == "" {
		ep = "127.0.0.1:9020"
	}
	ephost, _, err := net.SplitHostPort(ep)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %s", err)
	}
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	var capool *x509.CertPool
	if len(cabytes) > 0 {
		logrus.Debug("Appending custom CA certificate")
		capool = AppendCertPool(cabytes)
		conf = &tls.Config{
			RootCAs: capool,
		}
	}

	logrus.Debugf("Connecting to %s", ep)
	conn, err := tls.Dial("tcp", ep, conf)
	if err != nil {
		return fmt.Errorf("error connecting to %s: %s", ep, err)
	}
	defer conn.Close()

	// verify certs
	if certs := conn.ConnectionState().PeerCertificates; len(certs) > 0 {
		logrus.Debug("Checking TLS certificate")
		_, err = conn.ConnectionState().PeerCertificates[0].Verify(x509.VerifyOptions{
			DNSName:   ephost,
			Roots:     capool,
			KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		})
		if err != nil {
			return fmt.Errorf("TLS validation failed for %q: %s", ep, err)
		}
		if isSecure != nil {
			*isSecure = true
		}
		logrus.Debug("Got the following TLS certificates from the endpoint:")
		for _, c := range certs {
			logrus.Debugf("> %s   (isCA:%v, issud by %s)", c.Subject.String(), c.IsCA, c.Issuer.String())
		}
	}
	logrus.Debugf("Endpoint %s successfully validated", ep)
	return nil
}
