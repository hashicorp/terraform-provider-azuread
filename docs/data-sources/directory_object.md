---
subcategory: "Base"
---

# Data Source: azuread_directory_object

Retrieves the OData type for a generic directory object having the provided object ID.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up and output type of object by ID*
```terraform
data "azuread_directory_object" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}

output "object_type" {
  value = data.azuread_directory_object.example.type
}
```

## Argument Reference

The following arguments are supported:

* `object_id` - (Optional) Specifies the Object ID of the directory object to look up.

## Attributes Reference 

The following attributes are exported:

*`object_id` - The object ID of the directory object.
*`type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
