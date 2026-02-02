package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingChatHistoryDefaultMode string

const (
	MeetingChatHistoryDefaultMode_All  MeetingChatHistoryDefaultMode = "all"
	MeetingChatHistoryDefaultMode_None MeetingChatHistoryDefaultMode = "none"
)

func PossibleValuesForMeetingChatHistoryDefaultMode() []string {
	return []string{
		string(MeetingChatHistoryDefaultMode_All),
		string(MeetingChatHistoryDefaultMode_None),
	}
}

func (s *MeetingChatHistoryDefaultMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingChatHistoryDefaultMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingChatHistoryDefaultMode(input string) (*MeetingChatHistoryDefaultMode, error) {
	vals := map[string]MeetingChatHistoryDefaultMode{
		"all":  MeetingChatHistoryDefaultMode_All,
		"none": MeetingChatHistoryDefaultMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingChatHistoryDefaultMode(input)
	return &out, nil
}
