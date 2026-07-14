// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package groups_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

type GroupLicenseResource struct{}

func TestAccGroupLicense_basic(t *testing.T) {
	skuId := os.Getenv("AZUREAD_TEST_SKU_ID")
	if skuId == "" {
		t.Skip("AZUREAD_TEST_SKU_ID not set")
	}

	data := acceptance.BuildTestData(t, "azuread_group_license", "test")
	r := GroupLicenseResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, skuId),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_id").Exists(),
				check.That(data.ResourceName).Key("sku_id").HasValue(skuId),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupLicense_requiresImport(t *testing.T) {
	skuId := os.Getenv("AZUREAD_TEST_SKU_ID")
	if skuId == "" {
		t.Skip("AZUREAD_TEST_SKU_ID not set")
	}

	data := acceptance.BuildTestData(t, "azuread_group_license", "test")
	r := GroupLicenseResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, skuId),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data, skuId)),
	})
}

func TestAccGroupLicense_disabledPlans(t *testing.T) {
	skuId := os.Getenv("AZUREAD_TEST_SKU_ID")
	if skuId == "" {
		t.Skip("AZUREAD_TEST_SKU_ID not set")
	}
	disabledPlanId := os.Getenv("AZUREAD_TEST_DISABLED_PLAN_ID")
	if disabledPlanId == "" {
		t.Skip("AZUREAD_TEST_DISABLED_PLAN_ID not set")
	}

	data := acceptance.BuildTestData(t, "azuread_group_license", "test")
	r := GroupLicenseResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.disabledPlans(data, skuId, disabledPlanId),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("sku_id").HasValue(skuId),
				check.That(data.ResourceName).Key("disabled_plans.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func (r GroupLicenseResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Groups.GroupClientBeta

	id, err := parse.GroupLicenseID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetGroup(ctx, beta.NewGroupID(id.GroupId), group.GetGroupOperationOptions{
		Select: &[]string{"id", "assignedLicenses"},
	})
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	g := resp.Model
	if g == nil || g.AssignedLicenses == nil {
		return pointer.To(false), nil
	}

	for _, license := range *g.AssignedLicenses {
		if license.SkuId.GetOrZero() == id.SkuId {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (GroupLicenseResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test" {
  display_name     = "acctestGroupLicense-%[1]d"
  security_enabled = true
}
`, data.RandomInteger)
}

func (r GroupLicenseResource) basic(data acceptance.TestData, skuId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_license" "test" {
  group_id = azuread_group.test.object_id
  sku_id   = "%[2]s"
}
`, r.template(data), skuId)
}

func (r GroupLicenseResource) requiresImport(data acceptance.TestData, skuId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_license" "import" {
  group_id = azuread_group_license.test.group_id
  sku_id   = azuread_group_license.test.sku_id
}
`, r.basic(data, skuId))
}

func (r GroupLicenseResource) disabledPlans(data acceptance.TestData, skuId, disabledPlanId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_license" "test" {
  group_id       = azuread_group.test.object_id
  sku_id         = "%[2]s"
  disabled_plans = ["%[3]s"]
}
`, r.template(data), skuId, disabledPlanId)
}
