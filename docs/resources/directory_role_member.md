---
page_title: "Resource: azuread_directory_role_member"
subcategory: "Directory Roles"
description: |-
  Manages a single directory role membership (assignment) within Azure Active Directory.
  -> Deprecation Warning: This resource has been superseded by the azuread_directory_role_assignment https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role_assignment resource and will be removed in version 4.0 of the AzureAD provider.
---

# Resource: azuread_directory_role_member

Manages a single directory role membership (assignment) within Azure Active Directory.

-> **Deprecation Warning:** This resource has been superseded by the [azuread_directory_role_assignment](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/directory_role_assignment) resource and will be removed in version 4.0 of the AzureAD provider.

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

## Arguments Reference

The following arguments are supported:

* `member_object_id` - (Optional) The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this value forces a new resource to be created.
* `role_object_id` - (Optional) The object ID of the directory role you want to add the member to. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_directory_role_member.example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
