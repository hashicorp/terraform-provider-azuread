package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ImportedWindowsAutopilotDeviceIdentityUpload{}

type ImportedWindowsAutopilotDeviceIdentityUpload struct {
	// DateTime when the entity is created.
	CreatedDateTimeUtc *string `json:"createdDateTimeUtc,omitempty"`

	// Collection of all Autopilot devices as a part of this upload.
	DeviceIdentities *[]ImportedWindowsAutopilotDeviceIdentity `json:"deviceIdentities,omitempty"`

	Status *ImportedWindowsAutopilotDeviceIdentityUploadStatus `json:"status,omitempty"`

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

func (s ImportedWindowsAutopilotDeviceIdentityUpload) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ImportedWindowsAutopilotDeviceIdentityUpload{}

func (s ImportedWindowsAutopilotDeviceIdentityUpload) MarshalJSON() ([]byte, error) {
	type wrapper ImportedWindowsAutopilotDeviceIdentityUpload
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImportedWindowsAutopilotDeviceIdentityUpload: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportedWindowsAutopilotDeviceIdentityUpload: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.importedWindowsAutopilotDeviceIdentityUpload"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImportedWindowsAutopilotDeviceIdentityUpload: %+v", err)
	}

	return encoded, nil
}
