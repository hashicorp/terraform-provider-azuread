package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParentLabelDetails interface {
	ParentLabelDetails() BaseParentLabelDetailsImpl
}

var _ ParentLabelDetails = BaseParentLabelDetailsImpl{}

type BaseParentLabelDetailsImpl struct {
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

func (s BaseParentLabelDetailsImpl) ParentLabelDetails() BaseParentLabelDetailsImpl {
	return s
}

var _ ParentLabelDetails = RawParentLabelDetailsImpl{}

// RawParentLabelDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawParentLabelDetailsImpl struct {
	parentLabelDetails BaseParentLabelDetailsImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawParentLabelDetailsImpl) ParentLabelDetails() BaseParentLabelDetailsImpl {
	return s.parentLabelDetails
}

var _ json.Unmarshaler = &BaseParentLabelDetailsImpl{}

func (s *BaseParentLabelDetailsImpl) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling BaseParentLabelDetailsImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["parent"]; ok {
		impl, err := UnmarshalParentLabelDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Parent' for 'BaseParentLabelDetailsImpl': %+v", err)
		}
		s.Parent = impl
	}

	return nil
}

func UnmarshalParentLabelDetailsImplementation(input []byte) (ParentLabelDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ParentLabelDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.labelDetails") {
		var out LabelDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LabelDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseParentLabelDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseParentLabelDetailsImpl: %+v", err)
	}

	return RawParentLabelDetailsImpl{
		parentLabelDetails: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
