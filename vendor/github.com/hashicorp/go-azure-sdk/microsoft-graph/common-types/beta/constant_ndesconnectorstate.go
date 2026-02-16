package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NdesConnectorState string

const (
	NdesConnectorState_Active   NdesConnectorState = "active"
	NdesConnectorState_Inactive NdesConnectorState = "inactive"
	NdesConnectorState_None     NdesConnectorState = "none"
)

func PossibleValuesForNdesConnectorState() []string {
	return []string{
		string(NdesConnectorState_Active),
		string(NdesConnectorState_Inactive),
		string(NdesConnectorState_None),
	}
}

func (s *NdesConnectorState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNdesConnectorState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNdesConnectorState(input string) (*NdesConnectorState, error) {
	vals := map[string]NdesConnectorState{
		"active":   NdesConnectorState_Active,
		"inactive": NdesConnectorState_Inactive,
		"none":     NdesConnectorState_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NdesConnectorState(input)
	return &out, nil
}
