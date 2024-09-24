package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarkContent interface {
	LabelActionBase
	MarkContent() BaseMarkContentImpl
}

var _ MarkContent = BaseMarkContentImpl{}

type BaseMarkContentImpl struct {
	FontColor nullable.Type[string] `json:"fontColor,omitempty"`
	FontSize  nullable.Type[int64]  `json:"fontSize,omitempty"`
	Text      nullable.Type[string] `json:"text,omitempty"`

	// Fields inherited from LabelActionBase

	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMarkContentImpl) MarkContent() BaseMarkContentImpl {
	return s
}

func (s BaseMarkContentImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MarkContent = RawMarkContentImpl{}

// RawMarkContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMarkContentImpl struct {
	markContent BaseMarkContentImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawMarkContentImpl) MarkContent() BaseMarkContentImpl {
	return s.markContent
}

func (s RawMarkContentImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return s.markContent.LabelActionBase()
}

var _ json.Marshaler = BaseMarkContentImpl{}

func (s BaseMarkContentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMarkContentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMarkContentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMarkContentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.markContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMarkContentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalMarkContentImplementation(input []byte) (MarkContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MarkContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.addFooter") {
		var out AddFooter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddFooter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.addHeader") {
		var out AddHeader
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddHeader: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.addWatermark") {
		var out AddWatermark
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddWatermark: %+v", err)
		}
		return out, nil
	}

	var parent BaseMarkContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMarkContentImpl: %+v", err)
	}

	return RawMarkContentImpl{
		markContent: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
