---
subcategory: "Conditional Access"
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

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Specifies the display named of the named location to look up.

## Attributes Reference 

The following attributes are exported:

* `country` - A `country` block as documented below, which describes a country-based named location.
* `id` - The ID of the named location.
* `ip` - An `ip` block as documented below, which describes an IP-based named location.
* 
---

`country` block exports the following:

* `countries_and_regions` - List of countries and/or regions in two-letter format specified by ISO 3166-2.
* `include_unknown_countries_and_regions` - Whether IP addresses that don't map to a country or region are included in the named location.

---

`ip` block exports the following:

* `ip_ranges` - List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
* `trusted` - Whether the named location is trusted.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
