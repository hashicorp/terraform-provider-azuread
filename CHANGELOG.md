## 0.2.0 (Unreleased)

FEATURES:

* **New Data Source:** `azuread_domains` [GH-27]
* **New Data Source:** `azuread_group` [GH-14]
* **New Resource:** `azuread_group` [GH-14]

IMPROVEMENTS:

* dependencies: switching to use Go Modules [GH-26]
* dependencies: updating `github.com/Azure/azure-sdk-for-go` to v24.1.0 [GH-25]
* dependencies: updating `github.com/Azure/go-autorest` to v11.2.8 [GH-24]
* validation: adding validation to all fields [GH-30]
* `azuread_application` - support for `required_resource_access` property [GH-23]
* `azuread_service_principal` - support for the `tags` property [GH-31]

BUG FIXES:

* `azuread_application` - correctly reading back the `reply_urls` property into state [GH-21]


## 0.1.0 (January 09, 2019)

Initial release of the Azure Active Directory provider - featuring resources split out from the AzureRM Provider.

FEATURES:

* New Data Source: `azuread_application`
* New Data Source: `azuread_service_principal`
* New Resource: `azuread_application`
* New Resource: `azuread_service_principal`
* New Resource: `azuread_service_principal_password`
