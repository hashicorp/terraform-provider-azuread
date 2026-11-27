---
page_title: "Data Source: azuread_named_location"
subcategory: "Conditional Access"
description: |-
  Gets information about a Named Location within Azure Active Directory.
---

# Data Source: azuread_named_location

Gets information about a Named Location within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this data source.

When authenticated with a service principal, this resource requires the following application roles: `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Reader`

## Example Usage

```terraform
data "azuread_named_location" "example" {
  display_name = "My Named Location"
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) Specifies the display named of the named location to look up. Value must not be empty.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Data Source.
* `country` - A list of `country` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-636f756e747279))
* `ip` - A list of `ip` objects as defined below. (see [nested schema](#schema-6e6573746564-6f7574707574-6970))
* `object_id` - The object ID of the named location.

---

<a id="schema-6e6573746564-6f7574707574-636f756e747279"></a>
The `country` object exports the following:

* `countries_and_regions` - List of countries and/or regions in two-letter format specified by ISO 3166-2.
* `country_lookup_method` - Method of detecting country the user is located in.
* `include_unknown_countries_and_regions` - Whether IP addresses that don't map to a country or region are included in the named location.

---

<a id="schema-6e6573746564-6f7574707574-6970"></a>
The `ip` object exports the following:

* `ip_ranges` - List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
* `trusted` - Whether the named location is trusted.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the data source.
