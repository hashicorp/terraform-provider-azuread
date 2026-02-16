package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRestartBehavior string

const (
	Win32LobAppRestartBehavior_Allow             Win32LobAppRestartBehavior = "allow"
	Win32LobAppRestartBehavior_BasedOnReturnCode Win32LobAppRestartBehavior = "basedOnReturnCode"
	Win32LobAppRestartBehavior_Force             Win32LobAppRestartBehavior = "force"
	Win32LobAppRestartBehavior_Suppress          Win32LobAppRestartBehavior = "suppress"
)

func PossibleValuesForWin32LobAppRestartBehavior() []string {
	return []string{
		string(Win32LobAppRestartBehavior_Allow),
		string(Win32LobAppRestartBehavior_BasedOnReturnCode),
		string(Win32LobAppRestartBehavior_Force),
		string(Win32LobAppRestartBehavior_Suppress),
	}
}

func (s *Win32LobAppRestartBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRestartBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRestartBehavior(input string) (*Win32LobAppRestartBehavior, error) {
	vals := map[string]Win32LobAppRestartBehavior{
		"allow":             Win32LobAppRestartBehavior_Allow,
		"basedonreturncode": Win32LobAppRestartBehavior_BasedOnReturnCode,
		"force":             Win32LobAppRestartBehavior_Force,
		"suppress":          Win32LobAppRestartBehavior_Suppress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRestartBehavior(input)
	return &out, nil
}
