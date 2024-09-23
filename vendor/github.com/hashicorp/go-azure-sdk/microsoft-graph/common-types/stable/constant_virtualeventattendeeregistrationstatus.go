package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventAttendeeRegistrationStatus string

const (
	VirtualEventAttendeeRegistrationStatus_Canceled            VirtualEventAttendeeRegistrationStatus = "canceled"
	VirtualEventAttendeeRegistrationStatus_PendingApproval     VirtualEventAttendeeRegistrationStatus = "pendingApproval"
	VirtualEventAttendeeRegistrationStatus_Registered          VirtualEventAttendeeRegistrationStatus = "registered"
	VirtualEventAttendeeRegistrationStatus_RejectedByOrganizer VirtualEventAttendeeRegistrationStatus = "rejectedByOrganizer"
	VirtualEventAttendeeRegistrationStatus_Waitlisted          VirtualEventAttendeeRegistrationStatus = "waitlisted"
)

func PossibleValuesForVirtualEventAttendeeRegistrationStatus() []string {
	return []string{
		string(VirtualEventAttendeeRegistrationStatus_Canceled),
		string(VirtualEventAttendeeRegistrationStatus_PendingApproval),
		string(VirtualEventAttendeeRegistrationStatus_Registered),
		string(VirtualEventAttendeeRegistrationStatus_RejectedByOrganizer),
		string(VirtualEventAttendeeRegistrationStatus_Waitlisted),
	}
}

func (s *VirtualEventAttendeeRegistrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventAttendeeRegistrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventAttendeeRegistrationStatus(input string) (*VirtualEventAttendeeRegistrationStatus, error) {
	vals := map[string]VirtualEventAttendeeRegistrationStatus{
		"canceled":            VirtualEventAttendeeRegistrationStatus_Canceled,
		"pendingapproval":     VirtualEventAttendeeRegistrationStatus_PendingApproval,
		"registered":          VirtualEventAttendeeRegistrationStatus_Registered,
		"rejectedbyorganizer": VirtualEventAttendeeRegistrationStatus_RejectedByOrganizer,
		"waitlisted":          VirtualEventAttendeeRegistrationStatus_Waitlisted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventAttendeeRegistrationStatus(input)
	return &out, nil
}
