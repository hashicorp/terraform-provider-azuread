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
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalClaimsMappingPolicyAssignment struct{}

func TestClaimsMappingPolicyAssignment_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_claims_mapping_policy_assignment", "test")
	mappingPolicy := acceptance.TestData{
		ResourceName: "azuread_claims_mapping_policy.test2",
	}
	r := ServicePrincipalClaimsMappingPolicyAssignment{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicClaimsMappingPolicyAssignment(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.updateClaimsMappingPolicyAssignment(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key(
					"claims_mapping_policy_id",
				).MatchesOtherKey(
					check.That(mappingPolicy.ResourceName).Key(
						"id",
					),
				),
			),
		},
	})
}

func (ServicePrincipalClaimsMappingPolicyAssignment) basicClaimsMappingPolicyAssignment(data acceptance.TestData) string {
	c := ServicePrincipalClaimsMappingPolicy{}
	return c.basicClaimsMappingPolicy(data) + `
	data "azuread_application_published_app_ids" "well_known" {}
	
	resource "azuread_service_principal" "msgraph" {
	  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
	  use_existing   = true
	}
	resource "azuread_claims_mapping_policy_assignment" "test" {
	  service_principal_id     = azuread_service_principal.msgraph.id
	  claims_mapping_policy_id = azuread_claims_mapping_policy.test.id
	}
	`
}

func (ServicePrincipalClaimsMappingPolicyAssignment) updateClaimsMappingPolicyAssignment(data acceptance.TestData) string {
	return fmt.Sprintf(`
	provider "azuread" {}

	data "azuread_application_published_app_ids" "well_known" {}

	resource "azuread_service_principal" "msgraph" {
	  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
	  use_existing   = true
	}

	resource "azuread_claims_mapping_policy" "test2" {
	  definition = [
		  "{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\",\"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"employeeid\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name\",\"JwtClaimType\":\"name\"},{\"Source\":\"company\",\"ID\":\"tenantcountry\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country\",\"JwtClaimType\":\"country\"}]}}"
	  ]
	  description = "%[1]s"
	  display_name = "integration-%[1]s"
	}

	resource "azuread_claims_mapping_policy_assignment" "test" {
	  service_principal_id     = azuread_service_principal.msgraph.id
	  claims_mapping_policy_id = azuread_claims_mapping_policy.test2.id
	}
`, data.RandomString)
}

func (r ServicePrincipalClaimsMappingPolicyAssignment) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true

	spID := state.Attributes["service_principal_id"]
	policyList, status, err := client.ListClaimsMappingPolicy(ctx, spID)
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), fmt.Errorf("Service Policy with object ID %q does not exist", spID)
		}
		return utils.Bool(false), fmt.Errorf("failed to retrieve claims mapping policy assignments with service policy ID %q: %+v", spID, err)
	}

	policyID := state.Attributes["claims_mapping_policy_id"]

	// Check the assignment is found in the currently assigned policies
	for _, policy := range *policyList {
		if *policy.ID == policyID {
			return utils.Bool(true), nil
		}
	}
	return utils.Bool(false), nil
}
