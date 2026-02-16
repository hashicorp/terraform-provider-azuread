package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileVaultState string

const (
	FileVaultState_DriveEncryptedByUser   FileVaultState = "driveEncryptedByUser"
	FileVaultState_EscrowNotEnabled       FileVaultState = "escrowNotEnabled"
	FileVaultState_Success                FileVaultState = "success"
	FileVaultState_UserDeferredEncryption FileVaultState = "userDeferredEncryption"
)

func PossibleValuesForFileVaultState() []string {
	return []string{
		string(FileVaultState_DriveEncryptedByUser),
		string(FileVaultState_EscrowNotEnabled),
		string(FileVaultState_Success),
		string(FileVaultState_UserDeferredEncryption),
	}
}

func (s *FileVaultState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileVaultState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileVaultState(input string) (*FileVaultState, error) {
	vals := map[string]FileVaultState{
		"driveencryptedbyuser":   FileVaultState_DriveEncryptedByUser,
		"escrownotenabled":       FileVaultState_EscrowNotEnabled,
		"success":                FileVaultState_Success,
		"userdeferredencryption": FileVaultState_UserDeferredEncryption,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileVaultState(input)
	return &out, nil
}
