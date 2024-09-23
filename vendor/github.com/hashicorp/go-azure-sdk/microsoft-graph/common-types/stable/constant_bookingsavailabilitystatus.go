package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingsAvailabilityStatus string

const (
	BookingsAvailabilityStatus_Available      BookingsAvailabilityStatus = "available"
	BookingsAvailabilityStatus_Busy           BookingsAvailabilityStatus = "busy"
	BookingsAvailabilityStatus_OutOfOffice    BookingsAvailabilityStatus = "outOfOffice"
	BookingsAvailabilityStatus_SlotsAvailable BookingsAvailabilityStatus = "slotsAvailable"
)

func PossibleValuesForBookingsAvailabilityStatus() []string {
	return []string{
		string(BookingsAvailabilityStatus_Available),
		string(BookingsAvailabilityStatus_Busy),
		string(BookingsAvailabilityStatus_OutOfOffice),
		string(BookingsAvailabilityStatus_SlotsAvailable),
	}
}

func (s *BookingsAvailabilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingsAvailabilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingsAvailabilityStatus(input string) (*BookingsAvailabilityStatus, error) {
	vals := map[string]BookingsAvailabilityStatus{
		"available":      BookingsAvailabilityStatus_Available,
		"busy":           BookingsAvailabilityStatus_Busy,
		"outofoffice":    BookingsAvailabilityStatus_OutOfOffice,
		"slotsavailable": BookingsAvailabilityStatus_SlotsAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingsAvailabilityStatus(input)
	return &out, nil
}
