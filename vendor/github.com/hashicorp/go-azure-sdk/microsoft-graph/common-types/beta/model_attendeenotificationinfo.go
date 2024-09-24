package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeNotificationInfo struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The phone number of the external attendee. Required.
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`

	// The time zone of the external attendee. The timeZone property can be set to any of the time zones currently supported
	// by Windows. Required.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
