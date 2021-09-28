---
subcategory: "Groups"
---

# Data Source: azuread_groups

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`

When authenticated with a user principal, this data source does not require any additional roles.

## Example Usage

*Look up by group name*
```terraform
data "azuread_groups" "example" {
  display_names = ["group-a", "group-b"]
}
```

*Look up all groups*
```terraform
data "azuread_groups" "example" {
  return_all = true
}
```

*Look up all mail-enabled groups*
```terraform
data "azuread_groups" "example" {
  mail_enabled = true
  return_all   = true
}
```

*Look up all security-enabled groups that are not mail-enabled*
```terraform
data "azuread_groups" "example" {
  mail_enabled     = false
  return_all       = true
  security_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The display names of the groups.
* `mail_enabled` - (Optional) Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
* `object_ids` - (Optional) The object IDs of the groups.
* `return_all` - (Optional) A flag to denote if all groups should be fetched and returned.
* `security_enabled` - (Optional) Whether the returned groups should be security-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.

~> One of `display_names`, `object_ids` or `return_all` should be specified. Either of the first two _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The display names of the groups.
* `object_ids` - The object IDs of the groups.
