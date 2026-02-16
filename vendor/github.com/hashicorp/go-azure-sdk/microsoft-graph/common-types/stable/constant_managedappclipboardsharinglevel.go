package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppClipboardSharingLevel string

const (
	ManagedAppClipboardSharingLevel_AllApps                ManagedAppClipboardSharingLevel = "allApps"
	ManagedAppClipboardSharingLevel_Blocked                ManagedAppClipboardSharingLevel = "blocked"
	ManagedAppClipboardSharingLevel_ManagedApps            ManagedAppClipboardSharingLevel = "managedApps"
	ManagedAppClipboardSharingLevel_ManagedAppsWithPasteIn ManagedAppClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForManagedAppClipboardSharingLevel() []string {
	return []string{
		string(ManagedAppClipboardSharingLevel_AllApps),
		string(ManagedAppClipboardSharingLevel_Blocked),
		string(ManagedAppClipboardSharingLevel_ManagedApps),
		string(ManagedAppClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *ManagedAppClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppClipboardSharingLevel(input string) (*ManagedAppClipboardSharingLevel, error) {
	vals := map[string]ManagedAppClipboardSharingLevel{
		"allapps":                ManagedAppClipboardSharingLevel_AllApps,
		"blocked":                ManagedAppClipboardSharingLevel_Blocked,
		"managedapps":            ManagedAppClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": ManagedAppClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppClipboardSharingLevel(input)
	return &out, nil
}
