---
subcategory: "Groups"
layout: "azuread"
page_title: "Azure Active Directory: azuread_group"
description: |-
  Manages a Group within Azure Active Directory.

---

# azuread_group

Manages a Group within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.

## Example Usage

*Basic example*

```hcl
resource "azuread_group" "example" {
  name = "A-AD-Group"
}
```

*A group with members*

```hcl
resource "azuread_user" "example" {
  display_name          = "J Doe"
  password              = "notSecure123"
  user_principal_name   = "jdoe@hashicorp.com"
}

resource "azuread_group" "example" {
  name    = "MyGroup"
  members = [
    azuread_user.example.object_id,
    /* more users */
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The display name for the Group. Changing this forces a new resource to be created.
* `description` - (Optional) The description for the Group.  Changing this forces a new resource to be created.
* `members` (Optional) A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
* `owners` (Optional) A set of owners who own this Group. Supported Object types are Users or Service Principals.
* `prevent_duplicate_names` - (Optional) If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.

-> **NOTE:** Group names are not unique within Azure Active Directory.

~> **NOTE:** Do not use the `azuread_group_member` resource at the same time as the `members` argument.

## Attributes Reference

The following attributes are exported:

* `id` - The Object ID of the Group.

## Import

Azure Active Directory Groups can be imported using the `object id`, e.g.

```shell
terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000
```
