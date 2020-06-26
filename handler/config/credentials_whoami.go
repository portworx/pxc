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
package configcli

import (
	"fmt"
	"strings"
	"time"

	"github.com/portworx/pxc/pkg/auth"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/config"
	"github.com/portworx/pxc/pkg/portworx"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

type whoamiOptions struct {
	token string
}

var (
	whoamiArgs *whoamiOptions
	whoamiCmd  *cobra.Command
)

var _ = commander.RegisterCommandVar(func() {
	whoamiArgs = &whoamiOptions{}
	whoamiCmd = &cobra.Command{
		Use:   "whoami",
		Short: "Shows current authentication information",
		RunE:  whoAmIExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	CredentialsAddCommand(whoamiCmd)
	whoamiCmd.Flags().StringVar(&whoamiArgs.token,
		"auth-token", "", "Use this token instead of the token saved in the configuration. Useful for debugging tokens")
})

func WhoAmIAddCommand(cmd *cobra.Command) {
	whoamiCmd.AddCommand(cmd)
}

func whoAmIExec(cmd *cobra.Command, args []string) error {
	token := whoamiArgs.token
	if len(token) == 0 {
		authInfo := config.CM().GetCurrentAuthInfo()
		token = authInfo.Token

		if len(authInfo.KubernetesAuthInfo.SecretName) != 0 &&
			len(authInfo.KubernetesAuthInfo.SecretNamespace) != 0 {
			var err error
			token, err = portworx.PxGetTokenFromSecret(authInfo.KubernetesAuthInfo.SecretName, authInfo.KubernetesAuthInfo.SecretNamespace)
			if err != nil {
				return fmt.Errorf("Unable to retreive token from Kubernetes: %v", err)
			}
		}
		if len(token) == 0 {
			util.Printf("No authentication information provided\n")
			return nil
		}
	}

	expTime, err := auth.GetExpiration(token)
	if err != nil {
		return fmt.Errorf("Unable to get expiration information from token: %v", err)
	}
	iatTime, err := auth.GetIssuedAtTime(token)
	if err != nil {
		return fmt.Errorf("Unable to get issued time information from token: %v", err)
	}
	claims, err := auth.TokenClaims(token)
	if err != nil {
		return err
	}

	status := "Ok"
	err = auth.ValidateToken(token)
	if err != nil {
		status = fmt.Sprintf("%v", err)
	}

	util.Printf("Name: %s\n"+
		"Email: %s\n"+
		"Subject: %s\n"+
		"Groups: %s\n"+
		"Roles: %s\n"+
		"Issued At Time: %s\n"+
		"Expiration Time: %s\n"+
		"\n"+
		"Status: %s\n",
		claims.Name,
		claims.Email,
		claims.Subject,
		strings.Join(claims.Groups, ","),
		strings.Join(claims.Roles, ","),
		iatTime.Format(time.UnixDate),
		expTime.Format(time.UnixDate),
		status)

	return nil

}
