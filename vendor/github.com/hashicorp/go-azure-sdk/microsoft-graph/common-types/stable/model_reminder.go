package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reminder struct {
	// Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows
	// Exchange to apply changes to the correct version of the object.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The date, time and time zone that the event ends.
	EventEndTime *DateTimeTimeZone `json:"eventEndTime,omitempty"`

	// The unique ID of the event. Read only.
	EventId nullable.Type[string] `json:"eventId,omitempty"`

	// The location of the event.
	EventLocation Location `json:"eventLocation"`

	// The date, time, and time zone that the event starts.
	EventStartTime *DateTimeTimeZone `json:"eventStartTime,omitempty"`

	// The text of the event's subject line.
	EventSubject nullable.Type[string] `json:"eventSubject,omitempty"`

	// The URL to open the event in Outlook on the web.The event opens in the browser if you're logged in to your mailbox
	// via Outlook on the web. You're prompted to log in if you aren't already logged in with the browser.This URL can't be
	// accessed from within an iFrame.
	EventWebLink nullable.Type[string] `json:"eventWebLink,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date, time, and time zone that the reminder is set to occur.
	ReminderFireTime *DateTimeTimeZone `json:"reminderFireTime,omitempty"`
}

var _ json.Unmarshaler = &Reminder{}

func (s *Reminder) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChangeKey        nullable.Type[string] `json:"changeKey,omitempty"`
		EventEndTime     *DateTimeTimeZone     `json:"eventEndTime,omitempty"`
		EventId          nullable.Type[string] `json:"eventId,omitempty"`
		EventStartTime   *DateTimeTimeZone     `json:"eventStartTime,omitempty"`
		EventSubject     nullable.Type[string] `json:"eventSubject,omitempty"`
		EventWebLink     nullable.Type[string] `json:"eventWebLink,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
		ReminderFireTime *DateTimeTimeZone     `json:"reminderFireTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChangeKey = decoded.ChangeKey
	s.EventEndTime = decoded.EventEndTime
	s.EventId = decoded.EventId
	s.EventStartTime = decoded.EventStartTime
	s.EventSubject = decoded.EventSubject
	s.EventWebLink = decoded.EventWebLink
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReminderFireTime = decoded.ReminderFireTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Reminder into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["eventLocation"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EventLocation' for 'Reminder': %+v", err)
		}
		s.EventLocation = impl
	}

	return nil
}
