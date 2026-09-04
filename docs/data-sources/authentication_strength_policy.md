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

*Look up a built-in policy*

```terraform
data "azuread_authentication_strength_policy" "example" {
  display_name = "Multifactor authentication"
}
```

The display names of the built-in policies supplied by Microsoft are:

* `Multifactor authentication`
* `Passwordless MFA`
* `Phishing resistant MFA`

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the authentication strength policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allowed_combinations` - A list of allowed authentication methods combinations for this authentication strength policy.
* `description` - The description of this authentication strength policy.
* `id` - The ID of this authentication strength policy.
