package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentAvailabilityOptions string

const (
	EnrollmentAvailabilityOptions_AvailableWithPrompts    EnrollmentAvailabilityOptions = "availableWithPrompts"
	EnrollmentAvailabilityOptions_AvailableWithoutPrompts EnrollmentAvailabilityOptions = "availableWithoutPrompts"
	EnrollmentAvailabilityOptions_Unavailable             EnrollmentAvailabilityOptions = "unavailable"
)

func PossibleValuesForEnrollmentAvailabilityOptions() []string {
	return []string{
		string(EnrollmentAvailabilityOptions_AvailableWithPrompts),
		string(EnrollmentAvailabilityOptions_AvailableWithoutPrompts),
		string(EnrollmentAvailabilityOptions_Unavailable),
	}
}

func (s *EnrollmentAvailabilityOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentAvailabilityOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentAvailabilityOptions(input string) (*EnrollmentAvailabilityOptions, error) {
	vals := map[string]EnrollmentAvailabilityOptions{
		"availablewithprompts":    EnrollmentAvailabilityOptions_AvailableWithPrompts,
		"availablewithoutprompts": EnrollmentAvailabilityOptions_AvailableWithoutPrompts,
		"unavailable":             EnrollmentAvailabilityOptions_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentAvailabilityOptions(input)
	return &out, nil
}
