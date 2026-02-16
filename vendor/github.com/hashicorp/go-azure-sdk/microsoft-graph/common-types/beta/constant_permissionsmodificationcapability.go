package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsModificationCapability string

const (
	PermissionsModificationCapability_Enabled               PermissionsModificationCapability = "enabled"
	PermissionsModificationCapability_NoRecentDataCollected PermissionsModificationCapability = "noRecentDataCollected"
	PermissionsModificationCapability_NotConfigured         PermissionsModificationCapability = "notConfigured"
)

func PossibleValuesForPermissionsModificationCapability() []string {
	return []string{
		string(PermissionsModificationCapability_Enabled),
		string(PermissionsModificationCapability_NoRecentDataCollected),
		string(PermissionsModificationCapability_NotConfigured),
	}
}

func (s *PermissionsModificationCapability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsModificationCapability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsModificationCapability(input string) (*PermissionsModificationCapability, error) {
	vals := map[string]PermissionsModificationCapability{
		"enabled":               PermissionsModificationCapability_Enabled,
		"norecentdatacollected": PermissionsModificationCapability_NoRecentDataCollected,
		"notconfigured":         PermissionsModificationCapability_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsModificationCapability(input)
	return &out, nil
}
