package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MsiType string

const (
	MsiType_None           MsiType = "none"
	MsiType_SystemAssigned MsiType = "systemAssigned"
	MsiType_UserAssigned   MsiType = "userAssigned"
)

func PossibleValuesForMsiType() []string {
	return []string{
		string(MsiType_None),
		string(MsiType_SystemAssigned),
		string(MsiType_UserAssigned),
	}
}

func (s *MsiType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMsiType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMsiType(input string) (*MsiType, error) {
	vals := map[string]MsiType{
		"none":           MsiType_None,
		"systemassigned": MsiType_SystemAssigned,
		"userassigned":   MsiType_UserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MsiType(input)
	return &out, nil
}
