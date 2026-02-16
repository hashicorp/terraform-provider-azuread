package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttribute interface {
	Entity
	IdentityUserFlowAttribute() BaseIdentityUserFlowAttributeImpl
}

var _ IdentityUserFlowAttribute = BaseIdentityUserFlowAttributeImpl{}

type BaseIdentityUserFlowAttributeImpl struct {
	DataType *IdentityUserFlowAttributeDataType `json:"dataType,omitempty"`

	// The description of the user flow attribute that's shown to the user at the time of sign up.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the user flow attribute. Supports $filter (eq, ne).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	UserFlowAttributeType *IdentityUserFlowAttributeType `json:"userFlowAttributeType,omitempty"`

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

func (s BaseIdentityUserFlowAttributeImpl) IdentityUserFlowAttribute() BaseIdentityUserFlowAttributeImpl {
	return s
}

func (s BaseIdentityUserFlowAttributeImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IdentityUserFlowAttribute = RawIdentityUserFlowAttributeImpl{}

// RawIdentityUserFlowAttributeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityUserFlowAttributeImpl struct {
	identityUserFlowAttribute BaseIdentityUserFlowAttributeImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawIdentityUserFlowAttributeImpl) IdentityUserFlowAttribute() BaseIdentityUserFlowAttributeImpl {
	return s.identityUserFlowAttribute
}

func (s RawIdentityUserFlowAttributeImpl) Entity() BaseEntityImpl {
	return s.identityUserFlowAttribute.Entity()
}

var _ json.Marshaler = BaseIdentityUserFlowAttributeImpl{}

func (s BaseIdentityUserFlowAttributeImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIdentityUserFlowAttributeImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIdentityUserFlowAttributeImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIdentityUserFlowAttributeImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityUserFlowAttribute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIdentityUserFlowAttributeImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIdentityUserFlowAttributeImplementation(input []byte) (IdentityUserFlowAttribute, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityUserFlowAttribute into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityBuiltInUserFlowAttribute") {
		var out IdentityBuiltInUserFlowAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityBuiltInUserFlowAttribute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityCustomUserFlowAttribute") {
		var out IdentityCustomUserFlowAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityCustomUserFlowAttribute: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityUserFlowAttributeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityUserFlowAttributeImpl: %+v", err)
	}

	return RawIdentityUserFlowAttributeImpl{
		identityUserFlowAttribute: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
