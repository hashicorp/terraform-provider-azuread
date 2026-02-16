package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppGroupType string

const (
	TargetedManagedAppGroupType_AllApps              TargetedManagedAppGroupType = "allApps"
	TargetedManagedAppGroupType_AllCoreMicrosoftApps TargetedManagedAppGroupType = "allCoreMicrosoftApps"
	TargetedManagedAppGroupType_AllMicrosoftApps     TargetedManagedAppGroupType = "allMicrosoftApps"
	TargetedManagedAppGroupType_SelectedPublicApps   TargetedManagedAppGroupType = "selectedPublicApps"
)

func PossibleValuesForTargetedManagedAppGroupType() []string {
	return []string{
		string(TargetedManagedAppGroupType_AllApps),
		string(TargetedManagedAppGroupType_AllCoreMicrosoftApps),
		string(TargetedManagedAppGroupType_AllMicrosoftApps),
		string(TargetedManagedAppGroupType_SelectedPublicApps),
	}
}

func (s *TargetedManagedAppGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppGroupType(input string) (*TargetedManagedAppGroupType, error) {
	vals := map[string]TargetedManagedAppGroupType{
		"allapps":              TargetedManagedAppGroupType_AllApps,
		"allcoremicrosoftapps": TargetedManagedAppGroupType_AllCoreMicrosoftApps,
		"allmicrosoftapps":     TargetedManagedAppGroupType_AllMicrosoftApps,
		"selectedpublicapps":   TargetedManagedAppGroupType_SelectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppGroupType(input)
	return &out, nil
}
