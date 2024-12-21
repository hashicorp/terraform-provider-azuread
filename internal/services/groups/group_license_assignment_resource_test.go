// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

type GroupLicenseAssignmentResource struct{}

func TestAccGrouplicenseassignment_license(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_license_assignment", "testA")
	r := GroupLicenseAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.license(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_object_id").IsUuid(),
				check.That(data.ResourceName).Key("sku_id").IsUuid(),
				check.That(data.ResourceName).Key("disabled_plans").IsEmpty(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupLicenseAssignment_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_license_assignment", "test")
	r := GroupLicenseAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.license(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r GroupLicenseAssignmentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Groups.GroupClientBeta

	id, err := parse.GroupLicenseAssignmentID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Group License Assignment ID: %v", err)
	}

	resp, err := client.GetGroup(ctx, beta.NewGroupID(id.GroupId), group.GetGroupOperationOptions{
		Select: &[]string{
			"assignedLicenses",
		},
	})
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve license %q (group ID: %q): %+v", id.SKUId, id.GroupId, err)
	}

	if resp.Model != nil {
		for _, license := range *resp.Model.AssignedLicenses {
			if license.SkuId.GetOrZero() == id.SKUId {
				return pointer.To(true), nil
			}
		}
	}

	return pointer.To(false), nil
}

func (GroupLicenseAssignmentResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
}
`, data.RandomInteger)
}

func (r GroupLicenseAssignmentResource) license(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_license_assignment" "test" {
  group_object_id  = azuread_group.test.object_id
  sku_id           = "90d8b3f8-712e-4f7b-aa1e-62e7ae6cbe96" # SMB_APPS
}
`, r.template(data))
}

func (r GroupLicenseAssignmentResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_license_assignment" "import" {
  group_object_id = azuread_group_license_assignment.test.group_object_id
  sku_id          = azuread_group_license_assignment.test.sku_id
	disabled_plans  = azuread_group_license_assignment.test.disabled_plans
}
`, r.license(data))
}
