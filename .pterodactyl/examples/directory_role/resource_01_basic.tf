#> *Activate a directory role by its template ID*
resource "azuread_directory_role" "example" {
  template_id = "00000000-0000-0000-0000-000000000000"
}
