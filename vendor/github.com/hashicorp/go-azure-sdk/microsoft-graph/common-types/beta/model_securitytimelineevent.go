package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTimelineEvent struct {
	// The date and time when the event occurred.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// Additional details or context about the event.
	EventDetails nullable.Type[string] `json:"eventDetails,omitempty"`

	// The outcome or result of the event, such as delivery location or action taken.
	EventResult nullable.Type[string] `json:"eventResult,omitempty"`

	// The origin or actor that triggered the event. The possible values are: system, admin, user, unknownFutureValue.
	EventSource *SecurityEventSource `json:"eventSource,omitempty"`

	// Collection of threats identified or associated with this event.
	EventThreats *[]string `json:"eventThreats,omitempty"`

	// The type of event that occurred. The possible values are: originalDelivery, systemTimeTravel, dynamicDelivery,
	// userUrlClick, reprocessed, zap, quarantineRelease, air, unknown, unknownFutureValue.
	EventType *SecurityTimelineEventType `json:"eventType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
