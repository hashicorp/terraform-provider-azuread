---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role_member

Manages a single Directory Role Membership within Azure Active Directory.

-> **NOTE:** This resource is supported only if the  MS Graph API is enabled.

## Example Usage

```terraform

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Global Administrator"
}

resource "azuread_directory_role_member" "example" {
  directory_role_object_id = azuread_directory_role.example.id
  member_object_id = data.azuread_user.example.id
}
```

## Argument Reference

The following arguments are supported:

* `directory_role_object_id` - (Required) The Object ID of the Azure AD Directory Role you want to add the Member to. Changing this forces a new resource to be created.
* `member_object_id` - (Required) The Object ID of the Azure AD Object you want to add as a Member to the Directory Role. Supported Object types are Users or Groups. Changing this forces a new resource to be created.

-> **NOTE:** The Member object has to be present in your Azure Active Directory, either as a Member or a Guest.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Azure Active Directory Role Members can be imported using the `object id`, e.g.

```shell
terraform import azuread_directory_role_member.test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Azure AD Directory Role Object ID and the target Member Object ID in the format `{DirectoryRoleObjectID}/member/{MemberObjectID}`.
