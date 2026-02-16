package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeTimeZone struct {
	// A single point of time in a combined date and time representation ({date}T{time}). For example,
	// '2019-04-16T09:00:00'.
	DateTime *string `json:"dateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
