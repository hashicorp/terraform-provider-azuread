package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingReminderRecipients string

const (
	BookingReminderRecipients_AllAttendees BookingReminderRecipients = "allAttendees"
	BookingReminderRecipients_Customer     BookingReminderRecipients = "customer"
	BookingReminderRecipients_Staff        BookingReminderRecipients = "staff"
)

func PossibleValuesForBookingReminderRecipients() []string {
	return []string{
		string(BookingReminderRecipients_AllAttendees),
		string(BookingReminderRecipients_Customer),
		string(BookingReminderRecipients_Staff),
	}
}

func (s *BookingReminderRecipients) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingReminderRecipients(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingReminderRecipients(input string) (*BookingReminderRecipients, error) {
	vals := map[string]BookingReminderRecipients{
		"allattendees": BookingReminderRecipients_AllAttendees,
		"customer":     BookingReminderRecipients_Customer,
		"staff":        BookingReminderRecipients_Staff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingReminderRecipients(input)
	return &out, nil
}
