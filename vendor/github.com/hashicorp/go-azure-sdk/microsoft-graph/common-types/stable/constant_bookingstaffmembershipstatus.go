package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMembershipStatus string

const (
	BookingStaffMembershipStatus_Active            BookingStaffMembershipStatus = "active"
	BookingStaffMembershipStatus_PendingAcceptance BookingStaffMembershipStatus = "pendingAcceptance"
	BookingStaffMembershipStatus_RejectedByStaff   BookingStaffMembershipStatus = "rejectedByStaff"
)

func PossibleValuesForBookingStaffMembershipStatus() []string {
	return []string{
		string(BookingStaffMembershipStatus_Active),
		string(BookingStaffMembershipStatus_PendingAcceptance),
		string(BookingStaffMembershipStatus_RejectedByStaff),
	}
}

func (s *BookingStaffMembershipStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingStaffMembershipStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingStaffMembershipStatus(input string) (*BookingStaffMembershipStatus, error) {
	vals := map[string]BookingStaffMembershipStatus{
		"active":            BookingStaffMembershipStatus_Active,
		"pendingacceptance": BookingStaffMembershipStatus_PendingAcceptance,
		"rejectedbystaff":   BookingStaffMembershipStatus_RejectedByStaff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffMembershipStatus(input)
	return &out, nil
}
