package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRuleOperator string

const (
	Win32LobAppRuleOperator_Equal              Win32LobAppRuleOperator = "equal"
	Win32LobAppRuleOperator_GreaterThan        Win32LobAppRuleOperator = "greaterThan"
	Win32LobAppRuleOperator_GreaterThanOrEqual Win32LobAppRuleOperator = "greaterThanOrEqual"
	Win32LobAppRuleOperator_LessThan           Win32LobAppRuleOperator = "lessThan"
	Win32LobAppRuleOperator_LessThanOrEqual    Win32LobAppRuleOperator = "lessThanOrEqual"
	Win32LobAppRuleOperator_NotConfigured      Win32LobAppRuleOperator = "notConfigured"
	Win32LobAppRuleOperator_NotEqual           Win32LobAppRuleOperator = "notEqual"
)

func PossibleValuesForWin32LobAppRuleOperator() []string {
	return []string{
		string(Win32LobAppRuleOperator_Equal),
		string(Win32LobAppRuleOperator_GreaterThan),
		string(Win32LobAppRuleOperator_GreaterThanOrEqual),
		string(Win32LobAppRuleOperator_LessThan),
		string(Win32LobAppRuleOperator_LessThanOrEqual),
		string(Win32LobAppRuleOperator_NotConfigured),
		string(Win32LobAppRuleOperator_NotEqual),
	}
}

func (s *Win32LobAppRuleOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRuleOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRuleOperator(input string) (*Win32LobAppRuleOperator, error) {
	vals := map[string]Win32LobAppRuleOperator{
		"equal":              Win32LobAppRuleOperator_Equal,
		"greaterthan":        Win32LobAppRuleOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppRuleOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppRuleOperator_LessThan,
		"lessthanorequal":    Win32LobAppRuleOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppRuleOperator_NotConfigured,
		"notequal":           Win32LobAppRuleOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRuleOperator(input)
	return &out, nil
}
