package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlow interface {
	Entity
	IdentityUserFlow() BaseIdentityUserFlowImpl
}

var _ IdentityUserFlow = BaseIdentityUserFlowImpl{}

type BaseIdentityUserFlowImpl struct {
	UserFlowType *UserFlowType `json:"userFlowType,omitempty"`

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

func (s BaseIdentityUserFlowImpl) IdentityUserFlow() BaseIdentityUserFlowImpl {
	return s
}

func (s BaseIdentityUserFlowImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IdentityUserFlow = RawIdentityUserFlowImpl{}

// RawIdentityUserFlowImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityUserFlowImpl struct {
	identityUserFlow BaseIdentityUserFlowImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawIdentityUserFlowImpl) IdentityUserFlow() BaseIdentityUserFlowImpl {
	return s.identityUserFlow
}

func (s RawIdentityUserFlowImpl) Entity() BaseEntityImpl {
	return s.identityUserFlow.Entity()
}

var _ json.Marshaler = BaseIdentityUserFlowImpl{}

func (s BaseIdentityUserFlowImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIdentityUserFlowImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIdentityUserFlowImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIdentityUserFlowImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityUserFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIdentityUserFlowImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIdentityUserFlowImplementation(input []byte) (IdentityUserFlow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityUserFlow into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.b2cIdentityUserFlow") {
		var out B2cIdentityUserFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into B2cIdentityUserFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.b2xIdentityUserFlow") {
		var out B2xIdentityUserFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into B2xIdentityUserFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityUserFlowImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityUserFlowImpl: %+v", err)
	}

	return RawIdentityUserFlowImpl{
		identityUserFlow: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
