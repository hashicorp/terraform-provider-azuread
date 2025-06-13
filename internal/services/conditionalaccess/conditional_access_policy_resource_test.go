// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesspolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type ConditionalAccessPolicyResource struct{}

func TestAccConditionalAccessPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_signInFrequencyEveryTime(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.signinfrequencyintervalEverytime(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_signInFrequencyEveryTimeShouldFail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.signinfrequencyintervalEverytimeShouldFail(data),
			PlanOnly:    true,
			ExpectError: regexp.MustCompile("when `session_controls.sign_in_frequency_interval` is set to"),
		},
	})
}

func TestAccConditionalAccessPolicy_signInFrequencyEveryTimeUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.signinfrequencyintervalEverytime(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("enabledForReportingButNotEnforced"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_includedUserActions(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.includedUserActions(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.includedUserActions(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_sessionControls(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.sessionControls(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControls(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControls(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsDisabled(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsDisabled(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsApplicationEnforcedRestrictions(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsCloudAppSecurityPolicy(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsPersistentBrowserMode(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.sessionControlsDisabled(data),
			Check: acceptance.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.clientApplicationsIncluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.clientApplicationsExcluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
		{
			Config: r.clientApplicationsIncluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_authenticationStrength(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.authenticationStrengthPolicy(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("grant_controls.0.authentication_strength_policy_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_authenticationStrengthHardcoded(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.authenticationStrengthPolicyHardcoded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("grant_controls.0.authentication_strength_policy_id").HasValue("/policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000004"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_guestsOrExternalUsers(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.guestsOrExternalUsersAllServiceProvidersIncluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("conditions.0.users.0.included_guests_or_external_users.0.external_tenants.0.membership_kind").HasValue("all"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.guestsOrExternalUsersAllServiceProvidersExcluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("conditions.0.users.0.excluded_guests_or_external_users.0.external_tenants.0.membership_kind").HasValue("all"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_insiderRisk(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.insiderRisk(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("state").HasValue("disabled"),
				check.That(data.ResourceName).Key("conditions.0.insider_risk_levels").HasValue("moderate"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccConditionalAccessPolicy_guestsOrExternalUsersServiceProviderExternalTenantExcluded(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_conditional_access_policy", "test")
	r := ConditionalAccessPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.guestsOrExternalUsersServiceProviderExternalTenantExcluded(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-CONPOLICY-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("conditions.0.users.0.excluded_guests_or_external_users.0.external_tenants.0.membership_kind").HasValue("enumerated"),
				check.That(data.ResourceName).Key("conditions.0.users.0.excluded_guests_or_external_users.0.external_tenants.0.members.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func (r ConditionalAccessPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := stable.ParseIdentityConditionalAccessPolicyID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.ConditionalAccess.PolicyClient.GetConditionalAccessPolicy(ctx, *id, conditionalaccesspolicy.DefaultGetConditionalAccessPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(true), nil
}

func (ConditionalAccessPolicyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
provider "azuread" {}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "enabledForReportingButNotEnforced"

  conditions {
    client_app_types    = ["all"]
    sign_in_risk_levels = ["medium"]
    user_risk_levels    = ["medium"]
    insider_risk_levels = "elevated"

    applications {
      included_applications = ["All"]
      excluded_applications = []
    }

    devices {
      filter {
        mode = "exclude"
        rule = "device.operatingSystem eq \"Doors\""
      }
    }

    locations {
      included_locations = ["All"]
      excluded_locations = ["AllTrusted"]
    }

    platforms {
      included_platforms = ["all"]
      excluded_platforms = ["android", "iOS"]
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

  session_controls {
    application_enforced_restrictions_enabled = true
    cloud_app_security_policy                 = "blockDownloads"
    disable_resilience_defaults               = false
    persistent_browser_mode                   = "always"
    sign_in_frequency                         = 2
    sign_in_frequency_authentication_type     = "primaryAndSecondaryAuthentication"
    sign_in_frequency_interval                = "timeBased"
    sign_in_frequency_period                  = "days"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) includedUserActions(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
provider "azuread" {}

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

  session_controls {
    application_enforced_restrictions_enabled = true
    disable_resilience_defaults               = true
    cloud_app_security_policy                 = "monitorOnly"
    persistent_browser_mode                   = "never"
    sign_in_frequency                         = 10
    sign_in_frequency_authentication_type     = "primaryAndSecondaryAuthentication"
    sign_in_frequency_interval                = "timeBased"
    sign_in_frequency_period                  = "hours"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsDisabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsApplicationEnforcedRestrictions(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsCloudAppSecurityPolicy(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
    cloud_app_security_policy = "monitorOnly"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) sessionControlsPersistentBrowserMode(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
    persistent_browser_mode = "always"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) clientApplicationsIncluded(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
provider "azuread" {}

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

func (ConditionalAccessPolicyResource) authenticationStrengthPolicy(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["password"]
}

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
    operator                          = "OR"
    authentication_strength_policy_id = azuread_authentication_strength_policy.test.id
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) authenticationStrengthPolicyHardcoded(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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

  # Hard-code the Phishing resistant MFA policy
  grant_controls {
    operator                          = "OR"
    authentication_strength_policy_id = "/policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000004"
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) guestsOrExternalUsersAllServiceProvidersIncluded(data acceptance.TestData) string {
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
      included_guests_or_external_users {
        guest_or_external_user_types = ["internalGuest", "serviceProvider"]
        external_tenants {
          membership_kind = "all"
        }
      }
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) guestsOrExternalUsersAllServiceProvidersExcluded(data acceptance.TestData) string {
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
      included_users = ["None"]
      excluded_guests_or_external_users {
        guest_or_external_user_types = ["internalGuest", "serviceProvider"]
        external_tenants {
          membership_kind = "all"
        }
      }
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) insiderRisk(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "disabled"

  conditions {
    client_app_types    = ["browser"]
    insider_risk_levels = "moderate"

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

func (ConditionalAccessPolicyResource) guestsOrExternalUsersServiceProviderExternalTenantExcluded(data acceptance.TestData) string {
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
      included_users = ["None"]
      excluded_guests_or_external_users {
        guest_or_external_user_types = ["serviceProvider"]
        external_tenants {
          membership_kind = "enumerated"
          members = [
            "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
          ]
        }
      }
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) signinfrequencyintervalEverytime(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "enabledForReportingButNotEnforced"

  conditions {
    client_app_types = [
      "all",
    ]
    sign_in_risk_levels = []
    user_risk_levels = [
      "high",
    ]
    service_principal_risk_levels = []

    applications {
      excluded_applications = []
      included_applications = [
        "All",
      ]
    }

    users {
      excluded_groups = []
      excluded_roles  = []
      excluded_users  = []
      included_groups = []
      included_roles  = []
      included_users = [
        "None"
      ]
    }
  }

  grant_controls {
    built_in_controls = [
      "mfa",
      "passwordChange"
    ]
    custom_authentication_factors = []
    operator                      = "AND"
    terms_of_use                  = []
  }

  session_controls {
    cloud_app_security_policy             = null
    disable_resilience_defaults           = null
    persistent_browser_mode               = null
    sign_in_frequency_authentication_type = "primaryAndSecondaryAuthentication"
    sign_in_frequency_interval            = "everyTime" // NOTE: this precludes the use of sign_in_frequency_period etc
  }
}
`, data.RandomInteger)
}

func (ConditionalAccessPolicyResource) signinfrequencyintervalEverytimeShouldFail(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_conditional_access_policy" "test" {
  display_name = "acctest-CONPOLICY-%[1]d"
  state        = "enabledForReportingButNotEnforced"

  conditions {
    client_app_types = [
      "all",
    ]
    sign_in_risk_levels = []
    user_risk_levels = [
      "high",
    ]
    service_principal_risk_levels = []

    applications {
      excluded_applications = []
      included_applications = [
        "All",
      ]
    }

    users {
      excluded_groups = []
      excluded_roles  = []
      excluded_users  = []
      included_groups = []
      included_roles  = []
      included_users = [
        "None"
      ]
    }
  }

  grant_controls {
    built_in_controls = [
      "mfa",
      "passwordChange"
    ]
    custom_authentication_factors = []
    operator                      = "AND"
    terms_of_use                  = []
  }

  session_controls {
    cloud_app_security_policy             = null
    disable_resilience_defaults           = null
    persistent_browser_mode               = null
    sign_in_frequency_authentication_type = "primaryAndSecondaryAuthentication"
    sign_in_frequency_interval            = "everyTime" // NOTE: this precludes the use of sign_in_frequency_period etc
    sign_in_frequency_period              = "hours"     // This should fail because sign_in_frequency_interval is set to "everyTime"
    sign_in_frequency                     = 1           // This should fail because sign_in_frequency_interval is set to "everyTime"
  }
}
`, data.RandomInteger)
}
