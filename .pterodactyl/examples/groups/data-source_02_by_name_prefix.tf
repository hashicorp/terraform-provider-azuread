#> *Look up by display name prefix*
data "azuread_groups" "sales" {
  display_name_prefix = "sales-"
}
