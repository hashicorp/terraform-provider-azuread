package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnerSecuritySecurityScoreHistory{}

type PartnerSecuritySecurityScoreHistory struct {
	// The number of compliant security requirements at the time.
	CompliantRequirementsCount *int64 `json:"compliantRequirementsCount,omitempty"`

	// The date the history entry was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The total number of requirements at the time.
	TotalRequirementsCount *int64 `json:"totalRequirementsCount,omitempty"`

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

func (s PartnerSecuritySecurityScoreHistory) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecuritySecurityScoreHistory{}

func (s PartnerSecuritySecurityScoreHistory) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecuritySecurityScoreHistory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecuritySecurityScoreHistory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecuritySecurityScoreHistory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.securityScoreHistory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecuritySecurityScoreHistory: %+v", err)
	}

	return encoded, nil
}
