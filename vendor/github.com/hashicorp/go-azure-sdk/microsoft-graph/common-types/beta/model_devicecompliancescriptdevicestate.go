package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceComplianceScriptDeviceState{}

type DeviceComplianceScriptDeviceState struct {
	// Indicates the type of execution status of the device management script.
	DetectionState *RunState `json:"detectionState,omitempty"`

	// The next timestamp of when the device compliance script is expected to execute
	ExpectedStateUpdateDateTime nullable.Type[string] `json:"expectedStateUpdateDateTime,omitempty"`

	// The last timestamp of when the device compliance script executed
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// The last time that Intune Managment Extension synced with Intune
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// The managed device on which the device compliance script executed
	ManagedDevice *ManagedDevice `json:"managedDevice,omitempty"`

	// Error from the detection script
	ScriptError nullable.Type[string] `json:"scriptError,omitempty"`

	// Output of the detection script
	ScriptOutput nullable.Type[string] `json:"scriptOutput,omitempty"`

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

func (s DeviceComplianceScriptDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceComplianceScriptDeviceState{}

func (s DeviceComplianceScriptDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceComplianceScriptDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceComplianceScriptDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScriptDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceComplianceScriptDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceComplianceScriptDeviceState: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceComplianceScriptDeviceState{}

func (s *DeviceComplianceScriptDeviceState) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DetectionState              *RunState             `json:"detectionState,omitempty"`
		ExpectedStateUpdateDateTime nullable.Type[string] `json:"expectedStateUpdateDateTime,omitempty"`
		LastStateUpdateDateTime     *string               `json:"lastStateUpdateDateTime,omitempty"`
		LastSyncDateTime            *string               `json:"lastSyncDateTime,omitempty"`
		ScriptError                 nullable.Type[string] `json:"scriptError,omitempty"`
		ScriptOutput                nullable.Type[string] `json:"scriptOutput,omitempty"`
		Id                          *string               `json:"id,omitempty"`
		ODataId                     *string               `json:"@odata.id,omitempty"`
		ODataType                   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DetectionState = decoded.DetectionState
	s.ExpectedStateUpdateDateTime = decoded.ExpectedStateUpdateDateTime
	s.LastStateUpdateDateTime = decoded.LastStateUpdateDateTime
	s.LastSyncDateTime = decoded.LastSyncDateTime
	s.ScriptError = decoded.ScriptError
	s.ScriptOutput = decoded.ScriptOutput
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceComplianceScriptDeviceState into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedDevice"]; ok {
		impl, err := UnmarshalManagedDeviceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ManagedDevice' for 'DeviceComplianceScriptDeviceState': %+v", err)
		}
		s.ManagedDevice = &impl
	}

	return nil
}
