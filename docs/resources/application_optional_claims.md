---
page_title: "Resource: azuread_application_optional_claims"
subcategory: "Applications"
description: |-
  Manages optional claims for an application registration. This resource is analogous to the optional_claims block in the azuread_application resource. When using these resources together, you should use the ignore_changes lifecycle meta-argument https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle (see example below).
---

# Resource: azuread_application_optional_claims

Manages optional claims for an application registration.  This resource is analogous to the `optional_claims` block in the `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_optional_claims" "example" {
  application_id = azuread_application_registration.example.id

  access_token {
    name = "myclaim"
  }

  access_token {
    name = "otherclaim"
  }

  id_token {
    name                  = "userclaim"
    source                = "user"
    essential             = true
    additional_properties = ["emit_as_roles"]
  }

  saml2_token {
    name = "samlexample"
  }
}
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application to which these optional claims belong. Changing this value forces a new resource to be created.

---

* `access_token` - (Optional) The `access_token` block as defined below. At least one of `access_token`, `id_token`, or `saml2_token` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-6163636573735f746f6b656e))
* `id_token` - (Optional) The `id_token` block as defined below. At least one of `access_token`, `id_token`, or `saml2_token` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-69645f746f6b656e))
* `saml2_token` - (Optional) The `saml2_token` block as defined below. At least one of `access_token`, `id_token`, or `saml2_token` must be specified.
   (see [nested schema](#schema-6e6573746564-696e707574-73616d6c325f746f6b656e))

---

<a id="schema-6e6573746564-696e707574-6163636573735f746f6b656e"></a>
The `access_token` block supports the following:

* `name` - (Required) The name of the optional claim.

* `additional_properties` - (Optional) List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - (Optional) Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `source` - (Optional) The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.
---

<a id="schema-6e6573746564-696e707574-69645f746f6b656e"></a>
The `id_token` block supports the following:

* `name` - (Required) The name of the optional claim.

* `additional_properties` - (Optional) List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - (Optional) Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `source` - (Optional) The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.
---

<a id="schema-6e6573746564-696e707574-73616d6c325f746f6b656e"></a>
The `saml2_token` block supports the following:

* `name` - (Required) The name of the optional claim.

* `additional_properties` - (Optional) List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. For each element: Possible values include `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - (Optional) Whether the claim specified by the client is necessary to ensure a smooth authorization experience. Defaults to `false`.
* `source` - (Optional) The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object. The only possible value is `user`.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_application_optional_claims.example /applications/00000000-0000-0000-0000-000000000000
```
