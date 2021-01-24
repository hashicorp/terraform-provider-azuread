---
subcategory: "Applications"
---

# Resource: azuread_application_app_role

Manages an App Role associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_app_role" "example" {
  application_object_id = azuread_application.example.id
  allowed_member_types  = ["User"]
  description           = "Admins can manage roles and perform all task actions"
  display_name          = "Admin"
  is_enabled            = true
  value                 = "administer"
}
```

## Argument Reference

The following arguments are supported:

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.
* `application_object_id` - (Required) The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.
* `description` - (Required) Permission help text that appears in the admin app assignment and consent experiences.
* `display_name` - (Required) Display name for the permission that appears in the admin consent and app assignment experiences.
* `is_enabled` - (Optional) Determines if the app role is enabled. Defaults to `true`.
* `role_id` - (Optional) Specifies a custom UUID for the app role. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
* `value` - (Optional) Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

App Roles can be imported using the `object id` of an Application and the `id` of the App Role, e.g.

```shell
terraform import azuread_application_app_role.test 00000000-0000-0000-0000-000000000000/role/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID, the string "role" and the App Role's ID in the format `{ApplicationObjectId}/role/{AppRoleId}`.
