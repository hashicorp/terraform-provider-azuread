---
subcategory: "Policies"
---

# Data Source: azuread_authentication_strength_policy

Use this data source to retrieve information about an authentication strength policy within Azure Active Directory. This can be used to read either the built-in policies supplied by Microsoft, or custom policies created in the tenant.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Policy.Read.AuthenticationMethod` or `Policy.Read.All`

When authenticated with a user principal, this data source requires one of the following directory roles: `Conditional Access Administrator`, `Security Administrator`, `Security Reader` or `Global Administrator`

## Example Usage

*Look up a built-in policy by display name*

```terraform
data "azuread_authentication_strength_policy" "example" {
  display_name = "Multifactor authentication"
}
```

The display names of the built-in policies supplied by Microsoft are:

* `Multifactor authentication`
* `Passwordless MFA`
* `Phishing-resistant MFA`

*Look up a policy by object ID*

```terraform
data "azuread_authentication_strength_policy" "example" {
  object_id = "00000000-0000-0000-0000-000000000004"
}
```

*Look up a policy managed elsewhere in the same configuration*

```terraform
data "azuread_authentication_strength_policy" "example" {
  object_id = azuread_authentication_strength_policy.example.object_id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name of the authentication strength policy.
* `object_id` - (Optional) The object ID of the authentication strength policy.

~> One of `display_name` or `object_id` must be specified.

~> **Tip** Display names are expected to be unique within a tenant, however this is not guaranteed by the API. Specify `object_id` where you need to be certain of matching a specific policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allowed_combinations` - A list of allowed authentication methods combinations for this authentication strength policy.
* `combination_configurations` - A list of `combination_configurations` blocks as documented below, which further constrain the allowed combinations for this policy.
* `description` - The description of this authentication strength policy.
* `display_name` - The display name of this authentication strength policy.
* `id` - The ID of this authentication strength policy.
* `object_id` - The object ID of this authentication strength policy.

---

`combination_configurations` block exports the following:

* `allowed_aaguids` - A list of AAGUIDs allowed by this combination configuration. Only populated when `type` is `fido2CombinationConfiguration`.
* `allowed_issuer_skis` - A list of allowed certificate issuer subject key identifier values. Only populated when `type` is `x509CertificateCombinationConfiguration`.
* `allowed_policy_oids` - A list of allowed certificate policy OIDs. Only populated when `type` is `x509CertificateCombinationConfiguration`.
* `applies_to_combinations` - A list of authentication method combinations this configuration applies to, for example `fido2` or `x509CertificateSingleFactor`. This is distinct from `type`, which identifies the kind of combination configuration.
* `object_id` - The object ID of this combination configuration.
* `type` - The type of this combination configuration, either `fido2CombinationConfiguration` or `x509CertificateCombinationConfiguration`.

~> Combination configurations are managed using the `azuread_authentication_strength_policy_fido2_combination_configuration` and `azuread_authentication_strength_policy_x509_combination_configuration` resources.
