package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchOn string

const (
	MatchOn_DisplayName    MatchOn = "displayName"
	MatchOn_SamAccountName MatchOn = "samAccountName"
)

func PossibleValuesForMatchOn() []string {
	return []string{
		string(MatchOn_DisplayName),
		string(MatchOn_SamAccountName),
	}
}

func (s *MatchOn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMatchOn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMatchOn(input string) (*MatchOn, error) {
	vals := map[string]MatchOn{
		"displayname":    MatchOn_DisplayName,
		"samaccountname": MatchOn_SamAccountName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchOn(input)
	return &out, nil
}
