---
page_title: "Resource: azuread_authentication_strength_policy_x509_combination_configuration"
subcategory: "Policies"
description: |-
  Manages an x509 certificate combination configuration for an authentication strength policy within Azure Active Directory.
---

# Resource: azuread_authentication_strength_policy_x509_combination_configuration

Manages an x509 certificate combination configuration for an authentication strength policy within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_authentication_strength_policy" "example" {
  display_name = "Example Authentication Strength Policy"
  description  = "Policy for demo purposes"
  allowed_combinations = [
    "x509CertificateSingleFactor",
    "x509CertificateMultiFactor",
  ]
}

resource "azuread_authentication_strength_policy_x509_combination_configuration" "single_factor" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id
  applies_to_combinations           = ["x509CertificateSingleFactor"]
  allowed_issuer_skis               = ["9A4248C6AC8C2931AB2A86537818E92E7B6C97B6"]
}

resource "azuread_authentication_strength_policy_x509_combination_configuration" "multi_factor" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.example.object_id
  applies_to_combinations           = ["x509CertificateMultiFactor"]
  allowed_policy_oids               = ["1.2.3.4.5"]
}
```

## Arguments Reference

The following arguments are supported:

* `applies_to_combinations` - (Required) The x509 certificate authentication method combinations this configuration applies to. For each element: Possible values include `x509CertificateMultiFactor` and `x509CertificateSingleFactor`.
* `authentication_strength_policy_id` - (Required) The object ID of the authentication strength policy to which this combination configuration applies. Changing this value forces a new resource to be created.

---

* `allowed_issuer_skis` - (Optional) A set of allowed certificate issuer subject key identifier values. At least one of `allowed_issuer_skis` or `allowed_policy_oids` must be specified.
* `allowed_policy_oids` - (Optional) A set of allowed certificate policy OIDs. At least one of `allowed_issuer_skis` or `allowed_policy_oids` must be specified.

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
terraform import azuread_authentication_strength_policy_x509_combination_configuration.example /policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000000/combinationConfigurations/11111111-1111-1111-1111-111111111111
```
