resource "azuread_custom_directory_role" "example" {
  display_name = "My Custom Role"
  description  = "Allows reading applications and updating groups"
  enabled      = true
  version      = "1.0"

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/applications/basic/update",
      "microsoft.directory/applications/create",
      "microsoft.directory/applications/standard/read",
    ]
  }

  permissions {
    allowed_resource_actions = [
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/allProperties/read",
      "microsoft.directory/groups/basic/update",
      "microsoft.directory/groups/create",
      "microsoft.directory/groups/delete",
    ]
  }
}
