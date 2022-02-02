package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/manicminer/hamilton/odata"
)

type ServicePrincipalClaimsMappingPolicy struct{}

func TestClaimsMappingPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_claims_mapping_policy", "test")
	r := ServicePrincipalClaimsMappingPolicy{}
	updatedRegex, _ := regexp.Compile(`updated`)

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicClaimsMappingPolicy(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.updateClaimsMappingPolicy(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").MatchesRegex(updatedRegex),
			),
		},
	})
}

func (ServicePrincipalClaimsMappingPolicy) basicClaimsMappingPolicy(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_claims_mapping_policy" "test" {
  definition = [
	  "{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\",\"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"employeeid\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name\",\"JwtClaimType\":\"name\"},{\"Source\":\"company\",\"ID\":\"tenantcountry\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country\",\"JwtClaimType\":\"country\"}]}}"
  ]
  description = "%[1]s"
  display_name = "integration-%[1]s"
}
`, data.RandomString)
}

func (ServicePrincipalClaimsMappingPolicy) updateClaimsMappingPolicy(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_claims_mapping_policy" "test" {
  definition = [
	  "{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"true\",\"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"employeeid\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name\",\"JwtClaimType\":\"name\"},{\"Source\":\"company\",\"ID\":\"tenantcountry\",\"SamlClaimType\":\"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country\",\"JwtClaimType\":\"country\"}]}}"
  ]
  description = "%[1]s updated"
  display_name = "integration-%[1]s-updated"
}
`, data.RandomString)
}

func (r ServicePrincipalClaimsMappingPolicy) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ClaimsMappingPolicyClient
	client.BaseClient.DisableRetries = true

	exists := false
	_, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Claims mapping policy with object ID %q does not exist", state.ID)
		}
		return &exists, fmt.Errorf("failed to retrieve claims mapping policy with object ID %q: %+v", state.ID, err)
	}

	exists = true
	return &exists, nil
}
