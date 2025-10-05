---
subcategory: "Applications"
---

# Resource: azuread_application_permission_scope

Manages a permission scope for an application registration.

This resource is analogous to the `oauth2_permission_scope` block in the `api` block of the  `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "random_uuid" "example_administer" {}

resource "azuread_application_permission_scope" "example" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.example_administer.id
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
}
```

-> **Tip** For managing more permissions scopes, create additional instances of this resource

*Usage with azuread_application resource*

```terraform

resource "azuread_application" "example" {
  display_name = "example"

  lifecycle {
    ignore_changes = [
      api[0].oauth2_permission_scope,
    ]
  }
}

resource "azuread_application_permission_scope" "example" {
  application_id = azuread_application.example.id
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `admin_consent_description` - (Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - (Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `scope_id` - (Required) The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
* `type` - (Required) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
* `user_consent_description` - (Required) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - (Required) Display name for the delegated permission that appears in the end user consent experience.

-> **Tip** Use the `random_uuid` resource to generate UUIDs and save them to state for permission scopes within your Terraform configuration

* `value` - (Optional) The value that is used for the `scp` claim in OAuth access tokens.

-> **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.

## Attribute Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.

```shell
terraform import azuread_application_permission_scope.example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
```
