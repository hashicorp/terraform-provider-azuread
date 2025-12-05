// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/parse"
)

type GroupRoleManagementPolicyResource struct{}

func TestAccGroupRoleManagementPolicy_activationRules(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_role_management_policy", "test")
	r := GroupRoleManagementPolicyResource{}

	// Ignore the dangling resource post-test as the policy remains while the group is in a pending deletion state
	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.activationRules(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupRoleManagementPolicy_member(t *testing.T) {
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
		data.ImportStep(),
	})
}

func TestAccGroupRoleManagementPolicy_owner(t *testing.T) {
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
		data.ImportStep(),
	})

}

func (GroupRoleManagementPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Policies.RoleManagementPolicyClient
	policyId, err := parse.ParseRoleManagementPolicyID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("could not parse policy ID: %v", err)
	}

	id := stable.NewPolicyRoleManagementPolicyID(policyId.ID())

	policyOptions := rolemanagementpolicy.GetRoleManagementPolicyOperationOptions{
		Expand: &odata.Expand{
			Relationship: "*",
		},
	}

	policyResp, err := client.GetRoleManagementPolicy(ctx, id, policyOptions)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %v", id, err)
	}

	policy := policyResp.Model
	if policy == nil {
		return nil, fmt.Errorf("retrieving %s: API error, model was nil", id)
	}

	return pointer.To(true), nil
}

func (GroupRoleManagementPolicyResource) activationRules(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "this" {
  display_name     = "PAM Basic Test %[1]s"
  security_enabled = true
}

resource "azuread_group_role_management_policy" "test" {
  group_id = azuread_group.this.object_id
  role_id  = "member"

  activation_rules {
    maximum_duration = "PT12H"
  }
}
`, data.RandomString)
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
  group_id = azuread_group.pam.object_id
  role_id  = "member"

  eligible_assignment_rules {
    expiration_required = true
    expire_after        = "P365D"
  }

  active_assignment_rules {
    expiration_required = true
    expire_after        = "P180D"
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
  group_id = azuread_group.pam.object_id
  role_id  = "owner"

  eligible_assignment_rules {
    expiration_required = true
    expire_after        = "P365D"
  }

  active_assignment_rules {
    expiration_required = true
    expire_after        = "P90D"
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
