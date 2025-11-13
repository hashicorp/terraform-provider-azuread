---
subcategory: "Domains"
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

* `admin_managed` - (Optional) Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
* `include_unverified` - (Optional) Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
* `only_default` - (Optional) Set to `true` to only return the default domain.
* `only_initial` - (Optional) Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
* `only_root` - (Optional) Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
* `supports_services` - (Optional) A list of supported services that must be supported by a domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.

-> **Note on filters** If `include_unverified` is set to `true`, you cannot specify `only_default` or `only_initial`. Additionally, you cannot combine `only_default` with `only_initial`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `domains` - A list of tenant domains. Each `domain` object provides the attributes documented below.

---

`domain` object exports the following:

* `admin_managed` - Whether the DNS for the domain is managed by Microsoft 365.
* `authentication_type` - The authentication type of the domain. Possible values include `Managed` or `Federated`.
* `domain_name` - The name of the domain.
* `default` - Whether this is the default domain that is used for user creation.
* `initial` - Whether this is the initial domain created by Azure Active Directory.
* `root` - Whether the domain is a verified root domain (not a subdomain).
* `verified` - Whether the domain has completed domain ownership verification.
* `supported_services` - A list of capabilities / services supported by the domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
