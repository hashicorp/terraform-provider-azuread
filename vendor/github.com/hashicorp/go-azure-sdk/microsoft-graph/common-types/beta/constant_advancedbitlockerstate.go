package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedBitLockerState string

const (
	AdvancedBitLockerState_FixedDriveEncryptionMethodMismatch      AdvancedBitLockerState = "fixedDriveEncryptionMethodMismatch"
	AdvancedBitLockerState_FixedDriveNotEncrypted                  AdvancedBitLockerState = "fixedDriveNotEncrypted"
	AdvancedBitLockerState_LoggedOnUserNonAdmin                    AdvancedBitLockerState = "loggedOnUserNonAdmin"
	AdvancedBitLockerState_NetworkError                            AdvancedBitLockerState = "networkError"
	AdvancedBitLockerState_NoUserConsent                           AdvancedBitLockerState = "noUserConsent"
	AdvancedBitLockerState_OsVolumeEncryptionMethodMismatch        AdvancedBitLockerState = "osVolumeEncryptionMethodMismatch"
	AdvancedBitLockerState_OsVolumeTpmOnlyRequired                 AdvancedBitLockerState = "osVolumeTpmOnlyRequired"
	AdvancedBitLockerState_OsVolumeTpmPinRequired                  AdvancedBitLockerState = "osVolumeTpmPinRequired"
	AdvancedBitLockerState_OsVolumeTpmPinStartupKeyRequired        AdvancedBitLockerState = "osVolumeTpmPinStartupKeyRequired"
	AdvancedBitLockerState_OsVolumeTpmRequired                     AdvancedBitLockerState = "osVolumeTpmRequired"
	AdvancedBitLockerState_OsVolumeTpmStartupKeyRequired           AdvancedBitLockerState = "osVolumeTpmStartupKeyRequired"
	AdvancedBitLockerState_OsVolumeUnprotected                     AdvancedBitLockerState = "osVolumeUnprotected"
	AdvancedBitLockerState_RecoveryKeyBackupFailed                 AdvancedBitLockerState = "recoveryKeyBackupFailed"
	AdvancedBitLockerState_Success                                 AdvancedBitLockerState = "success"
	AdvancedBitLockerState_TpmNotAvailable                         AdvancedBitLockerState = "tpmNotAvailable"
	AdvancedBitLockerState_TpmNotReady                             AdvancedBitLockerState = "tpmNotReady"
	AdvancedBitLockerState_WindowsRecoveryEnvironmentNotConfigured AdvancedBitLockerState = "windowsRecoveryEnvironmentNotConfigured"
)

func PossibleValuesForAdvancedBitLockerState() []string {
	return []string{
		string(AdvancedBitLockerState_FixedDriveEncryptionMethodMismatch),
		string(AdvancedBitLockerState_FixedDriveNotEncrypted),
		string(AdvancedBitLockerState_LoggedOnUserNonAdmin),
		string(AdvancedBitLockerState_NetworkError),
		string(AdvancedBitLockerState_NoUserConsent),
		string(AdvancedBitLockerState_OsVolumeEncryptionMethodMismatch),
		string(AdvancedBitLockerState_OsVolumeTpmOnlyRequired),
		string(AdvancedBitLockerState_OsVolumeTpmPinRequired),
		string(AdvancedBitLockerState_OsVolumeTpmPinStartupKeyRequired),
		string(AdvancedBitLockerState_OsVolumeTpmRequired),
		string(AdvancedBitLockerState_OsVolumeTpmStartupKeyRequired),
		string(AdvancedBitLockerState_OsVolumeUnprotected),
		string(AdvancedBitLockerState_RecoveryKeyBackupFailed),
		string(AdvancedBitLockerState_Success),
		string(AdvancedBitLockerState_TpmNotAvailable),
		string(AdvancedBitLockerState_TpmNotReady),
		string(AdvancedBitLockerState_WindowsRecoveryEnvironmentNotConfigured),
	}
}

func (s *AdvancedBitLockerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvancedBitLockerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvancedBitLockerState(input string) (*AdvancedBitLockerState, error) {
	vals := map[string]AdvancedBitLockerState{
		"fixeddriveencryptionmethodmismatch":      AdvancedBitLockerState_FixedDriveEncryptionMethodMismatch,
		"fixeddrivenotencrypted":                  AdvancedBitLockerState_FixedDriveNotEncrypted,
		"loggedonusernonadmin":                    AdvancedBitLockerState_LoggedOnUserNonAdmin,
		"networkerror":                            AdvancedBitLockerState_NetworkError,
		"nouserconsent":                           AdvancedBitLockerState_NoUserConsent,
		"osvolumeencryptionmethodmismatch":        AdvancedBitLockerState_OsVolumeEncryptionMethodMismatch,
		"osvolumetpmonlyrequired":                 AdvancedBitLockerState_OsVolumeTpmOnlyRequired,
		"osvolumetpmpinrequired":                  AdvancedBitLockerState_OsVolumeTpmPinRequired,
		"osvolumetpmpinstartupkeyrequired":        AdvancedBitLockerState_OsVolumeTpmPinStartupKeyRequired,
		"osvolumetpmrequired":                     AdvancedBitLockerState_OsVolumeTpmRequired,
		"osvolumetpmstartupkeyrequired":           AdvancedBitLockerState_OsVolumeTpmStartupKeyRequired,
		"osvolumeunprotected":                     AdvancedBitLockerState_OsVolumeUnprotected,
		"recoverykeybackupfailed":                 AdvancedBitLockerState_RecoveryKeyBackupFailed,
		"success":                                 AdvancedBitLockerState_Success,
		"tpmnotavailable":                         AdvancedBitLockerState_TpmNotAvailable,
		"tpmnotready":                             AdvancedBitLockerState_TpmNotReady,
		"windowsrecoveryenvironmentnotconfigured": AdvancedBitLockerState_WindowsRecoveryEnvironmentNotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedBitLockerState(input)
	return &out, nil
}
