package stable

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

	// The permissions of the users with whom the calendar is shared.
	CalendarPermissions *[]CalendarPermission `json:"calendarPermissions,omitempty"`

	// The calendar view for the calendar. Navigation property. Read-only.
	CalendarView *[]Event `json:"calendarView,omitempty"`

	// true if the user can write to the calendar, false otherwise. This property is true for the user who created the
	// calendar. This property is also true for a user who shared a calendar and granted write access.
	CanEdit nullable.Type[bool] `json:"canEdit,omitempty"`

	// true if the user has permission to share the calendar, false otherwise. Only the user who created the calendar can
	// share it.
	CanShare nullable.Type[bool] `json:"canShare,omitempty"`

	// If true, the user can read calendar items that have been marked private, false otherwise.
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
	// set a color for the calendar, this property is empty. Read-only.
	HexColor nullable.Type[string] `json:"hexColor,omitempty"`

	// true if this is the default calendar where new events are created by default, false otherwise.
	IsDefaultCalendar nullable.Type[bool] `json:"isDefaultCalendar,omitempty"`

	// Indicates whether this user calendar can be deleted from the user mailbox.
	IsRemovable nullable.Type[bool] `json:"isRemovable,omitempty"`

	// Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users'
	// primary calendars support tracking of meeting responses.
	IsTallyingResponses nullable.Type[bool] `json:"isTallyingResponses,omitempty"`

	// The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The calendar name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// If set, this represents the user who created or added the calendar. For a calendar that the user created or added,
	// the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person
	// who shared that calendar with the user.
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
	delete(decoded, "changeKey")
	delete(decoded, "events")
	delete(decoded, "hexColor")
	delete(decoded, "multiValueExtendedProperties")
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
