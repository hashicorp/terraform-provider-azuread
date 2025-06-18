package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkplaceSensorDeviceTelemetry struct {
	// The value of the sensor can be true or false. Use it for sensors that report binary values, such as occupancy or
	// heartbeat.
	BoolValue nullable.Type[bool] `json:"boolValue,omitempty"`

	// The user-defined unique identifier of the device provided at the time of creation. Don't use the system generated
	// identifier of the device.
	DeviceId *string `json:"deviceId,omitempty"`

	// The extra values associated with badge signals.
	EventValue *WorkplaceSensorEventValue `json:"eventValue,omitempty"`

	// The value of the sensor as an integer. Use it for sensors that report numerical values, such as people count.
	IntValue nullable.Type[int64] `json:"intValue,omitempty"`

	// The additional information to indicate the location of the device.
	LocationHint nullable.Type[string] `json:"locationHint,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The user-defined unique identifier of the sensor on the device. Optional. If the device has multiple sensors of the
	// same type, the property must be provided to identify each sensor. If the device has unique sensor types, the property
	// can be omitted. The default value is the sensor type.
	SensorId nullable.Type[string] `json:"sensorId,omitempty"`

	SensorType *WorkplaceSensorType `json:"sensorType,omitempty"`

	// The date and time when the sensor measured and reported its value. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	Timestamp *string `json:"timestamp,omitempty"`
}
