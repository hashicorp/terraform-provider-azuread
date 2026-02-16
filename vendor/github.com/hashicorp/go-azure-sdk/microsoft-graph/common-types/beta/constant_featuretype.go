package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureType string

const (
	FeatureType_Registration FeatureType = "registration"
	FeatureType_Reset        FeatureType = "reset"
)

func PossibleValuesForFeatureType() []string {
	return []string{
		string(FeatureType_Registration),
		string(FeatureType_Reset),
	}
}

func (s *FeatureType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureType(input string) (*FeatureType, error) {
	vals := map[string]FeatureType{
		"registration": FeatureType_Registration,
		"reset":        FeatureType_Reset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureType(input)
	return &out, nil
}
