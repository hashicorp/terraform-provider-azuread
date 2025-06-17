package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesGroupAssignment interface {
	WindowsUpdatesGroupAssignment() BaseWindowsUpdatesGroupAssignmentImpl
}

var _ WindowsUpdatesGroupAssignment = BaseWindowsUpdatesGroupAssignmentImpl{}

type BaseWindowsUpdatesGroupAssignmentImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesGroupAssignmentImpl) WindowsUpdatesGroupAssignment() BaseWindowsUpdatesGroupAssignmentImpl {
	return s
}

var _ WindowsUpdatesGroupAssignment = RawWindowsUpdatesGroupAssignmentImpl{}

// RawWindowsUpdatesGroupAssignmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesGroupAssignmentImpl struct {
	windowsUpdatesGroupAssignment BaseWindowsUpdatesGroupAssignmentImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawWindowsUpdatesGroupAssignmentImpl) WindowsUpdatesGroupAssignment() BaseWindowsUpdatesGroupAssignmentImpl {
	return s.windowsUpdatesGroupAssignment
}

func UnmarshalWindowsUpdatesGroupAssignmentImplementation(input []byte) (WindowsUpdatesGroupAssignment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesGroupAssignment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.excludedGroupAssignment") {
		var out WindowsUpdatesExcludedGroupAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesExcludedGroupAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.includedGroupAssignment") {
		var out WindowsUpdatesIncludedGroupAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesIncludedGroupAssignment: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesGroupAssignmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesGroupAssignmentImpl: %+v", err)
	}

	return RawWindowsUpdatesGroupAssignmentImpl{
		windowsUpdatesGroupAssignment: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
