package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsManagementAppHealthState{}

type WindowsManagementAppHealthState struct {
	// Name of the device on which Windows management app is installed.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Windows 10 OS version of the device on which Windows management app is installed.
	DeviceOSVersion nullable.Type[string] `json:"deviceOSVersion,omitempty"`

	// Indicates health state of the Windows management app.
	HealthState *HealthState `json:"healthState,omitempty"`

	// Windows management app installed version.
	InstalledVersion nullable.Type[string] `json:"installedVersion,omitempty"`

	// Windows management app last check-in time.
	LastCheckInDateTime *string `json:"lastCheckInDateTime,omitempty"`

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

func (s WindowsManagementAppHealthState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsManagementAppHealthState{}

func (s WindowsManagementAppHealthState) MarshalJSON() ([]byte, error) {
	type wrapper WindowsManagementAppHealthState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsManagementAppHealthState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsManagementAppHealthState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsManagementAppHealthState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsManagementAppHealthState: %+v", err)
	}

	return encoded, nil
}
