package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementScriptUserState{}

type DeviceManagementScriptUserState struct {
	// List of run states for this script across all devices of specific user.
	DeviceRunStates *[]DeviceManagementScriptDeviceState `json:"deviceRunStates,omitempty"`

	// Error device count for specific user.
	ErrorDeviceCount *int64 `json:"errorDeviceCount,omitempty"`

	// Success device count for specific user.
	SuccessDeviceCount *int64 `json:"successDeviceCount,omitempty"`

	// User principle name of specific user.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s DeviceManagementScriptUserState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementScriptUserState{}

func (s DeviceManagementScriptUserState) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementScriptUserState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementScriptUserState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementScriptUserState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementScriptUserState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementScriptUserState: %+v", err)
	}

	return encoded, nil
}
