package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalReference struct {
	// A name alias to describe the reference.
	Alias nullable.Type[string] `json:"alias,omitempty"`

	// Read-only. User ID by which this is last modified.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Read-only. Date and time at which this is last modified. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Used to set the relative priority order in which the reference will be shown as a preview on the task.
	PreviewPriority nullable.Type[string] `json:"previewPriority,omitempty"`

	// Used to describe the type of the reference. Types include: PowerPoint, Word, Excel, Other.
	Type nullable.Type[string] `json:"type,omitempty"`
}

var _ json.Marshaler = PlannerExternalReference{}

func (s PlannerExternalReference) MarshalJSON() ([]byte, error) {
	type wrapper PlannerExternalReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerExternalReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerExternalReference: %+v", err)
	}

	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerExternalReference: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerExternalReference{}

func (s *PlannerExternalReference) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Alias                nullable.Type[string] `json:"alias,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
		PreviewPriority      nullable.Type[string] `json:"previewPriority,omitempty"`
		Type                 nullable.Type[string] `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Alias = decoded.Alias
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PreviewPriority = decoded.PreviewPriority
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerExternalReference into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'PlannerExternalReference': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
