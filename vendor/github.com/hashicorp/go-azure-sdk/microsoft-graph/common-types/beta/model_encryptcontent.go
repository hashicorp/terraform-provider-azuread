package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptContent interface {
	LabelActionBase
	EncryptContent() BaseEncryptContentImpl
}

var _ EncryptContent = BaseEncryptContentImpl{}

type BaseEncryptContentImpl struct {
	EncryptWith *EncryptWith `json:"encryptWith,omitempty"`

	// Fields inherited from LabelActionBase

	// The name of the action (for example, 'Encrypt', 'AddHeader').
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEncryptContentImpl) EncryptContent() BaseEncryptContentImpl {
	return s
}

func (s BaseEncryptContentImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EncryptContent = RawEncryptContentImpl{}

// RawEncryptContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEncryptContentImpl struct {
	encryptContent BaseEncryptContentImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawEncryptContentImpl) EncryptContent() BaseEncryptContentImpl {
	return s.encryptContent
}

func (s RawEncryptContentImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return s.encryptContent.LabelActionBase()
}

var _ json.Marshaler = BaseEncryptContentImpl{}

func (s BaseEncryptContentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEncryptContentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEncryptContentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEncryptContentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.encryptContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEncryptContentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalEncryptContentImplementation(input []byte) (EncryptContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EncryptContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptWithTemplate") {
		var out EncryptWithTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptWithTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptWithUserDefinedRights") {
		var out EncryptWithUserDefinedRights
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptWithUserDefinedRights: %+v", err)
		}
		return out, nil
	}

	var parent BaseEncryptContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEncryptContentImpl: %+v", err)
	}

	return RawEncryptContentImpl{
		encryptContent: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
