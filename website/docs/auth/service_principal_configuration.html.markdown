---
layout: "azuread"
page_title: "Azure Active Directory Provider: Configuring a Service Principal to manage an Azure Active Directory"
sidebar_current: "docs-azuread-authentication-configuring-service-principal"
description: |-
  This guide will cover how to use grant permissions to a Service Principal (Shared Account) to manage objects within an Azure Active Directory .

---

# Azure Active Directory Provider: Configuring a Service Principal for managing Azure Active Directory

Terraform supports a number of different methods for authenticating to Azure:

* [Authenticating to Azure using the Azure CLI](azure_cli.html)
* [Authenticating to Azure using Managed Service Identity](managed_service_identity.html)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* Authenticating to Azure using a Service Principal and a Client Secret (which is covered in this guide)

Further steps must be taken to grant a Service Principal permission to manage objects in an Azure Active Directory:

* Granting a Service Principal permission to manage AAD (which is covered in this guide) 

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

## Creating a Service Principal

A Service Principal is an application within Azure Active Directory whose authentication tokens can be used as the `client_id`, `client_secret`, and `tenant_id` fields needed by Terraform (`subscription_id` can be independently recovered from your Azure account details).

Depending on how the service principal authenticates to azure it can be created in a number of different ways: 
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

## Granting administrator permissions

~> **Note**: This requires the use of powershell cmdlets and is easiest to run in CloudShell.  


Firstly, connect to the directory using:

```shell
Connect-AzureAD -TenantID "00000000-0000-0000-0000-000000000000"
```

Next we want to get the correct role to assign, in this case `User Account Administrator`:

```shell
$role = Get-AzureADDirectoryRole | Where-Object {$_.displayName -eq 'User Account Administrator'}
Write-Host $role
```

Since this is a built-in Role, if this doesn't exist (returns `null` above) then we need to instantiate it from the Role Template:

```shell
if ($role -eq $null) {
    # Instantiate an instance of the role template
    $roleTemplate = Get-AzureADDirectoryRoleTemplate | Where-Object {$_.displayName -eq 'User Account Administrator'}
    Enable-AzureADDirectoryRole -RoleTemplateId $roleTemplate.ObjectId

    # Fetch User Account Administrator role instance again
    $role = Get-AzureADDirectoryRole | Where-Object {$_.displayName -eq 'User Account Administrator'}
}
```

Next we need the Client ID (sometimes referred to as the Application ID) of the Service Principal. We can look this up by it's display name:

```shell
$sp = Get-AzureADServicePrincipal | Where-Object {$_.displayName -eq 'Service Principal Name'}
$sp.ObjectId
```

Now that we have all the required information we can add the service principal to the role:

```shell
Add-AzureADDirectoryRoleMember -ObjectId $role.ObjectId -RefObjectId $sp.ObjectId

```

Finally we can repeat this for the `Company Administrator` role:

```shell
$role = Get-AzureADDirectoryRole | Where-Object {$_.displayName -eq 'Company Administrator'}
$role

if ($role -eq $null) {
    # Instantiate an instance of the role template
    $roleTemplate = Get-AzureADDirectoryRoleTemplate | Where-Object {$_.displayName -eq 'Company Administrator'}
    Enable-AzureADDirectoryRole -RoleTemplateId $roleTemplate.ObjectId

    # Fetch User Account Administrator role instance again
    $role = Get-AzureADDirectoryRole | Where-Object {$_.displayName -eq 'Company Administrator'}
}

$sp = Get-AzureADServicePrincipal | Where-Object {$_.displayName -eq 'Service Pricipal Name'}
$sp.ObjectId

Add-AzureADDirectoryRoleMember -ObjectId $role.ObjectId -RefObjectId $sp.ObjectId

```

At this point you should now be able to manage Users, Groups and other Azure Active Directory resources using Terraform.
