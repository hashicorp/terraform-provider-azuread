package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationState string

const (
	DeviceRegistrationState_ApprovalPending                DeviceRegistrationState = "approvalPending"
	DeviceRegistrationState_CertificateReset               DeviceRegistrationState = "certificateReset"
	DeviceRegistrationState_KeyConflict                    DeviceRegistrationState = "keyConflict"
	DeviceRegistrationState_NotRegistered                  DeviceRegistrationState = "notRegistered"
	DeviceRegistrationState_NotRegisteredPendingEnrollment DeviceRegistrationState = "notRegisteredPendingEnrollment"
	DeviceRegistrationState_Registered                     DeviceRegistrationState = "registered"
	DeviceRegistrationState_Revoked                        DeviceRegistrationState = "revoked"
	DeviceRegistrationState_Unknown                        DeviceRegistrationState = "unknown"
)

func PossibleValuesForDeviceRegistrationState() []string {
	return []string{
		string(DeviceRegistrationState_ApprovalPending),
		string(DeviceRegistrationState_CertificateReset),
		string(DeviceRegistrationState_KeyConflict),
		string(DeviceRegistrationState_NotRegistered),
		string(DeviceRegistrationState_NotRegisteredPendingEnrollment),
		string(DeviceRegistrationState_Registered),
		string(DeviceRegistrationState_Revoked),
		string(DeviceRegistrationState_Unknown),
	}
}

func (s *DeviceRegistrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceRegistrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceRegistrationState(input string) (*DeviceRegistrationState, error) {
	vals := map[string]DeviceRegistrationState{
		"approvalpending":                DeviceRegistrationState_ApprovalPending,
		"certificatereset":               DeviceRegistrationState_CertificateReset,
		"keyconflict":                    DeviceRegistrationState_KeyConflict,
		"notregistered":                  DeviceRegistrationState_NotRegistered,
		"notregisteredpendingenrollment": DeviceRegistrationState_NotRegisteredPendingEnrollment,
		"registered":                     DeviceRegistrationState_Registered,
		"revoked":                        DeviceRegistrationState_Revoked,
		"unknown":                        DeviceRegistrationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceRegistrationState(input)
	return &out, nil
}
