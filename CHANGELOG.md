## 1.2.1 (January 14, 2021)

BUGFIXES:

* `data.azuread_application` - correctly set the `display_name` attribute in state.
* `azuread_application` - correctly set the `display_name` attribute in state.

## 1.2.0 (January 14, 2021)

NOTES:

* **Terraform Plugin SDK Upgrade:** This version upgrades the Terraform Plugin SDK to v2.3.0. This does not provide any additional provider features or resources but is useful for developers and part of our development roadmap.
* **Refactor into multiple packages:** As part of our preparation for Microsoft Graph support, this release refactors resources and data sources into separate Go packages.

IMPROVEMENTS:

* `azuread_application` - support new values `include_externally_authenticated_upn`, `include_externally_authenticated_upn_without_hash`, and `use_guid` for the `additional_properties` property of the `optional_claims` block.

DEPRECATIONS:

* `data.azuread_application` - the `name` property has been renamed to `display_name` and will be removed in version 2.0.
* `data.azuread_group` - the `name` property has been renamed to `display_name` and will be removed in version 2.0.
* `data.azuread_groups` - the `names` property has been renamed to `display_names` and will be removed in version 2.0.
* `azuread_application` - the `name` property has been renamed to `display_name` and will be removed in version 2.0.
* `azuread_application` - the `type` property is now deprecated and will be removed in version 2.0, as there is no longer any distinction between native and webapp/api applications.
* `azuread_group` - the `name` property has been renamed to `display_name` and will be removed in version 2.0.

## 1.1.1 (November 26, 2020)

BUG FIXES:

* `azuread_application` - resolves an issue where setting `prevent_duplicate_names = true` causes an error for new applications ([#367](https://github.com/terraform-providers/terraform-provider-azuread/issues/367))
* `azuread_application` - fixes a bug where the default owner for a new application is removed ([#366](https://github.com/terraform-providers/terraform-provider-azuread/issues/366))

## 1.1.0 (November 25, 2020)

FEATURES:

* Added a flag to allow users to customize the Partner ID or opt-out of the default Terraform Partner ID ([#350](https://github.com/terraform-providers/terraform-provider-azuread/issues/350))
* This release includes updated support for working directly with tenants using Azure CLI authentication. We recommend the use of `az login --allow-no-subscription` to populate tenant-level accounts (which have no subscriptions).

IMPROVEMENTS:

* `data.azuread_user` - support the `given_name`, `surname`, `job_title`, `department`, `company_name`, `physical_delivery_office_name`, `street_address`, `city`, `state`, `country`, `postal_code` and `mobile` attribute ([#351](https://github.com/terraform-providers/terraform-provider-azuread/issues/351))
* `azuread_user` - support the `given_name`, `surname`, `job_title`, `department`, `company_name`, `physical_delivery_office_name`, `street_address`, `city`, `state`, `country`, `postal_code` and `mobile` properties ([#351](https://github.com/terraform-providers/terraform-provider-azuread/issues/351))

BUG FIXES:

* **Provider:** Fixed an issue where CLI authentication produced a `parsing json result` error during provider initialization ([#358](https://github.com/terraform-providers/terraform-provider-azuread/issues/358))
* `azuread_application` - enable removal of owners on existing applications, and creation of applications with no owners ([#355](https://github.com/terraform-providers/terraform-provider-azuread/issues/355))
* `azuread_application` - fixed a bug where specifying the `prevent_duplicate_names` property would report a false positive on update. ([#338](https://github.com/terraform-providers/terraform-provider-azuread/issues/338))

## 1.0.0 (September 03, 2020)

NOTES:

* **Major Version:** This is a major version upgrade which contains some breaking changes as detailed below.
* **Terraform 0.10/0.11:** This version of the provider requires Terraform 0.12.x or later and will not work with earlier versions.

FEATURES:

* New resource: `azuread_application_app_role` ([#150](https://github.com/terraform-providers/terraform-provider-azuread/issues/150)] [[#306](https://github.com/terraform-providers/terraform-provider-azuread/issues/306))
* New resource: `azuread_application_oauth2_permission` ([#267](https://github.com/terraform-providers/terraform-provider-azuread/issues/267))

BREAKING CHANGES:

* `azuread_application` - a default value for the `homepage` property is no longer derived when unspecified ([#268](https://github.com/terraform-providers/terraform-provider-azuread/issues/268))
* `azuread_application_password` - the deprecated `application_id` property has been removed
* `data.azuread_group` - the `name` property is now case-insensitive ([#246](https://github.com/terraform-providers/terraform-provider-azuread/issues/246))
* `data.azuread_groups` and `data.azuread_users` will not error if no results found

## 0.11.0 (July 09, 2020)

IMPROVEMENTS:

* Provider: no longer require configuring `subscription_id` (configuration value) / `ARM_SUBSCRIPTION_ID` (environment variable). ([#271](https://github.com/terraform-providers/terraform-provider-azuread/issues/271))
* `data.azuread_client_config` - deprecate the `subscription_id` property. For compatibility, still populates `subscription_id` if the provider is configured with a subscription ID ([#271](https://github.com/terraform-providers/terraform-provider-azuread/issues/271))
* `data.azuread_application` - support for the `application_id` property ([#274](https://github.com/terraform-providers/terraform-provider-azuread/issues/274))
* `data.azuread_users` - support the `ignore_missing` property ([#256](https://github.com/terraform-providers/terraform-provider-azuread/issues/256))
* `data.azuread_users` - export the `users` attribute containing a list of users with additional properties ([#256](https://github.com/terraform-providers/terraform-provider-azuread/issues/256))
* `azuread_application` - support the `prevent_duplicate_names` property ([#279](https://github.com/terraform-providers/terraform-provider-azuread/issues/279))
* `azuread_application` - validate `app_roles` and `oauth2_permissions` to check for duplicate `value`s ([#287](https://github.com/terraform-providers/terraform-provider-azuread/issues/287))
* `azuread_group` - support the `prevent_duplicate_names` property ([#279](https://github.com/terraform-providers/terraform-provider-azuread/issues/279))

BUG FIXES:

* `azuread_group` - remediate AAD replication delays when adding/removing group members ([#283](https://github.com/terraform-providers/terraform-provider-azuread/issues/283))
* `azuread_group` - remediate AAD replication delays after group creation, before setting owners/members ([#290](https://github.com/terraform-providers/terraform-provider-azuread/issues/290))

## 0.10.0 (June 05, 2020)

BREAKING CHANGES:

* `azuread_application` - the `oauth2_permissions` attribute has changed from a list to a set. If you are referencing this attribute with explicit list indexes, you will need to update your configuration to use a `for` expression. For example:

    ```hcl
    id = azuread_application.example.oauth2_permissions[0].id
    ```

    becomes

    ```hcl
    id = [for permission in azuread_application.example.oauth2_permissions : permission.id][0]
    ```

FEATURES:

* **New Resource:** `azuread_application_certificate` ([#262](https://github.com/terraform-providers/terraform-provider-azuread/issues/262))
* **New Resource:** `azuread_service_principal_certificate` ([#262](https://github.com/terraform-providers/terraform-provider-azuread/issues/262))

IMPROVEMENTS:

* `azuread_application` - support for the `optional_claims` property, for access tokens and ID tokens ([#260](https://github.com/terraform-providers/terraform-provider-azuread/issues/260))
* `azuread_application` - support for the `oauth2_permissions` property ([#252](https://github.com/terraform-providers/terraform-provider-azuread/issues/252))
* `azuread_application_password` - support the `description` property ([#253](https://github.com/terraform-providers/terraform-provider-azuread/issues/253))
* `azuread_service_principal_password` - support the `description` property ([#253](https://github.com/terraform-providers/terraform-provider-azuread/issues/253))
* `data.azuread_users` - support empty lists for `user_principal_names`/`object_ids`/`mail_nicknames` properties ([#258](https://github.com/terraform-providers/terraform-provider-azuread/issues/258))
* `data.azuread_groups` - support empty lists for `names`/`object_ids` properties ([#257](https://github.com/terraform-providers/terraform-provider-azuread/issues/257))

BUG FIXES:

* `azuread_application_password` and `azuread_service_principal_password` - Plan-time validation for `end_date` / `end_date_relative` ([#261](https://github.com/terraform-providers/terraform-provider-azuread/issues/261))
* `azuread_application_password` and `azuread_service_principal_password` - Change the resource ID format to mitigate potential UUID collision ([#264](https://github.com/terraform-providers/terraform-provider-azuread/issues/264))

## 0.9.0 (May 15, 2020)

DEPENDENCIES:

* upgrade `azure-sdk-for-go` to `v42.1.0` ([#247](https://github.com/terraform-providers/terraform-provider-azuread/issues/247))

IMPROVEMENTS:

* `azuread_application` - the `group_membership_claims` property now supports `ApplicationGroup` ([#238](https://github.com/terraform-providers/terraform-provider-azuread/issues/238))
* `azuread_service_principal` - changing the `tags` property no longer forces a new resource ([#245](https://github.com/terraform-providers/terraform-provider-azuread/issues/245))

BUG FIXES:

* `data.azuread_user` - use `equals` instead of `startsWith` when looking uo users by `mailNickname` ([#251](https://github.com/terraform-providers/terraform-provider-azuread/issues/251))
* `data.azuread_users` - use `equals` instead of `startsWith` when looking uo users by `mailNickname` ([#251](https://github.com/terraform-providers/terraform-provider-azuread/issues/251))

## 0.8.0 (March 16, 2020)

FEATURES:

* **New Data Source:** `azuread_client_config` ([#229](https://github.com/terraform-providers/terraform-provider-azuread/issues/229))

IMPROVEMENTS:

* dependencies: upgrade `azure-sdk-for-go` to `v40.3.0` ([#225](https://github.com/terraform-providers/terraform-provider-azuread/issues/225))
* dependencies: upgrade `go-autorest/autorest` to `v0.10.0` ([#225](https://github.com/terraform-providers/terraform-provider-azuread/issues/225))
* dependencies: upgrade `terraform-plugin-sdk` to `v1.6.0` ([#225](https://github.com/terraform-providers/terraform-provider-azuread/issues/225))
* `azuread_application` - support for the `logout_url` property ([#226](https://github.com/terraform-providers/terraform-provider-azuread/issues/226))
* `azuread_group` - support for the `description` property ([#216](https://github.com/terraform-providers/terraform-provider-azuread/issues/216))
* `azuread_user` - support for the `onpremises_sam_account_name` and `onpremises_user_principal_name` properties ([#222](https://github.com/terraform-providers/terraform-provider-azuread/issues/222))
* `azuread_user` - support for the `immutable_id` property ([#207](https://github.com/terraform-providers/terraform-provider-azuread/issues/207))

BUG FIXES:

* `azuread_application` - ensure all owners are added before removed ([#226](https://github.com/terraform-providers/terraform-provider-azuread/issues/226))
* `azuread_application_password` - validate the `length` property is less then `863` ([#228](https://github.com/terraform-providers/terraform-provider-azuread/issues/228))
* `azuread_group` - the `owners` property is now additive during creation allowing an existing owner to be provided ([#211](https://github.com/terraform-providers/terraform-provider-azuread/issues/211))
* `azuread_group_member` - mark as missing when member cannot be found instead of erroring ([#227](https://github.com/terraform-providers/terraform-provider-azuread/issues/227))
* `azuread_service_principal_password` - validate the `length` property is less then `863` ([#228](https://github.com/terraform-providers/terraform-provider-azuread/issues/228))

## 0.7.0 (November 15, 2019)

IMPROVEMENTS:

* provider: migrate to standalone plugin SDK v1.1.0 ([#154](https://github.com/terraform-providers/terraform-provider-azuread/issues/154))
* provider: using the current (rather than the vendored) version of Terraform Core in user agents ([#154](https://github.com/terraform-providers/terraform-provider-azuread/issues/154))
* `azuread_application` - adds ability to build homepage with HTTP in addition to HTTPS ([#155](https://github.com/terraform-providers/terraform-provider-azuread/issues/155))
* `azuread_application` - allow the `app_role` block `value` property to be nil ([#157](https://github.com/terraform-providers/terraform-provider-azuread/issues/157))
* `azuread_user` - support for the `usage_location` property ([#141](https://github.com/terraform-providers/terraform-provider-azuread/issues/141))
* `data.azuread_user` - support looking up a user with `mail_nickname` ([#161](https://github.com/terraform-providers/terraform-provider-azuread/issues/161))
* `data.azuread_users` - support looking up users with `mail_nicknames` ([#161](https://github.com/terraform-providers/terraform-provider-azuread/issues/161))

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

* `azuread_application` - support for the `app_roles` property ([#98](https://github.com/terraform-providers/terraform-provider-azuread/issues/98))
* `azuread_application` - the `identifier_uris` property now allows `api`,`urn`, and `ms-appx` URI schemas ([#115](https://github.com/terraform-providers/terraform-provider-azuread/issues/115))
* `azuread_application_password` - deprecation of `application_id` in favour of `application_object_id` ([#107](https://github.com/terraform-providers/terraform-provider-azuread/issues/107))
* `azuread_group` - support for the `members` property ([#100](https://github.com/terraform-providers/terraform-provider-azuread/issues/100))
* `azuread_group` - support for the `owners` property ([#62](https://github.com/terraform-providers/terraform-provider-azuread/issues/62))
* `azuread_service_principal` - export the `oauth2_permissions` property ([#103](https://github.com/terraform-providers/terraform-provider-azuread/issues/103))
* `data.azuread_application` - support for the `app_roles` property ([#110](https://github.com/terraform-providers/terraform-provider-azuread/issues/110))
* `data.azuread_service_principal` - export the `app_roles` property ([#110](https://github.com/terraform-providers/terraform-provider-azuread/issues/110))

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
* `data.azuread_application` - now exports the `group_membership_claims` property ([#78](https://github.com/terraform-providers/terraform-provider-azuread/issues/78))
* `data.azuread_application` - now exports the `oauth2_permissions` property ([#79](https://github.com/terraform-providers/terraform-provider-azuread/issues/79))

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
