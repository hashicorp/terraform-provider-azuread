package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticsActivityType string

const (
	AnalyticsActivityType_Call    AnalyticsActivityType = "Call"
	AnalyticsActivityType_Chat    AnalyticsActivityType = "Chat"
	AnalyticsActivityType_Email   AnalyticsActivityType = "Email"
	AnalyticsActivityType_Focus   AnalyticsActivityType = "Focus"
	AnalyticsActivityType_Meeting AnalyticsActivityType = "Meeting"
)

func PossibleValuesForAnalyticsActivityType() []string {
	return []string{
		string(AnalyticsActivityType_Call),
		string(AnalyticsActivityType_Chat),
		string(AnalyticsActivityType_Email),
		string(AnalyticsActivityType_Focus),
		string(AnalyticsActivityType_Meeting),
	}
}

func (s *AnalyticsActivityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAnalyticsActivityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAnalyticsActivityType(input string) (*AnalyticsActivityType, error) {
	vals := map[string]AnalyticsActivityType{
		"call":    AnalyticsActivityType_Call,
		"chat":    AnalyticsActivityType_Chat,
		"email":   AnalyticsActivityType_Email,
		"focus":   AnalyticsActivityType_Focus,
		"meeting": AnalyticsActivityType_Meeting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AnalyticsActivityType(input)
	return &out, nil
}
