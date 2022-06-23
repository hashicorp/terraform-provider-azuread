---
subcategory: "Directory Roles"
---

# Resource: azuread_directory_role_member

Manages a single directory role membership (assignment) within Azure Active Directory.

-> **Deprecation Warning:** This resource has been superseded by the [azuread_directory_role_assignment](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role_assignment) resource and will be removed in version 3.0 of the AzureAD provider

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

## Example Usage

```terraform

data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Security administrator"
}

resource "azuread_directory_role_member" "example" {
  role_object_id   = azuread_directory_role.example.object_id
  member_object_id = data.azuread_user.example.object_id
}
```

## Argument Reference

The following arguments are supported:

* `member_object_id` - (Required) The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
* `role_object_id` - (Required) The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.

```shell
terraform import azuread_directory_role_member.test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
