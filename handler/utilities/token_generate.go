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
package utilities

import (
	"fmt"
	"strings"
	"time"

	"github.com/portworx/pxc/pkg/auth"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type tokenInfo struct {
	issuer  string
	subject string
	name    string
	email   string
	roles   string
	groups  string
}

type tokenGenOptions struct {
	sharedSecret string
	rsaPem       string
	ecdsaPem     string
	duration     string
	output       string
	token        tokenInfo
}

// tokenGenCmd represents the tokenGen command
var (
	tokenGenArgs *tokenGenOptions
	tokenGenCmd  *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	tokenGenArgs = &tokenGenOptions{}
	tokenGenCmd = &cobra.Command{
		Use:   "token-generate",
		Short: "Generate a Portworx token",
		Example: `
  # Login to portworx using a secret in Kubernetes
  pxc utilities token-generate pxc util token-generate \
	--token-email=example.user@example.com \
	--token-name="Example User" \
	--token-roles=system.user \
	--token-groups=exampleGroup \
	--token-duration=7d \
	--token-issuer=myissuer \
	--token-subject="exampleCompany/example.user@example.com" \
	--shared-secret=mysecret`,
		RunE: tokenGenExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	UtilitiesAddCommand(tokenGenCmd)
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.sharedSecret,
		"shared-secret", "", "Shared secret to sign token")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.rsaPem,
		"rsa-private-keyfile", "", "RSA Private file to sign token")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.ecdsaPem,
		"ecdsa-private-keyfile", "", "ECDSA Private file to sign token")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.duration,
		"token-duration", "1d", "Duration of time where the token will be valid. "+
			"Postfix the duration by using "+
			auth.SecondDef+" for seconds, "+
			auth.MinuteDef+" for minutes, "+
			auth.HourDef+" for hours, "+
			auth.DayDef+" for days, and "+
			auth.YearDef+" for years.")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.issuer,
		"token-issuer", "portworx.com",
		"Issuer name of token. Do not use https:// in the issuer since it could indicate "+
			"that this is an OpenID Connect issuer.")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.name,
		"token-name", "", "Account name")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.subject,
		"token-subject", "", "Unique ID of this account")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.email,
		"token-email", "", "Unique ID of this account")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.roles,
		"token-roles", "", "Comma separated list of roles applied to this token")
	tokenGenCmd.Flags().StringVar(&tokenGenArgs.token.groups,
		"token-groups", "", "Comma separated list of groups which the token will be part of")

})

func TokenGenerateAddCommand(cmd *cobra.Command) {
	tokenGenCmd.AddCommand(cmd)
}

func tokenGenExec(cmd *cobra.Command, args []string) error {

	if len(tokenGenArgs.token.name) == 0 {
		return fmt.Errorf("Must supply an account name")
	} else if len(tokenGenArgs.token.email) == 0 {
		return fmt.Errorf("Must supply an email address")
	} else if len(tokenGenArgs.token.subject) == 0 {
		return fmt.Errorf("Must supply a unique identifier as the subject")
	}
	if len(tokenGenArgs.token.roles) == 0 {
		logrus.Warningf("Warning: No role provided")
	}
	if len(tokenGenArgs.token.groups) == 0 {
		logrus.Warningf("Warning: No role provided")
	}

	claims := &auth.Claims{
		Name:    tokenGenArgs.token.name,
		Email:   tokenGenArgs.token.email,
		Subject: tokenGenArgs.token.subject,
		Roles:   strings.Split(tokenGenArgs.token.roles, ","),
		Groups:  strings.Split(tokenGenArgs.token.groups, ","),
	}

	// Get duration
	options := &auth.Options{
		Issuer: tokenGenArgs.token.issuer,
	}
	expDuration, err := auth.ParseToDuration(tokenGenArgs.duration)
	if err != nil {
		return fmt.Errorf("Unable to parse duration")
	}
	options.Expiration = time.Now().Add(expDuration).Unix()

	// Get signature
	var signature *auth.Signature
	if len(tokenGenArgs.sharedSecret) != 0 {
		signature, err = auth.NewSignatureSharedSecret(tokenGenArgs.sharedSecret)
	} else if len(tokenGenArgs.rsaPem) != 0 {
		signature, err = auth.NewSignatureRSAFromFile(tokenGenArgs.rsaPem)
	} else if len(tokenGenArgs.ecdsaPem) != 0 {
		signature, err = auth.NewSignatureECDSAFromFile(tokenGenArgs.ecdsaPem)
	} else {
		return fmt.Errorf("Must provide a secret key to sign token")
	}
	if err != nil {
		return fmt.Errorf("Unable to generate signature: %v", err)
	}

	// Generate token
	token, err := auth.Token(claims, signature, options)
	if err != nil {
		return fmt.Errorf("Failed to create token: %v", err)
	}

	// Print token
	util.Printf("%s\n", token)

	return nil
}
