output "vnet_name" {
  value = "${azurerm_virtual_network.main.name}"
}

output "resource_group_name" {
  value = "${azurerm_resource_group.main.name}"
}