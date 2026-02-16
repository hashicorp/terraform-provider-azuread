package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiPackageType string

const (
	Win32LobAppMsiPackageType_DualPurpose Win32LobAppMsiPackageType = "dualPurpose"
	Win32LobAppMsiPackageType_PerMachine  Win32LobAppMsiPackageType = "perMachine"
	Win32LobAppMsiPackageType_PerUser     Win32LobAppMsiPackageType = "perUser"
)

func PossibleValuesForWin32LobAppMsiPackageType() []string {
	return []string{
		string(Win32LobAppMsiPackageType_DualPurpose),
		string(Win32LobAppMsiPackageType_PerMachine),
		string(Win32LobAppMsiPackageType_PerUser),
	}
}

func (s *Win32LobAppMsiPackageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppMsiPackageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppMsiPackageType(input string) (*Win32LobAppMsiPackageType, error) {
	vals := map[string]Win32LobAppMsiPackageType{
		"dualpurpose": Win32LobAppMsiPackageType_DualPurpose,
		"permachine":  Win32LobAppMsiPackageType_PerMachine,
		"peruser":     Win32LobAppMsiPackageType_PerUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppMsiPackageType(input)
	return &out, nil
}
