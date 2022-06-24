locals {
  # this local value is used by `resource "azuread_app_role_assignment"`
  # resources to assign AAD group to the Application.
  msiam_access_index = length(azuread_service_principal.example.app_roles) - 1
}
