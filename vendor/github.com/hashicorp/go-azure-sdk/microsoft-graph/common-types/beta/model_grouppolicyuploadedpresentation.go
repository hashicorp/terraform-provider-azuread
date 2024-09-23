package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedPresentation interface {
	Entity
	GroupPolicyPresentation
	GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl
}

var _ GroupPolicyUploadedPresentation = BaseGroupPolicyUploadedPresentationImpl{}

type BaseGroupPolicyUploadedPresentationImpl struct {

	// Fields inherited from GroupPolicyPresentation

	// The group policy definition associated with the presentation.
	Definition *GroupPolicyDefinition `json:"definition,omitempty"`

	// Localized text label for any presentation entity. The default value is empty.
	Label nullable.Type[string] `json:"label,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseGroupPolicyUploadedPresentationImpl) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return s
}

func (s BaseGroupPolicyUploadedPresentationImpl) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return BaseGroupPolicyPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s BaseGroupPolicyUploadedPresentationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GroupPolicyUploadedPresentation = RawGroupPolicyUploadedPresentationImpl{}

// RawGroupPolicyUploadedPresentationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupPolicyUploadedPresentationImpl struct {
	groupPolicyUploadedPresentation BaseGroupPolicyUploadedPresentationImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawGroupPolicyUploadedPresentationImpl) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return s.groupPolicyUploadedPresentation
}

func (s RawGroupPolicyUploadedPresentationImpl) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return s.groupPolicyUploadedPresentation.GroupPolicyPresentation()
}

func (s RawGroupPolicyUploadedPresentationImpl) Entity() BaseEntityImpl {
	return s.groupPolicyUploadedPresentation.Entity()
}

var _ json.Marshaler = BaseGroupPolicyUploadedPresentationImpl{}

func (s BaseGroupPolicyUploadedPresentationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGroupPolicyUploadedPresentationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGroupPolicyUploadedPresentationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGroupPolicyUploadedPresentationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyUploadedPresentation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGroupPolicyUploadedPresentationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGroupPolicyUploadedPresentationImplementation(input []byte) (GroupPolicyUploadedPresentation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyUploadedPresentation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationCheckBox") {
		var out GroupPolicyPresentationCheckBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationCheckBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationComboBox") {
		var out GroupPolicyPresentationComboBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationComboBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationDecimalTextBox") {
		var out GroupPolicyPresentationDecimalTextBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationDecimalTextBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationDropdownList") {
		var out GroupPolicyPresentationDropdownList
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationDropdownList: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationListBox") {
		var out GroupPolicyPresentationListBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationListBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationLongDecimalTextBox") {
		var out GroupPolicyPresentationLongDecimalTextBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationLongDecimalTextBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationMultiTextBox") {
		var out GroupPolicyPresentationMultiTextBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationMultiTextBox: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationText") {
		var out GroupPolicyPresentationText
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationText: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationTextBox") {
		var out GroupPolicyPresentationTextBox
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationTextBox: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupPolicyUploadedPresentationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupPolicyUploadedPresentationImpl: %+v", err)
	}

	return RawGroupPolicyUploadedPresentationImpl{
		groupPolicyUploadedPresentation: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
