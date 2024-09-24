package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreTimeRange string

const (
	RestoreTimeRange_After         RestoreTimeRange = "after"
	RestoreTimeRange_Before        RestoreTimeRange = "before"
	RestoreTimeRange_BeforeOrAfter RestoreTimeRange = "beforeOrAfter"
)

func PossibleValuesForRestoreTimeRange() []string {
	return []string{
		string(RestoreTimeRange_After),
		string(RestoreTimeRange_Before),
		string(RestoreTimeRange_BeforeOrAfter),
	}
}

func (s *RestoreTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestoreTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestoreTimeRange(input string) (*RestoreTimeRange, error) {
	vals := map[string]RestoreTimeRange{
		"after":         RestoreTimeRange_After,
		"before":        RestoreTimeRange_Before,
		"beforeorafter": RestoreTimeRange_BeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreTimeRange(input)
	return &out, nil
}
