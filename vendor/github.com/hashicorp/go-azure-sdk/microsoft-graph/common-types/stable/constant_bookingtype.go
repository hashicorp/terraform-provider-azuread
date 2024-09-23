package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingType string

const (
	BookingType_Reserved BookingType = "reserved"
	BookingType_Standard BookingType = "standard"
	BookingType_Unknown  BookingType = "unknown"
)

func PossibleValuesForBookingType() []string {
	return []string{
		string(BookingType_Reserved),
		string(BookingType_Standard),
		string(BookingType_Unknown),
	}
}

func (s *BookingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingType(input string) (*BookingType, error) {
	vals := map[string]BookingType{
		"reserved": BookingType_Reserved,
		"standard": BookingType_Standard,
		"unknown":  BookingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingType(input)
	return &out, nil
}
