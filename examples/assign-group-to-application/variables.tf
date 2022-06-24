variable "app-users" {
  description = "List of UPNs which should be members of the group. eg name@mytenant.onmicrosoft.com"
  type        = list(string)
}
