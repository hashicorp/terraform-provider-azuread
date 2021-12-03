package conditionalaccess

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func flattenConditionalAccessConditionSet(in *msgraph.ConditionalAccessConditionSet) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"applications":        flattenConditionalAccessApplications(in.Applications),
			"users":               flattenConditionalAccessUsers(in.Users),
			"client_app_types":    tf.FlattenStringSlicePtr(in.ClientAppTypes),
			"devices":             flattenConditionalAccessDevices(in.Devices),
			"locations":           flattenConditionalAccessLocations(in.Locations),
			"platforms":           flattenConditionalAccessPlatforms(in.Platforms),
			"sign_in_risk_levels": tf.FlattenStringSlicePtr(in.SignInRiskLevels),
			"user_risk_levels":    tf.FlattenStringSlicePtr(in.UserRiskLevels),
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

func flattenConditionalAccessUsers(in *msgraph.ConditionalAccessUsers) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"included_users":  tf.FlattenStringSlicePtr(in.IncludeUsers),
			"excluded_users":  tf.FlattenStringSlicePtr(in.ExcludeUsers),
			"included_groups": tf.FlattenStringSlicePtr(in.IncludeGroups),
			"excluded_groups": tf.FlattenStringSlicePtr(in.ExcludeGroups),
			"included_roles":  tf.FlattenStringSlicePtr(in.IncludeRoles),
			"excluded_roles":  tf.FlattenStringSlicePtr(in.ExcludeRoles),
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

	return []interface{}{
		map[string]interface{}{
			"operator":                      in.Operator,
			"built_in_controls":             tf.FlattenStringSlicePtr(in.BuiltInControls),
			"custom_authentication_factors": tf.FlattenStringSlicePtr(in.CustomAuthenticationFactors),
			"terms_of_use":                  tf.FlattenStringSlicePtr(in.TermsOfUse),
		},
	}
}

func flattenConditionalAccessSessionControls(in *msgraph.ConditionalAccessSessionControls) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	applicationEnforceRestrictions := false
	if in.ApplicationEnforcedRestrictions != nil {
		applicationEnforceRestrictions = *in.ApplicationEnforcedRestrictions.IsEnabled
	}

	cloudAppSecurity := ""
	if in.CloudAppSecurity != nil && in.CloudAppSecurity.CloudAppSecurityType != nil {
		cloudAppSecurity = *in.CloudAppSecurity.CloudAppSecurityType
	}

	signInFrequency := 0
	signInFrequencyPeriod := ""
	if in.SignInFrequency != nil && in.SignInFrequency.Value != nil && in.SignInFrequency.Type != nil {
		signInFrequency = int(*in.SignInFrequency.Value)
		signInFrequencyPeriod = *in.SignInFrequency.Type
	}

	persistentBrowserMode := ""
	if in.PersistentBrowser != nil && in.PersistentBrowser.Mode != nil {
		persistentBrowserMode = *in.PersistentBrowser.Mode
	}

	return []interface{}{
		map[string]interface{}{
			"application_enforced_restrictions_enabled": applicationEnforceRestrictions,
			"cloud_app_security_policy":                 cloudAppSecurity,
			"persistent_browser_mode":                   persistentBrowserMode,
			"sign_in_frequency":                         signInFrequency,
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
	signInRiskLevels := config["sign_in_risk_levels"].([]interface{})
	userRiskLevels := config["user_risk_levels"].([]interface{})

	result.Applications = expandConditionalAccessApplications(applications)
	result.Users = expandConditionalAccessUsers(users)
	result.ClientAppTypes = tf.ExpandStringSlicePtr(clientAppTypes)
	result.Devices = expandConditionalAccessDevices(devices)
	result.Locations = expandConditionalAccessLocations(locations)
	result.Platforms = expandConditionalAccessPlatforms(platforms)
	result.SignInRiskLevels = tf.ExpandStringSlicePtr(signInRiskLevels)
	result.UserRiskLevels = tf.ExpandStringSlicePtr(userRiskLevels)

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

	return &result
}

func expandConditionalAccessPlatforms(in []interface{}) *msgraph.ConditionalAccessPlatforms {
	result := msgraph.ConditionalAccessPlatforms{}
	if len(in) == 0 || in[0] == nil {
		return &result
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
	builtInControls := config["built_in_controls"].([]interface{})
	customAuthenticationFactors := config["custom_authentication_factors"].([]interface{})
	termsOfUse := config["terms_of_use"].([]interface{})

	result.Operator = &operator
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
		IsEnabled: utils.Bool(config["application_enforced_restrictions_enabled"].(bool)),
	}

	if cloudAppSecurity := config["cloud_app_security_policy"].(string); cloudAppSecurity != "" {
		result.CloudAppSecurity = &msgraph.CloudAppSecurityControl{
			IsEnabled:            utils.Bool(true),
			CloudAppSecurityType: utils.String(cloudAppSecurity),
		}
	}

	if persistentBrowserMode := config["persistent_browser_mode"].(string); persistentBrowserMode != "" {
		result.PersistentBrowser = &msgraph.PersistentBrowserSessionControl{
			IsEnabled: utils.Bool(true),
			Mode:      utils.String(persistentBrowserMode),
		}
	}

	if signInFrequency := config["sign_in_frequency"].(int); signInFrequency > 0 {
		result.SignInFrequency = &msgraph.SignInFrequencySessionControl{
			IsEnabled: utils.Bool(true),
			Type:      utils.String(config["sign_in_frequency_period"].(string)),
			Value:     utils.Int32(int32(signInFrequency)),
		}
	}

	return &result
}

func expandConditionalAccessFilter(in []interface{}) *msgraph.ConditionalAccessFilter {
	result := msgraph.ConditionalAccessFilter{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	config := in[0].(map[string]interface{})

	result.Mode = utils.String(config["mode"].(string))
	result.Rule = utils.String(config["rule"].(string))

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
	result.IncludeUnknownCountriesAndRegions = utils.Bool(includeUnknown.(bool))

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
	result.IsTrusted = utils.Bool(trusted.(bool))

	return &result
}

func expandIPNamedLocationIPRange(in []interface{}) *[]msgraph.IPNamedLocationIPRange {
	if len(in) == 0 {
		return nil
	}

	result := make([]msgraph.IPNamedLocationIPRange, 0)
	for _, cidr := range in {
		result = append(result, msgraph.IPNamedLocationIPRange{
			CIDRAddress: utils.String(cidr.(string)),
		})
	}

	return &result
}
