package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendanceInterval struct {
	// Duration of the meeting interval in seconds; that is, the difference between joinDateTime and leaveDateTime.
	DurationInSeconds nullable.Type[int64] `json:"durationInSeconds,omitempty"`

	// The time the attendee joined in UTC.
	JoinDateTime nullable.Type[string] `json:"joinDateTime,omitempty"`

	// The time the attendee left in UTC.
	LeaveDateTime nullable.Type[string] `json:"leaveDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
