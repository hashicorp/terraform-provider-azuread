package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PartnerSecurityPartnerSecurityScore{}

type PartnerSecurityPartnerSecurityScore struct {
	// Contains customer-specific information for certain requirements.
	CustomerInsights *[]PartnerSecurityCustomerInsight `json:"customerInsights,omitempty"`

	// Contains a list of recent score changes.
	History *[]PartnerSecuritySecurityScoreHistory `json:"history,omitempty"`

	// The last time the data was checked.
	LastRefreshDateTime *string `json:"lastRefreshDateTime,omitempty"`

	// Contains the list of security requirements that make up the score.
	Requirements *[]PartnerSecuritySecurityRequirement `json:"requirements,omitempty"`

	// The last time the security score or related properties changed.
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

func (s PartnerSecurityPartnerSecurityScore) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PartnerSecurityPartnerSecurityScore{}

func (s PartnerSecurityPartnerSecurityScore) MarshalJSON() ([]byte, error) {
	type wrapper PartnerSecurityPartnerSecurityScore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PartnerSecurityPartnerSecurityScore: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerSecurityPartnerSecurityScore: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.partner.security.partnerSecurityScore"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PartnerSecurityPartnerSecurityScore: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PartnerSecurityPartnerSecurityScore{}

func (s *PartnerSecurityPartnerSecurityScore) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CustomerInsights    *[]PartnerSecurityCustomerInsight      `json:"customerInsights,omitempty"`
		History             *[]PartnerSecuritySecurityScoreHistory `json:"history,omitempty"`
		LastRefreshDateTime *string                                `json:"lastRefreshDateTime,omitempty"`
		UpdatedDateTime     *string                                `json:"updatedDateTime,omitempty"`
		Id                  *string                                `json:"id,omitempty"`
		ODataId             *string                                `json:"@odata.id,omitempty"`
		ODataType           *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CustomerInsights = decoded.CustomerInsights
	s.History = decoded.History
	s.LastRefreshDateTime = decoded.LastRefreshDateTime
	s.UpdatedDateTime = decoded.UpdatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PartnerSecurityPartnerSecurityScore into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["requirements"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Requirements into list []json.RawMessage: %+v", err)
		}

		output := make([]PartnerSecuritySecurityRequirement, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPartnerSecuritySecurityRequirementImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Requirements' for 'PartnerSecurityPartnerSecurityScore': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Requirements = &output
	}

	return nil
}
