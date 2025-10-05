---
subcategory: "Conditional Access"
---

# Resource: azuread_conditional_access_policy

Manages a Conditional Access Policy within Azure Active Directory.

-> **Licensing Requirements** Specifying `client_applications` property requires the activation of Microsoft Entra on your tenant and the availability of sufficient Workload Identities Premium licences (one per service principal managed by a conditional access).

-> **API Limits** This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to [reduce parallelism](https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options) or specify extended [custom resource timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

## Example Usage

### All users except guests or external users

```terraform
resource "azuread_conditional_access_policy" "example" {
  display_name = "example policy"
  state        = "disabled"

  conditions {
    client_app_types    = ["all"]
    sign_in_risk_levels = ["medium"]
    user_risk_levels    = ["medium"]

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

  session_controls {
    application_enforced_restrictions_enabled = true
    disable_resilience_defaults               = false
    sign_in_frequency                         = 10
    sign_in_frequency_period                  = "hours"
    cloud_app_security_policy                 = "monitorOnly"
  }
}
```

### Included client applications / service principals

```terraform

data "azuread_client_config" "current" {}

resource "azuread_conditional_access_policy" "example" {
  display_name = "example policy"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_applications = ["All"]
    }

    client_applications {
      included_service_principals = [data.azuread_client_config.current.object_id]
      excluded_service_principals = []
    }

    users {
      included_users = ["None"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
```

### Excluded client applications / service principals

```terraform

data "azuread_client_config" "current" {}

resource "azuread_conditional_access_policy" "example" {
  display_name = "example policy"
  state        = "disabled"

  conditions {
    client_app_types = ["all"]

    applications {
      included_applications = ["All"]
    }

    client_applications {
      included_service_principals = ["ServicePrincipalsInMyTenant"]
      excluded_service_principals = [data.azuread_client_config.current.object_id]
    }

    users {
      included_users = ["None"]
    }
  }

  grant_controls {
    operator          = "OR"
    built_in_controls = ["block"]
  }
}
```
## Argument Reference

The following arguments are supported:

* `conditions` - (Required) A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
* `display_name` - (Required) The friendly name for this Conditional Access Policy.
* `grant_controls` - (Optional) A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
* `session_controls` - (Optional) A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.

~> Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.

* `state` - (Required) Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`

---

`conditions` block supports the following:

* `applications` - (Required) An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
* `client_app_types` - (Required) A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
* `client_applications` - (Optional) An `client_applications` block as documented below, which specifies service principals included in and excluded from the policy.
* `devices` - (Optional) A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
* `insider_risk_levels` - (Optional) The insider risk level in the policy. Possible values are: `minor`, `moderate`, `elevated`, `unknownFutureValue`.
* `locations` - (Optional) A `locations` block as documented below, which specifies locations included in and excluded from the policy.
* `platforms` - (Optional) A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
* `service_principal_risk_levels` - (Optional) A list of service principal sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `none`, `unknownFutureValue`.
* `sign_in_risk_levels` - (Optional) A list of user sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
* `user_risk_levels` - (Optional) A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
* `users` - (Required) A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.

---

`applications` block supports the following:

* `excluded_applications` - (Optional) A list of application IDs explicitly excluded from the policy. Can also be set to `Office365`.
* `included_applications` - (Optional) A list of application IDs the policy applies to, unless explicitly excluded (in `excluded_applications`). Can also be set to `All`, `None` or `Office365`. Cannot be specified with `included_user_actions`. One of `included_applications` or `included_user_actions` must be specified.
* `included_user_actions` - (Optional) A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `included_applications`. One of `included_applications` or `included_user_actions` must be specified.

---

`client_applications` block supports the following:

* `excluded_service_principals` - (Optional) A list of service principal IDs explicitly excluded in the policy.
* `filter` - (Optional) A `filter` block as documented below.

~> **Note:** Specifying `filter` requires the `Attribute Definition Reader` role, this is not included in the `Global Administrator` or other administrator roles and must be separately assigned.

* `included_service_principals` - (Optional) A list of service principal IDs explicitly included in the policy. Can be set to `ServicePrincipalsInMyTenant` to include all service principals. This is mandatory value when at least one `excluded_service_principals` is set.

---

`devices` block supports the following:

* `filter` - (Optional) A `filter` block as documented below.

-> **Note:** For more information on device filters, see the [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).

---

`filter` block supports the following:

* `mode` - (Required) Whether to include in, or exclude from, matching items from the policy. Supported values are `include` or `exclude`.
* `rule` - (Required) Condition filter to match items.

---

`users` block supports the following:

* `excluded_groups` - (Optional) A list of group IDs excluded from scope of policy.
* `excluded_guests_or_external_users` - (Optional) A `guests_or_external_users` block as documented below, which specifies internal guests and external users excluded from scope of policy.
* `excluded_roles` - (Optional) A list of role IDs excluded from scope of policy.
* `excluded_users` - (Optional) A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
* `included_groups` - (Optional) A list of group IDs in scope of policy unless explicitly excluded.
* `included_guests_or_external_users` - (Optional) A `guests_or_external_users` block as documented below, which specifies internal guests and external users in scope of policy.
* `included_roles` - (Optional) A list of role IDs in scope of policy unless explicitly excluded.
* `included_users` - (Optional) A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.

-> At least one of `included_groups`, `included_guests_or_external_users`, `included_roles` or `included_users` must be specified.

---

`guests_or_external_users` block supports the following:

* `external_tenants` - (Optional) An `external_tenants` block as documented below, which specifies external tenants in a policy scope.
* `guest_or_external_user_types` - (Required) A list of guest or external user types. Possible values are: `b2bCollaborationGuest`, `b2bCollaborationMember`, `b2bDirectConnectUser`, `internalGuest`, `none`, `otherExternalUser`, `serviceProvider`, `unknownFutureValue`.

---

`external_tenants` block supports the following:

* `members` - (Optional) A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
* `membership_kind` - (Required) The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.

---

`locations` block supports the following:

* `excluded_locations` - (Optional) A list of location IDs excluded from scope of policy. Can also be set to `AllTrusted`.
* `included_locations` - (Required) A list of location IDs in scope of policy unless explicitly excluded. Can also be set to `All`, or `AllTrusted`.

---

`platforms` block supports the following:

* `excluded_platforms` - (Optional) A list of platforms explicitly excluded from the policy. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.
* `included_platforms` - (Required) A list of platforms the policy applies to, unless explicitly excluded. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.

---

`grant_controls` block supports the following:

* `authentication_strength_policy_id` - (Optional) ID of an Authentication Strength Policy to use in this policy. When using a hard-coded ID, the UUID value should be prefixed with: `/policies/authenticationStrengthPolicies/`.
* `built_in_controls` - (Optional) List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
* `custom_authentication_factors` - (Optional) List of custom controls IDs required by the policy.
* `operator` - (Required) Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
* `terms_of_use` - (Optional) List of terms of use IDs required by the policy.

-> At least one of `authentication_strength_policy_id`, `built_in_controls` or `terms_of_use` must be specified.

---

`session_controls` block supports the following:

* `application_enforced_restrictions_enabled` - (Optional) Whether application enforced restrictions are enabled. Defaults to `false`.

-> Only Office 365, Exchange Online and Sharepoint Online support application enforced restrictions.

* `cloud_app_security_policy` - (Optional) Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
* `disable_resilience_defaults` - (Optional) Disables [resilience defaults](https://learn.microsoft.com/en-us/azure/active-directory/conditional-access/resilience-defaults). Defaults to `false`.
* `persistent_browser_mode` - (Optional) Session control to define whether to persist cookies. Possible values are: `always` or `never`.
* `sign_in_frequency` - (Optional) Number of days or hours to enforce sign-in frequency. Required when `sign_in_frequency_period` is specified.
* `sign_in_frequency_authentication_type` - (Optional) Authentication type for enforcing sign-in frequency. Possible values are: `primaryAndSecondaryAuthentication` or `secondaryAuthentication`. Defaults to `primaryAndSecondaryAuthentication`.
* `sign_in_frequency_interval` - (Optional) The interval to apply to sign-in frequency control. Possible values are: `timeBased` or `everyTime`. Defaults to `timeBased`.
* `sign_in_frequency_period` - (Optional) The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `sign_in_frequency_period` is specified.

---

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the Conditional Access Policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 15 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Conditional Access Policies can be imported using the `id`, e.g.

```shell
terraform import azuread_conditional_access_policy.my_location /identity/conditionalAccess/policies/00000000-0000-0000-0000-000000000000
```
