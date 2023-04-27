package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalClaimsMappingPolicyAssignmentResource struct{}

func TestClaimsMappingPolicyAssignment_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_claims_mapping_policy_assignment", "test")
	r := ServicePrincipalClaimsMappingPolicyAssignmentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicClaimsMappingPolicyAssignment(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
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

func (r ServicePrincipalClaimsMappingPolicyAssignmentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ClaimsMappingPolicyAssignmentID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing CLaims Mapping Policy Assignment ID: %v", err)
	}

	policyList, status, err := client.ListClaimsMappingPolicy(ctx, id.ServicePrincipalId)
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), fmt.Errorf("Service Policy with object ID %q does not exist", id.ServicePrincipalId)
		}
		return utils.Bool(false), fmt.Errorf("failed to retrieve claims mapping policy assignments with service policy ID %q: %+v", id.ServicePrincipalId, err)
	}

	// Check the assignment is found in the currently assigned policies
	for _, policy := range *policyList {
		if policy.ID() != nil && *policy.ID() == id.ClaimsMappingPolicyId {
			return utils.Bool(true), nil
		}
	}

	return utils.Bool(false), nil
}
