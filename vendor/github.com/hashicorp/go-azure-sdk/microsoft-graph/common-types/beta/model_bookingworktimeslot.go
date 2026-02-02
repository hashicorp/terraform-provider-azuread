package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingWorkTimeSlot struct {
	// The time of the day when work stops. For example, 17:00:00.0000000.
	End *string `json:"end,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time of the day when work starts. For example, 08:00:00.0000000.
	Start *string `json:"start,omitempty"`
}
