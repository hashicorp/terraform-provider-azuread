package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryHive string

const (
	RegistryHive_CurrentConfig        RegistryHive = "currentConfig"
	RegistryHive_CurrentUser          RegistryHive = "currentUser"
	RegistryHive_LocalMachineSam      RegistryHive = "localMachineSam"
	RegistryHive_LocalMachineSecurity RegistryHive = "localMachineSecurity"
	RegistryHive_LocalMachineSoftware RegistryHive = "localMachineSoftware"
	RegistryHive_LocalMachineSystem   RegistryHive = "localMachineSystem"
	RegistryHive_Unknown              RegistryHive = "unknown"
	RegistryHive_UsersDefault         RegistryHive = "usersDefault"
)

func PossibleValuesForRegistryHive() []string {
	return []string{
		string(RegistryHive_CurrentConfig),
		string(RegistryHive_CurrentUser),
		string(RegistryHive_LocalMachineSam),
		string(RegistryHive_LocalMachineSecurity),
		string(RegistryHive_LocalMachineSoftware),
		string(RegistryHive_LocalMachineSystem),
		string(RegistryHive_Unknown),
		string(RegistryHive_UsersDefault),
	}
}

func (s *RegistryHive) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryHive(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryHive(input string) (*RegistryHive, error) {
	vals := map[string]RegistryHive{
		"currentconfig":        RegistryHive_CurrentConfig,
		"currentuser":          RegistryHive_CurrentUser,
		"localmachinesam":      RegistryHive_LocalMachineSam,
		"localmachinesecurity": RegistryHive_LocalMachineSecurity,
		"localmachinesoftware": RegistryHive_LocalMachineSoftware,
		"localmachinesystem":   RegistryHive_LocalMachineSystem,
		"unknown":              RegistryHive_Unknown,
		"usersdefault":         RegistryHive_UsersDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryHive(input)
	return &out, nil
}
