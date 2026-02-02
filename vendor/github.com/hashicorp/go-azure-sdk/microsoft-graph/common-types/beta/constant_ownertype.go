package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OwnerType string

const (
	OwnerType_Company  OwnerType = "company"
	OwnerType_Personal OwnerType = "personal"
	OwnerType_Unknown  OwnerType = "unknown"
)

func PossibleValuesForOwnerType() []string {
	return []string{
		string(OwnerType_Company),
		string(OwnerType_Personal),
		string(OwnerType_Unknown),
	}
}

func (s *OwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOwnerType(input string) (*OwnerType, error) {
	vals := map[string]OwnerType{
		"company":  OwnerType_Company,
		"personal": OwnerType_Personal,
		"unknown":  OwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OwnerType(input)
	return &out, nil
}
