package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeAvailability struct {
	// The email address and type of attendee - whether it's a person or a resource, and whether required or optional if
	// it's a person.
	Attendee *AttendeeBase `json:"attendee,omitempty"`

	// The availability status of the attendee. Possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
	Availability *FreeBusyStatus `json:"availability,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &AttendeeAvailability{}

func (s *AttendeeAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Availability *FreeBusyStatus `json:"availability,omitempty"`
		ODataId      *string         `json:"@odata.id,omitempty"`
		ODataType    *string         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Availability = decoded.Availability
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AttendeeAvailability into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attendee"]; ok {
		impl, err := UnmarshalAttendeeBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Attendee' for 'AttendeeAvailability': %+v", err)
		}
		s.Attendee = &impl
	}

	return nil
}
