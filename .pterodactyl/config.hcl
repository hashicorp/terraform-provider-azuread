provider {
  name = "azuread"
  type = "sdkv2"
}

registry_format = "docs" # Use "legacy" if the provider uses the old `website/docs` structure

integration {
  sdkv2 {
    package = "github.com/hashicorp/terraform-provider-azuread/internal/provider"
    factory = "AzureADProvider"
  }
}

extraction {
  sdkv2_validators = true
}

# Generation is opt-in by entity category. Enable only the documents owned by this tool.
generation {
  provider_index      = false
  resources           = true
  data_sources        = true
  ephemeral_resources = false
  list_resources      = false
  actions             = false
  functions           = false
}

templates {
  # Use "tfplugindocs-enhanced" to add extracted constraints, or "flat" for separate Arguments/Attributes Reference sections.
  format = "flat"
  # Opt in to Number and singular collection element names, such as Map of String, matching tfplugindocs.
  # tfplugindocs_types = true
  templates_directory = ".pterodactyl/templates"
  examples_directory = ".pterodactyl/examples"
}
