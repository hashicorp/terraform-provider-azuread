---
subcategory: "Policies"
---

# Resource: azuread_authentication_strength_policy_fido2_combination_configuration

Manages a FIDO2 combination configuration for an authentication strength policy within Azure Active Directory.

A FIDO2 combination configuration restricts which security keys, identified by their Authenticator Attestation GUID (AAGUID), satisfy the `fido2` authentication method combination. A policy supports at most one FIDO2 combination configuration.

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

## Argument Reference

The following arguments are supported:

* `allowed_aaguids` - (Required) A set of Authenticator Attestation GUIDs (AAGUIDs) allowed to satisfy the `fido2` combination.
* `authentication_strength_policy_id` - (Required) The object ID of the authentication strength policy to which this combination configuration applies. Changing this field forces a new resource to be created.

-> The associated authentication strength policy must specify `fido2` in its `allowed_combinations`, otherwise the API rejects the combination configuration.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the combination configuration.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

FIDO2 combination configurations can be imported using the `id`, e.g.

```shell
terraform import azuread_authentication_strength_policy_fido2_combination_configuration.example /policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000000/combinationConfigurations/11111111-1111-1111-1111-111111111111
```
