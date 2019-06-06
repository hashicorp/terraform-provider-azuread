---
layout: "azuread"
page_title: "Azure Active Directory: azuread_application"
sidebar_current: "docs-azuread-resource-azuread-application"
description: |-
  Manages an Application within Azure Active Directory.

---

# azuread_application

Manages an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "test" {
  name                       = "example"
  homepage                   = "https://homepage"
  identifier_uris            = ["https://uri"]
  reply_urls                 = ["https://replyurl"]
  available_to_other_tenants = false
  oauth2_allow_implicit_flow = true
  type                       = "webapp/api"

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"
  
    resource_access {
      id = "..."
      type = "Role"
    }
    resource_access {
      id = "..."
      type = "Scope"
    }
    
    resource_access {
      id = "..."
      type = "Scope"
    }
  }
    
  required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"
    
    resource_access {
      id = "..."
      type = "Scope"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The display name for the application.

* `homepage` - (optional) The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.

* `identifier_uris` - (Optional) A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.

* `reply_urls` - (Optional) A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.

* `available_to_other_tenants` - (Optional) Is this Azure AD Application available to other tenants? Defaults to `false`.

* `oauth2_allow_implicit_flow` - (Optional) Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.

* `group_membership_claims` - (Optional) Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup` or `All`.

* `required_resource_access` - (Optional) A collection of `required_resource_access` blocks as documented below.

* `type` - (Optional) Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.  

---

`required_resource_access` supports the following:

* `resource_app_id` - (Required) The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.

* `resource_access` - (Required) A collection of `resource_access` blocks as documented below.

---

`resource_access` supports the following:

* `id` - (Required) The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes. 

* `type` - (Required) Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

## Attributes Reference

The following attributes are exported:

* `application_id` - The Application ID.

* `object_id` - The Application's Object ID.

* `oauth2_permissions` - A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.

---

`oauth2_permission` block exports the following:

* `id` - The unique identifier for one of the `OAuth2Permission`.

* `type` - The type of the permission.

* `admin_consent_description` - The description of the admin consent.

* `admin_consent_display_name` - The display name of the admin consent.

* `is_enabled` - Is this permission enabled?

* `user_consent_description` - The description of the user consent.

* `user_consent_display_name` - The display name of the user consent.

* `value` - The name of this permission.

## Import

Azure Active Directory Applications can be imported using the `object id`, e.g.

```shell
terraform import azuread_application.test 00000000-0000-0000-0000-000000000000
```
