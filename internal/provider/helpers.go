// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

// logEntry avoids log entries showing up in test output
func logEntry(f string, v ...interface{}) {
	if os.Getenv("TF_LOG") == "" {
		return
	}

	if os.Getenv("TF_ACC") != "" {
		return
	}

	log.Printf(f, v...)
}

func decodeCertificate(clientCertificate string) ([]byte, error) {
	var pfx []byte
	if clientCertificate != "" {
		out := make([]byte, base64.StdEncoding.DecodedLen(len(clientCertificate)))
		n, err := base64.StdEncoding.Decode(out, []byte(clientCertificate))
		if err != nil {
			return pfx, fmt.Errorf("could not decode client certificate data: %v", err)
		}
		pfx = out[:n]
	}
	return pfx, nil
}

// readTrimmedFile reads the file at path and returns its contents with surrounding
// whitespace removed.
func readTrimmedFile(path string) (string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(contents)), nil
}

// reconcile returns candidate when existing is empty or identical, and an error
// describing the conflict otherwise. name is the human readable name of the value,
// e.g. "Client ID", and source describes where candidate came from.
func reconcile(name, existing, candidate, source string) (string, error) {
	if existing != "" && existing != candidate {
		return "", fmt.Errorf("mismatch between supplied %[1]s and %[1]s %[2]s - please either remove one or ensure they match", name, source)
	}

	return candidate, nil
}

// reconcileAksWorkloadIdentityEnv reconciles existing with the value of the environment
// variable envVar, which is projected into the pod by the AKS Workload Identity mutating
// admission webhook. When the environment variable is unset, existing is returned as-is.
func reconcileAksWorkloadIdentityEnv(d *pluginsdk.ResourceData, name, existing, envVar string) (string, error) {
	if !d.Get("use_aks_workload_identity").(bool) {
		return existing, nil
	}

	candidate := strings.TrimSpace(os.Getenv(envVar))
	if candidate == "" {
		return existing, nil
	}

	return reconcile(name, existing, candidate, fmt.Sprintf("provided by AKS Workload Identity in %s", envVar))
}

func getOidcToken(d *pluginsdk.ResourceData) (*string, error) {
	idToken := d.Get("oidc_token").(string)

	if path := d.Get("oidc_token_file_path").(string); path != "" {
		fileToken, err := readTrimmedFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading OIDC Token from file %q: %v", path, err)
		}

		if idToken, err = reconcile("OIDC token", idToken, fileToken, fmt.Sprintf("read from the file %q", path)); err != nil {
			return nil, err
		}
	}

	if d.Get("use_aks_workload_identity").(bool) {
		path := strings.TrimSpace(os.Getenv("AZURE_FEDERATED_TOKEN_FILE"))
		if path == "" && idToken == "" {
			return nil, fmt.Errorf("no OIDC token was found: use_aks_workload_identity is enabled but AZURE_FEDERATED_TOKEN_FILE is not set - please ensure the pod is labelled with `azure.workload.identity/use: \"true\"`, or configure oidc_token or oidc_token_file_path")
		}

		if path != "" {
			fileToken, err := readTrimmedFile(path)
			if err != nil {
				return nil, fmt.Errorf("reading OIDC Token from file %q provided by AKS Workload Identity: %v", path, err)
			}

			if idToken, err = reconcile("OIDC token", idToken, fileToken, fmt.Sprintf("read from the file %q provided by AKS Workload Identity in AZURE_FEDERATED_TOKEN_FILE", path)); err != nil {
				return nil, err
			}
		}
	}

	return &idToken, nil
}

func getClientId(d *pluginsdk.ResourceData) (*string, error) {
	clientId := strings.TrimSpace(d.Get("client_id").(string))

	if path := d.Get("client_id_file_path").(string); path != "" {
		fileClientId, err := readTrimmedFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading Client ID from file %q: %v", path, err)
		}

		if clientId, err = reconcile("Client ID", clientId, fileClientId, fmt.Sprintf("read from the file %q", path)); err != nil {
			return nil, err
		}
	}

	clientId, err := reconcileAksWorkloadIdentityEnv(d, "Client ID", clientId, "AZURE_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	return &clientId, nil
}

func getClientSecret(d *pluginsdk.ResourceData) (*string, error) {
	clientSecret := strings.TrimSpace(d.Get("client_secret").(string))

	if path := d.Get("client_secret_file_path").(string); path != "" {
		fileSecret, err := readTrimmedFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading Client Secret from file %q: %v", path, err)
		}

		if clientSecret, err = reconcile("Client Secret", clientSecret, fileSecret, fmt.Sprintf("read from the file %q", path)); err != nil {
			return nil, err
		}
	}

	return &clientSecret, nil
}

func getTenantId(d *pluginsdk.ResourceData) (*string, error) {
	tenantId := strings.TrimSpace(d.Get("tenant_id").(string))

	tenantId, err := reconcileAksWorkloadIdentityEnv(d, "Tenant ID", tenantId, "AZURE_TENANT_ID")
	if err != nil {
		return nil, err
	}

	return &tenantId, nil
}
