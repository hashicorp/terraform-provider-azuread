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
* `Phishing resistant MFA`

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
* `description` - The description of this authentication strength policy.
* `display_name` - The display name of this authentication strength policy.
* `id` - The ID of this authentication strength policy.
* `object_id` - The object ID of this authentication strength policy.
