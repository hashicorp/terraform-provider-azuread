package userflows_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

type B2CUserFlowResource struct{}

func TestAccUserflow_signIn(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_b2c_user_flow", "test")
	r := B2CUserFlowResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.signIn(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccUserflow_signUp(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_b2c_user_flow", "test_v1")
	r := B2CUserFlowResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.signUp(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.updateSignUp(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That("azuread_b2c_user_flow.test_v2").ExistsInAzure(r),
			),
		},
	})
}

func (r B2CUserFlowResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.UserFlows.UserFlowClient
	client.BaseClient.DisableRetries = true

	userFlow, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Userflow with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Userflow with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(userFlow.ID != nil && *userFlow.ID == state.ID), nil
}

func (r B2CUserFlowResource) signIn(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {
  tenant_id = "%[1]s"
}

resource "azuread_b2c_user_flow" "test" {
  name    = "B2C_1_SignIn-%[2]s"
  type    = "signIn"
  version = "1"
}
`, os.Getenv("ARM_TEST_B2C_TENANT_ID"), data.RandomString)
}

func (r B2CUserFlowResource) signUp(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {
  tenant_id = "%[1]s"
}

resource "azuread_b2c_user_flow" "test_v1" {
  name    = "B2C_1_SignUp-%[2]s"
  type    = "signUp"
  version = "1"
}
`, os.Getenv("ARM_TEST_B2C_TENANT_ID"), data.RandomString)
}

func (r B2CUserFlowResource) updateSignUp(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {
  tenant_id = "%[1]s"
}

resource "azuread_b2c_user_flow" "test_v1" {
  name    = "B2C_1_SignUp-%[2]s"
  type    = "signUp"
  version = "1"

  default_language_tag           = "en"
  language_customization_enabled = false
}

resource "azuread_b2c_user_flow" "test_v2" {
  name    = "B2C_1_SignUp-%[2]s"
  type    = "signUp"
  version = "2"

  default_language_tag           = "fr"
  language_customization_enabled = true
}
`, os.Getenv("ARM_TEST_B2C_TENANT_ID"), data.RandomString)
}
