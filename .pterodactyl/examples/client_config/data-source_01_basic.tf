data "azuread_client_config" "current" {}

output "object_id" {
  value = data.azuread_client_config.current.object_id
}
