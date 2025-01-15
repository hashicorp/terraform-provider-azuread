---
subcategory: "Policies"
---

# Resource: azuread_authentication_strength_policy

Manages a Authentication Strength Policy within Azure Active Directory.

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
    "fido2",
    "password",
  ]
}

resource "azuread_authentication_strength_policy" "example2" {
  display_name = "Example Authentication Strength Policy"
  description  = "Policy for demo purposes with all possible combinations"
  allowed_combinations = [
    "fido2",
    "password",
    "deviceBasedPush",
    "temporaryAccessPassOneTime",
    "federatedMultiFactor",
    "federatedSingleFactor",
    "hardwareOath,federatedSingleFactor",
    "microsoftAuthenticatorPush,federatedSingleFactor",
    "password,hardwareOath",
    "password,microsoftAuthenticatorPush",
    "password,sms",
    "password,softwareOath",
    "password,voice",
    "sms",
    "sms,federatedSingleFactor",
    "softwareOath,federatedSingleFactor",
    "temporaryAccessPassMultiUse",
    "voice,federatedSingleFactor",
    "windowsHelloForBusiness",
    "x509CertificateMultiFactor",
    "x509CertificateSingleFactor",
  ]
}
```

## Argument Reference

The following arguments are supported:

- `allowed_combinations` - (Required) List of allowed authentication methods for this authentication strength policy.
- `description` - (Optional) The description for this authentication strength policy.
- `display_name` - (Required) The friendly name for this authentication strength policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The ID of the authentication strength policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Authentication Strength Policies can be imported using the `id`, e.g.

```shell
terraform import azuread_authentication_strength_policy.my_policy /policies/authenticationStrengthPolicies/00000000-0000-0000-0000-000000000000
```
