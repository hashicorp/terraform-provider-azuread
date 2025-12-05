// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package invitations_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type InvitationResource struct{}

func TestAccInvitation_guest(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_invitation", "test")
	r := InvitationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.member(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withMessage(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withMessageHavingCustomizedBody(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withMessageHavingLanguage(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withGroupMembership(data, count),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
	})
}

func (r InvitationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Invitations.UserClient
	userId := stable.NewUserID(state.Attributes["user_id"])

	resp, err := client.GetUser(ctx, userId, user.DefaultGetUserOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve invited %s: %+v", userId, err)
	}

	return pointer.To(true), nil
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
