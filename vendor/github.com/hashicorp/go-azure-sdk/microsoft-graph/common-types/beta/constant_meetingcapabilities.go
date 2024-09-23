package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingCapabilities string

const (
	MeetingCapabilities_QuestionAndAnswer MeetingCapabilities = "questionAndAnswer"
)

func PossibleValuesForMeetingCapabilities() []string {
	return []string{
		string(MeetingCapabilities_QuestionAndAnswer),
	}
}

func (s *MeetingCapabilities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingCapabilities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingCapabilities(input string) (*MeetingCapabilities, error) {
	vals := map[string]MeetingCapabilities{
		"questionandanswer": MeetingCapabilities_QuestionAndAnswer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingCapabilities(input)
	return &out, nil
}
