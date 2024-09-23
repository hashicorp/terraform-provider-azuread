package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardTimeZoneOffset interface {
	StandardTimeZoneOffset() BaseStandardTimeZoneOffsetImpl
}

var _ StandardTimeZoneOffset = BaseStandardTimeZoneOffsetImpl{}

type BaseStandardTimeZoneOffsetImpl struct {
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

func (s BaseStandardTimeZoneOffsetImpl) StandardTimeZoneOffset() BaseStandardTimeZoneOffsetImpl {
	return s
}

var _ StandardTimeZoneOffset = RawStandardTimeZoneOffsetImpl{}

// RawStandardTimeZoneOffsetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawStandardTimeZoneOffsetImpl struct {
	standardTimeZoneOffset BaseStandardTimeZoneOffsetImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawStandardTimeZoneOffsetImpl) StandardTimeZoneOffset() BaseStandardTimeZoneOffsetImpl {
	return s.standardTimeZoneOffset
}

func UnmarshalStandardTimeZoneOffsetImplementation(input []byte) (StandardTimeZoneOffset, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StandardTimeZoneOffset into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.daylightTimeZoneOffset") {
		var out DaylightTimeZoneOffset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DaylightTimeZoneOffset: %+v", err)
		}
		return out, nil
	}

	var parent BaseStandardTimeZoneOffsetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStandardTimeZoneOffsetImpl: %+v", err)
	}

	return RawStandardTimeZoneOffsetImpl{
		standardTimeZoneOffset: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
