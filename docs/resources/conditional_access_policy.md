---
page_title: "Resource: azuread_conditional_access_policy"
subcategory: "Conditional Access"
description: |-
  Manages a Conditional Access Policy within Azure Active Directory. -> Licensing Requirements Specifying client_applications property requires the activation of Microsoft Entra on your tenant and the availability of sufficient Workload Identities Premium licences (one per service principal managed by a conditional access). -> API Limits This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to reduce parallelism https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options or specify extended custom resource timeouts https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts.
---

# Resource: azuread_conditional_access_policy

Manages a Conditional Access Policy within Azure Active Directory.  -> **Licensing Requirements** Specifying `client_applications` property requires the activation of Microsoft Entra on your tenant and the availability of sufficient Workload Identities Premium licences (one per service principal managed by a conditional access).  -> **API Limits** This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to [reduce parallelism](https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options) or specify extended [custom resource timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts).

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
* `display_name` - (Required) No description provided.
* `state` - (Required) No description provided. Possible values include `disabled`, `enabled`, and `enabledForReportingButNotEnforced`.

---

* `grant_controls` - (Optional) The `grant_controls` block as defined below. At least one of `grant_controls` or `session_controls` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-6772616e745f636f6e74726f6c73))
* `session_controls` - (Optional) The `session_controls` block as defined below. At least one of `grant_controls` or `session_controls` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-73657373696f6e5f636f6e74726f6c73))

---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73"></a>
The `conditions` block supports the following:

* `applications` - (Required) The `applications` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73))
* `client_app_types` - (Required) No description provided. For each element: Possible values include `all`, `browser`, `easSupported`, `exchangeActiveSync`, `mobileAppsAndDesktopClients`, and `other`.
* `users` - (Required) The `users` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273))

* `authentication_flow_transfer_methods` - (Optional) No description provided. For each element: Possible values include `authenticationTransfer` and `deviceCodeFlow`.
* `client_applications` - (Optional) The `client_applications` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73))
* `devices` - (Optional) The `devices` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573))
* `insider_risk_levels` - (Optional) No description provided. Possible values include `elevated`, `minor`, and `moderate`.
* `locations` - (Optional) The `locations` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6c6f636174696f6e73))
* `platforms` - (Optional) The `platforms` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-706c6174666f726d73))
* `service_principal_risk_levels` - (Optional) No description provided. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.
* `sign_in_risk_levels` - (Optional) No description provided. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.
* `user_risk_levels` - (Optional) No description provided. For each element: Possible values include `hidden`, `high`, `low`, `medium`, and `none`.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73"></a>
The `conditions.applications` block supports the following:

* `excluded_applications` - (Optional) No description provided.
* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73-66696c746572))
* `included_applications` - (Optional) No description provided. Exactly one of `conditions.0.applications.0.included_applications` or `conditions.0.applications.0.included_user_actions` must be specified.
* `included_user_actions` - (Optional) No description provided. Exactly one of `conditions.0.applications.0.included_applications` or `conditions.0.applications.0.included_user_actions` must be specified.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6170706c69636174696f6e73-66696c746572"></a>
The `conditions.applications.filter` block supports the following:

* `mode` - (Required) No description provided. Possible values include `exclude` and `include`.
* `rule` - (Required) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73"></a>
The `conditions.client_applications` block supports the following:

* `excluded_service_principals` - (Optional) No description provided.
* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73-66696c746572))
* `included_service_principals` - (Optional) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-636c69656e745f6170706c69636174696f6e73-66696c746572"></a>
The `conditions.client_applications.filter` block supports the following:

* `mode` - (Required) No description provided. Possible values include `exclude` and `include`.
* `rule` - (Required) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573"></a>
The `conditions.devices` block supports the following:

* `filter` - (Optional) The `filter` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573-66696c746572))
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-64657669636573-66696c746572"></a>
The `conditions.devices.filter` block supports the following:

* `mode` - (Required) No description provided. Possible values include `exclude` and `include`.
* `rule` - (Required) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-6c6f636174696f6e73"></a>
The `conditions.locations` block supports the following:

* `included_locations` - (Required) No description provided.

* `excluded_locations` - (Optional) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-706c6174666f726d73"></a>
The `conditions.platforms` block supports the following:

* `included_platforms` - (Required) No description provided. For each element: Possible values include `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, and `windowsPhone`.

* `excluded_platforms` - (Optional) No description provided. For each element: Possible values include `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, and `windowsPhone`.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273"></a>
The `conditions.users` block supports the following:

* `excluded_groups` - (Optional) No description provided.
* `excluded_guests_or_external_users` - (Optional) The `excluded_guests_or_external_users` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273))
* `excluded_roles` - (Optional) No description provided.
* `excluded_users` - (Optional) No description provided.
* `included_groups` - (Optional) No description provided. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
* `included_guests_or_external_users` - (Optional) The `included_guests_or_external_users` block as defined below. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273))
* `included_roles` - (Optional) No description provided. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
* `included_users` - (Optional) No description provided. At least one of `conditions.0.users.0.included_groups`, `conditions.0.users.0.included_roles`, `conditions.0.users.0.included_users`, or `conditions.0.users.0.included_guests_or_external_users` must be specified.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273"></a>
The `conditions.users.excluded_guests_or_external_users` block supports the following:

* `guest_or_external_user_types` - (Required) No description provided. For each element: Possible values include `b2bCollaborationGuest`, `b2bCollaborationMember`, `b2bDirectConnectUser`, `internalGuest`, `none`, `otherExternalUser`, and `serviceProvider`.

* `external_tenants` - (Optional) The `external_tenants` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473))
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-6578636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473"></a>
The `conditions.users.excluded_guests_or_external_users.external_tenants` block supports the following:

* `membership_kind` - (Required) No description provided. Possible values include `all` and `enumerated`.

* `members` - (Optional) No description provided.
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273"></a>
The `conditions.users.included_guests_or_external_users` block supports the following:

* `guest_or_external_user_types` - (Required) No description provided. For each element: Possible values include `b2bCollaborationGuest`, `b2bCollaborationMember`, `b2bDirectConnectUser`, `internalGuest`, `none`, `otherExternalUser`, and `serviceProvider`.

* `external_tenants` - (Optional) The `external_tenants` block as defined below. (see [nested schema](#schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473))
---

<a id="schema-6e6573746564-696e707574-636f6e646974696f6e73-7573657273-696e636c756465645f6775657374735f6f725f65787465726e616c5f7573657273-65787465726e616c5f74656e616e7473"></a>
The `conditions.users.included_guests_or_external_users.external_tenants` block supports the following:

* `membership_kind` - (Required) No description provided. Possible values include `all` and `enumerated`.

* `members` - (Optional) No description provided.
---

<a id="schema-6e6573746564-696e707574-6772616e745f636f6e74726f6c73"></a>
The `grant_controls` block supports the following:

* `operator` - (Required) No description provided. Possible values include `AND` and `OR`.

* `authentication_strength_policy_id` - (Optional) No description provided. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.
* `built_in_controls` - (Optional) No description provided. For each element: Possible values include `approvedApplication`, `block`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `mfa`, and `passwordChange`. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.
* `custom_authentication_factors` - (Optional) No description provided.
* `terms_of_use` - (Optional) No description provided. At least one of `grant_controls.0.built_in_controls`, `grant_controls.0.authentication_strength_policy_id`, or `grant_controls.0.terms_of_use` must be specified.
---

<a id="schema-6e6573746564-696e707574-73657373696f6e5f636f6e74726f6c73"></a>
The `session_controls` block supports the following:

* `application_enforced_restrictions_enabled` - (Optional) No description provided.
* `cloud_app_security_policy` - (Optional) No description provided. Possible values include `blockDownloads`, `mcasConfigured`, and `monitorOnly`.
* `disable_resilience_defaults` - (Optional) No description provided.
* `persistent_browser_mode` - (Optional) No description provided. Possible values include `always` and `never`.
* `sign_in_frequency` - (Optional) No description provided. Value must be at least `0`. Required when using `session_controls.0.sign_in_frequency_period`.
* `sign_in_frequency_authentication_type` - (Optional) No description provided. Possible values include `primaryAndSecondaryAuthentication` and `secondaryAuthentication`.
* `sign_in_frequency_interval` - (Optional) No description provided. Possible values include `everyTime` and `timeBased`.
* `sign_in_frequency_period` - (Optional) No description provided. Possible values include `days` and `hours`. Required when using `session_controls.0.sign_in_frequency`.

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
