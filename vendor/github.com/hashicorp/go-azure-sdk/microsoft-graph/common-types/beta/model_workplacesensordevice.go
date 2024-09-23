package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkplaceSensorDevice{}

type WorkplaceSensorDevice struct {
	// The description of the device.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The user-defined unique identifier of the device provided at the time of creation.
	DeviceId *string `json:"deviceId,omitempty"`

	// The display name of the device.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The IPv4 address of the device.
	IPV4Address nullable.Type[string] `json:"ipV4Address,omitempty"`

	// The IPv6 address of the device.
	IPV6Address nullable.Type[string] `json:"ipV6Address,omitempty"`

	// The MAC address of the device.
	MacAddress nullable.Type[string] `json:"macAddress,omitempty"`

	// The manufacturer of the device.
	Manufacturer *string `json:"manufacturer,omitempty"`

	// The unique identifier of the place where the device is located. If the device is installed in a room equipped with a
	// mailbox, this property should match the ExternalDirectoryObjectId or Microsoft Entra object ID of the room mailbox.
	PlaceId nullable.Type[string] `json:"placeId,omitempty"`

	// A list of sensors associated with the device that collect and report data about physical or environmental conditions,
	// such as occupancy, people count, inferred occupancy, temperature, and more.
	Sensors *[]WorkplaceSensor `json:"sensors,omitempty"`

	// A list of custom tags associated with the device.
	Tags *[]string `json:"tags,omitempty"`

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

func (s WorkplaceSensorDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkplaceSensorDevice{}

func (s WorkplaceSensorDevice) MarshalJSON() ([]byte, error) {
	type wrapper WorkplaceSensorDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkplaceSensorDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkplaceSensorDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workplaceSensorDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkplaceSensorDevice: %+v", err)
	}

	return encoded, nil
}
