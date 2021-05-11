---
subcategory: "Applications"
---

# Resource: azuread_application_app_role

Manages an App Role associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_app_role" "example" {
  application_object_id = azuread_application.example.id
  allowed_member_types  = ["User"]
  description           = "Admins can manage roles and perform all task actions"
  display_name          = "Admin"
  enabled               = true
  value                 = "administer"
}
```

## Argument Reference

The following arguments are supported:

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
* `application_object_id` - (Required) The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.
* `description` - (Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - (Required) Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - (Optional) Determines if the app role is enabled: Defaults to `true`.
* `role_id` - (Optional) The unique identifier for the app role. If omitted, a random UUID will be automatically generated. Must be a valid UUID. Changing this field forces a new resource to be created.
* `value` - (Optional) The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

App Roles can be imported using the `object_id` of an Application and the `id` of the App Role, e.g.

```shell
terraform import azuread_application_app_role.test 00000000-0000-0000-0000-000000000000/role/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID, the string "role" and the App Role's ID in the format `{ApplicationObjectId}/role/{AppRoleId}`.
