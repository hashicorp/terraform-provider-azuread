package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementFeatures string

const (
	ManagedDeviceManagementFeatures_MicrosoftManagedDesktop ManagedDeviceManagementFeatures = "microsoftManagedDesktop"
	ManagedDeviceManagementFeatures_None                    ManagedDeviceManagementFeatures = "none"
)

func PossibleValuesForManagedDeviceManagementFeatures() []string {
	return []string{
		string(ManagedDeviceManagementFeatures_MicrosoftManagedDesktop),
		string(ManagedDeviceManagementFeatures_None),
	}
}

func (s *ManagedDeviceManagementFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementFeatures(input string) (*ManagedDeviceManagementFeatures, error) {
	vals := map[string]ManagedDeviceManagementFeatures{
		"microsoftmanageddesktop": ManagedDeviceManagementFeatures_MicrosoftManagedDesktop,
		"none":                    ManagedDeviceManagementFeatures_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementFeatures(input)
	return &out, nil
}
