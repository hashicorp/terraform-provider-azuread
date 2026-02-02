package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfirmedBy string

const (
	ConfirmedBy_Manager ConfirmedBy = "manager"
	ConfirmedBy_None    ConfirmedBy = "none"
	ConfirmedBy_User    ConfirmedBy = "user"
)

func PossibleValuesForConfirmedBy() []string {
	return []string{
		string(ConfirmedBy_Manager),
		string(ConfirmedBy_None),
		string(ConfirmedBy_User),
	}
}

func (s *ConfirmedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfirmedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfirmedBy(input string) (*ConfirmedBy, error) {
	vals := map[string]ConfirmedBy{
		"manager": ConfirmedBy_Manager,
		"none":    ConfirmedBy_None,
		"user":    ConfirmedBy_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfirmedBy(input)
	return &out, nil
}
