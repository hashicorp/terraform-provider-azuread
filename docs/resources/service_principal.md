---
subcategory: "Service Principals"
---

# Resource: azuread_service_principal

Manages a service principal associated with an application within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

It is not currently possible to manage service principals whilst having only the `Application.ReadWrite.OwnedBy` role granted.

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

*Create a service principal for an application*

```terraform
data "azuread_client_config" "current" {}

resource "azuread_application" "example" {
  display_name = "example"
  owners       = [data.azuread_client_config.current.object_id]
}

resource "azuread_service_principal" "example" {
  application_id               = azuread_application.example.application_id
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
  application_id               = azuread_application.example.application_id
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
  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing   = true
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
  application_id = azuread_application.example.application_id
  use_existing   = true
}
```

*Create a service principal with service-principal-only AppRoles, eg for AWS SSO federation*

```terraform
resource "random_uuid" "example_approle1" {}

resource "azuread_application "example" {
  display_name = "example"
  ...
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application.example.application_id

  sp_app_role {
    id = random_uuid.example_approle1.result
    allowed_member_types = [ "User" ]
    description          = "111111111111-myrole,WAAD"
    display_name         = "myrole"
    enabled              = true
    value                = "arn:aws:iam::111111111111:role/myrole,arn:aws:iam::111111111111:saml-privder/WAAD"
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_enabled` - (Optional) Whether or not the service principal account is enabled. Defaults to `true`.
* `alternative_names` - (Optional) A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
* `app_role_assignment_required` - (Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
* `application_id` - (Required) The application ID (client ID) of the application for which to create a service principal.
* `description` - (Optional) A description of the service principal provided for internal end-users.
* `feature_tags` - (Optional) A `feature_tags` block as described below. Cannot be used together with the `tags` property.

-> **Features and Tags** Features are configured for a service principal using tags, and are provided as a shortcut to set the corresponding magic tag value for each feature. You cannot configure `feature_tags` and `tags` for a service principal at the same time, so if you need to assign additional custom tags it's recommended to use the `tags` property instead. Any tags configured for the linked application will propagate to this service principal.

* `login_url` - (Optional) The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
* `notes` - (Optional) A free text field to capture information about the service principal, typically used for operational purposes.
* `notification_email_addresses` - (Optional) A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
* `owners` - (Optional) A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.

-> **Ownership of Service Principals** It's recommended to always specify one or more service principal owners, including the principal being used to execute Terraform, such as in the example above.

* `preferred_single_sign_on_mode` - (Optional) The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
* `saml_single_sign_on` - (Optional) A `saml_single_sign_on` block as documented below.
* `sp_app_role` - (Optional)  A collection of `sp_app_role` blocks as documented below. Note that `sp_app_role` are AppRoles which only exist on the Service Principal object. The resultant attribute `approles` below will contain both the underlying Application AppRoles and the Service Principal AppRoles.
* `tags` - (Optional) A set of tags to apply to the service principal. Cannot be used together with the `feature_tags` block.

-> **Tags and Features** Azure Active Directory uses special tag values to configure the behavior of service principals. These can be specified using either the `tags` property or with the `feature_tags` block. If you need to set any custom tag values not supported by the `feature_tags` block, it's recommended to use the `tags` property. Tag values set for the linked application will also propagate to this service principal.

* `use_existing` - (Optional) When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.

-> **Caveats of `use_existing`** Enabling this behaviour is useful for managing existing service principals that may already be installed in your tenant for Microsoft-published APIs, as it allows you to make changes where permitted, and then also reference them in your Terraform configuration. However, the behaviour of delete operations is also affected - when `use_existing` is `true`, Terraform will still attempt to delete the service principal on destroy, although it will not raise an error if the deletion fails (as it often the case for first-party Microsoft applications).

---

`feature_tags` block supports the following:

* `custom_single_sign_on` - (Optional) Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
* `enterprise` - (Optional) Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
* `gallery` - (Optional) Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
* `hide` - (Optional) Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.

---

`saml_single_sign_on` supports the following:

* `relay_state` - (Optional) The relative URI the service provider would redirect to after completion of the single sign-on flow.

---

`sp_app_role` block supports the following:

* `allowed_member_types` - (Required) Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
* `description` - (Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
* `display_name` - (Required) Display name for the app role that appears during app role assignment and in consent experiences.
* `enabled` - (Optional) Determines if the app role is enabled. Defaults to `true`.
* `id` - (Required) The unique identifier of the app role. Must be a valid UUID.

-> **Tip: Generating a UUID for the `id` field** To generate a value for the `id` field in cases where the actual UUID is not important, you can use the `random_uuid` resource. See the [application example](https://github.com/hashicorp/terraform-provider-azuread/tree/main/examples/application) in the provider repository.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_role_ids` - A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
* `app_roles` - A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `display_name` - The display name of the application associated with this service principal.
* `homepage_url` - Home page or landing page of the associated application.
* `logout_url` - The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
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
