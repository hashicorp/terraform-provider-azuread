// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type ApplicationPublishedAppIdsDataSource struct{}

func TestAccApplicationPublishedAppIdsDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application_published_app_ids", "test")
	r := ApplicationPublishedAppIdsDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("result.%").Exists(),
			),
		},
	})
}

func (ApplicationPublishedAppIdsDataSource) basic(data acceptance.TestData) string {
	return `provider azuread {}
data "azuread_application_published_app_ids" "test" {}`
}
