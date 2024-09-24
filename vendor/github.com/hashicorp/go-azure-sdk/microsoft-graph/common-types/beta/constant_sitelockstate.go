package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteLockState string

const (
	SiteLockState_LockedNoAccess    SiteLockState = "lockedNoAccess"
	SiteLockState_LockedNoAdditions SiteLockState = "lockedNoAdditions"
	SiteLockState_LockedReadOnly    SiteLockState = "lockedReadOnly"
	SiteLockState_Unlocked          SiteLockState = "unlocked"
)

func PossibleValuesForSiteLockState() []string {
	return []string{
		string(SiteLockState_LockedNoAccess),
		string(SiteLockState_LockedNoAdditions),
		string(SiteLockState_LockedReadOnly),
		string(SiteLockState_Unlocked),
	}
}

func (s *SiteLockState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteLockState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteLockState(input string) (*SiteLockState, error) {
	vals := map[string]SiteLockState{
		"lockednoaccess":    SiteLockState_LockedNoAccess,
		"lockednoadditions": SiteLockState_LockedNoAdditions,
		"lockedreadonly":    SiteLockState_LockedReadOnly,
		"unlocked":          SiteLockState_Unlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteLockState(input)
	return &out, nil
}
