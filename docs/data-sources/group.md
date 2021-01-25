---
subcategory: "Groups"
---

# Data Source: azuread_group

Gets information about an Azure Active Directory group.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage (by Group Display Name)

```hcl
data "azuread_group" "example" {
  display_name = "A-AD-Group"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The splay name of the Group within Azure Active Directory.
* `object_id` - (Optional) Specifies the Object ID of the Group within Azure Active Directory.

~> **NOTE:** One of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `description` - The description of the AD Group.
* `display_name` - The name of the Azure AD Group.
* `id` - The Object ID of the Azure AD Group.
* `members` - The Object IDs of the Azure AD Group members.
* `owners` - The Object IDs of the Azure AD Group owners.

