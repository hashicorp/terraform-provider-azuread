---
page_title: "Data Source: azuread_service_principals"
subcategory: "Service Principals"
description: |-
  Gets basic information for multiple Azure Active Directory service principals.
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

## Arguments Reference

The following arguments are supported:

* `client_ids` - (Optional) A list of client IDs of the applications associated with the service principals. Exactly one of `client_ids`, `display_names`, `object_ids`, or `return_all` must be specified.
* `display_names` - (Optional) A list of display names of the applications associated with the service principals. Exactly one of `client_ids`, `display_names`, `object_ids`, or `return_all` must be specified.
* `ignore_missing` - (Optional) Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to `false`. Conflicts with `return_all`.
* `object_ids` - (Optional) The object IDs of the service principals. Exactly one of `client_ids`, `display_names`, `object_ids`, or `return_all` must be specified.
* `return_all` - (Optional) When `true`, the data source will return all service principals. Defaults to `false`. Conflicts with `ignore_missing`. Exactly one of `client_ids`, `display_names`, `object_ids`, or `return_all` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `service_principals` - A list of service principals. (see [nested schema](#schema-6e6573746564-6f7574707574-736572766963655f7072696e636970616c73))

---

<a id="schema-6e6573746564-6f7574707574-736572766963655f7072696e636970616c73"></a>
The `service_principals` object exports the following:

* `account_enabled` - Whether the service principal account is enabled.
* `app_role_assignment_required` - Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
* `application_tenant_id` - The tenant ID where the associated application is registered.
* `client_id` - The client ID of the application associated with this service principal.
* `display_name` - The display name of the application associated with this service principal.
* `object_id` - The object ID of the service principal.
* `preferred_single_sign_on_mode` - The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
* `saml_metadata_url` - The URL where the service exposes SAML metadata for federation.
* `service_principal_names` - A list of identifier URI(s), copied over from the associated application.
* `sign_in_audience` - The Microsoft account types that are supported for the associated application.
* `tags` - A list of tags applied to the service principal.
* `type` - Identifies whether the service principal represents an application or a managed identity.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
