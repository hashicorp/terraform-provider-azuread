package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerChecklistItem struct {
	// Value is true if the item is checked and false otherwise.
	IsChecked nullable.Type[bool] `json:"isChecked,omitempty"`

	// Read-only. User ID by which this is last modified.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Read-only. Date and time at which this is last modified. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Used to set the relative order of items in the checklist. The format is defined as outlined here.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

	// Title of the checklist item.
	Title nullable.Type[string] `json:"title,omitempty"`
}

var _ json.Marshaler = PlannerChecklistItem{}

func (s PlannerChecklistItem) MarshalJSON() ([]byte, error) {
	type wrapper PlannerChecklistItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerChecklistItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerChecklistItem: %+v", err)
	}

	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerChecklistItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerChecklistItem{}

func (s *PlannerChecklistItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsChecked            nullable.Type[bool]   `json:"isChecked,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
		OrderHint            nullable.Type[string] `json:"orderHint,omitempty"`
		Title                nullable.Type[string] `json:"title,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsChecked = decoded.IsChecked
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OrderHint = decoded.OrderHint
	s.Title = decoded.Title

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerChecklistItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PlannerChecklistItem': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
