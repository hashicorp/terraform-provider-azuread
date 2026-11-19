---
page_title: "Data Source: azuread_authentication_strength_policy"
subcategory: "Policies"
description: |-
  Use this data source to retrieve information about an authentication strength policy within Azure Active Directory. This can be used to read either the built-in policies supplied by Microsoft, or custom policies created in the tenant.
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

## Arguments Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the authentication strength policy. Exactly one of `display_name` or `object_id` must be specified.
* `object_id` - (Optional) The object ID of the authentication strength policy. Exactly one of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `allowed_combinations` - The allowed MFA methods for this policy.
* `combination_configurations` - The combination configurations which further constrain the allowed combinations for this policy. (see [nested schema](#schema-6e6573746564-6f7574707574-636f6d62696e6174696f6e5f636f6e66696775726174696f6e73))
* `description` - The description for the authentication strength policy.

---

<a id="schema-6e6573746564-6f7574707574-636f6d62696e6174696f6e5f636f6e66696775726174696f6e73"></a>
The `combination_configurations` object exports the following:

* `allowed_aaguids` - The AAGUIDs allowed by this configuration, for the `fido2CombinationConfiguration` type.
* `allowed_issuer_skis` - The certificate issuer subject key identifiers allowed by this configuration, for the `x509CertificateCombinationConfiguration` type.
* `allowed_policy_oids` - The certificate policy OIDs allowed by this configuration, for the `x509CertificateCombinationConfiguration` type.
* `applies_to_combinations` - The authentication method combinations this configuration applies to.
* `object_id` - The object ID of the combination configuration.
* `type` - The type of the combination configuration.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 0s) Used when retrieving the data-source.
