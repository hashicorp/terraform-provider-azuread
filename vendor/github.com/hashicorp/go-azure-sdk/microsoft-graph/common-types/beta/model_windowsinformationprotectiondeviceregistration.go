package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsInformationProtectionDeviceRegistration{}

type WindowsInformationProtectionDeviceRegistration struct {
	// Device Mac address.
	DeviceMacAddress nullable.Type[string] `json:"deviceMacAddress,omitempty"`

	// Device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Device identifier for this device registration record.
	DeviceRegistrationId nullable.Type[string] `json:"deviceRegistrationId,omitempty"`

	// Device type, for example, Windows laptop VS Windows phone.
	DeviceType nullable.Type[string] `json:"deviceType,omitempty"`

	// Last checkin time of the device.
	LastCheckInDateTime *string `json:"lastCheckInDateTime,omitempty"`

	// UserId associated with this device registration record.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s WindowsInformationProtectionDeviceRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsInformationProtectionDeviceRegistration{}

func (s WindowsInformationProtectionDeviceRegistration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionDeviceRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionDeviceRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionDeviceRegistration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionDeviceRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionDeviceRegistration: %+v", err)
	}

	return encoded, nil
}
