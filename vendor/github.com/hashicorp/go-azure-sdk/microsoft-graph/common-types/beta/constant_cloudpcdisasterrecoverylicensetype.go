package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryLicenseType string

const (
	CloudPCDisasterRecoveryLicenseType_None     CloudPCDisasterRecoveryLicenseType = "none"
	CloudPCDisasterRecoveryLicenseType_Plus     CloudPCDisasterRecoveryLicenseType = "plus"
	CloudPCDisasterRecoveryLicenseType_Standard CloudPCDisasterRecoveryLicenseType = "standard"
)

func PossibleValuesForCloudPCDisasterRecoveryLicenseType() []string {
	return []string{
		string(CloudPCDisasterRecoveryLicenseType_None),
		string(CloudPCDisasterRecoveryLicenseType_Plus),
		string(CloudPCDisasterRecoveryLicenseType_Standard),
	}
}

func (s *CloudPCDisasterRecoveryLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDisasterRecoveryLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDisasterRecoveryLicenseType(input string) (*CloudPCDisasterRecoveryLicenseType, error) {
	vals := map[string]CloudPCDisasterRecoveryLicenseType{
		"none":     CloudPCDisasterRecoveryLicenseType_None,
		"plus":     CloudPCDisasterRecoveryLicenseType_Plus,
		"standard": CloudPCDisasterRecoveryLicenseType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDisasterRecoveryLicenseType(input)
	return &out, nil
}
