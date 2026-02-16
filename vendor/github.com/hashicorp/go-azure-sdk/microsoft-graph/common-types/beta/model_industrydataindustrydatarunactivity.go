package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunActivity interface {
	Entity
	IndustryDataIndustryDataRunActivity() BaseIndustryDataIndustryDataRunActivityImpl
}

var _ IndustryDataIndustryDataRunActivity = BaseIndustryDataIndustryDataRunActivityImpl{}

type BaseIndustryDataIndustryDataRunActivityImpl struct {
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

func (s BaseIndustryDataIndustryDataRunActivityImpl) IndustryDataIndustryDataRunActivity() BaseIndustryDataIndustryDataRunActivityImpl {
	return s
}

func (s BaseIndustryDataIndustryDataRunActivityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataIndustryDataRunActivity = RawIndustryDataIndustryDataRunActivityImpl{}

// RawIndustryDataIndustryDataRunActivityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataIndustryDataRunActivityImpl struct {
	industryDataIndustryDataRunActivity BaseIndustryDataIndustryDataRunActivityImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawIndustryDataIndustryDataRunActivityImpl) IndustryDataIndustryDataRunActivity() BaseIndustryDataIndustryDataRunActivityImpl {
	return s.industryDataIndustryDataRunActivity
}

func (s RawIndustryDataIndustryDataRunActivityImpl) Entity() BaseEntityImpl {
	return s.industryDataIndustryDataRunActivity.Entity()
}

var _ json.Marshaler = BaseIndustryDataIndustryDataRunActivityImpl{}

func (s BaseIndustryDataIndustryDataRunActivityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataIndustryDataRunActivityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataIndustryDataRunActivityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataIndustryDataRunActivityImpl: %+v", err)
	}

	delete(decoded, "blockingError")
	delete(decoded, "displayName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.industryDataRunActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataIndustryDataRunActivityImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseIndustryDataIndustryDataRunActivityImpl{}

func (s *BaseIndustryDataIndustryDataRunActivityImpl) UnmarshalJSON(bytes []byte) error {
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
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseIndustryDataIndustryDataRunActivityImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activity"]; ok {
		impl, err := UnmarshalIndustryDataIndustryDataActivityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Activity' for 'BaseIndustryDataIndustryDataRunActivityImpl': %+v", err)
		}
		s.Activity = &impl
	}

	return nil
}

func UnmarshalIndustryDataIndustryDataRunActivityImplementation(input []byte) (IndustryDataIndustryDataRunActivity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRunActivity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFlowActivity") {
		var out IndustryDataInboundFlowActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFlowActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.outboundFlowActivity") {
		var out IndustryDataOutboundFlowActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataOutboundFlowActivity: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataIndustryDataRunActivityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataIndustryDataRunActivityImpl: %+v", err)
	}

	return RawIndustryDataIndustryDataRunActivityImpl{
		industryDataIndustryDataRunActivity: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
