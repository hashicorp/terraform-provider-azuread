---
subcategory: "Directory Roles"
---

# Data Source: azuread_directory_roles

Gets Object IDs or Display Names for multiple Azure Active Directory Roles.

~> **NOTE:** This data source is supported only if `EnableMsGraphBeta` feature is enabled.

-> **NOTE:** This data source can retrieve only information about activated Directory Role.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `RoleManagement.ReadWrite.Directory` within the `Azure Active Directory` API. 
Please refer to [this documentation](https://docs.microsoft.com/en-us/graph/api/resources/directoryrole?view=graph-rest-1.0) for more details about directory roles. 
List of available built-in directory roles can be found in [this documentation](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference).


## Example Usage

```terraform
data "azuread_directory_roles" "dir_roles" {
  display_names = ["Global Administrator", "Application Administrator"]
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The Display Names of the Azure AD Directory Roles.
* `object_ids` - (Optional) The Object IDs of the Azure AD Directory Roles.

~> **NOTE:** Either `display_names` or `object_ids` should be specified. Either of these _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The Display Names of the Azure AD Directory Roles.
* `object_ids` - The Object IDs of the Azure AD Directory Roles.
