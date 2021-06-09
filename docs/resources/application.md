---
subcategory: "Applications"
---

# Resource: azuread_application

Manages an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

## Example Usage

```terraform
data "azuread_client_config" "current" {}

resource "azuread_application" "example" {
  display_name     = "example"
  identifier_uris  = ["api://example-app"]
  owners           = [data.azuread_client_config.current.object_id]
  sign_in_audience = "AzureADMultipleOrgs"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access example on behalf of the signed-in user."
      admin_consent_display_name = "Access example"
      enabled                    = true
      id                         = "96183846-204b-4b43-82e1-5d2222eb4b9b"
      type                       = "User"
      user_consent_description   = "Allow the application to access example on your behalf."
      user_consent_display_name  = "Access example"
      value                      = "user_impersonation"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Administer the example application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "be98fa3e-ab5b-4b11-83d9-04ba2b7946bc"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User", "Application"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
    value                = "admin"
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

  web {
    homepage_url  = "https://app.example.net"
    logout_url    = "https://app.example.net/logout"
    redirect_uris = ["https://app.example.net/account"]

    implicit_grant {
      access_token_issuance_enabled = true
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `api` - (Optional) An `api` block as documented below, which configures API related settings for this Application.
* `app_role` - (Optional) A collection of `app_role` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
* `available_to_other_tenants` - (Optional, **Deprecated**) Is this Azure AD Application available to other tenants? Defaults to `false`. This property is deprecated and has been replaced by the `sign_in_audience` property.
* `display_name` - (Required) The display name for the application.
* `fallback_public_client_enabled` - (Optional) The fallback application type as public client, such as an installed application running on a mobile device. Defaults to `false`.
* `group_membership_claims` - (Optional) Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
* `homepage` - (Optional, **Deprecated**) The URL to the application's home page. This property is deprecated and has been replaced by the `homepage_url` property in the `web` block.
* `identifier_uris` - (Optional) The user-defined URI(s) that uniquely identify an application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
* `logout_url` - (Optional, **Deprecated**) The URL of the logout page. This property is deprecated and has been replaced by the `logout_url` property in the `web` block.
* `oauth2_allow_implicit_flow` - (Optional, **Deprecated**) Does this Azure AD Application allow OAuth 2.0 implicit flow tokens? Defaults to `false`. This property is deprecated and has been replaced by the `access_token_issuance_enabled` property in the `implicit_grant` block.
* `oauth2_permissions` - (Optional, **Deprecated**) A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2_permissions` blocks as documented below. This block is deprecated and has been replaced by the `oauth2_permission_scope` block in the `api` block.
* `optional_claims` - (Optional) A collection of `access_token` or `id_token` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
* `owners` - (Optional) A list of object IDs of principals that will be granted ownership of the application. It's recommended to specify the object ID of the authenticated principal running Terraform, to ensure sufficient permissions that the application can be subsequently updated.
* `prevent_duplicate_names` - (Optional) If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
* `public_client` - (Optional, **Deprecates**) Is this Azure AD Application a public client? Defaults to `false`. This property is deprecated and has been replaced by the `fallback_public_client_enabled` property.
* `reply_urls` - (Optional, **Deprecated**) A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to. This property is deprecated and has been replaced by the `redirect_uris` property in the `web` block.
* `required_resource_access` - (Optional) A collection of `required_resource_access` blocks as documented below.
* `sign_in_audience` - (Optional) The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg` or `AzureADMultipleOrgs`. Defaults to `AzureADMyOrg`.
* `type` - (Optional, **Deprecated**) The type of the application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not be set. **This legacy property is deprecated and will be removed in version 2.0 of the provider**.

~> **Note:** The `type` attribute is deprecated and will be removed in version 2.0 of the provider, along with the associated constraints of this attribute's values. Applications in Azure Active Directory are no longer differentiated by their type, instead you will be able to set native client specific attributes.

* `web` - (Optional) A `web` block as documented below, which configures web related settings for this Application.

---

`access_token` and/or `id_token` blocks support the following:

* `additional_properties` - List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.

---

`api` block supports the following:

* `oauth2_permission_scope` - (Optional) One or more `oauth2_permission_scope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this Application.

---

`app_role` block supports the following:

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
* `description` - (Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - (Required) Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - (Optional) Determines if the app role is enabled: Defaults to `true`.
* `id` - The unique identifier of the app role. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom `id`, it's recommended to use the [azuread_application_app_role](application_app_role.html) resource.

~> In version 2.0 of the provider, the `id` property will become mandatory. For more information, see the [Upgrade Guide for v2.0](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/microsoft-graph.html).

* `value` - (Optional) The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

-> **Note on roles and permission scopes:** In Azure Active Directory, roles (`app_role`) and permission scopes (`oauth2_permission_scope`) exported by an Application share the same namespace and cannot contain duplicate `value`s. Terraform will attempt to detect this at plan time.

---

`implicit_grant` block supports the following:

* `access_token_issuance_enabled` - (Optional) Whether this web application can request an access token using OAuth 2.0 implicit flow.

---

`oauth2_permission_scope` block supports the following:

* `admin_consent_description` - (Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - (Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - (Optional) Determines if the permission scope is enabled. Defaults to `true`.
* `id` - (Required) The unique identifier of the delegated permission. Must be a valid UUID.

-> **Tip: Generating a UUID for the `id` field** To generate a value for the `id` field in cases where the actual UUID is not important, you can use the `random_uuid` resource. See the [application example](https://github.com/hashicorp/terraform-provider-azuread/tree/main/examples/application) in the provider repository.

* `type` - (Required) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
* `user_consent_description` - (Optional) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - (Optional) Display name for the delegated permission that appears in the end user consent experience.
* `value` - (Optional) The value that is used for the `scp` claim in OAuth 2.0 access tokens.

If you don't specify any `oauth2_permission_scope` blocks, your Application will be assigned the default `user_impersonation` scope by Azure Active Directory. However, due to the declarative nature of Terraform configuration, if you do specify any `oauth2_permission_scope` blocks, you will need to include a block for the `user_impersonation` scope if you need it, or it will be removed (see the example above).

~> The behaviour of the default `user_impersonation` scope will change in version 2.0 of the provider. For more information, see the [Upgrade Guide for v2.0](../guides/microsoft-graph.html).

-> **Note on roles and permission scopes:** In Azure Active Directory, roles (`app_role`) and permission scopes (`oauth2_permission_scope`) exported by an Application share the same namespace and cannot contain duplicate `value`s. Terraform will attempt to detect this at plan time.

---

`oauth2_permissions` block (deprecated) supports the following:

* `admin_consent_description` - (Required) Permission help text that appears in the admin consent and app assignment experiences.
* `admin_consent_display_name` - (Required) Display name for the permission that appears in the admin consent and app assignment experiences.
* `id` - The unique identifier of the permision. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom `id`, it's recommended to migrate to the newer `oauth2_permission_scope` block, or use the [azuread_application_oauth2_permission](application_oauth2_permission.html) resource.
* `is_enabled` - (Optional) Determines if the permission is enabled: defaults to `true`.
* `type` - (Required) Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by a Company Administrator. Possible values are "User" or "Admin".
* `user_consent_description` - (Optional) Permission help text that appears in the end user consent experience.
* `user_consent_display_name` - (Optional) Display name for the permission that appears in the end user consent experience.
* `value` - (Required) The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.

If you don't specify any `oauth2_permissions` blocks, your Application will be assigned the default `user_impersonation` scope by Azure Active Directory. However, due to the declarative nature of Terraform configuration, if you do specify any `oauth2_permissions` blocks, you will need to include a block for the `user_impersonation` scope if you need it, or it will be removed (see the example above). To ensure that no OAuth 2.0 permission scopes are present for your Application, specify `oauth2_permissions = []` in your Application resource.

---

`required_resource_access` block supports the following:

* `resource_access` - (Required) A collection of `resource_access` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
* `resource_app_id` - (Required) The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.

-> **Note:** Documentation on `resource_app_id` values for Microsoft APIs can be difficult to find, but you can use the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/ad/sp?view=azure-cli-latest#az_ad_sp_list) to find them. (e.g. `az ad sp list --display-name "Microsoft Graph" --query '[].{appDisplayName:appDisplayName, appId:appId}'`)

---

`resource_access` block supports the following:

* `id` - (Required) The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
* `type` - (Required) Specifies whether the `id` property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

---

`web` block supports the following:

* `homepage_url` - (Optional) Home page or landing page of the application.
* `implicit_grant` - (Optional) An `implicit_grant` block as documented above.
* `logout_url` - (Optional) The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
* `redirect_uris` - (Optional) A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_id` - The Application ID (Also called Client ID).
* `object_id` - The application's Object ID.

## Import

Azure Active Directory Applications can be imported using the `object id`, e.g.

```shell
terraform import azuread_application.test 00000000-0000-0000-0000-000000000000
```
