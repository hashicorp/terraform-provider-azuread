# Terraform AzureAD assigning Azure AD group to Enterprise Application example

Create a basic Application and Service Principal then assign an AzureAD Group to it so group members can access the Application.

1. `widgets-service` represents a Widgets API service
2. `widgets-app` represents a Widgets web application

To view the generated passwords after running `terraform apply`, use:

```sh
terraform output widgets_service_client_secret
```

**WARNING:** The passwords will be persisted to state
