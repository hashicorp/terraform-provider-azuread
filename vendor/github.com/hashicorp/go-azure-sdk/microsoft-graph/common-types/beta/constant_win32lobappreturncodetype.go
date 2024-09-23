package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppReturnCodeType string

const (
	Win32LobAppReturnCodeType_Failed     Win32LobAppReturnCodeType = "failed"
	Win32LobAppReturnCodeType_HardReboot Win32LobAppReturnCodeType = "hardReboot"
	Win32LobAppReturnCodeType_Retry      Win32LobAppReturnCodeType = "retry"
	Win32LobAppReturnCodeType_SoftReboot Win32LobAppReturnCodeType = "softReboot"
	Win32LobAppReturnCodeType_Success    Win32LobAppReturnCodeType = "success"
)

func PossibleValuesForWin32LobAppReturnCodeType() []string {
	return []string{
		string(Win32LobAppReturnCodeType_Failed),
		string(Win32LobAppReturnCodeType_HardReboot),
		string(Win32LobAppReturnCodeType_Retry),
		string(Win32LobAppReturnCodeType_SoftReboot),
		string(Win32LobAppReturnCodeType_Success),
	}
}

func (s *Win32LobAppReturnCodeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppReturnCodeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppReturnCodeType(input string) (*Win32LobAppReturnCodeType, error) {
	vals := map[string]Win32LobAppReturnCodeType{
		"failed":     Win32LobAppReturnCodeType_Failed,
		"hardreboot": Win32LobAppReturnCodeType_HardReboot,
		"retry":      Win32LobAppReturnCodeType_Retry,
		"softreboot": Win32LobAppReturnCodeType_SoftReboot,
		"success":    Win32LobAppReturnCodeType_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppReturnCodeType(input)
	return &out, nil
}
