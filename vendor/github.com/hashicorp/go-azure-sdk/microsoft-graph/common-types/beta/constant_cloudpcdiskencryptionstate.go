package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDiskEncryptionState string

const (
	CloudPCDiskEncryptionState_EncryptedUsingCustomerManagedKey CloudPCDiskEncryptionState = "encryptedUsingCustomerManagedKey"
	CloudPCDiskEncryptionState_EncryptedUsingPlatformManagedKey CloudPCDiskEncryptionState = "encryptedUsingPlatformManagedKey"
	CloudPCDiskEncryptionState_NotAvailable                     CloudPCDiskEncryptionState = "notAvailable"
	CloudPCDiskEncryptionState_NotEncrypted                     CloudPCDiskEncryptionState = "notEncrypted"
)

func PossibleValuesForCloudPCDiskEncryptionState() []string {
	return []string{
		string(CloudPCDiskEncryptionState_EncryptedUsingCustomerManagedKey),
		string(CloudPCDiskEncryptionState_EncryptedUsingPlatformManagedKey),
		string(CloudPCDiskEncryptionState_NotAvailable),
		string(CloudPCDiskEncryptionState_NotEncrypted),
	}
}

func (s *CloudPCDiskEncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDiskEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDiskEncryptionState(input string) (*CloudPCDiskEncryptionState, error) {
	vals := map[string]CloudPCDiskEncryptionState{
		"encryptedusingcustomermanagedkey": CloudPCDiskEncryptionState_EncryptedUsingCustomerManagedKey,
		"encryptedusingplatformmanagedkey": CloudPCDiskEncryptionState_EncryptedUsingPlatformManagedKey,
		"notavailable":                     CloudPCDiskEncryptionState_NotAvailable,
		"notencrypted":                     CloudPCDiskEncryptionState_NotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDiskEncryptionState(input)
	return &out, nil
}
