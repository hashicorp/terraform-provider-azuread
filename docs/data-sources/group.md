---
subcategory: "Groups"
---

# Data Source: azuread_group

Gets information about an Azure Active Directory group.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

## Example Usage (by Group Display Name)

```hcl
data "azuread_group" "example" {
  display_name     = "MyGroupName"
  security_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) The display name for the Group.
* `mail_enabled` - (Optional) Whether the group is mail-enabled.
* `object_id` - (Optional) Specifies the Object ID of the Group.
* `security_enabled` - (Optional) Whether the group is a security group.

~> **NOTE:** One of `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `description` - The optional description of the Group.
* `display_name` - The display name for the Group.
* `id` - The Object ID of the Azure AD Group.
* `mail_enabled` - Whether the group is mail-enabled.
* `members` - The Object IDs of the Group members.
* `owners` - The Object IDs of the Group owners.
* `security_enabled` - Whether the group is a security group.
