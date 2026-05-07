// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package customsecurityattributes_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	sdkClient "github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

// ── Test resource struct ──────────────────────────────────────────────────────

type CustomSecurityAttributeAssignmentResource struct{}

// ── Acceptance tests ──────────────────────────────────────────────────────────

// TestAccCustomSecurityAttributeAssignment_singleStringValue creates a resource
// with a single-string attribute and verifies it is readable.
func TestAccCustomSecurityAttributeAssignment_singleStringValue(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	// Custom security attributes cannot be truly deleted via the API (only nulled),
	// so we skip the destroy check.
	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.singleStringValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("attribute_set").HasValue("TestAttributeSet"),
				check.That(data.ResourceName).Key("attribute.#").HasValue("1"),
			),
		},
		// Import by <principal_id>/<attribute_set>
		data.ImportStep(),
	})
}

// TestAccCustomSecurityAttributeAssignment_multiStringValues creates a resource
// with a multi-value string collection attribute.
func TestAccCustomSecurityAttributeAssignment_multiStringValues(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.multiStringValues(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("attribute.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

// TestAccCustomSecurityAttributeAssignment_booleanValue creates a resource
// with a boolean attribute.
func TestAccCustomSecurityAttributeAssignment_booleanValue(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.booleanValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("attribute.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

// TestAccCustomSecurityAttributeAssignment_multipleAttributes creates a resource
// with multiple attribute blocks of different types in the same attribute set.
func TestAccCustomSecurityAttributeAssignment_multipleAttributes(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.multipleAttributes(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("attribute.#").HasValue("3"),
			),
		},
		data.ImportStep(),
	})
}

// TestAccCustomSecurityAttributeAssignment_update verifies that updating attribute
// values in-place works without recreating the resource.
func TestAccCustomSecurityAttributeAssignment_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.singleStringValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.singleStringValueUpdated(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

// TestAccCustomSecurityAttributeAssignment_requiresImport verifies that attempting
// to create a duplicate resource produces the expected error.
func TestAccCustomSecurityAttributeAssignment_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_security_attribute_assignment", "test")
	r := CustomSecurityAttributeAssignmentResource{}

	data.ResourceTestSkipCheckDestroyed(t, []acceptance.TestStep{
		{
			Config: r.singleStringValue(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

// ── Exists helper ─────────────────────────────────────────────────────────────

// Exists checks that the custom security attribute assignment exists in Azure by
// reading the customSecurityAttributes property from the service principal.
func (r CustomSecurityAttributeAssignmentResource) Exists(ctx context.Context, c *clients.Client, state *terraform.InstanceState) (*bool, error) {
	// Parse the composite ID: <principalId>/<attributeSet>
	id := state.ID
	var principalId, attributeSet string
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == '/' {
			principalId = id[:i]
			attributeSet = id[i+1:]
			break
		}
	}
	if principalId == "" || attributeSet == "" {
		return nil, fmt.Errorf("could not parse resource ID %q: expected <principalId>/<attributeSet>", id)
	}

	spId := beta.NewServicePrincipalID(principalId)
	opts := serviceprincipalBeta.GetServicePrincipalOperationOptions{
		Select: pointer.To([]string{"id", "customSecurityAttributes"}),
	}

	resp, err := c.CustomSecurityAttributes.ServicePrincipalClientBeta.GetServicePrincipal(ctx, spId, opts)
	if err != nil {
		return pointer.To(false), fmt.Errorf("retrieving service principal %q: %+v", principalId, err)
	}
	if resp.Model == nil {
		return pointer.To(false), nil
	}

	// Read customSecurityAttributes via raw JSON since the SDK type is a stub
	csaClient := c.CustomSecurityAttributes.ServicePrincipalClientBeta
	req, err := csaClient.Client.NewRequest(ctx, sdkClient.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                fmt.Sprintf("/servicePrincipals/%s", principalId),
		OptionsObject: &selectOpts{
			fields: []string{"id", "customSecurityAttributes"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("building GET request: %+v", err)
	}

	httpResp, err := req.Execute(ctx)
	if err != nil {
		return pointer.To(false), nil
	}

	var body map[string]interface{}
	if err := httpResp.Unmarshal(&body); err != nil {
		return nil, fmt.Errorf("decoding response: %+v", err)
	}

	csaRaw, ok := body["customSecurityAttributes"].(map[string]interface{})
	if !ok {
		return pointer.To(false), nil
	}

	setVal, ok := csaRaw[attributeSet].(map[string]interface{})
	if !ok {
		return pointer.To(false), nil
	}

	// The attribute set exists and has at least one non-metadata key
	for k := range setVal {
		if k != "@odata.type" {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

// selectOpts implements sdkClient.Options to append a $select query parameter.
type selectOpts struct {
	fields []string
}

func (o *selectOpts) ToHeaders() *sdkClient.Headers { return nil }
func (o *selectOpts) ToOData() *odata.Query {
	return &odata.Query{Select: o.fields}
}
func (o *selectOpts) ToQuery() *sdkClient.QueryParams { return nil }

// ── Terraform config templates ────────────────────────────────────────────────

// template returns the shared service principal used by all test configs.
// NOTE: The attribute set "TestAttributeSet" and its attributes must be
// pre-created in the tenant before running these tests. Custom security
// attribute definitions cannot be managed by this provider.
func (r CustomSecurityAttributeAssignmentResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctestCSA-%[1]d"
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}
`, data.RandomInteger)
}

func (r CustomSecurityAttributeAssignmentResource) singleStringValue(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "test" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name  = "Environment"
    value = "Production"
  }
}
`, r.template(data))
}

func (r CustomSecurityAttributeAssignmentResource) singleStringValueUpdated(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "test" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name  = "Environment"
    value = "Staging"
  }
}
`, r.template(data))
}

func (r CustomSecurityAttributeAssignmentResource) multiStringValues(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "test" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name   = "AllowedLocations"
    values = ["US", "EU"]
  }
}
`, r.template(data))
}

func (r CustomSecurityAttributeAssignmentResource) booleanValue(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "test" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name          = "IsCompliant"
    boolean_value = true
  }
}
`, r.template(data))
}

func (r CustomSecurityAttributeAssignmentResource) multipleAttributes(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "test" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name  = "Environment"
    value = "Production"
  }

  attribute {
    name   = "AllowedLocations"
    values = ["US", "EU"]
  }

  attribute {
    name          = "IsCompliant"
    boolean_value = true
  }
}
`, r.template(data))
}

func (r CustomSecurityAttributeAssignmentResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_custom_security_attribute_assignment" "import" {
  principal_id  = azuread_service_principal.test.object_id
  attribute_set = "TestAttributeSet"

  attribute {
    name  = "Environment"
    value = "Production"
  }
}
`, r.singleStringValue(data))
}

// ── JSON debug helper (used in Exists) ────────────────────────────────────────

var _ = json.Marshal // ensure json import is used
