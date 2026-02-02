package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecuritySensitivityLabel{}

type SecuritySensitivityLabel struct {
	// The color that the UI should display for the label, if configured.
	Color nullable.Type[string] `json:"color,omitempty"`

	// Returns the supported content formats for the label.
	ContentFormats *[]string `json:"contentFormats,omitempty"`

	// The admin-defined description for the label.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates whether the label has protection actions configured.
	HasProtection *bool `json:"hasProtection,omitempty"`

	// Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
	IsActive *bool `json:"isActive,omitempty"`

	// Indicates whether the label can be applied to content. False if the label is a parent with child labels.
	IsAppliable *bool `json:"isAppliable,omitempty"`

	// The plaintext name of the label.
	Name nullable.Type[string] `json:"name,omitempty"`

	Parent *SecuritySensitivityLabel `json:"parent,omitempty"`

	// The sensitivity value of the label, where lower is less sensitive.
	Sensitivity *int64 `json:"sensitivity,omitempty"`

	// The tooltip that should be displayed for the label in a UI.
	Tooltip nullable.Type[string] `json:"tooltip,omitempty"`

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

func (s SecuritySensitivityLabel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySensitivityLabel{}

func (s SecuritySensitivityLabel) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySensitivityLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySensitivityLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySensitivityLabel: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.sensitivityLabel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySensitivityLabel: %+v", err)
	}

	return encoded, nil
}
