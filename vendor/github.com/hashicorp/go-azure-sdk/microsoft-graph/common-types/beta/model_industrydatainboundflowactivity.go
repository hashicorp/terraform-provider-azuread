package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataIndustryDataRunActivity = IndustryDataInboundFlowActivity{}

type IndustryDataInboundFlowActivity struct {

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

func (s IndustryDataInboundFlowActivity) IndustryDataIndustryDataRunActivity() BaseIndustryDataIndustryDataRunActivityImpl {
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

func (s IndustryDataInboundFlowActivity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataInboundFlowActivity{}

func (s IndustryDataInboundFlowActivity) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataInboundFlowActivity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataInboundFlowActivity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataInboundFlowActivity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.inboundFlowActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataInboundFlowActivity: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataInboundFlowActivity{}

func (s *IndustryDataInboundFlowActivity) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling IndustryDataInboundFlowActivity into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activity"]; ok {
		impl, err := UnmarshalIndustryDataIndustryDataActivityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Activity' for 'IndustryDataInboundFlowActivity': %+v", err)
		}
		s.Activity = &impl
	}

	return nil
}
