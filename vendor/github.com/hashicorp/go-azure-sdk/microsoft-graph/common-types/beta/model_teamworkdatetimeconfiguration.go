package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDateTimeConfiguration struct {
	// The date format for the device.
	DateFormat nullable.Type[string] `json:"dateFormat,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time of the day when the device is turned off.
	OfficeHoursEndTime nullable.Type[string] `json:"officeHoursEndTime,omitempty"`

	// The time of the day when the device is turned on.
	OfficeHoursStartTime nullable.Type[string] `json:"officeHoursStartTime,omitempty"`

	// The time format for the device.
	TimeFormat nullable.Type[string] `json:"timeFormat,omitempty"`

	// The time zone to which the office hours apply.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
