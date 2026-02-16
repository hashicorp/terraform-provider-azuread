package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsRequestOccurrenceStatus string

const (
	PermissionsRequestOccurrenceStatus_Granted        PermissionsRequestOccurrenceStatus = "granted"
	PermissionsRequestOccurrenceStatus_Granting       PermissionsRequestOccurrenceStatus = "granting"
	PermissionsRequestOccurrenceStatus_GrantingFailed PermissionsRequestOccurrenceStatus = "grantingFailed"
	PermissionsRequestOccurrenceStatus_Revoked        PermissionsRequestOccurrenceStatus = "revoked"
	PermissionsRequestOccurrenceStatus_Revoking       PermissionsRequestOccurrenceStatus = "revoking"
	PermissionsRequestOccurrenceStatus_RevokingFailed PermissionsRequestOccurrenceStatus = "revokingFailed"
)

func PossibleValuesForPermissionsRequestOccurrenceStatus() []string {
	return []string{
		string(PermissionsRequestOccurrenceStatus_Granted),
		string(PermissionsRequestOccurrenceStatus_Granting),
		string(PermissionsRequestOccurrenceStatus_GrantingFailed),
		string(PermissionsRequestOccurrenceStatus_Revoked),
		string(PermissionsRequestOccurrenceStatus_Revoking),
		string(PermissionsRequestOccurrenceStatus_RevokingFailed),
	}
}

func (s *PermissionsRequestOccurrenceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsRequestOccurrenceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsRequestOccurrenceStatus(input string) (*PermissionsRequestOccurrenceStatus, error) {
	vals := map[string]PermissionsRequestOccurrenceStatus{
		"granted":        PermissionsRequestOccurrenceStatus_Granted,
		"granting":       PermissionsRequestOccurrenceStatus_Granting,
		"grantingfailed": PermissionsRequestOccurrenceStatus_GrantingFailed,
		"revoked":        PermissionsRequestOccurrenceStatus_Revoked,
		"revoking":       PermissionsRequestOccurrenceStatus_Revoking,
		"revokingfailed": PermissionsRequestOccurrenceStatus_RevokingFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsRequestOccurrenceStatus(input)
	return &out, nil
}
