// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/hashicorp/terraform-provider-azuread/version"
)

func TestUserAgent(t *testing.T) {
	baseUserAgent := fmt.Sprintf("Terraform/1.2.3 (+https://www.terraform.io) Terraform-Plugin-SDK/%s terraform-provider-azuread/%s", meta.SDKVersionString(), version.ProviderVersion) //nolint:staticcheck

	testCases := []struct {
		name         string
		options      ClientOptions
		sdkUserAgent string
		env          map[string]string
		expected     string
	}{
		{
			name:     "no additions",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			expected: baseUserAgent,
		},
		{
			name:         "sdk user agent",
			options:      ClientOptions{TerraformVersion: "1.2.3"},
			sdkUserAgent: "HashiCorp/go-azure-sdk (Go-http-client/1.1)",
			expected:     fmt.Sprintf("%s HashiCorp/go-azure-sdk (Go-http-client/1.1)", baseUserAgent),
		},
		{
			name:     "appended user agent",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			env:      map[string]string{"TF_APPEND_USER_AGENT": "my-custom-agent/1.0"},
			expected: fmt.Sprintf("%s my-custom-agent/1.0", baseUserAgent),
		},
		{
			name:     "appended user agent is trimmed",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			env:      map[string]string{"TF_APPEND_USER_AGENT": "  my-custom-agent/1.0  "},
			expected: fmt.Sprintf("%s my-custom-agent/1.0", baseUserAgent),
		},
		{
			name:     "empty appended user agent is ignored",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			env:      map[string]string{"TF_APPEND_USER_AGENT": "   "},
			expected: baseUserAgent,
		},
		{
			name:     "cloudshell agent",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			env:      map[string]string{"AZURE_HTTP_USER_AGENT": "cloud-shell/1.0"},
			expected: fmt.Sprintf("%s cloud-shell/1.0", baseUserAgent),
		},
		{
			name:     "empty cloudshell agent is ignored",
			options:  ClientOptions{TerraformVersion: "1.2.3"},
			env:      map[string]string{"AZURE_HTTP_USER_AGENT": ""},
			expected: baseUserAgent,
		},
		{
			name:         "cloudshell agent follows go-azure-sdk agent and precedes partner id",
			options:      ClientOptions{TerraformVersion: "1.2.3", PartnerID: "222b7c9b-8f2c-4b8c-9b3f-a1b2c3d4e5f6"},
			sdkUserAgent: "HashiCorp/go-azure-sdk (Go-http-client/1.1)",
			env:          map[string]string{"AZURE_HTTP_USER_AGENT": "cloud-shell/1.0"},
			expected:     fmt.Sprintf("%s HashiCorp/go-azure-sdk (Go-http-client/1.1) cloud-shell/1.0 pid-222b7c9b-8f2c-4b8c-9b3f-a1b2c3d4e5f6", baseUserAgent),
		},
		{
			name:         "appended user agent precedes go-azure-sdk agent, cloudshell agent and partner id",
			options:      ClientOptions{TerraformVersion: "1.2.3", PartnerID: "222b7c9b-8f2c-4b8c-9b3f-a1b2c3d4e5f6"},
			sdkUserAgent: "HashiCorp/go-azure-sdk (Go-http-client/1.1)",
			env: map[string]string{
				"TF_APPEND_USER_AGENT":  "my-custom-agent/1.0",
				"AZURE_HTTP_USER_AGENT": "cloud-shell/1.0",
			},
			expected: fmt.Sprintf("%s my-custom-agent/1.0 HashiCorp/go-azure-sdk (Go-http-client/1.1) cloud-shell/1.0 pid-222b7c9b-8f2c-4b8c-9b3f-a1b2c3d4e5f6", baseUserAgent),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Setenv("TF_APPEND_USER_AGENT", "")
			t.Setenv("AZURE_HTTP_USER_AGENT", "")
			for k, v := range testCase.env {
				t.Setenv(k, v)
			}

			if actual := testCase.options.userAgent(testCase.sdkUserAgent); actual != testCase.expected {
				t.Fatalf("expected user agent %q but got %q", testCase.expected, actual)
			}
		})
	}
}

func TestUserAgentAppendedValueIsNotDuplicated(t *testing.T) {
	t.Setenv("TF_APPEND_USER_AGENT", "my-custom-agent/1.0")

	o := ClientOptions{TerraformVersion: "1.2.3"}
	userAgent := o.userAgent("")

	if count := strings.Count(userAgent, "my-custom-agent/1.0"); count != 1 {
		t.Fatalf("expected appended user agent to appear once, but it appeared %d times in %q", count, userAgent)
	}
}
