package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ParentLabelDetails = LabelDetails{}

type LabelDetails struct {

	// Fields inherited from ParentLabelDetails

	// The color that the user interface should display for the label, if configured.
	Color nullable.Type[string] `json:"color,omitempty"`

	// The admin-defined description for the label.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The label ID is a globally unique identifier (GUID).
	Id nullable.Type[string] `json:"id,omitempty"`

	// Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
	IsActive *bool `json:"isActive,omitempty"`

	// The plaintext name of the label.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Parent ParentLabelDetails `json:"parent"`

	// The sensitivity value of the label, where lower is less sensitive.
	Sensitivity *int64 `json:"sensitivity,omitempty"`

	// The tooltip that should be displayed for the label in a user interface.
	Tooltip nullable.Type[string] `json:"tooltip,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s LabelDetails) ParentLabelDetails() BaseParentLabelDetailsImpl {
	return BaseParentLabelDetailsImpl{
		Color:       s.Color,
		Description: s.Description,
		Id:          s.Id,
		IsActive:    s.IsActive,
		Name:        s.Name,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Parent:      s.Parent,
		Sensitivity: s.Sensitivity,
		Tooltip:     s.Tooltip,
	}
}

var _ json.Marshaler = LabelDetails{}

func (s LabelDetails) MarshalJSON() ([]byte, error) {
	type wrapper LabelDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LabelDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.labelDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LabelDetails: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LabelDetails{}

func (s *LabelDetails) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Color       nullable.Type[string] `json:"color,omitempty"`
		Description nullable.Type[string] `json:"description,omitempty"`
		Id          nullable.Type[string] `json:"id,omitempty"`
		IsActive    *bool                 `json:"isActive,omitempty"`
		Name        nullable.Type[string] `json:"name,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
		Sensitivity *int64                `json:"sensitivity,omitempty"`
		Tooltip     nullable.Type[string] `json:"tooltip,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Color = decoded.Color
	s.Description = decoded.Description
	s.Id = decoded.Id
	s.IsActive = decoded.IsActive
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Sensitivity = decoded.Sensitivity
	s.Tooltip = decoded.Tooltip

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LabelDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["parent"]; ok {
		impl, err := UnmarshalParentLabelDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Parent' for 'LabelDetails': %+v", err)
		}
		s.Parent = impl
	}

	return nil
}
