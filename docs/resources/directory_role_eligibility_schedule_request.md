---
page_title: "Resource: azuread_directory_role_eligibility_schedule_request"
subcategory: "Directory Roles"
description: |-
  Manages a single directory role eligibility schedule request within Azure Active Directory.
---

# Resource: azuread_directory_role_eligibility_schedule_request

Manages a single directory role eligibility schedule request within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

The calling principal requires one of the following application roles: `RoleEligibilitySchedule.ReadWrite.Directory` or `RoleManagement.ReadWrite.Directory`.

The calling principal requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`.

## Example Usage

```terraform
data "azuread_user" "example" {
  user_principal_name = "jdoe@hashicorp.com"
}

resource "azuread_directory_role" "example" {
  display_name = "Application Administrator"
}

resource "azuread_directory_role_eligibility_schedule_request" "example" {
  role_definition_id = azuread_directory_role.example.template_id
  principal_id       = data.azuread_user.example.object_id
  directory_scope_id = "/"
  justification      = "Example"
}
```

## Arguments Reference

The following arguments are supported:

* `directory_scope_id` - (Required) Identifier of the directory object representing the scope of the role eligibility schedule request. Changing this value forces a new resource to be created.
* `justification` - (Required) Justification for why the role is assigned. Changing this value forces a new resource to be created.
* `principal_id` - (Required) The object ID of the member principal. Changing this value forces a new resource to be created.
* `role_definition_id` - (Required) The object ID of the directory role for this role eligibility schedule request. Changing this value forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_directory_role_eligibility_schedule_request.example 822ec710-4c9f-4f71-a27a-451759cc7522
```
