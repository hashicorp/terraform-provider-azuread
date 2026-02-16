package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryDetectionType string

const (
	Win32LobAppRegistryDetectionType_DoesNotExist  Win32LobAppRegistryDetectionType = "doesNotExist"
	Win32LobAppRegistryDetectionType_Exists        Win32LobAppRegistryDetectionType = "exists"
	Win32LobAppRegistryDetectionType_Integer       Win32LobAppRegistryDetectionType = "integer"
	Win32LobAppRegistryDetectionType_NotConfigured Win32LobAppRegistryDetectionType = "notConfigured"
	Win32LobAppRegistryDetectionType_String        Win32LobAppRegistryDetectionType = "string"
	Win32LobAppRegistryDetectionType_Version       Win32LobAppRegistryDetectionType = "version"
)

func PossibleValuesForWin32LobAppRegistryDetectionType() []string {
	return []string{
		string(Win32LobAppRegistryDetectionType_DoesNotExist),
		string(Win32LobAppRegistryDetectionType_Exists),
		string(Win32LobAppRegistryDetectionType_Integer),
		string(Win32LobAppRegistryDetectionType_NotConfigured),
		string(Win32LobAppRegistryDetectionType_String),
		string(Win32LobAppRegistryDetectionType_Version),
	}
}

func (s *Win32LobAppRegistryDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryDetectionType(input string) (*Win32LobAppRegistryDetectionType, error) {
	vals := map[string]Win32LobAppRegistryDetectionType{
		"doesnotexist":  Win32LobAppRegistryDetectionType_DoesNotExist,
		"exists":        Win32LobAppRegistryDetectionType_Exists,
		"integer":       Win32LobAppRegistryDetectionType_Integer,
		"notconfigured": Win32LobAppRegistryDetectionType_NotConfigured,
		"string":        Win32LobAppRegistryDetectionType_String,
		"version":       Win32LobAppRegistryDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryDetectionType(input)
	return &out, nil
}
