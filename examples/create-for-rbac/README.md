# Terraform AzureAD / AzureRM RBAC Example

Create a basic Application and Service Principal and then assign a role of Contributor to the default subscription.

This mimics the behaviour of the following commands:

- `az ad sp create-for-rbac --name example`
- `az ad sp create-for-rbac --name example --create-cert`

Note the use of AzureRM resources for subscription data and role assignment.

To view the generated credentials after running `terraform apply`, use `terraform output client_key` and `terraform output client_secret`.

**WARNING:** The Application private key and password will be persisted to state
