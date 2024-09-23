package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestSchedule struct {
	// In entitlement management, when the access should expire.
	Expiration *ExpirationPattern `json:"expiration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For recurring access, or eligible or active assignment. This property is currently unsupported in both PIM and
	// entitlement management.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. In PIM, when the eligible or active assignment becomes active.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
