---
subcategory: "Authentication"
layout: "azuread"
page_title: "Authenticating via Managed Service Identity"
description: |-
  This guide will cover how to use Managed Service Identity as authentication for the Azure Active Directory Provider.

---

# Azure Active Directory Provider: Authenticating using Managed Service Identity

Terraform supports a number of different methods for authenticating to Azure:

* [Authenticating to Azure using the Azure CLI](azure_cli.html)
* Authenticating to Azure using Managed Service Identity (covered in this guide)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

Once you have configured a Service Principal as described in this guide, you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

## What is Managed Service Identity?

Certain services within Azure (for example Virtual Machines and Virtual Machine Scale Sets) can be assigned an Azure Active Directory identity. This identity can then be granted permissions to manage objects in Azure Active Directory.

Once a resource is configured with an identity, a local metadata service exposes credentials which can be used by applications such as Terraform.

When using a Managed Service Identity, you can only manage resources in the tenant where the corresponding service principal is homed. If you need to manage multiple tenants from the same location, we suggest using Service Principal Authentication with a [client certificate](service_principal_client_certificate.html) or [client secret](service_principal_client_secret.html) so that you can specify different credentials for each tenant.

## Configuring Managed Service Identity

The (simplified) Terraform Configuration below configures a Virtual Machine with Managed Service Identity, and then outputs the Object ID of the corresponding Service Principal:

```hcl
data "azurerm_subscription" "current" {}

resource "azurerm_linux_virtual_machine" "test" {
  name = "test-vm"

  # ...

  identity = {
    type = "SystemAssigned"
  }
}

output "test_msi_object_id" {
  value = azurerm_linux_virtual_machine.test.identity.0.principal_id
}
```

The implicitly created Service Principal should have the same or similar name as your virtual machine.

Refer to the [azurerm_linux_virtual_machine][azurerm_linux_virtual_machine] and [azurerm_windows_virtual_machine][azurerm_windows_virtual_machine] documentation for more information on how to use these resources to launch a new virtual machine.

## Configuring Managed Service Identity in Terraform

At this point we assume that Managed Service Identity is configured on the resource (e.g. Virtual Machine) being used - and that you are running Terraform on that resource.

Terraform can be configured to use Managed Service Identity for authentication in one of two ways: using Environment Variables or by defining the fields within the Provider block.

You can configure Terraform to use Managed Service Identity by setting the Environment Variable `ARM_USE_MSI` to `true`; as shown below:

```shell
$ export ARM_USE_MSI=true
```

-> **Using a Custom MSI Endpoint?** In the unlikely event you're using a custom endpoint for Managed Service Identity - this can be configured using the `ARM_MSI_ENDPOINT` Environment Variable - however this shouldn't need to be configured in regular use.

Whilst a Provider block is _technically_ optional when using Environment Variables - we'd strongly recommend defining one to be able to pin the version of the Provider to be used:

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"
}
```

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using Managed Service Identity.

Next you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.

---

It's also possible to enable Managed Service Identity within the Provider Block:

```hcl
provider "azuread" {
  # Whilst version is optional, we /strongly recommend/ using it to pin the version of the Provider to be used
  version = "=1.1.0"

  use_msi = true
}
```

-> **Using a Custom MSI Endpoint?** In the unlikely event you're using a custom endpoint for Managed Service Identity - this can be configured using the `msi_endpoint` field - however this shouldn't need to be configured in regular use.

More information on [the fields supported in the Provider block can be found here](../index.html#argument-reference).

At this point running either `terraform plan` or `terraform apply` should allow Terraform to run using Managed Service Identity.

Next you should follow the [Configuring a Service Principal for managing Azure Active Directory](service_principal_configuration.html) guide to grant the Service Principal necessary permissions to create and modify Azure Active Directory objects such as users and groups.


[azurerm_linux_virtual_machine]: https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine.html
[azurerm_windows_virtual_machine]: https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine.html
