// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type GroupRoleManagementPolicyDataSource struct{}

func TestGroupRoleManagementPolicyDataSource_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group_role_management_policy", "test")
	r := GroupRoleManagementPolicyDataSource{}

	// Ignore the dangling resource post-test as the policy remains while the group is in a pending deletion state
	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.member(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (GroupRoleManagementPolicyDataSource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Policies.RoleManagementPolicyClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	_, status, err := client.Get(ctx, state.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve role management policy with ID %q: %+v", state.ID, err)
	}

	return pointer.To(true), nil
}

func (GroupRoleManagementPolicyDataSource) member(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "pam" {
  display_name     = "PAM Member Test %[1]s"
  mail_enabled     = false
  security_enabled = true
}

data "azuread_group_role_management_policy" "test" {
  group_id = azuread_group_role_management_policy.test.group_id
  role_id  = "member"
}
`, data.RandomString)
}
