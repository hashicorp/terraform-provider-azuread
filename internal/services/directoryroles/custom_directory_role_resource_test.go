// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type CustomDirectoryRoleResource struct{}

func TestAccCustomDirectoryRole_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_directory_role", "test")
	r := CustomDirectoryRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCustomDirectoryRole_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_directory_role", "test")
	r := CustomDirectoryRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCustomDirectoryRole_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_directory_role", "test")
	r := CustomDirectoryRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCustomDirectoryRole_disable(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_directory_role", "test")
	r := CustomDirectoryRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.disabled(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCustomDirectoryRole_templateId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_custom_directory_role", "test")
	r := CustomDirectoryRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.templateId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r CustomDirectoryRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.RoleDefinitionsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	role, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Custom Directory Role with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Custom Directory Role with object ID %q: %+v", state.ID, err)
	}

	return utils.Bool(role.ID() != nil && *role.ID() == state.ID), nil
}

func (r CustomDirectoryRoleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_custom_directory_role" "test" {
  display_name = "acctestCustomRole-%[1]d"
  enabled      = true
  version      = "1.0"

  permissions {
    allowed_resource_actions = ["microsoft.directory/applications/standard/read"]
  }
}
`, data.RandomInteger)
}

func (r CustomDirectoryRoleResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_custom_directory_role" "test" {
  display_name = "acctestCustomRoleComplete-%[1]d"
  description  = "test role for testing"
  enabled      = true
  version      = "v1.5"

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/applications/basic/update",
      "microsoft.directory/applications/create",
      "microsoft.directory/applications/standard/read",
    ]
  }

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/basic/update",
      "microsoft.directory/groups/create",
      "microsoft.directory/groups/delete",
    ]
  }
}
`, data.RandomInteger)
}

func (r CustomDirectoryRoleResource) disabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_custom_directory_role" "test" {
  display_name = "acctestCustomRole-%[1]d"
  enabled      = false
  version      = "1.0"

  permissions {
    allowed_resource_actions = ["microsoft.directory/applications/standard/read"]
  }
}
`, data.RandomInteger)
}

func (r CustomDirectoryRoleResource) templateId(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_custom_directory_role" "test" {
  display_name = "acctestCustomRole-%[1]d"
  enabled      = true
  template_id  = "%[2]s"
  version      = "1.0"

  permissions {
    allowed_resource_actions = ["microsoft.directory/applications/standard/read"]
  }
}
`, data.RandomInteger, data.RandomID)
}
