package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSyncSetting string

const (
	BrowserSyncSetting_Blocked                 BrowserSyncSetting = "blocked"
	BrowserSyncSetting_BlockedWithUserOverride BrowserSyncSetting = "blockedWithUserOverride"
	BrowserSyncSetting_NotConfigured           BrowserSyncSetting = "notConfigured"
)

func PossibleValuesForBrowserSyncSetting() []string {
	return []string{
		string(BrowserSyncSetting_Blocked),
		string(BrowserSyncSetting_BlockedWithUserOverride),
		string(BrowserSyncSetting_NotConfigured),
	}
}

func (s *BrowserSyncSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSyncSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSyncSetting(input string) (*BrowserSyncSetting, error) {
	vals := map[string]BrowserSyncSetting{
		"blocked":                 BrowserSyncSetting_Blocked,
		"blockedwithuseroverride": BrowserSyncSetting_BlockedWithUserOverride,
		"notconfigured":           BrowserSyncSetting_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSyncSetting(input)
	return &out, nil
}
