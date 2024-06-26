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

*Look up by display name prefix*
```terraform
data "azuread_groups" "sales" {
  display_name_prefix = "sales-"
}
```

*Look up all groups*
```terraform
data "azuread_groups" "all" {
  return_all = true
}
```

*Look up all mail-enabled groups*
```terraform
data "azuread_groups" "mail_enabled" {
  mail_enabled = true
  return_all   = true
}
```

*Look up all security-enabled groups that are not mail-enabled*
```terraform
data "azuread_groups" "security_only" {
  mail_enabled     = false
  return_all       = true
  security_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `display_names` - (Optional) The display names of the groups.
* `display_name_prefix` - (Optional) A common display name prefix to match when returning groups.
* `ignore_missing` - (Optional) Ignore missing groups and return groups that were found. The data source will still fail if no groups are found. Cannot be specified with `return_all`. Defaults to `false`.
* `mail_enabled` - (Optional) Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
* `object_ids` - (Optional) The object IDs of the groups.
* `return_all` - (Optional) A flag to denote if all groups should be fetched and returned. Cannot be specified wth `ignore_missing`. Defaults to `false`.
* `security_enabled` - (Optional) Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.

~> One of `display_names`, `display_name_prefix`, `object_ids` or `return_all` should be specified. Either `display_name` or `object_ids` _may_ be specified as an empty list, in which case no results will be returned.

## Attributes Reference

The following attributes are exported:

* `display_names` - The display names of the groups.
* `object_ids` - The object IDs of the groups.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
