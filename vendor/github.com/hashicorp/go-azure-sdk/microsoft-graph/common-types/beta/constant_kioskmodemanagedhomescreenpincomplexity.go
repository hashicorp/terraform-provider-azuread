package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KioskModeManagedHomeScreenPinComplexity string

const (
	KioskModeManagedHomeScreenPinComplexity_Complex       KioskModeManagedHomeScreenPinComplexity = "complex"
	KioskModeManagedHomeScreenPinComplexity_NotConfigured KioskModeManagedHomeScreenPinComplexity = "notConfigured"
	KioskModeManagedHomeScreenPinComplexity_Simple        KioskModeManagedHomeScreenPinComplexity = "simple"
)

func PossibleValuesForKioskModeManagedHomeScreenPinComplexity() []string {
	return []string{
		string(KioskModeManagedHomeScreenPinComplexity_Complex),
		string(KioskModeManagedHomeScreenPinComplexity_NotConfigured),
		string(KioskModeManagedHomeScreenPinComplexity_Simple),
	}
}

func (s *KioskModeManagedHomeScreenPinComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKioskModeManagedHomeScreenPinComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKioskModeManagedHomeScreenPinComplexity(input string) (*KioskModeManagedHomeScreenPinComplexity, error) {
	vals := map[string]KioskModeManagedHomeScreenPinComplexity{
		"complex":       KioskModeManagedHomeScreenPinComplexity_Complex,
		"notconfigured": KioskModeManagedHomeScreenPinComplexity_NotConfigured,
		"simple":        KioskModeManagedHomeScreenPinComplexity_Simple,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KioskModeManagedHomeScreenPinComplexity(input)
	return &out, nil
}
