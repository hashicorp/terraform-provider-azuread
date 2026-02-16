package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementProvider string

const (
	ManagedTenantsManagementProvider_Community        ManagedTenantsManagementProvider = "community"
	ManagedTenantsManagementProvider_IndirectProvider ManagedTenantsManagementProvider = "indirectProvider"
	ManagedTenantsManagementProvider_Microsoft        ManagedTenantsManagementProvider = "microsoft"
	ManagedTenantsManagementProvider_Self             ManagedTenantsManagementProvider = "self"
)

func PossibleValuesForManagedTenantsManagementProvider() []string {
	return []string{
		string(ManagedTenantsManagementProvider_Community),
		string(ManagedTenantsManagementProvider_IndirectProvider),
		string(ManagedTenantsManagementProvider_Microsoft),
		string(ManagedTenantsManagementProvider_Self),
	}
}

func (s *ManagedTenantsManagementProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementProvider(input string) (*ManagedTenantsManagementProvider, error) {
	vals := map[string]ManagedTenantsManagementProvider{
		"community":        ManagedTenantsManagementProvider_Community,
		"indirectprovider": ManagedTenantsManagementProvider_IndirectProvider,
		"microsoft":        ManagedTenantsManagementProvider_Microsoft,
		"self":             ManagedTenantsManagementProvider_Self,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementProvider(input)
	return &out, nil
}
