package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OutlookItem = Event{}

type Event struct {
	// true if the meeting organizer allows invitees to propose a new time when responding; otherwise false. Optional.
	// Default is true.
	AllowNewTimeProposals nullable.Type[bool] `json:"allowNewTimeProposals,omitempty"`

	// The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation
	// property. Read-only. Nullable.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// The collection of attendees for the event.
	Attendees *[]Attendee `json:"attendees,omitempty"`

	// The body of the message associated with the event. It can be in HTML or text format.
	Body *ItemBody `json:"body,omitempty"`

	// The preview of the message associated with the event. It is in text format.
	BodyPreview nullable.Type[string] `json:"bodyPreview,omitempty"`

	// The calendar that contains the event. Navigation property. Read-only.
	Calendar *Calendar `json:"calendar,omitempty"`

	// Contains occurrenceId property values of canceled instances in a recurring series, if the event is the series master.
	// Instances in a recurring series that are canceled are called cancelledOccurences.Returned only on $select in a Get
	// operation which specifies the id of a series master event (that is, the seriesMasterId property value).
	CancelledOccurrences *[]string `json:"cancelledOccurrences,omitempty"`

	// The date, time, and time zone that the event ends. By default, the end time is in UTC.
	End *DateTimeTimeZone `json:"end,omitempty"`

	ExceptionOccurrences *[]Event `json:"exceptionOccurrences,omitempty"`

	// The collection of open extensions defined for the event. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// Set to true if the event has attachments.
	HasAttachments nullable.Type[bool] `json:"hasAttachments,omitempty"`

	// When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is
	// false.
	HideAttendees nullable.Type[bool] `json:"hideAttendees,omitempty"`

	// A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series.
	// Read-only.
	ICalUId nullable.Type[string] `json:"iCalUId,omitempty"`

	// The importance of the event. Possible values are: low, normal, high.
	Importance *Importance `json:"importance,omitempty"`

	// The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are
	// part of the recurrence pattern and exceptions that have been modified. It doesn't include occurrences that have been
	// canceled from the series. Navigation property. Read-only. Nullable.
	Instances *[]Event `json:"instances,omitempty"`

	// Set to true if the event lasts all day. If true, regardless of whether it's a single-day or multi-day event, start
	// and end time must be set to midnight and be in the same time zone.
	IsAllDay nullable.Type[bool] `json:"isAllDay,omitempty"`

	// Set to true if the event has been canceled.
	IsCancelled nullable.Type[bool] `json:"isCancelled,omitempty"`

	// Set to true if the user has updated the meeting in Outlook but hasn't sent the updates to attendees. Set to false if
	// all changes have been sent, or if the event is an appointment without any attendees.
	IsDraft nullable.Type[bool] `json:"isDraft,omitempty"`

	// True if this event has online meeting information (that is, onlineMeeting points to an onlineMeetingInfo resource),
	// false otherwise. Default is false (onlineMeeting is null). Optional. After you set isOnlineMeeting to true, Microsoft
	// Graph initializes onlineMeeting. Outlook then ignores any further changes to isOnlineMeeting, and the meeting remains
	// available online.
	IsOnlineMeeting nullable.Type[bool] `json:"isOnlineMeeting,omitempty"`

	// Set to true if the calendar owner (specified by the owner property of the calendar) is the organizer of the event
	// (specified by the organizer property of the event). This also applies if a delegate organized the event on behalf of
	// the owner.
	IsOrganizer nullable.Type[bool] `json:"isOrganizer,omitempty"`

	// Set to true if an alert is set to remind the user of the event.
	IsReminderOn nullable.Type[bool] `json:"isReminderOn,omitempty"`

	// The location of the event.
	Location Location `json:"location"`

	// The locations where the event is held or attended from. The location and locations properties always correspond with
	// each other. If you update the location property, any prior locations in the locations collection would be removed and
	// replaced by the new location value.
	Locations *[]Location `json:"locations,omitempty"`

	// The collection of multi-value extended properties defined for the event. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// An identifier for an occurrence in a recurring event series. Null if the event isn't part of a recurring series.The
	// format of the property value is OID.{seriesMasterId-value}.{occurrence-start-date}. The time zone for
	// {occurrence-start-date} is the recurrenceTimeZone property defined for the corresponding recurrenceRange.This
	// property can identify any occurrence in a recurring series, including an occurrence that has been modified or
	// canceled. You can use this property to perform all operations supported by occurrences in the recurring series.
	OccurrenceId nullable.Type[string] `json:"occurrenceId,omitempty"`

	// Details for an attendee to join the meeting online. Default is null. Read-only. After you set the isOnlineMeeting and
	// onlineMeetingProvider properties to enable a meeting online, Microsoft Graph initializes onlineMeeting. When set, the
	// meeting remains available online, and you can't change the isOnlineMeeting, onlineMeetingProvider, and onlneMeeting
	// properties again.
	OnlineMeeting *OnlineMeetingInfo `json:"onlineMeeting,omitempty"`

	// Represents the online meeting service provider. By default, onlineMeetingProvider is unknown. The possible values are
	// unknown, teamsForBusiness, skypeForBusiness, and skypeForConsumer. Optional. After you set onlineMeetingProvider,
	// Microsoft Graph initializes onlineMeeting. Subsequently you can't change onlineMeetingProvider again, and the meeting
	// remains available online.
	OnlineMeetingProvider *OnlineMeetingProviderType `json:"onlineMeetingProvider,omitempty"`

	// A URL for an online meeting. The property is set only when an organizer specifies in Outlook that an event is an
	// online meeting such as Skype. Read-only.To access the URL to join an online meeting, use joinUrl which is exposed via
	// the onlineMeeting property of the event. The onlineMeetingUrl property will be deprecated in the future.
	OnlineMeetingUrl nullable.Type[string] `json:"onlineMeetingUrl,omitempty"`

	// The organizer of the event.
	Organizer Recipient `json:"organizer"`

	// The end time zone that was set when the event was created. A value of tzone://Microsoft/Custom indicates that a
	// legacy custom time zone was set in desktop Outlook.
	OriginalEndTimeZone nullable.Type[string] `json:"originalEndTimeZone,omitempty"`

	// Represents the start time of an event when it's initially created as an occurrence or exception in a recurring
	// series. This property isn't returned for events that are single instances. Its date and time information is expressed
	// in ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	OriginalStart nullable.Type[string] `json:"originalStart,omitempty"`

	// The start time zone that was set when the event was created. A value of tzone://Microsoft/Custom indicates that a
	// legacy custom time zone was set in desktop Outlook.
	OriginalStartTimeZone nullable.Type[string] `json:"originalStartTimeZone,omitempty"`

	// The recurrence pattern for the event.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// The number of minutes before the event start time that the reminder alert occurs.
	ReminderMinutesBeforeStart nullable.Type[int64] `json:"reminderMinutesBeforeStart,omitempty"`

	// Default is true, which represents the organizer would like an invitee to send a response to the event.
	ResponseRequested nullable.Type[bool] `json:"responseRequested,omitempty"`

	// Indicates the type of response sent in response to an event message.
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`

	// Possible values are: normal, personal, private, confidential.
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`

	// The ID for the recurring series master item, if this event is part of a recurring series.
	SeriesMasterId nullable.Type[string] `json:"seriesMasterId,omitempty"`

	// The status to show. Possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
	ShowAs *FreeBusyStatus `json:"showAs,omitempty"`

	// The collection of single-value extended properties defined for the event. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The start date, time, and time zone of the event. By default, the start time is in UTC.
	Start *DateTimeTimeZone `json:"start,omitempty"`

	// The text of the event's subject line.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// A custom identifier specified by a client app for the server to avoid redundant POST operations if the client retries
	// to create the same event. This is useful when low network connectivity causes the client to time out before receiving
	// a response from the server for the client's prior create-event request. After you set transactionId when creating an
	// event, you can't change transactionId in a subsequent update. This property is only returned in a response payload if
	// an app has set it. Optional.
	TransactionId nullable.Type[string] `json:"transactionId,omitempty"`

	// The event type. Possible values are: singleInstance, occurrence, exception, seriesMaster. Read-only
	Type *EventType `json:"type,omitempty"`

	// A unique identifier for calendar events. For recurring events, the value is the same for the series master and all of
	// its occurrences including exceptions.
	Uid nullable.Type[string] `json:"uid,omitempty"`

	// The URL to open the event in Outlook on the web.Outlook on the web opens the event in the browser if you're signed in
	// to your mailbox. Otherwise, Outlook on the web prompts you to sign in.This URL can't be accessed from within an
	// iFrame.
	WebLink nullable.Type[string] `json:"webLink,omitempty"`

	// Fields inherited from OutlookItem

	// The categories associated with the item.
	Categories *[]string `json:"categories,omitempty"`

	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange
	// to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Event) OutlookItem() BaseOutlookItemImpl {
	return BaseOutlookItemImpl{
		Categories:           s.Categories,
		ChangeKey:            s.ChangeKey,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Event) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Event{}

func (s Event) MarshalJSON() ([]byte, error) {
	type wrapper Event
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Event: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Event: %+v", err)
	}

	delete(decoded, "attachments")
	delete(decoded, "calendar")
	delete(decoded, "iCalUId")
	delete(decoded, "instances")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "onlineMeeting")
	delete(decoded, "onlineMeetingUrl")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.event"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Event: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Event{}

func (s *Event) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowNewTimeProposals         nullable.Type[bool]                  `json:"allowNewTimeProposals,omitempty"`
		Attendees                     *[]Attendee                          `json:"attendees,omitempty"`
		Body                          *ItemBody                            `json:"body,omitempty"`
		BodyPreview                   nullable.Type[string]                `json:"bodyPreview,omitempty"`
		Calendar                      *Calendar                            `json:"calendar,omitempty"`
		CancelledOccurrences          *[]string                            `json:"cancelledOccurrences,omitempty"`
		End                           *DateTimeTimeZone                    `json:"end,omitempty"`
		ExceptionOccurrences          *[]Event                             `json:"exceptionOccurrences,omitempty"`
		HasAttachments                nullable.Type[bool]                  `json:"hasAttachments,omitempty"`
		HideAttendees                 nullable.Type[bool]                  `json:"hideAttendees,omitempty"`
		ICalUId                       nullable.Type[string]                `json:"iCalUId,omitempty"`
		Importance                    *Importance                          `json:"importance,omitempty"`
		Instances                     *[]Event                             `json:"instances,omitempty"`
		IsAllDay                      nullable.Type[bool]                  `json:"isAllDay,omitempty"`
		IsCancelled                   nullable.Type[bool]                  `json:"isCancelled,omitempty"`
		IsDraft                       nullable.Type[bool]                  `json:"isDraft,omitempty"`
		IsOnlineMeeting               nullable.Type[bool]                  `json:"isOnlineMeeting,omitempty"`
		IsOrganizer                   nullable.Type[bool]                  `json:"isOrganizer,omitempty"`
		IsReminderOn                  nullable.Type[bool]                  `json:"isReminderOn,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		OccurrenceId                  nullable.Type[string]                `json:"occurrenceId,omitempty"`
		OnlineMeeting                 *OnlineMeetingInfo                   `json:"onlineMeeting,omitempty"`
		OnlineMeetingProvider         *OnlineMeetingProviderType           `json:"onlineMeetingProvider,omitempty"`
		OnlineMeetingUrl              nullable.Type[string]                `json:"onlineMeetingUrl,omitempty"`
		OriginalEndTimeZone           nullable.Type[string]                `json:"originalEndTimeZone,omitempty"`
		OriginalStart                 nullable.Type[string]                `json:"originalStart,omitempty"`
		OriginalStartTimeZone         nullable.Type[string]                `json:"originalStartTimeZone,omitempty"`
		Recurrence                    *PatternedRecurrence                 `json:"recurrence,omitempty"`
		ReminderMinutesBeforeStart    nullable.Type[int64]                 `json:"reminderMinutesBeforeStart,omitempty"`
		ResponseRequested             nullable.Type[bool]                  `json:"responseRequested,omitempty"`
		ResponseStatus                *ResponseStatus                      `json:"responseStatus,omitempty"`
		Sensitivity                   *Sensitivity                         `json:"sensitivity,omitempty"`
		SeriesMasterId                nullable.Type[string]                `json:"seriesMasterId,omitempty"`
		ShowAs                        *FreeBusyStatus                      `json:"showAs,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		Start                         *DateTimeTimeZone                    `json:"start,omitempty"`
		Subject                       nullable.Type[string]                `json:"subject,omitempty"`
		TransactionId                 nullable.Type[string]                `json:"transactionId,omitempty"`
		Type                          *EventType                           `json:"type,omitempty"`
		Uid                           nullable.Type[string]                `json:"uid,omitempty"`
		WebLink                       nullable.Type[string]                `json:"webLink,omitempty"`
		Categories                    *[]string                            `json:"categories,omitempty"`
		ChangeKey                     nullable.Type[string]                `json:"changeKey,omitempty"`
		CreatedDateTime               nullable.Type[string]                `json:"createdDateTime,omitempty"`
		LastModifiedDateTime          nullable.Type[string]                `json:"lastModifiedDateTime,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowNewTimeProposals = decoded.AllowNewTimeProposals
	s.Attendees = decoded.Attendees
	s.Body = decoded.Body
	s.BodyPreview = decoded.BodyPreview
	s.Calendar = decoded.Calendar
	s.CancelledOccurrences = decoded.CancelledOccurrences
	s.End = decoded.End
	s.ExceptionOccurrences = decoded.ExceptionOccurrences
	s.HasAttachments = decoded.HasAttachments
	s.HideAttendees = decoded.HideAttendees
	s.ICalUId = decoded.ICalUId
	s.Importance = decoded.Importance
	s.Instances = decoded.Instances
	s.IsAllDay = decoded.IsAllDay
	s.IsCancelled = decoded.IsCancelled
	s.IsDraft = decoded.IsDraft
	s.IsOnlineMeeting = decoded.IsOnlineMeeting
	s.IsOrganizer = decoded.IsOrganizer
	s.IsReminderOn = decoded.IsReminderOn
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.OccurrenceId = decoded.OccurrenceId
	s.OnlineMeeting = decoded.OnlineMeeting
	s.OnlineMeetingProvider = decoded.OnlineMeetingProvider
	s.OnlineMeetingUrl = decoded.OnlineMeetingUrl
	s.OriginalEndTimeZone = decoded.OriginalEndTimeZone
	s.OriginalStart = decoded.OriginalStart
	s.OriginalStartTimeZone = decoded.OriginalStartTimeZone
	s.Recurrence = decoded.Recurrence
	s.ReminderMinutesBeforeStart = decoded.ReminderMinutesBeforeStart
	s.ResponseRequested = decoded.ResponseRequested
	s.ResponseStatus = decoded.ResponseStatus
	s.Sensitivity = decoded.Sensitivity
	s.SeriesMasterId = decoded.SeriesMasterId
	s.ShowAs = decoded.ShowAs
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.Start = decoded.Start
	s.Subject = decoded.Subject
	s.TransactionId = decoded.TransactionId
	s.Type = decoded.Type
	s.Uid = decoded.Uid
	s.WebLink = decoded.WebLink
	s.Categories = decoded.Categories
	s.ChangeKey = decoded.ChangeKey
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Event into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attachments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attachments into list []json.RawMessage: %+v", err)
		}

		output := make([]Attachment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAttachmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attachments' for 'Event': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attachments = &output
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Event': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["location"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'Event': %+v", err)
		}
		s.Location = impl
	}

	if v, ok := temp["locations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Locations into list []json.RawMessage: %+v", err)
		}

		output := make([]Location, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLocationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Locations' for 'Event': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Locations = &output
	}

	if v, ok := temp["organizer"]; ok {
		impl, err := UnmarshalRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Organizer' for 'Event': %+v", err)
		}
		s.Organizer = impl
	}

	return nil
}
