// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/claimsmappingpolicy"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
)

type ServicePrincipalClaimsMappingPolicyAssignmentResource struct{}

func TestClaimsMappingPolicyAssignment_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_claims_mapping_policy_assignment", "test")
	r := ServicePrincipalClaimsMappingPolicyAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicClaimsMappingPolicyAssignment(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r ServicePrincipalClaimsMappingPolicyAssignmentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ClaimsMappingPolicyClient

	id, err := parse.ClaimsMappingPolicyAssignmentID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing CLaims Mapping Policy Assignment ID: %v", err)
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ServicePrincipalId)

	resp, err := client.ListClaimsMappingPolicies(ctx, servicePrincipalId, claimsmappingpolicy.DefaultListClaimsMappingPoliciesOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, fmt.Errorf("%s does not exist", servicePrincipalId)
		}
		return nil, fmt.Errorf("failed to retrieve claims mapping policy assignments for %s: %+v", servicePrincipalId, err)
	}

	if resp.Model != nil {
		for _, p := range *resp.Model {
			if pointer.From(p.Id) == id.ClaimsMappingPolicyId {
				return pointer.To(true), nil
			}
		}
	}

	return pointer.To(false), nil
}

func (ServicePrincipalClaimsMappingPolicyAssignmentResource) basicClaimsMappingPolicyAssignment(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_claims_mapping_policy" "test" {
  definition = [
    "{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\",\"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"employeeid\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name\",\"JwtClaimType\":\"name\"},{\"Source\":\"company\",\"ID\":\"tenantcountry\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country\",\"JwtClaimType\":\"country\"}]}}"
  ]
  display_name = "acctest-%[1]s"
}

resource "azuread_service_principal" "msgraph" {
  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing   = true
}

resource "azuread_service_principal_claims_mapping_policy_assignment" "test" {
  claims_mapping_policy_id = azuread_claims_mapping_policy.test.id
  service_principal_id     = azuread_service_principal.msgraph.id
}
`, data.RandomString)
}
