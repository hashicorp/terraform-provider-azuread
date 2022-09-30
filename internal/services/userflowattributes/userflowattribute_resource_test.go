package userflowattributes_test

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

type UserflowAttributeResource struct{}

func TestAccUserflowAttribute(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_userflow_attribute", "test")
	r := UserflowAttributeResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			// Create.
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			// Update.
			Config: r.basic_newDescription(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			// Update again.
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r UserflowAttributeResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.UserflowAttributes.Client
	client.BaseClient.DisableRetries = true

	userflowAttr, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Userflow attribute with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Userflow attribute with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(userflowAttr.ID != nil && *userflowAttr.ID == state.ID), nil
}

func (r UserflowAttributeResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_userflow_attribute" "test" {
  display_name        				= "acctestUserflowAttr%s"
  description 						= "acc test description %[1]s"
  userflow_attribute_type           = "custom"
  data_type 						= "string"
}
`, data.RandomString)
}

func (r UserflowAttributeResource) basic_newDescription(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_userflow_attribute" "test" {
  display_name        				= "acctestUserflowAttr%s"
  description 						= "acc test description %[1]s new"
  userflow_attribute_type           = "custom"
  data_type 						= "string"
}
`, data.RandomString)
}
