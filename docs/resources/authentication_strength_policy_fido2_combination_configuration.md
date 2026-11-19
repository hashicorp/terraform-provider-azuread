---
page_title: "Resource: azuread_authentication_strength_policy_fido2_combination_configuration"
subcategory: "Policies"
description: |-
  Manages a FIDO2 combination configuration for an authentication strength policy within Azure Active Directory.
---

# Resource: azuread_authentication_strength_policy_fido2_combination_configuration

Manages a FIDO2 combination configuration for an authentication strength policy within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_authentication_strength_policy" "example" {
  display_name         = "Example Authentication Strength Policy"
  description          = "Policy for demo purposes"
  allowed_combinations = ["fido2"]
}

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "example" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id

  allowed_aaguids = [
    "de1e552d-db1d-4423-a619-566b625cdc84",
    "90a3ccdf-635c-4729-a248-9b709135078f",
  ]
}
```

## Arguments Reference

The following arguments are supported:

* `allowed_aaguids` - (Required) A set of AAGUIDs allowed to be used as part of the `fido2` combination.
* `authentication_strength_policy_id` - (Required) The object ID of the authentication strength policy to which this combination configuration applies. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_authentication_strength_policy_fido2_combination_configuration.example /policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000000/combinationConfigurations/11111111-1111-1111-1111-111111111111
```
