// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userflows_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type UserflowAttributeResource struct{}

func TestAccUserFlowAttribute_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user_flow_attribute", "test")
	r := UserflowAttributeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccUserFlowAttribute_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user_flow_attribute", "test")
	r := UserflowAttributeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.update(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func TestAccUserFlowAttribute_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user_flow_attribute", "test")
	r := UserflowAttributeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r UserflowAttributeResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.UserFlows.UserFlowAttributesClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	userFlowAttr, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("user flow attribute with ID %q does not exist", state.ID)
		}

		return nil, fmt.Errorf("failed to retrieve User Flow attribute with ID %q: %+v", state.ID, err)
	}

	return pointer.To(userFlowAttr.ID != nil && *userFlowAttr.ID == state.ID), nil
}

func (r UserflowAttributeResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_user_flow_attribute" "test" {
  display_name = "acctestUserFlowAttr%s"
  description  = "acctest description %[1]s"
  data_type    = "string"
}
`, data.RandomString)
}

func (r UserflowAttributeResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_user_flow_attribute" "test" {
  display_name = "acctestUserFlowAttr%s"
  description  = "acctest new description %[1]s"
  data_type    = "string"
}
`, data.RandomString)
}

func (r UserflowAttributeResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_user_flow_attribute" "import" {
  display_name = azuread_user_flow_attribute.test.display_name
  description  = azuread_user_flow_attribute.test.description
  data_type    = azuread_user_flow_attribute.test.data_type
}
`, r.basic(data), data.RandomString)
}
