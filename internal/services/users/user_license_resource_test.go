// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package users_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/users/parse"
)

type UserLicenseResource struct{}

// These acceptance tests require a license SKU that is available in the test tenant. As license
// availability is tenant-specific, the SKU ID (and optionally a service plan ID to disable) must be
// provided via environment variables, otherwise the tests are skipped.
const (
	skuIdEnvVar         = "AZUREAD_TEST_SKU_ID"
	disabledPlanEnvVar  = "AZUREAD_TEST_DISABLED_PLAN_ID"
	skipMessageNoSkuId  = "skipping as `" + skuIdEnvVar + "` is not specified"
	skipMessageNoPlanId = "skipping as `" + disabledPlanEnvVar + "` is not specified"
)

func TestAccUserLicense_basic(t *testing.T) {
	skuId := os.Getenv(skuIdEnvVar)
	if skuId == "" {
		t.Skip(skipMessageNoSkuId)
	}

	data := acceptance.BuildTestData(t, "azuread_user_license", "test")
	r := UserLicenseResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data, skuId),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("sku_id").HasValue(skuId),
			),
		},
		data.ImportStep(),
	})
}

func TestAccUserLicense_requiresImport(t *testing.T) {
	skuId := os.Getenv(skuIdEnvVar)
	if skuId == "" {
		t.Skip(skipMessageNoSkuId)
	}

	data := acceptance.BuildTestData(t, "azuread_user_license", "test")
	r := UserLicenseResource{}

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

func TestAccUserLicense_disabledPlans(t *testing.T) {
	skuId := os.Getenv(skuIdEnvVar)
	if skuId == "" {
		t.Skip(skipMessageNoSkuId)
	}
	disabledPlanId := os.Getenv(disabledPlanEnvVar)
	if disabledPlanId == "" {
		t.Skip(skipMessageNoPlanId)
	}

	data := acceptance.BuildTestData(t, "azuread_user_license", "test")
	r := UserLicenseResource{}

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

func (r UserLicenseResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Users.UserClient

	id, err := parse.UserLicenseID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetUser(ctx, stable.NewUserID(id.UserId), user.GetUserOperationOptions{
		Select: &[]string{"id", "assignedLicenses", "licenseAssignmentStates"},
	})
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	u := resp.Model
	if u == nil || u.LicenseAssignmentStates == nil {
		return pointer.To(false), nil
	}

	for _, assignmentState := range *u.LicenseAssignmentStates {
		if assignmentState.SkuId.GetOrZero() == id.SkuId && assignmentState.AssignedByGroup.GetOrZero() == "" {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (UserLicenseResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUserLicense.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUserLicense-%[1]d"
  password            = "%[2]s"
  usage_location      = "US"
}
`, data.RandomInteger, data.RandomPassword)
}

func (r UserLicenseResource) basic(data acceptance.TestData, skuId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_user_license" "test" {
  user_id = azuread_user.test.object_id
  sku_id  = "%[2]s"
}
`, r.template(data), skuId)
}

func (r UserLicenseResource) requiresImport(data acceptance.TestData, skuId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_user_license" "import" {
  user_id = azuread_user_license.test.user_id
  sku_id  = azuread_user_license.test.sku_id
}
`, r.basic(data, skuId))
}

func (r UserLicenseResource) disabledPlans(data acceptance.TestData, skuId, disabledPlanId string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_user_license" "test" {
  user_id        = azuread_user.test.object_id
  sku_id         = "%[2]s"
  disabled_plans = ["%[3]s"]
}
`, r.template(data), skuId, disabledPlanId)
}
