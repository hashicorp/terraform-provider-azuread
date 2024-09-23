package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceHealthScriptAssignment{}

type DeviceHealthScriptAssignment struct {
	// Determine whether we want to run detection script only or run both detection script and remediation script
	RunRemediationScript *bool `json:"runRemediationScript,omitempty"`

	// Script run schedule for the target group
	RunSchedule DeviceHealthScriptRunSchedule `json:"runSchedule"`

	// The Azure Active Directory group we are targeting the script to
	Target DeviceAndAppManagementAssignmentTarget `json:"target"`

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

func (s DeviceHealthScriptAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptAssignment{}

func (s DeviceHealthScriptAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceHealthScriptAssignment{}

func (s *DeviceHealthScriptAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		RunRemediationScript *bool   `json:"runRemediationScript,omitempty"`
		Id                   *string `json:"id,omitempty"`
		ODataId              *string `json:"@odata.id,omitempty"`
		ODataType            *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.RunRemediationScript = decoded.RunRemediationScript
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceHealthScriptAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["runSchedule"]; ok {
		impl, err := UnmarshalDeviceHealthScriptRunScheduleImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RunSchedule' for 'DeviceHealthScriptAssignment': %+v", err)
		}
		s.RunSchedule = impl
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'DeviceHealthScriptAssignment': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
