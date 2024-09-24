package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyUploadedPresentation = GroupPolicyPresentationMultiTextBox{}

type GroupPolicyPresentationMultiTextBox struct {
	// An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
	MaxLength *int64 `json:"maxLength,omitempty"`

	// An unsigned integer that specifies the maximum number of strings. Default value is 0.
	MaxStrings *int64 `json:"maxStrings,omitempty"`

	// Requirement to enter a value in the text box. Default value is false.
	Required *bool `json:"required,omitempty"`

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

func (s GroupPolicyPresentationMultiTextBox) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return BaseGroupPolicyUploadedPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationMultiTextBox) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return BaseGroupPolicyPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationMultiTextBox) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationMultiTextBox{}

func (s GroupPolicyPresentationMultiTextBox) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationMultiTextBox
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationMultiTextBox: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationMultiTextBox: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationMultiTextBox"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationMultiTextBox: %+v", err)
	}

	return encoded, nil
}
