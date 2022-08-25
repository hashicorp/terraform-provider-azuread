---
subcategory: "Base"
---

# Data Source: azuread_directory_object

Retrieves the OData type for a generic directory object having the provided object ID.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires either `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up and output type of object by ID*
```terraform
data "azuread_directory_object" "example" {
  object_id = "object-id"
}

output "object_type" {
  value = data.azuread_directory_object.example.type
}
```

## Attributes Reference 

The following attributes are exported:

*`object_id` - The object_id specified.
*`type` - The type of the object
