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

resource "azuread_authentication_strength_policy" "example3" {
  display_name         = "Example Authentication Strength Policy with FIDO2 restrictions"
  description          = "Only allow specific security keys"
  allowed_combinations = ["fido2"]

  fido2_combination_configuration {
    allowed_aaguids = [
      "de1e552d-db1d-4423-a619-566b625cdc84",
      "90a3ccdf-635c-4729-a248-9b709135078f",
    ]
  }
}

resource "azuread_authentication_strength_policy" "example4" {
  display_name = "Example Authentication Strength Policy with x509 restrictions"
  description  = "Only allow certificates from specific issuers"
  allowed_combinations = [
    "x509CertificateSingleFactor",
    "x509CertificateMultiFactor",
  ]

  x509_certificate_combination_configuration {
    applies_to_combinations = [
      "x509CertificateSingleFactor",
      "x509CertificateMultiFactor",
    ]
    allowed_issuer_skis = ["9af52a26d8e4bd7d5e8f43e9c7c5e2f4a3b1c0d9"]
    allowed_policy_oids = ["1.2.3.4.5"]
  }
}
```

## Argument Reference

The following arguments are supported:

- `allowed_combinations` - (Required) List of allowed authentication methods for this authentication strength policy.
- `description` - (Optional) The description for this authentication strength policy.
- `display_name` - (Required) The friendly name for this authentication strength policy.
- `fido2_combination_configuration` - (Optional) A `fido2_combination_configuration` block as documented below. Requires `fido2` to be present in `allowed_combinations`.
- `x509_certificate_combination_configuration` - (Optional) A `x509_certificate_combination_configuration` block as documented below. Requires a matching `x509Certificate*` value in `allowed_combinations`.

---

A `fido2_combination_configuration` block supports the following:

- `allowed_aaguids` - (Required) A list of Authenticator Attestation GUIDs (AAGUIDs) allowed to satisfy the `fido2` combination.

---

A `x509_certificate_combination_configuration` block supports the following:

- `allowed_issuer_skis` - (Optional) A list of allowed certificate issuer subject key identifier (SKI) values. At least one of `allowed_issuer_skis` or `allowed_policy_oids` must be specified.
- `allowed_policy_oids` - (Optional) A list of allowed certificate policy OIDs. At least one of `allowed_issuer_skis` or `allowed_policy_oids` must be specified.
- `applies_to_combinations` - (Required) A list of x509 certificate authentication method combinations this configuration applies to. Possible values are `x509CertificateSingleFactor` and `x509CertificateMultiFactor`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `fido2_combination_configuration.id` - The system-generated ID of the FIDO2 combination configuration.
- `id` - The ID of the authentication strength policy.
- `x509_certificate_combination_configuration.id` - The system-generated ID of the x509 certificate combination configuration.

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
