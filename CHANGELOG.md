## 0.8.0 (Unreleased)

IMPROVEMENTS:

* `azuread_user` - support for the `onpremises_sam_account_name` and `onpremises_user_principal_name` properties [GH-222]
* `azuread_user` - support for the `immutable_id` property [GH-207]

## 0.7.0 (November 15, 2019)

IMPROVEMENTS:

* provider: migrate to standalone plugin SDK v1.1.0 ([#154](https://github.com/terraform-providers/terraform-provider-azuread/issues/154))
* provider: using the current (rather than the vendored) version of Terraform Core in user agents ([#154](https://github.com/terraform-providers/terraform-provider-azuread/issues/154))
* Data Source `azuread_user` - support looking up a user with `mail_nickname` ([#161](https://github.com/terraform-providers/terraform-provider-azuread/issues/161))
* Data Source `azuread_users` - support looking up users with `mail_nicknames` ([#161](https://github.com/terraform-providers/terraform-provider-azuread/issues/161))
* `azuread_application` - adds ability to build homepage with HTTP in addition to HTTPS ([#155](https://github.com/terraform-providers/terraform-provider-azuread/issues/155))
* `azuread_application` - allow the `app_role` block `value` property to be nil ([#157](https://github.com/terraform-providers/terraform-provider-azuread/issues/157))
* `azuread_user` - support for the `usage_location` property ([#141](https://github.com/terraform-providers/terraform-provider-azuread/issues/141))

## 0.6.0 (August 21, 2019)

IMPROVEMENTS:

* dependencies: upgrading `github.com/Azure/azure-sdk-for-go` to `v32.5.0` ([#140](https://github.com/terraform-providers/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/Azure/go-autorest` to `v13.0.0` ([#140](https://github.com/terraform-providers/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/hashicorp/go-azure-helpers` to `v0.7.0` ([#140](https://github.com/terraform-providers/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/hashicorp/terraform` to `0.12.6` ([#133](https://github.com/terraform-providers/terraform-provider-azuread/issues/133))
* `azuread_service_principal` - support for the `app_role_assignment_required` property ([#127](https://github.com/terraform-providers/terraform-provider-azuread/issues/127))


## 0.5.1 (July 24, 2019)

BUG FIXES:

* `azuread_application_password` - fix incorrect conflicts with ([#129](https://github.com/terraform-providers/terraform-provider-azuread/issues/129))

## 0.5.0 (July 24, 2019)

FEATURES:

* **New Data Source:** `azuread_users` ([#109](https://github.com/terraform-providers/terraform-provider-azuread/issues/109))
* **New Resource:** `azuread_group_member` ([#100](https://github.com/terraform-providers/terraform-provider-azuread/issues/100))

IMPROVEMENTS:

* Data Source `azuread_application` - support for the `app_roles` property ([#110](https://github.com/terraform-providers/terraform-provider-azuread/issues/110))
* Data Source `azuread_service_principal` - export the `app_roles` property ([#110](https://github.com/terraform-providers/terraform-provider-azuread/issues/110))
* `azuread_application` - support for the `app_roles` property ([#98](https://github.com/terraform-providers/terraform-provider-azuread/issues/98))
* `azuread_application` - the `identifier_uris` property now allows `api`,`urn`, and `ms-appx` URI schemas ([#115](https://github.com/terraform-providers/terraform-provider-azuread/issues/115))
* `azuread_application_password` - deprecation of `application_id` in favour of `application_object_id` ([#107](https://github.com/terraform-providers/terraform-provider-azuread/issues/107))
* `azuread_group` - support for the `members` property ([#100](https://github.com/terraform-providers/terraform-provider-azuread/issues/100))
* `azuread_group` - support for the `owners` property ([#62](https://github.com/terraform-providers/terraform-provider-azuread/issues/62))
* `azuread_service_principal` - export the `oauth2_permissions` property ([#103](https://github.com/terraform-providers/terraform-provider-azuread/issues/103))

BUG FIXES:

* `azuread_application_password` - will now wait for replication on resource creation ([#118](https://github.com/terraform-providers/terraform-provider-azuread/issues/118))
* `azuread_service_principal_password` - will now wait for replication on resource creation ([#117](https://github.com/terraform-providers/terraform-provider-azuread/issues/117))

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
* `azuread_user` - increase the maximum allowed length of `password` to 256 ([#81](https://github.com/terraform-providers/terraform-provider-azuread/issues/81))
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
