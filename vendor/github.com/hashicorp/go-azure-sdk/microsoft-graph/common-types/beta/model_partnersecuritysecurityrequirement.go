package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityRequirement interface {
	Entity
	PartnerSecuritySecurityRequirement() BasePartnerSecuritySecurityRequirementImpl
}

var _ PartnerSecuritySecurityRequirement = BasePartnerSecuritySecurityRequirementImpl{}

type BasePartnerSecuritySecurityRequirementImpl struct {
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

func (s BasePartnerSecuritySecurityRequirementImpl) PartnerSecuritySecurityRequirement() BasePartnerSecuritySecurityRequirementImpl {
	return s
}

func (s BasePartnerSecuritySecurityRequirementImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PartnerSecuritySecurityRequirement = RawPartnerSecuritySecurityRequirementImpl{}

// RawPartnerSecuritySecurityRequirementImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartnerSecuritySecurityRequirementImpl struct {
	partnerSecuritySecurityRequirement BasePartnerSecuritySecurityRequirementImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawPartnerSecuritySecurityRequirementImpl) PartnerSecuritySecurityRequirement() BasePartnerSecuritySecurityRequirementImpl {
	return s.partnerSecuritySecurityRequirement
}

func (s RawPartnerSecuritySecurityRequirementImpl) Entity() BaseEntityImpl {
	return s.partnerSecuritySecurityRequirement.Entity()
}

var _ json.Marshaler = BasePartnerSecuritySecurityRequirementImpl{}

func (s BasePartnerSecuritySecurityRequirementImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePartnerSecuritySecurityRequirementImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePartnerSecuritySecurityRequirementImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePartnerSecuritySecurityRequirementImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.securityRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePartnerSecuritySecurityRequirementImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPartnerSecuritySecurityRequirementImplementation(input []byte) (PartnerSecuritySecurityRequirement, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecuritySecurityRequirement into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.adminsMfaEnforcedSecurityRequirement") {
		var out PartnerSecurityAdminsMfaEnforcedSecurityRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityAdminsMfaEnforcedSecurityRequirement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.customersMfaEnforcedSecurityRequirement") {
		var out PartnerSecurityCustomersMfaEnforcedSecurityRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityCustomersMfaEnforcedSecurityRequirement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.customersSpendingBudgetSecurityRequirement") {
		var out PartnerSecurityCustomersSpendingBudgetSecurityRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityCustomersSpendingBudgetSecurityRequirement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.responseTimeSecurityRequirement") {
		var out PartnerSecurityResponseTimeSecurityRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityResponseTimeSecurityRequirement: %+v", err)
		}
		return out, nil
	}

	var parent BasePartnerSecuritySecurityRequirementImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartnerSecuritySecurityRequirementImpl: %+v", err)
	}

	return RawPartnerSecuritySecurityRequirementImpl{
		partnerSecuritySecurityRequirement: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
