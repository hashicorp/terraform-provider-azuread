---
subcategory: "Applications"
---

# Resource: azuread_application_flexible_federated_identity_credential

Manages a flexible federated identity credential associated with an application within Azure Active Directory.

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

resource "azuread_application_flexible_federated_identity_credential" "example" {
  application_id             = azuread_application_registration.example.id
  claims_matching_expression = "claims['sub'] matches 'repo:contoso/contoso-repo:ref:refs/heads/*' and claims['job_workflow_ref'] matches 'contoso/contoso-prod/.github/workflows/*.yml@refs/heads/main'"
  display_name               = "my-repo-deploy"
  description                = "Deployments for my-repo"
  audience                  = "api://AzureADTokenExchange"
  issuer                     = "https://token.actions.githubusercontent.com"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The resource ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
* `audience` - (Required) The audience that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
* `claims_matching_expression` - (Required) The expression to match for claims. See the [Preview Documentation](https://learn.microsoft.com/en-us/entra/workload-id/workload-identities-flexible-federated-identity-credentials?tabs=terraformcloud#flexible-federated-identity-credential-expression-language-functionality) for more information.
* `description` - (Optional) A description for the federated identity credential.
* `display_name` - (Required) A unique display name for the federated identity credential. Changing this forces a new resource to be created.
* `issuer` - (Required) The URL of the external identity provider, which must match the issuer claim of the external token being exchanged.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `credential_id` - A UUID used to uniquely identify this federated identity credential.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 15 minutes) Used when creating the resource.
* `read` - (Defaults to 5 minutes) Used when retrieving the resource.
* `update` - (Defaults to 5 minutes) Used when updating the resource.
* `delete` - (Defaults to 5 minutes) Used when deleting the resource.

## Import

Flexible Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the flexible federated identity credential, e.g.

```shell
terraform import azuread_application_flexible_federated_identity_credential.example 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
```

-> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.
