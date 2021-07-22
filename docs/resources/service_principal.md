---
subcategory: "Service Principals"
---

# Resource: azuread_service_principal

Manages a service principal associated with an application within Azure Active Directory.

## Example Usage

```terraform
resource "azuread_application" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  application_id               = azuread_application.example.application_id
  app_role_assignment_required = false

  tags = ["example", "tags", "here"]
}
```

## Argument Reference

The following arguments are supported:

* `account_enabled` - (Optional) Whether or not the service principal account is enabled. Defaults to `true`.
* `alternative_names` - (Optional) A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
* `app_role_assignment_required` - (Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
* `application_id` - (Required) The application ID (client ID) of the application for which to create a service principal.
* `description` - (Optional) A description of the service principal provided for internal end-users.
* `login_url` - (Optional) The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
* `notes` - (Optional) A free text field to capture information about the service principal, typically used for operational purposes.
* `notification_email_addresses` - (Optional) A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
* `preferred_single_sign_on_mode` - (Optional) The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
* `tags` - (Optional) A set of tags to apply to the service principal.
* `use_existing` - (Optional) When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.

-> **Caveats of `use_existing`** Enabling this behaviour is useful for managing existing service principals that may already be installed in your tenant for Microsoft-published APIs, as it allows you to make changes where permitted, and then also reference them in your Terraform configuration. However, the behaviour of delete operations is also affected - when `use_existing` is `true`, Terraform will still attempt to delete the service principal on destroy, although it will not raise an error if the deletion fails (as it often the case for first-party Microsoft applications).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_role_ids` - A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
* `app_roles` - A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `display_name` - The display name of the application associated with this service principal.
* `homepage_url` - Home page or landing page of the associated application.
* `logout_url` - The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
* `oauth2_permission_scope_ids` - A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
* `oauth2_permission_scopes` - A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
* `object_id` - The object ID of the service principal.
* `redirect_uris` - A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
* `type` - Identifies whether the service principal represents an application or a managed identity. Possible values include `Application` or `ManagedIdentity`.

---

`app_roles` is a list of objects with the following attributes:

* `allowed_member_types` - Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: `User` and `Application`, or both.
* `description` - Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - Determines if the app role is enabled.
* `id` - The unique identifier of the `app_role`.
* `value` - The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

---

`oauth2_permission_scopes` is a list of objects with the following attributes:

* `admin_consent_description` - Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - Specifies whether the permission scope is enabled.
* `id` - The unique identifier of the delegated permission.
* `type` - Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
* `user_consent_description` - Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - Display name for the delegated permission that appears in the end user consent experience.
* `value` - The value that is used for the `scp` claim in OAuth 2.0 access tokens.

## Import

Service principals can be imported using their object ID, e.g.

```shell
terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000
```
