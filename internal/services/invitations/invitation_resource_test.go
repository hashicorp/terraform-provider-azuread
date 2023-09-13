// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package invitations_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type InvitationResource struct{}

func TestAccInvitation_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("redeem_url").Exists(),
				check.That(data.ResourceName).Key("redirect_url").HasValue("https://portal.azure.com"),
				check.That(data.ResourceName).Key("user_email_address").HasValue(fmt.Sprintf("acctest-user-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("user_type").HasValue("Guest"),
			),
		},
	})
}

func TestAccInvitation_member(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.member(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("redeem_url").Exists(),
				check.That(data.ResourceName).Key("redirect_url").HasValue("https://portal.azure.com"),
				check.That(data.ResourceName).Key("user_email_address").HasValue(fmt.Sprintf("acctest-user-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("user_type").HasValue("Member"),
			),
		},
	})
}

func TestAccInvitation_message(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withMessage(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("redeem_url").Exists(),
				check.That(data.ResourceName).Key("redirect_url").HasValue("https://portal.azure.com"),
				check.That(data.ResourceName).Key("user_display_name").HasValue("Test user"),
				check.That(data.ResourceName).Key("user_email_address").HasValue(fmt.Sprintf("acctest-user-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("message.#").HasValue("1"),
				check.That(data.ResourceName).Key("user_type").HasValue("Guest"),
			),
		},
	})
}

func TestAccInvitation_messageWithCustomizedBody(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withMessageHavingCustomizedBody(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("redeem_url").Exists(),
				check.That(data.ResourceName).Key("redirect_url").HasValue("https://portal.azure.com"),
				check.That(data.ResourceName).Key("user_display_name").HasValue("Test user"),
				check.That(data.ResourceName).Key("user_email_address").HasValue(fmt.Sprintf("acctest-user-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("message.#").HasValue("1"),
				check.That(data.ResourceName).Key("message.0.additional_recipients.#").HasValue("1"),
				check.That(data.ResourceName).Key("message.0.additional_recipients.0").HasValue(fmt.Sprintf("acctest-another-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("message.0.body").HasValue("Hello there! You are invited to join my Azure tenant."),
				check.That(data.ResourceName).Key("user_type").HasValue("Guest"),
			),
		},
	})
}

func TestAccInvitation_messageWithLanguage(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withMessageHavingLanguage(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("redeem_url").Exists(),
				check.That(data.ResourceName).Key("redirect_url").HasValue("https://portal.azure.com"),
				check.That(data.ResourceName).Key("user_display_name").HasValue("Test user"),
				check.That(data.ResourceName).Key("user_email_address").HasValue(fmt.Sprintf("acctest-user-%s@test.com", data.RandomString)),
				check.That(data.ResourceName).Key("user_id").Exists(),
				check.That(data.ResourceName).Key("message.#").HasValue("1"),
				check.That(data.ResourceName).Key("message.0.language").HasValue("fr-CA"),
				check.That(data.ResourceName).Key("user_type").HasValue("Guest"),
			),
		},
	})
}

func TestAccInvitation_withGroupMembership(t *testing.T) {
	count := 10
	data := acceptance.BuildTestData(t, "azuread_invitation", fmt.Sprintf("test.%d", count-1))
	r := InvitationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.withGroupMembership(data, count),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r InvitationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Invitations.UsersClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	userID := state.Attributes["user_id"]

	user, status, err := client.Get(ctx, userID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Invited user with object ID %q does not exist", userID)
		}
		return nil, fmt.Errorf("failed to retrieve invited user with object ID %q: %+v", userID, err)
	}

	return utils.Bool(user.ID() != nil && *user.ID() == userID), nil
}

func (InvitationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_invitation" "test" {
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[1]s@test.com"
}
`, data.RandomString)
}

func (InvitationResource) member(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_invitation" "test" {
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[1]s@test.com"
  user_type          = "Member"
}
`, data.RandomString)
}

func (InvitationResource) withMessage(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_invitation" "test" {
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[1]s@test.com"
  user_display_name  = "Test user"

  message {}
}
`, data.RandomString)
}

func (InvitationResource) withMessageHavingCustomizedBody(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_invitation" "test" {
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[1]s@test.com"
  user_display_name  = "Test user"

  message {
    additional_recipients = ["acctest-another-%[1]s@test.com"]
    body                  = "Hello there! You are invited to join my Azure tenant."
  }
}
`, data.RandomString)
}

func (InvitationResource) withMessageHavingLanguage(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_invitation" "test" {
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[1]s@test.com"
  user_display_name  = "Test user"

  message {
    language = "fr-CA"
  }
}
`, data.RandomString)
}

func (InvitationResource) withGroupMembership(data acceptance.TestData, count int) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
}

resource "azuread_invitation" "test" {
  count              = %[3]d
  redirect_url       = "https://portal.azure.com"
  user_email_address = "acctest-user-%[2]s-groupMember-${count.index}@test.com"
}

resource "azuread_group_member" "test" {
  count            = %[3]d
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_invitation.test[count.index].user_id
}
`, data.RandomInteger, data.RandomString, count)
}
