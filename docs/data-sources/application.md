---
page_title: "Data Source: azuread_application"
subcategory: "Applications"
description: |-
  Use this data source to access information about an existing Application within Azure Active Directory.
---

# Data Source: azuread_application

Use this data source to access information about an existing Application within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_application" "example" {
  display_name = "My First AzureAD Application"
}

output "application_object_id" {
  value = data.azuread_application.example.object_id
}
```

## Arguments Reference

The following arguments are supported:

* `client_id` - (Optional) The Client ID (also called Application ID). Exactly one of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
* `display_name` - (Optional) The display name for the application. Exactly one of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
* `identifier_uri` - (Optional) One of the application's identifier URIs. Exactly one of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
* `object_id` - (Optional) The application's object ID. Exactly one of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `api` - A list of `api` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-617069))
* `app_role_ids` - Mapping of app role names to UUIDs.
* `app_roles` - List of app roles published by the application. (see [nested schema](#schema-6e6573746564-6f7574707574-6170705f726f6c6573))
* `description` - Description of the application as shown to end users.
* `device_only_auth_enabled` - Specifies whether this application supports device authentication without a user.
* `disabled_by_microsoft` - Whether Microsoft has disabled the registered application.
* `fallback_public_client_enabled` - The fallback application type as public client, such as an installed application running on a mobile device.
* `feature_tags` - Block of features configured for this application using tags. (see [nested schema](#schema-6e6573746564-6f7574707574-666561747572655f74616773))
* `group_membership_claims` - The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
* `identifier_uris` - A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
* `logo_url` - CDN URL to the application's logo.
* `marketing_url` - URL of the application's marketing page.
* `notes` - User-specified notes relevant for the management of the application.
* `oauth2_permission_scope_ids` - Mapping of OAuth2.0 permission scope names to UUIDs.
* `oauth2_post_response_required` - Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.
* `optional_claims` - A list of `optional_claims` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73))
* `owners` - A list of object IDs of principals that are assigned ownership of the application.
* `privacy_statement_url` - URL of the application's privacy statement.
* `public_client` - A list of `public_client` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-7075626c69635f636c69656e74))
* `publisher_domain` - The verified publisher domain for the application.
* `required_resource_access` - A list of `required_resource_access` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-72657175697265645f7265736f757263655f616363657373))
* `service_management_reference` - References application or service contact information from a Service or Asset Management database.
* `sign_in_audience` - The Microsoft account types that are supported for the current application.
* `single_page_application` - A list of `single_page_application` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-73696e676c655f706167655f6170706c69636174696f6e))
* `support_url` - URL of the application's support page.
* `tags` - A set of tags applied to the application.
* `terms_of_service_url` - URL of the application's terms of service statement.
* `web` - A list of `web` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-776562))

---

<a id="schema-6e6573746564-6f7574707574-617069"></a>
The `api` object exports the following:

* `known_client_applications` - Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
* `mapped_claims_enabled` - Allows an application to use claims mapping without specifying a custom signing key.
* `oauth2_permission_scopes` - List of OAuth2 permission scopes published by the application. (see [nested schema](#schema-6e6573746564-6f7574707574-617069-6f61757468325f7065726d697373696f6e5f73636f706573))
* `requested_access_token_version` - Specifies the access token version expected by this resource.
---

<a id="schema-6e6573746564-6f7574707574-617069-6f61757468325f7065726d697373696f6e5f73636f706573"></a>
The `api.oauth2_permission_scopes` object exports the following:

* `admin_consent_description` - Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
* `admin_consent_display_name` - Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
* `enabled` - Determines if the permission scope is enabled.
* `id` - The unique identifier of the delegated permission. Must be a valid UUID.
* `type` - Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
* `user_consent_description` - Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
* `user_consent_display_name` - Display name for the delegated permission that appears in the end user consent experience.
* `value` - The value that is used for the `scp` claim in OAuth 2.0 access tokens.
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

* `custom_single_sign_on` - Whether this application principal represents a custom SAML application for linked service principals.
* `enterprise` - Whether this application represents an Enterprise Application for linked service principals.
* `gallery` - Whether this application represents a gallery application for linked service principals.
* `hide` - Whether this app is invisible to users in My Apps and Office 365 Launcher. Defaults to `true`.
---

<a id="schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73"></a>
The `optional_claims` object exports the following:

* `access_token` - The `access_token` block as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-6163636573735f746f6b656e))
* `id_token` - The `id_token` block as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-69645f746f6b656e))
* `saml2_token` - The `saml2_token` block as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-73616d6c325f746f6b656e))
---

<a id="schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-6163636573735f746f6b656e"></a>
The `optional_claims.access_token` block exports the following:

* `additional_properties` - List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.
---

<a id="schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-69645f746f6b656e"></a>
The `optional_claims.id_token` block exports the following:

* `additional_properties` - List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.
---

<a id="schema-6e6573746564-6f7574707574-6f7074696f6e616c5f636c61696d73-73616d6c325f746f6b656e"></a>
The `optional_claims.saml2_token` block exports the following:

* `additional_properties` - List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.
---

<a id="schema-6e6573746564-6f7574707574-7075626c69635f636c69656e74"></a>
The `public_client` object exports the following:

* `redirect_uris` - The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
---

<a id="schema-6e6573746564-6f7574707574-72657175697265645f7265736f757263655f616363657373"></a>
The `required_resource_access` object exports the following:

* `resource_access` - A collection of `resource_access` blocks describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource. (see [nested schema](#schema-6e6573746564-6f7574707574-72657175697265645f7265736f757263655f616363657373-7265736f757263655f616363657373))
* `resource_app_id` - The unique identifier for the resource that the application requires access to. This is the Client ID (also called Application ID) of the target application.
---

<a id="schema-6e6573746564-6f7574707574-72657175697265645f7265736f757263655f616363657373-7265736f757263655f616363657373"></a>
The `required_resource_access.resource_access` object exports the following:

* `id` - The unique identifier for an app role or OAuth2 permission scope published by the resource application.
* `type` - Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
---

<a id="schema-6e6573746564-6f7574707574-73696e676c655f706167655f6170706c69636174696f6e"></a>
The `single_page_application` object exports the following:

* `redirect_uris` - The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
---

<a id="schema-6e6573746564-6f7574707574-776562"></a>
The `web` object exports the following:

* `homepage_url` - Home page or landing page of the application.
* `implicit_grant` - A list of `implicit_grant` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-776562-696d706c696369745f6772616e74))
* `logout_url` - The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
* `redirect_uris` - A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
---

<a id="schema-6e6573746564-6f7574707574-776562-696d706c696369745f6772616e74"></a>
The `web.implicit_grant` object exports the following:

* `access_token_issuance_enabled` - Whether this web application can request an access token using OAuth 2.0 implicit flow.
* `id_token_issuance_enabled` - Whether this web application can request an ID token using OAuth 2.0 implicit flow.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
