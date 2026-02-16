package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryOperation string

const (
	RegistryOperation_Create  RegistryOperation = "create"
	RegistryOperation_Delete  RegistryOperation = "delete"
	RegistryOperation_Modify  RegistryOperation = "modify"
	RegistryOperation_Unknown RegistryOperation = "unknown"
)

func PossibleValuesForRegistryOperation() []string {
	return []string{
		string(RegistryOperation_Create),
		string(RegistryOperation_Delete),
		string(RegistryOperation_Modify),
		string(RegistryOperation_Unknown),
	}
}

func (s *RegistryOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistryOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistryOperation(input string) (*RegistryOperation, error) {
	vals := map[string]RegistryOperation{
		"create":  RegistryOperation_Create,
		"delete":  RegistryOperation_Delete,
		"modify":  RegistryOperation_Modify,
		"unknown": RegistryOperation_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryOperation(input)
	return &out, nil
}
