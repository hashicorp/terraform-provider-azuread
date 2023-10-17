// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationOwnerResource struct{}

func TestAccApplicationOwner_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_owner", "test")
	r := ApplicationOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("owner_object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationOwner_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_owner", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_owner", "test2")
	r := ApplicationOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("owner_object_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("owner_object_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationOwner_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_owner", "test")
	r := ApplicationOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("owner_object_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationOwnerResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ParseOwnerID(state.ID)
	if err != nil {
		return nil, err
	}

	result, status, err := client.GetOwner(ctx, id.ApplicationId, id.OwnerID)
	if err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}
	if result == nil {
		return nil, fmt.Errorf("retrieving %s: result was nil", id)
	}

	return pointer.To(true), nil
}

func (ApplicationOwnerResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application_registration" "test" {
  display_name = "acctest-Owner-%[1]d"
}

resource "azuread_user" "test" {
  user_principal_name = "acctestAppOwner.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAppOwner-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application_owner" "test" {
  application_id  = azuread_application_registration.test.id
  owner_object_id = azuread_user.test.object_id
}
`, data.RandomInteger, data.RandomPassword)
}

func (ApplicationOwnerResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_user" "test" {
  user_principal_name = "acctestAppOwner.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAppOwner-%[1]d"
  password            = "%[2]s"
}

resource "azuread_user" "test2" {
  user_principal_name = "acctestAppOwner2.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAppOwner2-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application_owner" "test" {
  application_id  = azuread_application_registration.test.id
  owner_object_id = azuread_user.test.object_id
}

resource "azuread_application_owner" "test2" {
  application_id  = azuread_application_registration.test.id
  owner_object_id = azuread_user.test2.object_id
}
`, data.RandomInteger, data.RandomPassword)
}

func (ApplicationOwnerResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_user" "test" {
  user_principal_name = "acctestAppOwner.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAppOwner-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application_owner" "test" {
  application_id  = azuread_application_registration.test.id
  owner_object_id = azuread_user.test.object_id
}

resource "azuread_application_owner" "import" {
  application_id  = azuread_application_owner.test.application_id
  owner_object_id = azuread_application_owner.test.owner_object_id
}
`, data.RandomInteger, data.RandomPassword)
}
