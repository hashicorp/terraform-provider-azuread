---
subcategory: "Delegated Permission Grants"
---

# Resource: azuread_service_principal_delegated_permission_grant

Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`

When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`

## Example Usage

*Delegated permission grant for all users*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}

resource "azuread_application" "example" {
  display_name = "example"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["openid"]
      type = "Scope"
    }

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.Read"]
      type = "Scope"
    }
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_service_principal_delegated_permission_grant" "example" {
  service_principal_object_id          = azuread_service_principal.example.object_id
  resource_service_principal_object_id = azuread_service_principal.msgraph.object_id
  claim_values                         = ["openid", "User.Read.All"]
}
```

*Delegated permission grant for a single user*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}

resource "azuread_application" "example" {
  display_name = "example"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["openid"]
      type = "Scope"
    }

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.Read"]
      type = "Scope"
    }
  }
}

resource "azuread_service_principal" "example" {
  client_id = azuread_application.example.client_id
}

resource "azuread_user" "example" {
  display_name        = "J. Doe"
  user_principal_name = "jdoe@hashicorp.com"
  mail_nickname       = "jdoe"
  password            = "SecretP@sswd99!"
}

resource "azuread_service_principal_delegated_permission_grant" "example" {
  service_principal_object_id          = azuread_service_principal.example.object_id
  resource_service_principal_object_id = azuread_service_principal.msgraph.object_id
  claim_values                         = ["openid", "User.Read.All"]
  user_object_id                       = azuread_user.example.object_id
}
```

## Argument Reference

The following arguments are supported:

* `claim_values` - (Required) - A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
* `resource_service_principal_object_id` - (Required) The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
* `service_principal_object_id` - (Required) The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
* `user_object_id` - (Optional) - The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.

-> **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the delegated permission grant.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Delegated permission grants can be imported using their ID, e.g.

```shell
terraform import azuread_service_principal_delegated_permission_grant.example /oauth2PermissionGrants/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
```
