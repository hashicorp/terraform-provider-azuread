package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyUploadedPresentation = GroupPolicyPresentationLongDecimalTextBox{}

type GroupPolicyPresentationLongDecimalTextBox struct {
	// An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
	DefaultValue *int64 `json:"defaultValue,omitempty"`

	// An unsigned long that specifies the maximum allowed value. The default value is 9999.
	MaxValue *int64 `json:"maxValue,omitempty"`

	// An unsigned long that specifies the minimum allowed value. The default value is 0.
	MinValue *int64 `json:"minValue,omitempty"`

	// Requirement to enter a value in the parameter box. The default value is false.
	Required *bool `json:"required,omitempty"`

	// If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
	Spin *bool `json:"spin,omitempty"`

	// An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
	SpinStep *int64 `json:"spinStep,omitempty"`

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

func (s GroupPolicyPresentationLongDecimalTextBox) GroupPolicyUploadedPresentation() BaseGroupPolicyUploadedPresentationImpl {
	return BaseGroupPolicyUploadedPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationLongDecimalTextBox) GroupPolicyPresentation() BaseGroupPolicyPresentationImpl {
	return BaseGroupPolicyPresentationImpl{
		Definition:           s.Definition,
		Label:                s.Label,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationLongDecimalTextBox) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationLongDecimalTextBox{}

func (s GroupPolicyPresentationLongDecimalTextBox) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationLongDecimalTextBox
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationLongDecimalTextBox: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationLongDecimalTextBox: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationLongDecimalTextBox"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationLongDecimalTextBox: %+v", err)
	}

	return encoded, nil
}
