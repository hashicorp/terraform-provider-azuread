---
page_title: "Data Source: azuread_group_role_management_policy"
subcategory: "Policies"
description: |-
  Use this data source to retrieve a role policy for an Azure AD group.
---

# Data Source: azuread_group_role_management_policy

Use this data source to retrieve a role policy for an Azure AD group.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the `RoleManagementPolicy.Read.AzureADGroup` Microsoft Graph API permissions.

When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.

## Example Usage

```terraform
resource "azuread_group" "example" {
  display_name     = "group-name"
  security_enabled = true
}

data "azuread_group_role_management_policy" "owners_policy" {
  group_id = azuread_group.example.id
  role_id  = "owner"
}
```

## Arguments Reference

The following arguments are supported:

* `group_id` - (Required) ID of the group to which this policy is assigned.
* `role_id` - (Required) The ID of the role of this policy to the group. Possible values include `member` and `owner`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `description` - Description of the policy.
* `display_name` - The display name of the policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 0s) Used when retrieving the data-source.
