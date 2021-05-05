package conditionalAccessPolicies_test

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

type ConditionalAccessPolicyResource struct{}

func TestAccConditionalAccessPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-POLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_basicEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicEnabled(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-POLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("enabled"),
			),
		},
		data.ImportStep(),
	})
}

func (r ConditionalAccessPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	var id *string

	switch clients.EnableMsGraphBeta {
	case true:
		app, status, err := clients.ConditionalAccessPolicies.MsClient.Get(ctx, state.ID)
		if err != nil {
			if status == http.StatusNotFound {
				return nil, fmt.Errorf("Conditional Access Policy with ID %q does not exist", state.ID)
			}
			return nil, fmt.Errorf("failed to retrieve Conditional Access Policy with ID %q: %+v", state.ID, err)
		}
		id = app.ID

	case false:
		return nil, fmt.Errorf("Resource does not support AAD Client")
	}

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (ConditionalAccessPolicyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-POLICY-%[1]d"
  state = "disabled"
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) basicEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-POLICY-%[1]d"
  state = "enabled"
}
`, data.RandomInteger)
}
