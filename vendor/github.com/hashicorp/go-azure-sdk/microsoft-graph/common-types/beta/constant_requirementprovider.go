package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequirementProvider string

const (
	RequirementProvider_AccountCompromisePolicies                         RequirementProvider = "accountCompromisePolicies"
	RequirementProvider_AuthenticationStrengths                           RequirementProvider = "authenticationStrengths"
	RequirementProvider_BaselineProtection                                RequirementProvider = "baselineProtection"
	RequirementProvider_CrossTenantOutboundRule                           RequirementProvider = "crossTenantOutboundRule"
	RequirementProvider_EnforcedForCspAdmins                              RequirementProvider = "enforcedForCspAdmins"
	RequirementProvider_GpsLocationCondition                              RequirementProvider = "gpsLocationCondition"
	RequirementProvider_MfaRegistrationRequiredByBaselineProtection       RequirementProvider = "mfaRegistrationRequiredByBaselineProtection"
	RequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy RequirementProvider = "mfaRegistrationRequiredByIdentityProtectionPolicy"
	RequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess   RequirementProvider = "mfaRegistrationRequiredByMultiConditionalAccess"
	RequirementProvider_MfaRegistrationRequiredBySecurityDefaults         RequirementProvider = "mfaRegistrationRequiredBySecurityDefaults"
	RequirementProvider_MultiConditionalAccess                            RequirementProvider = "multiConditionalAccess"
	RequirementProvider_ProofUpCodeRequest                                RequirementProvider = "proofUpCodeRequest"
	RequirementProvider_Request                                           RequirementProvider = "request"
	RequirementProvider_RiskBasedPolicy                                   RequirementProvider = "riskBasedPolicy"
	RequirementProvider_ScopeBasedAuthRequirementPolicy                   RequirementProvider = "scopeBasedAuthRequirementPolicy"
	RequirementProvider_SecurityDefaults                                  RequirementProvider = "securityDefaults"
	RequirementProvider_ServicePrincipal                                  RequirementProvider = "servicePrincipal"
	RequirementProvider_TenantSessionRiskPolicy                           RequirementProvider = "tenantSessionRiskPolicy"
	RequirementProvider_User                                              RequirementProvider = "user"
	RequirementProvider_V1ConditionalAccess                               RequirementProvider = "v1ConditionalAccess"
	RequirementProvider_V1ConditionalAccessDependency                     RequirementProvider = "v1ConditionalAccessDependency"
	RequirementProvider_V1ConditionalAccessPolicyIdRequested              RequirementProvider = "v1ConditionalAccessPolicyIdRequested"
)

func PossibleValuesForRequirementProvider() []string {
	return []string{
		string(RequirementProvider_AccountCompromisePolicies),
		string(RequirementProvider_AuthenticationStrengths),
		string(RequirementProvider_BaselineProtection),
		string(RequirementProvider_CrossTenantOutboundRule),
		string(RequirementProvider_EnforcedForCspAdmins),
		string(RequirementProvider_GpsLocationCondition),
		string(RequirementProvider_MfaRegistrationRequiredByBaselineProtection),
		string(RequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy),
		string(RequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess),
		string(RequirementProvider_MfaRegistrationRequiredBySecurityDefaults),
		string(RequirementProvider_MultiConditionalAccess),
		string(RequirementProvider_ProofUpCodeRequest),
		string(RequirementProvider_Request),
		string(RequirementProvider_RiskBasedPolicy),
		string(RequirementProvider_ScopeBasedAuthRequirementPolicy),
		string(RequirementProvider_SecurityDefaults),
		string(RequirementProvider_ServicePrincipal),
		string(RequirementProvider_TenantSessionRiskPolicy),
		string(RequirementProvider_User),
		string(RequirementProvider_V1ConditionalAccess),
		string(RequirementProvider_V1ConditionalAccessDependency),
		string(RequirementProvider_V1ConditionalAccessPolicyIdRequested),
	}
}

func (s *RequirementProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRequirementProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRequirementProvider(input string) (*RequirementProvider, error) {
	vals := map[string]RequirementProvider{
		"accountcompromisepolicies":                         RequirementProvider_AccountCompromisePolicies,
		"authenticationstrengths":                           RequirementProvider_AuthenticationStrengths,
		"baselineprotection":                                RequirementProvider_BaselineProtection,
		"crosstenantoutboundrule":                           RequirementProvider_CrossTenantOutboundRule,
		"enforcedforcspadmins":                              RequirementProvider_EnforcedForCspAdmins,
		"gpslocationcondition":                              RequirementProvider_GpsLocationCondition,
		"mfaregistrationrequiredbybaselineprotection":       RequirementProvider_MfaRegistrationRequiredByBaselineProtection,
		"mfaregistrationrequiredbyidentityprotectionpolicy": RequirementProvider_MfaRegistrationRequiredByIdentityProtectionPolicy,
		"mfaregistrationrequiredbymulticonditionalaccess":   RequirementProvider_MfaRegistrationRequiredByMultiConditionalAccess,
		"mfaregistrationrequiredbysecuritydefaults":         RequirementProvider_MfaRegistrationRequiredBySecurityDefaults,
		"multiconditionalaccess":                            RequirementProvider_MultiConditionalAccess,
		"proofupcoderequest":                                RequirementProvider_ProofUpCodeRequest,
		"request":                                           RequirementProvider_Request,
		"riskbasedpolicy":                                   RequirementProvider_RiskBasedPolicy,
		"scopebasedauthrequirementpolicy":                   RequirementProvider_ScopeBasedAuthRequirementPolicy,
		"securitydefaults":                                  RequirementProvider_SecurityDefaults,
		"serviceprincipal":                                  RequirementProvider_ServicePrincipal,
		"tenantsessionriskpolicy":                           RequirementProvider_TenantSessionRiskPolicy,
		"user":                                              RequirementProvider_User,
		"v1conditionalaccess":                               RequirementProvider_V1ConditionalAccess,
		"v1conditionalaccessdependency":                     RequirementProvider_V1ConditionalAccessDependency,
		"v1conditionalaccesspolicyidrequested":              RequirementProvider_V1ConditionalAccessPolicyIdRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequirementProvider(input)
	return &out, nil
}
