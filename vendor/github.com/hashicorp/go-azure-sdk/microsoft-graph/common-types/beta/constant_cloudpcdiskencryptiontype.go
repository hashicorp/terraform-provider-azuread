package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDiskEncryptionType string

const (
	CloudPCDiskEncryptionType_CustomerManagedKey CloudPCDiskEncryptionType = "customerManagedKey"
	CloudPCDiskEncryptionType_PlatformManagedKey CloudPCDiskEncryptionType = "platformManagedKey"
)

func PossibleValuesForCloudPCDiskEncryptionType() []string {
	return []string{
		string(CloudPCDiskEncryptionType_CustomerManagedKey),
		string(CloudPCDiskEncryptionType_PlatformManagedKey),
	}
}

func (s *CloudPCDiskEncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDiskEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDiskEncryptionType(input string) (*CloudPCDiskEncryptionType, error) {
	vals := map[string]CloudPCDiskEncryptionType{
		"customermanagedkey": CloudPCDiskEncryptionType_CustomerManagedKey,
		"platformmanagedkey": CloudPCDiskEncryptionType_PlatformManagedKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDiskEncryptionType(input)
	return &out, nil
}
