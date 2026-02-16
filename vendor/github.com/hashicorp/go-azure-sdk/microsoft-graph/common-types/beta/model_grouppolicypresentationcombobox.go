package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyUploadedPresentation = GroupPolicyPresentationComboBox{}

type GroupPolicyPresentationComboBox struct {
	// Localized default string displayed in the combo box. The default value is empty.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is
	// 1023.
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Specifies whether a value must be specified for the parameter. The default value is false.
	Required *bool `json:"required,omitempty"`

	// Localized strings listed in the drop-down list of the combo box. The default value is empty.
	Suggestions *[]string `json:"suggestions,omitempty"`

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

func (s GroupPolicyPresentationComboBox) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return BaseGroupPolicyUploadedPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationComboBox) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return BaseGroupPolicyPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationComboBox) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationComboBox{}

func (s GroupPolicyPresentationComboBox) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationComboBox
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationComboBox: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationComboBox: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationComboBox"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationComboBox: %+v", err)
	}

	return encoded, nil
}
