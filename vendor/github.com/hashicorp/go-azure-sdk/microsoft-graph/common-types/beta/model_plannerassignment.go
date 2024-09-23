package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerAssignment struct {
	// The identity of the user that performed the assignment of the task, that is, the assignor.
	AssignedBy IdentitySet `json:"assignedBy"`

	// The time at which the task was assigned. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	AssignedDateTime nullable.Type[string] `json:"assignedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Hint used to order assignees in a task. The format is defined as outlined here.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`
}

var _ json.Unmarshaler = &PlannerAssignment{}

func (s *PlannerAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedDateTime nullable.Type[string] `json:"assignedDateTime,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		OrderHint        nullable.Type[string] `json:"orderHint,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedDateTime = decoded.AssignedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OrderHint = decoded.OrderHint

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignedBy' for 'PlannerAssignment': %+v", err)
		}
		s.AssignedBy = impl
	}

	return nil
}
