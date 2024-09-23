package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsExpiration struct {
	// Represents the frequency at which the terms will expire, after its first expiration as set in startDateTime. The
	// value is represented in ISO 8601 format for durations. For example, PT1M represents a time period of one month.
	Frequency nullable.Type[string] `json:"frequency,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The DateTime when the agreement is set to expire for all users. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
