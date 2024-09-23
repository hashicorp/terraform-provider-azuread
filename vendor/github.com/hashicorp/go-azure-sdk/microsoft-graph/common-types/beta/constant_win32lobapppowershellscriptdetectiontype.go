package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptDetectionType string

const (
	Win32LobAppPowerShellScriptDetectionType_Boolean       Win32LobAppPowerShellScriptDetectionType = "boolean"
	Win32LobAppPowerShellScriptDetectionType_DateTime      Win32LobAppPowerShellScriptDetectionType = "dateTime"
	Win32LobAppPowerShellScriptDetectionType_Float         Win32LobAppPowerShellScriptDetectionType = "float"
	Win32LobAppPowerShellScriptDetectionType_Integer       Win32LobAppPowerShellScriptDetectionType = "integer"
	Win32LobAppPowerShellScriptDetectionType_NotConfigured Win32LobAppPowerShellScriptDetectionType = "notConfigured"
	Win32LobAppPowerShellScriptDetectionType_String        Win32LobAppPowerShellScriptDetectionType = "string"
	Win32LobAppPowerShellScriptDetectionType_Version       Win32LobAppPowerShellScriptDetectionType = "version"
)

func PossibleValuesForWin32LobAppPowerShellScriptDetectionType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptDetectionType_Boolean),
		string(Win32LobAppPowerShellScriptDetectionType_DateTime),
		string(Win32LobAppPowerShellScriptDetectionType_Float),
		string(Win32LobAppPowerShellScriptDetectionType_Integer),
		string(Win32LobAppPowerShellScriptDetectionType_NotConfigured),
		string(Win32LobAppPowerShellScriptDetectionType_String),
		string(Win32LobAppPowerShellScriptDetectionType_Version),
	}
}

func (s *Win32LobAppPowerShellScriptDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptDetectionType(input string) (*Win32LobAppPowerShellScriptDetectionType, error) {
	vals := map[string]Win32LobAppPowerShellScriptDetectionType{
		"boolean":       Win32LobAppPowerShellScriptDetectionType_Boolean,
		"datetime":      Win32LobAppPowerShellScriptDetectionType_DateTime,
		"float":         Win32LobAppPowerShellScriptDetectionType_Float,
		"integer":       Win32LobAppPowerShellScriptDetectionType_Integer,
		"notconfigured": Win32LobAppPowerShellScriptDetectionType_NotConfigured,
		"string":        Win32LobAppPowerShellScriptDetectionType_String,
		"version":       Win32LobAppPowerShellScriptDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptDetectionType(input)
	return &out, nil
}
