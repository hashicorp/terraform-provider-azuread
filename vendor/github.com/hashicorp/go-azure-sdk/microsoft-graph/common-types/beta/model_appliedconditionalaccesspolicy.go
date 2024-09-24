package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicy struct {
	// The custom authentication strength enforced in a Conditional Access policy.
	AuthenticationStrength *AuthenticationStrength `json:"authenticationStrength,omitempty"`

	// Refers to the conditional access policy conditions that aren't satisfied. The possible values are: none, application,
	// users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState,
	// client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk,
	// authenticationFlows, insiderRisk . You must use the Prefer: include-unknown-enum-members request header to get the
	// following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk.
	// conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a
	// comma-separated list.
	ConditionsNotSatisfied *ConditionalAccessConditions `json:"conditionsNotSatisfied,omitempty"`

	// Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application,
	// users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState,
	// client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk,
	// authenticationFlows, insiderRisk. You must use the Prefer: include-unknown-enum-members request header to get the
	// following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk.
	// conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a
	// comma-separated list.
	ConditionsSatisfied *ConditionalAccessConditions `json:"conditionsSatisfied,omitempty"`

	// Name of the conditional access policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Refers to the grant controls enforced by the conditional access policy (example: 'Require multifactor
	// authentication').
	EnforcedGrantControls *[]string `json:"enforcedGrantControls,omitempty"`

	// Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
	EnforcedSessionControls *[]string `json:"enforcedSessionControls,omitempty"`

	// List of key-value pairs containing each matched exclude condition in the conditional access policy. Example:
	// [{'devicePlatform' : 'DevicePlatform'}] means the policy didn't apply, because the DevicePlatform condition was a
	// match.
	ExcludeRulesSatisfied *[]ConditionalAccessRuleSatisfied `json:"excludeRulesSatisfied,omitempty"`

	// Identifier of the conditional access policy.
	Id nullable.Type[string] `json:"id,omitempty"`

	// List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{
	// 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are
	// included and Users condition was a match because the user was part of the included Group rule.
	IncludeRulesSatisfied *[]ConditionalAccessRuleSatisfied `json:"includeRulesSatisfied,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (policy
	// isn't applied because policy conditions weren't met), notEnabled (this is due to the policy in a disabled state),
	// unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You
	// must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum:
	// reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
	Result *AppliedConditionalAccessPolicyResult `json:"result,omitempty"`

	// Refers to the session controls that a sign-in activity didn't satisfy. (Example: Application enforced Restrictions).
	SessionControlsNotSatisfied *[]string `json:"sessionControlsNotSatisfied,omitempty"`
}
