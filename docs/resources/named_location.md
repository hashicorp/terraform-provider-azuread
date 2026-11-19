---
page_title: "Resource: azuread_named_location"
subcategory: "Conditional Access"
description: |-
  Manages a Named Location within Azure Active Directory. -> API Limits This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to reduce parallelism https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options or specify extended custom resource timeouts https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts.
---

# Resource: azuread_named_location

Manages a Named Location within Azure Active Directory.  -> **API Limits** This resource is subject to a restrictive API request limit of 1 request/second. Whilst Terraform will automatically back-off and retry throttled requests, if you have a large number of resource changes to make, you may wish to [reduce parallelism](https://developer.hashicorp.com/terraform/cli/commands/apply#apply-options) or specify extended [custom resource timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_named_location" "example-ip" {
  display_name = "IP Named Location"
  ip {
    ip_ranges = [
      "1.1.1.1/32",
      "2.2.2.2/32",
    ]
    trusted = true
  }
}

resource "azuread_named_location" "example-country" {
  display_name = "Country Named Location"
  country {
    countries_and_regions = [
      "GB",
      "US",
    ]
    include_unknown_countries_and_regions = false
  }
}
```

## Arguments Reference

The following arguments are supported:

* `display_name` - (Required) No description provided.

---

* `country` - (Optional) The `country` block as defined below. Exactly one of `ip` or `country` must be specified. Changing this value forces a new resource to be created.
   (see [nested schema](#schema-6e6573746564-696e707574-636f756e747279))
* `ip` - (Optional) The `ip` block as defined below. Exactly one of `ip` or `country` must be specified. Changing this value forces a new resource to be created.
   (see [nested schema](#schema-6e6573746564-696e707574-6970))

---

<a id="schema-6e6573746564-696e707574-636f756e747279"></a>
The `country` block supports the following:

* `countries_and_regions` - (Required) No description provided.

* `country_lookup_method` - (Optional) No description provided. Possible values include `authenticatorAppGps` and `clientIpAddress`. Defaults to `clientIpAddress`.
* `include_unknown_countries_and_regions` - (Optional) No description provided.
---

<a id="schema-6e6573746564-696e707574-6970"></a>
The `ip` block supports the following:

* `ip_ranges` - (Required) No description provided.

* `trusted` - (Optional) No description provided.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `object_id` - The object ID of the named location.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 5 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_named_location.my_location /identity/conditionalAccess/namedLocations/00000000-0000-0000-0000-000000000000
```
