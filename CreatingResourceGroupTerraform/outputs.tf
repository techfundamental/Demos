output "id" {
  description = "Resource Group id"
  value = azurerm_resource_group.this.id
}
output "name" {
  description = "Resource Group name"
  value = azurerm_resource_group.this.name
}
output "resource" {
  description = "Resource Group object"
  value = azurerm_resource_group.this
}
output "location" {
  value       = azurerm_resource_group.this.location
  description = "The resource group location "
}
