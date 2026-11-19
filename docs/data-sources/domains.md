---
page_title: "Data Source: azuread_domains"
subcategory: "Domains"
description: |-
  Use this data source to access information about existing Domains within Azure Active Directory.
---

# Data Source: azuread_domains

Use this data source to access information about existing Domains within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Domain.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

```terraform
data "azuread_domains" "aad_domains" {}

output "domain_names" {
  value = data.azuread_domains.aad_domains.domains.*.domain_name
}
```

## Arguments Reference

The following arguments are supported:

* `admin_managed` - (Optional) Set to `true` to only return domains whose DNS is managed by Microsoft 365.
* `include_unverified` - (Optional) Set to `true` if unverified Azure AD domains should be included. Conflicts with `only_default` and `only_initial`.
* `only_default` - (Optional) Set to `true` to only return the default domain. Conflicts with `only_initial` and `only_root`.
* `only_initial` - (Optional) Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Conflicts with `only_default` and `only_root`.
* `only_root` - (Optional) Set to `true` to only return verified root domains. Excludes subdomains and unverified domains. Conflicts with `only_default` and `only_initial`.
* `supports_services` - (Optional) A list of supported services that must be supported by a domain.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `domains` - A list of tenant domains. (see [nested schema](#schema-6e6573746564-6f7574707574-646f6d61696e73))

---

<a id="schema-6e6573746564-6f7574707574-646f6d61696e73"></a>
The `domains` object exports the following:

* `admin_managed` - Whether the DNS for the domain is managed by Microsoft 365.
* `authentication_type` - The authentication type of the domain. Possible values include `Managed` or `Federated`.
* `default` - Whether this is the default domain that is used for user creation.
* `domain_name` - The name of the domain.
* `initial` - Whether this is the initial domain created by Azure Active Directory.
* `root` - Whether the domain is a verified root domain (not a subdomain).
* `supported_services` - A list of capabilities / services supported by the domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
* `verified` - Whether the domain has completed domain ownership verification.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data-source.
