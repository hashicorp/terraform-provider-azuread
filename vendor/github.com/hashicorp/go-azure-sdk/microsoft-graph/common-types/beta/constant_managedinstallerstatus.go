package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstallerStatus string

const (
	ManagedInstallerStatus_Disabled ManagedInstallerStatus = "disabled"
	ManagedInstallerStatus_Enabled  ManagedInstallerStatus = "enabled"
)

func PossibleValuesForManagedInstallerStatus() []string {
	return []string{
		string(ManagedInstallerStatus_Disabled),
		string(ManagedInstallerStatus_Enabled),
	}
}

func (s *ManagedInstallerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedInstallerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedInstallerStatus(input string) (*ManagedInstallerStatus, error) {
	vals := map[string]ManagedInstallerStatus{
		"disabled": ManagedInstallerStatus_Disabled,
		"enabled":  ManagedInstallerStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedInstallerStatus(input)
	return &out, nil
}
