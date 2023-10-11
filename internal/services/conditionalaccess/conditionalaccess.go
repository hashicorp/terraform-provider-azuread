// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
)

func flattenConditionalAccessConditionSet(in *msgraph.ConditionalAccessConditionSet) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"applications":                  flattenConditionalAccessApplications(in.Applications),
			"client_applications":           flattenConditionalAccessClientApplications(in.ClientApplications),
			"users":                         flattenConditionalAccessUsers(in.Users),
			"client_app_types":              tf.FlattenStringSlicePtr(in.ClientAppTypes),
			"devices":                       flattenConditionalAccessDevices(in.Devices),
			"locations":                     flattenConditionalAccessLocations(in.Locations),
			"platforms":                     flattenConditionalAccessPlatforms(in.Platforms),
			"service_principal_risk_levels": tf.FlattenStringSlicePtr(in.ServicePrincipalRiskLevels),
			"sign_in_risk_levels":           tf.FlattenStringSlicePtr(in.SignInRiskLevels),
			"user_risk_levels":              tf.FlattenStringSlicePtr(in.UserRiskLevels),
		},
	}
}

func flattenConditionalAccessApplications(in *msgraph.ConditionalAccessApplications) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_applications": tf.FlattenStringSlicePtr(in.IncludeApplications),
			"excluded_applications": tf.FlattenStringSlicePtr(in.ExcludeApplications),
			"included_user_actions": tf.FlattenStringSlicePtr(in.IncludeUserActions),
		},
	}
}

func flattenConditionalAccessClientApplications(in *msgraph.ConditionalAccessClientApplications) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_service_principals": tf.FlattenStringSlicePtr(in.IncludeServicePrincipals),
			"excluded_service_principals": tf.FlattenStringSlicePtr(in.ExcludeServicePrincipals),
		},
	}
}

func flattenConditionalAccessUsers(in *msgraph.ConditionalAccessUsers) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_users":                    tf.FlattenStringSlicePtr(in.IncludeUsers),
			"excluded_users":                    tf.FlattenStringSlicePtr(in.ExcludeUsers),
			"included_groups":                   tf.FlattenStringSlicePtr(in.IncludeGroups),
			"excluded_groups":                   tf.FlattenStringSlicePtr(in.ExcludeGroups),
			"included_roles":                    tf.FlattenStringSlicePtr(in.IncludeRoles),
			"excluded_roles":                    tf.FlattenStringSlicePtr(in.ExcludeRoles),
			"included_guests_or_external_users": flattenGuestsOrExternalUsers(in.IncludeGuestsOrExternalUsers),
			"excluded_guests_or_external_users": flattenGuestsOrExternalUsers(in.ExcludeGuestsOrExternalUsers),
		},
	}
}

func flattenConditionalAccessDevices(in *msgraph.ConditionalAccessDevices) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"filter": flattenConditionalAccessDeviceFilter(in.DeviceFilter),
		},
	}
}

func flattenConditionalAccessLocations(in *msgraph.ConditionalAccessLocations) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_locations": tf.FlattenStringSlicePtr(in.IncludeLocations),
			"excluded_locations": tf.FlattenStringSlicePtr(in.ExcludeLocations),
		},
	}
}

func flattenConditionalAccessPlatforms(in *msgraph.ConditionalAccessPlatforms) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_platforms": tf.FlattenStringSlicePtr(in.IncludePlatforms),
			"excluded_platforms": tf.FlattenStringSlicePtr(in.ExcludePlatforms),
		},
	}
}

func flattenConditionalAccessGrantControls(in *msgraph.ConditionalAccessGrantControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	var authenticationStrengthPolicyId string
	if in.AuthenticationStrength != nil {
		authenticationStrengthPolicyId = pointer.From(in.AuthenticationStrength.ID)
	}

	return []interface{}{
		map[string]interface{}{
			"operator":                          in.Operator,
			"built_in_controls":                 tf.FlattenStringSlicePtr(in.BuiltInControls),
			"authentication_strength_policy_id": authenticationStrengthPolicyId,
			"custom_authentication_factors":     tf.FlattenStringSlicePtr(in.CustomAuthenticationFactors),
			"terms_of_use":                      tf.FlattenStringSlicePtr(in.TermsOfUse),
		},
	}
}

func flattenConditionalAccessSessionControls(in *msgraph.ConditionalAccessSessionControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	applicationEnforceRestrictions := false
	if in.ApplicationEnforcedRestrictions != nil {
		applicationEnforceRestrictions = pointer.From(in.ApplicationEnforcedRestrictions.IsEnabled)
	}

	cloudAppSecurity := ""
	if in.CloudAppSecurity != nil {
		cloudAppSecurity = pointer.From(in.CloudAppSecurity.CloudAppSecurityType)
	}

	disableResilienceDefaults := false
	if in.DisableResilienceDefaults != nil {
		disableResilienceDefaults = *in.DisableResilienceDefaults
	}

	signInFrequency := 0
	signInFrequencyAuthenticationType := ""
	signInFrequencyInterval := ""
	signInFrequencyPeriod := ""
	if in.SignInFrequency != nil {
		if in.SignInFrequency.Value != nil && in.SignInFrequency.Type != nil {
			signInFrequency = int(*in.SignInFrequency.Value)
			signInFrequencyPeriod = *in.SignInFrequency.Type
		}

		signInFrequencyAuthenticationType = pointer.From(in.SignInFrequency.AuthenticationType)
		signInFrequencyInterval = pointer.From(in.SignInFrequency.FrequencyInterval)
	}

	persistentBrowserMode := ""
	if in.PersistentBrowser != nil {
		persistentBrowserMode = pointer.From(in.PersistentBrowser.Mode)
	}

	return []interface{}{
		map[string]interface{}{
			"application_enforced_restrictions_enabled": applicationEnforceRestrictions,
			"cloud_app_security_policy":                 cloudAppSecurity,
			"disable_resilience_defaults":               disableResilienceDefaults,
			"persistent_browser_mode":                   persistentBrowserMode,
			"sign_in_frequency":                         signInFrequency,
			"sign_in_frequency_authentication_type":     signInFrequencyAuthenticationType,
			"sign_in_frequency_interval":                signInFrequencyInterval,
			"sign_in_frequency_period":                  signInFrequencyPeriod,
		},
	}
}

func flattenConditionalAccessDeviceFilter(in *msgraph.ConditionalAccessFilter) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"mode": in.Mode,
			"rule": in.Rule,
		},
	}
}

func flattenGuestsOrExternalUsers(in *msgraph.ConditionalAccessGuestsOrExternalUsers) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"guest_or_external_user_types": tf.FlattenStringSlicePtr(in.GuestOrExternalUserTypes),
			"external_tenants":             flattenExternalTenants(in.ExternalTenants),
		},
	}
}

func flattenExternalTenants(in *msgraph.ConditionalAccessExternalTenants) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"membership_kind": in.MembershipKind,
			"members":         tf.FlattenStringSlicePtr(in.Members),
		},
	}
}

func flattenCountryNamedLocation(in *msgraph.CountryNamedLocation) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	includeUnknown := false
	if in.IncludeUnknownCountriesAndRegions != nil {
		includeUnknown = *in.IncludeUnknownCountriesAndRegions
	}

	return []interface{}{
		map[string]interface{}{
			"countries_and_regions":                 tf.FlattenStringSlicePtr(in.CountriesAndRegions),
			"include_unknown_countries_and_regions": includeUnknown,
		},
	}
}

func flattenIPNamedLocation(in *msgraph.IPNamedLocation) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	trusted := false
	if in.IsTrusted != nil {
		trusted = *in.IsTrusted
	}

	return []interface{}{
		map[string]interface{}{
			"ip_ranges": flattenIPNamedLocationIPRange(in.IPRanges),
			"trusted":   trusted,
		},
	}
}

func flattenIPNamedLocationIPRange(in *[]msgraph.IPNamedLocationIPRange) []interface{} {
	if in == nil || len(*in) == 0 {
		return []interface{}{}
	}

	result := make([]string, 0)
	for _, cidr := range *in {
		if cidr.CIDRAddress != nil {
			result = append(result, *cidr.CIDRAddress)
		}
	}

	return tf.FlattenStringSlice(result)
}

func expandConditionalAccessConditionSet(in []interface{}) *msgraph.ConditionalAccessConditionSet {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessConditionSet{}
	config := in[0].(map[string]interface{})

	applications := config["applications"].([]interface{})
	users := config["users"].([]interface{})
	clientAppTypes := config["client_app_types"].([]interface{})
	devices := config["devices"].([]interface{})
	locations := config["locations"].([]interface{})
	platforms := config["platforms"].([]interface{})
	servicePrincipalRiskLevels := config["service_principal_risk_levels"].([]interface{})
	signInRiskLevels := config["sign_in_risk_levels"].([]interface{})
	userRiskLevels := config["user_risk_levels"].([]interface{})
	clientApplications := config["client_applications"].([]interface{})

	result.Applications = expandConditionalAccessApplications(applications)
	result.Users = expandConditionalAccessUsers(users)
	result.ClientAppTypes = tf.ExpandStringSlicePtr(clientAppTypes)
	result.Devices = expandConditionalAccessDevices(devices)
	result.Locations = expandConditionalAccessLocations(locations)
	result.Platforms = expandConditionalAccessPlatforms(platforms)
	result.ServicePrincipalRiskLevels = tf.ExpandStringSlicePtr(servicePrincipalRiskLevels)
	result.SignInRiskLevels = tf.ExpandStringSlicePtr(signInRiskLevels)
	result.UserRiskLevels = tf.ExpandStringSlicePtr(userRiskLevels)
	result.ClientApplications = expandConditionalAccessClientApplications(clientApplications)

	return &result
}

func expandConditionalAccessClientApplications(in []interface{}) *msgraph.ConditionalAccessClientApplications {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessClientApplications{}
	config := in[0].(map[string]interface{})

	includeServicePrincipals := config["included_service_principals"].([]interface{})
	excludeServicePrincipals := config["excluded_service_principals"].([]interface{})

	result.IncludeServicePrincipals = tf.ExpandStringSlicePtr(includeServicePrincipals)
	result.ExcludeServicePrincipals = tf.ExpandStringSlicePtr(excludeServicePrincipals)

	return &result
}

func expandConditionalAccessApplications(in []interface{}) *msgraph.ConditionalAccessApplications {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessApplications{}
	config := in[0].(map[string]interface{})

	includeApplications := config["included_applications"].([]interface{})
	excludeApplications := config["excluded_applications"].([]interface{})
	includeUserActions := config["included_user_actions"].([]interface{})

	result.IncludeApplications = tf.ExpandStringSlicePtr(includeApplications)
	result.ExcludeApplications = tf.ExpandStringSlicePtr(excludeApplications)
	result.IncludeUserActions = tf.ExpandStringSlicePtr(includeUserActions)

	return &result
}

func expandConditionalAccessUsers(in []interface{}) *msgraph.ConditionalAccessUsers {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessUsers{}
	config := in[0].(map[string]interface{})

	includeUsers := config["included_users"].([]interface{})
	excludeUsers := config["excluded_users"].([]interface{})
	includeGroups := config["included_groups"].([]interface{})
	excludeGroups := config["excluded_groups"].([]interface{})
	includeRoles := config["included_roles"].([]interface{})
	excludeRoles := config["excluded_roles"].([]interface{})

	result.IncludeUsers = tf.ExpandStringSlicePtr(includeUsers)
	result.ExcludeUsers = tf.ExpandStringSlicePtr(excludeUsers)
	result.IncludeGroups = tf.ExpandStringSlicePtr(includeGroups)
	result.ExcludeGroups = tf.ExpandStringSlicePtr(excludeGroups)
	result.IncludeRoles = tf.ExpandStringSlicePtr(includeRoles)
	result.ExcludeRoles = tf.ExpandStringSlicePtr(excludeRoles)

	result.IncludeGuestsOrExternalUsers = expandGuestsOrExternalUsers(config["included_guests_or_external_users"].([]interface{}))
	result.ExcludeGuestsOrExternalUsers = expandGuestsOrExternalUsers(config["excluded_guests_or_external_users"].([]interface{}))

	return &result
}

func expandConditionalAccessPlatforms(in []interface{}) *msgraph.ConditionalAccessPlatforms {
	result := msgraph.ConditionalAccessPlatforms{}
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	config := in[0].(map[string]interface{})

	includePlatforms := config["included_platforms"].([]interface{})
	excludePlatforms := config["excluded_platforms"].([]interface{})

	result.IncludePlatforms = tf.ExpandStringSlicePtr(includePlatforms)
	result.ExcludePlatforms = tf.ExpandStringSlicePtr(excludePlatforms)

	return &result
}

func expandConditionalAccessDevices(in []interface{}) *msgraph.ConditionalAccessDevices {
	result := msgraph.ConditionalAccessDevices{}

	if len(in) == 0 || in[0] == nil {
		// The devices field cannot be empty on POST, and is currently totally ignored when empty on PATCH,
		// so for now we'll just return nil here and revisit later.
		return nil
	}

	config := in[0].(map[string]interface{})

	deviceFilter := config["filter"].([]interface{})

	if len(deviceFilter) > 0 {
		result.DeviceFilter = expandConditionalAccessFilter(deviceFilter)
	}

	return &result
}

func expandConditionalAccessLocations(in []interface{}) *msgraph.ConditionalAccessLocations {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessLocations{}
	config := in[0].(map[string]interface{})

	includeLocations := config["included_locations"].([]interface{})
	excludeLocations := config["excluded_locations"].([]interface{})

	result.IncludeLocations = tf.ExpandStringSlicePtr(includeLocations)
	result.ExcludeLocations = tf.ExpandStringSlicePtr(excludeLocations)

	return &result
}

func expandConditionalAccessGrantControls(in []interface{}) *msgraph.ConditionalAccessGrantControls {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessGrantControls{}
	config := in[0].(map[string]interface{})

	operator := config["operator"].(string)
	authenticationStrengthId := config["authentication_strength_policy_id"].(string)
	builtInControls := config["built_in_controls"].([]interface{})
	customAuthenticationFactors := config["custom_authentication_factors"].([]interface{})
	termsOfUse := config["terms_of_use"].([]interface{})

	result.Operator = &operator

	if authenticationStrengthId != "" {
		result.AuthenticationStrength = &msgraph.AuthenticationStrengthPolicy{
			ID: &authenticationStrengthId,
		}
	}

	result.BuiltInControls = tf.ExpandStringSlicePtr(builtInControls)
	result.CustomAuthenticationFactors = tf.ExpandStringSlicePtr(customAuthenticationFactors)
	result.TermsOfUse = tf.ExpandStringSlicePtr(termsOfUse)

	return &result
}

func expandConditionalAccessSessionControls(in []interface{}) *msgraph.ConditionalAccessSessionControls {
	result := msgraph.ConditionalAccessSessionControls{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	config := in[0].(map[string]interface{})

	result.ApplicationEnforcedRestrictions = &msgraph.ApplicationEnforcedRestrictionsSessionControl{
		IsEnabled: pointer.To(config["application_enforced_restrictions_enabled"].(bool)),
	}

	if cloudAppSecurity := config["cloud_app_security_policy"].(string); cloudAppSecurity != "" {
		result.CloudAppSecurity = &msgraph.CloudAppSecurityControl{
			IsEnabled:            pointer.To(true),
			CloudAppSecurityType: pointer.To(cloudAppSecurity),
		}
	}

	DisableResilienceDefaults := config["disable_resilience_defaults"]
	result.DisableResilienceDefaults = pointer.To(DisableResilienceDefaults.(bool))

	if persistentBrowserMode := config["persistent_browser_mode"].(string); persistentBrowserMode != "" {
		result.PersistentBrowser = &msgraph.PersistentBrowserSessionControl{
			IsEnabled: pointer.To(true),
			Mode:      pointer.To(persistentBrowserMode),
		}
	}

	signInFrequency := msgraph.SignInFrequencySessionControl{}
	if frequencyValue := config["sign_in_frequency"].(int); frequencyValue > 0 {
		signInFrequency.IsEnabled = pointer.To(true)
		signInFrequency.Type = pointer.To(config["sign_in_frequency_period"].(string))
		signInFrequency.Value = pointer.To(int32(frequencyValue))
	}

	if authenticationType, ok := config["sign_in_frequency_authentication_type"]; ok {
		signInFrequency.AuthenticationType = pointer.To(authenticationType.(string))
	}

	if interval, ok := config["sign_in_frequency_interval"]; ok {
		signInFrequency.FrequencyInterval = pointer.To(interval.(string))
	}

	// API returns 400 error if signInFrequency is set with all default/zero values
	if pointer.From(signInFrequency.IsEnabled) || pointer.From(signInFrequency.FrequencyInterval) != msgraph.ConditionalAccessFrequencyIntervalTimeBased ||
		pointer.From(signInFrequency.AuthenticationType) != msgraph.ConditionalAccessAuthenticationTypePrimaryAndSecondaryAuthentication {
		result.SignInFrequency = &signInFrequency
	}

	// API does not accept ineffectual and sessionControls object, and it will not remove any existing sessionControls unless the entire object is set to null
	if (result.ApplicationEnforcedRestrictions == nil || !pointer.From(result.ApplicationEnforcedRestrictions.IsEnabled)) &&
		result.CloudAppSecurity == nil && !pointer.From(result.DisableResilienceDefaults) &&
		result.PersistentBrowser == nil && result.SignInFrequency == nil {
		return nil
	}

	return &result
}

func expandConditionalAccessFilter(in []interface{}) *msgraph.ConditionalAccessFilter {
	result := msgraph.ConditionalAccessFilter{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	config := in[0].(map[string]interface{})

	result.Mode = pointer.To(config["mode"].(string))
	result.Rule = pointer.To(config["rule"].(string))

	return &result
}

func expandGuestsOrExternalUsers(in []interface{}) *msgraph.ConditionalAccessGuestsOrExternalUsers {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessGuestsOrExternalUsers{}

	config := in[0].(map[string]interface{})

	types := config["guest_or_external_user_types"].([]interface{})

	result.GuestOrExternalUserTypes = tf.ExpandStringSlicePtr(types)
	result.ExternalTenants = expandExternalTenants(config["external_tenants"].([]interface{}))

	return &result
}

func expandExternalTenants(in []interface{}) *msgraph.ConditionalAccessExternalTenants {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.ConditionalAccessExternalTenants{}

	config := in[0].(map[string]interface{})

	members := config["members"].([]interface{})

	result.MembershipKind = pointer.To(config["membership_kind"].(string))

	// only membership_kind enumerated is allowed to have members field set
	// so we omit setting an empty array when no members configured
	if len(members) > 0 {
		result.Members = tf.ExpandStringSlicePtr(members)
	}

	return &result
}

func expandCountryNamedLocation(in []interface{}) *msgraph.CountryNamedLocation {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.CountryNamedLocation{}
	config := in[0].(map[string]interface{})

	countriesAndRegions := config["countries_and_regions"].([]interface{})
	includeUnknown := config["include_unknown_countries_and_regions"]

	result.CountriesAndRegions = tf.ExpandStringSlicePtr(countriesAndRegions)
	result.IncludeUnknownCountriesAndRegions = pointer.To(includeUnknown.(bool))

	return &result
}

func expandIPNamedLocation(in []interface{}) *msgraph.IPNamedLocation {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := msgraph.IPNamedLocation{}
	config := in[0].(map[string]interface{})

	ipRanges := config["ip_ranges"].([]interface{})
	trusted := config["trusted"]

	result.IPRanges = expandIPNamedLocationIPRange(ipRanges)
	result.IsTrusted = pointer.To(trusted.(bool))

	return &result
}

func expandIPNamedLocationIPRange(in []interface{}) *[]msgraph.IPNamedLocationIPRange {
	if len(in) == 0 {
		return nil
	}

	result := make([]msgraph.IPNamedLocationIPRange, 0)
	for _, cidr := range in {
		result = append(result, msgraph.IPNamedLocationIPRange{
			CIDRAddress: pointer.To(cidr.(string)),
		})
	}

	return &result
}
