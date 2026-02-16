package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteEntityBaseModel interface {
	Entity
	OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl
}

var _ OnenoteEntityBaseModel = BaseOnenoteEntityBaseModelImpl{}

type BaseOnenoteEntityBaseModelImpl struct {
	// The endpoint where you can get details about the page. Read-only.
	Self nullable.Type[string] `json:"self,omitempty"`

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

func (s BaseOnenoteEntityBaseModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return s
}

func (s BaseOnenoteEntityBaseModelImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OnenoteEntityBaseModel = RawOnenoteEntityBaseModelImpl{}

// RawOnenoteEntityBaseModelImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnenoteEntityBaseModelImpl struct {
	onenoteEntityBaseModel BaseOnenoteEntityBaseModelImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawOnenoteEntityBaseModelImpl) OnenoteEntityBaseModel() BaseOnenoteEntityBaseModelImpl {
	return s.onenoteEntityBaseModel
}

func (s RawOnenoteEntityBaseModelImpl) Entity() BaseEntityImpl {
	return s.onenoteEntityBaseModel.Entity()
}

var _ json.Marshaler = BaseOnenoteEntityBaseModelImpl{}

func (s BaseOnenoteEntityBaseModelImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOnenoteEntityBaseModelImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOnenoteEntityBaseModelImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOnenoteEntityBaseModelImpl: %+v", err)
	}

	delete(decoded, "self")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onenoteEntityBaseModel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOnenoteEntityBaseModelImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOnenoteEntityBaseModelImplementation(input []byte) (OnenoteEntityBaseModel, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnenoteEntityBaseModel into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteEntitySchemaObjectModel") {
		var out OnenoteEntitySchemaObjectModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteEntitySchemaObjectModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteResource") {
		var out OnenoteResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteResource: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnenoteEntityBaseModelImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnenoteEntityBaseModelImpl: %+v", err)
	}

	return RawOnenoteEntityBaseModelImpl{
		onenoteEntityBaseModel: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
