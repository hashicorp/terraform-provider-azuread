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

type GroupRoleManagementPolicyResource struct{}

func TestGroupRoleManagementPolicy_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_role_management_policy", "test")
	r := GroupRoleManagementPolicyResource{}

	// Ignore the dangling resource post-test as the policy remains while the group is in a pending deletion state
	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.member(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("active_assignment_rules.0.expire_after").Exists(),
				check.That(data.ResourceName).Key("eligible_assignment_rules.0.expiration_required").Exists(),
			),
		},
	})
}

func TestGroupRoleManagementPolicy_owner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_role_management_policy", "test")
	r := GroupRoleManagementPolicyResource{}

	// Ignore the dangling resource post-test as the policy remains while the group is in a pending deletion state
	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.owner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("active_assignment_rules.0.expire_after").Exists(),
				check.That(data.ResourceName).Key("eligible_assignment_rules.0.expiration_required").Exists(),
			),
		},
	})

}

func (GroupRoleManagementPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
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

func (GroupRoleManagementPolicyResource) member(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "pam" {
  display_name     = "PAM Member Test %[1]s"
  mail_enabled     = false
  security_enabled = true
}

resource "azuread_group_role_management_policy" "test" {
  group_id        = azuread_group.pam.id
  assignment_type = "member"

  eligible_assignment_rules {
    expiration_required = false
  }

  active_assignment_rules {
    expire_after = "P365D"
  }

	notification_rules {
		eligible_assignments {
		  approver_notifications {
				notification_level    = "Critical"
				default_recipients    = false
				additional_recipients = ["someone@example.com"]
			}
		}
		eligible_activations {
		  assignee_notifications {
				notification_level    = "All"
				default_recipients    = true
				additional_recipients = ["someone@example.com"]
			}
		}
	}
}
`, data.RandomString)
}

func (GroupRoleManagementPolicyResource) owner(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}
 
resource "azuread_user" "approver" {
	user_principal_name = "pam-approver-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Approver Test %[1]s"
  password            = "%[2]s"
}

resource "azuread_group" "pam" {
  display_name     = "PAM Owner Test %[1]s"
  mail_enabled     = false
  security_enabled = true
}

resource "azuread_group_role_management_policy" "test" {
  group_id        = azuread_group.pam.id
  assignment_type = "owner"

  eligible_assignment_rules {
    expiration_required = false
  }

  active_assignment_rules {
    expire_after = "P90D"
  }

	activation_rules {
		maximum_duration = "PT1H"
		require_approval = true
		approval_stage {
			primary_approver {
				object_id = azuread_user.approver.object_id
				type      = "singleUser" 
			}
		}
	}

	notification_rules {
		active_assignments {
			admin_notifications {
				notification_level    = "Critical"
				default_recipients    = false
				additional_recipients = ["someone@example.com"]
			}
		}
	}
}
`, data.RandomString, data.RandomPassword)
}
