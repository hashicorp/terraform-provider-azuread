package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkOnlineMeetingInfo struct {
	// The identifier of the calendar event associated with the meeting.
	CalendarEventId nullable.Type[string] `json:"calendarEventId,omitempty"`

	// The URL which can be clicked on to join or uniquely identify the meeting.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The organizer of the meeting.
	Organizer *TeamworkUserIdentity `json:"organizer,omitempty"`
}
