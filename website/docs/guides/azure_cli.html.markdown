---
subcategory: "Authentication"
layout: "azuread"
page_title: "Authenticating via the Azure CLI"
description: |-
  This guide will cover how to use the Azure CLI as authentication for the Azure Active Directory Provider.

---

TEST

# Azure Active Directory Provider: Authenticating using the Azure CLI

Terraform supports a number of different methods for authenticating to Azure:

* Authenticating to Azure using the Azure CLI (covered in this guide)
* [Authenticating to Azure using Managed Service Identity](managed_service_identity.html)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

## Important Notes about Authenticating using the Azure CLI

* Terraform only supports authenticating using the `az` CLI (and this must be available on your PATH) - authenticating using the older `azure` CLI or PowerShell Az / AzureRM Cmdlets is not supported.
* Authenticating via the Azure CLI is only supported when using a User Account. If you're using a Service Principal (for example via `az login --service-principal`) you should instead authenticate via the Service Principal directly, either using a [Client Certificate](service_principal_client_certificate.html) or a [Client Secret](service_principal_client_secret.html).

---

## Logging into the Azure CLI

~> **Using other clouds** If you're using the **China**, **German** or **Government** Azure Clouds - you'll need to first configure the Azure CLI to work with that Cloud.  You can do this by running:

```shell
$ az cloud set --name AzureChinaCloud|AzureGermanCloud|AzureUSGovernment
```

---

Firstly, login to the Azure CLI using:

```shell
$ az login --allow-no-subscriptions
```

Once logged in - it's possible to list the Subscriptions and Tenants associated with the account via:

```shell
$ az account list
```

The output (similar to below) will display one or more Tenants and/or Subscriptions.

```json
[
  {
    "cloudName": "AzureCloud",
    "id": "00000000-0000-0000-0000-000000000000",
    "isDefault": true,
    "name": "PAYG Subscription",
    "state": "Enabled",
    "tenantId": "00000000-0000-0000-0000-000000000000",
    "user": {
      "name": "user@example.com",
      "type": "user"
    }
  }
]
```

The provider will select the tenant ID from your default Azure CLI account. If you have more than one tenant listed in the output of `az account list`, for example if you are a guest user in other tenants, you can specify the tenant to use.

```shell
$ export ARM_TENANT_ID=00000000-0000-0000-0000-000000000000
```

You can also configure the tenant ID from within the provider block.

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider being used
  version = "=1.1.0"

  tenant_id = "00000000-0000-0000-0000-000000000000"
}
```

Alternatively, you can configure the Azure CLI to authenticate against the tenant you are managing with Terraform.

```bash
$ az login --allow-no-subscriptions --tenant "TENANT_ID_OR_DOMAIN"
```

-> **Tenants and Subscriptions** The AzureAD provider operates on tenants and not on subscriptions. We recommend always specifying `az login --allow-no-subscription` as it will force the Azure CLI to report tenants with no associated subscriptions, or if your user account does not have any roles assigned against your subscriptions.

---

## Configuring Azure CLI authentication in Terraform

No specific configuration is required for the provider to use Azure CLI authentication. A Provider block is _technically_ optional, however we'd strongly recommend defining one to be able to pin the version of the Provider being used:

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"
}
```

If you're looking to use Terraform across Tenants - it's possible to do this by configuring the `tenant_id` field in the Provider block, as shown below:

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"

  tenant_id = "10000000-2000-3000-4000-500000000000"
}
```

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using the Azure CLI to authenticate.
