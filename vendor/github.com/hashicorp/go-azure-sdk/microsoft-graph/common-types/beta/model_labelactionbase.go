package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelActionBase interface {
	LabelActionBase() BaseLabelActionBaseImpl
}

var _ LabelActionBase = BaseLabelActionBaseImpl{}

type BaseLabelActionBaseImpl struct {
	// The name of the action (for example, 'Encrypt', 'AddHeader').
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseLabelActionBaseImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return s
}

var _ LabelActionBase = RawLabelActionBaseImpl{}

// RawLabelActionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLabelActionBaseImpl struct {
	labelActionBase BaseLabelActionBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawLabelActionBaseImpl) LabelActionBase() BaseLabelActionBaseImpl {
	return s.labelActionBase
}

func UnmarshalLabelActionBaseImplementation(input []byte) (LabelActionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelActionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.encryptContent") {
		var out EncryptContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EncryptContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.markContent") {
		var out MarkContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MarkContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectGroup") {
		var out ProtectGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectOnlineMeetingAction") {
		var out ProtectOnlineMeetingAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectOnlineMeetingAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.protectSite") {
		var out ProtectSite
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProtectSite: %+v", err)
		}
		return out, nil
	}

	var parent BaseLabelActionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLabelActionBaseImpl: %+v", err)
	}

	return RawLabelActionBaseImpl{
		labelActionBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
