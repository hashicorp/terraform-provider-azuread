package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiamondModel string

const (
	DiamondModel_Adversary      DiamondModel = "adversary"
	DiamondModel_Capability     DiamondModel = "capability"
	DiamondModel_Infrastructure DiamondModel = "infrastructure"
	DiamondModel_Unknown        DiamondModel = "unknown"
	DiamondModel_Victim         DiamondModel = "victim"
)

func PossibleValuesForDiamondModel() []string {
	return []string{
		string(DiamondModel_Adversary),
		string(DiamondModel_Capability),
		string(DiamondModel_Infrastructure),
		string(DiamondModel_Unknown),
		string(DiamondModel_Victim),
	}
}

func (s *DiamondModel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiamondModel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiamondModel(input string) (*DiamondModel, error) {
	vals := map[string]DiamondModel{
		"adversary":      DiamondModel_Adversary,
		"capability":     DiamondModel_Capability,
		"infrastructure": DiamondModel_Infrastructure,
		"unknown":        DiamondModel_Unknown,
		"victim":         DiamondModel_Victim,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiamondModel(input)
	return &out, nil
}
