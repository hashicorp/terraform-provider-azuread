package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataIndustryDataRunActivity = IndustryDataOutboundFlowActivity{}

type IndustryDataOutboundFlowActivity struct {

	// Fields inherited from IndustryDataIndustryDataRunActivity

	// The flow that was run by this activity.
	Activity *IndustryDataIndustryDataActivity `json:"activity,omitempty"`

	// An error object to diagnose critical failures in an activity.
	BlockingError *PublicError `json:"blockingError,omitempty"`

	// The name of the running flow.
	DisplayName *string `json:"displayName,omitempty"`

	Status *IndustryDataIndustryDataActivityStatus `json:"status,omitempty"`

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

func (s IndustryDataOutboundFlowActivity) IndustryDataIndustryDataRunActivity() BaseIndustryDataIndustryDataRunActivityImpl {
	return BaseIndustryDataIndustryDataRunActivityImpl{
		Activity:      s.Activity,
		BlockingError: s.BlockingError,
		DisplayName:   s.DisplayName,
		Status:        s.Status,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s IndustryDataOutboundFlowActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataOutboundFlowActivity{}

func (s IndustryDataOutboundFlowActivity) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataOutboundFlowActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataOutboundFlowActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOutboundFlowActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.outboundFlowActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataOutboundFlowActivity: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataOutboundFlowActivity{}

func (s *IndustryDataOutboundFlowActivity) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BlockingError *PublicError                            `json:"blockingError,omitempty"`
		DisplayName   *string                                 `json:"displayName,omitempty"`
		Status        *IndustryDataIndustryDataActivityStatus `json:"status,omitempty"`
		Id            *string                                 `json:"id,omitempty"`
		ODataId       *string                                 `json:"@odata.id,omitempty"`
		ODataType     *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BlockingError = decoded.BlockingError
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataOutboundFlowActivity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activity"]; ok {
		impl, err := UnmarshalIndustryDataIndustryDataActivityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Activity' for 'IndustryDataOutboundFlowActivity': %+v", err)
		}
		s.Activity = &impl
	}

	return nil
}
