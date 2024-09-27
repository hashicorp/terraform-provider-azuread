---
subcategory: "Service Principals"
---

# Data Source: azuread_service_principals

Gets basic information for multiple Azure Active Directory service principals.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up by application display names*

```terraform
data "azuread_service_principals" "example" {
  display_names = [
    "example-app",
    "another-app",
  ]
}
```

*Look up by application IDs (client IDs)*

```terraform
data "azuread_service_principals" "example" {
  client_ids = [
    "11111111-0000-0000-0000-000000000000",
    "22222222-0000-0000-0000-000000000000",
    "33333333-0000-0000-0000-000000000000",
  ]
}
```

*Look up by service principal object IDs*

```terraform
data "azuread_service_principals" "example" {
  object_ids = [
    "00000000-0000-0000-0000-000000000000",
    "00000000-0000-0000-0000-111111111111",
    "00000000-0000-0000-0000-222222222222",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `client_ids` - (Optional) A list of client IDs of the applications associated with the service principals.
* `display_names` - (Optional) A list of display names of the applications associated with the service principals.
* `ignore_missing` - (Optional) Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
* `object_ids` - (Optional) The object IDs of the service principals.
* `return_all` - (Optional) When `true`, the data source will return all service principals. Cannot be used with `ignore_missing`. Defaults to false.

~> Either `return_all`, or one of `client_ids`, `display_names` or `object_ids` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `client_ids` - A list of client IDs of the applications associated with the service principals.
* `display_names` - A list of display names of the applications associated with the service principals.
* `object_ids` - The object IDs of the service principals.
* `service_principals` - A list of service principals. Each `service_principal` object provides the attributes documented below.

---

`service_principal` object exports the following:

* `account_enabled` - Whether the service principal account is enabled.
* `app_role_assignment_required` - Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `client_ids` - The client ID of the application associated with this service principal.
* `display_name` - The display name of the application associated with this service principal.
* `object_id` - The object ID of the service principal.
* `preferred_single_sign_on_mode` - The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
* `tags` - A list of tags applied to the service principal.
* `type` - Identifies whether the service principal represents an application or a managed identity. Possible values include `Application` or `ManagedIdentity`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
