package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSchedule struct {
	// Date and time when this job will expire. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	Expiration nullable.Type[string] `json:"expiration,omitempty"`

	// The interval between synchronization iterations. The value is represented in ISO 8601 format for durations. For
	// example, P1M represents a period of one month and PT1M represents a period of one minute.
	Interval *string `json:"interval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *SynchronizationScheduleState `json:"state,omitempty"`
}
