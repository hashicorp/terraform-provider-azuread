#> *Look up by group name*
data "azuread_groups" "example" {
  display_names = ["group-a", "group-b"]
}
