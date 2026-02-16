package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsInformationProtectionAppLockerFile{}

type WindowsInformationProtectionAppLockerFile struct {
	// The friendly name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// File as a byte array
	File nullable.Type[string] `json:"file,omitempty"`

	// SHA256 hash of the file
	FileHash nullable.Type[string] `json:"fileHash,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s WindowsInformationProtectionAppLockerFile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsInformationProtectionAppLockerFile{}

func (s WindowsInformationProtectionAppLockerFile) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionAppLockerFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionAppLockerFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionAppLockerFile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionAppLockerFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionAppLockerFile: %+v", err)
	}

	return encoded, nil
}
