package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnerSecuritySecurityRequirement = PartnerSecurityAdminsMfaEnforcedSecurityRequirement{}

type PartnerSecurityAdminsMfaEnforcedSecurityRequirement struct {
	// The number of admins who are required to use MFA, but haven't completed registration.
	AdminsRequiredNotUsingMfaCount *int64 `json:"adminsRequiredNotUsingMfaCount,omitempty"`

	LegacyPerUserMfaStatus           *PartnerSecurityPolicyStatus `json:"legacyPerUserMfaStatus,omitempty"`
	MfaConditionalAccessPolicyStatus *PartnerSecurityPolicyStatus `json:"mfaConditionalAccessPolicyStatus,omitempty"`

	// The number of admins who are using MFA.
	MfaEnabledAdminsCount *int64 `json:"mfaEnabledAdminsCount,omitempty"`

	// The number of users who are using MFA.
	MfaEnabledUsersCount *int64 `json:"mfaEnabledUsersCount,omitempty"`

	SecurityDefaultsStatus *PartnerSecurityPolicyStatus `json:"securityDefaultsStatus,omitempty"`

	// The total number of admins in the partner's tenant.
	TotalAdminsCount *int64 `json:"totalAdminsCount,omitempty"`

	// The total number of users in the partner's tenant.
	TotalUsersCount *int64 `json:"totalUsersCount,omitempty"`

	// The number of users who are required to use MFA, but haven't completed registration.
	UsersRequiredNotUsingMfaCount *int64 `json:"usersRequiredNotUsingMfaCount,omitempty"`

	// Fields inherited from PartnerSecuritySecurityRequirement

	// The link to the site where the admin can take action on the requirement.
	ActionUrl *string `json:"actionUrl,omitempty"`

	ComplianceStatus *PartnerSecurityComplianceStatus `json:"complianceStatus,omitempty"`

	// The link to documentation for the requirement.
	HelpUrl *string `json:"helpUrl,omitempty"`

	// The maximum score possible for the requirement.
	MaxScore *int64 `json:"maxScore,omitempty"`

	RequirementType *PartnerSecuritySecurityRequirementType `json:"requirementType,omitempty"`

	// The score received for this requirement.
	Score *int64 `json:"score,omitempty"`

	State *PartnerSecuritySecurityRequirementState `json:"state,omitempty"`

	// The date the requirement properties were last updated.
	UpdatedDateTime *string `json:"updatedDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PartnerSecurityAdminsMfaEnforcedSecurityRequirement) PartnerSecuritySecurityRequirement() BasePartnerSecuritySecurityRequirementImpl {
	return BasePartnerSecuritySecurityRequirementImpl{
		ActionUrl:        s.ActionUrl,
		ComplianceStatus: s.ComplianceStatus,
		HelpUrl:          s.HelpUrl,
		MaxScore:         s.MaxScore,
		RequirementType:  s.RequirementType,
		Score:            s.Score,
		State:            s.State,
		UpdatedDateTime:  s.UpdatedDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s PartnerSecurityAdminsMfaEnforcedSecurityRequirement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityAdminsMfaEnforcedSecurityRequirement{}

func (s PartnerSecurityAdminsMfaEnforcedSecurityRequirement) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityAdminsMfaEnforcedSecurityRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityAdminsMfaEnforcedSecurityRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityAdminsMfaEnforcedSecurityRequirement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.adminsMfaEnforcedSecurityRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityAdminsMfaEnforcedSecurityRequirement: %+v", err)
	}

	return encoded, nil
}
