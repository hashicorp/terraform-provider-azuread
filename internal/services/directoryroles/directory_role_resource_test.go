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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type DirectoryRoleResource struct{}

func TestAccDirectoryRole_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role", "test")
	r := DirectoryRoleResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.byDisplayName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
				check.That(data.ResourceName).Key("template_id").IsUuid(),
			),
		},
	})
}

func TestAccDirectoryRole_byTemplateId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role", "test")
	r := DirectoryRoleResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.byTemplateId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").HasValue("Printer Administrator"),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
	})
}

func (r DirectoryRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.DirectoryRoleClient
	id := stable.NewDirectoryRoleID(state.ID)

	resp, err := client.GetDirectoryRole(ctx, id, directoryrole.DefaultGetDirectoryRoleOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve Directory Role with object ID %q: %+v", state.ID, err)
	}

	return pointer.To(true), nil
}

func (DirectoryRoleResource) byDisplayName(_ acceptance.TestData) string {
	return `
provider "azuread" {}

resource "azuread_directory_role" "test" {
  display_name = "Teams administrator"
}
`
}

func (DirectoryRoleResource) byTemplateId(_ acceptance.TestData) string {
	return `
provider "azuread" {}

resource "azuread_directory_role" "test" {
  template_id = "644ef478-e28f-4e28-b9dc-3fdde9aa0b1f" // Printer administrator
}
`
}
