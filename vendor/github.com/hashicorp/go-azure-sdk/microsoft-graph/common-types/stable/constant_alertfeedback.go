package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertFeedback string

const (
	AlertFeedback_BenignPositive AlertFeedback = "benignPositive"
	AlertFeedback_FalsePositive  AlertFeedback = "falsePositive"
	AlertFeedback_TruePositive   AlertFeedback = "truePositive"
	AlertFeedback_Unknown        AlertFeedback = "unknown"
)

func PossibleValuesForAlertFeedback() []string {
	return []string{
		string(AlertFeedback_BenignPositive),
		string(AlertFeedback_FalsePositive),
		string(AlertFeedback_TruePositive),
		string(AlertFeedback_Unknown),
	}
}

func (s *AlertFeedback) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertFeedback(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertFeedback(input string) (*AlertFeedback, error) {
	vals := map[string]AlertFeedback{
		"benignpositive": AlertFeedback_BenignPositive,
		"falsepositive":  AlertFeedback_FalsePositive,
		"truepositive":   AlertFeedback_TruePositive,
		"unknown":        AlertFeedback_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertFeedback(input)
	return &out, nil
}
