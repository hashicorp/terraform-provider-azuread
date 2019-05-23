---
layout: "azuread"
page_title: "Provider: Azure Active Directory"
sidebar_current: "docs-azuread-index"
description: |-
  The AzureAD Provider is used to interact with the many resources supported by Azure Active Directory.

---

# Azure Active Directory Provider

The Azure Provider can be used to configure infrastructure in [Azure Active Directory](https://azure.microsoft.com/en-us/services/active-directory/) using the Azure Resource Manager API's. Documentation regarding the [Data Sources](/docs/configuration/data-sources.html) and [Resources](/docs/configuration/resources.html) supported by the Azure Active Directory Provider can be found in the navigation to the left.

Interested in the provider's latest features, or want to make sure you're up to date? Check out the [changelog](https://github.com/terraform-providers/terraform-provider-azuread/blob/master/CHANGELOG.md) for version information and release notes.

## Authenticating to Azure Active Directory

Terraform supports a number of different methods for authenticating to Azure Active Directory:

* [Authenticating to Azure Active Directory using the Azure CLI](auth/azure_cli.html)
* [Authenticating to Azure Active Directory using Managed Service Identity](auth/managed_service_identity.html)
* [Authenticating to Azure Active Directory using a Service Principal and a Client Certificate](auth/service_principal_client_certificate.html)
* [Authenticating to Azure Active Directory using a Service Principal and a Client Secret](auth/service_principal_client_secret.html)

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

## Example Usage

```hcl
# Configure the Microsoft Azure Active Directory Provider
provider "azuread" {
  version = "=0.3.0"
}

# Create an application
resource "azuread_application" "example" {
  name = "ExampleApp"
}

# Create a service principal
resource "azuread_service_principal" "example" {
  application_id = "${azuread_application.example.application_id}"
}
```

## Features and Bug Requests

The Azure Active Directory provider's bugs and feature requests can be found in the [GitHub repo issues](https://github.com/terraform-providers/terraform-provider-azuread/issues).
Please avoid "me too" or "+1" comments. Instead, use a thumbs up [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/)
on enhancement requests. Provider maintainers will often prioritize work based on the number of thumbs on an issue.

Community input is appreciated on outstanding issues! We love to hear what use
cases you have for new features, and want to provide the best possible
experience for you using the Azure Active Directory provider.

If you have a bug or feature request without an existing issue

* if an existing resource or field is working in an unexpected way, [file a bug](https://github.com/terraform-providers/terraform-provider-azuread/issues/new?template=bug.md).

* if you'd like the provider to support a new resource or field, [file an enhancement/feature request](https://github.com/terraform-providers/terraform-provider-azuread/issues/new?template=enhancement.md).

The provider maintainers will often use the assignee field on an issue to mark
who is working on it.

* An issue assigned to an individual maintainer indicates that maintainer is working
on the issue

* If you're interested in working on an issue please leave a comment in that issue

---

If you have configuration questions, or general questions about using the provider, try checking out:

* [Terraform's community resources](https://www.terraform.io/docs/extend/community/index.html)
* [HashiCorp support](https://support.hashicorp.com) for Terraform Enterprise customers

## Argument Reference

The following arguments are supported:

* `client_id` - (Optional) The Client ID which should be used. This can also be sourced from the `ARM_CLIENT_ID` Environment Variable.

* `environment` - (Optional) The Cloud Environment which be used. Possible values are `public`, `usgovernment`, `german` and `china`. Defaults to `public`. This can also be sourced from the `ARM_ENVIRONMENT` environment variable.

* `subscription_id` - (Optional) The Subscription ID which should be used. This can also be sourced from the `ARM_SUBSCRIPTION_ID` Environment Variable.

* `tenant_id` - (Optional) The Tenant ID which should be used. This can also be sourced from the `ARM_TENANT_ID` Environment Variable.

---

When authenticating as a Service Principal using a Client Certificate, the following fields can be set:

* `client_certificate_password` - (Optional) The password associated with the Client Certificate. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PASSWORD` Environment Variable.

* `client_certificate_path` - (Optional) The path to the Client Certificate associated with the Service Principal which should be used. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` Environment Variable.

More information on [how to configure a Service Principal using a Client Certificate can be found in this guide](auth/service_principal_client_certificate.html).

---

When authenticating as a Service Principal using a Client Secret, the following fields can be set:

* `client_secret` - (Optional) The Client Secret which should be used. This can also be sourced from the `ARM_CLIENT_SECRET` Environment Variable.

More information on [how to configure a Service Principal using a Client Secret can be found in this guide](auth/service_principal_client_secret.html).

---

When authenticating using Managed Service Identity, the following fields can be set:

* `msi_endpoint` - (Optional) The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically. This can also be sourced from the `ARM_MSI_ENDPOINT` Environment Variable.

* `use_msi` - (Optional) Should Managed Service Identity be used for Authentication? This can also be sourced from the `ARM_USE_MSI` Environment Variable. Defaults to `false`.

More information on [how to configure a Service Principal using Managed Service Identity can be found in this guide](auth/managed_service_identity.html).

---

It's also possible to use multiple Provider blocks within a single Terraform configuration, for example to work with resources across multiple Azure Active Directory Environments - more information can be found [in the documentation for Providers](https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances).
