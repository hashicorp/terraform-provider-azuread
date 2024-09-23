package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDataTransferLevel string

const (
	ManagedAppDataTransferLevel_AllApps     ManagedAppDataTransferLevel = "allApps"
	ManagedAppDataTransferLevel_ManagedApps ManagedAppDataTransferLevel = "managedApps"
	ManagedAppDataTransferLevel_None        ManagedAppDataTransferLevel = "none"
)

func PossibleValuesForManagedAppDataTransferLevel() []string {
	return []string{
		string(ManagedAppDataTransferLevel_AllApps),
		string(ManagedAppDataTransferLevel_ManagedApps),
		string(ManagedAppDataTransferLevel_None),
	}
}

func (s *ManagedAppDataTransferLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppDataTransferLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppDataTransferLevel(input string) (*ManagedAppDataTransferLevel, error) {
	vals := map[string]ManagedAppDataTransferLevel{
		"allapps":     ManagedAppDataTransferLevel_AllApps,
		"managedapps": ManagedAppDataTransferLevel_ManagedApps,
		"none":        ManagedAppDataTransferLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppDataTransferLevel(input)
	return &out, nil
}
