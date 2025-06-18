package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesGroupAssignment = WindowsUpdatesIncludedGroupAssignment{}

type WindowsUpdatesIncludedGroupAssignment struct {
	Assignments *[]WindowsUpdatesAssignedGroup `json:"assignments,omitempty"`

	// Fields inherited from WindowsUpdatesGroupAssignment

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesIncludedGroupAssignment) WindowsUpdatesGroupAssignment() BaseWindowsUpdatesGroupAssignmentImpl {
	return BaseWindowsUpdatesGroupAssignmentImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesIncludedGroupAssignment{}

func (s WindowsUpdatesIncludedGroupAssignment) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesIncludedGroupAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesIncludedGroupAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesIncludedGroupAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.includedGroupAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesIncludedGroupAssignment: %+v", err)
	}

	return encoded, nil
}
