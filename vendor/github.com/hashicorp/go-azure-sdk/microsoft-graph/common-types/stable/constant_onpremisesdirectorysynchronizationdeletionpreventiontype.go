package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationDeletionPreventionType string

const (
	OnPremisesDirectorySynchronizationDeletionPreventionType_Disabled             OnPremisesDirectorySynchronizationDeletionPreventionType = "disabled"
	OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForCount      OnPremisesDirectorySynchronizationDeletionPreventionType = "enabledForCount"
	OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForPercentage OnPremisesDirectorySynchronizationDeletionPreventionType = "enabledForPercentage"
)

func PossibleValuesForOnPremisesDirectorySynchronizationDeletionPreventionType() []string {
	return []string{
		string(OnPremisesDirectorySynchronizationDeletionPreventionType_Disabled),
		string(OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForCount),
		string(OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForPercentage),
	}
}

func (s *OnPremisesDirectorySynchronizationDeletionPreventionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesDirectorySynchronizationDeletionPreventionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesDirectorySynchronizationDeletionPreventionType(input string) (*OnPremisesDirectorySynchronizationDeletionPreventionType, error) {
	vals := map[string]OnPremisesDirectorySynchronizationDeletionPreventionType{
		"disabled":             OnPremisesDirectorySynchronizationDeletionPreventionType_Disabled,
		"enabledforcount":      OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForCount,
		"enabledforpercentage": OnPremisesDirectorySynchronizationDeletionPreventionType_EnabledForPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesDirectorySynchronizationDeletionPreventionType(input)
	return &out, nil
}
