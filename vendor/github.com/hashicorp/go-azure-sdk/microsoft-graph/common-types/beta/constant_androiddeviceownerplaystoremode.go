package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPlayStoreMode string

const (
	AndroidDeviceOwnerPlayStoreMode_AllowList     AndroidDeviceOwnerPlayStoreMode = "allowList"
	AndroidDeviceOwnerPlayStoreMode_BlockList     AndroidDeviceOwnerPlayStoreMode = "blockList"
	AndroidDeviceOwnerPlayStoreMode_NotConfigured AndroidDeviceOwnerPlayStoreMode = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerPlayStoreMode() []string {
	return []string{
		string(AndroidDeviceOwnerPlayStoreMode_AllowList),
		string(AndroidDeviceOwnerPlayStoreMode_BlockList),
		string(AndroidDeviceOwnerPlayStoreMode_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerPlayStoreMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPlayStoreMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPlayStoreMode(input string) (*AndroidDeviceOwnerPlayStoreMode, error) {
	vals := map[string]AndroidDeviceOwnerPlayStoreMode{
		"allowlist":     AndroidDeviceOwnerPlayStoreMode_AllowList,
		"blocklist":     AndroidDeviceOwnerPlayStoreMode_BlockList,
		"notconfigured": AndroidDeviceOwnerPlayStoreMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPlayStoreMode(input)
	return &out, nil
}
