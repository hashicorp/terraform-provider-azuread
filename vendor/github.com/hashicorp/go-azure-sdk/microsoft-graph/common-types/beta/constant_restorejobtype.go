package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreJobType string

const (
	RestoreJobType_Bulk     RestoreJobType = "bulk"
	RestoreJobType_Standard RestoreJobType = "standard"
)

func PossibleValuesForRestoreJobType() []string {
	return []string{
		string(RestoreJobType_Bulk),
		string(RestoreJobType_Standard),
	}
}

func (s *RestoreJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestoreJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestoreJobType(input string) (*RestoreJobType, error) {
	vals := map[string]RestoreJobType{
		"bulk":     RestoreJobType_Bulk,
		"standard": RestoreJobType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreJobType(input)
	return &out, nil
}
