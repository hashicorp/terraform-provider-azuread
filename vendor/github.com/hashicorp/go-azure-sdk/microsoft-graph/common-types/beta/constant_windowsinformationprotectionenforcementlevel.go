package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionEnforcementLevel string

const (
	WindowsInformationProtectionEnforcementLevel_EncryptAndAuditOnly   WindowsInformationProtectionEnforcementLevel = "encryptAndAuditOnly"
	WindowsInformationProtectionEnforcementLevel_EncryptAuditAndBlock  WindowsInformationProtectionEnforcementLevel = "encryptAuditAndBlock"
	WindowsInformationProtectionEnforcementLevel_EncryptAuditAndPrompt WindowsInformationProtectionEnforcementLevel = "encryptAuditAndPrompt"
	WindowsInformationProtectionEnforcementLevel_NoProtection          WindowsInformationProtectionEnforcementLevel = "noProtection"
)

func PossibleValuesForWindowsInformationProtectionEnforcementLevel() []string {
	return []string{
		string(WindowsInformationProtectionEnforcementLevel_EncryptAndAuditOnly),
		string(WindowsInformationProtectionEnforcementLevel_EncryptAuditAndBlock),
		string(WindowsInformationProtectionEnforcementLevel_EncryptAuditAndPrompt),
		string(WindowsInformationProtectionEnforcementLevel_NoProtection),
	}
}

func (s *WindowsInformationProtectionEnforcementLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionEnforcementLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionEnforcementLevel(input string) (*WindowsInformationProtectionEnforcementLevel, error) {
	vals := map[string]WindowsInformationProtectionEnforcementLevel{
		"encryptandauditonly":   WindowsInformationProtectionEnforcementLevel_EncryptAndAuditOnly,
		"encryptauditandblock":  WindowsInformationProtectionEnforcementLevel_EncryptAuditAndBlock,
		"encryptauditandprompt": WindowsInformationProtectionEnforcementLevel_EncryptAuditAndPrompt,
		"noprotection":          WindowsInformationProtectionEnforcementLevel_NoProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionEnforcementLevel(input)
	return &out, nil
}
