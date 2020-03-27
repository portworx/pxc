/*
Copyright 2018 Portworx

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
package auth

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims provides information about the claims in the token
// See https://openid.net/specs/openid-connect-core-1_0.html#IDToken
// for more information.
type Claims struct {
	// Issuer is the token issuer. For selfsigned token do not prefix
	// with `https://`.
	Issuer string `json:"iss"`
	// Subject identifier. Unique ID of this account
	Subject string `json:"sub" yaml:"sub"`
	// Account name
	Name string `json:"name" yaml:"name"`
	// Account email
	Email string `json:"email" yaml:"email"`
	// Roles of this account
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`
	// (optional) Groups in which this account is part of
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`
}

// Options provide any options to apply to the token
type Options struct {
	// Expiration time in Unix format as per JWT standard
	Expiration int64
	// Issuer of the claims
	Issuer string
}

// Token returns a signed JWT containing the claims provided
func Token(
	claims *Claims,
	signature *Signature,
	options *Options,
) (string, error) {

	mapclaims := jwt.MapClaims{
		"sub":   claims.Subject,
		"iss":   options.Issuer,
		"email": claims.Email,
		"name":  claims.Name,
		"roles": claims.Roles,
		"iat":   time.Now().Unix(),
		"exp":   options.Expiration,
	}
	if claims.Groups != nil {
		mapclaims["groups"] = claims.Groups
	}
	token := jwt.NewWithClaims(signature.Type, mapclaims)
	signedtoken, err := token.SignedString(signature.Key)
	if err != nil {
		return "", err
	}

	return signedtoken, nil
}

// TokenClaims returns the claims for the raw JWT token.
func TokenClaims(rawtoken string) (*Claims, error) {
	parts := strings.Split(rawtoken, ".")

	// There are supposed to be three parts for the token
	if len(parts) < 3 {
		return nil, fmt.Errorf("Token is invalid: %v", rawtoken)
	}

	// Access claims in the token
	claimBytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Failed to decode claims: %v", err)
	}
	var claims *Claims

	// Unmarshal claims
	err = json.Unmarshal(claimBytes, &claims)
	if err != nil {
		return nil, fmt.Errorf("Unable to get information from the claims in the token: %v", err)
	}

	return claims, nil
}

// TokenIssuer returns the issuer for the raw JWT token.
func TokenIssuer(rawtoken string) (string, error) {
	claims, err := TokenClaims(rawtoken)
	if err != nil {
		return "", err
	}

	// Return issuer
	if len(claims.Issuer) != 0 {
		return claims.Issuer, nil
	} else {
		return "", fmt.Errorf("Issuer was not specified in the token")
	}
}

// IsJwtToken returns true if the provided string is a valid jwt token
func IsJwtToken(authstring string) bool {
	_, _, err := new(jwt.Parser).ParseUnverified(authstring, jwt.MapClaims{})
	return err == nil
}

func ValidateToken(rawtoken string) error {
	var mapClaims jwt.MapClaims
	_, _, err := new(jwt.Parser).
		ParseUnverified(rawtoken, &mapClaims)
	if err != nil {
		return err
	}
	return mapClaims.Valid()
}

// This function is similar to jwt-go's (m MapClaims) VerifyExpiresAt
// Copyright (c) 2012 Dave Grijalva
func GetExpiration(rawtoken string) (time.Time, error) {
	var mapClaims jwt.MapClaims
	_, _, err := new(jwt.Parser).
		ParseUnverified(rawtoken, &mapClaims)
	if err != nil {
		return time.Time{}, err
	}

	switch exp := mapClaims["exp"].(type) {
	case float64:
		return time.Unix(int64(exp), 0), nil
	case json.Number:
		v, _ := exp.Int64()
		return time.Unix(v, 0), nil
	}
	return time.Time{}, fmt.Errorf("Unable to get expiration time from token")
}

// This function is similar to jwt-go's (m MapClaims) VerifyIssedAt
// Copyright (c) 2012 Dave Grijalva
func GetIssuedAtTime(rawtoken string) (time.Time, error) {
	var mapClaims jwt.MapClaims
	_, _, err := new(jwt.Parser).
		ParseUnverified(rawtoken, &mapClaims)
	if err != nil {
		return time.Time{}, err
	}

	switch iat := mapClaims["iat"].(type) {
	case float64:
		return time.Unix(int64(iat), 0), nil
	case json.Number:
		v, _ := iat.Int64()
		return time.Unix(v, 0), nil
	}
	return time.Time{}, fmt.Errorf("Unable to get expiration time from token")
}
