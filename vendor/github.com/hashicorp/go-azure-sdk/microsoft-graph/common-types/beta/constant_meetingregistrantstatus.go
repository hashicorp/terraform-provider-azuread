package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantStatus string

const (
	MeetingRegistrantStatus_Canceled   MeetingRegistrantStatus = "canceled"
	MeetingRegistrantStatus_Processing MeetingRegistrantStatus = "processing"
	MeetingRegistrantStatus_Registered MeetingRegistrantStatus = "registered"
)

func PossibleValuesForMeetingRegistrantStatus() []string {
	return []string{
		string(MeetingRegistrantStatus_Canceled),
		string(MeetingRegistrantStatus_Processing),
		string(MeetingRegistrantStatus_Registered),
	}
}

func (s *MeetingRegistrantStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRegistrantStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRegistrantStatus(input string) (*MeetingRegistrantStatus, error) {
	vals := map[string]MeetingRegistrantStatus{
		"canceled":   MeetingRegistrantStatus_Canceled,
		"processing": MeetingRegistrantStatus_Processing,
		"registered": MeetingRegistrantStatus_Registered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrantStatus(input)
	return &out, nil
}
