// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/identitygovernance/helpers"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedule"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type PrivilegedAccessGroupEligibilityScheduleResource struct{}

func TestAccPrivilegedAccessGroupEligibilitySchedule_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_privileged_access_group_eligibility_schedule", "member")
	r := PrivilegedAccessGroupEligibilityScheduleResource{}

	endTime := time.Now().AddDate(0, 2, 0).UTC()

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.member(data, endTime),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				// There is a minimum life of 5 minutes for a schedule request to exist.
				// Attempting to delete the request within this time frame will result in
				// a 400 error on destroy, which we can't trap.
				helpers.SleepCheck(5*time.Minute+15*time.Second),
			),
		},
		data.ImportStep(),
	})
}

func TestAccPrivilegedAccessGroupEligibilitySchedule_owner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_privileged_access_group_eligibility_schedule", "owner")
	r := PrivilegedAccessGroupEligibilityScheduleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.owner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				// There is a minimum life of 5 minutes for a schedule request to exist.
				// Attempting to delete the request within this time frame will result in
				// a 400 error on destroy, which we can't trap.
				helpers.SleepCheck(5*time.Minute+15*time.Second),
			),
		},
		data.ImportStep(),
	})
}

func (PrivilegedAccessGroupEligibilityScheduleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.PrivilegedAccessGroupEligibilityScheduleClient

	resourceId, err := parse.ParsePrivilegedAccessGroupScheduleID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse privileged group assignment schedule ID %q: %+v", state.ID, err)
	}

	id := stable.NewIdentityGovernancePrivilegedAccessGroupEligibilityScheduleID(resourceId.ID())

	resp, err := client.GetPrivilegedAccessGroupEligibilitySchedule(ctx, id, privilegedaccessgroupeligibilityschedule.DefaultGetPrivilegedAccessGroupEligibilityScheduleOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(true), nil
}

func (PrivilegedAccessGroupEligibilityScheduleResource) member(data acceptance.TestData, endTime time.Time) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "pam" {
  display_name     = "Privileged Eligibility %[1]s"
  mail_enabled     = false
  security_enabled = true
}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "member" {
  user_principal_name = "pam-member-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Member %[1]s"
  password            = "%[2]s"
}

resource "azuread_privileged_access_group_eligibility_schedule" "member" {
  group_id        = azuread_group.pam.object_id
  principal_id    = azuread_user.member.object_id
  assignment_type = "member"
  expiration_date = "%[3]s"
  justification   = "required"
}
`, data.RandomString, data.RandomPassword, endTime.Format(time.RFC3339))
}

func (PrivilegedAccessGroupEligibilityScheduleResource) owner(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "manual_owner" {
  user_principal_name = "pam-eligible-owner-manual-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Owner (Manual) %[1]s"
  password            = "%[2]s"
}

resource "azuread_group" "pam" {
  display_name     = "Privileged Eligibility %[1]s"
  mail_enabled     = false
  security_enabled = true

  owners = [azuread_user.manual_owner.object_id]

  lifecycle {
    ignore_changes = [owners]
  }
}

resource "azuread_user" "eligibile_owner" {
  user_principal_name = "pam-eligible-owner-eligibile-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Owner (Eligibile) %[1]s"
  password            = "%[2]s"
}


resource "azuread_privileged_access_group_eligibility_schedule" "owner" {
  group_id        = azuread_group.pam.object_id
  principal_id    = azuread_user.eligibile_owner.object_id
  assignment_type = "owner"
  duration        = "P30D"
  justification   = "required"
}
`, data.RandomString, data.RandomPassword)
}
