// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
)

func flattenConditionalAccessConditionSet(in *stable.ConditionalAccessConditionSet) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	clientAppTypes := make([]string, 0)
	for _, v := range in.ClientAppTypes {
		clientAppTypes = append(clientAppTypes, string(v))
	}

	servicePrincipalRiskLevels := make([]string, 0)
	for _, v := range pointer.From(in.ServicePrincipalRiskLevels) {
		servicePrincipalRiskLevels = append(servicePrincipalRiskLevels, string(v))
	}

	signInRiskLevels := make([]string, 0)
	for _, v := range in.SignInRiskLevels {
		signInRiskLevels = append(signInRiskLevels, string(v))
	}

	userRiskLevels := make([]string, 0)
	for _, v := range in.UserRiskLevels {
		userRiskLevels = append(userRiskLevels, string(v))
	}

	insiderRiskLevels := ""
	if in.InsiderRiskLevels != nil {
		insiderRiskLevels = string(pointer.From(in.InsiderRiskLevels))
	}

	return []interface{}{
		map[string]interface{}{
			"applications":                  flattenConditionalAccessApplications(in.Applications),
			"client_applications":           flattenConditionalAccessClientApplications(in.ClientApplications),
			"users":                         flattenConditionalAccessUsers(in.Users),
			"client_app_types":              clientAppTypes,
			"devices":                       flattenConditionalAccessDevices(in.Devices),
			"locations":                     flattenConditionalAccessLocations(in.Locations),
			"platforms":                     flattenConditionalAccessPlatforms(in.Platforms),
			"service_principal_risk_levels": servicePrincipalRiskLevels,
			"sign_in_risk_levels":           signInRiskLevels,
			"user_risk_levels":              userRiskLevels,
			"insider_risk_levels":           insiderRiskLevels,
		},
	}
}

func flattenConditionalAccessApplications(in stable.ConditionalAccessApplications) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"included_applications": tf.FlattenStringSlicePtr(in.IncludeApplications),
			"excluded_applications": tf.FlattenStringSlicePtr(in.ExcludeApplications),
			"included_user_actions": tf.FlattenStringSlicePtr(in.IncludeUserActions),
		},
	}
}

func flattenConditionalAccessClientApplications(in *stable.ConditionalAccessClientApplications) []interface{} {
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

func flattenConditionalAccessUsers(in *stable.ConditionalAccessUsers) []interface{} {
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

func flattenConditionalAccessDevices(in *stable.ConditionalAccessDevices) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"filter": flattenConditionalAccessDeviceFilter(in.DeviceFilter),
		},
	}
}

func flattenConditionalAccessLocations(in *stable.ConditionalAccessLocations) []interface{} {
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

func flattenConditionalAccessPlatforms(in *stable.ConditionalAccessPlatforms) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	includePlatforms := make([]string, 0)
	for _, v := range pointer.From(in.IncludePlatforms) {
		includePlatforms = append(includePlatforms, string(v))
	}

	excludePlatforms := make([]string, 0)
	for _, v := range pointer.From(in.ExcludePlatforms) {
		excludePlatforms = append(excludePlatforms, string(v))
	}

	return []interface{}{
		map[string]interface{}{
			"included_platforms": includePlatforms,
			"excluded_platforms": excludePlatforms,
		},
	}
}

func flattenConditionalAccessGrantControls(in *stable.ConditionalAccessGrantControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	var authenticationStrengthPolicyId string
	if in.AuthenticationStrength != nil {
		authenticationStrengthPolicyId = stable.NewPolicyAuthenticationStrengthPolicyID(pointer.From(in.AuthenticationStrength.Id)).ID()
	}

	builtInControls := make([]string, 0)
	for _, v := range pointer.From(in.BuiltInControls) {
		builtInControls = append(builtInControls, string(v))
	}

	return []interface{}{
		map[string]interface{}{
			"operator":                          in.Operator.GetOrZero(),
			"built_in_controls":                 builtInControls,
			"authentication_strength_policy_id": authenticationStrengthPolicyId,
			"custom_authentication_factors":     tf.FlattenStringSlicePtr(in.CustomAuthenticationFactors),
			"terms_of_use":                      tf.FlattenStringSlicePtr(in.TermsOfUse),
		},
	}
}

func flattenConditionalAccessSessionControls(in *stable.ConditionalAccessSessionControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	applicationEnforceRestrictions := false
	if in.ApplicationEnforcedRestrictions != nil {
		applicationEnforceRestrictions = in.ApplicationEnforcedRestrictions.IsEnabled.GetOrZero()
	}

	cloudAppSecurity := ""
	if in.CloudAppSecurity != nil {
		cloudAppSecurity = string(pointer.From(in.CloudAppSecurity.CloudAppSecurityType))
	}

	disableResilienceDefaults := false
	if in.DisableResilienceDefaults != nil {
		disableResilienceDefaults = in.DisableResilienceDefaults.GetOrZero()
	}

	signInFrequency := 0
	signInFrequencyAuthenticationType := ""
	signInFrequencyInterval := ""
	signInFrequencyPeriod := ""
	if in.SignInFrequency != nil {
		if !in.SignInFrequency.Value.IsNull() && in.SignInFrequency.Type != nil {
			signInFrequency = int(in.SignInFrequency.Value.GetOrZero())
			signInFrequencyPeriod = string(*in.SignInFrequency.Type)
		}

		signInFrequencyAuthenticationType = string(pointer.From(in.SignInFrequency.AuthenticationType))
		signInFrequencyInterval = string(pointer.From(in.SignInFrequency.FrequencyInterval))
	}

	persistentBrowserMode := ""
	if in.PersistentBrowser != nil {
		persistentBrowserMode = string(pointer.From(in.PersistentBrowser.Mode))
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

func flattenConditionalAccessDeviceFilter(in *stable.ConditionalAccessFilter) []interface{} {
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

func flattenGuestsOrExternalUsers(in *stable.ConditionalAccessGuestsOrExternalUsers) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	guestOrExternalUserTypes := make([]string, 0)
	for _, v := range strings.Split(string(pointer.From(in.GuestOrExternalUserTypes)), ",") {
		guestOrExternalUserTypes = append(guestOrExternalUserTypes, strings.TrimSpace(v))
	}

	return []interface{}{
		map[string]interface{}{
			"guest_or_external_user_types": guestOrExternalUserTypes,
			"external_tenants":             flattenExternalTenants(in.ExternalTenants),
		},
	}
}

func flattenExternalTenants(in stable.ConditionalAccessExternalTenants) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	externalTenants := in.ConditionalAccessExternalTenants()

	return []interface{}{
		map[string]interface{}{
			"membership_kind": externalTenants.MembershipKind,
			"members":         tf.FlattenStringSlicePtr(externalTenants.Members),
		},
	}
}

func flattenCountryNamedLocation(in *stable.CountryNamedLocation) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	includeUnknown := false
	if in.IncludeUnknownCountriesAndRegions != nil {
		includeUnknown = *in.IncludeUnknownCountriesAndRegions
	}

	countryLookupMethod := stable.CountryLookupMethodType_ClientIPAddress
	if in.CountryLookupMethod != nil {
		countryLookupMethod = *in.CountryLookupMethod
	}

	return []interface{}{
		map[string]interface{}{
			"countries_and_regions":                 tf.FlattenStringSlice(in.CountriesAndRegions),
			"include_unknown_countries_and_regions": includeUnknown,
			"country_lookup_method":                 countryLookupMethod,
		},
	}
}

func flattenIPNamedLocation(in *stable.IPNamedLocation) []interface{} {
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

func flattenIPNamedLocationIPRange(in []stable.IPRange) []interface{} {
	if len(in) == 0 {
		return []interface{}{}
	}

	result := make([]string, 0)
	for _, i := range in {
		switch model := i.(type) {
		case stable.IPv4CIDRRange:
			if model.CIDRAddress != nil {
				result = append(result, *model.CIDRAddress)
			}
		case stable.IPv6CIDRRange:
			if model.CIDRAddress != nil {
				result = append(result, *model.CIDRAddress)
			}
		}
	}

	return tf.FlattenStringSlice(result)
}

func expandConditionalAccessConditionSet(in []interface{}) *stable.ConditionalAccessConditionSet {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessConditionSet{}
	config := in[0].(map[string]interface{})

	applications := config["applications"].([]interface{})
	clientApplications := config["client_applications"].([]interface{})
	devices := config["devices"].([]interface{})
	locations := config["locations"].([]interface{})
	platforms := config["platforms"].([]interface{})
	users := config["users"].([]interface{})

	clientAppTypes := make([]stable.ConditionalAccessClientApp, 0)
	for _, elem := range config["client_app_types"].([]interface{}) {
		clientAppTypes = append(clientAppTypes, stable.ConditionalAccessClientApp(elem.(string)))
	}

	servicePrincipalRiskLevels := make([]stable.RiskLevel, 0)
	for _, elem := range config["service_principal_risk_levels"].([]interface{}) {
		servicePrincipalRiskLevels = append(servicePrincipalRiskLevels, stable.RiskLevel(elem.(string)))
	}

	signInRiskLevels := make([]stable.RiskLevel, 0)
	for _, elem := range config["sign_in_risk_levels"].([]interface{}) {
		signInRiskLevels = append(signInRiskLevels, stable.RiskLevel(elem.(string)))
	}

	userRiskLevels := make([]stable.RiskLevel, 0)
	for _, elem := range config["user_risk_levels"].([]interface{}) {
		userRiskLevels = append(userRiskLevels, stable.RiskLevel(elem.(string)))
	}

	if insiderRiskLevel, ok := config["insider_risk_levels"]; ok && insiderRiskLevel.(string) != "" {
		result.InsiderRiskLevels = pointer.To(stable.ConditionalAccessInsiderRiskLevels(insiderRiskLevel.(string)))
	}

	result.Applications = expandConditionalAccessApplications(applications)
	result.ClientAppTypes = clientAppTypes
	result.ClientApplications = expandConditionalAccessClientApplications(clientApplications)
	result.Devices = expandConditionalAccessDevices(devices)
	result.Locations = expandConditionalAccessLocations(locations)
	result.Platforms = expandConditionalAccessPlatforms(platforms)
	result.ServicePrincipalRiskLevels = &servicePrincipalRiskLevels
	result.SignInRiskLevels = signInRiskLevels
	result.UserRiskLevels = userRiskLevels
	result.Users = expandConditionalAccessUsers(users)

	return &result
}

func expandConditionalAccessClientApplications(in []interface{}) *stable.ConditionalAccessClientApplications {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessClientApplications{}
	config := in[0].(map[string]interface{})

	includeServicePrincipals := config["included_service_principals"].([]interface{})
	excludeServicePrincipals := config["excluded_service_principals"].([]interface{})

	result.IncludeServicePrincipals = tf.ExpandStringSlicePtr(includeServicePrincipals)
	result.ExcludeServicePrincipals = tf.ExpandStringSlicePtr(excludeServicePrincipals)

	return &result
}

func expandConditionalAccessApplications(in []interface{}) stable.ConditionalAccessApplications {
	result := stable.ConditionalAccessApplications{}
	if len(in) == 0 || in[0] == nil {
		return result
	}

	config := in[0].(map[string]interface{})

	includeApplications := config["included_applications"].([]interface{})
	excludeApplications := config["excluded_applications"].([]interface{})
	includeUserActions := config["included_user_actions"].([]interface{})

	result.IncludeApplications = tf.ExpandStringSlicePtr(includeApplications)
	result.ExcludeApplications = tf.ExpandStringSlicePtr(excludeApplications)
	result.IncludeUserActions = tf.ExpandStringSlicePtr(includeUserActions)

	return result
}

func expandConditionalAccessUsers(in []interface{}) *stable.ConditionalAccessUsers {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessUsers{}
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

func expandConditionalAccessDevices(in []interface{}) *stable.ConditionalAccessDevices {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessDevices{}
	config := in[0].(map[string]interface{})

	deviceFilter := config["filter"].([]interface{})
	if len(deviceFilter) > 0 {
		result.DeviceFilter = expandConditionalAccessFilter(deviceFilter)
	}

	return &result
}

func expandConditionalAccessLocations(in []interface{}) *stable.ConditionalAccessLocations {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessLocations{}
	config := in[0].(map[string]interface{})

	includeLocations := config["included_locations"].([]interface{})
	excludeLocations := config["excluded_locations"].([]interface{})

	result.IncludeLocations = tf.ExpandStringSlicePtr(includeLocations)
	result.ExcludeLocations = tf.ExpandStringSlicePtr(excludeLocations)

	return &result
}

func expandConditionalAccessPlatforms(in []interface{}) *stable.ConditionalAccessPlatforms {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.ConditionalAccessPlatforms{}
	config := in[0].(map[string]interface{})

	includePlatforms := make([]stable.ConditionalAccessDevicePlatform, 0)
	for _, elem := range config["included_platforms"].([]interface{}) {
		includePlatforms = append(includePlatforms, stable.ConditionalAccessDevicePlatform(elem.(string)))
	}

	excludePlatforms := make([]stable.ConditionalAccessDevicePlatform, 0)
	for _, elem := range config["excluded_platforms"].([]interface{}) {
		excludePlatforms = append(excludePlatforms, stable.ConditionalAccessDevicePlatform(elem.(string)))
	}

	result.IncludePlatforms = &includePlatforms
	result.ExcludePlatforms = &excludePlatforms

	return &result
}

func expandConditionalAccessGrantControls(in []interface{}) (*stable.ConditionalAccessGrantControls, error) {
	if len(in) == 0 || in[0] == nil {
		return nil, nil
	}

	result := stable.ConditionalAccessGrantControls{}
	config := in[0].(map[string]interface{})

	if id := config["authentication_strength_policy_id"].(string); id != "" {
		policyId, err := stable.ParsePolicyAuthenticationStrengthPolicyID(id)
		if err != nil {
			return nil, err
		}
		result.AuthenticationStrength = &stable.AuthenticationStrengthPolicy{
			Id: pointer.To(policyId.AuthenticationStrengthPolicyId),
		}
	}

	builtInControls := make([]stable.ConditionalAccessGrantControl, 0)
	for _, elem := range config["built_in_controls"].([]interface{}) {
		builtInControls = append(builtInControls, stable.ConditionalAccessGrantControl(elem.(string)))
	}

	result.BuiltInControls = &builtInControls
	result.CustomAuthenticationFactors = tf.ExpandStringSlicePtr(config["custom_authentication_factors"].([]interface{}))
	result.Operator = nullable.Value(config["operator"].(string))
	result.TermsOfUse = tf.ExpandStringSlicePtr(config["terms_of_use"].([]interface{}))

	return &result, nil
}

func expandConditionalAccessSessionControls(in []interface{}) *stable.ConditionalAccessSessionControls {
	result := stable.ConditionalAccessSessionControls{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	config := in[0].(map[string]interface{})

	if cloudAppSecurity := config["cloud_app_security_policy"]; cloudAppSecurity.(string) != "" {
		result.CloudAppSecurity = &stable.CloudAppSecuritySessionControl{
			IsEnabled:            nullable.Value(true),
			CloudAppSecurityType: pointer.To(stable.CloudAppSecuritySessionControlType(cloudAppSecurity.(string))),
		}
	}

	if persistentBrowserMode := config["persistent_browser_mode"]; persistentBrowserMode.(string) != "" {
		result.PersistentBrowser = &stable.PersistentBrowserSessionControl{
			IsEnabled: nullable.Value(true),
			Mode:      pointer.To(stable.PersistentBrowserSessionMode(persistentBrowserMode.(string))),
		}
	}

	signInFrequency := stable.SignInFrequencySessionControl{}
	if frequencyValue := config["sign_in_frequency"].(int); frequencyValue > 0 {
		signInFrequency.IsEnabled = nullable.Value(true)
		signInFrequency.Type = pointer.To(stable.SigninFrequencyType(config["sign_in_frequency_period"].(string)))
		signInFrequency.Value = nullable.Value(int64(frequencyValue))

		signInFrequency.AuthenticationType = pointer.To(stable.SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication)
		signInFrequency.FrequencyInterval = pointer.To(stable.SignInFrequencyInterval_TimeBased)
	}

	if authenticationType, ok := config["sign_in_frequency_authentication_type"]; ok && authenticationType.(string) != "" {
		signInFrequency.AuthenticationType = pointer.To(stable.SignInFrequencyAuthenticationType(authenticationType.(string)))
	}

	if interval, ok := config["sign_in_frequency_interval"]; ok && interval.(string) != "" {
		signInFrequency.IsEnabled = nullable.Value(true)
		signInFrequency.AuthenticationType = pointer.To(stable.SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication)
		if authType := config["sign_in_frequency_authentication_type"].(string); authType != "" {
			signInFrequency.AuthenticationType = pointer.ToEnum[stable.SignInFrequencyAuthenticationType](authType)
		}
		signInFrequency.FrequencyInterval = pointer.To(stable.SignInFrequencyInterval(interval.(string)))
	}

	applicationEnforcedRestrictions := config["application_enforced_restrictions_enabled"].(bool)
	if pointer.From(signInFrequency.FrequencyInterval) != stable.SignInFrequencyInterval_EveryTime { // application enforced restrictions are not allowed for everyTime sign-in frequency see https://github.com/hashicorp/terraform-provider-azuread/issues/1225
		result.ApplicationEnforcedRestrictions = &stable.ApplicationEnforcedRestrictionsSessionControl{
			IsEnabled: nullable.Value(applicationEnforcedRestrictions),
		}
	}

	DisableResilienceDefaults := config["disable_resilience_defaults"].(bool)
	if pointer.From(signInFrequency.FrequencyInterval) != stable.SignInFrequencyInterval_EveryTime { // disable resilience defaults are not allowed for everyTime sign-in frequency see https://github.com/hashicorp/terraform-provider-azuread/issues/1225
		result.DisableResilienceDefaults = nullable.Value(DisableResilienceDefaults)
	}

	// API returns 400 error if signInFrequency is set with all default/zero values
	if (signInFrequency.IsEnabled.GetOrZero()) ||
		(signInFrequency.FrequencyInterval != nil && *signInFrequency.FrequencyInterval != stable.SignInFrequencyInterval_TimeBased) ||
		(signInFrequency.AuthenticationType != nil && *signInFrequency.AuthenticationType != stable.SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication) {
		result.SignInFrequency = &signInFrequency
	}

	// API does not accept ineffectual and sessionControls object, and it will not remove any existing sessionControls unless the entire object is set to null
	if (result.ApplicationEnforcedRestrictions == nil || !result.ApplicationEnforcedRestrictions.IsEnabled.GetOrZero()) &&
		result.CloudAppSecurity == nil && !result.DisableResilienceDefaults.GetOrZero() &&
		result.PersistentBrowser == nil && result.SignInFrequency == nil {
		return nil
	}

	return &result
}

func expandConditionalAccessFilter(in []interface{}) *stable.ConditionalAccessFilter {
	result := stable.ConditionalAccessFilter{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	config := in[0].(map[string]interface{})

	result.Mode = pointer.To(stable.FilterMode(config["mode"].(string)))
	result.Rule = pointer.To(config["rule"].(string))

	return &result
}

func expandGuestsOrExternalUsers(in []interface{}) *stable.ConditionalAccessGuestsOrExternalUsers {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	config := in[0].(map[string]interface{})
	result := stable.ConditionalAccessGuestsOrExternalUsers{}

	var guestOrExternalUserTypes *stable.ConditionalAccessGuestOrExternalUserTypes
	if len(config["guest_or_external_user_types"].([]interface{})) > 0 {
		values := strings.Join(tf.ExpandStringSlice(config["guest_or_external_user_types"].([]interface{})), ",")
		guestOrExternalUserTypes = pointer.To(stable.ConditionalAccessGuestOrExternalUserTypes(values))
	}

	result.GuestOrExternalUserTypes = guestOrExternalUserTypes
	result.ExternalTenants = expandExternalTenants(config["external_tenants"].([]interface{}))

	return &result
}

func expandExternalTenants(in []interface{}) stable.ConditionalAccessExternalTenants {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	config := in[0].(map[string]interface{})

	members := config["members"].([]interface{})
	membershipKind := stable.ConditionalAccessExternalTenantsMembershipKind(config["membership_kind"].(string))

	// only membership_kind enumerated is allowed to have members field set
	if membershipKind == stable.ConditionalAccessExternalTenantsMembershipKind_Enumerated {
		result := stable.ConditionalAccessEnumeratedExternalTenants{}

		result.MembershipKind = pointer.To(membershipKind)
		result.Members = tf.ExpandStringSlicePtr(members)

		return &result
	}

	result := stable.BaseConditionalAccessExternalTenantsImpl{}
	result.MembershipKind = pointer.To(membershipKind)

	return &result
}

func expandCountryNamedLocation(in []interface{}) *stable.CountryNamedLocation {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.CountryNamedLocation{}
	config := in[0].(map[string]interface{})

	countriesAndRegions := config["countries_and_regions"].([]interface{})
	includeUnknown := config["include_unknown_countries_and_regions"]

	result.CountriesAndRegions = tf.ExpandStringSlice(countriesAndRegions)
	result.IncludeUnknownCountriesAndRegions = pointer.To(includeUnknown.(bool))

	if countryLookupMethodType, ok := config["country_lookup_method"]; ok && countryLookupMethodType.(string) != "" {
		result.CountryLookupMethod = pointer.To(stable.CountryLookupMethodType(countryLookupMethodType.(string)))
	}

	return &result
}

func expandIPNamedLocation(in []interface{}) *stable.IPNamedLocation {
	if len(in) == 0 || in[0] == nil {
		return nil
	}

	result := stable.IPNamedLocation{}
	config := in[0].(map[string]interface{})

	ipRanges := config["ip_ranges"].([]interface{})
	trusted := config["trusted"]

	result.IPRanges = expandIPNamedLocationIPRange(ipRanges)
	result.IsTrusted = pointer.To(trusted.(bool))

	return &result
}

func expandIPNamedLocationIPRange(in []interface{}) []stable.IPRange {
	result := make([]stable.IPRange, 0)
	for _, cidr := range in {
		result = append(result, stable.IPv4CIDRRange{
			CIDRAddress: pointer.To(cidr.(string)),
		})
	}

	return result
}
