data "azuread_directory_role_templates" "current" {}

output "roles" {
  value = data.azuread_directory_role_templates.current.object_ids
}
