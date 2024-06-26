---
subcategory: "Applications"
---

# Resource: azuread_application_optional_claims

Manages optional claims for an application registration.

This resource is analogous to the `optional_claims` block in the `azuread_application` resource. When using these resources together, you should use the `ignore_changes` [lifecycle meta-argument](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) (see example below).

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

## Argument Reference

The following arguments are supported:

* `access_token` - (Optional) One or more `access_token` blocks as documented below.
* `application_id` - (Required) The resource ID of the application registration. Changing this forces a new resource to be created.
* `id_token` - (Optional) One or more `id_token` blocks as documented below.
* `saml2_token` - (Optional) One or more `saml2_token` blocks as documented below.

-> At least one of `access_token`, `id_token` or `saml2_token` must be specified

---

`access_token`, `id_token` and `saml2_token` blocks support the following:

* `additional_properties` - List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
* `essential` - Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
* `name` - The name of the optional claim.
* `source` - The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.

## Attributes Reference

No additional attributes are exported.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 10 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 10 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Application Optional Claims can be imported using the object ID of the application, in the following format.

```shell
terraform import azuread_application_optional_claims.example /applications/00000000-0000-0000-0000-000000000000
```
