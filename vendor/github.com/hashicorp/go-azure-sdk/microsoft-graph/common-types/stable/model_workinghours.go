package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkingHours struct {
	// The days of the week on which the user works.
	DaysOfWeek *[]DayOfWeek `json:"daysOfWeek,omitempty"`

	// The time of the day that the user stops working.
	EndTime nullable.Type[string] `json:"endTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time of the day that the user starts working.
	StartTime nullable.Type[string] `json:"startTime,omitempty"`

	// The time zone to which the working hours apply.
	TimeZone TimeZoneBase `json:"timeZone"`
}

var _ json.Unmarshaler = &WorkingHours{}

func (s *WorkingHours) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DaysOfWeek *[]DayOfWeek          `json:"daysOfWeek,omitempty"`
		EndTime    nullable.Type[string] `json:"endTime,omitempty"`
		ODataId    *string               `json:"@odata.id,omitempty"`
		ODataType  *string               `json:"@odata.type,omitempty"`
		StartTime  nullable.Type[string] `json:"startTime,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DaysOfWeek = decoded.DaysOfWeek
	s.EndTime = decoded.EndTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StartTime = decoded.StartTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WorkingHours into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["timeZone"]; ok {
		impl, err := UnmarshalTimeZoneBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TimeZone' for 'WorkingHours': %+v", err)
		}
		s.TimeZone = impl
	}

	return nil
}
