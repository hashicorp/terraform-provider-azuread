---
page_title: "Authenticating via a Service Principal and OpenID Connect"
subcategory: "Authentication"
---

# Authenticating using a Service Principal and OpenID Connect

* [Authenticating to Azure using the Azure CLI](azure_cli.html)
* [Authenticating to Azure using Managed Identity](managed_service_identity.html)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)
* Authenticating to Azure using a Service Principal and OpenID Connect (covered in this guide)

---

We recommend using either a Service Principal or Managed Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

Once you have configured a Service Principal as described in this guide, you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

---

## Setting up an Application and Service Principal

A Service Principal is a security principal within Azure Active Directory which can be granted permissions to manage objects in Azure Active Directory. To authenticate with a Service Principal, you will need to create an Application object within Azure Active Directory, which you will use as a means of authentication, either using a [Client Secret](service_principal_client_secret.html), [Client Certificate](service_principal_client_certificate.html), or OpenID Connect (which is documented in this guide). This can be done using the Azure Portal.

This guide will cover how to create an Application and linked Service Principal, and then how to assign federated identity credentials to the Application so that it can be used for authentication via OpenID Connect.

## Creating the Application and Service Principal

We're going to create the Application in the Azure Portal - to do this, navigate to the [**Azure Active Directory** overview][azure-portal-aad-overview] within the [Azure Portal][azure-portal] - then select the [**App Registrations** blade][azure-portal-applications-blade]. Click the **New registration** button at the top to add a new Application within Azure Active Directory. On this page, set the following values then press **Create**:

- **Name** - this is a friendly identifier and can be anything (e.g. "Terraform")
- **Supported Account Types** - this should be set to "Accounts in this organisational directory only (single tenant)"
- **Redirect URI** - you should choose "Web" in for the URI type. the actual value can be left blank

At this point the newly created Azure Active Directory application should be visible on-screen - if it's not, navigate to the [**App Registration** blade][azure-portal-applications-blade] and select the Azure Active Directory application.

At the top of this page, you'll need to take note of the "Application (client) ID" and the "Directory (tenant) ID", which you can use for the values of `client_id` and `tenant_id` respectively.

## Configure Azure Active Directory Application to Trust a GitHub Repository

An application will need a federated credential specified for each GitHub Environment, Branch Name, Pull Request, or Tag based on your use case. For this example, we'll give permission to `main` branch workflow runs.

-> **Tip:** You can also configure the Application using the [azuread_application](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application) and [azuread_application_federated_identity_credential](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_federated_identity_credential) resources in the AzureAD Terraform Provider.


### Via the Portal

On the Azure Active Directory application page, go to **Certificates and secrets**.

In the Federated credentials tab, select Add credential. The Add a credential blade opens. In the **Federated credential scenario** drop-down box select **GitHub actions deploying Azure resources**.

Specify the **Organization** and **Repository** for your GitHub Actions workflow. For **Entity type**, select **Environment**, **Branch**, **Pull request**, or **Tag** and specify the value. The values must exactly match the configuration in the GitHub workflow. For our example, let's select **Branch** and specify `main`.

Add a **Name** for the federated credential.

The **Issuer**, **Audiences**, and **Subject identifier** fields autopopulate based on the values you entered.

Click **Add** to configure the federated credential.

### Via the Azure API

```sh
az rest --method POST \
        --uri https://graph.microsoft.com/beta/applications/${APP_OBJ_ID}/federatedIdentityCredentials \
        --headers Content-Type='application/json' \
        --body @body.json
```

Where the body is:

```json
{
  "name":"${REPO_NAME}-pull-request",
  "issuer":"https://token.actions.githubusercontent.com",
  "subject":"repo:${REPO_OWNER}/${REPO_NAME}:refs:refs/heads/main",
  "description":"${REPO_OWNER} PR",
  "audiences":["api://AzureADTokenExchange"],
}
```

See the [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation-create-trust-github) for more details.

### Configure Azure Active Directory Application / Managed Identity to Trust an Azure DevOps Service Connection

An application or managed identity requires a federated credential for each Azure DevOps service connection. In common scenarios, there will be one registration/identity per environment with one credential for the environment's service connection.

#### Automatic configuration - App registration
The simplest method for setting up federation is to create a new **Workload Identity federation (automatic)** in Azure DevOps. This will automatically create a new app registration in your tenant. Alternatively, if you want to retain your existing connection, you can convert an existing secret-based connection to a federated one using the provided `Convert` option in the service connection overview. This may have some implications for your pipelines, but there is a rollback option available.

For more details, refer to [the official documentation](https://learn.microsoft.com/en-gb/azure/devops/pipelines/library/connect-to-azure?view=azure-devops#create-a-new-workload-identity-federation-service-connection) for more details.

#### Manual Configuration - Managed Identity / App Registration
To configure a Managed Identity for federation, select the **Workload Identity federation (manual)** option in the creation wizard. After providing a name for the new connection, you will be presented with the issuer URL and subject identifier values required to configure federated credentials in the Managed Identity resource.

In Azure Managed Identity resource settings, select **Other** from the **Federated credential scenario** options and provide the issuer URL, subject identifier provided by Azure DevOps, and a display name for your credentials. Then, proceed with the **Verify and save** option in the Azure DevOps `New Azure service connection` wizard.

#### Terraform Configuration - Managed Identity / App registration

Please refer to examples in [azuredevops_serviceendpoint_azurerm](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs/resources/serviceendpoint_azurerm#workload-identity-federation-manual-azurerm-service-endpoint-subscription-scoped) resource documentation.

### Configure Azure Active Directory Application to Trust a Generic Issuer

On the Azure Active Directory application page, go to **Certificates and secrets**.

In the Federated credentials tab, select **Add credential**. The 'Add a credential' blade opens. Refer to the instructions from your OIDC provider for completing the form, before choosing a **Name** for the federated credential and clicking the **Add** button.

## Configuring Terraform to use OIDC

As we've obtained the credentials for this Service Principal - it's possible to configure them in a few different ways.

When storing the credentials as Environment Variables, for example:

```bash
$ export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
$ export ARM_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
$ export ARM_TENANT_ID="00000000-0000-0000-0000-000000000000"
$ export ARM_USE_OIDC=true
```

### OIDC token
The provider will use the `ARM_OIDC_TOKEN` environment variable as an OIDC token. You can use this variable to specify the token provided by your OIDC provider.

**GitHub Actions**

When running Terraform in GitHub Actions, the provider will detect the `ACTIONS_ID_TOKEN_REQUEST_URL` and `ACTIONS_ID_TOKEN_REQUEST_TOKEN` environment variables set by the GitHub Actions runtime. You can also specify the `ARM_OIDC_REQUEST_TOKEN` and `ARM_OIDC_REQUEST_URL` environment variables.

For GitHub Actions workflows, you'll need to ensure the workflow has `write` permissions for the `id-token`.

```yaml
permissions:
  id-token: write
  contents: read
```

For more information about OIDC in GitHub Actions, see [official documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-cloud-providers).


**Azure DevOps Pipelines**

When running Terraform in Azure DevOps Pipelines, the provider use `ARM_OIDC_REQUEST_TOKEN` and `ARM_OIDC_REQUEST_URL` environment variables, if these are absent, it will attempt to fall back on `SYSTEM_ACCESSTOKEN` and `SYSTEM_OIDCREQUESTURI` environment variables respectively.

The ADO service connection ID is required in combination with these and can be specified via environment variable `ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID`, or in the provider configuration directly via `ado_pipeline_service_connection_id`. For users of the AzAPI provider AzureAD will also fall back on `ARM_OIDC_AZURE_SERVICE_CONNECTION_ID` for compatibility.

~> **Note:** If the `ado_pipeline_service_connection_id` value is not set, either directly via config or by environment variable, the presence of the remaining OIDC configuration will result in the provider falling back on the GitHub OIDC authoriser instead.

For Azure DevOps Pipelines, at least one task in the pipeline has Service Connection support and has your service connection specified. Without this, the agent will fail to load the Service Connection and results in a `No service connection found with identifier "..."`  error.

It is recommend to use the `AzureCLI@2` task as below (note the `azureSubscription` input parameter):

```yaml
- task: AzureCLI@2
  inputs:
    azureSubscription: $(SERVICE_CONNECTION_ID)
    scriptType: bash
    scriptLocation: "inlineScript"
    inlineScript: |
      # Terraform commands
  env:
    #...
    ARM_USE_OIDC: true
    SYSTEM_ACCESSTOKEN: $(System.AccessToken)
    SYSTEM_OIDCREQUESTURI: $(System.OidcRequestUri)
    ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID: $(SERVICE_CONNECTION_ID)
```

---

-> **Note:** Support for OpenID Connect was added in version 2.23.0 of the Terraform AzureAD provider.

~> **Note:** If using the AzureRM Backend you may also need to configure OIDC there too, see [the documentation for the AzureRM Backend](https://www.terraform.io/language/settings/backends/azurerm) for more information.

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

---

The following Terraform and Provider blocks can be specified - where `2.23.0` is the version of the Azure Provider that you'd like to use:

```hcl
# We strongly recommend using the required_providers block to set the
# Azure Provider source and version being used
terraform {
  required_providers {
    azuread = {
      source  = "hashicorp/azuread"
      version = "=2.23.0"
    }
  }
}
# Configure the Microsoft Azure Provider
provider "azuread" {
  use_oidc = true # or use the environment variable "ARM_USE_OIDC=true"
  features {}
}
```

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using the Service Principal to authenticate.

Next you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

[azure-portal]: https://portal.azure.com/
[azure-portal-aad-overview]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview
[azure-portal-applications-blade]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/RegisteredApps/RegisteredApps/Overview
