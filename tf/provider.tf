provider "azurerm" {
  version = "~>1.37"
}

terraform {
  backend "azurerm" {
    resource_group_name  = "rg-oss-paris"
    storage_account_name = "ossparisterraform"
    container_name       = "tfstate"
    key                  = "demo.terraform.tfstate"
  }
}

