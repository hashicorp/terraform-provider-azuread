package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleOperationType string

const (
	Win32LobAppPowerShellScriptRuleOperationType_Boolean       Win32LobAppPowerShellScriptRuleOperationType = "boolean"
	Win32LobAppPowerShellScriptRuleOperationType_DateTime      Win32LobAppPowerShellScriptRuleOperationType = "dateTime"
	Win32LobAppPowerShellScriptRuleOperationType_Float         Win32LobAppPowerShellScriptRuleOperationType = "float"
	Win32LobAppPowerShellScriptRuleOperationType_Integer       Win32LobAppPowerShellScriptRuleOperationType = "integer"
	Win32LobAppPowerShellScriptRuleOperationType_NotConfigured Win32LobAppPowerShellScriptRuleOperationType = "notConfigured"
	Win32LobAppPowerShellScriptRuleOperationType_String        Win32LobAppPowerShellScriptRuleOperationType = "string"
	Win32LobAppPowerShellScriptRuleOperationType_Version       Win32LobAppPowerShellScriptRuleOperationType = "version"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleOperationType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleOperationType_Boolean),
		string(Win32LobAppPowerShellScriptRuleOperationType_DateTime),
		string(Win32LobAppPowerShellScriptRuleOperationType_Float),
		string(Win32LobAppPowerShellScriptRuleOperationType_Integer),
		string(Win32LobAppPowerShellScriptRuleOperationType_NotConfigured),
		string(Win32LobAppPowerShellScriptRuleOperationType_String),
		string(Win32LobAppPowerShellScriptRuleOperationType_Version),
	}
}

func (s *Win32LobAppPowerShellScriptRuleOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRuleOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRuleOperationType(input string) (*Win32LobAppPowerShellScriptRuleOperationType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleOperationType{
		"boolean":       Win32LobAppPowerShellScriptRuleOperationType_Boolean,
		"datetime":      Win32LobAppPowerShellScriptRuleOperationType_DateTime,
		"float":         Win32LobAppPowerShellScriptRuleOperationType_Float,
		"integer":       Win32LobAppPowerShellScriptRuleOperationType_Integer,
		"notconfigured": Win32LobAppPowerShellScriptRuleOperationType_NotConfigured,
		"string":        Win32LobAppPowerShellScriptRuleOperationType_String,
		"version":       Win32LobAppPowerShellScriptRuleOperationType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleOperationType(input)
	return &out, nil
}
