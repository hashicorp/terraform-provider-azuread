package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRuleOperationType string

const (
	Win32LobAppRegistryRuleOperationType_AppVersion    Win32LobAppRegistryRuleOperationType = "appVersion"
	Win32LobAppRegistryRuleOperationType_DoesNotExist  Win32LobAppRegistryRuleOperationType = "doesNotExist"
	Win32LobAppRegistryRuleOperationType_Exists        Win32LobAppRegistryRuleOperationType = "exists"
	Win32LobAppRegistryRuleOperationType_Integer       Win32LobAppRegistryRuleOperationType = "integer"
	Win32LobAppRegistryRuleOperationType_NotConfigured Win32LobAppRegistryRuleOperationType = "notConfigured"
	Win32LobAppRegistryRuleOperationType_String        Win32LobAppRegistryRuleOperationType = "string"
	Win32LobAppRegistryRuleOperationType_Version       Win32LobAppRegistryRuleOperationType = "version"
)

func PossibleValuesForWin32LobAppRegistryRuleOperationType() []string {
	return []string{
		string(Win32LobAppRegistryRuleOperationType_AppVersion),
		string(Win32LobAppRegistryRuleOperationType_DoesNotExist),
		string(Win32LobAppRegistryRuleOperationType_Exists),
		string(Win32LobAppRegistryRuleOperationType_Integer),
		string(Win32LobAppRegistryRuleOperationType_NotConfigured),
		string(Win32LobAppRegistryRuleOperationType_String),
		string(Win32LobAppRegistryRuleOperationType_Version),
	}
}

func (s *Win32LobAppRegistryRuleOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryRuleOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryRuleOperationType(input string) (*Win32LobAppRegistryRuleOperationType, error) {
	vals := map[string]Win32LobAppRegistryRuleOperationType{
		"appversion":    Win32LobAppRegistryRuleOperationType_AppVersion,
		"doesnotexist":  Win32LobAppRegistryRuleOperationType_DoesNotExist,
		"exists":        Win32LobAppRegistryRuleOperationType_Exists,
		"integer":       Win32LobAppRegistryRuleOperationType_Integer,
		"notconfigured": Win32LobAppRegistryRuleOperationType_NotConfigured,
		"string":        Win32LobAppRegistryRuleOperationType_String,
		"version":       Win32LobAppRegistryRuleOperationType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRuleOperationType(input)
	return &out, nil
}
