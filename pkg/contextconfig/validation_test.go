package contextconfig_test

/*
import (
	"strings"
	"testing"

	"github.com/portworx/pxc/pkg/contextconfig"
)

func TestAddClaimsInfo(t *testing.T) {
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppbXN0ZXZlbnNAZ21haWwuY29tIiwiZXhwIjo0NzE3NTk4NTI4LCJncm91cHMiOlsiKiJdLCJpYXQiOjE1NjM5OTg1MjgsImlzcyI6Im9wZW5zdG9yYWdlLmlvIiwibmFtZSI6ImppbSIsInJvbGVzIjpbInN5c3RlbS5hZG1pbiJdLCJzdWIiOiJqaW1zdGV2ZW5zQGdtYWlsLmNvbSJ9.99zV2XyuY53_2Nf5qkpM7nGNFzu8M3Y0jQaxPWkSozA"

	validContext := &contextconfig.ContextConfig{
		Current: "validctx",
		Configurations: []contextconfig.ClientContext{
			contextconfig.ClientContext{
				Name:  "validctx",
				Token: validToken,
			},
		},
	}
	cfgWithClaimsInfo := contextconfig.AddClaimsInfo(validContext)

	if cfgWithClaimsInfo.Configurations[0].Identity.Name != "jim" {
		t.Errorf("got %s, expected Grant", cfgWithClaimsInfo.Configurations[0].Identity.Name)
	}
	if cfgWithClaimsInfo.Configurations[0].Identity.Email != "jimstevens@gmail.com" {
		t.Errorf("got %s, expected grant@portworx.com", cfgWithClaimsInfo.Configurations[0].Identity.Email)
	}
	if cfgWithClaimsInfo.Configurations[0].Identity.Subject != "jimstevens@gmail.com" {
		t.Errorf("got %s, expected grant@portworx.com", cfgWithClaimsInfo.Configurations[0].Identity.Subject)
	}
}

func TestMarkInvalidTokens(t *testing.T) {
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppbXN0ZXZlbnNAZ21haWwuY29tIiwiZXhwIjoxNTYzOTk4NDg0LCJncm91cHMiOlsiKiJdLCJpYXQiOjE1NjM5OTg0ODMsImlzcyI6Im9wZW5zdG9yYWdlLmlvIiwibmFtZSI6ImppbSIsInJvbGVzIjpbInN5c3RlbS5hZG1pbiJdLCJzdWIiOiJqaW1zdGV2ZW5zQGdtYWlsLmNvbSJ9.rYvraLijzV-PdCfr1SJqHQ1T-ne-AK9032NK7mQ2srw"
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImppbXN0ZXZlbnNAZ21haWwuY29tIiwiZXhwIjo0NzE3NTk4NTI4LCJncm91cHMiOlsiKiJdLCJpYXQiOjE1NjM5OTg1MjgsImlzcyI6Im9wZW5zdG9yYWdlLmlvIiwibmFtZSI6ImppbSIsInJvbGVzIjpbInN5c3RlbS5hZG1pbiJdLCJzdWIiOiJqaW1zdGV2ZW5zQGdtYWlsLmNvbSJ9.99zV2XyuY53_2Nf5qkpM7nGNFzu8M3Y0jQaxPWkSozA"
	invalidToken := "abcd"
	validContext := &contextconfig.ContextConfig{
		Current: "validctx",
		Configurations: []contextconfig.ClientContext{
			contextconfig.ClientContext{
				Name:  "validctx",
				Token: validToken,
			},
		},
	}
	invalidContext := &contextconfig.ContextConfig{
		Current: "invalidctx",
		Configurations: []contextconfig.ClientContext{
			contextconfig.ClientContext{
				Name:  "invalidctx",
				Token: invalidToken,
			},
		},
	}
	expiredContext := &contextconfig.ContextConfig{
		Current: "expiredctx",
		Configurations: []contextconfig.ClientContext{
			contextconfig.ClientContext{
				Name:  "expiredctx",
				Token: expiredToken,
			},
		},
	}

	tt := []struct {
		context            *contextconfig.ContextConfig
		expectError        bool
		expectedNameSuffix string
	}{
		{
			context:     validContext,
			expectError: false,
		},
		{
			context:            invalidContext,
			expectError:        true,
			expectedNameSuffix: "(token invalid)",
		},
		{
			context:            expiredContext,
			expectError:        true,
			expectedNameSuffix: "(token expired)",
		},
	}

	for _, tc := range tt {
		oldName := tc.context.Configurations[0].Name
		tc.context = contextconfig.MarkInvalidTokens(tc.context)

		if tc.expectError {
			if !strings.Contains(tc.context.Configurations[0].Name, tc.expectedNameSuffix) {
				t.Errorf("got %s; should contain %s", tc.context.Configurations[0].Name, tc.expectedNameSuffix)
			}
		} else {
			if oldName != tc.context.Configurations[0].Name {
				t.Errorf("original valid context name should not chage. old: %s, validated: %s", oldName, tc.context.Configurations[0].Name)
			}
		}
	}
}
*/
