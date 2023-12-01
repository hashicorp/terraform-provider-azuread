---
page_title: "Authenticating via the Azure CLI"
subcategory: "Authentication"
---

# Authenticating using the Azure CLI

Terraform supports a number of different methods for authenticating to Azure:

* Authenticating to Azure using the Azure CLI (covered in this guide)
* [Authenticating to Azure using Managed Identity](managed_service_identity.html)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)
* [Authenticating to Azure using a Service Principal and OpenID Connect](service_principal_oidc.html)

---

We recommend using either a Service Principal or Managed Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

## Important Notes about Authenticating using the Azure CLI

* Terraform only supports authenticating using the `az` CLI (and this must be available on your PATH) - authenticating using the older `azure` CLI or PowerShell Az / AzureRM Cmdlets is not supported.
* Prior to version 2.35, authenticating via the Azure CLI was only supported when using a User Account. For example `az login --service-principal` was not supported and you had to use either a [Client Secret](service_principal_client_secret.html) or a [Client Certificate](service_principal_client_certificate.html). From 2.35 upwards, authenticating via the Azure CLI is supported when using a Service Principal or Managed Identity.

---

## Logging into the Azure CLI

-> **Using other clouds** If you're using the **China**, **German** or **Government** Azure Clouds - you'll need to first configure the Azure CLI to work with that Cloud, so that the correct authentication service is used.  You can do this by running: <br><br>`$ az cloud set --name AzureChinaCloud|AzureGermanCloud|AzureUSGovernment`

---

Firstly, login to the Azure CLI using a User, Service Principal or Managed Identity.

User Account:

```shell
az login --allow-no-subscriptions
```

Service Principal with a Secret:

```shell
az login --service-principal -u "CLIENT_ID" -p "CLIENT_SECRET" --tenant "TENANT_ID" --allow-no-subscriptions
```

Service Principal with a Certificate:

```shell
az login --service-principal -u "CLIENT_ID" -p "CERTIFICATE_PEM" --tenant "TENANT_ID" --allow-no-subscriptions
```

Service Principal with Open ID Connect (for use in CI / CD):

```shell
az login --service-principal -u "CLIENT_ID" --tenant "TENANT_ID" --allow-no-subscriptions
```

Managed Identity:

```shell
az login --identity --allow-no-subscriptions

or

az login --identity --username "CLIENT_ID" --allow-no-subscriptions
```

The `--allow-no-subscriptions` argument enables access to tenants that have no linked subscriptions, in addition to tenants that do.

---

Once logged in - it's possible to list the Subscriptions and Tenants associated with the account via:

```shell-session
$ az account list -o table --all --query "[].{TenantID: tenantId, Subscription: name, Default: isDefault}"
```

The output (similar to below) will display one or more Tenants and/or Subscriptions.

```
TenantID                              Subscription                         Default
------------------------------------  -----------------------------------  ---------
00000000-0000-1111-1111-111111111111  N/A(tenant level account)            False
00000000-0000-2222-2222-222222222222  N/A(tenant level account)            False
00000000-0000-1111-1111-111111111111  My Subscription                      True
00000000-0000-1111-1111-111111111111  My Other Subscription                False
```

Each entry shown is referred to as an `Azure CLI account`, which represents either a subscription with its linked tenant, or a tenant without any accessible subscriptions (Azure CLI does not show tenant names or domains). The provider will select the tenant ID from your default Azure CLI account. If you have more than one tenant listed in the output of `az account list`, for example if you are a guest user in other tenants, you can specify the tenant to use.

```shell-session
# sh
export ARM_TENANT_ID=00000000-0000-2222-2222-222222222222
```
```powershell
# PowerShell
$env:ARM_TENANT_ID = 00000000-0000-2222-2222-222222222222
```

You can also configure the tenant ID from within the provider block.

```hcl
provider "azuread" {
  tenant_id = "00000000-0000-2222-2222-222222222222"
}
```

Alternatively, you can configure the Azure CLI to default to the tenant you are managing with Terraform.

```shell-session
$ az login --allow-no-subscriptions --tenant "TENANT_ID_OR_DOMAIN"
```

<br>

-> **Tenants and Subscriptions** The AzureAD provider operates on tenants and not on subscriptions. We recommend always specifying `az login --allow-no-subscriptions` as it will force the Azure CLI to report tenants with no associated subscriptions, or where your user account does not have any roles assigned for a subscription.

---

## Configuring Azure CLI authentication in Terraform

No specific configuration is required for the provider to use Azure CLI authentication. If you're looking to use Terraform across Tenants - it's possible to do this by configuring the `tenant_id` field in the Provider block, as shown below:

```hcl
provider "azuread" {
  tenant_id = "00000000-0000-1111-1111-111111111111"
}
```

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using the Azure CLI to authenticate.

## Disabling Azure CLI authentication

For compatibility reasons and to ensure a positive user experience when running Terraform interactively, Azure CLI authentication is enabled by default. It's possible to disable authentication using Azure CLI, which you may wish to do in automated environments such as CI/CD pipelines or when scripting operations with Terraform.

To do so, add the `use_cli` configuration property in the Provider block.

```hcl
provider "azuread" {
  use_cli = false
}
```

Alternatively, you can set the `ARM_USE_CLI` environment variable.

```shell
# sh
export ARM_USE_CLI=false

# PowerShell
$env:ARM_USE_CLI = false
```
