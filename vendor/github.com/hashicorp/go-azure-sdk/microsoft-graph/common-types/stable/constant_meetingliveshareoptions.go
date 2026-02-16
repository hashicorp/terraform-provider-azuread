package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingLiveShareOptions string

const (
	MeetingLiveShareOptions_Disabled MeetingLiveShareOptions = "disabled"
	MeetingLiveShareOptions_Enabled  MeetingLiveShareOptions = "enabled"
)

func PossibleValuesForMeetingLiveShareOptions() []string {
	return []string{
		string(MeetingLiveShareOptions_Disabled),
		string(MeetingLiveShareOptions_Enabled),
	}
}

func (s *MeetingLiveShareOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingLiveShareOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingLiveShareOptions(input string) (*MeetingLiveShareOptions, error) {
	vals := map[string]MeetingLiveShareOptions{
		"disabled": MeetingLiveShareOptions_Disabled,
		"enabled":  MeetingLiveShareOptions_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingLiveShareOptions(input)
	return &out, nil
}
