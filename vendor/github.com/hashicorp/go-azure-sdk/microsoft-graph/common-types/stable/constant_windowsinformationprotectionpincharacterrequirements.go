package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPinCharacterRequirements string

const (
	WindowsInformationProtectionPinCharacterRequirements_Allow             WindowsInformationProtectionPinCharacterRequirements = "allow"
	WindowsInformationProtectionPinCharacterRequirements_NotAllow          WindowsInformationProtectionPinCharacterRequirements = "notAllow"
	WindowsInformationProtectionPinCharacterRequirements_RequireAtLeastOne WindowsInformationProtectionPinCharacterRequirements = "requireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPinCharacterRequirements() []string {
	return []string{
		string(WindowsInformationProtectionPinCharacterRequirements_Allow),
		string(WindowsInformationProtectionPinCharacterRequirements_NotAllow),
		string(WindowsInformationProtectionPinCharacterRequirements_RequireAtLeastOne),
	}
}

func (s *WindowsInformationProtectionPinCharacterRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionPinCharacterRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionPinCharacterRequirements(input string) (*WindowsInformationProtectionPinCharacterRequirements, error) {
	vals := map[string]WindowsInformationProtectionPinCharacterRequirements{
		"allow":             WindowsInformationProtectionPinCharacterRequirements_Allow,
		"notallow":          WindowsInformationProtectionPinCharacterRequirements_NotAllow,
		"requireatleastone": WindowsInformationProtectionPinCharacterRequirements_RequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPinCharacterRequirements(input)
	return &out, nil
}
