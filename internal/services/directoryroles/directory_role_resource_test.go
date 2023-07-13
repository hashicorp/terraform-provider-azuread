// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type DirectoryRoleResource struct{}

func TestAccDirectoryRole_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role", "test")
	r := DirectoryRoleResource{}

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.byDisplayName(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.byTemplateId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").HasValue("Printer Administrator"),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
	})
}

func (r DirectoryRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.DirectoryRolesClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	role, status, err := client.Get(ctx, state.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Directory Role with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Directory Role with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(role.ID() != nil && *role.ID() == state.ID), nil
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
