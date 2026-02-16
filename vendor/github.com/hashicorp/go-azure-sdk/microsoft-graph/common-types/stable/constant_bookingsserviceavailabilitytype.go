package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingsServiceAvailabilityType string

const (
	BookingsServiceAvailabilityType_BookWhenStaffAreFree BookingsServiceAvailabilityType = "bookWhenStaffAreFree"
	BookingsServiceAvailabilityType_CustomWeeklyHours    BookingsServiceAvailabilityType = "customWeeklyHours"
	BookingsServiceAvailabilityType_NotBookable          BookingsServiceAvailabilityType = "notBookable"
)

func PossibleValuesForBookingsServiceAvailabilityType() []string {
	return []string{
		string(BookingsServiceAvailabilityType_BookWhenStaffAreFree),
		string(BookingsServiceAvailabilityType_CustomWeeklyHours),
		string(BookingsServiceAvailabilityType_NotBookable),
	}
}

func (s *BookingsServiceAvailabilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingsServiceAvailabilityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingsServiceAvailabilityType(input string) (*BookingsServiceAvailabilityType, error) {
	vals := map[string]BookingsServiceAvailabilityType{
		"bookwhenstaffarefree": BookingsServiceAvailabilityType_BookWhenStaffAreFree,
		"customweeklyhours":    BookingsServiceAvailabilityType_CustomWeeklyHours,
		"notbookable":          BookingsServiceAvailabilityType_NotBookable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingsServiceAvailabilityType(input)
	return &out, nil
}
