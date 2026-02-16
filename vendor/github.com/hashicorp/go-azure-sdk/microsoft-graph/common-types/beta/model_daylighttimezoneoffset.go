package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StandardTimeZoneOffset = DaylightTimeZoneOffset{}

type DaylightTimeZoneOffset struct {
	// The time offset from Coordinated Universal Time (UTC) for daylight saving time. This value is in minutes.
	DaylightBias nullable.Type[int64] `json:"daylightBias,omitempty"`

	// Fields inherited from StandardTimeZoneOffset

	// Represents the nth occurrence of the day of week that the transition from daylight saving time to standard time
	// occurs.
	DayOccurrence nullable.Type[int64] `json:"dayOccurrence,omitempty"`

	// Represents the day of the week when the transition from daylight saving time to standard time.
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`

	// Represents the month of the year when the transition from daylight saving time to standard time occurs.
	Month nullable.Type[int64] `json:"month,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the time of day when the transition from daylight saving time to standard time occurs.
	Time nullable.Type[string] `json:"time,omitempty"`

	// Represents how frequently in terms of years the change from daylight saving time to standard time occurs. For
	// example, a value of 0 means every year.
	Year nullable.Type[int64] `json:"year,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DaylightTimeZoneOffset) StandardTimeZoneOffset() BaseStandardTimeZoneOffsetImpl {
	return BaseStandardTimeZoneOffsetImpl{
		DayOccurrence: s.DayOccurrence,
		DayOfWeek:     s.DayOfWeek,
		Month:         s.Month,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		Time:          s.Time,
		Year:          s.Year,
	}
}

var _ json.Marshaler = DaylightTimeZoneOffset{}

func (s DaylightTimeZoneOffset) MarshalJSON() ([]byte, error) {
	type wrapper DaylightTimeZoneOffset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DaylightTimeZoneOffset: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DaylightTimeZoneOffset: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.daylightTimeZoneOffset"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DaylightTimeZoneOffset: %+v", err)
	}

	return encoded, nil
}
