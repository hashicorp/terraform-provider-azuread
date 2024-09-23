package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffRole string

const (
	BookingStaffRole_Administrator BookingStaffRole = "administrator"
	BookingStaffRole_ExternalGuest BookingStaffRole = "externalGuest"
	BookingStaffRole_Guest         BookingStaffRole = "guest"
	BookingStaffRole_Scheduler     BookingStaffRole = "scheduler"
	BookingStaffRole_TeamMember    BookingStaffRole = "teamMember"
	BookingStaffRole_Viewer        BookingStaffRole = "viewer"
)

func PossibleValuesForBookingStaffRole() []string {
	return []string{
		string(BookingStaffRole_Administrator),
		string(BookingStaffRole_ExternalGuest),
		string(BookingStaffRole_Guest),
		string(BookingStaffRole_Scheduler),
		string(BookingStaffRole_TeamMember),
		string(BookingStaffRole_Viewer),
	}
}

func (s *BookingStaffRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBookingStaffRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBookingStaffRole(input string) (*BookingStaffRole, error) {
	vals := map[string]BookingStaffRole{
		"administrator": BookingStaffRole_Administrator,
		"externalguest": BookingStaffRole_ExternalGuest,
		"guest":         BookingStaffRole_Guest,
		"scheduler":     BookingStaffRole_Scheduler,
		"teammember":    BookingStaffRole_TeamMember,
		"viewer":        BookingStaffRole_Viewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BookingStaffRole(input)
	return &out, nil
}
