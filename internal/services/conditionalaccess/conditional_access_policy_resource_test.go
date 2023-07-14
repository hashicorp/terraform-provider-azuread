// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess_test

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

type ConditionalAccessPolicyResource struct{}

func TestAccConditionalAccessPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

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

func TestAccConditionalAccessPolicy_deviceFilter(t *testing.T) {
	// This is a separate test for two reasons:
	// 1. To accommodate future properties included_devices/excluded_devices which both conflict with devices.0.filter
	// 2. Because devices.0.filter is conditionally ForceNew, as the API ignores removal of this property

	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.deviceFilter(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
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
			Config: r.deviceFilter(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.deviceFilterUpdate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_includedUserActions(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.includedUserActions(data),
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
		{
			Config: r.includedUserActions(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_sessionControls(t *testing.T) {
	// This is in a separate test to avoid ForceNew in the update test due to https://github.com/microsoftgraph/msgraph-metadata/issues/93
	// session_controls can be added to the complete config, and this rest removed, when this issue is resolved

	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.sessionControls(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
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
			Config: r.sessionControls(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsUpdate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControls(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
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
		{
			Config: r.sessionControls(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_sessionControlsDisabled(t *testing.T) {
	// Remove this test when https://github.com/microsoftgraph/msgraph-metadata/issues/93 is resolved

	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.sessionControlsDisabled(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
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
			Config: r.sessionControlsDisabled(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_clientApplications(t *testing.T) {
	// This is a separate test for two reasons:
	// - conditional access policies applies either to users/groups or to client applications (workload identities)
	// - conditional access policies using client applications requires special licensing (Microsoft Entra Workload Identities)

	// Due to eventual consistency issues making it difficult to create a service principal on demand for inclusion in this
	// test policy, the config for this test requires a pre-existing service principal named "Terraform Acceptance Tests (Single Tenant)"
	// which should be linked to a single tenant application in the same tenant.

	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.clientApplicationsIncluded(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.clientApplicationsExcluded(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.clientApplicationsIncluded(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func (r ConditionalAccessPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	var id *string

	app, status, err := clients.ConditionalAccess.PoliciesClient.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Conditional Access Policy with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Conditional Access Policy with ID %q: %+v", state.ID, err)
	}
	id = app.ID

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (ConditionalAccessPolicyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["None"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types    = ["all"]
    sign_in_risk_levels = ["medium"]
    user_risk_levels    = ["medium"]

    applications {
      included_applications = ["All"]
      excluded_applications = ["00000004-0000-0ff1-ce00-000000000000"]
    }

    locations {
      included_locations = ["All"]
      excluded_locations = ["AllTrusted"]
    }

    platforms {
      included_platforms = ["android"]
      excluded_platforms = ["iOS"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["mfa"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) deviceFilter(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["All"]
    }

    devices {
      filter {
        mode = "exclude"
        rule = "device.operatingSystem eq \"Doors\""
      }
    }

    locations {
      included_locations = ["All"]
    }

    platforms {
      included_platforms = ["all"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) deviceFilterUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["All"]
    }

    devices {
      filter {
        mode = "exclude"
        rule = "device.model eq \"yPhone Z\""
      }
    }

    locations {
      included_locations = ["All"]
    }

    platforms {
      included_platforms = ["all"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }

}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) includedUserActions(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_user_actions = [
        "urn:user:registerdevice",
        "urn:user:registersecurityinfo",
      ]
    }

    locations {
      included_locations = ["All"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["mfa"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControls(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["All"]
    }

    locations {
      included_locations = ["All"]
    }

    platforms {
      included_platforms = ["all"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }

  session_controls {
    application_enforced_restrictions_enabled = true
    disable_resilience_defaults               = true
    cloud_app_security_policy                 = "monitorOnly"
    persistent_browser_mode                   = "never"
    sign_in_frequency                         = 10
    sign_in_frequency_period                  = "hours"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["All"]
    }

    locations {
      included_locations = ["All"]
    }

    platforms {
      included_platforms = ["all"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }

  session_controls {
    application_enforced_restrictions_enabled = true
    disable_resilience_defaults               = false
    cloud_app_security_policy                 = "blockDownloads"
    persistent_browser_mode                   = "always"
    sign_in_frequency                         = 2
    sign_in_frequency_period                  = "days"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsDisabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["browser"]

    applications {
      included_applications = ["All"]
    }

    locations {
      included_locations = ["All"]
    }

    platforms {
      included_platforms = ["all"]
    }

    users {
      included_users = ["All"]
      excluded_users = ["GuestsOrExternalUsers"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }

  session_controls {
    application_enforced_restrictions_enabled = false
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) clientApplicationsIncluded(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_service_principal" "test" {
  display_name = "Terraform Acceptance Tests (Single Tenant)"
}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_applications = ["All"]
    }

    client_applications {
      included_service_principals = [data.azuread_service_principal.test.object_id]
    }

	service_principal_risk_levels = ["medium"]

    users {
      included_users = ["None"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) clientApplicationsExcluded(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_service_principal" "test" {
  display_name = "Terraform Acceptance Tests (Single Tenant)"
}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_applications = ["All"]
    }

    client_applications {
      included_service_principals = ["ServicePrincipalsInMyTenant"]
      excluded_service_principals = [data.azuread_service_principal.test.object_id]
    }

	service_principal_risk_levels = ["medium"]

    users {
      included_users = ["None"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}
