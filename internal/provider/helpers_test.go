// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func testProviderData(t *testing.T, raw map[string]interface{}) *pluginsdk.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, AzureADProvider().Schema, raw)
}

func testWriteTokenFile(t *testing.T, contents string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "azure-identity-token")
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("writing token file: %v", err)
	}
	return path
}

func TestGetOidcToken_fromAksWorkloadIdentity(t *testing.T) {
	path := testWriteTokenFile(t, "  aks-token\n")

	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", path)

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	idToken, err := getOidcToken(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *idToken != "aks-token" {
		t.Fatalf("expected %q, got %q", "aks-token", *idToken)
	}
}

func TestGetOidcToken_aksWorkloadIdentityDisabled(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", testWriteTokenFile(t, "aks-token"))

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": false,
	})

	idToken, err := getOidcToken(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *idToken != "" {
		t.Fatalf("expected empty token, got %q", *idToken)
	}
}

func TestGetOidcToken_aksWorkloadIdentityMismatch(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", testWriteTokenFile(t, "aks-token"))

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"oidc_token":                "some-other-token",
	})

	if _, err := getOidcToken(d); err == nil {
		t.Fatal("expected an error for mismatched OIDC token, got none")
	}
}

func TestGetOidcToken_aksWorkloadIdentityMissingFile(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", filepath.Join(t.TempDir(), "does-not-exist"))

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	if _, err := getOidcToken(d); err == nil {
		t.Fatal("expected an error for unreadable OIDC token file, got none")
	}
}

// TestAksWorkloadIdentity_satisfiesOidcAuthorizerPreconditions is a regression test for
// "no Authorizer could be configured". auth.NewAuthorizerFromCredentials only selects the
// OIDC authorizer when TenantID, ClientID and OIDCAssertionToken are all non-empty; before
// AZURE_FEDERATED_TOKEN_FILE and AZURE_CLIENT_ID were honoured, enabling
// use_aks_workload_identity on its own left all three unset and every authorizer was skipped.
func TestAksWorkloadIdentity_satisfiesOidcAuthorizerPreconditions(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", testWriteTokenFile(t, "aks-token"))
	t.Setenv("AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("AZURE_TENANT_ID", "11111111-1111-1111-1111-111111111111")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	idToken, err := getOidcToken(d)
	if err != nil {
		t.Fatalf("getOidcToken: %v", err)
	}
	clientId, err := getClientId(d)
	if err != nil {
		t.Fatalf("getClientId: %v", err)
	}
	tenantId, err := getTenantId(d)
	if err != nil {
		t.Fatalf("getTenantId: %v", err)
	}

	for name, value := range map[string]string{
		"OIDCAssertionToken": *idToken,
		"ClientID":           *clientId,
		"TenantID":           *tenantId,
	} {
		if strings.TrimSpace(value) == "" {
			t.Errorf("%s is empty - the OIDC authorizer would be skipped and authentication would fail with \"no Authorizer could be configured\"", name)
		}
	}
}

func TestGetOidcToken_aksWorkloadIdentityMatchesConfiguredToken(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", testWriteTokenFile(t, "aks-token"))

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"oidc_token":                "aks-token",
	})

	idToken, err := getOidcToken(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *idToken != "aks-token" {
		t.Fatalf("expected %q, got %q", "aks-token", *idToken)
	}
}

func TestGetOidcToken_aksWorkloadIdentityWithoutTokenFileEnv(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", "")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	if _, err := getOidcToken(d); err == nil {
		t.Fatal("expected an error when AZURE_FEDERATED_TOKEN_FILE is unset and no token is configured, got none")
	}
}

func TestGetOidcToken_aksWorkloadIdentityWithoutTokenFileEnvButConfiguredToken(t *testing.T) {
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", "")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"oidc_token":                "configured-token",
	})

	idToken, err := getOidcToken(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *idToken != "configured-token" {
		t.Fatalf("expected %q, got %q", "configured-token", *idToken)
	}
}

func TestGetClientId_fromAksWorkloadIdentity(t *testing.T) {
	t.Setenv("AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	clientId, err := getClientId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *clientId != "00000000-0000-0000-0000-000000000000" {
		t.Fatalf("unexpected client ID %q", *clientId)
	}
}

func TestGetClientId_aksWorkloadIdentityMismatch(t *testing.T) {
	t.Setenv("AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"client_id":                 "11111111-1111-1111-1111-111111111111",
	})

	if _, err := getClientId(d); err == nil {
		t.Fatal("expected an error for mismatched Client ID, got none")
	}
}

func TestGetClientId_aksWorkloadIdentityMatchesConfiguredClientId(t *testing.T) {
	t.Setenv("AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"client_id":                 "00000000-0000-0000-0000-000000000000",
	})

	clientId, err := getClientId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *clientId != "00000000-0000-0000-0000-000000000000" {
		t.Fatalf("unexpected client ID %q", *clientId)
	}
}

func TestGetClientId_aksWorkloadIdentityDisabled(t *testing.T) {
	t.Setenv("AZURE_CLIENT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": false,
	})

	clientId, err := getClientId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *clientId != "" {
		t.Fatalf("expected empty client ID, got %q", *clientId)
	}
}

func TestGetClientId_aksWorkloadIdentityEnvUnset(t *testing.T) {
	t.Setenv("AZURE_CLIENT_ID", "")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"client_id":                 "11111111-1111-1111-1111-111111111111",
	})

	clientId, err := getClientId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *clientId != "11111111-1111-1111-1111-111111111111" {
		t.Fatalf("unexpected client ID %q", *clientId)
	}
}

func TestGetTenantId_fromAksWorkloadIdentity(t *testing.T) {
	t.Setenv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
	})

	tenantId, err := getTenantId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *tenantId != "00000000-0000-0000-0000-000000000000" {
		t.Fatalf("unexpected tenant ID %q", *tenantId)
	}
}

func TestGetTenantId_aksWorkloadIdentityMismatch(t *testing.T) {
	t.Setenv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": true,
		"tenant_id":                 "11111111-1111-1111-1111-111111111111",
	})

	if _, err := getTenantId(d); err == nil {
		t.Fatal("expected an error for mismatched Tenant ID, got none")
	}
}

func TestGetTenantId_aksWorkloadIdentityDisabled(t *testing.T) {
	t.Setenv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")

	d := testProviderData(t, map[string]interface{}{
		"use_aks_workload_identity": false,
		"tenant_id":                 "11111111-1111-1111-1111-111111111111",
	})

	tenantId, err := getTenantId(d)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if *tenantId != "11111111-1111-1111-1111-111111111111" {
		t.Fatalf("unexpected tenant ID %q", *tenantId)
	}
}
