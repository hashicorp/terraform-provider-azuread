package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Calendar{}

type Calendar struct {
	// Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible
	// values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
	AllowedOnlineMeetingProviders *[]OnlineMeetingProviderType `json:"allowedOnlineMeetingProviders,omitempty"`

	// The calendarGroup in which to create the calendar. If the user has never explicitly set a group for the calendar,
	// this property is null.
	CalendarGroupId nullable.Type[string] `json:"calendarGroupId,omitempty"`

	// The permissions of the users with whom the calendar is shared.
	CalendarPermissions *[]CalendarPermission `json:"calendarPermissions,omitempty"`

	// The calendar view for the calendar. Navigation property. Read-only.
	CalendarView *[]Event `json:"calendarView,omitempty"`

	// true if the user can write to the calendar, false otherwise. This property is true for the user who created the
	// calendar. This property is also true for a user who has been shared a calendar and granted write access, through an
	// Outlook client or the corresponding calendarPermission resource. Read-only.
	CanEdit nullable.Type[bool] `json:"canEdit,omitempty"`

	// true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar
	// can share it. Read-only.
	CanShare nullable.Type[bool] `json:"canShare,omitempty"`

	// true if the user can read calendar items that have been marked private, false otherwise. This property is set through
	// an Outlook client or the corresponding calendarPermission resource. Read-only.
	CanViewPrivateItems nullable.Type[bool] `json:"canViewPrivateItems,omitempty"`

	// Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This
	// allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto,
	// lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
	Color *CalendarColor `json:"color,omitempty"`

	// The default online meeting provider for meetings sent from this calendar. Possible values are: unknown,
	// skypeForBusiness, skypeForConsumer, teamsForBusiness.
	DefaultOnlineMeetingProvider *OnlineMeetingProviderType `json:"defaultOnlineMeetingProvider,omitempty"`

	// The events in the calendar. Navigation property. Read-only.
	Events *[]Event `json:"events,omitempty"`

	// The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and
	// representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly
	// set a color for the calendar, this property is empty.
	HexColor nullable.Type[string] `json:"hexColor,omitempty"`

	// true if this is the default calendar where new events are created by default, false otherwise.
	IsDefaultCalendar nullable.Type[bool] `json:"isDefaultCalendar,omitempty"`

	// Indicates whether this user calendar can be deleted from the user mailbox.
	IsRemovable nullable.Type[bool] `json:"isRemovable,omitempty"`

	// true if the user has shared the calendar with other users, false otherwise. Since only the user who created the
	// calendar can share it, isShared and isSharedWithMe cannot be true for the same user. This property is set when
	// sharing is initiated in an Outlook client, and can be reset when the sharing is cancelled through the client or the
	// corresponding calendarPermission resource. Read-only.
	IsShared nullable.Type[bool] `json:"isShared,omitempty"`

	// true if the user has been shared this calendar, false otherwise. This property is always false for a calendar owner.
	// This property is set when sharing is initiated in an Outlook client, and can be reset when the sharing is cancelled
	// through the client or the corresponding calendarPermission resource. Read-only.
	IsSharedWithMe nullable.Type[bool] `json:"isSharedWithMe,omitempty"`

	// Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users'
	// primary calendars support tracking of meeting responses.
	IsTallyingResponses nullable.Type[bool] `json:"isTallyingResponses,omitempty"`

	// The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The calendar name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// If set, this represents the user who created or added the calendar. For a calendar that the user created or added,
	// the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person
	// who shared that calendar with the user. Read-only.
	Owner *EmailAddress `json:"owner,omitempty"`

	// The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

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

func (s Calendar) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Calendar{}

func (s Calendar) MarshalJSON() ([]byte, error) {
	type wrapper Calendar
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Calendar: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Calendar: %+v", err)
	}

	delete(decoded, "calendarView")
	delete(decoded, "canEdit")
	delete(decoded, "canShare")
	delete(decoded, "canViewPrivateItems")
	delete(decoded, "changeKey")
	delete(decoded, "events")
	delete(decoded, "isShared")
	delete(decoded, "isSharedWithMe")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "owner")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.calendar"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Calendar: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Calendar{}

func (s *Calendar) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowedOnlineMeetingProviders *[]OnlineMeetingProviderType         `json:"allowedOnlineMeetingProviders,omitempty"`
		CalendarGroupId               nullable.Type[string]                `json:"calendarGroupId,omitempty"`
		CalendarPermissions           *[]CalendarPermission                `json:"calendarPermissions,omitempty"`
		CalendarView                  *[]Event                             `json:"calendarView,omitempty"`
		CanEdit                       nullable.Type[bool]                  `json:"canEdit,omitempty"`
		CanShare                      nullable.Type[bool]                  `json:"canShare,omitempty"`
		CanViewPrivateItems           nullable.Type[bool]                  `json:"canViewPrivateItems,omitempty"`
		ChangeKey                     nullable.Type[string]                `json:"changeKey,omitempty"`
		Color                         *CalendarColor                       `json:"color,omitempty"`
		DefaultOnlineMeetingProvider  *OnlineMeetingProviderType           `json:"defaultOnlineMeetingProvider,omitempty"`
		Events                        *[]Event                             `json:"events,omitempty"`
		HexColor                      nullable.Type[string]                `json:"hexColor,omitempty"`
		IsDefaultCalendar             nullable.Type[bool]                  `json:"isDefaultCalendar,omitempty"`
		IsRemovable                   nullable.Type[bool]                  `json:"isRemovable,omitempty"`
		IsShared                      nullable.Type[bool]                  `json:"isShared,omitempty"`
		IsSharedWithMe                nullable.Type[bool]                  `json:"isSharedWithMe,omitempty"`
		IsTallyingResponses           nullable.Type[bool]                  `json:"isTallyingResponses,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		Name                          nullable.Type[string]                `json:"name,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowedOnlineMeetingProviders = decoded.AllowedOnlineMeetingProviders
	s.CalendarGroupId = decoded.CalendarGroupId
	s.CalendarPermissions = decoded.CalendarPermissions
	s.CalendarView = decoded.CalendarView
	s.CanEdit = decoded.CanEdit
	s.CanShare = decoded.CanShare
	s.CanViewPrivateItems = decoded.CanViewPrivateItems
	s.ChangeKey = decoded.ChangeKey
	s.Color = decoded.Color
	s.DefaultOnlineMeetingProvider = decoded.DefaultOnlineMeetingProvider
	s.Events = decoded.Events
	s.HexColor = decoded.HexColor
	s.IsDefaultCalendar = decoded.IsDefaultCalendar
	s.IsRemovable = decoded.IsRemovable
	s.IsShared = decoded.IsShared
	s.IsSharedWithMe = decoded.IsSharedWithMe
	s.IsTallyingResponses = decoded.IsTallyingResponses
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.Name = decoded.Name
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Calendar into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["owner"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Owner' for 'Calendar': %+v", err)
		}
		s.Owner = &impl
	}

	return nil
}
