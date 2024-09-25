// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type RoleEligibilityScheduleRequestResource struct{}

func TestAccRoleEligibilityScheduleRequest_builtin(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_eligibility_schedule_request", "test")
	r := RoleEligibilityScheduleRequestResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.builtin(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r RoleEligibilityScheduleRequestResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.DirectoryRoleEligibilityScheduleRequestClient
	id := stable.NewRoleManagementDirectoryRoleEligibilityScheduleRequestID(state.ID)

	resp, err := client.GetDirectoryRoleEligibilityScheduleRequest(ctx, id, directoryroleeligibilityschedulerequest.DefaultGetDirectoryRoleEligibilityScheduleRequestOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", id, err)
	}

	return pointer.To(true), nil
}

func (r RoleEligibilityScheduleRequestResource) builtin(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestManager.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestManager-%[1]d"
  password            = "%[2]s"
}

resource "azuread_directory_role" "test" {
  display_name = "Application Administrator"
}

resource "azuread_directory_role_eligibility_schedule_request" "test" {
  role_definition_id = azuread_directory_role.test.template_id
  principal_id       = azuread_user.test.object_id
  directory_scope_id = "/"
  justification      = "abc"
}
`, data.RandomInteger, data.RandomPassword)
}
