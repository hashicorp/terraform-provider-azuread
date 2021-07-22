---
subcategory: "Service Principals"
---

# Data Source: azuread_service_principal

Gets information about an existing service principal associated with an application within Azure Active Directory.

## Example Usage

**Look up by application display name**

```terraform
data "azuread_service_principal" "example" {
  display_name = "my-awesome-application"
}
```

**Look up by application ID**

```terraform
data "azuread_service_principal" "example" {
  application_id = "00000000-0000-0000-0000-000000000000"
}
```

**Look up by service principal object ID**

```terraform
data "azuread_service_principal" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The application ID (client ID) of the application associated with this service principal.
* `display_name` - (Optional) The display name of the application associated with this service principal.
* `object_id` - (Optional) The object ID of the service principal.

~> **NOTE:** At least one of `application_id`, `display_name` or `object_id` must be specified.

## Attributes Reference

The following attributes are exported:

* `account_enabled` - - Whether or not the service principal account is enabled.
* `alternative_names` - A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
* `app_role_assignment_required` - Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
* `app_role_ids` - A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
* `app_roles` - A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `description` - A description of the service principal provided for internal end-users.
* `homepage_url` - Home page or landing page of the associated application.
* `login_url` - The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
* `logout_url` - The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
* `notes` - A free text field to capture information about the service principal, typically used for operational purposes.
* `notification_email_addresses` - A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
* `object_id` - The object ID for the service principal.
* `oauth2_permission_scope_ids` - A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
* `oauth2_permission_scopes` - A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
* `redirect_uris` - A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
* `tags` - A list of tags applied to the service principal.
* `type` - Identifies whether the service principal represents an application or a managed identity. Possible values include `Application` or `ManagedIdentity`.

---

`app_roles` block exports the following:

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: `User` and `Application`, or both.
* `description` - Permission help text that appears in the admin app assignment and consent experiences.
* `display_name` - Display name for the permission that appears in the admin consent and app assignment experiences.
* `id` - The unique identifier of the app role.
* `is_enabled` - Determines if the app role is enabled.
* `value` - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

---

`oauth2_permission_scopes` block exports the following:

* `admin_consent_description` - Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - Determines if the permission scope is enabled.
* `id` - The unique identifier of the delegated permission. Must be a valid UUID.
* `type` - Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
* `user_consent_description` - Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - Display name for the delegated permission that appears in the end user consent experience.
* `value` - The value that is used for the `scp` claim in OAuth 2.0 access tokens.
