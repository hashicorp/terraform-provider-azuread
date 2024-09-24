package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftActivity struct {
	// Customer defined code for the shiftActivity. Required.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The name of the shiftActivity. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The end date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Indicates whether the microsoft.graph.user should be paid for the activity during their shift. Required.
	IsPaid nullable.Type[bool] `json:"isPaid,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The start date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}
