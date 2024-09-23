package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteInstallProgressDisplayLevel string

const (
	OfficeSuiteInstallProgressDisplayLevel_Full OfficeSuiteInstallProgressDisplayLevel = "full"
	OfficeSuiteInstallProgressDisplayLevel_None OfficeSuiteInstallProgressDisplayLevel = "none"
)

func PossibleValuesForOfficeSuiteInstallProgressDisplayLevel() []string {
	return []string{
		string(OfficeSuiteInstallProgressDisplayLevel_Full),
		string(OfficeSuiteInstallProgressDisplayLevel_None),
	}
}

func (s *OfficeSuiteInstallProgressDisplayLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteInstallProgressDisplayLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteInstallProgressDisplayLevel(input string) (*OfficeSuiteInstallProgressDisplayLevel, error) {
	vals := map[string]OfficeSuiteInstallProgressDisplayLevel{
		"full": OfficeSuiteInstallProgressDisplayLevel_Full,
		"none": OfficeSuiteInstallProgressDisplayLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteInstallProgressDisplayLevel(input)
	return &out, nil
}
