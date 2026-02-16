package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceLocalCredential{}

type DeviceLocalCredential struct {
	// The name of the local admin account for which LAPS is enabled.
	AccountName *string `json:"accountName,omitempty"`

	// The SID of the local admin account for which LAPS is enabled.
	AccountSid *string `json:"accountSid,omitempty"`

	// When the local administrator account credential for the device object was backed up to Azure Active Directory.
	BackupDateTime *string `json:"backupDateTime,omitempty"`

	// The password for the local administrator account that is backed up to Azure Active Directory and returned as a Base64
	// encoded value.
	PasswordBase64 *string `json:"passwordBase64,omitempty"`

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

func (s DeviceLocalCredential) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceLocalCredential{}

func (s DeviceLocalCredential) MarshalJSON() ([]byte, error) {
	type wrapper DeviceLocalCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceLocalCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceLocalCredential: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceLocalCredential"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceLocalCredential: %+v", err)
	}

	return encoded, nil
}
