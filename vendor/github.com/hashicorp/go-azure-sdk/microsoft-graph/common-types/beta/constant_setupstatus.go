package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetupStatus string

const (
	SetupStatus_Disabled                      SetupStatus = "disabled"
	SetupStatus_NotRegisteredYet              SetupStatus = "notRegisteredYet"
	SetupStatus_RegisteredSetupInProgress     SetupStatus = "registeredSetupInProgress"
	SetupStatus_RegisteredSetupNotStarted     SetupStatus = "registeredSetupNotStarted"
	SetupStatus_RegistrationAndSetupCompleted SetupStatus = "registrationAndSetupCompleted"
	SetupStatus_RegistrationFailed            SetupStatus = "registrationFailed"
	SetupStatus_RegistrationTimedOut          SetupStatus = "registrationTimedOut"
	SetupStatus_Unknown                       SetupStatus = "unknown"
)

func PossibleValuesForSetupStatus() []string {
	return []string{
		string(SetupStatus_Disabled),
		string(SetupStatus_NotRegisteredYet),
		string(SetupStatus_RegisteredSetupInProgress),
		string(SetupStatus_RegisteredSetupNotStarted),
		string(SetupStatus_RegistrationAndSetupCompleted),
		string(SetupStatus_RegistrationFailed),
		string(SetupStatus_RegistrationTimedOut),
		string(SetupStatus_Unknown),
	}
}

func (s *SetupStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSetupStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSetupStatus(input string) (*SetupStatus, error) {
	vals := map[string]SetupStatus{
		"disabled":                      SetupStatus_Disabled,
		"notregisteredyet":              SetupStatus_NotRegisteredYet,
		"registeredsetupinprogress":     SetupStatus_RegisteredSetupInProgress,
		"registeredsetupnotstarted":     SetupStatus_RegisteredSetupNotStarted,
		"registrationandsetupcompleted": SetupStatus_RegistrationAndSetupCompleted,
		"registrationfailed":            SetupStatus_RegistrationFailed,
		"registrationtimedout":          SetupStatus_RegistrationTimedOut,
		"unknown":                       SetupStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SetupStatus(input)
	return &out, nil
}
