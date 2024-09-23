package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerArchivalInfo struct {
	// Read-only. Reason why the entity was archived or unarchived.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Read-only. Identity of the user who archived or unarchived the entity
	StatusChangedBy *IdentitySet `json:"statusChangedBy,omitempty"`

	// Read-only. Date and time at which the entity's archive status changed.
	StatusChangedDateTime nullable.Type[string] `json:"statusChangedDateTime,omitempty"`
}

var _ json.Marshaler = PlannerArchivalInfo{}

func (s PlannerArchivalInfo) MarshalJSON() ([]byte, error) {
	type wrapper PlannerArchivalInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerArchivalInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerArchivalInfo: %+v", err)
	}

	delete(decoded, "justification")
	delete(decoded, "statusChangedBy")
	delete(decoded, "statusChangedDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerArchivalInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerArchivalInfo{}

func (s *PlannerArchivalInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Justification         nullable.Type[string] `json:"justification,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
		StatusChangedDateTime nullable.Type[string] `json:"statusChangedDateTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Justification = decoded.Justification
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StatusChangedDateTime = decoded.StatusChangedDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerArchivalInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["statusChangedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StatusChangedBy' for 'PlannerArchivalInfo': %+v", err)
		}
		s.StatusChangedBy = &impl
	}

	return nil
}
