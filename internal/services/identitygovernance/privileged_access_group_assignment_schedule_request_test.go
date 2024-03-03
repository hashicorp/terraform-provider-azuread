// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupAssignmentScheduleRequestResource struct{}

func TestPrivilegedAccessGroupAssignmentScheduleRequest_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_privileged_access_group_assignment_schedule_request", "member")
	r := PrivilegedAccessGroupAssignmentScheduleRequestResource{}

	endTime := time.Now().AddDate(0, 2, 0).UTC()

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.member(data, endTime),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				// There is a minimum life of 5 minutes for a schedule request to exist.
				// Attempting to delete the request within this time frame will result in
				// a 400 error on destroy, which we can't trap.
				sleepCheck(5*time.Minute+1*time.Second),
			),
		},
	})
}

func TestPrivilegedAccessGroupAssignmentScheduleRequest_owner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_privileged_access_group_assignment_schedule_request", "owner")
	r := PrivilegedAccessGroupAssignmentScheduleRequestResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.owner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				// There is a minimum life of 5 minutes for a schedule request to exist.
				// Attempting to delete the request within this time frame will result in
				// a 400 error on destroy, which we can't trap.
				sleepCheck(5*time.Minute+1*time.Second),
			),
		},
	})

}

func (PrivilegedAccessGroupAssignmentScheduleRequestResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.PrivilegedAccessGroupAssignmentScheduleRequestsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	request, status, err := client.Get(ctx, state.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve privileged group assignment schedule request with ID %q: %+v", state.ID, err)
	}

	// Requests are not deleted, but marked as canceled or revoked.
	if slices.Contains([]string{
		msgraph.PrivilegedAccessGroupAssignmentStatusCanceled,
		msgraph.PrivilegedAccessGroupAssignmentStatusRevoked,
	}, request.Status) {
		return pointer.To(false), nil
	} else {
		return pointer.To(true), nil
	}
}

func (PrivilegedAccessGroupAssignmentScheduleRequestResource) member(data acceptance.TestData, endTime time.Time) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "pam" {
  display_name     = "Privileged %[1]s"
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

resource "azuread_privileged_access_group_assignment_schedule_request" "member" {
  group_id        = azuread_group.pam.id
  principal_id    = azuread_user.member.id
  assignment_type = "member"
  expiration_date = "%[3]s"
  justification   = "required"
}
`, data.RandomString, data.RandomPassword, endTime.Format(time.RFC3339))
}

func (PrivilegedAccessGroupAssignmentScheduleRequestResource) owner(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "manual_owner" {
  user_principal_name = "pam-owner-manual-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Owner (Manual) %[1]s"
  password            = "%[2]s"
}

resource "azuread_group" "pam" {
  display_name     = "Privileged %[1]s"
  mail_enabled     = false
  security_enabled = true

	owners = [azuread_user.manual_owner.object_id]

	lifecycle {
		ignore_changes = [owners]
	}
}

resource "azuread_user" "assigned_owner" {
  user_principal_name = "pam-owner-assigned-%[1]s@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "PAM Owner (Assigned) %[1]s"
  password            = "%[2]s"
}


resource "azuread_privileged_access_group_assignment_schedule_request" "owner" {
  group_id        = azuread_group.pam.id
  principal_id    = azuread_user.assigned_owner.id
  assignment_type = "owner"
  duration        = "P30D"
  justification   = "required"
}
`, data.RandomString, data.RandomPassword)
}

func sleepCheck(d time.Duration) acceptance.TestCheckFunc {
	return func(s *terraform.State) error {
		time.Sleep(d)
		return nil
	}
}
