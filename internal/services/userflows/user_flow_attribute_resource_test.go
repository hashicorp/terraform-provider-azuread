// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userflows_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/userflowattribute"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type UserFlowAttributeResource struct{}

func TestAccUserFlowAttribute_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user_flow_attribute", "test")
	r := UserFlowAttributeResource{}

	data.ResourceSequentialTest(t, r, []acceptance.TestStep{
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
	r := UserFlowAttributeResource{}

	data.ResourceSequentialTest(t, r, []acceptance.TestStep{
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
	r := UserFlowAttributeResource{}

	data.ResourceSequentialTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r UserFlowAttributeResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.UserFlows.UserFlowAttributeClient

	id, err := stable.ParseIdentityUserFlowAttributeID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetUserFlowAttribute(ctx, *id, userflowattribute.DefaultGetUserFlowAttributeOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}

		return nil, fmt.Errorf("retrieving %s: %v", id, err)
	}

	return pointer.To(true), nil
}

func (r UserFlowAttributeResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_user_flow_attribute" "test" {
  display_name = "acctestUserFlowAttr%s"
  description  = "acctest description %[1]s"
  data_type    = "string"
}
`, data.RandomString)
}

func (r UserFlowAttributeResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_user_flow_attribute" "test" {
  display_name = "acctestUserFlowAttr%s"
  description  = "acctest new description %[1]s"
  data_type    = "string"
}
`, data.RandomString)
}

func (r UserFlowAttributeResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_user_flow_attribute" "import" {
  display_name = azuread_user_flow_attribute.test.display_name
  description  = azuread_user_flow_attribute.test.description
  data_type    = azuread_user_flow_attribute.test.data_type
}
`, r.basic(data), data.RandomString)
}
