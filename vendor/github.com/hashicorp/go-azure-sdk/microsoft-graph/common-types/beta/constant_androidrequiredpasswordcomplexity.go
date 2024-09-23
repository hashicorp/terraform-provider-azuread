package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidRequiredPasswordComplexity string

const (
	AndroidRequiredPasswordComplexity_High   AndroidRequiredPasswordComplexity = "high"
	AndroidRequiredPasswordComplexity_Low    AndroidRequiredPasswordComplexity = "low"
	AndroidRequiredPasswordComplexity_Medium AndroidRequiredPasswordComplexity = "medium"
	AndroidRequiredPasswordComplexity_None   AndroidRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidRequiredPasswordComplexity_High),
		string(AndroidRequiredPasswordComplexity_Low),
		string(AndroidRequiredPasswordComplexity_Medium),
		string(AndroidRequiredPasswordComplexity_None),
	}
}

func (s *AndroidRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidRequiredPasswordComplexity(input string) (*AndroidRequiredPasswordComplexity, error) {
	vals := map[string]AndroidRequiredPasswordComplexity{
		"high":   AndroidRequiredPasswordComplexity_High,
		"low":    AndroidRequiredPasswordComplexity_Low,
		"medium": AndroidRequiredPasswordComplexity_Medium,
		"none":   AndroidRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidRequiredPasswordComplexity(input)
	return &out, nil
}
