---
subcategory: "Applications"
---

# Resource: azuread_application_oauth2_permission

Manages an OAuth2 Permission (also known as a Scope) associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_oauth2_permission" "example" {
  application_object_id      = azuread_application.example.id
  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  is_enabled                 = true
  type                       = "User"
  user_consent_description   = "Administer the application"
  user_consent_display_name  = "Administer"
  value                      = "administer"
}
```

## Argument Reference

The following arguments are supported:

* `admin_consent_description` - (Required) Permission help text that appears in the admin consent and app assignment experiences.
* `admin_consent_display_name` - (Required) Display name for the permission that appears in the admin consent and app assignment experiences.
* `application_object_id` - (Required) The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
* `is_enabled` - (Optional) Determines if the Permission is enabled. Defaults to `true`.
* `permission_id` - (Optional) Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
* `type` - (Required) Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
* `user_consent_description` - (Optional) Permission help text that appears in the end user consent experience.
* `user_consent_display_name` - (Optional) Display name for the permission that appears in the end user consent experience.
* `value` - (Required) The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

OAuth2 Permissions can be imported using the `object id` of an Application and the `id` of the Permission, e.g.

```shell
terraform import azuread_application_oauth2_permission.test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application's Object ID, the string "scope" and the Permission's ID in the format `{ApplicationObjectId}/scope/{PermissionId}`.
