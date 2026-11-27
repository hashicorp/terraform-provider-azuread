---
page_title: "Resource: azuread_conditional_access_policy"
subcategory: "Conditional Access"
description: |-
  Manages a Conditional Access Policy within Azure Active Directory.
  -> Licensing Requirements Specifying client_applications property requires the activation of Microsoft Entra on your tenant and the availability of sufficient Workload Identities Premium licences (one per service principal managed by a conditional access).
  -> API Limits This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to reduce parallelism https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options or specify extended custom resource timeouts https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts.
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

## Arguments Reference

The following arguments are supported:

* `conditions` - (Required) The `conditions` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73))
* `display_name` - (Required) The friendly name for this Conditional Access Policy. Value must not be empty.
* `state` - (Required) Specifies the state of the policy object. Possible values include `disabled`, `enabled`, and `enabledForReportingButNotEnforced`.

---

* `grant_controls` - (Optional) The `grant_controls` block as defined below. At least one of `grant_controls` or `session_controls` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-6772616e745f636f6e74726f6c73))
* `session_controls` - (Optional) The `session_controls` block as defined below. At least one of `grant_controls` or `session_controls` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-73657373696f6e5f636f6e74726f6c73))

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73"></a>
The `conditions` block supports the following:

* `applications` - (Required) The `applications` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73))
* `client_app_types` - (Required) A list of client application types included in the policy. For each element: Possible values include `all`, `browser`, `easSupported`, `exchangeActiveSync`, `mobileAppsAndDesktopClients`, and `other`.
* `users` - (Required) The `users` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273))
* `authentication_flow_transfer_methods` - (Optional) A list of authentication flow transfer methods included in the policy. For each element: Possible values include `authenticationTransfer` and `deviceCodeFlow`.
* `client_applications` - (Optional) The `client_applications` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73))
* `devices` - (Optional) A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573))
* `insider_risk_levels` - (Optional) The insider risk level in the policy. Possible values include `elevated`, `minor`, and `moderate`.
* `locations` - (Optional) A `locations` block as documented below, which specifies locations included in and excluded from the policy. In the Microsoft Entra admin center this condition was renamed from **Location** to **Network**; this is a portal label change only and the underlying API is unchanged. See [Conditional Access Policy: Using Network Signals](https://learn.microsoft.com/en-us/entra/identity/conditional-access/concept-assignment-network) for more information. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6c6f636174696f6e73))

~> **Note:** In the Microsoft Entra admin center, the **Location** condition was renamed to **Network** and now also surfaces at the assignment level. This is a portal label change only - the Microsoft Graph API and this provider continue to use the `locations` naming, and existing configurations continue to work unchanged. A "Global Secure Access" compliant network location (shown as **All Compliant Network locations** in the portal) is referenced like any other location, by supplying its location ID in `included_locations` or `excluded_locations`. See [Conditional Access Policy: Using Network Signals](https://learn.microsoft.com/en-us/entra/identity/conditional-access/concept-assignment-network) for more information.
* `platforms` - (Optional) A `platforms` block as documented below, which specifies platforms included in and excluded from the policy. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-706c6174666f726d73))
* `service_principal_risk_levels` - (Optional) A list of service principal sign-in risk levels included in the policy. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.
* `sign_in_risk_levels` - (Optional) A list of user sign-in risk levels included in the policy. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.
* `user_risk_levels` - (Optional) A list of user risk levels included in the policy. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73"></a>
The `conditions.applications` block supports the following:

* `excluded_applications` - (Optional) A list of application IDs explicitly excluded from the policy. Can also be set to `Office365`. For each element: Value must not be empty.
* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73-66696c746572))

~> **Note:** Specifying `filter` requires the `Attribute Definition Reader` role, this is not included in the `Global Administrator` or other administrator roles and must be separately assigned.
* `included_applications` - (Optional) A list of application IDs the policy applies to, unless explicitly excluded (in `excluded_applications`). Can also be set to `All`, `None` or `Office365`. For each element: Value must not be empty. Exactly one of `conditions.0.applications.0.included_applications` or `conditions.0.applications.0.included_user_actions` must be specified.
* `included_user_actions` - (Optional) A list of user actions to include. For each element: Value must not be empty. Exactly one of `conditions.0.applications.0.included_applications` or `conditions.0.applications.0.included_user_actions` must be specified.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73-66696c746572"></a>
The `conditions.applications.filter` block supports the following:

* `mode` - (Required) Whether to include in, or exclude from, matching items from the policy. Possible values include `exclude` and `include`.
* `rule` - (Required) Condition filter to match items. Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73"></a>
The `conditions.client_applications` block supports the following:

* `excluded_service_principals` - (Optional) A list of service principal IDs explicitly excluded in the policy. For each element: Value must not be empty.
* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73-66696c746572))

~> **Note:** Specifying `filter` requires the `Attribute Definition Reader` role, this is not included in the `Global Administrator` or other administrator roles and must be separately assigned.
* `included_service_principals` - (Optional) A list of service principal IDs explicitly included in the policy. Can be set to `ServicePrincipalsInMyTenant` to include all service principals. This is mandatory value when at least one `excluded_service_principals` is set. For each element: Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73-66696c746572"></a>
The `conditions.client_applications.filter` block supports the following:

* `mode` - (Required) Whether to include in, or exclude from, matching items from the policy. Possible values include `exclude` and `include`.
* `rule` - (Required) Condition filter to match items. Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573"></a>
The `conditions.devices` block supports the following:

* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573-66696c746572))

-> **Note:** For more information on device filters, see the [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573-66696c746572"></a>
The `conditions.devices.filter` block supports the following:

* `mode` - (Required) Whether to include in, or exclude from, matching items from the policy. Possible values include `exclude` and `include`.
* `rule` - (Required) Condition filter to match items. Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6c6f636174696f6e73"></a>
The `conditions.locations` block supports the following:

* `included_locations` - (Required) A list of location IDs in scope of policy unless explicitly excluded. Can also be set to `All`, or `AllTrusted`. For each element: Value must not be empty.
* `excluded_locations` - (Optional) A list of location IDs excluded from scope of policy. Can also be set to `AllTrusted`. For each element: Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-706c6174666f726d73"></a>
The `conditions.platforms` block supports the following:

* `included_platforms` - (Required) A list of platforms the policy applies to, unless explicitly excluded. For each element: Possible values include `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, and `windowsPhone`.
* `excluded_platforms` - (Optional) A list of platforms explicitly excluded from the policy. For each element: Possible values include `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, and `windowsPhone`.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273"></a>
The `conditions.users` block supports the following:

* `excluded_groups` - (Optional) A list of group IDs excluded from scope of policy. For each element: Value must not be empty.
* `excluded_guests_or_external_users` - (Optional) The `excluded_guests_or_external_users` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273))
* `excluded_roles` - (Optional) A list of role template IDs excluded from scope of policy. For each element: Value must not be empty.
* `excluded_users` - (Optional) A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`. For each element: Value must not be empty.
* `included_groups` - (Optional) A list of group IDs in scope of policy unless explicitly excluded. For each element: Value must not be empty. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
* `included_guests_or_external_users` - (Optional) The `included_guests_or_external_users` block as defined below. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273))
* `included_roles` - (Optional) A list of role template IDs in scope of policy unless explicitly excluded. For each element: Value must not be empty. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
* `included_users` - (Optional) A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`. For each element: Value must not be empty. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273"></a>
The `conditions.users.excluded_guests_or_external_users` block supports the following:

* `guest_or_external_user_types` - (Required) A list of guest or external user types. For each element: Possible values include `b2bCollaborationGuest`, `b2bCollaborationMember`, `b2bDirectConnectUser`, `internalGuest`, `none`, `otherExternalUser`, and `serviceProvider`.
* `external_tenants` - (Optional) The `external_tenants` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473))

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473"></a>
The `conditions.users.excluded_guests_or_external_users.external_tenants` block supports the following:

* `membership_kind` - (Required) The external tenant membership kind. Possible values include `all` and `enumerated`.
* `members` - (Optional) A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`. For each element: Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273"></a>
The `conditions.users.included_guests_or_external_users` block supports the following:

* `guest_or_external_user_types` - (Required) A list of guest or external user types. For each element: Possible values include `b2bCollaborationGuest`, `b2bCollaborationMember`, `b2bDirectConnectUser`, `internalGuest`, `none`, `otherExternalUser`, and `serviceProvider`.
* `external_tenants` - (Optional) The `external_tenants` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473))

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473"></a>
The `conditions.users.included_guests_or_external_users.external_tenants` block supports the following:

* `membership_kind` - (Required) The external tenant membership kind. Possible values include `all` and `enumerated`.
* `members` - (Optional) A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`. For each element: Value must not be empty.

---

<a id="schema-6e6573746564-696e707574-6772616e745f636f6e74726f6c73"></a>
The `grant_controls` block supports the following:

* `operator` - (Required) Defines the relationship of the grant controls. Possible values include `AND` and `OR`.
* `authentication_strength_policy_id` - (Optional) ID of an Authentication Strength Policy to use in this policy. When using a hard-coded ID, the UUID value should be prefixed with: `/policies/authenticationStrengthPolicies/`. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.
* `built_in_controls` - (Optional) List of built-in controls required by the policy. For each element: Possible values include `approvedApplication`, `block`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `mfa`, and `passwordChange`. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.
* `custom_authentication_factors` - (Optional) List of custom controls IDs required by the policy. For each element: Value must not be empty.
* `terms_of_use` - (Optional) List of terms of use IDs required by the policy. For each element: Value must not be empty. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.

---

<a id="schema-6e6573746564-696e707574-73657373696f6e5f636f6e74726f6c73"></a>
The `session_controls` block supports the following:

* `application_enforced_restrictions_enabled` - (Optional) Whether application enforced restrictions are enabled.

-> Only Office 365, Exchange Online and Sharepoint Online support application enforced restrictions.
* `cloud_app_security_policy` - (Optional) Enables cloud app security and specifies the cloud app security policy to use. Possible values include `blockDownloads`, `mcasConfigured`, and `monitorOnly`.
* `disable_resilience_defaults` - (Optional) Disables [resilience defaults](https://learn.microsoft.com/en-us/azure/active-directory/conditional-access/resilience-defaults).
* `persistent_browser_mode` - (Optional) Session control to define whether to persist cookies. Possible values include `always` and `never`.
* `sign_in_frequency` - (Optional) Number of days or hours to enforce sign-in frequency. Value must be at least `0`. Required when using `session_controls.0.sign_in_frequency_period`.
* `sign_in_frequency_authentication_type` - (Optional) Authentication type for enforcing sign-in frequency. Possible values include `primaryAndSecondaryAuthentication` and `secondaryAuthentication`.
* `sign_in_frequency_interval` - (Optional) The interval to apply to sign-in frequency control. Possible values include `everyTime` and `timeBased`.
* `sign_in_frequency_period` - (Optional) The time period to enforce sign-in frequency. Possible values include `days` and `hours`. Required when using `session_controls.0.sign_in_frequency`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `object_id` - The object ID of the policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 15 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_conditional_access_policy.my_location /identity/conditionalAccess/policies/00000000-0000-0000-0000-000000000000
```
