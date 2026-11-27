---
page_title: "Data Source: azuread_groups"
subcategory: "Groups"
description: |-
  Gets Object IDs or Display Names for multiple Azure Active Directory groups.
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

## Arguments Reference

The following arguments are supported:

* `display_name_prefix` - (Optional) A common display name prefix to match when returning groups. Value must not be empty. Exactly one of `display_names`, `display_name_prefix`, `object_ids`, or `return_all` must be specified.
* `display_names` - (Optional) The display names of the groups. For each element: Value must not be empty. Exactly one of `display_names`, `display_name_prefix`, `object_ids`, or `return_all` must be specified.
* `ignore_missing` - (Optional) Ignore missing groups and return groups that were found. The data source will still fail if no groups are found. Defaults to `false`. Conflicts with `return_all`.
* `mail_enabled` - (Optional) Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Conflicts with `object_ids`.
* `object_ids` - (Optional) The object IDs of the groups. Exactly one of `display_names`, `display_name_prefix`, `object_ids`, or `return_all` must be specified.
* `return_all` - (Optional) A flag to denote if all groups should be fetched and returned. Conflicts with `ignore_missing`. Exactly one of `display_names`, `display_name_prefix`, `object_ids`, or `return_all` must be specified.
* `security_enabled` - (Optional) Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Conflicts with `object_ids`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `groups` - A list of groups. (see [nested schema](#schema-6e6573746564-6f7574707574-67726f757073))

---

<a id="schema-6e6573746564-6f7574707574-67726f757073"></a>
The `groups` object exports the following:

* `display_name` - The display name of the group.
* `mail_enabled` - Whether the group is mail-enabled.
* `object_id` - The object ID of the group.
* `security_enabled` - Whether the group is security-enabled.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
