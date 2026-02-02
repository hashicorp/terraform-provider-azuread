package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingReminder struct {
	// The message in the reminder.
	Message *string `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601
	// format.
	Offset *string `json:"offset,omitempty"`

	Recipients *BookingReminderRecipients `json:"recipients,omitempty"`
}
