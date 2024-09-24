package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkplaceSensor struct {
	// The display name of the sensor. Optional.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of the place that the sensor detects. If the device is installed in a room equipped with a
	// mailbox, this property should match the ExternalDirectoryObjectId or Microsoft Entra object ID of the room mailbox.
	// If the sensor detects the same place as the location of the device, the property can be omitted. The default value is
	// the place identifier of the device. Optional.
	PlaceId nullable.Type[string] `json:"placeId,omitempty"`

	// The user-defined unique identifier of the sensor on the device. If the device has multiple sensors of the same type,
	// the property must be provided to identify each sensor. If the device has only one sensor of a type, the property can
	// be omitted. The default value is the sensor type. Optional.
	SensorId nullable.Type[string] `json:"sensorId,omitempty"`

	SensorType *WorkplaceSensorType `json:"sensorType,omitempty"`
}
