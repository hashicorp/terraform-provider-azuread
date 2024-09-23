package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingAudience string

const (
	BroadcastMeetingAudience_Everyone       BroadcastMeetingAudience = "everyone"
	BroadcastMeetingAudience_Organization   BroadcastMeetingAudience = "organization"
	BroadcastMeetingAudience_RoleIsAttendee BroadcastMeetingAudience = "roleIsAttendee"
)

func PossibleValuesForBroadcastMeetingAudience() []string {
	return []string{
		string(BroadcastMeetingAudience_Everyone),
		string(BroadcastMeetingAudience_Organization),
		string(BroadcastMeetingAudience_RoleIsAttendee),
	}
}

func (s *BroadcastMeetingAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBroadcastMeetingAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBroadcastMeetingAudience(input string) (*BroadcastMeetingAudience, error) {
	vals := map[string]BroadcastMeetingAudience{
		"everyone":       BroadcastMeetingAudience_Everyone,
		"organization":   BroadcastMeetingAudience_Organization,
		"roleisattendee": BroadcastMeetingAudience_RoleIsAttendee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BroadcastMeetingAudience(input)
	return &out, nil
}
