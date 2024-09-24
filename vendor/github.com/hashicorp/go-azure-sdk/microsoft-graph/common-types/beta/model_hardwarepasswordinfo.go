package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwarePasswordInfo{}

type HardwarePasswordInfo struct {
	// Current device password. This property is read-only.
	CurrentPassword nullable.Type[string] `json:"currentPassword,omitempty"`

	// List of previous device passwords. This property is read-only.
	PreviousPasswords *[]string `json:"previousPasswords,omitempty"`

	// Associated device's serial number . This property is read-only.
	SerialNumber *string `json:"serialNumber,omitempty"`

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

func (s HardwarePasswordInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwarePasswordInfo{}

func (s HardwarePasswordInfo) MarshalJSON() ([]byte, error) {
	type wrapper HardwarePasswordInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwarePasswordInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwarePasswordInfo: %+v", err)
	}

	delete(decoded, "currentPassword")
	delete(decoded, "previousPasswords")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwarePasswordInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwarePasswordInfo: %+v", err)
	}

	return encoded, nil
}
