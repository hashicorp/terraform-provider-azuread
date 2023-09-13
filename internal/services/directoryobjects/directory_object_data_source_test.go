// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryobjects_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type PrincipalTypeDataSource struct{}

func TestAccPrincipalTypeDataSource_groupByObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_object", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: PrincipalTypeDataSource{}.objectTypeFromGroup(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("type").HasValue("Group"),
			),
		},
	})
}

func TestAccPrincipalTypeDataSource_userByObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_object", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: PrincipalTypeDataSource{}.objectTypeFromUser(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("type").HasValue("User"),
			),
		},
	})
}

func TestAccPrincipalTypeDataSource_servicePrincipalByObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_object", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: PrincipalTypeDataSource{}.objectTypeFromServicePrincipal(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("type").HasValue("ServicePrincipal"),
			),
		},
	})
}

func (PrincipalTypeDataSource) basicGroup(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  display_name     = "acctestGroup-%d"
  security_enabled = true
}
`, data.RandomInteger)
}

func (PrincipalTypeDataSource) basicUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword)
}

func (PrincipalTypeDataSource) basicServicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (PrincipalTypeDataSource) objectTypeFromGroup(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_directory_object" "test" {
  object_id = azuread_group.test.object_id
}
`, PrincipalTypeDataSource{}.basicGroup(data))
}

func (PrincipalTypeDataSource) objectTypeFromUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_directory_object" "test" {
  object_id = azuread_user.test.object_id
}
`, PrincipalTypeDataSource{}.basicUser(data))
}

func (PrincipalTypeDataSource) objectTypeFromServicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_directory_object" "test" {
  object_id = azuread_service_principal.test.object_id
}
`, PrincipalTypeDataSource{}.basicServicePrincipal(data))
}
