package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceInstallState{}

type DeviceInstallState struct {
	// Device Id.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The error code for install failures.
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// Possible values for install state.
	InstallState *InstallState `json:"installState,omitempty"`

	// Last sync date and time.
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// OS Description.
	OsDescription nullable.Type[string] `json:"osDescription,omitempty"`

	// OS Version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Device User Name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

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

func (s DeviceInstallState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceInstallState{}

func (s DeviceInstallState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceInstallState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceInstallState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceInstallState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceInstallState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceInstallState: %+v", err)
	}

	return encoded, nil
}
