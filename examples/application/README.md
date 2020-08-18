# Terraform AzureAD / AzureRM Application Example

Create some Application registrations in Azure Active Directory for use in a Widgets application that utilises OAuth2 on the Microsoft Identity Platform.

1. `widgets-service` represents a Widgets API service
2. `widgets-app` represents a Widgets web application

To view the generated passwords after running `terraform apply`, use:

```sh
terraform output widgets_service_client_secret
```

**WARNING:** The passwords will be persisted to state
