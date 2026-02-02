package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReleaseType string

const (
	ReleaseType_GenerallyAvailable ReleaseType = "generallyAvailable"
	ReleaseType_Preview            ReleaseType = "preview"
)

func PossibleValuesForReleaseType() []string {
	return []string{
		string(ReleaseType_GenerallyAvailable),
		string(ReleaseType_Preview),
	}
}

func (s *ReleaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReleaseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReleaseType(input string) (*ReleaseType, error) {
	vals := map[string]ReleaseType{
		"generallyavailable": ReleaseType_GenerallyAvailable,
		"preview":            ReleaseType_Preview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReleaseType(input)
	return &out, nil
}
