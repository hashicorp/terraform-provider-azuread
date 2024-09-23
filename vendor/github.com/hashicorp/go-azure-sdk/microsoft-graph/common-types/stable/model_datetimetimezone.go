package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeTimeZone struct {
	// A single point of time in a combined date and time representation ({date}T{time}; for example,
	// 2017-08-29T04:00:00.0000000).
	DateTime *string `json:"dateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
