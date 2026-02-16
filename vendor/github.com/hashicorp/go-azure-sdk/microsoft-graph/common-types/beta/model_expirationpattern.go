package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationPattern struct {
	// The requestor's desired duration of access represented in ISO 8601 format for durations. For example, PT3H refers to
	// three hours. If specified in a request, endDateTime shouldn't be present and the type property should be set to
	// afterDuration.
	Duration nullable.Type[string] `json:"duration,omitempty"`

	// Timestamp of date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014, is 2014-01-01T00:00:00Z.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The requestor's desired expiration pattern type.
	Type *ExpirationPatternType `json:"type,omitempty"`
}
