#> *Look up and output type of object by ID*
data "azuread_directory_object" "example" {
  object_id = "00000000-0000-0000-0000-000000000000"
}

output "object_type" {
  value = data.azuread_directory_object.example.type
}
