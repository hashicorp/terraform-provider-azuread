---
subcategory: "Directory Object"
---

# Data Source: azuread_directory_object

Returns the OData type of the specified Azure Active Directory Object using the ObjectId.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up by ObjectId*
```terraform
data "azuread_directory_object" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
```

## Argument Reference

The following arguments are supported:

* `object_id` - The ObjectId of the Azure Active Directory Object.

## Attributes Reference

The following attributes are exported:

* `object_id` - The object ID of the Azure Active Directory Object.
* `type` - The OData Type of the object.