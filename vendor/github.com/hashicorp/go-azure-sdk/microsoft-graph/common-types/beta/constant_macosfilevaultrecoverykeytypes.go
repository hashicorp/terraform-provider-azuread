package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSFileVaultRecoveryKeyTypes string

const (
	MacOSFileVaultRecoveryKeyTypes_InstitutionalRecoveryKey MacOSFileVaultRecoveryKeyTypes = "institutionalRecoveryKey"
	MacOSFileVaultRecoveryKeyTypes_NotConfigured            MacOSFileVaultRecoveryKeyTypes = "notConfigured"
	MacOSFileVaultRecoveryKeyTypes_PersonalRecoveryKey      MacOSFileVaultRecoveryKeyTypes = "personalRecoveryKey"
)

func PossibleValuesForMacOSFileVaultRecoveryKeyTypes() []string {
	return []string{
		string(MacOSFileVaultRecoveryKeyTypes_InstitutionalRecoveryKey),
		string(MacOSFileVaultRecoveryKeyTypes_NotConfigured),
		string(MacOSFileVaultRecoveryKeyTypes_PersonalRecoveryKey),
	}
}

func (s *MacOSFileVaultRecoveryKeyTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSFileVaultRecoveryKeyTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSFileVaultRecoveryKeyTypes(input string) (*MacOSFileVaultRecoveryKeyTypes, error) {
	vals := map[string]MacOSFileVaultRecoveryKeyTypes{
		"institutionalrecoverykey": MacOSFileVaultRecoveryKeyTypes_InstitutionalRecoveryKey,
		"notconfigured":            MacOSFileVaultRecoveryKeyTypes_NotConfigured,
		"personalrecoverykey":      MacOSFileVaultRecoveryKeyTypes_PersonalRecoveryKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSFileVaultRecoveryKeyTypes(input)
	return &out, nil
}
