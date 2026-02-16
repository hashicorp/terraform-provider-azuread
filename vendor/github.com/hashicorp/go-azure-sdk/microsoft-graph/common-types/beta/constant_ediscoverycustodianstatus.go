package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCustodianStatus string

const (
	EdiscoveryCustodianStatus_Active   EdiscoveryCustodianStatus = "active"
	EdiscoveryCustodianStatus_Released EdiscoveryCustodianStatus = "released"
)

func PossibleValuesForEdiscoveryCustodianStatus() []string {
	return []string{
		string(EdiscoveryCustodianStatus_Active),
		string(EdiscoveryCustodianStatus_Released),
	}
}

func (s *EdiscoveryCustodianStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCustodianStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCustodianStatus(input string) (*EdiscoveryCustodianStatus, error) {
	vals := map[string]EdiscoveryCustodianStatus{
		"active":   EdiscoveryCustodianStatus_Active,
		"released": EdiscoveryCustodianStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCustodianStatus(input)
	return &out, nil
}
