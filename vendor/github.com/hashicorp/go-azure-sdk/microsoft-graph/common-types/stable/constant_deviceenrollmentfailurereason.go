package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentFailureReason string

const (
	DeviceEnrollmentFailureReason_AccountValidation              DeviceEnrollmentFailureReason = "accountValidation"
	DeviceEnrollmentFailureReason_Authentication                 DeviceEnrollmentFailureReason = "authentication"
	DeviceEnrollmentFailureReason_Authorization                  DeviceEnrollmentFailureReason = "authorization"
	DeviceEnrollmentFailureReason_BadRequest                     DeviceEnrollmentFailureReason = "badRequest"
	DeviceEnrollmentFailureReason_ClientDisconnected             DeviceEnrollmentFailureReason = "clientDisconnected"
	DeviceEnrollmentFailureReason_DeviceNotSupported             DeviceEnrollmentFailureReason = "deviceNotSupported"
	DeviceEnrollmentFailureReason_EnrollmentRestrictionsEnforced DeviceEnrollmentFailureReason = "enrollmentRestrictionsEnforced"
	DeviceEnrollmentFailureReason_FeatureNotSupported            DeviceEnrollmentFailureReason = "featureNotSupported"
	DeviceEnrollmentFailureReason_InMaintenance                  DeviceEnrollmentFailureReason = "inMaintenance"
	DeviceEnrollmentFailureReason_Unknown                        DeviceEnrollmentFailureReason = "unknown"
	DeviceEnrollmentFailureReason_UserAbandonment                DeviceEnrollmentFailureReason = "userAbandonment"
	DeviceEnrollmentFailureReason_UserValidation                 DeviceEnrollmentFailureReason = "userValidation"
)

func PossibleValuesForDeviceEnrollmentFailureReason() []string {
	return []string{
		string(DeviceEnrollmentFailureReason_AccountValidation),
		string(DeviceEnrollmentFailureReason_Authentication),
		string(DeviceEnrollmentFailureReason_Authorization),
		string(DeviceEnrollmentFailureReason_BadRequest),
		string(DeviceEnrollmentFailureReason_ClientDisconnected),
		string(DeviceEnrollmentFailureReason_DeviceNotSupported),
		string(DeviceEnrollmentFailureReason_EnrollmentRestrictionsEnforced),
		string(DeviceEnrollmentFailureReason_FeatureNotSupported),
		string(DeviceEnrollmentFailureReason_InMaintenance),
		string(DeviceEnrollmentFailureReason_Unknown),
		string(DeviceEnrollmentFailureReason_UserAbandonment),
		string(DeviceEnrollmentFailureReason_UserValidation),
	}
}

func (s *DeviceEnrollmentFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentFailureReason(input string) (*DeviceEnrollmentFailureReason, error) {
	vals := map[string]DeviceEnrollmentFailureReason{
		"accountvalidation":              DeviceEnrollmentFailureReason_AccountValidation,
		"authentication":                 DeviceEnrollmentFailureReason_Authentication,
		"authorization":                  DeviceEnrollmentFailureReason_Authorization,
		"badrequest":                     DeviceEnrollmentFailureReason_BadRequest,
		"clientdisconnected":             DeviceEnrollmentFailureReason_ClientDisconnected,
		"devicenotsupported":             DeviceEnrollmentFailureReason_DeviceNotSupported,
		"enrollmentrestrictionsenforced": DeviceEnrollmentFailureReason_EnrollmentRestrictionsEnforced,
		"featurenotsupported":            DeviceEnrollmentFailureReason_FeatureNotSupported,
		"inmaintenance":                  DeviceEnrollmentFailureReason_InMaintenance,
		"unknown":                        DeviceEnrollmentFailureReason_Unknown,
		"userabandonment":                DeviceEnrollmentFailureReason_UserAbandonment,
		"uservalidation":                 DeviceEnrollmentFailureReason_UserValidation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentFailureReason(input)
	return &out, nil
}
