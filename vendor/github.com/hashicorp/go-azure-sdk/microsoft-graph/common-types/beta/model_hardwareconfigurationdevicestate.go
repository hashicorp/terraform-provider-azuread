package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwareConfigurationDeviceState{}

type HardwareConfigurationDeviceState struct {
	// A list of identifier strings of different assignment filters applied
	AssignmentFilterIds nullable.Type[string] `json:"assignmentFilterIds,omitempty"`

	// Error from the hardware configuration execution
	ConfigurationError nullable.Type[string] `json:"configurationError,omitempty"`

	// Output of the hardware configuration execution
	ConfigurationOutput nullable.Type[string] `json:"configurationOutput,omitempty"`

	// Indicates the type of execution status of the device management script.
	ConfigurationState *RunState `json:"configurationState,omitempty"`

	// The name of the device
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The Policy internal version
	InternalVersion *int64 `json:"internalVersion,omitempty"`

	// The last timestamp of when the hardware configuration executed
	LastStateUpdateDateTime *string `json:"lastStateUpdateDateTime,omitempty"`

	// Operating system version of the device (E.g. 10.0.19042.1165, 10.0.19042.1288 etc.)
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// User Principal Name (UPN).
	Upn nullable.Type[string] `json:"upn,omitempty"`

	// The unique identifier of the Entra user associated with the device for which policy is applied. Read-Only.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s HardwareConfigurationDeviceState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwareConfigurationDeviceState{}

func (s HardwareConfigurationDeviceState) MarshalJSON() ([]byte, error) {
	type wrapper HardwareConfigurationDeviceState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareConfigurationDeviceState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareConfigurationDeviceState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareConfigurationDeviceState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareConfigurationDeviceState: %+v", err)
	}

	return encoded, nil
}
