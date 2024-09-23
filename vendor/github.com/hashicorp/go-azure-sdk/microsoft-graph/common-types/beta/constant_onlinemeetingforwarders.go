package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingForwarders string

const (
	OnlineMeetingForwarders_Everyone  OnlineMeetingForwarders = "everyone"
	OnlineMeetingForwarders_Organizer OnlineMeetingForwarders = "organizer"
)

func PossibleValuesForOnlineMeetingForwarders() []string {
	return []string{
		string(OnlineMeetingForwarders_Everyone),
		string(OnlineMeetingForwarders_Organizer),
	}
}

func (s *OnlineMeetingForwarders) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingForwarders(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingForwarders(input string) (*OnlineMeetingForwarders, error) {
	vals := map[string]OnlineMeetingForwarders{
		"everyone":  OnlineMeetingForwarders_Everyone,
		"organizer": OnlineMeetingForwarders_Organizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingForwarders(input)
	return &out, nil
}
