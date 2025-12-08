// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"fmt"
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type GroupRoleManagementPolicyDataSource struct{}

func TestAccGroupRoleManagementPolicyDataSource_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group_role_management_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupRoleManagementPolicyDataSource{}.member(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").Exists(),
			),
		},
	})
}

func (GroupRoleManagementPolicyDataSource) member(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group_role_management_policy" "test" {
  group_id = azuread_group_role_management_policy.test.group_id
  role_id  = "member"
}
`, GroupRoleManagementPolicyResource{}.member(data))
}
