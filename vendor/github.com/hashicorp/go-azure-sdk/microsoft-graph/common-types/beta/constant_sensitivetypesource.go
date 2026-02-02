package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeSource string

const (
	SensitiveTypeSource_OutOfBox SensitiveTypeSource = "outOfBox"
	SensitiveTypeSource_Tenant   SensitiveTypeSource = "tenant"
)

func PossibleValuesForSensitiveTypeSource() []string {
	return []string{
		string(SensitiveTypeSource_OutOfBox),
		string(SensitiveTypeSource_Tenant),
	}
}

func (s *SensitiveTypeSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitiveTypeSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitiveTypeSource(input string) (*SensitiveTypeSource, error) {
	vals := map[string]SensitiveTypeSource{
		"outofbox": SensitiveTypeSource_OutOfBox,
		"tenant":   SensitiveTypeSource_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeSource(input)
	return &out, nil
}
