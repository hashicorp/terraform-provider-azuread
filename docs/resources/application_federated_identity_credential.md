---
page_title: "Resource: azuread_application_federated_identity_credential"
subcategory: "Applications"
description: |-
  Manages a federated identity credential associated with an application within Azure Active Directory.
---

# Resource: azuread_application_federated_identity_credential

Manages a federated identity credential associated with an application within Azure Active Directory.

## API Permissions

The following API permissions are required in order to use this resource.

When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

-> When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Terraform must be an owner of the application.

When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

## Example Usage

```terraform
resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_application_federated_identity_credential" "example" {
  application_id = azuread_application_registration.example.id
  display_name   = "my-repo-deploy"
  description    = "Deployments for my-repo"
  audiences      = ["api://AzureADTokenExchange"]
  issuer         = "https://token.actions.githubusercontent.com"
  subject        = "repo:my-organization/my-repo:environment:prod"
}
```

## Arguments Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application for which this federated identity credential should be created. Changing this value forces a new resource to be created.
* `audiences` - (Required) List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
* `display_name` - (Required) A unique display name for the federated identity credential. String length must be between `1` and `120` characters inclusive. Changing this value forces a new resource to be created.
* `issuer` - (Required) The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
* `subject` - (Required) The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.

---

* `description` - (Optional) A description for the federated identity credential.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Resource.
* `credential_id` - A UUID used to uniquely identify this federated identity credential.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 15 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

This resource can be imported using its resource ID, e.g.

```shell
terraform import azuread_application_federated_identity_credential.example 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
```
