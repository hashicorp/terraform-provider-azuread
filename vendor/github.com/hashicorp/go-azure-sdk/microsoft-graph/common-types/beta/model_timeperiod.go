package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimePeriod struct {
	// The date time of the end of the time period.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date time of the start of the time period.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
