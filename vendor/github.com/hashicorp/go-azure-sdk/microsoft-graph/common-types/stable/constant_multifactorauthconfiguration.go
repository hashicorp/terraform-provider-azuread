package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiFactorAuthConfiguration string

const (
	MultiFactorAuthConfiguration_NotRequired MultiFactorAuthConfiguration = "notRequired"
	MultiFactorAuthConfiguration_Required    MultiFactorAuthConfiguration = "required"
)

func PossibleValuesForMultiFactorAuthConfiguration() []string {
	return []string{
		string(MultiFactorAuthConfiguration_NotRequired),
		string(MultiFactorAuthConfiguration_Required),
	}
}

func (s *MultiFactorAuthConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiFactorAuthConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiFactorAuthConfiguration(input string) (*MultiFactorAuthConfiguration, error) {
	vals := map[string]MultiFactorAuthConfiguration{
		"notrequired": MultiFactorAuthConfiguration_NotRequired,
		"required":    MultiFactorAuthConfiguration_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiFactorAuthConfiguration(input)
	return &out, nil
}
