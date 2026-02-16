package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentation interface {
	Entity
	GroupPolicyPresentation() BaseGroupPolicyPresentationImpl
}

var _ GroupPolicyPresentation = BaseGroupPolicyPresentationImpl{}

type BaseGroupPolicyPresentationImpl struct {
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

func (s BaseGroupPolicyPresentationImpl) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return s
}

func (s BaseGroupPolicyPresentationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GroupPolicyPresentation = RawGroupPolicyPresentationImpl{}

// RawGroupPolicyPresentationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupPolicyPresentationImpl struct {
	groupPolicyPresentation BaseGroupPolicyPresentationImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawGroupPolicyPresentationImpl) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return s.groupPolicyPresentation
}

func (s RawGroupPolicyPresentationImpl) Entity() BaseEntityImpl {
	return s.groupPolicyPresentation.Entity()
}

var _ json.Marshaler = BaseGroupPolicyPresentationImpl{}

func (s BaseGroupPolicyPresentationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGroupPolicyPresentationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGroupPolicyPresentationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGroupPolicyPresentationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGroupPolicyPresentationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGroupPolicyPresentationImplementation(input []byte) (GroupPolicyPresentation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyUploadedPresentation") {
		var out GroupPolicyUploadedPresentation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyUploadedPresentation: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupPolicyPresentationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupPolicyPresentationImpl: %+v", err)
	}

	return RawGroupPolicyPresentationImpl{
		groupPolicyPresentation: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
