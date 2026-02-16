package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessType string

const (
	AccessType_Deny  AccessType = "deny"
	AccessType_Grant AccessType = "grant"
)

func PossibleValuesForAccessType() []string {
	return []string{
		string(AccessType_Deny),
		string(AccessType_Grant),
	}
}

func (s *AccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessType(input string) (*AccessType, error) {
	vals := map[string]AccessType{
		"deny":  AccessType_Deny,
		"grant": AccessType_Grant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessType(input)
	return &out, nil
}
