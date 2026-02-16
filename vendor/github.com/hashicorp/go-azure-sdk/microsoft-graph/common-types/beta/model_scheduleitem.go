package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleItem struct {
	// The date, time, and time zone that the corresponding event ends.
	End *DateTimeTimeZone `json:"end,omitempty"`

	// The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
	IsPrivate nullable.Type[bool] `json:"isPrivate,omitempty"`

	// The location where the corresponding event is held or attended from. Optional.
	Location nullable.Type[string] `json:"location,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date, time, and time zone that the corresponding event starts.
	Start *DateTimeTimeZone `json:"start,omitempty"`

	// The availability status of the user or resource during the corresponding event. The possible values are: free,
	// tentative, busy, oof, workingElsewhere, unknown.
	Status *FreeBusyStatus `json:"status,omitempty"`

	// The corresponding event's subject line. Optional.
	Subject nullable.Type[string] `json:"subject,omitempty"`
}
