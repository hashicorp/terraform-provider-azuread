package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteUserRole string

const (
	OnenoteUserRole_Contributor OnenoteUserRole = "Contributor"
	OnenoteUserRole_None        OnenoteUserRole = "None"
	OnenoteUserRole_Owner       OnenoteUserRole = "Owner"
	OnenoteUserRole_Reader      OnenoteUserRole = "Reader"
)

func PossibleValuesForOnenoteUserRole() []string {
	return []string{
		string(OnenoteUserRole_Contributor),
		string(OnenoteUserRole_None),
		string(OnenoteUserRole_Owner),
		string(OnenoteUserRole_Reader),
	}
}

func (s *OnenoteUserRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenoteUserRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenoteUserRole(input string) (*OnenoteUserRole, error) {
	vals := map[string]OnenoteUserRole{
		"contributor": OnenoteUserRole_Contributor,
		"none":        OnenoteUserRole_None,
		"owner":       OnenoteUserRole_Owner,
		"reader":      OnenoteUserRole_Reader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenoteUserRole(input)
	return &out, nil
}
