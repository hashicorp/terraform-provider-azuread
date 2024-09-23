package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PresenceStatusMessage struct {
	// Time in which the status message expires.If not provided, the status message doesn't expire.expiryDateTime.dateTime
	// shouldn't include time zone.expiryDateTime isn't available when you request the presence of another user.
	ExpiryDateTime *DateTimeTimeZone `json:"expiryDateTime,omitempty"`

	// Status message item. The only supported format currently is message.contentType = 'text'.
	Message *ItemBody `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time in which the status message was published.Read-only.publishedDateTime isn't available when you request the
	// presence of another user.
	PublishedDateTime nullable.Type[string] `json:"publishedDateTime,omitempty"`
}
