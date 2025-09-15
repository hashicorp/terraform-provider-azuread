---
subcategory: "Applications"
---

# Resource: azuread_application

Manages an application registration within Azure Active Directory.

For a more lightweight alternative, please see the [azuread_application_registration](application_registration.html) resource. Please note that this resource should not be used together with the `azuread_application_registration` resource when managing the same application.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, you should ensure that the principal being used to run Terraform is included in the `owners` property.

Additionally, you may need the `User.Read.All` application role when including user principals in the `owners` property.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Create an application*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_application" "example" {
  display_name     = "example"
  identifier_uris  = ["api://example-app"]
  logo_image       = filebase64("/path/to/logo.png")
  owners           = [data.azuread_client_config.current.object_id]
  sign_in_audience = "AzureADMultipleOrgs"

  api {
    mapped_claims_enabled          = true
    requested_access_token_version = 2

    known_client_applications = [
      azuread_application.known1.client_id,
      azuread_application.known2.client_id,
    ]

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
    enabled              = true
    id                   = "1b19509b-32b1-4e9f-b71d-4992aa991967"
    value                = "admin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "497406e4-012a-4267-bf18-45a1cb148a01"
    value                = "User"
  }

  feature_tags {
    enterprise = true
    gallery    = true
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

    saml2_token {
      name = "samlexample"
    }
  }

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000" # Microsoft Graph

    resource_access {
      id   = "df021288-bdef-4463-88db-98f22de89214" # User.Read.All
      type = "Role"
    }

    resource_access {
      id   = "b4e74841-8e56-480b-be8b-910348b18b4c" # User.ReadWrite
      type = "Scope"
    }
  }

  required_resource_access {
    resource_app_id = "c5393580-f805-4401-95e8-94b7a6ef2fc2" # Office 365 Management

    resource_access {
      id   = "594c1fb6-4f81-4475-ae41-0c394909246c" # ActivityFeed.Read
      type = "Role"
    }
  }

  web {
    homepage_url  = "https://app.example.net"
    logout_url    = "https://app.example.net/logout"
    redirect_uris = ["https://app.example.net/account"]

    implicit_grant {
      access_token_issuance_enabled = true
      id_token_issuance_enabled     = true
    }
  }
}
```

*Create application and generate a password*

```terraform
data "azuread_client_config" "current" {}

resource "time_rotating" "example" {
  rotation_days = 180
}

resource "azuread_application" "example" {
  display_name = "example"
  owners       = [data.azuread_client_config.current.object_id]

  password {
    display_name = "MySecret-1"
    start_date   = time_rotating.example.id
    end_date     = timeadd(time_rotating.example.id, "4320h")
  }
}

output "example_password" {
  sensitive = true
  value     = tolist(azuread_application.example.password).0.value
}

```

*Create application from a gallery template*

```terraform
data "azuread_application_template" "example" {
  display_name = "Marketo"
}

resource "azuread_application" "example" {
  display_name = "example"
  template_id  = data.azuread_application_template.example.template_id
}

resource "azuread_service_principal" "example" {
  client_id    = azuread_application.example.client_id
  use_existing = true
}
```

## Argument Reference

The following arguments are supported:

* `api` - (Optional) An `api` block as documented below, which configures API related settings for this application.
* `app_role` - (Optional) A collection of `app_role` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
* `description` - (Optional) A description of the application, as shown to end users.
* `device_only_auth_enabled` - (Optional) Specifies whether this application supports device authentication without a user. Defaults to `false`.
* `display_name` - (Required) The display name for the application.
* `fallback_public_client_enabled` - (Optional) Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
* `native_authentication_apis_enabled` - (Optional )The native application has native authentication enabled, allowing public clients to directly interact with users to collect credentials without rediecting to browsers for authentication. Defaults to `none`. Possible Values are `all` and `none`
* `feature_tags` - (Optional) A `feature_tags` block as described below. Cannot be used together with the `tags` property.

-> **Features and Tags** Features are configured for an application using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for an application at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Tag values also propagate to any linked service principals.

* `group_membership_claims` - (Optional) A set of strings containing membership claims issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
* `identifier_uris` - (Optional) A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
* `logo_image` - (Optional) A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
* `marketing_url` - (Optional) URL of the application's marketing page.
* `notes` - (Optional) User-specified notes relevant for the management of the application.
* `oauth2_post_response_required` - (Optional) Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
* `optional_claims` - (Optional) An `optional_claims` block as documented below.
* `owners` - (Optional) A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.

-> **Ownership of Applications** It's recommended to always specify one or more application owners, including the principal being used to execute Terraform, such as in the example above.

* `password` - (Optional) A single `password` block as documented below. The password is generated during creation. By default, no password is generated.

-> **Creating a Password** The `password` block supports a single password for the application, and is provided so that a password can be generated when a new application is created. This helps to make new applications available for authentication more quickly. To add additional passwords to an application, see the [azuread_application_password](application_password.html) resource.

* `prevent_duplicate_names` - (Optional) If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
* `privacy_statement_url` - (Optional) URL of the application's privacy statement.
* `public_client` - (Optional) A `public_client` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
* `required_resource_access` - (Optional) A collection of `required_resource_access` blocks as documented below.
* `service_management_reference` - (Optional) References application context information from a Service or Asset Management database.
* `sign_in_audience` - (Optional) The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.

~> **Changing `sign_in_audience` for existing applications** When updating an existing application to use a `sign_in_audience` value of `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`, your configuration may no longer be valid. Refer to [official documentation](https://docs.microsoft.com/en-gb/azure/active-directory/develop/supported-accounts-validation) to understand the differences in supported configurations. Where possible, the provider will attempt to validate your configuration and try to avoid applying unsupported settings to your application.

* `single_page_application` - (Optional) A `single_page_application` block as documented below, which configures single-page application (SPA) related settings for this application.
* `support_url` - (Optional) URL of the application's support page.
* `tags` - (Optional) A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the `feature_tags` block.

-> **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of applications. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it's recommended to use the `tags` property. Tag values also propagate to any linked service principals.

* `template_id` - (Optional) Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.

-> **Tip for Gallery Applications** This resource can  be used to instantiate a gallery application, however it will also attempt to manage the properties of the resulting application. If this is not desired, consider using the [azuread_application_registration](application_registration.html) resource instead.

* `terms_of_service_url` - (Optional) URL of the application's terms of service statement.
* `web` - (Optional) A `web` block as documented below, which configures web related settings for this application.

-> **Application Name Uniqueness** Application names are not unique within Azure Active Directory. Use the `prevent_duplicate_names` argument to check for existing applications if you want to avoid name collisions.

---

`api` block supports the following:

* `known_client_applications` - (Optional) A set of client IDs, used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
* `mapped_claims_enabled` - (Optional) Allows an application to use claims mapping without specifying a custom signing key. Defaults to `false`.
* `oauth2_permission_scope` - (Optional) One or more `oauth2_permission_scope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
* `requested_access_token_version` - (Optional) The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `1`.

---

`oauth2_permission_scope` blocks support the following:

* `admin_consent_description` - (Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - (Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - (Optional) Determines if the permission scope is enabled. Defaults to `true`.
* `id` - (Required) The unique identifier of the delegated permission. Must be a valid UUID.

-> **Tip: Generating a UUID for the `id` field** To generate a value for the `id` field in cases where the actual UUID is not important, you can use the `random_uuid` resource. See the [application example](https://github.com/hashicorp/terraform-provider-azuread/tree/main/examples/application) in the provider repository.

* `type` - (Optional) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
* `user_consent_description` - (Optional) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - (Optional) Display name for the delegated permission that appears in the end user consent experience.
* `value` - (Optional) The value that is used for the `scp` claim in OAuth 2.0 access tokens.

~> **Default `user_impersonation` Scope** Unlike the Azure Portal, applications created with the Terraform AzureAD provider do not get assigned a default `user_impersonation` scope. You will need to include a block for the `user_impersonation` scope if you need it for your application.

-> **Roles and Permission Scopes** In Azure Active Directory, application roles (`app_role`) and permission scopes (`oauth2_permission_scope`) exported by an application share the same namespace and cannot contain duplicate `value`s. Terraform will attempt to detect this during a plan or apply operation.

---

`app_role` block supports the following:

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
* `description` - (Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - (Required) Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - (Optional) Determines if the app role is enabled. Defaults to `true`.
* `id` - (Required) The unique identifier of the app role. Must be a valid UUID.

-> **Tip: Generating a UUID for the `id` field** To generate a value for the `id` field in cases where the actual UUID is not important, you can use the `random_uuid` resource. See the [application example](https://github.com/hashicorp/terraform-provider-azuread/tree/main/examples/application) in the provider repository.

* `value` - (Optional) The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

-> **Roles and Permission Scopes** In Azure Active Directory, application roles (`app_role`) and permission scopes (`oauth2_permission_scope`) exported by an application share the same namespace and cannot contain duplicate `value`s. Terraform will attempt to detect this during a plan or apply operation.

---

`feature_tags` block supports the following:

* `custom_single_sign_on` - (Optional) Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
* `enterprise` - (Optional) Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
* `gallery` - (Optional) Whether this application represents a gallery application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
* `hide` - (Optional) Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.

---

`optional_claims` block supports the following:

* `access_token` - (Optional) One or more `access_token` blocks as documented below.
* `id_token` - (Optional) One or more `id_token` blocks as documented below.
* `saml2_token` - (Optional) One or more `saml2_token` blocks as documented below.

---

`access_token`, `id_token` and `saml2_token` blocks support the following:

* `additional_properties` - List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.

---

`password` block supports the following:

* `display_name` - (Required) A display name for the password. Changing this field forces a new resource to be created.
* `end_date` - (Optional) The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
* `start_date` - (Optional) The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.

---

`public_client` block supports the following:

* `redirect_uris` - (Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` or `ms-appx-web` URL.

---

`required_resource_access` block supports the following:

* `resource_access` - (Required) A collection of `resource_access` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
* `resource_app_id` - (Required) The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.

-> **Note:** Documentation on `resource_app_id` values for Microsoft APIs can be difficult to find, but you can use the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/ad/sp?view=azure-cli-latest#az_ad_sp_list) to find them. (e.g. `az ad sp list --display-name "Microsoft Graph" --query '[].{appDisplayName:appDisplayName, appId:appId}'`)

---

`resource_access` block supports the following:

* `id` - (Required) The unique identifier for an app role or OAuth2 permission scope published by the resource application.
* `type` - (Required) Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.

---

`single_page_application` block supports the following:

* `redirect_uris` - (Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` URL.

---

`web` block supports the following:

* `homepage_url` - (Optional) Home page or landing page of the application.
* `implicit_grant` - (Optional) An `implicit_grant` block as documented above.
* `logout_url` - (Optional) The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
* `redirect_uris` - (Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `http` URL or a URN.

---

`implicit_grant` block supports the following:

* `access_token_issuance_enabled` - (Optional) Whether this web application can request an access token using OAuth 2.0 implicit flow.
* `id_token_issuance_enabled` - (Optional) Whether this web application can request an ID token using OAuth 2.0 implicit flow.

---

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_role_ids` - A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
* `client_id` - The Client ID for the application.
* `disabled_by_microsoft` - Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
* `id` - The Terraform resource ID for the application, for use when referencing this resource in your Terraform configuration.
* `logo_url` - CDN URL to the application's logo, as uploaded with the `logo_image` property.
* `oauth2_permission_scope_ids` - A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
* `object_id` - The application's object ID.
* `password` - A `password` block as documented below. Note that this block is a set rather than a list, and you will need to convert or iterate it to address its attributes (see the usage example above).
* `publisher_domain` - The verified publisher domain for the application.

---

`password` block exports the following:

* `key_id` - (Required) The unique key ID for the generated password.
* `value` - (Required) The generated password for the application.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Applications can be imported using the object ID of the application, in the following format.

```shell
terraform import azuread_application.example /applications/00000000-0000-0000-0000-000000000000
```
