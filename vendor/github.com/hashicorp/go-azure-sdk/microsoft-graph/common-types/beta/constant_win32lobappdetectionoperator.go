package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppDetectionOperator string

const (
	Win32LobAppDetectionOperator_Equal              Win32LobAppDetectionOperator = "equal"
	Win32LobAppDetectionOperator_GreaterThan        Win32LobAppDetectionOperator = "greaterThan"
	Win32LobAppDetectionOperator_GreaterThanOrEqual Win32LobAppDetectionOperator = "greaterThanOrEqual"
	Win32LobAppDetectionOperator_LessThan           Win32LobAppDetectionOperator = "lessThan"
	Win32LobAppDetectionOperator_LessThanOrEqual    Win32LobAppDetectionOperator = "lessThanOrEqual"
	Win32LobAppDetectionOperator_NotConfigured      Win32LobAppDetectionOperator = "notConfigured"
	Win32LobAppDetectionOperator_NotEqual           Win32LobAppDetectionOperator = "notEqual"
)

func PossibleValuesForWin32LobAppDetectionOperator() []string {
	return []string{
		string(Win32LobAppDetectionOperator_Equal),
		string(Win32LobAppDetectionOperator_GreaterThan),
		string(Win32LobAppDetectionOperator_GreaterThanOrEqual),
		string(Win32LobAppDetectionOperator_LessThan),
		string(Win32LobAppDetectionOperator_LessThanOrEqual),
		string(Win32LobAppDetectionOperator_NotConfigured),
		string(Win32LobAppDetectionOperator_NotEqual),
	}
}

func (s *Win32LobAppDetectionOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppDetectionOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppDetectionOperator(input string) (*Win32LobAppDetectionOperator, error) {
	vals := map[string]Win32LobAppDetectionOperator{
		"equal":              Win32LobAppDetectionOperator_Equal,
		"greaterthan":        Win32LobAppDetectionOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppDetectionOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppDetectionOperator_LessThan,
		"lessthanorequal":    Win32LobAppDetectionOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppDetectionOperator_NotConfigured,
		"notequal":           Win32LobAppDetectionOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppDetectionOperator(input)
	return &out, nil
}
