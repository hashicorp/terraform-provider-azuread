---
page_title: "Resource: azuread_service_principal"
subcategory: "Service Principals"
description: |-
  Manages a service principal associated with an application within Azure Active Directory.
---

# Resource: azuread_service_principal

Manages a service principal associated with an application within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of _both_ the linked application registration, _and_ the service principal being managed.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Create a service principal for an application*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_application" "example" {
  display_name = "example"
  owners       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal" "example" {
  client_id                    = azuread_application.example.client_id
  app_role_assignment_required = false
  owners                       = [data.azuread_client_config.current.object_id]
}
```

*Create a service principal for an enterprise application*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_application" "example" {
  display_name = "example"
  owners       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal" "example" {
  client_id                    = azuread_application.example.client_id
  app_role_assignment_required = false
  owners                       = [data.azuread_client_config.current.object_id]

  feature_tags {
    enterprise = true
    gallery    = true
  }
}
```

*Manage a service principal for a first-party Microsoft application*

```terraform
data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}
```

*Create a service principal for an application created from a gallery template*

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

## Arguments Reference

The following arguments are supported:

* `client_id` - (Required) The client ID of the application for which to create a service principal. Changing this value forces a new resource to be created.

---

* `account_enabled` - (Optional) Whether the service principal account is enabled. Defaults to `true`.
* `alternative_names` - (Optional) A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. For each element: Value must not be empty.
* `app_role_assignment_required` - (Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
* `description` - (Optional) A description of the service principal provided for internal end-users. String length must be between `0` and `1024` characters inclusive.
* `feature_tags` - (Optional) Block of features to configure for this service principal using tags. Conflicts with `features` and `tags`.
   (see [nested schema](#schema-6e6573746564-696e707574-666561747572655f74616773))

-> **Features and Tags** Features are configured for a service principal using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for a service principal at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Any tags configured for the linked application will propagate to this service principal.
* `features` - (Optional) Block of features to configure for this service principal using tags. Conflicts with `feature_tags` and `tags`.
   (see [nested schema](#schema-6e6573746564-696e707574-6665617475726573))
* `login_url` - (Optional) The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
* `notes` - (Optional) A free text field to capture information about the service principal, typically used for operational purposes. String length must be between `0` and `1024` characters inclusive.
* `notification_email_addresses` - (Optional) A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications. For each element: Value must not be empty.
* `owners` - (Optional) A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.

-> **Ownership of Service Principals** It's recommended to always specify one or more service principal owners, including the principal being used to execute Terraform, such as in the example above.
* `preferred_single_sign_on_mode` - (Optional) The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Omit this property or specify a blank string to unset. Possible values include `""`, `notSupported`, `oidc`, `password`, and `saml`.
* `saml_single_sign_on` - (Optional) Settings related to SAML single sign-on. (see [nested schema](#schema-6e6573746564-696e707574-73616d6c5f73696e676c655f7369676e5f6f6e))
* `tags` - (Optional) A set of tags to apply to the service principal for configuring specific behaviours of the service principal. Note that these are not provided for use by practitioners. Conflicts with `features` and `feature_tags`.

-> **Tags and Features** Azure Active Directory uses special tag values to configure the behaviour of service principals. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it's recommended to use the `tags` property. Tag values set for the linked application will also propagate to this service principal.
* `use_existing` - (Optional) When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.

-> **Caveats of `use_existing`** Enabling this behaviour is useful for managing existing service principals that may already be installed in your tenant for Microsoft-published APIs, as it allows you to make changes where permitted, and then also reference them in your Terraform configuration. However, the behaviour of delete operations is also affected - when `use_existing` is `true`, Terraform will still attempt to delete the service principal on destroy, although it will not raise an error if the deletion fails (as it often the case for first-party Microsoft applications).

---

<a id="schema-6e6573746564-696e707574-666561747572655f74616773"></a>
The `feature_tags` block supports the following:

* `custom_single_sign_on` - (Optional) Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag.
* `enterprise` - (Optional) Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag.
* `gallery` - (Optional) Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag.
* `hide` - (Optional) Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag.

---

<a id="schema-6e6573746564-696e707574-6665617475726573"></a>
The `features` block supports the following:

* `custom_single_sign_on_app` - (Optional) Whether this service principal represents a custom SAML application.
* `enterprise_application` - (Optional) Whether this service principal represents an Enterprise Application.
* `gallery_application` - (Optional) Whether this service principal represents a gallery application.
* `visible_to_users` - (Optional) Whether this app is visible to users in My Apps and Office 365 Launcher. Defaults to `true`.

---

<a id="schema-6e6573746564-696e707574-73616d6c5f73696e676c655f7369676e5f6f6e"></a>
The `saml_single_sign_on` block supports the following:

* `relay_state` - (Optional) The relative URI the service provider would redirect to after completion of the single sign-on flow. Value must not be empty.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `app_role_ids` - A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
* `app_roles` - A list of app roles published by the associated application. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles). (see [nested schema](#schema-6e6573746564-6f7574707574-6170705f726f6c6573))
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `display_name` - The display name of the application associated with this service principal.
* `homepage_url` - Home page or landing page of the associated application.
* `logout_url` - The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
* `oauth2_permission_scope_ids` - A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
* `oauth2_permission_scopes` - A collection of OAuth 2.0 delegated permissions exposed by the associated application. (see [nested schema](#schema-6e6573746564-6f7574707574-6f61757468325f7065726d697373696f6e5f73636f706573))
* `object_id` - The object ID of the service principal.
* `redirect_uris` - A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application.
* `type` - Identifies whether the service principal represents an application or a managed identity.

---

<a id="schema-6e6573746564-6f7574707574-6170705f726f6c6573"></a>
The `app_roles` object exports the following:

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios).
* `description` - Permission help text that appears in the admin app assignment and consent experiences.
* `display_name` - Display name for the permission that appears in the admin consent and app assignment experiences.
* `enabled` - Whether the app role is enabled.
* `id` - The unique identifier of the app role.
* `value` - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

---

<a id="schema-6e6573746564-6f7574707574-6f61757468325f7065726d697373696f6e5f73636f706573"></a>
The `oauth2_permission_scopes` object exports the following:

* `admin_consent_description` - Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - Determines if the permission scope is enabled.
* `id` - The unique identifier of the delegated permission.
* `type` - Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
* `user_consent_description` - Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - Display name for the delegated permission that appears in the end user consent experience.
* `value` - The value that is used for the `scp` claim in OAuth 2.0 access tokens.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_service_principal.example 00000000-0000-0000-0000-000000000000
```
