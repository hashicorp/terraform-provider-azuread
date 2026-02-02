package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicy struct {
	// Refers to the name of the conditional access policy (example: 'Require MFA for Salesforce').
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Refers to the grant controls enforced by the conditional access policy (example: 'Require multifactor
	// authentication').
	EnforcedGrantControls *[]string `json:"enforcedGrantControls,omitempty"`

	// Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
	EnforcedSessionControls *[]string `json:"enforcedSessionControls,omitempty"`

	// An identifier of the conditional access policy. Supports $filter (eq).
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (policy
	// isn't applied because policy conditions weren't met), notEnabled (This is due to the policy in a disabled state),
	// unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. Use
	// the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum:
	// reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
	Result *AppliedConditionalAccessPolicyResult `json:"result,omitempty"`
}
