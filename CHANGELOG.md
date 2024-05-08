## 2.49.0 (Unreleased)

FEATURES:

* **New Data Source:** `azuread_group_role_management_policy` [GH-1327]
* **New Resource:** `azuread_group_role_management_policy` [GH-1327]
* **New Resource:** `azuread_privileged_access_group_assignment_schedule` [GH-1327]
* **New Resource:** `azuread_privileged_access_group_eligibility_schedule` [GH-1327]

ENHANCEMENTS:

* `data.azuread_group` - support for the `include_transitive_members` property [GH-1300]
* `azuread_application` - relax validation for the `identifier_uris` property to allow more values [GH-1351]
* `azuread_application_identifier_uri` - relax validation for the `identifier_uri` property to allow more values [GH-1351]
* `azuread_user` - relax validation for the `employee_type` property to allow more values [GH-1328]

## 2.48.0 (April 11, 2024)

ENHANCEMENTS:

* dependencies: updating to `v0.20240411.1104331` of `github.com/hashicorp/go-azure-sdk/sdk` ([#1353](https://github.com/hashicorp/terraform-provider-azuread/issues/1353))

BUG FIXES:

* provider: fix an issue where the provider was not correctly configured when using a custom metadata host ([#1353](https://github.com/hashicorp/terraform-provider-azuread/issues/1353))

## 2.47.0 (December 14, 2023)

BUG FIXES:

* `azuread_access_package_assignment_policy` - fix a potential crash when removing the `question` block ([#1273](https://github.com/hashicorp/terraform-provider-azuread/issues/1273))
* `data.azuread_named_location` - fix a potential crash if the named location is not found ([#1274](https://github.com/hashicorp/terraform-provider-azuread/issues/1274))

## 2.46.0 (November 16, 2023)

ENHANCEMENTS:

* `data.azuread_application` - update the resource ID format to match the `azuread_application` resource ([#1255](https://github.com/hashicorp/terraform-provider-azuread/issues/1255))
* `azuread_named_location` - add validation for the `ip_ranges` property in the `ip` block ([#1254](https://github.com/hashicorp/terraform-provider-azuread/issues/1254))

## 2.45.0 (October 27, 2023)

FEATURES:

* **New Resource:** `azuread_application_optional_claims` ([#1223](https://github.com/hashicorp/terraform-provider-azuread/issues/1223))

ENHANCEMENTS:

* `azuread_conditional_access_policy` - improved plan-time validation for the `session_controls` block ([#1229](https://github.com/hashicorp/terraform-provider-azuread/issues/1229))
* `azuread_conditional_access_policy` - support for the `sign_in_frequency_authentication_type` and `sign_in_frequency_interval` properties in the `session_controls` block ([#1229](https://github.com/hashicorp/terraform-provider-azuread/issues/1229))
* `azuread_conditional_access_policy` - support for the `included_guests_or_external_users` and `excluded_guests_or_external_users` blocks in the `users` block ([#1222](https://github.com/hashicorp/terraform-provider-azuread/issues/1222))

BUG FIXES:

* `azuread_conditional_access_policy` - removing the `devices` or `session_controls` blocks will no longer force a new resource to be created ([#1229](https://github.com/hashicorp/terraform-provider-azuread/issues/1229))

## 2.44.1 (October 23, 2023)

BUG FIXES:

* `azuread_application_certificate` - work around an unexpected diff with the `application_object_id` property ([#1221](https://github.com/hashicorp/terraform-provider-azuread/issues/1221))
* `azuread_application_federated_identity_credential` - work around an unexpected diff with the `application_object_id` property ([#1221](https://github.com/hashicorp/terraform-provider-azuread/issues/1221))
* `azuread_application_password` - work around an unexpected diff with the `application_object_id` property ([#1221](https://github.com/hashicorp/terraform-provider-azuread/issues/1221))
* `azuread_application_pre_authorized` - work around an unexpected diff with the `application_object_id` property ([#1221](https://github.com/hashicorp/terraform-provider-azuread/issues/1221))

## 2.44.0 (October 20, 2023)

* Developer Note: the Typed Resource SDK, as also used in the AzureRM provider, is now the preferred way of introducing new resources ([#1188](https://github.com/hashicorp/terraform-provider-azuread/issues/1188))

FEATURES:

* **New Resource:** `azuread_application_api_access` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_app_role` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_fallback_public_client` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_from_template` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_identifier_uri` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_known_clients` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_owner` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_permission_scope` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_redirect_uris` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_application_registration` ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* **New Resource:** `azuread_authentication_strength_policy` ([#1171](https://github.com/hashicorp/terraform-provider-azuread/issues/1171))

ENHANCEMENTS:

* `data.azuread_application` - export the `client_id` attribute, deprecate the `application_id` attribute ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `data.azuread_service_principal` - support for the `client_id` property, deprecate the `application_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `data.azuread_service_principals` - support for the `client_ids` property, deprecate the `application_ids` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `data.azuread_service_principals` - export the `client_id` attribute in the `service_principals` block, deprecate the `application_id` attribute ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_application` - export the `client_id` attribute, deprecate the `application_id` attribute ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_application_federated_identity_credential` - support for the `application_id` property, deprecate the `application_object_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_application_certificate` - support for the `application_id` property, deprecate the `application_object_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_application_password` - support for the `application_id` property, deprecate the `application_object_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_application_pre_authorized` - support for the `application_id` property, deprecate the `application_object_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_service_principal` - support for the `client_id` property, deprecate the `application_id` property ([#1214](https://github.com/hashicorp/terraform-provider-azuread/issues/1214))
* `azuread_conditional_access_policy` - support for the `authentication_strength_policy_id` property in the `grant_controls` block [GH_1171]

BUG FIXES:

* `azuread_group_member` - resolve a bug when refreshing state if the group is missing ([#1198](https://github.com/hashicorp/terraform-provider-azuread/issues/1198))

## 2.43.0 (September 22, 2023)

FEATURES:

* **New Resource:** `azuread_directory_role_eligibility_schedule_request` ([#974](https://github.com/hashicorp/terraform-provider-azuread/issues/974))

## 2.42.0 (September 15, 2023)

IMPROVEMENTS:

* provider: support for the `client_id_file_path` and `client_secret_file_path` provider properties ([#1189](https://github.com/hashicorp/terraform-provider-azuread/issues/1189))
* `data.azuread_group` - support for looking up a group with the `mail_nickname` property ([#1173](https://github.com/hashicorp/terraform-provider-azuread/issues/1173))

BUG FIXES:

* `azuread_conditional_access_policy` - allow specifying `terms_of_use` in place of `built_in_controls` in the `grant_controls` block ([#1168](https://github.com/hashicorp/terraform-provider-azuread/issues/1168))

## 2.41.0 (July 27, 2023)

FEATURES:

* **New Data Source:** `azuread_directory_role_templates` ([#1152](https://github.com/hashicorp/terraform-provider-azuread/issues/1152))
* **New Data Source:** `azuread_named_location` ([#1156](https://github.com/hashicorp/terraform-provider-azuread/issues/1156))

IMPROVEMENTS:

* `azuread_access_package_assignment_policy` - support the `Manager` value for the `review_type` property in the `assignment_review_settings` block ([#1159](https://github.com/hashicorp/terraform-provider-azuread/issues/1159))
* `azuread_conditional_access_policy` - support for the `service_principal_risk_levels` property in the `conditions` block ([#1145](https://github.com/hashicorp/terraform-provider-azuread/issues/1145))
* `azuread_conditional_access_policy` - the `grant_controls` block is now optional ([#1155](https://github.com/hashicorp/terraform-provider-azuread/issues/1155))

BUG FIXES:

* `azuread_access_package_resource_package_association` - support destruction of this resource ([#1124](https://github.com/hashicorp/terraform-provider-azuread/issues/1124))
* `azuread_application` - set the `display_name` property correctly on creation to improve UX in the event of failure ([#1160](https://github.com/hashicorp/terraform-provider-azuread/issues/1160))

## 2.40.0 (July 14, 2023)

IMPROVEMENTS:

* dependencies: updating to `v0.62.0` of `github.com/manicminer/hamilton`
* `data.azuread_user` - supporting looking up a user using the `employee_id` property ([#1040](https://github.com/hashicorp/terraform-provider-azuread/issues/1040))
* `data.azuread_users` - supporting looking up users using the `employee_ids` property ([#1040](https://github.com/hashicorp/terraform-provider-azuread/issues/1040))
* `azuread_conditional_access_policy` - support for the `client_applications` block in the `conditions` block ([#1047](https://github.com/hashicorp/terraform-provider-azuread/issues/1047))
* `azuread_conditional_access_policy` - support for the `disable_resilience_defaults` property in the `session_controls` block ([#1135](https://github.com/hashicorp/terraform-provider-azuread/issues/1135))
* `azuread_group` - the `behaviors` property now supports the `CalendarMemberReadOnly` and `ConnectorsDisabled` values ([#1144](https://github.com/hashicorp/terraform-provider-azuread/issues/1144))

## 2.39.0 (May 12, 2023)

IMPROVEMENTS:

* dependencies: updating to `v0.20230511.1094507` of `github.com/hashicorp/go-azure-sdk` ([#1100](https://github.com/hashicorp/terraform-provider-azuread/issues/1100))

BUG FIXES:

* **provider:** fix a token refresh bug that could cause authentication errors after initial token expiry ([#1100](https://github.com/hashicorp/terraform-provider-azuread/issues/1100))

## 2.38.0 (April 27, 2023)

FEATURES:

* **New Data Source:** `azuread_access_package_catalog_role` ([#1033](https://github.com/hashicorp/terraform-provider-azuread/issues/1033))
* **New Resource:** `azuread_access_package_catalog_role_assignment` ([#1033](https://github.com/hashicorp/terraform-provider-azuread/issues/1033))

BUG FIXES:

* **Provider:** fix an issue where API requests might not be retried correctly ([#1090](https://github.com/hashicorp/terraform-provider-azuread/issues/1090))
* `azuread_service_principal_token_signing_certificate` - fix a crash when importing legacy certificates ([#1082](https://github.com/hashicorp/terraform-provider-azuread/issues/1082))

## 2.37.2 (April 20, 2023)

BUG FIXES:

* `azuread_group` - remove conditional ForceNew for the `onpremises_group_type` property, resolve breaking change in v2.37.1 ([#1076](https://github.com/hashicorp/terraform-provider-azuread/issues/1076))
* `azuread_group` - improve a workaround for reading Microsoft 365-only properties for groups in a non-M365 tenant ([#1076](https://github.com/hashicorp/terraform-provider-azuread/issues/1076))
* `azuread_group` - improve a workaround for detecting unwanted changes to the `description` property ([#1074](https://github.com/hashicorp/terraform-provider-azuread/issues/1074))

## 2.37.1 (April 17, 2023)

NOTES:

* This release contains a breaking change with the `azuread_group` resource, in order to fix a regression. Please see [#1072](https://github.com/hashicorp/terraform-provider-azuread/issues/1072) for workaround information.

BUG FIXES:

* `azuread_group` - fix a regression that caused `onpremises_group_type` to be set when not configured, and unsetting this property now forces replacement of the resource ([#1070](https://github.com/hashicorp/terraform-provider-azuread/issues/1070))

## 2.37.0 (April 13, 2023)

FEATURES:

* **New Data Source:** `azuread_access_package` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Data Source:** `azuread_access_package_catalog` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_access_package` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_access_package_assignment_policy` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_access_package_catalog` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_access_package_resource_catalog_association` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_access_package_resource_package_association` ([#903](https://github.com/hashicorp/terraform-provider-azuread/issues/903))
* **New Resource:** `azuread_administrative_unit_role_member` ([#983](https://github.com/hashicorp/terraform-provider-azuread/issues/983))
* **New Resource:** `azuread_user_flow_attribute` ([#1063](https://github.com/hashicorp/terraform-provider-azuread/issues/1063))

IMPROVEMENTS:

* dependencies: updating to `v0.60.0` of `github.com/manicminer/hamilton` ([#1062](https://github.com/hashicorp/terraform-provider-azuread/issues/1062))
* `data.azuread_application` - support for the `service_management_reference` attribute ([#1046](https://github.com/hashicorp/terraform-provider-azuread/issues/1046))
* `data.azuread_group` - support for the `onpremises_group_type` and `writeback_enabled` attributes ([#964](https://github.com/hashicorp/terraform-provider-azuread/issues/964))
* `data.azuread_user` - support for the `mail` property ([#996](https://github.com/hashicorp/terraform-provider-azuread/issues/996))
* `azuread_application` - support for the `service_management_reference` property ([#1046](https://github.com/hashicorp/terraform-provider-azuread/issues/1046))
* `azuread_group` - support for the `onpremises_group_type` and `writeback_enabled` properties ([#964](https://github.com/hashicorp/terraform-provider-azuread/issues/964))

## 2.36.0 (March 03, 2023)

IMPROVEMENTS:

* **Provider:** requests to Microsoft Graph no longer include the tenant ID as part of the URI path ([#1039](https://github.com/hashicorp/terraform-provider-azuread/issues/1039))

BUG FIXES:

* `azuread_group` - work around an API issue that prevented group creation for some configurations where the calling principal is specified as an owner ([#1037](https://github.com/hashicorp/terraform-provider-azuread/issues/1037))

## 2.35.0 (February 24, 2023)

BUG FIXES:

* `azuread_application_federated_identity_credential` - the `audiences` property now only supports a single value due to a breaking API change ([#1027](https://github.com/hashicorp/terraform-provider-azuread/issues/1027))
* `azuread_group` - only try to set additional fields when explicitly configured, to work around an API bug when application-only permissions are used ([#1028](https://github.com/hashicorp/terraform-provider-azuread/issues/1028))
* `azuread_service_principal` - resolve an issue where newly created service principals might not be found when specifying `use_existing = true` ([#1025](https://github.com/hashicorp/terraform-provider-azuread/issues/1025))

IMPROVEMENTS:

* **Provider:** support for the `metadata_host` property ([#1026](https://github.com/hashicorp/terraform-provider-azuread/issues/1026))
* **Provider:** authentication now uses the `github.com/hashicorp/go-azure-sdk/sdk/auth` package ([#1026](https://github.com/hashicorp/terraform-provider-azuread/issues/1026))
* **Provider:** cloud configuration now uses the `github.com/hashicorp/go-azure-sdk/sdk/environments` package ([#1026](https://github.com/hashicorp/terraform-provider-azuread/issues/1026))
* `data.azuread_application` - support for the `notes` attribute ([#1027](https://github.com/hashicorp/terraform-provider-azuread/issues/1027))
* `data.azuread_directory_roles` - support for the `template_ids` attribute ([#1011](https://github.com/hashicorp/terraform-provider-azuread/issues/1011))
* `azuread_application` - support for the `notes` property ([#1027](https://github.com/hashicorp/terraform-provider-azuread/issues/1027))
* `azuread_group` - support for the `administrative_unit_ids` property ([#984](https://github.com/hashicorp/terraform-provider-azuread/issues/984))
* `azuread_synchronization_job` - fix a bug where the incorrect API version was used, preventing this resource from working properly ([#1030](https://github.com/hashicorp/terraform-provider-azuread/issues/1030))
* `azuread_synchronization_secret` - fix a bug where the incorrect API version was used, preventing this resource from working properly ([#1030](https://github.com/hashicorp/terraform-provider-azuread/issues/1030))

## 2.34.1 (February 17, 2023)

BUG FIXES:

* `azuread_administrative_unit` - revert to the Microsoft Graph beta API version to resolve an API error when using this resource ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_application` - revert to the Microsoft Graph beta API version to resolve an issue preventing creation of new applications ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_application` - revert to the Microsoft Graph beta API version to resolve an issue preventing setting the `oauth2_post_response_required` property ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_application_pre_authorized` - revert to the Microsoft Graph beta API version to resolve an issue creating this resource ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_group` - revert to the Microsoft Graph beta API version to resolve an issue when managing group members ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_group_member` - revert to the Microsoft Graph beta API version to resolve an issue when managing group members ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))
* `azuread_user` - revert to the Microsoft Graph beta API version to resolve a persistent diff for the `account_enabled` and `show_in_address_list` properties ([#1023](https://github.com/hashicorp/terraform-provider-azuread/issues/1023))

## 2.34.0 (February 16, 2023)

IMPROVEMENTS:

* **Provider:** All resources will now explicitly use the Microsoft Graph v1.0 API unless stated otherwise in the provider documentation ([#990](https://github.com/hashicorp/terraform-provider-azuread/issues/990))
* `data.azuread_application` - support the `description` attribute ([#991](https://github.com/hashicorp/terraform-provider-azuread/issues/991))
* `azuread_application` - support app role and scope values up to 249 characters ([#1010](https://github.com/hashicorp/terraform-provider-azuread/issues/1010))

BUG FIXES:

* **Provider:** Support authentication scenarios where the `oid` claim is missing from the access token ([#1014](https://github.com/hashicorp/terraform-provider-azuread/issues/1014))
* `data.azuread_application_template` - revert a workaround from v2.31.0 and no longer use the beta API for this data source ([#987](https://github.com/hashicorp/terraform-provider-azuread/issues/987))
* `azuread_application` - work around an API bug where `mapped_claims_enabled` could be set on create when holding the `Application.ReadWrite.OwnedBy` role ([#1008](https://github.com/hashicorp/terraform-provider-azuread/issues/1008))

## 2.33.0 (January 19, 2023)

FEATURES:

* **New Resource:** `azuread_service_principal_token_signing_certificate` ([#968](https://github.com/hashicorp/terraform-provider-azuread/issues/968))

IMPROVEMENTS:

* `azuread_application` - support the `description` property ([#977](https://github.com/hashicorp/terraform-provider-azuread/issues/977))

BUG FIXES:

* `azuread_service_principal_delegated_permission_grant` - fix a bug that caused state refreshes to fail if the resource is edited outside Terraform ([#981](https://github.com/hashicorp/terraform-provider-azuread/issues/981))
* `azuread_group` - fix a validation bug to allow periods (.) in the `mail_nickname` property ([#979](https://github.com/hashicorp/terraform-provider-azuread/issues/979))
* `azuread_group` - fix a bug that prevents replacing a group when `prevent_duplicate_names = true` ([#980](https://github.com/hashicorp/terraform-provider-azuread/issues/980))
* `azuread_group` - set the `display_name` property correctly on creation so that triggered notification emails are correct ([#982](https://github.com/hashicorp/terraform-provider-azuread/issues/982))

## 2.32.0 (January 12, 2023)

FEATURES:

* **New Data Source:** `azuread_directory_roles` ([#945](https://github.com/hashicorp/terraform-provider-azuread/issues/945))

IMPROVEMENTS:

* `azuread_application` - support the `cloud_displayname` optional claim ([#967](https://github.com/hashicorp/terraform-provider-azuread/issues/967))

BUG FIXES:

* `azuread_application` - improve validation when checking for duplicate app roles and permission scopes ([#971](https://github.com/hashicorp/terraform-provider-azuread/issues/971))

## 2.31.0 (December 01, 2022)

IMPROVEMENTS:

* `azuread_application` - validation for `identifier_uris` to detect trailing slash with no path ([#928](https://github.com/hashicorp/terraform-provider-azuread/issues/928))

BUG FIXES:

* `data.azuread_application_template` - work around an API bug in the US Government cloud, by using the beta API ([#936](https://github.com/hashicorp/terraform-provider-azuread/issues/936))
* `azuread_application` - fix a bug where `owners` where not correctly removed ([#916](https://github.com/hashicorp/terraform-provider-azuread/issues/916))
* `azuread_application` - work around an API bug in the US Government cloud, by using the beta API when `template_id` is specified ([#936](https://github.com/hashicorp/terraform-provider-azuread/issues/936))

## 2.30.0 (October 28, 2022)

FEATURES:

* **New Resource:** `azuread_synchronization_job` ([#830](https://github.com/hashicorp/terraform-provider-azuread/issues/830))
* **New Resource:** `azuread_synchronization_secret` ([#830](https://github.com/hashicorp/terraform-provider-azuread/issues/830))

## 2.29.0 (September 29, 2022)

IMPROVEMENTS:

* Provider: support for the `oidc_token_file_path` property & `ARM_OIDC_TOKEN_FILE_PATH` environment variable ([#897](https://github.com/hashicorp/terraform-provider-azuread/issues/897))
* `data.azuread_service_principal` - this resource now makes use of the MS Graph v1.0 API instead of the beta API ([#896](https://github.com/hashicorp/terraform-provider-azuread/issues/896))
* `azuread_service_principal` - this resource now makes use of the MS Graph v1.0 API instead of the beta API ([#896](https://github.com/hashicorp/terraform-provider-azuread/issues/896))

## 2.28.1 (August 30, 2022)

BUG FIXES:

* **Provider:** fix a bug that could cause GitHub OIDC authentication to fail ([#876](https://github.com/hashicorp/terraform-provider-azuread/issues/876))

## 2.28.0 (August 25, 2022)

FEATURES

* **Provider:** support for generic OIDC authentication providers ([#874](https://github.com/hashicorp/terraform-provider-azuread/issues/874))
* **New Data Source:** `azuread_directory_object` ([#847](https://github.com/hashicorp/terraform-provider-azuread/issues/847))

IMPROVEMENTS:

* `azuread_application` - support `max_size_limit` as a value for the `additional_properties` property in the `optional_claims` block ([#864](https://github.com/hashicorp/terraform-provider-azuread/issues/864))

## 2.27.0 (August 05, 2022)

NOTES:

* This release contains a behavioral change for application/service principal passwords and certificates, when using a relative end date.

BUG FIXES:

* `data.azuread_group` - ensure security/mail enabled groups are excluded when explicitly `false` in config ([#841](https://github.com/hashicorp/terraform-provider-azuread/issues/841))
* `azuread_application_certificate` - calculate `end_date_relative` from the `start_date` and not the current timestamp ([#844](https://github.com/hashicorp/terraform-provider-azuread/issues/844))
* `azuread_application_password` - calculate `end_date_relative` from the `start_date` and not the current timestamp ([#844](https://github.com/hashicorp/terraform-provider-azuread/issues/844))
* `azuread_service_principal_certificate` - calculate `end_date_relative` from the `start_date` and not the current timestamp ([#844](https://github.com/hashicorp/terraform-provider-azuread/issues/844))
* `azuread_service_principal_password` - calculate `end_date_relative` from the `start_date` and not the current timestamp ([#844](https://github.com/hashicorp/terraform-provider-azuread/issues/844))

## 2.26.1 (July 11, 2022)

BUG FIXES:

* `azuread_directory_role_assignment` - fix a bug that required `directory_scope_id` to be set for unscoped assignments ([#840](https://github.com/hashicorp/terraform-provider-azuread/issues/840))

## 2.26.0 (July 08, 2022)

IMPROVEMENTS:

* `azuread_directory_role_assignment` - deprecate the `app_scope_object_id` property in favor of the `app_scope_id` property ([#837](https://github.com/hashicorp/terraform-provider-azuread/issues/837))
* `azuread_directory_role_assignment` - deprecate the `directory_scope_object_id` property in favor of the `directory_scope_id` property ([#837](https://github.com/hashicorp/terraform-provider-azuread/issues/837))

BUG FIXES:

* `azuread_directory_role_assignment` - fix incorrect schema validation for scoped role assignments ([#837](https://github.com/hashicorp/terraform-provider-azuread/issues/837))
* `azuread_directory_role_assignment` - fix a bug that was preventing the creation of some scoped role assignments ([#837](https://github.com/hashicorp/terraform-provider-azuread/issues/837))
* `azuread_group` - fix a bug where new group creation can error out before the timeout due to API inconsistency ([#838](https://github.com/hashicorp/terraform-provider-azuread/issues/838))
* `azuread_user` - only set `show_in_address_list` when changed in config as it is a potentially read-only attribute ([#831](https://github.com/hashicorp/terraform-provider-azuread/issues/831))

## 2.25.0 (June 23, 2022)

FEATURES:

* **New Resource:** `azuread_directory_role_assignment` (deprecates the `azuread_directory_role_member` resource) ([#826](https://github.com/hashicorp/terraform-provider-azuread/issues/826))

## 2.24.0 (June 16, 2022)

BUG FIXES:

* **Provider:** Fix a bug causing GitHub OIDC authentication to fail when consuming default environment variables ([#822](https://github.com/hashicorp/terraform-provider-azuread/issues/822))

## 2.23.0 (June 10, 2022)

FEATURES:

* **New Authentication Method:** Support for authenticating via OIDC with GitHub Actions ([#805](https://github.com/hashicorp/terraform-provider-azuread/issues/805))

IMPROVEMENTS:

* `azuread_user` - allow changing the `user_principal_name` property without recreating the user account ([#815](https://github.com/hashicorp/terraform-provider-azuread/issues/815))

BUG FIXES:

* **Provider:** Fix an Azure CLI authentication issue that could fail to autodetect the current tenant ID ([#819](https://github.com/hashicorp/terraform-provider-azuread/issues/819))
* `azuread_application_federated_identity_credential` - fix overly restrictive validation for the `audiences` property ([#808](https://github.com/hashicorp/terraform-provider-azuread/issues/808))
* `azuread_group` - fix a bug that could cause a crash when creating unified groups ([#816](https://github.com/hashicorp/terraform-provider-azuread/issues/816))

## 2.22.0 (April 28, 2022)

IMPROVEMENTS:

* `data.azuread_groups` - support the `ignore_missing` property ([#783](https://github.com/hashicorp/terraform-provider-azuread/issues/783))
* `azuread_conditional_access_policy` - support `linux` in the `included_platforms` and `excluded_platforms` properties ([#784](https://github.com/hashicorp/terraform-provider-azuread/issues/784))
* `azuread_group` - support the `SubscribeMembersToCalendarEventsDisabled` value in the `behaviors` property ([#785](https://github.com/hashicorp/terraform-provider-azuread/issues/785))

BUG FIXES:

* `data.azuread_service_principal` - raise an error when multiple results are found for the same `display_name` ([#781](https://github.com/hashicorp/terraform-provider-azuread/issues/781))
* `azuread_group` - ensure that unified groups can be created without a `description` ([#783](https://github.com/hashicorp/terraform-provider-azuread/issues/783))

## 2.21.0 (April 21, 2022)

BUG FIXES:

* `azuread_conditional_access_policy` - the `included_applications` property in the `conditions` block is now optional ([#775](https://github.com/hashicorp/terraform-provider-azuread/issues/775))
* `azuread_conditional_access_policy` - the `locations` and `platforms` blocks are now optional ([#775](https://github.com/hashicorp/terraform-provider-azuread/issues/775))

## 2.20.0 (April 08, 2022)

FEATURES:

* **New Resource:** `azuread_claims_mapping_policy` ([#733](https://github.com/hashicorp/terraform-provider-azuread/issues/733)) ([#766](https://github.com/hashicorp/terraform-provider-azuread/issues/766))
* **New Resource:** `azuread_service_principal_claims_mapping_policy_assignment` ([#733](https://github.com/hashicorp/terraform-provider-azuread/issues/733)) ([#766](https://github.com/hashicorp/terraform-provider-azuread/issues/766))

## 2.19.1 (March 11, 2022)

BUG FIXES:

* `azuread_application` - revert an earlier change for validation of role/scope values ([#756](https://github.com/terraform-providers/terraform-provider-azuread/issues/756))

## 2.19.0 (March 11, 2022)

IMPROVEMENTS:

* `data.azuread_service_principals` - export the `object_id` property in the `service_principals` list ([#749](https://github.com/terraform-providers/terraform-provider-azuread/issues/749))

BUG FIXES:

* `azuread_application` - add a missing validation check for role/scope values ([#750](https://github.com/terraform-providers/terraform-provider-azuread/issues/750))
* `azuread_conditional_access_policy` - fix a crash during the plan phase when `session_controls` is empty ([#747](https://github.com/terraform-providers/terraform-provider-azuread/issues/747))

## 2.18.0 (February 11, 2022)

BUG FIXES:

* `azuread_group` - make the `auto_subscribe_new_members`, `external_senders_allowed`, `hide_from_address_lists` and `hide_from_outlook_clients` properties Computed to avoid setting them unnecessarily ([#731](https://github.com/terraform-providers/terraform-provider-azuread/issues/731))

## 2.17.0 (February 03, 2022)

FEATURES:

* **New Resource:** `azuread_custom_directory_role` ([#728](https://github.com/terraform-providers/terraform-provider-azuread/issues/728))

## 2.16.0 (January 28, 2022)

IMPROVEMENTS:

* `data.azuread_group` - support for the `allow_external_senders`, `auto_subscribe_new_members`, `hide_from_address_lists` and `hide_from_outlook_clients` attributes ([#723](https://github.com/terraform-providers/terraform-provider-azuread/issues/723))
* `azuread_group` - support for the `allow_external_senders`, `auto_subscribe_new_members`, `hide_from_address_lists` and `hide_from_outlook_clients` properties ([#723](https://github.com/terraform-providers/terraform-provider-azuread/issues/723))

## 2.15.0 (January 14, 2022)

IMPROVEMENTS:

* `data.azuread_group` - support the `display_name_prefix` property ([#716](https://github.com/terraform-providers/terraform-provider-azuread/issues/716))

BUG FIXES:

* `azuread_application` - remove an unnecessary API call that may require additional permissions, when assigning owners ([#713](https://github.com/terraform-providers/terraform-provider-azuread/issues/713))
* `azuread_service_principal` - remove an unnecessary API call that may require additional permissions, when assigning owners ([#713](https://github.com/terraform-providers/terraform-provider-azuread/issues/713))

## 2.14.0 (January 07, 2022)

FEATURES:

* **New Resource:** `azuread_application_federated_identity_credential` ([#705](https://github.com/terraform-providers/terraform-provider-azuread/issues/705))

IMPROVEMENTS:

* `azuread_service_principal_password`: re-add support for `display_name`, `start_date`, `end_date` and `end_date_relative` properties ([#706](https://github.com/terraform-providers/terraform-provider-azuread/issues/706))

## 2.13.0 (December 15, 2021)

IMPROVEMENTS:

* `azuread_group`: support for `dynamic_memberships` ([#695](https://github.com/terraform-providers/terraform-provider-azuread/issues/695))

## 2.12.0 (December 03, 2021)

IMPROVEMENTS:

* `azuread_conditional_access_policy` - support the `persistent_browser_mode` in the `session_controls` block ([#677](https://github.com/terraform-providers/terraform-provider-azuread/issues/677))

BUG FIXES:

* `azuread_application` - allow URNs to be used in `redirect_uris` in the `public_client` block ([#684](https://github.com/terraform-providers/terraform-provider-azuread/issues/684))
* `azuread_service_principal_delegated_permission_grant` - add missing support for importing this resource ([#685](https://github.com/terraform-providers/terraform-provider-azuread/issues/685))

## 2.11.0 (November 25, 2021)

BREAKING CHANGES:

* **Provider:** support for the German national cloud, which was [closed down as of October 29, 2021](https://www.microsoft.com/en-us/cloud-platform/germany-cloud-regions), has been removed in this release ([#670](https://github.com/terraform-providers/terraform-provider-azuread/issues/670))

FEATURES:

* **New Data Source:** `azuread_administrative_unit` ([#672](https://github.com/terraform-providers/terraform-provider-azuread/issues/672))
* **New Resource:** `azuread_administrative_unit` ([#672](https://github.com/terraform-providers/terraform-provider-azuread/issues/672))
* **New Resource:** `azuread_administrative_unit_member` ([#672](https://github.com/terraform-providers/terraform-provider-azuread/issues/672))
* **New Resource:** `azuread_service_principal_delegated_permission_grant` ([#676](https://github.com/terraform-providers/terraform-provider-azuread/issues/676))

IMPROVEMENTS:

* `azuread_conditional_access_policy` - support the `devices` block ([#673](https://github.com/terraform-providers/terraform-provider-azuread/issues/673))

BUG FIXES:

* `azuread_conditional_access_policy` - fix a bug when removing the `session_controls` block from a policy ([#673](https://github.com/terraform-providers/terraform-provider-azuread/issues/673))

## 2.10.0 (November 19, 2021)

BUG FIXES:

* `azuread_group` - fix a bug that prevented removing all `members` of a group ([#666](https://github.com/terraform-providers/terraform-provider-azuread/issues/666))

## 2.9.0 (November 12, 2021)

BUG FIXES:

* **Provider:** fix an authentication bug that prevented authorizing using a Managed Identity when running in Azure Cloud Shell ([#660](https://github.com/terraform-providers/terraform-provider-azuread/issues/660))
* `data.azuread_user` - ensure apostrophes are correctly quoted when matching by `mail_nickname` or `user_principal_name` ([#643](https://github.com/terraform-providers/terraform-provider-azuread/issues/643))
* `data.azuread_users` - ensure apostrophes are correctly quoted when matching by `mail_nicknames` or `user_principal_names` ([#643](https://github.com/terraform-providers/terraform-provider-azuread/issues/643))
* `azuread_application_certificate` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_application_password` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_application` - add mitigation for replication delay when creating new applications ([#656](https://github.com/terraform-providers/terraform-provider-azuread/issues/656))
* `azuread_directory_role_member` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_group_member` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_group` - add mitigation for replication delay when creating new groups ([#656](https://github.com/terraform-providers/terraform-provider-azuread/issues/656))
* `azuread_group` - work around an API consistency issue when creating and deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_invitation` - work around an API consistency issue when creating and deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_service_principal_certificate` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_service_principal_password` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_service_principal` - add mitigation for replication delay when creating new service principals ([#656](https://github.com/terraform-providers/terraform-provider-azuread/issues/656))
* `azuread_service_principal` - work around an API consistency issue when creating and deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))
* `azuread_user` - add mitigation for replication delay when creating new users ([#656](https://github.com/terraform-providers/terraform-provider-azuread/issues/656))
* `azuread_user` - work around an API consistency issue when deleting resources ([#659](https://github.com/terraform-providers/terraform-provider-azuread/issues/659))

## 2.8.0 (October 28, 2021)

BUG FIXES:

* `azuread_application` - allow custom URI schemes for public client redirect URIs ([#647](https://github.com/terraform-providers/terraform-provider-azuread/issues/647))
* `azuread_group` - ensure `mail_nickname` is set for all groups when specified in configuration ([#645](https://github.com/terraform-providers/terraform-provider-azuread/issues/645))

## 2.7.0 (October 15, 2021)

IMPROVEMENTS:

* **Provider:** log the claims from access tokens for improved debugging ability ([#623](https://github.com/terraform-providers/terraform-provider-azuread/issues/623))
* `azuread_user` - support for the `manager_id` property ([#628](https://github.com/terraform-providers/terraform-provider-azuread/issues/628))
* `azuread_application` - support for the `feature_tags` block and the `tags` property ([#630](https://github.com/terraform-providers/terraform-provider-azuread/issues/630))
* `azuread_service_principal` - the `features` block has been deprecated in favour of the `feature_tags` for clarity ([#630](https://github.com/terraform-providers/terraform-provider-azuread/issues/630))

## 2.6.0 (October 07, 2021)

IMPROVEMENTS:

* **Provider:** Generate and log request/response correlation IDs for improved inspection ability in HTTP traces ([#621](https://github.com/terraform-providers/terraform-provider-azuread/issues/621))

BUG FIXES:

* **Provider:** Implement a workaround for a breaking API change affecting all resources having relationships such as `members` and `owners` ([#616](https://github.com/terraform-providers/terraform-provider-azuread/issues/616))
* `azuread_application_certificate` - fix an eventual consistency issue when creating new certificates ([#618](https://github.com/terraform-providers/terraform-provider-azuread/issues/618))
* `azuread_application_password` - fix an eventual consistency issue when creating new passwords ([#618](https://github.com/terraform-providers/terraform-provider-azuread/issues/618))
* `azuread_service_principal_certificate` - fix an eventual consistency issue when creating new certificates ([#618](https://github.com/terraform-providers/terraform-provider-azuread/issues/618))
* `azuread_service_principal_password` - fix an eventual consistency issue when creating new passwords ([#618](https://github.com/terraform-providers/terraform-provider-azuread/issues/618))

## 2.5.0 (September 30, 2021)

IMPROVEMENTS:

* `data.azuread_groups` - support the `mail_enabled` and `security_enabled` properties ([#603](https://github.com/terraform-providers/terraform-provider-azuread/issues/603))
* `data.azuread_user` - support the `cost_center`, `division` and `employee_type` attributes ([#597](https://github.com/terraform-providers/terraform-provider-azuread/issues/597))
* `azuread_user` - support the `cost_center`, `division` and `employee_type` properties ([#597](https://github.com/terraform-providers/terraform-provider-azuread/issues/597))

BUG FIXES:

* `azuread_application` - support for "myapp://auth" as a public client redirect URI, to support B2C IEF applications ([#607](https://github.com/terraform-providers/terraform-provider-azuread/issues/607))
* `azuread_application` - ensure that `prevent_duplicate_names` does not fail incorrectly when `display_name` is not known at plan time ([#596](https://github.com/terraform-providers/terraform-provider-azuread/issues/596))
* `azuread_group` - ensure that `prevent_duplicate_names` does not fail incorrectly when `display_name` is not known at plan time ([#596](https://github.com/terraform-providers/terraform-provider-azuread/issues/596))
* `azuread_service_principal` - fix a bug that prevented `features` from being empty or having all disabled properties ([#602](https://github.com/terraform-providers/terraform-provider-azuread/issues/602))

## 2.4.0 (September 23, 2021)

FEATURES:

* **New Resource:** `azuread_app_role_assignment` ([#584](https://github.com/terraform-providers/terraform-provider-azuread/issues/584))

IMPROVEMENTS:

* `azuread_application_password` - support the `rotate_when_changed` property (this was previously available as an undocumented property `keepers`) ([#572](https://github.com/terraform-providers/terraform-provider-azuread/issues/572))
* `azuread_service_principal_password` - support the `rotate_when_changed` property (this was previously available as an undocumented property `keepers`) ([#572](https://github.com/terraform-providers/terraform-provider-azuread/issues/572))

## 2.3.0 (September 16, 2021)

FEATURES:

* **New Resource:** `azuread_directory_role` ([#573](https://github.com/terraform-providers/terraform-provider-azuread/issues/573))
* **New Resource:** `azuread_directory_role_member` ([#573](https://github.com/terraform-providers/terraform-provider-azuread/issues/573))

IMPROVEMENTS:

* `data.azuread_service_principal` - support the `features` block ([#571](https://github.com/terraform-providers/terraform-provider-azuread/issues/571))
* `azuread_application` - support the `logo_image` property ([#574](https://github.com/terraform-providers/terraform-provider-azuread/issues/574))
* `azuread_application` - allow URNs to be specified for web redirect URIs ([#577](https://github.com/terraform-providers/terraform-provider-azuread/issues/577))
* `azuread_service_principal` - support the `features` block ([#571](https://github.com/terraform-providers/terraform-provider-azuread/issues/571))

BUG FIXES:

* `azuread_conditional_access_policy` - resolve a number of bugs related to updating an existing conditional access policy ([#569](https://github.com/terraform-providers/terraform-provider-azuread/issues/569))

## 2.2.1 (September 10, 2021)

BUG FIXES:

* **Provider:** fix a bug in handling retried requests that could cause errors when attempting to read a resource that no longer exists ([#564](https://github.com/terraform-providers/terraform-provider-azuread/issues/564))

## 2.2.0 (September 10, 2021)

FEATURES:

* **New Data Source:** `azuread_application_template` ([#554](https://github.com/terraform-providers/terraform-provider-azuread/issues/554))
* **New Data Source:** `azuread_service_principals` ([#555](https://github.com/terraform-providers/terraform-provider-azuread/issues/555))
* **New Resource:** `azuread_conditional_access_policy` ([#466](https://github.com/terraform-providers/terraform-provider-azuread/issues/466))
* **New Resource:** `azuread_named_location` ([#441](https://github.com/terraform-providers/terraform-provider-azuread/issues/441))

IMPROVEMENTS:

* `azuread_application` - support for the `template_id` property for creating applications (and service principals) from a template ([#554](https://github.com/terraform-providers/terraform-provider-azuread/issues/554))
* `azuread_service_principal` - support the `saml_single_sign_on` block containing the `relay_state` property ([#557](https://github.com/terraform-providers/terraform-provider-azuread/issues/557))
* `azuread_user` - support the `disable_password_expiration` and `disable_strong_password` properties ([#550](https://github.com/terraform-providers/terraform-provider-azuread/issues/550))

BUG FIXES:

* **Provider:** fix a decoding bug when parsing claims from an access token ([#560](https://github.com/terraform-providers/terraform-provider-azuread/issues/560))
* **Provider:** attempt to detect when using Azure CLI authentication in Azure Cloud Shell and avoid specifying the tenant ID ([#560](https://github.com/terraform-providers/terraform-provider-azuread/issues/560))
* `azuread_group` - fix an API error caused by duplicate `owners` being mistakenly sent when creating new groups ([#553](https://github.com/terraform-providers/terraform-provider-azuread/issues/553))

## 2.1.0 (September 02, 2021)

FEATURES:

* **New Resource:** `azuread_invitation` ([#445](https://github.com/terraform-providers/terraform-provider-azuread/issues/445))

BUG FIXES:

* `data.azuread_client_config` - populate the `tenant_id` and `client_id` attributes when authenticating via Azure CLI ([#539](https://github.com/terraform-providers/terraform-provider-azuread/issues/539))
* `azuread_service_principal` - fix a bug that prevented creation of service principals in some cases due to `owners` being applied incorrectly ([#539](https://github.com/terraform-providers/terraform-provider-azuread/issues/539))
* `azuread_user` - fix a validation bug for the `password` property ([#543](https://github.com/terraform-providers/terraform-provider-azuread/issues/543))

IMPROVEMENTS:

* `data.azuread_groups` - support the `return_all` property ([#520](https://github.com/terraform-providers/terraform-provider-azuread/issues/520))
* `data.azuread_users` - support the `return_all` property ([#513](https://github.com/terraform-providers/terraform-provider-azuread/issues/513))
* `azuread_application` - allow `redirect_uris` with a scheme of `ms-appx-web` ([#540](https://github.com/terraform-providers/terraform-provider-azuread/issues/540))

## 2.0.1 (August 26, 2021)

BUG FIXES:

* `azuread_application` - fix a bug where unknown IDs or values for roles/scopes were incorrectly flagged as duplicates ([#528](https://github.com/terraform-providers/terraform-provider-azuread/issues/528))

## 2.0.0 (August 26, 2021)

NOTES:

* **Major Version:** This is a major version upgrade which contains breaking changes. Please read the [Upgrade Guide](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/microsoft-graph) before upgrading, which details all the known breaking changes that practitioners should be aware of.
* **Microsoft Graph:** The upstream API for Azure Active Directory is now Microsoft Graph, and the deprecated Azure Active Directory Graph API is no longer supported.

FEATURES:

* **Provider:** Client Certificate authentication now supports specifying an inline certificate ([#490](https://github.com/terraform-providers/terraform-provider-azuread/issues/490))
* **New Data Source:** `azuread_application_published_app_ids` ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* **New Resource:** `application_pre_authorized` ([#472](https://github.com/terraform-providers/terraform-provider-azuread/issues/472))

IMPROVEMENTS:

* `data.azuread_application` - the `api` block now supports the `accept_mapped_claims`, `known_client_applications` and `requested_access_token_version` attributes ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - the `implicit_grant` block now supports the `id_token_issuance_enabled` attribute ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the `optional_claims` block now supports the `saml2_token` attribute ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - export the `disabled_by_microsoft` attribute ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `device_only_auth_enabled` and `oauth2_post_response_required` attributes ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `logo_url`, `marketing_url`, `privacy_statement_url` and `terms_of_service_url` attributes ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `publisher_domain` attribute ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `public_client` block ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `single_page_application` block ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_application` - export the `app_role_ids` and `oauth2_permission_scope_ids` attributes ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `data.azuread_domains` - export the `admin_managed`, `root` and `supported_services` attributes for each domain ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_domains` - support the `admin_managed`, `only_root` and `supports_services` properties ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_group` - export the `assignable_to_role`, `behaviors`, `mail_nickname`, `theme` and `visibility` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_group` - export the `mail`, `preferred_language` and `proxy_addresses` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_group` - export the `onpremises_domain_name`, `onpremises_netbios_name`, `onpremises_sam_account_name`, `onpremises_security_identifier` and `onpremises_sync_enabled` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_service_principal` - export the `account_enabled`, `login_url` and `preferred_single_sign_on_mode` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `data.azuread_service_principal` - export the `alternative_names`, `description`, `notes` and `notification_email_addresses` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `data.azuread_service_principal` - export the `app_role_ids` and `oauth2_permission_scope_ids` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `data.azuread_service_principal` - export the `application_tenant_id`, `display_name`, `service_principal_names`, `sign_in_audience` and `type` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `data.azuread_service_principal` - export the `homepage_url`, `logout_url`, `redirect_uris` and `saml_metadata_url` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `data.azuread_user` - export the `age_group` and `consent_provided_for_minor` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_user` - export the `business_phones`, `employee_id`, `fax_number` and `preferred_language` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_user` - export the `mail`, `other_mails` and `show_in_address_list` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_user` - export the `creation_type`, `external_user_state`, `im_addresses` and `proxy_addresses` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `data.azuread_user` - export the `onpremises_distinguished_name`, `onpremises_domain_name`, `onpremises_security_identifier` and `onpremises_sync_enabled` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_application` - the `api` block now supports the `accept_mapped_claims`, `known_client_applications` and `requested_access_token_version` properties ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - the `implicit_grant` block now supports the `id_token_issuance_enabled` property ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `optional_claims` block now supports the `saml2_token` block ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `sign_in_audience` property now supports the `AzureADandPersonalMicrosoftAccount` and `PersonalMicrosoftAccount` values ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - export the `disabled_by_microsoft` attribute ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - export the `publisher_domain` attribute ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - support the `device_only_auth_enabled` and `oauth2_post_response_required` properties ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - support the `logo_url`, `marketing_url`, `privacy_statement_url` and `terms_of_service_url` properties ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - support for the `public_client` block ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - support for the `single_page_application` block ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application` - export the `app_role_ids` and `oauth2_permission_scope_ids` attributes ([#474](https://github.com/terraform-providers/terraform-provider-azuread/issues/474))
* `azuread_application_password` - support the `keepers` property ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_group` - support for creating mail-enabled groups ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - support for creating Microsoft 365 groups ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - support for updating groups without recreating them ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - support the `assignable_to_role`, `behaviors`, `mail_nickname`, `theme` and `visibility` properties ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_group` - export the `mail`, `preferred_language` and `proxy_addresses` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_group` - export the `onpremises_domain_name`, `onpremises_netbios_name`, `onpremises_sam_account_name`, `onpremises_security_identifier` and `onpremises_sync_enabled` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_service_principal` - support the `account_enabled`, `login_url` and `preferred_single_sign_on_mode` properties ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal` - support the `alternative_names`, `description`, `notes` and `notification_email_addresses` properties ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal` - support the `owners` property ([#519](https://github.com/terraform-providers/terraform-provider-azuread/issues/519))
* `azuread_service_principal` - support the `use_existing` property ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal` - export the `app_role_ids` and `oauth2_permission_scope_ids` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal` - export the `application_tenant_id`, `display_name`, `service_principal_names`, `sign_in_audience` and `type` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal` - export the `homepage_url`, `logout_url`, `redirect_uris` and `saml_metadata_url` attributes ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_service_principal_password` - support the `keepers` property ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_user` - support the `age_group` and `consent_provided_for_minor` properties ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_user` - support the `business_phones`, `employee_id`, `fax_number` and `preferred_language` properties ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_user` - support the `mail`, `other_mails` and `show_in_address_list` properties ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_user` - export the `creation_type`, `external_user_state`, `im_addresses` and `proxy_addresses` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))
* `azuread_user` - export the `onpremises_distinguished_name`, `onpremises_domain_name`, `onpremises_security_identifier` and `onpremises_sync_enabled` attributes ([#476](https://github.com/terraform-providers/terraform-provider-azuread/issues/476))

BUG FIXES:

* `azuread_application` - resolved an issue where `identifier_uris` could be reordered and cause a persistent diff ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `identifier_uris` property can now be set for all applications regardless of target platform ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - fixed a bug where app roles could be duplicated or left in a disabled state ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - fixed a bug where app roles could not be removed from an application ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - fixed a bug where the `enabled` property of app roles could be ignored ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - fixed a bug where the `id` property of app roles could be undesirably changed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - resolved an issue where the default scope could not be removed from an application ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - resolved an issue where multiple `group_membership_claims` could not be specified ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application_password` - the `display_name` / `description` properties are no longer stored using the `customKeyIdentifier` API field, lifting the 32 byte limit ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - fix a bug where `owners` or `members` would sometimes not be updated ([#519](https://github.com/terraform-providers/terraform-provider-azuread/issues/519))
* `azuread_group` - fix some ownership-related bugs where groups could sometimes not be created or updated ([#519](https://github.com/terraform-providers/terraform-provider-azuread/issues/519))
* `azuread_user` - resolved an issue where importing users would inadvertently reset their password ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))

BREAKING CHANGES:

* `data.azuread_domains` - the `is_` prefix has been dropped from all exported attributes ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the `display_name` property is now matched case-insensitively which mirrors the behaviour of Azure Active Directory ([#492](https://github.com/terraform-providers/terraform-provider-azuread/issues/492))
* `data.azuread_application` - the deprecated property `name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `available_to_other_tenants` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the `group_membership_claims` attribute has changed from a string to a list of strings ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `homepage` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `logout_url` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `oauth2_allow_implicit_flow` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `oauth2_permissions` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the `public_client` attribute is now a block containing public client settings ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `reply_urls` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_application` - the deprecated attribute `type` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_group` - the deprecated property `name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_groups` - the deprecated property `names` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_service_principal` - the deprecated attribute `oauth2_permissions` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_user` - the deprecated attribute `immutable_id` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_user` - the deprecated attribute `physical_delivery_office_name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_user` - the deprecated attribute `mobile` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `data.azuread_users` - the deprecated attribute `immutable_id` in the `users` block has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `api` block is no longer Computed, omitting this block will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `app_role` block is no longer Computed, omitting this block will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `id` property in the `app_role` block is now Required ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `available_to_other_tenants` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `fallback_public_client_enabled` property is no longer Computed, omitting this property will cause the default value to be applied ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `group_membership_claims` property has changed from a string to a set of strings ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `homepage` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `identifier_uris` property is no longer Computed, omitting this property will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `identifier_uris` property has changed from a List to a Set to resolve an API ordering issue ([#481](https://github.com/terraform-providers/terraform-provider-azuread/issues/481))
* `azuread_application` - the deprecated property `logout_url` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `oauth2_allow_implicit_flow` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `oauth2_permission_scope` block is no longer Computed, omitting this block will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated block `oauth2_permissions` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `owners` property is no longer Computed, omitting this property will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `public_client` property is now a block containing public client settings ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `reply_urls` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `sign_in_audience` property is no longer Computed, omitting this property will cause the default value to be applied ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the deprecated property `type` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application` - the `web` block is no longer Computed, omitting this block will cause it to be removed from your configuration ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_application_password` - the `key_id` and `value` properties are now Computed, due to API changes it is no longer possible to specify these values ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - the deprecated property `name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_group` - at least one of the `mail_enabled` or `security_enabled` properties are now Required ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_service_principal` - the deprecated attribute `oauth2_permissions` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_service_principal_password` - the `key_id` and `value` properties are now Computed, due to API changes it is no longer possible to specify these values ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_service_principal_password` - the `start_date` and `end_date` properties are now Computed, due to an API issue it is no longer possible to specify these values ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_user` - the deprecated property `immutable_id` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_user` - the deprecated property `physical_delivery_office_name` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))
* `azuread_user` - the deprecated property `mobile` has been removed ([#461](https://github.com/terraform-providers/terraform-provider-azuread/issues/461))

## 1.6.0 (June 24, 2021)

DEPRECATIONS:

* `azuread_application_app_role` - this resource is deprecated and will be removed in version 2.0 ([#465](https://github.com/terraform-providers/terraform-provider-azuread/issues/465))
* `azuread_application_oauth2_permission` - this resource is deprecated and will be removed in version 2.0 ([#465](https://github.com/terraform-providers/terraform-provider-azuread/issues/465))
* `azuread_application_oauth2_permission_scope` - this resource is deprecated and will be removed in version 2.0 ([#465](https://github.com/terraform-providers/terraform-provider-azuread/issues/465))

## 1.5.1 (June 10, 2021)

BUG FIXES:

* **Provider:** Suppress a spurious deprecation notice for the `metadata_host` provider field ([#439](https://github.com/terraform-providers/terraform-provider-azuread/issues/439))
* `azuread_application_password` - fix a bug that prevented specifying the `display_name`, `start_date`, `end_date` or `end_date_relative` properties when using Microsoft Graph ([#444](https://github.com/terraform-providers/terraform-provider-azuread/issues/444))
* `azuread_group` - fix a bug that prevented creating a group with more than 20 owners or members ([#454](https://github.com/terraform-providers/terraform-provider-azuread/issues/454))
* `azuread_service_principal_password` - fix a bug that prevented specifying the `display_name`, `start_date`, `end_date` or `end_date_relative` properties when using Microsoft Graph ([#444](https://github.com/terraform-providers/terraform-provider-azuread/issues/444))

## 1.5.0 (May 20, 2021)

NOTES:

* **Support for Microsoft Graph:** This release introduces beta support for [Microsoft Graph](https://docs.microsoft.com/en-us/graph/overview) in a way that is forward (and backward) compatible with the current [Azure Active Directory Graph](https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-graph-api) API implementation. We do not recommend enabling this beta _in production_ at this time, but encourage you to try it out in test environments where minimal impact can occur if something doesn't work as expected. See the [Migration Guide](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/microsoft-graph#beta-support-for-microsoft-graph-in-v150) for more details.

* **Deprecations:** This release contains a number of additional deprecations to aid in future upgrades to version 2.0 of this provider. These will be flagged when running Terraform, and are documented in detail in the [Migration Guide](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/guides/microsoft-graph). Existing configurations will continue to work unchanged for any v1.x release, regardless of which API is used.

IMPROVEMENTS:

* `data.azuread_user` - export the `user_type` attribute ([#406](https://github.com/terraform-providers/terraform-provider-azuread/issues/406))
* `azuread_user` - export the `user_type` attribute ([#401](https://github.com/terraform-providers/terraform-provider-azuread/issues/401)] / [[#413](https://github.com/terraform-providers/terraform-provider-azuread/issues/413))

BUG FIXES:

* `azuread_application` - validation for the `identifier_uris` property now supports URNs ([#426](https://github.com/terraform-providers/terraform-provider-azuread/issues/426))

## 1.4.0 (February 18, 2021)

IMPROVEMENTS:

* dependencies: updating to build using Go 1.16 which adds support for `darwin/arm64` (Apple Silicon) ([#403](https://github.com/hashicorp/terraform-provider-azuread/issues/403))
* Data Source: `azuread_group` - support for the `mail_enabled` and `security_enabled` properties ([#393](https://github.com/hashicorp/terraform-provider-azuread/issues/393))
* `azuread_group` - support for the `mail_enabled` and `security_enabled` attributes ([#393](https://github.com/hashicorp/terraform-provider-azuread/issues/393))

## 1.3.0 (January 28, 2021)

IMPROVEMENTS:

* `azuread_application_certificate` - support for base64 and hex encoded certificate values ([#386](https://github.com/hashicorp/terraform-provider-azuread/issues/386))
* `azuread_service_principal_certificate` - support for base64 and hex encoded certificate values ([#386](https://github.com/hashicorp/terraform-provider-azuread/issues/386))

## 1.2.2 (January 16, 2021)

BUGFIXES:

* `azuread_application` - set the display name correctly when creating/updating applications using the `display_name` property

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

* `azuread_application` - resolves an issue where setting `prevent_duplicate_names = true` causes an error for new applications ([#367](https://github.com/hashicorp/terraform-provider-azuread/issues/367))
* `azuread_application` - fixes a bug where the default owner for a new application is removed ([#366](https://github.com/hashicorp/terraform-provider-azuread/issues/366))

## 1.1.0 (November 25, 2020)

FEATURES:

* Added a flag to allow users to customize the Partner ID or opt-out of the default Terraform Partner ID ([#350](https://github.com/hashicorp/terraform-provider-azuread/issues/350))
* This release includes updated support for working directly with tenants using Azure CLI authentication. We recommend the use of `az login --allow-no-subscription` to populate tenant-level accounts (which have no subscriptions).

IMPROVEMENTS:

* `data.azuread_user` - support the `given_name`, `surname`, `job_title`, `department`, `company_name`, `physical_delivery_office_name`, `street_address`, `city`, `state`, `country`, `postal_code` and `mobile` attribute ([#351](https://github.com/hashicorp/terraform-provider-azuread/issues/351))
* `azuread_user` - support the `given_name`, `surname`, `job_title`, `department`, `company_name`, `physical_delivery_office_name`, `street_address`, `city`, `state`, `country`, `postal_code` and `mobile` properties ([#351](https://github.com/hashicorp/terraform-provider-azuread/issues/351))

BUG FIXES:

* **Provider:** Fixed an issue where CLI authentication produced a `parsing json result` error during provider initialization ([#358](https://github.com/hashicorp/terraform-provider-azuread/issues/358))
* `azuread_application` - enable removal of owners on existing applications, and creation of applications with no owners ([#355](https://github.com/hashicorp/terraform-provider-azuread/issues/355))
* `azuread_application` - fixed a bug where specifying the `prevent_duplicate_names` property would report a false positive on update. ([#338](https://github.com/hashicorp/terraform-provider-azuread/issues/338))

## 1.0.0 (September 03, 2020)

NOTES:

* **Major Version:** This is a major version upgrade which contains some breaking changes as detailed below.
* **Terraform 0.10/0.11:** This version of the provider requires Terraform 0.12.x or later and will not work with earlier versions.

FEATURES:

* New resource: `azuread_application_app_role` ([#150](https://github.com/hashicorp/terraform-provider-azuread/issues/150)] [[#306](https://github.com/hashicorp/terraform-provider-azuread/issues/306))
* New resource: `azuread_application_oauth2_permission` ([#267](https://github.com/hashicorp/terraform-provider-azuread/issues/267))

BREAKING CHANGES:

* `azuread_application` - a default value for the `homepage` property is no longer derived when unspecified ([#268](https://github.com/hashicorp/terraform-provider-azuread/issues/268))
* `azuread_application_password` - the deprecated `application_id` property has been removed
* `data.azuread_group` - the `name` property is now case-insensitive ([#246](https://github.com/hashicorp/terraform-provider-azuread/issues/246))
* `data.azuread_groups` and `data.azuread_users` will not error if no results found

## 0.11.0 (July 09, 2020)

IMPROVEMENTS:

* Provider: no longer require configuring `subscription_id` (configuration value) / `ARM_SUBSCRIPTION_ID` (environment variable). ([#271](https://github.com/hashicorp/terraform-provider-azuread/issues/271))
* `data.azuread_client_config` - deprecate the `subscription_id` property. For compatibility, still populates `subscription_id` if the provider is configured with a subscription ID ([#271](https://github.com/hashicorp/terraform-provider-azuread/issues/271))
* `data.azuread_application` - support for the `application_id` property ([#274](https://github.com/hashicorp/terraform-provider-azuread/issues/274))
* `data.azuread_users` - support the `ignore_missing` property ([#256](https://github.com/hashicorp/terraform-provider-azuread/issues/256))
* `data.azuread_users` - export the `users` attribute containing a list of users with additional properties ([#256](https://github.com/hashicorp/terraform-provider-azuread/issues/256))
* `azuread_application` - support the `prevent_duplicate_names` property ([#279](https://github.com/hashicorp/terraform-provider-azuread/issues/279))
* `azuread_application` - validate `app_roles` and `oauth2_permissions` to check for duplicate `value`s ([#287](https://github.com/hashicorp/terraform-provider-azuread/issues/287))
* `azuread_group` - support the `prevent_duplicate_names` property ([#279](https://github.com/hashicorp/terraform-provider-azuread/issues/279))

BUG FIXES:

* `azuread_group` - remediate AAD replication delays when adding/removing group members ([#283](https://github.com/hashicorp/terraform-provider-azuread/issues/283))
* `azuread_group` - remediate AAD replication delays after group creation, before setting owners/members ([#290](https://github.com/hashicorp/terraform-provider-azuread/issues/290))

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

* **New Resource:** `azuread_application_certificate` ([#262](https://github.com/hashicorp/terraform-provider-azuread/issues/262))
* **New Resource:** `azuread_service_principal_certificate` ([#262](https://github.com/hashicorp/terraform-provider-azuread/issues/262))

IMPROVEMENTS:

* `azuread_application` - support for the `optional_claims` property, for access tokens and ID tokens ([#260](https://github.com/hashicorp/terraform-provider-azuread/issues/260))
* `azuread_application` - support for the `oauth2_permissions` property ([#252](https://github.com/hashicorp/terraform-provider-azuread/issues/252))
* `azuread_application_password` - support the `description` property ([#253](https://github.com/hashicorp/terraform-provider-azuread/issues/253))
* `azuread_service_principal_password` - support the `description` property ([#253](https://github.com/hashicorp/terraform-provider-azuread/issues/253))
* `data.azuread_users` - support empty lists for `user_principal_names`/`object_ids`/`mail_nicknames` properties ([#258](https://github.com/hashicorp/terraform-provider-azuread/issues/258))
* `data.azuread_groups` - support empty lists for `names`/`object_ids` properties ([#257](https://github.com/hashicorp/terraform-provider-azuread/issues/257))

BUG FIXES:

* `azuread_application_password` and `azuread_service_principal_password` - Plan-time validation for `end_date` / `end_date_relative` ([#261](https://github.com/hashicorp/terraform-provider-azuread/issues/261))
* `azuread_application_password` and `azuread_service_principal_password` - Change the resource ID format to mitigate potential UUID collision ([#264](https://github.com/hashicorp/terraform-provider-azuread/issues/264))

## 0.9.0 (May 15, 2020)

DEPENDENCIES:

* upgrade `azure-sdk-for-go` to `v42.1.0` ([#247](https://github.com/hashicorp/terraform-provider-azuread/issues/247))

IMPROVEMENTS:

* `azuread_application` - the `group_membership_claims` property now supports `ApplicationGroup` ([#238](https://github.com/hashicorp/terraform-provider-azuread/issues/238))
* `azuread_service_principal` - changing the `tags` property no longer forces a new resource ([#245](https://github.com/hashicorp/terraform-provider-azuread/issues/245))

BUG FIXES:

* `data.azuread_user` - use `equals` instead of `startsWith` when looking uo users by `mailNickname` ([#251](https://github.com/hashicorp/terraform-provider-azuread/issues/251))
* `data.azuread_users` - use `equals` instead of `startsWith` when looking uo users by `mailNickname` ([#251](https://github.com/hashicorp/terraform-provider-azuread/issues/251))

## 0.8.0 (March 16, 2020)

FEATURES:

* **New Data Source:** `azuread_client_config` ([#229](https://github.com/hashicorp/terraform-provider-azuread/issues/229))

IMPROVEMENTS:

* dependencies: upgrade `azure-sdk-for-go` to `v40.3.0` ([#225](https://github.com/hashicorp/terraform-provider-azuread/issues/225))
* dependencies: upgrade `go-autorest/autorest` to `v0.10.0` ([#225](https://github.com/hashicorp/terraform-provider-azuread/issues/225))
* dependencies: upgrade `terraform-plugin-sdk` to `v1.6.0` ([#225](https://github.com/hashicorp/terraform-provider-azuread/issues/225))
* `azuread_application` - support for the `logout_url` property ([#226](https://github.com/hashicorp/terraform-provider-azuread/issues/226))
* `azuread_group` - support for the `description` property ([#216](https://github.com/hashicorp/terraform-provider-azuread/issues/216))
* `azuread_user` - support for the `onpremises_sam_account_name` and `onpremises_user_principal_name` properties ([#222](https://github.com/hashicorp/terraform-provider-azuread/issues/222))
* `azuread_user` - support for the `immutable_id` property ([#207](https://github.com/hashicorp/terraform-provider-azuread/issues/207))

BUG FIXES:

* `azuread_application` - ensure all owners are added before removed ([#226](https://github.com/hashicorp/terraform-provider-azuread/issues/226))
* `azuread_application_password` - validate the `length` property is less then `863` ([#228](https://github.com/hashicorp/terraform-provider-azuread/issues/228))
* `azuread_group` - the `owners` property is now additive during creation allowing an existing owner to be provided ([#211](https://github.com/hashicorp/terraform-provider-azuread/issues/211))
* `azuread_group_member` - mark as missing when member cannot be found instead of erroring ([#227](https://github.com/hashicorp/terraform-provider-azuread/issues/227))
* `azuread_service_principal_password` - validate the `length` property is less then `863` ([#228](https://github.com/hashicorp/terraform-provider-azuread/issues/228))

## 0.7.0 (November 15, 2019)

IMPROVEMENTS:

* provider: migrate to standalone plugin SDK v1.1.0 ([#154](https://github.com/hashicorp/terraform-provider-azuread/issues/154))
* provider: using the current (rather than the vendored) version of Terraform Core in user agents ([#154](https://github.com/hashicorp/terraform-provider-azuread/issues/154))
* `azuread_application` - adds ability to build homepage with HTTP in addition to HTTPS ([#155](https://github.com/hashicorp/terraform-provider-azuread/issues/155))
* `azuread_application` - allow the `app_role` block `value` property to be nil ([#157](https://github.com/hashicorp/terraform-provider-azuread/issues/157))
* `azuread_user` - support for the `usage_location` property ([#141](https://github.com/hashicorp/terraform-provider-azuread/issues/141))
* `data.azuread_user` - support looking up a user with `mail_nickname` ([#161](https://github.com/hashicorp/terraform-provider-azuread/issues/161))
* `data.azuread_users` - support looking up users with `mail_nicknames` ([#161](https://github.com/hashicorp/terraform-provider-azuread/issues/161))

## 0.6.0 (August 21, 2019)

IMPROVEMENTS:

* dependencies: upgrading `github.com/Azure/azure-sdk-for-go` to `v32.5.0` ([#140](https://github.com/hashicorp/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/Azure/go-autorest` to `v13.0.0` ([#140](https://github.com/hashicorp/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/hashicorp/go-azure-helpers` to `v0.7.0` ([#140](https://github.com/hashicorp/terraform-provider-azuread/issues/140))
* dependencies: upgrading `github.com/hashicorp/terraform` to `0.12.6` ([#133](https://github.com/hashicorp/terraform-provider-azuread/issues/133))
* `azuread_service_principal` - support for the `app_role_assignment_required` property ([#127](https://github.com/hashicorp/terraform-provider-azuread/issues/127))


## 0.5.1 (July 24, 2019)

BUG FIXES:

* `azuread_application_password` - fix incorrect conflicts with ([#129](https://github.com/hashicorp/terraform-provider-azuread/issues/129))

## 0.5.0 (July 24, 2019)

FEATURES:

* **New Data Source:** `azuread_users` ([#109](https://github.com/hashicorp/terraform-provider-azuread/issues/109))
* **New Resource:** `azuread_group_member` ([#100](https://github.com/hashicorp/terraform-provider-azuread/issues/100))

IMPROVEMENTS:

* `azuread_application` - support for the `app_roles` property ([#98](https://github.com/hashicorp/terraform-provider-azuread/issues/98))
* `azuread_application` - the `identifier_uris` property now allows `api`,`urn`, and `ms-appx` URI schemas ([#115](https://github.com/hashicorp/terraform-provider-azuread/issues/115))
* `azuread_application_password` - deprecation of `application_id` in favour of `application_object_id` ([#107](https://github.com/hashicorp/terraform-provider-azuread/issues/107))
* `azuread_group` - support for the `members` property ([#100](https://github.com/hashicorp/terraform-provider-azuread/issues/100))
* `azuread_group` - support for the `owners` property ([#62](https://github.com/hashicorp/terraform-provider-azuread/issues/62))
* `azuread_service_principal` - export the `oauth2_permissions` property ([#103](https://github.com/hashicorp/terraform-provider-azuread/issues/103))
* `data.azuread_application` - support for the `app_roles` property ([#110](https://github.com/hashicorp/terraform-provider-azuread/issues/110))
* `data.azuread_service_principal` - export the `app_roles` property ([#110](https://github.com/hashicorp/terraform-provider-azuread/issues/110))

BUG FIXES:

* `azuread_application_password` - will now wait for replication on resource creation ([#118](https://github.com/hashicorp/terraform-provider-azuread/issues/118))
* `azuread_service_principal_password` - will now wait for replication on resource creation ([#117](https://github.com/hashicorp/terraform-provider-azuread/issues/117))

## 0.4.0 (June 06, 2019)

NOTES:

* Resource creation potentially could take longer after this release as the provider will now attempt to wait for replication like the az cli tool. 

FEATURES:

* **New Resource:** `azuread_application_password` ([#71](https://github.com/hashicorp/terraform-provider-azuread/issues/71))

IMPROVEMENTS:

* dependencies: upgrading to `v0.12.0` of `github.com/hashicorp/terraform` ([#82](https://github.com/hashicorp/terraform-provider-azuread/issues/82))
* `azuread_application` - support for the `group_membership_claims` property ([#78](https://github.com/hashicorp/terraform-provider-azuread/issues/78))
* `azuread_application` - now exports the `oauth2_permissions` property ([#79](https://github.com/hashicorp/terraform-provider-azuread/issues/79))
* `azuread_application` - now exports the `object_id` property ([#99](https://github.com/hashicorp/terraform-provider-azuread/issues/99))
* `azuread_application` - support for the `type` property enabling the creation of `native` applications ([#74](https://github.com/hashicorp/terraform-provider-azuread/issues/74))
* `azuread_application` - will now wait for replication by waiting for 10 successful reads after creation ([#93](https://github.com/hashicorp/terraform-provider-azuread/issues/93))
* `azuread_group` - will now wait for replication by waiting for 10 successful reads after creation ([#91](https://github.com/hashicorp/terraform-provider-azuread/issues/91))
* `azuread_group` - now exports the `object_id` property ([#99](https://github.com/hashicorp/terraform-provider-azuread/issues/99))
* `azuread_service_principal` - will now wait for replication by waiting for 10 successful reads after creation ([#93](https://github.com/hashicorp/terraform-provider-azuread/issues/93))
* `azuread_service_principal` - now exports the `object_id` property ([#99](https://github.com/hashicorp/terraform-provider-azuread/issues/99))
* `azuread_user` - will now wait for replication by waiting for 10 successful reads after creation ([#91](https://github.com/hashicorp/terraform-provider-azuread/issues/91))
* `azuread_user` - increase the maximum allowed length of `password` to 256 ([#81](https://github.com/hashicorp/terraform-provider-azuread/issues/81))
* `azuread_user` - now exports the `object_id` property ([#99](https://github.com/hashicorp/terraform-provider-azuread/issues/99))
* `data.azuread_application` - now exports the `group_membership_claims` property ([#78](https://github.com/hashicorp/terraform-provider-azuread/issues/78))
* `data.azuread_application` - now exports the `oauth2_permissions` property ([#79](https://github.com/hashicorp/terraform-provider-azuread/issues/79))

## 0.3.1 (April 18, 2019)

BUG FIXES:

* Release fixing metadata to register the provider as compatible with Terraform 0.12.

## 0.3.0 (April 18, 2019)

NOTES:

* This release includes a Terraform SDK upgrade with compatibility for Terraform v0.12. The provider remains backwards compatible with Terraform v0.11 and there should not be any significant behavioural changes. ([#56](https://github.com/hashicorp/terraform-provider-azuread/issues/56))

BUG FIXES:

* `azuread_application` - the order of the `reply_urls` property no longer matters ([#61](https://github.com/hashicorp/terraform-provider-azuread/issues/61))

## 0.2.0 (March 12, 2019)

FEATURES:

* **New Data Source:** `azuread_domains` ([#27](https://github.com/hashicorp/terraform-provider-azuread/issues/27))
* **New Data Source:** `azuread_group` ([#14](https://github.com/hashicorp/terraform-provider-azuread/issues/14))
* **New Resource:** `azuread_group` ([#14](https://github.com/hashicorp/terraform-provider-azuread/issues/14))

IMPROVEMENTS:

* dependencies: switching to use Go Modules ([#26](https://github.com/hashicorp/terraform-provider-azuread/issues/26))
* dependencies: updating `github.com/Azure/azure-sdk-for-go` to v24.1.0 ([#25](https://github.com/hashicorp/terraform-provider-azuread/issues/25))
* dependencies: updating `github.com/Azure/go-autorest` to v11.2.8 ([#24](https://github.com/hashicorp/terraform-provider-azuread/issues/24))
* validation: adding validation to all fields ([#30](https://github.com/hashicorp/terraform-provider-azuread/issues/30))
* `azuread_application` - support for `required_resource_access` property ([#23](https://github.com/hashicorp/terraform-provider-azuread/issues/23))
* `azuread_service_principal` - support for the `tags` property ([#31](https://github.com/hashicorp/terraform-provider-azuread/issues/31))
* `azuread_service_principal_password` - support for realitive ends dates with the `end_date_relative` property ([#53](https://github.com/hashicorp/terraform-provider-azuread/issues/53))

BUG FIXES:

* `azuread_application` - correctly reading back the `reply_urls` property into state ([#21](https://github.com/hashicorp/terraform-provider-azuread/issues/21))


## 0.1.0 (January 09, 2019)

Initial release of the Azure Active Directory provider - featuring resources split out from the AzureRM Provider.

FEATURES:

* New Data Source: `azuread_application`
* New Data Source: `azuread_service_principal`
* New Resource: `azuread_application`
* New Resource: `azuread_service_principal`
* New Resource: `azuread_service_principal_password`
