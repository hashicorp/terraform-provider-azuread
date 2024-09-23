package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SigninFrequencyType string

const (
	SigninFrequencyType_Days  SigninFrequencyType = "days"
	SigninFrequencyType_Hours SigninFrequencyType = "hours"
)

func PossibleValuesForSigninFrequencyType() []string {
	return []string{
		string(SigninFrequencyType_Days),
		string(SigninFrequencyType_Hours),
	}
}

func (s *SigninFrequencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSigninFrequencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSigninFrequencyType(input string) (*SigninFrequencyType, error) {
	vals := map[string]SigninFrequencyType{
		"days":  SigninFrequencyType_Days,
		"hours": SigninFrequencyType_Hours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SigninFrequencyType(input)
	return &out, nil
}
