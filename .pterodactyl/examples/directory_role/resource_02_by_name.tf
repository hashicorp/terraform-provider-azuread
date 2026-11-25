#> *Activate a directory role by display name*
resource "azuread_directory_role" "example" {
  display_name = "Printer administrator"
}
