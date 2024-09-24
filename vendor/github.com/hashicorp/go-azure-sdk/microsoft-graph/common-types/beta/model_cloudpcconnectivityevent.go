package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEvent struct {
	// Indicates the date and time when this event was created. The timestamp is shown in ISO 8601 format and Coordinated
	// Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// Name of the event.
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	EventResult *CloudPCConnectivityEventResult `json:"eventResult,omitempty"`
	EventType   *CloudPCConnectivityEventType   `json:"eventType,omitempty"`

	// Additional message for this event.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
