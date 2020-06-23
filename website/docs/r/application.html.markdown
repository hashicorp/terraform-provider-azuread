---
subcategory: "Application"
layout: "azuread"
page_title: "Azure Active Directory: azuread_application"
description: |-
  Manages an Application within Azure Active Directory.

---

# azuread_application

Manages an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
resource "azuread_application" "example" {
  name                       = "example"
  homepage                   = "https://homepage"
  identifier_uris            = ["https://uri"]
  reply_urls                 = ["https://replyurl"]
  available_to_other_tenants = false
  oauth2_allow_implicit_flow = true
  type                       = "webapp/api"
  owners                     = ["00000004-0000-0000-c000-000000000000"]

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"

    resource_access {
      id   = "..."
      type = "Role"
    }

    resource_access {
      id   = "..."
      type = "Scope"
    }

    resource_access {
      id   = "..."
      type = "Scope"
    }
  }

  required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"

    resource_access {
      id   = "..."
      type = "Scope"
    }
  }

  app_role {
    allowed_member_types = [
      "User",
      "Application",
    ]

    description  = "Admins can manage roles and perform all task actions"
    display_name = "Admin"
    is_enabled   = true
    value        = "Admin"
  }

  oauth2_permissions {
    admin_consent_description  = "Allow the application to access example on behalf of the signed-in user."
    admin_consent_display_name = "Access example"
    is_enabled                 = true
    type                       = "User"
    user_consent_description   = "Allow the application to access example on your behalf."
    user_consent_display_name  = "Access example"
    value                      = "user_impersonation"
  }

  oauth2_permissions {
    admin_consent_description  = "Administer the example application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }

  optional_claims {
    access_token {
      name = "myclaim"
    }

    access_token {
      name = "otherclaim"
    }

    id_token {
      name                  = "userclaim"
      source                = "user"
      essential             = true
      additional_properties = ["emit_as_roles"]
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

* `logout_url` - (Optional) The URL of the logout page.

* `available_to_other_tenants` - (Optional) Is this Azure AD Application available to other tenants? Defaults to `false`.

* `public_client` - (Optional) Is this Azure AD Application a public client? Defaults to `false`.

* `oauth2_allow_implicit_flow` - (Optional) Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.

* `group_membership_claims` - (Optional) Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.

* `optional_claims` - A collection of `access_token` or `id_token` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims

* `owners` - (Optional) A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list. 

* `required_resource_access` - (Optional) A collection of `required_resource_access` blocks as documented below.

* `type` - (Optional) Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.

* `app_role` - (Optional) A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles

* `oauth2_permissions` - (Optional) A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2_permissions` blocks as documented below.

* `prevent_duplicate_names` - (Optional) If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.

---

`required_resource_access` supports the following:

* `resource_app_id` - (Required) The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.

* `resource_access` - (Required) A collection of `resource_access` blocks as documented below.

---

`resource_access` supports the following:

* `id` - (Required) The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.

* `type` - (Required) Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

---

`access_token` and/or `id_token` blocks support the following:

* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
* `additional_properties` - List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.

---

`app_role` supports the following:

* `id` - The unique identifier of the `app_role`.

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.

* `description` - (Required) Permission help text that appears in the admin app assignment and consent experiences.

* `display_name` - (Required) Display name for the permission that appears in the admin consent and app assignment experiences.

* `is_enabled` - (Optional) Determines if the app role is enabled: Defaults to `true`.

* `value` - (Optional) Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

---

`oauth2_permissions` supports the following:

* `admin_consent_description` - (Required) Permission help text that appears in the admin consent and app assignment experiences.

* `admin_consent_display_name` - (Required) Display name for the permission that appears in the admin consent and app assignment experiences.

* `value` - (Required) The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.

* `type` - (Required) Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by a Company Administrator. Possible values are "User" or "Admin".

* `is_enabled` - (Optional) Determines if the permission is enabled: defaults to `true`.

* `user_consent_description` - (Optional) Permission help text that appears in the end user consent experience.

* `user_consent_display_name` - (Optional) Display name for the permission that appears in the end user consent experience.

If you don't specify any `oauth2_permissions` blocks, your Application will be assigned the default `user_impersonation` scope by Azure Active Directory. However, due to the declarative nature of Terraform configuration, if you do specify any `oauth2_permissions` blocks, you will need to include a block for the `user_impersonation` scope if you need it, or it will be removed (see the example above). To ensure that no OAuth 2.0 permission scopes are present for your Application, specify `oauth2_permissions = []` in your Application resource.

## Attributes Reference

The following attributes are exported:

* `application_id` - The Application ID.

* `object_id` - The Application's Object ID.

## Import

Azure Active Directory Applications can be imported using the `object id`, e.g.

```shell
terraform import azuread_application.test 00000000-0000-0000-0000-000000000000
```
