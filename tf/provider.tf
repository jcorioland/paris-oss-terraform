provider "azurerm" {
  version = "~>1.37"

  subscription_id = "3ac6a0c9-5f38-4f50-b98d-8f77cae6b872"
  tenant_id       = "cc996d09-21a0-477a-87c8-b38d17f6982a"
}

terraform {
  backend "azurerm" {
    resource_group_name  = "rg-oss-paris"
    storage_account_name = "ossparisterraform"
    container_name       = "tfstate"
    key                  = "demo.terraform.tfstate"
  }
}