package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidSafetyNetEvaluationType string

const (
	AndroidSafetyNetEvaluationType_Basic          AndroidSafetyNetEvaluationType = "basic"
	AndroidSafetyNetEvaluationType_HardwareBacked AndroidSafetyNetEvaluationType = "hardwareBacked"
)

func PossibleValuesForAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidSafetyNetEvaluationType_Basic),
		string(AndroidSafetyNetEvaluationType_HardwareBacked),
	}
}

func (s *AndroidSafetyNetEvaluationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidSafetyNetEvaluationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidSafetyNetEvaluationType(input string) (*AndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidSafetyNetEvaluationType{
		"basic":          AndroidSafetyNetEvaluationType_Basic,
		"hardwarebacked": AndroidSafetyNetEvaluationType_HardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidSafetyNetEvaluationType(input)
	return &out, nil
}
