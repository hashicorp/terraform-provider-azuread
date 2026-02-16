package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidProfileApplicability string

const (
	AndroidProfileApplicability_AndroidDeviceOwner AndroidProfileApplicability = "androidDeviceOwner"
	AndroidProfileApplicability_AndroidWorkProfile AndroidProfileApplicability = "androidWorkProfile"
	AndroidProfileApplicability_Default            AndroidProfileApplicability = "default"
)

func PossibleValuesForAndroidProfileApplicability() []string {
	return []string{
		string(AndroidProfileApplicability_AndroidDeviceOwner),
		string(AndroidProfileApplicability_AndroidWorkProfile),
		string(AndroidProfileApplicability_Default),
	}
}

func (s *AndroidProfileApplicability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidProfileApplicability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidProfileApplicability(input string) (*AndroidProfileApplicability, error) {
	vals := map[string]AndroidProfileApplicability{
		"androiddeviceowner": AndroidProfileApplicability_AndroidDeviceOwner,
		"androidworkprofile": AndroidProfileApplicability_AndroidWorkProfile,
		"default":            AndroidProfileApplicability_Default,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidProfileApplicability(input)
	return &out, nil
}
