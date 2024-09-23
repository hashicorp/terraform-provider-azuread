package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwarePasswordDetail{}

type HardwarePasswordDetail struct {
	// The current device's BIOS password. Supports: $filter, $select, $top, $OrderBy, $skip. This property is read-only.
	CurrentPassword nullable.Type[string] `json:"currentPassword,omitempty"`

	// The list of all the previous device BIOS passwords. Supports: $filter, $select, $top, $skip. This property is
	// read-only.
	PreviousPasswords *[]string `json:"previousPasswords,omitempty"`

	// The device serial number as defined by the device manufacturer. Supports: $filter, $select, $top, $OrderBy, $skip.
	// This property is read-only.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

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

func (s HardwarePasswordDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwarePasswordDetail{}

func (s HardwarePasswordDetail) MarshalJSON() ([]byte, error) {
	type wrapper HardwarePasswordDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwarePasswordDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwarePasswordDetail: %+v", err)
	}

	delete(decoded, "currentPassword")
	delete(decoded, "previousPasswords")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwarePasswordDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwarePasswordDetail: %+v", err)
	}

	return encoded, nil
}
