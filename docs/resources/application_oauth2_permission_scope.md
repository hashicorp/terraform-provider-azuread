---
subcategory: "Applications"
---

# Resource: azuread_application_oauth2_permission_scope

Manages an OAuth 2.0 Permission Scope associated with an application.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
resource "azuread_application" "example" {
  name = "example"
}

resource "azuread_application_oauth2_permission_scope" "example" {
  application_object_id      = azuread_application.example.id
  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  enabled                    = true
  type                       = "User"
  user_consent_description   = "Administer the application"
  user_consent_display_name  = "Administer"
  value                      = "administer"
}
```

## Argument Reference

The following arguments are supported:

* `admin_consent_description` - (Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - (Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `application_object_id` - (Required) The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
* `enabled` - (Optional) Determines if the permission scope is enabled. Defaults to `true`.
* `scope_id` - (Optional) Specifies a custom UUID for the permission scope. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
* `type` - (Required) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
* `user_consent_description` - (Optional) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - (Optional) Display name for the delegated permission that appears in the end user consent experience.
* `value` - (Required) The value that is used for the `scp` claim in OAuth 2.0 access tokens.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

*No additional attributes are exported*

## Import

OAuth2 Permission Scopes can be imported using the `object_id` of an Application and the `id` of the Permission Scope, e.g.

```shell
terraform import azuread_application_oauth2_permission_scope.test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
```

-> **NOTE:** This ID format is unique to Terraform and is composed of the Application Object ID, the string "scope", and the Permission Scope ID in the format `{ApplicationObjectId}/scope/{PermissionScopeId}`.
