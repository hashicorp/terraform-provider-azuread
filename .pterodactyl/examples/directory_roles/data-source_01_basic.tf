data "azuread_directory_roles" "current" {}

output "roles" {
  value = data.azuread_directory_roles.current.object_ids
}
