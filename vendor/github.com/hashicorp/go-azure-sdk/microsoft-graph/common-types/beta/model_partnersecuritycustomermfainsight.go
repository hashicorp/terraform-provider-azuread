package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecurityCustomerMfaInsight struct {
	// The number of admins that are compliant with the MFA requirements
	CompliantAdminsCount *int64 `json:"compliantAdminsCount,omitempty"`

	// The number of users that are compliant with the MFA requirements
	CompliantNonAdminsCount *int64 `json:"compliantNonAdminsCount,omitempty"`

	LegacyPerUserMfaStatus           *PartnerSecurityPolicyStatus `json:"legacyPerUserMfaStatus,omitempty"`
	MfaConditionalAccessPolicyStatus *PartnerSecurityPolicyStatus `json:"mfaConditionalAccessPolicyStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SecurityDefaultsStatus *PartnerSecurityPolicyStatus `json:"securityDefaultsStatus,omitempty"`

	// The total number of users in the tenant
	TotalUsersCount *int64 `json:"totalUsersCount,omitempty"`
}
