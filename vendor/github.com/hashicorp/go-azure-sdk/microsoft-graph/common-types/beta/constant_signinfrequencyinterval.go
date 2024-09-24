package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencyInterval string

const (
	SignInFrequencyInterval_EveryTime SignInFrequencyInterval = "everyTime"
	SignInFrequencyInterval_TimeBased SignInFrequencyInterval = "timeBased"
)

func PossibleValuesForSignInFrequencyInterval() []string {
	return []string{
		string(SignInFrequencyInterval_EveryTime),
		string(SignInFrequencyInterval_TimeBased),
	}
}

func (s *SignInFrequencyInterval) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInFrequencyInterval(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInFrequencyInterval(input string) (*SignInFrequencyInterval, error) {
	vals := map[string]SignInFrequencyInterval{
		"everytime": SignInFrequencyInterval_EveryTime,
		"timebased": SignInFrequencyInterval_TimeBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencyInterval(input)
	return &out, nil
}
