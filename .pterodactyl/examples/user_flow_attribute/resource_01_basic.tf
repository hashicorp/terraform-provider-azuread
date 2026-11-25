#> *Basic example*
resource "azuread_user_flow_attribute" "example" {
  display_name = "Hobby"
  description  = "Your hobby"
  data_type    = "string"
}
