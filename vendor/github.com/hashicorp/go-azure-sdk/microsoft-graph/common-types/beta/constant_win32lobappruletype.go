package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRuleType string

const (
	Win32LobAppRuleType_Detection   Win32LobAppRuleType = "detection"
	Win32LobAppRuleType_Requirement Win32LobAppRuleType = "requirement"
)

func PossibleValuesForWin32LobAppRuleType() []string {
	return []string{
		string(Win32LobAppRuleType_Detection),
		string(Win32LobAppRuleType_Requirement),
	}
}

func (s *Win32LobAppRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRuleType(input string) (*Win32LobAppRuleType, error) {
	vals := map[string]Win32LobAppRuleType{
		"detection":   Win32LobAppRuleType_Detection,
		"requirement": Win32LobAppRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRuleType(input)
	return &out, nil
}
