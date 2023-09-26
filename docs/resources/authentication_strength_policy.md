---
subcategory: "Conditional Access"
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
```

## Argument Reference

The following arguments are supported:

- `allowed_combinations` - (Required) List of allowed authentication methods for this authentication strength policy.
- `description` - (Optional) The description for this authentication strength policy.
- `display_name` - (Required) The friendly name for this authentication strength policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The ID of the authentication strength policy.

## Import

Authentication Strength Policies can be imported using the `id`, e.g.

```shell
terraform import azuread_authentication_strength_policy.my_policy 00000000-0000-0000-0000-000000000000
```
