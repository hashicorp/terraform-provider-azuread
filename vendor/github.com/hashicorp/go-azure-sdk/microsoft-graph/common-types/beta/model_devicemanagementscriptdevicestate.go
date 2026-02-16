package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementScriptDeviceState{}

type DeviceManagementScriptDeviceState struct {
	// Error code corresponding to erroneous execution of the device management script.
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Error description corresponding to erroneous execution of the device management script.
	ErrorDescription nullable.Type[string] `json:"errorDescription,omitempty"`

	// Latest time the device management script executes.
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// The managed devices that executes the device management script.
	ManagedDevice *ManagedDevice `json:"managedDevice,omitempty"`

	// Details of execution output.
	ResultMessage nullable.Type[string] `json:"resultMessage,omitempty"`

	// Indicates the type of execution status of the device management script.
	RunState *RunState `json:"runState,omitempty"`

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

func (s DeviceManagementScriptDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementScriptDeviceState{}

func (s DeviceManagementScriptDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementScriptDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementScriptDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementScriptDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementScriptDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementScriptDeviceState: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementScriptDeviceState{}

func (s *DeviceManagementScriptDeviceState) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ErrorCode               *int64                `json:"errorCode,omitempty"`
		ErrorDescription        nullable.Type[string] `json:"errorDescription,omitempty"`
		LastStateUpdateDateTime *string               `json:"lastStateUpdateDateTime,omitempty"`
		ResultMessage           nullable.Type[string] `json:"resultMessage,omitempty"`
		RunState                *RunState             `json:"runState,omitempty"`
		Id                      *string               `json:"id,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ErrorCode = decoded.ErrorCode
	s.ErrorDescription = decoded.ErrorDescription
	s.LastStateUpdateDateTime = decoded.LastStateUpdateDateTime
	s.ResultMessage = decoded.ResultMessage
	s.RunState = decoded.RunState
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementScriptDeviceState into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["managedDevice"]; ok {
		impl, err := UnmarshalManagedDeviceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ManagedDevice' for 'DeviceManagementScriptDeviceState': %+v", err)
		}
		s.ManagedDevice = &impl
	}

	return nil
}
