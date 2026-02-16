package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertClassification string

const (
	SecurityAlertClassification_FalsePositive                 SecurityAlertClassification = "falsePositive"
	SecurityAlertClassification_InformationalExpectedActivity SecurityAlertClassification = "informationalExpectedActivity"
	SecurityAlertClassification_TruePositive                  SecurityAlertClassification = "truePositive"
	SecurityAlertClassification_Unknown                       SecurityAlertClassification = "unknown"
)

func PossibleValuesForSecurityAlertClassification() []string {
	return []string{
		string(SecurityAlertClassification_FalsePositive),
		string(SecurityAlertClassification_InformationalExpectedActivity),
		string(SecurityAlertClassification_TruePositive),
		string(SecurityAlertClassification_Unknown),
	}
}

func (s *SecurityAlertClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertClassification(input string) (*SecurityAlertClassification, error) {
	vals := map[string]SecurityAlertClassification{
		"falsepositive":                 SecurityAlertClassification_FalsePositive,
		"informationalexpectedactivity": SecurityAlertClassification_InformationalExpectedActivity,
		"truepositive":                  SecurityAlertClassification_TruePositive,
		"unknown":                       SecurityAlertClassification_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertClassification(input)
	return &out, nil
}
