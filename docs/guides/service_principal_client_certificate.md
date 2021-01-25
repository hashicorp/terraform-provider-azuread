---
page_title: "Authenticating via a Service Principal and a Client Certificate"
subcategory: "Authentication"
---

# Azure Active Directory Provider: Authenticating using a Service Principal with a Client Certificate

Terraform supports a number of different methods for authenticating to Azure:

* [Authenticating to Azure using the Azure CLI](azure_cli.html)
* [Authenticating to Azure using Managed Service Identity](managed_service_identity.html)
* Authenticating to Azure using a Service Principal and a Client Certificate (covered in this guide)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

Once you have configured a Service Principal as described in this guide, you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

---

## Setting up an Application and Service Principal

A Service Principal is a security principal within Azure Active Directory which can be granted permissions to manage objects in Azure Active Directory. To authenticate with a Service Principal, you will need to create an Application object within Azure Active Directory, which you will use as a means of authentication, either [using a Client Secret](service_principal_client_secret.html) or a Client Certificate (which is documented in this guide). This can be done using the Azure Portal.

This guide will cover how to generate a client certificate, how to create an Application and linked Service Principal, and then how to assign the Client Certificate to the Application so that it can be used for authentication.

---

## Generating a Client Certificate

Firstly we need to create a certificate which can be used for authentication. To do that we're going to generate a Certificate Signing Request (also known as a CSR) using `openssl` (this can also be achieved using PowerShell, however that's outside the scope of this document):

```shell
$ openssl req -newkey rsa:4096 -nodes -keyout "service-principal.key" -out "service-principal.csr"
```

-> During the generation of the certificate you'll be prompted for various bits of information required for the certificate signing request - at least one item has to be specified for this to complete.

We can now sign that Certificate Signing Request, in this example we're going to self-sign this certificate using the Key we just generated; however it's also possible to do this using a Certificate Authority. In order to do that we're again going to use `openssl`:

```shell
$ openssl x509 -signkey "service-principal.key" -in "service-principal.csr" -req -days 365 -out "service-principal.crt"
```

Finally we can generate a PFX file which can be used to authenticate with Azure:

```shell
$ openssl pkcs12 -export -out "service-principal.pfx" -inkey "service-principal.key" -in "service-principal.crt"
```

Now that we've generated a certificate, we can create the Azure Active Directory Application.

---

### Creating the Application and Service Principal

We're going to create the Application in the Azure Portal - to do this navigate to the [**Azure Active Directory** overview][azure-portal-aad-overview] within the [Azure Portal][azure-portal] - then select the [**App Registrations** blade][azure-portal-applications-blade]. Click the **New registration** button at the top to add a new Application within Azure Active Directory. On this page, set the following values then press **Create**:

- **Name** - this is a friendly identifier and can be anything (e.g. "Terraform")
- **Supported Account Types** - this should be set to "Accounts in this organisational directory only (single tenant)"
- **Redirect URI** - you should choose "Web" in for the URI type. the actual value can be left blank

At this point the newly created Azure Active Directory application should be visible on-screen - if it's not, navigate to the [**App Registration** blade][azure-portal-applications-blade] and select the newly created Azure Active Directory application.

At the top of this page, you'll need to take note of the "Application (client) ID" and the "Directory (tenant) ID", which you can use for the values of `client_id` and `tenant_id` respectively.

### Assigning the Client Certificate to the Azure Active Directory Application

To associate the public portion of the Client Certificate (the `*.crt` file) with the Azure Active Directory Application - to do this select **Certificates & secrets**. This screen displays the Certificates and Client Secrets (i.e. passwords) which are associated with this Azure Active Directory Application.

The Public Key associated with the generated Certificate can be uploaded by selecting **Upload Certificate**, selecting the file which should be uploaded (in the example above, this'd be `service-principal.crt`) - and then hitting **Add**.

---

### Configuring the Service Principal in Terraform

As we've obtained the credentials for this Service Principal - it's possible to configure them in a few different ways.

When storing the credentials as Environment Variables, for example:

```bash
$ export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
$ export ARM_CLIENT_CERTIFICATE_PATH="/path/to/my/client/certificate.pfx"
$ export ARM_CLIENT_CERTIFICATE_PASSWORD="Pa55w0rd123"
$ export ARM_TENANT_ID="10000000-2000-3000-4000-500000000000"
```

The following Provider block can be specified - where `1.1.0` is the version of the AzureAD Provider that you'd like to use:

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"
}
```

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using the Service Principal to authenticate.

Next you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

---

It's also possible to configure these variables either in-line or from using variables in Terraform (as the `client_certificate_path` and `client_certificate_password` are in this example), like so:

~> We recommend not defining these variables in-line since they could easily be checked into Source Control.

```hcl
variable "client_certificate_path" {}
variable "client_certificate_password" {}

provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"

  client_id                   = "00000000-0000-0000-0000-000000000000"
  client_certificate_path     = var.client_certificate_path
  client_certificate_password = var.client_certificate_password
  tenant_id                   = "10000000-2000-3000-4000-500000000000"
}
```

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using the Service Principal to authenticate.

Next you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

[azure-portal]: https://portal.azure.com/
[azure-portal-aad-overview]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview
[azure-portal-applications-blade]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/RegisteredApps/RegisteredApps/Overview
