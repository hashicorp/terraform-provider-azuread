package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointPreference string

const (
	RestorePointPreference_Latest RestorePointPreference = "latest"
	RestorePointPreference_Oldest RestorePointPreference = "oldest"
)

func PossibleValuesForRestorePointPreference() []string {
	return []string{
		string(RestorePointPreference_Latest),
		string(RestorePointPreference_Oldest),
	}
}

func (s *RestorePointPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestorePointPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestorePointPreference(input string) (*RestorePointPreference, error) {
	vals := map[string]RestorePointPreference{
		"latest": RestorePointPreference_Latest,
		"oldest": RestorePointPreference_Oldest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestorePointPreference(input)
	return &out, nil
}
