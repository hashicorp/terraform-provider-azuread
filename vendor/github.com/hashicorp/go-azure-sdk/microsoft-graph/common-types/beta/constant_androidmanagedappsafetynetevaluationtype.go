package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppSafetyNetEvaluationType string

const (
	AndroidManagedAppSafetyNetEvaluationType_Basic          AndroidManagedAppSafetyNetEvaluationType = "basic"
	AndroidManagedAppSafetyNetEvaluationType_HardwareBacked AndroidManagedAppSafetyNetEvaluationType = "hardwareBacked"
)

func PossibleValuesForAndroidManagedAppSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidManagedAppSafetyNetEvaluationType_Basic),
		string(AndroidManagedAppSafetyNetEvaluationType_HardwareBacked),
	}
}

func (s *AndroidManagedAppSafetyNetEvaluationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppSafetyNetEvaluationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppSafetyNetEvaluationType(input string) (*AndroidManagedAppSafetyNetEvaluationType, error) {
	vals := map[string]AndroidManagedAppSafetyNetEvaluationType{
		"basic":          AndroidManagedAppSafetyNetEvaluationType_Basic,
		"hardwarebacked": AndroidManagedAppSafetyNetEvaluationType_HardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppSafetyNetEvaluationType(input)
	return &out, nil
}
