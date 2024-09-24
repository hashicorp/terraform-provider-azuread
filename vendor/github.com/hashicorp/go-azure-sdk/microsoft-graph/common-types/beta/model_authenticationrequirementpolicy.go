package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationRequirementPolicy struct {
	// Provides additional detail on the feature identified in requirementProvider.
	Detail nullable.Type[string] `json:"detail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Identifies what Microsoft Entra feature requires MFA in this policy. Possible values are: user, request,
	// servicePrincipal, v1ConditionalAccess, multiConditionalAccess, tenantSessionRiskPolicy, accountCompromisePolicies,
	// v1ConditionalAccessDependency, v1ConditionalAccessPolicyIdRequested,
	// mfaRegistrationRequiredByIdentityProtectionPolicy, baselineProtection, mfaRegistrationRequiredByBaselineProtection,
	// mfaRegistrationRequiredByMultiConditionalAccess, enforcedForCspAdmins, securityDefaults,
	// mfaRegistrationRequiredBySecurityDefaults, proofUpCodeRequest, crossTenantOutboundRule, gpsLocationCondition,
	// riskBasedPolicy, unknownFutureValue, scopeBasedAuthRequirementPolicy, authenticationStrengths . Also, note that you
	// must use the Prefer: include-unknown-enum-members request header to get the following value or values in this
	// evolvable enum: scopeBasedAuthRequirementPolicy, authenticationStrengths.
	RequirementProvider *RequirementProvider `json:"requirementProvider,omitempty"`
}
