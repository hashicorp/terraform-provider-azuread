// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationAppRoleResource struct{}

func TestAccApplicationAppRole_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationAppRole_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_app_role", "test2")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationAppRole_multipleUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_app_role", "test2")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
		{
			Config: r.multipleUpdate(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationAppRole_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationAppRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseAppRoleID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", applicationId, err)
	}

	app := resp.Model
	if app == nil {
		return nil, fmt.Errorf("retrieving %s: model was nil", applicationId)
	}

	if app.AppRoles == nil {
		return pointer.To(false), nil
	}

	for _, role := range *app.AppRoles {
		if strings.EqualFold(*role.Id, id.RoleID) {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (ApplicationAppRoleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_app_role" "test" {
  application_id = azuread_application_registration.test.id
  role_id        = "%[2]s"

  allowed_member_types = ["User"]
  description          = "Admins can manage roles and perform all task actions"
  display_name         = "Admin"
  value                = "admin"
}
`, data.RandomInteger, data.RandomID)
}

func (ApplicationAppRoleResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "random_uuid" "test" {}
resource "random_uuid" "test2" {}

resource "azuread_application_app_role" "test" {
  application_id = azuread_application_registration.test.id
  role_id        = random_uuid.test.id

  allowed_member_types = ["User"]
  description          = "Admins can manage roles and perform all task actions"
  display_name         = "Administrator"
  value                = "admin"
}

resource "azuread_application_app_role" "test2" {
  application_id = azuread_application_registration.test.id
  role_id        = random_uuid.test2.id

  allowed_member_types = ["User"]
  description          = "Read-Only roles have limited query access"
  display_name         = "Read-Only"
  value                = "user"
}
`, data.RandomInteger)
}

func (ApplicationAppRoleResource) multipleUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "random_uuid" "test" {}
resource "random_uuid" "test2" {}

resource "azuread_application_app_role" "test" {
  application_id = azuread_application_registration.test.id
  role_id        = random_uuid.test.id

  allowed_member_types = ["User"]
  description          = "Supers can manage roles and perform all task actions"
  display_name         = "SuperUser"
  value                = "super"
}

resource "azuread_application_app_role" "test2" {
  application_id = azuread_application_registration.test.id
  role_id        = random_uuid.test2.id

  allowed_member_types = ["User"]
  description          = "Users can look but not touch"
  display_name         = "User"
  value                = "luser"
}
`, data.RandomInteger)
}

func (ApplicationAppRoleResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_app_role" "test" {
  application_id = azuread_application_registration.test.id
  role_id        = "%[2]s"

  allowed_member_types = ["User"]
  description          = "Admins can manage roles and perform all task actions"
  display_name         = "Admin"
  value                = "admin"
}

resource "azuread_application_app_role" "import" {
  application_id = azuread_application_app_role.test.application_id
  role_id        = azuread_application_app_role.test.role_id

  allowed_member_types = azuread_application_app_role.test.allowed_member_types
  description          = azuread_application_app_role.test.description
  display_name         = azuread_application_app_role.test.display_name
  value                = azuread_application_app_role.test.value
}
`, data.RandomInteger, data.RandomID)
}
