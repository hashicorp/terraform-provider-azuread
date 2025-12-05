// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
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

type ApplicationOptionalClaimsResource struct{}

func TestAccApplicationOptionalClaims_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_optional_claims", "test")
	r := ApplicationOptionalClaimsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.applicationOnly(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That("azuread_application_registration.test").DoesNotExistInAzure(r),
			),
		},
	})
}

func TestAccApplicationOptionalClaims_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_optional_claims", "test")
	r := ApplicationOptionalClaimsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationOptionalClaims_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_optional_claims", "test")
	r := ApplicationOptionalClaimsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.applicationOnly(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That("azuread_application_registration.test").DoesNotExistInAzure(r),
			),
		},
	})
}

func TestAccApplicationOptionalClaims_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_optional_claims", "test")
	r := ApplicationOptionalClaimsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationOptionalClaimsResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseOptionalClaimsID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	app := resp.Model
	if app == nil {
		return nil, fmt.Errorf("retrieving %s: model was nil", id)
	}

	if claims := app.OptionalClaims; claims == nil {
		return pointer.To(false), nil
	} else if (claims.AccessToken == nil || len(*claims.AccessToken) == 0) &&
		(claims.IdToken == nil || len(*claims.IdToken) == 0) &&
		(claims.Saml2Token == nil || len(*claims.Saml2Token) == 0) {
		return pointer.To(false), nil
	}

	return pointer.To(true), nil
}

func (ApplicationOptionalClaimsResource) applicationOnly(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-OptionalClaims-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationOptionalClaimsResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-OptionalClaims-%[1]d"
}

resource "azuread_application_optional_claims" "test" {
  application_id = azuread_application_registration.test.id

  access_token {
    name = "myclaim"
  }

  id_token {
    name = "userclaim"
  }
}
`, data.RandomInteger)
}

func (ApplicationOptionalClaimsResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-OptionalClaims-%[1]d"
}

resource "azuread_application_optional_claims" "test" {
  application_id = azuread_application_registration.test.id

  access_token {
    name                  = "userclaim"
    source                = "user"
    essential             = true
    additional_properties = ["emit_as_roles"]
  }

  access_token {
    name      = "otherclaim"
    essential = false
  }

  id_token {
    name                  = "idclaim"
    source                = "user"
    essential             = true
    additional_properties = ["emit_as_roles"]
  }

  saml2_token {
    name      = "saml2claim"
    source    = "user"
    essential = true
    additional_properties = [
      "dns_domain_and_sam_account_name",
      "on_premise_security_identifier",
    ]
  }

  saml2_token {
    name                  = "saml2claim2"
    source                = "user"
    essential             = true
    additional_properties = ["netbios_domain_and_sam_account_name"]
  }
}
`, data.RandomInteger)
}

func (ApplicationOptionalClaimsResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-OptionalClaims-%[1]d"
}

resource "azuread_application_optional_claims" "test" {
  application_id = azuread_application_registration.test.id

  access_token {
    name = "myclaim"
  }
}

resource "azuread_application_optional_claims" "import" {
  application_id = azuread_application_optional_claims.test.application_id

  access_token {
    name = "myclaim"
  }
}
`, data.RandomInteger, data.RandomID)
}
