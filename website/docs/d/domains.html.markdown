---
layout: "azuread"
page_title: "Azure Active Directory: azuread_domains"
sidebar_current: "docs-azuread-datasource-azuread-domains"
description: |-
  Gets information about an existing Domains within Azure Active Directory.
---

# Data Source: azuread_domains

Use this data source to access information about an existing Domains within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.Read.All` within the `Windows Azure Active Directory` API.

## Example Usage

```hcl
data "azuread_domains" "aad_domains" {}

output "domains" {
  value = "${data.azuread_domains.aad_domains.domains}"
}
```

## Argument Reference

* `include_unverified` - (Optional) Set to `true` if unverified Azure AD Domains should be included. Defaults to `false`.
* `only_default` - (Optional) Set to `true` to only return the default domain.
* `only_initial` - (Optional) Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.

-> **NOTE:** If `include_unverified` is set to `true` you cannot specify `only_default` or `only_initial`. Additionally you cannot combine `only_default` with `only_initial`.

## Attributes Reference

* `domains` - One or more `domain` blocks as defined below.

The `domain` block contains:

* `domain_name` - The name of the domain.
* `authentication_type` - The authentication type of the domain (Managed or Federated).
* `is_default` - `True` if this is the default domain that is used for user creation.
* `is_initial` - `True` if this is the initial domain created by Azure Activie Directory.
* `is_verified` - `True` if the domain has completed domain ownership verification.
