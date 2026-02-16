package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesGroupAssignment = WindowsUpdatesExcludedGroupAssignment{}

type WindowsUpdatesExcludedGroupAssignment struct {
	Assignments *[]WindowsUpdatesAssignedGroup `json:"assignments,omitempty"`

	// Fields inherited from WindowsUpdatesGroupAssignment

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesExcludedGroupAssignment) WindowsUpdatesGroupAssignment() BaseWindowsUpdatesGroupAssignmentImpl {
	return BaseWindowsUpdatesGroupAssignmentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesExcludedGroupAssignment{}

func (s WindowsUpdatesExcludedGroupAssignment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesExcludedGroupAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesExcludedGroupAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesExcludedGroupAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.excludedGroupAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesExcludedGroupAssignment: %+v", err)
	}

	return encoded, nil
}
