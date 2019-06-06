## 0.4.0 (June 06, 2019)

NOTES:

* Resource creation potentially could take longer after this release as the provider will now attempt to wait for replication like the az cli tool. 

FEATURES:

* **New Resource:** `azuread_application_password` ([#71](https://github.com/terraform-providers/terraform-provider-azuread/issues/71))

IMPROVEMENTS:

* dependencies: upgrading to `v0.12.0` of `github.com/hashicorp/terraform` ([#82](https://github.com/terraform-providers/terraform-provider-azuread/issues/82))
* Data Source `azuread_application` - now exports the `group_membership_claims` property ([#78](https://github.com/terraform-providers/terraform-provider-azuread/issues/78))
* Data Source `azuread_application` - now exports the `oauth2_permissions` property ([#79](https://github.com/terraform-providers/terraform-provider-azuread/issues/79))
* `azuread_application` - support for the `group_membership_claims` property ([#78](https://github.com/terraform-providers/terraform-provider-azuread/issues/78))
* `azuread_application` - now exports the `oauth2_permissions` property ([#79](https://github.com/terraform-providers/terraform-provider-azuread/issues/79))
* `azuread_application` - now exports the `object_id` property ([#99](https://github.com/terraform-providers/terraform-provider-azuread/issues/99))
* `azuread_application` - support for the `type` property enabling the creation of `native` applications ([#74](https://github.com/terraform-providers/terraform-provider-azuread/issues/74))
* `azuread_application` - will now wait for replication by waiting for 10 successful reads after creation ([#93](https://github.com/terraform-providers/terraform-provider-azuread/issues/93))
* `azuread_group` - will now wait for replication by waiting for 10 successful reads after creation ([#91](https://github.com/terraform-providers/terraform-provider-azuread/issues/91))
* `azuread_group` - now exports the `object_id` property ([#99](https://github.com/terraform-providers/terraform-provider-azuread/issues/99))
* `azuread_service_principal` - will now wait for replication by waiting for 10 successful reads after creation ([#93](https://github.com/terraform-providers/terraform-provider-azuread/issues/93))
* `azuread_service_principal` - now exports the `object_id` property ([#99](https://github.com/terraform-providers/terraform-provider-azuread/issues/99))
* `azuread_user` - will now wait for replication by waiting for 10 successful reads after creation ([#91](https://github.com/terraform-providers/terraform-provider-azuread/issues/91))
* `azuread_user` - increase the maximum allowed lengh of `password` to 256 ([#81](https://github.com/terraform-providers/terraform-provider-azuread/issues/81))
* `azuread_user` - now exports the `object_id` property ([#99](https://github.com/terraform-providers/terraform-provider-azuread/issues/99))

## 0.3.1 (April 18, 2019)

BUG FIXES:

* Release fixing metadata to register the provider as compatible with Terraform 0.12.

## 0.3.0 (April 18, 2019)

NOTES:

* This release includes a Terraform SDK upgrade with compatibility for Terraform v0.12. The provider remains backwards compatible with Terraform v0.11 and there should not be any significant behavioural changes. ([#56](https://github.com/terraform-providers/terraform-provider-azuread/issues/56))

BUG FIXES:

* `azuread_application` - the order of the `reply_urls` property no longer matters ([#61](https://github.com/terraform-providers/terraform-provider-azuread/issues/61))

## 0.2.0 (March 12, 2019)

FEATURES:

* **New Data Source:** `azuread_domains` ([#27](https://github.com/terraform-providers/terraform-provider-azuread/issues/27))
* **New Data Source:** `azuread_group` ([#14](https://github.com/terraform-providers/terraform-provider-azuread/issues/14))
* **New Resource:** `azuread_group` ([#14](https://github.com/terraform-providers/terraform-provider-azuread/issues/14))

IMPROVEMENTS:

* dependencies: switching to use Go Modules ([#26](https://github.com/terraform-providers/terraform-provider-azuread/issues/26))
* dependencies: updating `github.com/Azure/azure-sdk-for-go` to v24.1.0 ([#25](https://github.com/terraform-providers/terraform-provider-azuread/issues/25))
* dependencies: updating `github.com/Azure/go-autorest` to v11.2.8 ([#24](https://github.com/terraform-providers/terraform-provider-azuread/issues/24))
* validation: adding validation to all fields ([#30](https://github.com/terraform-providers/terraform-provider-azuread/issues/30))
* `azuread_application` - support for `required_resource_access` property ([#23](https://github.com/terraform-providers/terraform-provider-azuread/issues/23))
* `azuread_service_principal` - support for the `tags` property ([#31](https://github.com/terraform-providers/terraform-provider-azuread/issues/31))
* `azuread_service_principal_password` - support for realitive ends dates with the `end_date_relative` property ([#53](https://github.com/terraform-providers/terraform-provider-azuread/issues/53))

BUG FIXES:

* `azuread_application` - correctly reading back the `reply_urls` property into state ([#21](https://github.com/terraform-providers/terraform-provider-azuread/issues/21))


## 0.1.0 (January 09, 2019)

Initial release of the Azure Active Directory provider - featuring resources split out from the AzureRM Provider.

FEATURES:

* New Data Source: `azuread_application`
* New Data Source: `azuread_service_principal`
* New Resource: `azuread_application`
* New Resource: `azuread_service_principal`
* New Resource: `azuread_service_principal_password`
