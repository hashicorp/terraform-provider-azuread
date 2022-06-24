terraform {
  required_version = "~> 1.2"
  required_providers {
    # https://registry.terraform.io/providers/hashicorp/azuread/latest
    azuread = {
      version = "=2.22.0"
    }
    # https://registry.terraform.io/providers/hashicorp/tls/3.4.0
    tls = {
      source  = "hashicorp/tls"
      version = "3.4.0"
    }
  }
}

provider "azuread" {}

provider "tls" {}
