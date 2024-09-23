package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingAudience string

const (
	MeetingAudience_Everyone     MeetingAudience = "everyone"
	MeetingAudience_Organization MeetingAudience = "organization"
)

func PossibleValuesForMeetingAudience() []string {
	return []string{
		string(MeetingAudience_Everyone),
		string(MeetingAudience_Organization),
	}
}

func (s *MeetingAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingAudience(input string) (*MeetingAudience, error) {
	vals := map[string]MeetingAudience{
		"everyone":     MeetingAudience_Everyone,
		"organization": MeetingAudience_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingAudience(input)
	return &out, nil
}
