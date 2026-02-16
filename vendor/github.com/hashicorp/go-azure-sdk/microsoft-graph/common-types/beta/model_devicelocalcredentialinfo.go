package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceLocalCredentialInfo{}

type DeviceLocalCredentialInfo struct {
	// The credentials of the device's local administrator account backed up to Azure Active Directory.
	Credentials *[]DeviceLocalCredential `json:"credentials,omitempty"`

	// Display name of the device that the local credentials are associated with.
	DeviceName *string `json:"deviceName,omitempty"`

	// When the local administrator account credential was backed up to Microsoft Entra ID.
	LastBackupDateTime *string `json:"lastBackupDateTime,omitempty"`

	// When the local administrator account credential will be refreshed and backed up to Microsoft Entra ID.
	RefreshDateTime *string `json:"refreshDateTime,omitempty"`

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

func (s DeviceLocalCredentialInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceLocalCredentialInfo{}

func (s DeviceLocalCredentialInfo) MarshalJSON() ([]byte, error) {
	type wrapper DeviceLocalCredentialInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceLocalCredentialInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceLocalCredentialInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceLocalCredentialInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceLocalCredentialInfo: %+v", err)
	}

	return encoded, nil
}
