package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleInformation struct {
	// Represents a merged view of availability of all the items in scheduleItems. The view consists of time slots.
	// Availability during each time slot is indicated with: 0= free or working elswhere, 1= tentative, 2= busy, 3= out of
	// office.Note: Working elsewhere is set to 0 instead of 4 for backward compatibility. For details, see the Q&A and
	// Exchange 2007 and Exchange 2010 do not use the WorkingElsewhere value.
	AvailabilityView nullable.Type[string] `json:"availabilityView,omitempty"`

	// Error information from attempting to get the availability of the user, distribution list, or resource.
	Error *FreeBusyError `json:"error,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An SMTP address of the user, distribution list, or resource, identifying an instance of scheduleInformation.
	ScheduleId nullable.Type[string] `json:"scheduleId,omitempty"`

	// Contains the items that describe the availability of the user or resource.
	ScheduleItems *[]ScheduleItem `json:"scheduleItems,omitempty"`

	// The days of the week and hours in a specific time zone that the user works. These are set as part of the user's
	// mailboxSettings.
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}
