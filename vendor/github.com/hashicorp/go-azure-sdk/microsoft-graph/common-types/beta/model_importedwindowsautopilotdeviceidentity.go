package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ImportedWindowsAutopilotDeviceIdentity{}

type ImportedWindowsAutopilotDeviceIdentity struct {
	// UPN of the user the device will be assigned
	AssignedUserPrincipalName nullable.Type[string] `json:"assignedUserPrincipalName,omitempty"`

	// Group Tag of the Windows autopilot device.
	GroupTag nullable.Type[string] `json:"groupTag,omitempty"`

	// Hardware Blob of the Windows autopilot device.
	HardwareIdentifier nullable.Type[string] `json:"hardwareIdentifier,omitempty"`

	// The Import Id of the Windows autopilot device.
	ImportId nullable.Type[string] `json:"importId,omitempty"`

	// Product Key of the Windows autopilot device.
	ProductKey nullable.Type[string] `json:"productKey,omitempty"`

	// Serial number of the Windows autopilot device.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// Current state of the imported device.
	State *ImportedWindowsAutopilotDeviceIdentityState `json:"state,omitempty"`

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

func (s ImportedWindowsAutopilotDeviceIdentity) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ImportedWindowsAutopilotDeviceIdentity{}

func (s ImportedWindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	type wrapper ImportedWindowsAutopilotDeviceIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImportedWindowsAutopilotDeviceIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedWindowsAutopilotDeviceIdentity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedWindowsAutopilotDeviceIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImportedWindowsAutopilotDeviceIdentity: %+v", err)
	}

	return encoded, nil
}
