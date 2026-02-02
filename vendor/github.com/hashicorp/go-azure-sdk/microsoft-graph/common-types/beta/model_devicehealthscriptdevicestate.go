package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceHealthScriptDeviceState{}

type DeviceHealthScriptDeviceState struct {
	// A list of the assignment filter ids used for health script applicability evaluation
	AssignmentFilterIds *[]string `json:"assignmentFilterIds,omitempty"`

	// Indicates the type of execution status of the device management script.
	DetectionState *RunState `json:"detectionState,omitempty"`

	// The next timestamp of when the device health script is expected to execute
	ExpectedStateUpdateDateTime nullable.Type[string] `json:"expectedStateUpdateDateTime,omitempty"`

	// The last timestamp of when the device health script executed
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// The last time that Intune Managment Extension synced with Intune
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// The managed device on which the device health script executed
	ManagedDevice *ManagedDevice `json:"managedDevice,omitempty"`

	// Error from the detection script after remediation
	PostRemediationDetectionScriptError nullable.Type[string] `json:"postRemediationDetectionScriptError,omitempty"`

	// Detection script output after remediation
	PostRemediationDetectionScriptOutput nullable.Type[string] `json:"postRemediationDetectionScriptOutput,omitempty"`

	// Error from the detection script before remediation
	PreRemediationDetectionScriptError nullable.Type[string] `json:"preRemediationDetectionScriptError,omitempty"`

	// Output of the detection script before remediation
	PreRemediationDetectionScriptOutput nullable.Type[string] `json:"preRemediationDetectionScriptOutput,omitempty"`

	// Error output of the remediation script
	RemediationScriptError nullable.Type[string] `json:"remediationScriptError,omitempty"`

	// Indicates the type of execution status of the device management script.
	RemediationState *RemediationState `json:"remediationState,omitempty"`

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

func (s DeviceHealthScriptDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptDeviceState{}

func (s DeviceHealthScriptDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptDeviceState: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceHealthScriptDeviceState{}

func (s *DeviceHealthScriptDeviceState) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignmentFilterIds                  *[]string             `json:"assignmentFilterIds,omitempty"`
		DetectionState                       *RunState             `json:"detectionState,omitempty"`
		ExpectedStateUpdateDateTime          nullable.Type[string] `json:"expectedStateUpdateDateTime,omitempty"`
		LastStateUpdateDateTime              *string               `json:"lastStateUpdateDateTime,omitempty"`
		LastSyncDateTime                     *string               `json:"lastSyncDateTime,omitempty"`
		PostRemediationDetectionScriptError  nullable.Type[string] `json:"postRemediationDetectionScriptError,omitempty"`
		PostRemediationDetectionScriptOutput nullable.Type[string] `json:"postRemediationDetectionScriptOutput,omitempty"`
		PreRemediationDetectionScriptError   nullable.Type[string] `json:"preRemediationDetectionScriptError,omitempty"`
		PreRemediationDetectionScriptOutput  nullable.Type[string] `json:"preRemediationDetectionScriptOutput,omitempty"`
		RemediationScriptError               nullable.Type[string] `json:"remediationScriptError,omitempty"`
		RemediationState                     *RemediationState     `json:"remediationState,omitempty"`
		Id                                   *string               `json:"id,omitempty"`
		ODataId                              *string               `json:"@odata.id,omitempty"`
		ODataType                            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignmentFilterIds = decoded.AssignmentFilterIds
	s.DetectionState = decoded.DetectionState
	s.ExpectedStateUpdateDateTime = decoded.ExpectedStateUpdateDateTime
	s.LastStateUpdateDateTime = decoded.LastStateUpdateDateTime
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.PostRemediationDetectionScriptError = decoded.PostRemediationDetectionScriptError
	s.PostRemediationDetectionScriptOutput = decoded.PostRemediationDetectionScriptOutput
	s.PreRemediationDetectionScriptError = decoded.PreRemediationDetectionScriptError
	s.PreRemediationDetectionScriptOutput = decoded.PreRemediationDetectionScriptOutput
	s.RemediationScriptError = decoded.RemediationScriptError
	s.RemediationState = decoded.RemediationState
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceHealthScriptDeviceState into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedDevice"]; ok {
		impl, err := UnmarshalManagedDeviceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ManagedDevice' for 'DeviceHealthScriptDeviceState': %+v", err)
		}
		s.ManagedDevice = &impl
	}

	return nil
}
