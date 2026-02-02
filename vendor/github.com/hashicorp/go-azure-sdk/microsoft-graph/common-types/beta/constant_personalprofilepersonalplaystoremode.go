package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonalProfilePersonalPlayStoreMode string

const (
	PersonalProfilePersonalPlayStoreMode_AllowedApps   PersonalProfilePersonalPlayStoreMode = "allowedApps"
	PersonalProfilePersonalPlayStoreMode_BlockedApps   PersonalProfilePersonalPlayStoreMode = "blockedApps"
	PersonalProfilePersonalPlayStoreMode_NotConfigured PersonalProfilePersonalPlayStoreMode = "notConfigured"
)

func PossibleValuesForPersonalProfilePersonalPlayStoreMode() []string {
	return []string{
		string(PersonalProfilePersonalPlayStoreMode_AllowedApps),
		string(PersonalProfilePersonalPlayStoreMode_BlockedApps),
		string(PersonalProfilePersonalPlayStoreMode_NotConfigured),
	}
}

func (s *PersonalProfilePersonalPlayStoreMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonalProfilePersonalPlayStoreMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonalProfilePersonalPlayStoreMode(input string) (*PersonalProfilePersonalPlayStoreMode, error) {
	vals := map[string]PersonalProfilePersonalPlayStoreMode{
		"allowedapps":   PersonalProfilePersonalPlayStoreMode_AllowedApps,
		"blockedapps":   PersonalProfilePersonalPlayStoreMode_BlockedApps,
		"notconfigured": PersonalProfilePersonalPlayStoreMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonalProfilePersonalPlayStoreMode(input)
	return &out, nil
}
