package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailabilityItem struct {
	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the service ID for 1:n appointments. If the appointment is of type 1:n, this field is present, otherwise,
	// null.
	ServiceId nullable.Type[string] `json:"serviceId,omitempty"`

	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

	// The status of the staff member. Possible values are: available, busy, slotsAvailable, outOfOffice,
	// unknownFutureValue.
	Status *BookingsAvailabilityStatus `json:"status,omitempty"`
}
