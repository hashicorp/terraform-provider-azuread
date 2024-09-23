package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationConfiguration interface {
	Entity
	VirtualEventRegistrationConfiguration() BaseVirtualEventRegistrationConfigurationImpl
}

var _ VirtualEventRegistrationConfiguration = BaseVirtualEventRegistrationConfigurationImpl{}

type BaseVirtualEventRegistrationConfigurationImpl struct {
	Capacity           nullable.Type[int64]                    `json:"capacity,omitempty"`
	Questions          *[]VirtualEventRegistrationQuestionBase `json:"questions,omitempty"`
	RegistrationWebUrl nullable.Type[string]                   `json:"registrationWebUrl,omitempty"`

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

func (s BaseVirtualEventRegistrationConfigurationImpl) VirtualEventRegistrationConfiguration() BaseVirtualEventRegistrationConfigurationImpl {
	return s
}

func (s BaseVirtualEventRegistrationConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ VirtualEventRegistrationConfiguration = RawVirtualEventRegistrationConfigurationImpl{}

// RawVirtualEventRegistrationConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVirtualEventRegistrationConfigurationImpl struct {
	virtualEventRegistrationConfiguration BaseVirtualEventRegistrationConfigurationImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawVirtualEventRegistrationConfigurationImpl) VirtualEventRegistrationConfiguration() BaseVirtualEventRegistrationConfigurationImpl {
	return s.virtualEventRegistrationConfiguration
}

func (s RawVirtualEventRegistrationConfigurationImpl) Entity() BaseEntityImpl {
	return s.virtualEventRegistrationConfiguration.Entity()
}

var _ json.Marshaler = BaseVirtualEventRegistrationConfigurationImpl{}

func (s BaseVirtualEventRegistrationConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseVirtualEventRegistrationConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseVirtualEventRegistrationConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseVirtualEventRegistrationConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventRegistrationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseVirtualEventRegistrationConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseVirtualEventRegistrationConfigurationImpl{}

func (s *BaseVirtualEventRegistrationConfigurationImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Capacity           nullable.Type[int64]  `json:"capacity,omitempty"`
		RegistrationWebUrl nullable.Type[string] `json:"registrationWebUrl,omitempty"`
		Id                 *string               `json:"id,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Capacity = decoded.Capacity
	s.RegistrationWebUrl = decoded.RegistrationWebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseVirtualEventRegistrationConfigurationImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["questions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Questions into list []json.RawMessage: %+v", err)
		}

		output := make([]VirtualEventRegistrationQuestionBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalVirtualEventRegistrationQuestionBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Questions' for 'BaseVirtualEventRegistrationConfigurationImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Questions = &output
	}

	return nil
}

func UnmarshalVirtualEventRegistrationConfigurationImplementation(input []byte) (VirtualEventRegistrationConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventRegistrationConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventWebinarRegistrationConfiguration") {
		var out VirtualEventWebinarRegistrationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventWebinarRegistrationConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseVirtualEventRegistrationConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVirtualEventRegistrationConfigurationImpl: %+v", err)
	}

	return RawVirtualEventRegistrationConfigurationImpl{
		virtualEventRegistrationConfiguration: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
