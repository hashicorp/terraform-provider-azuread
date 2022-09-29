package b2cuserflow_test

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
	"github.com/manicminer/hamilton/odata"
)

type B2CUserflowResource struct{}

func TestAccUserflow_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_b2c_userflow", "test")
	r := B2CUserflowResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccUserflow_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_b2c_userflow", "test")
	r := B2CUserflowResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccUserflow_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_b2c_userflow", "test")
	r := B2CUserflowResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r B2CUserflowResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.B2CUserFlow.UserFlowClient
	client.BaseClient.DisableRetries = true

	userflow, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Userflow with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Userflow with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(userflow.ID != nil && *userflow.ID == state.ID), nil
}

func (r B2CUserflowResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_b2c_userflow" "test" {
  name = "%sacctest-userflow"
  user_flow_type = "signUp"
  user_flow_type_version = "3"
  default_language_tag = "en"
}`, data.RandomString)
}

func (r B2CUserflowResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_b2c_userflow" "test" {
  name = "%sacctest-userflow"
  user_flow_type = "signUp"
  user_flow_type_version = "3"
  default_language_tag = "en"
  is_language_customization_enabled = true
}`, data.RandomString)
}
