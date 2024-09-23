package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionBase interface {
	Entity
	VirtualEventRegistrationQuestionBase() BaseVirtualEventRegistrationQuestionBaseImpl
}

var _ VirtualEventRegistrationQuestionBase = BaseVirtualEventRegistrationQuestionBaseImpl{}

type BaseVirtualEventRegistrationQuestionBaseImpl struct {
	// Display name of the registration question.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether an answer to the question is required. The default value is false.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

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

func (s BaseVirtualEventRegistrationQuestionBaseImpl) VirtualEventRegistrationQuestionBase() BaseVirtualEventRegistrationQuestionBaseImpl {
	return s
}

func (s BaseVirtualEventRegistrationQuestionBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ VirtualEventRegistrationQuestionBase = RawVirtualEventRegistrationQuestionBaseImpl{}

// RawVirtualEventRegistrationQuestionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVirtualEventRegistrationQuestionBaseImpl struct {
	virtualEventRegistrationQuestionBase BaseVirtualEventRegistrationQuestionBaseImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawVirtualEventRegistrationQuestionBaseImpl) VirtualEventRegistrationQuestionBase() BaseVirtualEventRegistrationQuestionBaseImpl {
	return s.virtualEventRegistrationQuestionBase
}

func (s RawVirtualEventRegistrationQuestionBaseImpl) Entity() BaseEntityImpl {
	return s.virtualEventRegistrationQuestionBase.Entity()
}

var _ json.Marshaler = BaseVirtualEventRegistrationQuestionBaseImpl{}

func (s BaseVirtualEventRegistrationQuestionBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseVirtualEventRegistrationQuestionBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseVirtualEventRegistrationQuestionBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseVirtualEventRegistrationQuestionBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventRegistrationQuestionBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseVirtualEventRegistrationQuestionBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalVirtualEventRegistrationQuestionBaseImplementation(input []byte) (VirtualEventRegistrationQuestionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventRegistrationQuestionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationCustomQuestion") {
		var out VirtualEventRegistrationCustomQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationCustomQuestion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventRegistrationPredefinedQuestion") {
		var out VirtualEventRegistrationPredefinedQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventRegistrationPredefinedQuestion: %+v", err)
		}
		return out, nil
	}

	var parent BaseVirtualEventRegistrationQuestionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVirtualEventRegistrationQuestionBaseImpl: %+v", err)
	}

	return RawVirtualEventRegistrationQuestionBaseImpl{
		virtualEventRegistrationQuestionBase: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
