package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HorizontalSectionColumn{}

type HorizontalSectionColumn struct {
	// The collection of WebParts in this column.
	Webparts *[]WebPart `json:"webparts,omitempty"`

	// Width of the column. A horizontal section is divided into 12 grids. A column should have a value of 1-12 to represent
	// its range spans. For example, there can be two columns both have a width of 6 in a section.
	Width nullable.Type[int64] `json:"width,omitempty"`

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

func (s HorizontalSectionColumn) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HorizontalSectionColumn{}

func (s HorizontalSectionColumn) MarshalJSON() ([]byte, error) {
	type wrapper HorizontalSectionColumn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HorizontalSectionColumn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HorizontalSectionColumn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.horizontalSectionColumn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HorizontalSectionColumn: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &HorizontalSectionColumn{}

func (s *HorizontalSectionColumn) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Width     nullable.Type[int64] `json:"width,omitempty"`
		Id        *string              `json:"id,omitempty"`
		ODataId   *string              `json:"@odata.id,omitempty"`
		ODataType *string              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Width = decoded.Width
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HorizontalSectionColumn into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["webparts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Webparts into list []json.RawMessage: %+v", err)
		}

		output := make([]WebPart, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWebPartImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Webparts' for 'HorizontalSectionColumn': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Webparts = &output
	}

	return nil
}
