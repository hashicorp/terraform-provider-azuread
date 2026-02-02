package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementCategory string

const (
	ManagedTenantsManagementCategory_Custom   ManagedTenantsManagementCategory = "custom"
	ManagedTenantsManagementCategory_Data     ManagedTenantsManagementCategory = "data"
	ManagedTenantsManagementCategory_Devices  ManagedTenantsManagementCategory = "devices"
	ManagedTenantsManagementCategory_Identity ManagedTenantsManagementCategory = "identity"
)

func PossibleValuesForManagedTenantsManagementCategory() []string {
	return []string{
		string(ManagedTenantsManagementCategory_Custom),
		string(ManagedTenantsManagementCategory_Data),
		string(ManagedTenantsManagementCategory_Devices),
		string(ManagedTenantsManagementCategory_Identity),
	}
}

func (s *ManagedTenantsManagementCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementCategory(input string) (*ManagedTenantsManagementCategory, error) {
	vals := map[string]ManagedTenantsManagementCategory{
		"custom":   ManagedTenantsManagementCategory_Custom,
		"data":     ManagedTenantsManagementCategory_Data,
		"devices":  ManagedTenantsManagementCategory_Devices,
		"identity": ManagedTenantsManagementCategory_Identity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementCategory(input)
	return &out, nil
}
