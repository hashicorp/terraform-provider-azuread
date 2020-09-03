# Terraform AzureAD / AzureRM RBAC Example

Create a basic Application and Service Principal and then assign a role of Contributor to the default subscription.

This mimics the behaviour of `az ad sp create-for-rbac --name example --years 2`.

Note the use of AzureRM resources for subscription data and role assignment.

To view the generated password after running `terraform apply`, use `terraform output client_secret`.

**WARNING:** The Application password will be persisted to state
