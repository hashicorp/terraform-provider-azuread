---
page_title: "Configuring a Service Principal to manage an Azure Active Directory"
subcategory: "Authentication"
---

# Azure Active Directory Provider: Configuring a Service Principal for managing Azure Active Directory

Terraform supports a number of different methods for authenticating to Azure:

* [Authenticating to Azure using the Azure CLI](azure_cli.html)
* [Authenticating to Azure using Managed Service Identity](managed_service_identity.html)
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

---

We recommend using either a Service Principal or Managed Service Identity when running Terraform non-interactively (such as when running Terraform in a CI server) - and authenticating using the Azure CLI when running Terraform locally.

## Creating a Service Principal

A Service Principal represents an application within Azure Active Directory whose authentication tokens can be used as the `client_id`, `client_secret`, and `tenant_id` fields needed by Terraform.

Depending on how the service principal authenticates to azure it can be created in a number of different ways: 
* [Authenticating to Azure using a Service Principal and a Client Certificate](service_principal_client_certificate.html)
* [Authenticating to Azure using a Service Principal and a Client Secret](service_principal_client_secret.html)

## Azure Active Directory permissions

Now that you have created and authenticated an Application / Service Principal pair, you will need to grant some permissions to administer Azure Active Directory. You can choose either of the following methods to achieve similar results.

### Method 1: Directory Roles (recommended)

With this method, you will assign directory roles to the Service Principal you created, to grant the desired permissions to administer objects in your Azure Active Directory tenant.

Navigate to the [**Azure Active Directory** overview][azure-portal-aad-overview] within the [Azure Portal][azure-portal]. Go to the [**Roles and Administrators** blade][azure-portal-aad-roles-blade].

Locate the role you wish to assign and click on it. Consult the [documentation for administrator role permissions][admin-roles-docs] from Microsoft for more information about the available roles and the permissions they grant.

Click "Add assignments" and type the name of your Service Principal in the search box to locate it. If you know the Object ID of the Service Principal, verify that it is the same. Select it and click the "Add" button to assign the role.

The choice of which directory roles to assign will be specific to your organisations Commonly used roles include:

Role                        | Description
--------------------------- | -----------------------------------------------------------------------------------------------------------------------
`Global Administrator`      | Effective superuser permissions to administer any object in your AAD tenant. Sometimes called `Company Administrator`.
`Global Reader`             | Commonly used in conjunction with other roles to allow reading, but not writing, of directory data.
`Application Administrator` | Create and manage applications, service principals (enterprise applications) and application proxy.
`Groups Administrator`      | Create and manage groups.
`User Administrator`        | Create and manage users _and_ groups.

### Method 2: API access with admin consent

This method involves granting API scopes to your Application, and then granting consent for your Application to access the APIs in its own capacity (i.e. not on behalf of a user).

Navigate to the [**Azure Active Directory** overview][azure-portal-aad-overview] within the [Azure Portal][azure-portal] and select the [**App Registrations** blade][azure-portal-aad-applications-blade]. Locate your registered Application and click on its display name to manage it.

Go to the API Permissions blade for the Application and click the "Add a permission" button. In the pane that opens, select "Azure Active Directory Graph" (under the Supported Legacy APIs subheading). Do not select "Microsoft Graph", as the provider does not currently make use of this API.

Choose "Application Permissions" for the permission type, and check the permissions you would like to assign. The permissions you need will depend on which directory objects you are seeking to manage with Terraform. We suggest the following permissions:

- `Application.ReadWrite.All`
- `Directory.ReadWrite.All`

Note that these permissions do not cover all use cases. Notable actions you cannot perform with permissions granted in this way include deleting groups and deleting users.

Once you have assigned the permissions, you will need to grant admin consent. This requires that you are signed in to the Portal as a Global Administrator. Click the "Grant admin consent" button and confirm the action.

The Application now has permission to administer your Azure Active Directory tenant.


[admin-roles-docs]: https://docs.microsoft.com/en-us/azure/active-directory/users-groups-roles/directory-assign-admin-roles
[azure-portal]: https://portal.azure.com/
[azure-portal-aad-overview]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview
[azure-portal-aad-roles-blade]: https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/RolesAndAdministrators
[azure-portal-aad-applications-blade]: https://portal.azure.com#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/RegisteredApps
