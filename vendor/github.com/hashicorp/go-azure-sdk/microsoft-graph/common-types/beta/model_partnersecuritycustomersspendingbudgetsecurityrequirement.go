package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnerSecuritySecurityRequirement = PartnerSecurityCustomersSpendingBudgetSecurityRequirement{}

type PartnerSecurityCustomersSpendingBudgetSecurityRequirement struct {
	// The number of customers with a spending budget set.
	CustomersWithSpendBudgetCount *int64 `json:"customersWithSpendBudgetCount,omitempty"`

	// The total number of customers associated with the partner.
	TotalCustomersCount *int64 `json:"totalCustomersCount,omitempty"`

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

func (s PartnerSecurityCustomersSpendingBudgetSecurityRequirement) PartnerSecuritySecurityRequirement() BasePartnerSecuritySecurityRequirementImpl {
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

func (s PartnerSecurityCustomersSpendingBudgetSecurityRequirement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityCustomersSpendingBudgetSecurityRequirement{}

func (s PartnerSecurityCustomersSpendingBudgetSecurityRequirement) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityCustomersSpendingBudgetSecurityRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityCustomersSpendingBudgetSecurityRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityCustomersSpendingBudgetSecurityRequirement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.customersSpendingBudgetSecurityRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityCustomersSpendingBudgetSecurityRequirement: %+v", err)
	}

	return encoded, nil
}
