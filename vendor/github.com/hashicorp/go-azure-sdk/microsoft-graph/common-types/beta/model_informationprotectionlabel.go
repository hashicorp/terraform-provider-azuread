package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InformationProtectionLabel{}

type InformationProtectionLabel struct {
	// The color that the UI should display for the label, if configured.
	Color nullable.Type[string] `json:"color,omitempty"`

	// The admin-defined description for the label.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
	IsActive *bool `json:"isActive,omitempty"`

	// The plaintext name of the label.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The parent label associated with a child label. Null if label has no parent.
	Parent ParentLabelDetails `json:"parent"`

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

func (s InformationProtectionLabel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InformationProtectionLabel{}

func (s InformationProtectionLabel) MarshalJSON() ([]byte, error) {
	type wrapper InformationProtectionLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InformationProtectionLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InformationProtectionLabel: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.informationProtectionLabel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InformationProtectionLabel: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &InformationProtectionLabel{}

func (s *InformationProtectionLabel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Color       nullable.Type[string] `json:"color,omitempty"`
		Description nullable.Type[string] `json:"description,omitempty"`
		IsActive    *bool                 `json:"isActive,omitempty"`
		Name        nullable.Type[string] `json:"name,omitempty"`
		Sensitivity *int64                `json:"sensitivity,omitempty"`
		Tooltip     nullable.Type[string] `json:"tooltip,omitempty"`
		Id          *string               `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Color = decoded.Color
	s.Description = decoded.Description
	s.IsActive = decoded.IsActive
	s.Name = decoded.Name
	s.Sensitivity = decoded.Sensitivity
	s.Tooltip = decoded.Tooltip
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InformationProtectionLabel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["parent"]; ok {
		impl, err := UnmarshalParentLabelDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Parent' for 'InformationProtectionLabel': %+v", err)
		}
		s.Parent = impl
	}

	return nil
}
