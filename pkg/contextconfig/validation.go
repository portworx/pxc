package contextconfig

/*
import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type claims struct {
	Subject string `json:"sub" yaml:"sub"`
	Name    string `json:"name" yaml:"name"`
	Email   string `json:"email" yaml:"email"`
}

// GetTokenClaims returns the claims for the raw JWT token.
func GetTokenClaims(rawtoken string) (*claims, error) {
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
	var claims *claims

	// Unmarshal claims
	err = json.Unmarshal(claimBytes, &claims)
	if err != nil {
		return nil, fmt.Errorf("Unable to get information from the claims in the token: %v", err)
	}

	return claims, nil
}

// AddClaimsInfo adds additional claims information to a contextconfig
func AddClaimsInfo(contextCfg *ContextConfig) *ContextConfig {
	for i := range contextCfg.Configurations {
		if contextCfg.Configurations[i].Token != "" {
			claims, err := GetTokenClaims(contextCfg.Configurations[i].Token)
			if err != nil {
				continue
			}
			contextCfg.Configurations[i].Identity.Subject = claims.Subject
			contextCfg.Configurations[i].Identity.Name = claims.Name
			contextCfg.Configurations[i].Identity.Email = claims.Email
		}
	}

	return contextCfg
}

// AddTokenValidity checks and marks if a token is invalid.
func AddTokenValidity(clientContext ClientContext) ClientContext {
	var mapClaims jwt.MapClaims
	_, _, err := new(jwt.Parser).
		ParseUnverified(clientContext.Token, &mapClaims)
	if err != nil {
		clientContext.Error = err.Error()
		clientContext.Name += " (token invalid)"
		return clientContext
	}

	if mapClaims.VerifyExpiresAt(time.Now().Unix(), false) == false {
		clientContext.Error = "Token is expired"
		clientContext.Name += " (token expired)"
	}

	return clientContext
}

// MarkInvalidTokens will mark all invalid tokens for a given context config.
func MarkInvalidTokens(contextCfg *ContextConfig) *ContextConfig {
	for i := range contextCfg.Configurations {
		if contextCfg.Configurations[i].Token != "" {
			contextCfg.Configurations[i] = AddTokenValidity(contextCfg.Configurations[i])
		}
	}

	return contextCfg
}
*/
