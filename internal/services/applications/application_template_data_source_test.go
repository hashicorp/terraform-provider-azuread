// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

const (
	testApplicationTemplateDisplayName = "Marketo"                              // The display name of the template
	testApplicationTemplateId          = "4601ed45-8ff3-4599-8377-b6649007e876" // The template ID for the Marketo app template
	testApplicationTemplateAppRoleId   = "dfd0e7dd-26fb-4b2c-98d2-e444486c1e37" // The app role provided by the template
)

type ApplicationTemplateDataSource struct{}

func TestAccApplicationTemplateDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application_template", "test")
	r := ApplicationTemplateDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.byDisplayName(data),
			Check:  r.testCheck(data),
		},
	})
}

func TestAccApplicationTemplateDataSource_byTemplateId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application_template", "test")
	r := ApplicationTemplateDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.byTemplateId(data),
			Check:  r.testCheck(data),
		},
	})
}

func (ApplicationTemplateDataSource) testCheck(data acceptance.TestData) acceptance.TestCheckFunc {
	return acceptance.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("template_id").HasValue(testApplicationTemplateId),
		check.That(data.ResourceName).Key("display_name").HasValue(testApplicationTemplateDisplayName),
		check.That(data.ResourceName).Key("categories.#").Exists(),
		check.That(data.ResourceName).Key("homepage_url").Exists(),
		check.That(data.ResourceName).Key("logo_url").Exists(),
		check.That(data.ResourceName).Key("publisher").Exists(),
		check.That(data.ResourceName).Key("supported_provisioning_types.#").Exists(),
		check.That(data.ResourceName).Key("supported_single_sign_on_modes.#").Exists(),
	)
}

func (ApplicationTemplateDataSource) byTemplateId(_ acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application_template" "test" {
  template_id = "%[1]s"
}
`, testApplicationTemplateId)
}

func (ApplicationTemplateDataSource) byDisplayName(_ acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application_template" "test" {
  display_name = "%[1]s"
}
`, testApplicationTemplateDisplayName)
}
