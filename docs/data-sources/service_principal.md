---
page_title: "Data Source: azuread_service_principal"
subcategory: "Service Principals"
description: |-
  Gets information about an existing service principal associated with an application within Azure Active Directory.
---

# Data Source: azuread_service_principal

Gets information about an existing service principal associated with an application within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up by application display name*

```terraform
data "azuread_service_principal" "example" {
  display_name = "my-awesome-application"
}
```

*Look up by client ID*

```terraform
data "azuread_service_principal" "example" {
  client_id = "00000000-0000-0000-0000-000000000000"
}
```

*Look up by service principal object ID*

```terraform
data "azuread_service_principal" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Arguments Reference

The following arguments are supported:

* `client_id` - (Optional) The client ID of the application associated with this service principal. Exactly one of `client_id`, `display_name`, or `object_id` must be specified.
* `display_name` - (Optional) The display name of the application associated with this service principal. Exactly one of `client_id`, `display_name`, or `object_id` must be specified.
* `object_id` - (Optional) The object ID of the service principal. Exactly one of `client_id`, `display_name`, or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `account_enabled` - Whether or not the service principal account is enabled.
* `alternative_names` - A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
* `app_role_assignment_required` - Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
* `app_role_ids` - Mapping of app role names to UUIDs.
* `app_roles` - A list of `app_roles` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6170705f726f6c6573))
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `description` - Description of the service principal provided for internal end-users.
* `feature_tags` - Block of features configured for this service principal using tags. (see [nested schema](#schema-6e6573746564-6f7574707574-666561747572655f74616773))
* `features` - Block of features configured for this service principal using tags. (see [nested schema](#schema-6e6573746564-6f7574707574-6665617475726573))
* `homepage_url` - Home page or landing page of the application.
* `login_url` - The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
* `logout_url` - The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
* `notes` - Free text field to capture information about the service principal, typically used for operational purposes.
* `notification_email_addresses` - List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
* `oauth2_permission_scope_ids` - Mapping of OAuth2.0 permission scope names to UUIDs.
* `oauth2_permission_scopes` - A list of `oauth2_permission_scopes` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6f61757468325f7065726d697373696f6e5f73636f706573))
* `preferred_single_sign_on_mode` - The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
* `redirect_uris` - The URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `saml_single_sign_on` - Settings related to SAML single sign-on. (see [nested schema](#schema-6e6573746564-6f7574707574-73616d6c5f73696e676c655f7369676e5f6f6e))
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application.
* `tags` - A set of tags to apply to the service principal.
* `type` - Identifies whether the service principal represents an application or a managed identity.

---

<a id="schema-6e6573746564-6f7574707574-6170705f726f6c6573"></a>
The `app_roles` object exports the following:

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are `User` or `Application`, or both.
* `description` - Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - The unique identifier of the app role.
* `id` - The unique identifier of the app role.
* `value` - The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
---

<a id="schema-6e6573746564-6f7574707574-666561747572655f74616773"></a>
The `feature_tags` object exports the following:

* `custom_single_sign_on` - Whether this service principal represents a custom SAML application.
* `enterprise` - Whether this service principal represents an Enterprise Application.
* `gallery` - Whether this service principal represents a gallery application.
* `hide` - Whether this app is invisible to users in My Apps and Office 365 Launcher.
---

<a id="schema-6e6573746564-6f7574707574-6665617475726573"></a>
The `features` object exports the following:

* `custom_single_sign_on_app` - Whether this service principal represents a custom SAML application.
* `enterprise_application` - Whether this service principal represents an Enterprise Application.
* `gallery_application` - Whether this service principal represents a gallery application.
* `visible_to_users` - Whether this app is visible to users in My Apps and Office 365 Launcher.
---

<a id="schema-6e6573746564-6f7574707574-6f61757468325f7065726d697373696f6e5f73636f706573"></a>
The `oauth2_permission_scopes` object exports the following:

* `admin_consent_description` - Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - Determines if the permission scope is enabled.
* `id` - The unique identifier of the delegated permission. Must be a valid UUID.
* `type` - Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
* `user_consent_description` - Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - Display name for the delegated permission that appears in the end user consent experience.
* `value` - The value that is used for the `scp` claim in OAuth 2.0 access tokens.
---

<a id="schema-6e6573746564-6f7574707574-73616d6c5f73696e676c655f7369676e5f6f6e"></a>
The `saml_single_sign_on` object exports the following:

* `relay_state` - The relative URI the service provider would redirect to after completion of the single sign-on flow.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
