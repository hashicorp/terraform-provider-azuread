package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyStatus string

const (
	AuthenticationAppPolicyStatus_AppContextNotShown                AuthenticationAppPolicyStatus = "appContextNotShown"
	AuthenticationAppPolicyStatus_AppContextOutOfDate               AuthenticationAppPolicyStatus = "appContextOutOfDate"
	AuthenticationAppPolicyStatus_AppContextShown                   AuthenticationAppPolicyStatus = "appContextShown"
	AuthenticationAppPolicyStatus_AppLockDisabled                   AuthenticationAppPolicyStatus = "appLockDisabled"
	AuthenticationAppPolicyStatus_AppLockEnabled                    AuthenticationAppPolicyStatus = "appLockEnabled"
	AuthenticationAppPolicyStatus_AppLockOutOfDate                  AuthenticationAppPolicyStatus = "appLockOutOfDate"
	AuthenticationAppPolicyStatus_LocationContextNotShown           AuthenticationAppPolicyStatus = "locationContextNotShown"
	AuthenticationAppPolicyStatus_LocationContextOutOfDate          AuthenticationAppPolicyStatus = "locationContextOutOfDate"
	AuthenticationAppPolicyStatus_LocationContextShown              AuthenticationAppPolicyStatus = "locationContextShown"
	AuthenticationAppPolicyStatus_NumberMatchCorrectNumberEntered   AuthenticationAppPolicyStatus = "numberMatchCorrectNumberEntered"
	AuthenticationAppPolicyStatus_NumberMatchDeny                   AuthenticationAppPolicyStatus = "numberMatchDeny"
	AuthenticationAppPolicyStatus_NumberMatchIncorrectNumberEntered AuthenticationAppPolicyStatus = "numberMatchIncorrectNumberEntered"
	AuthenticationAppPolicyStatus_NumberMatchOutOfDate              AuthenticationAppPolicyStatus = "numberMatchOutOfDate"
	AuthenticationAppPolicyStatus_TamperResistantHardwareNotUsed    AuthenticationAppPolicyStatus = "tamperResistantHardwareNotUsed"
	AuthenticationAppPolicyStatus_TamperResistantHardwareOutOfDate  AuthenticationAppPolicyStatus = "tamperResistantHardwareOutOfDate"
	AuthenticationAppPolicyStatus_TamperResistantHardwareUsed       AuthenticationAppPolicyStatus = "tamperResistantHardwareUsed"
	AuthenticationAppPolicyStatus_Unknown                           AuthenticationAppPolicyStatus = "unknown"
)

func PossibleValuesForAuthenticationAppPolicyStatus() []string {
	return []string{
		string(AuthenticationAppPolicyStatus_AppContextNotShown),
		string(AuthenticationAppPolicyStatus_AppContextOutOfDate),
		string(AuthenticationAppPolicyStatus_AppContextShown),
		string(AuthenticationAppPolicyStatus_AppLockDisabled),
		string(AuthenticationAppPolicyStatus_AppLockEnabled),
		string(AuthenticationAppPolicyStatus_AppLockOutOfDate),
		string(AuthenticationAppPolicyStatus_LocationContextNotShown),
		string(AuthenticationAppPolicyStatus_LocationContextOutOfDate),
		string(AuthenticationAppPolicyStatus_LocationContextShown),
		string(AuthenticationAppPolicyStatus_NumberMatchCorrectNumberEntered),
		string(AuthenticationAppPolicyStatus_NumberMatchDeny),
		string(AuthenticationAppPolicyStatus_NumberMatchIncorrectNumberEntered),
		string(AuthenticationAppPolicyStatus_NumberMatchOutOfDate),
		string(AuthenticationAppPolicyStatus_TamperResistantHardwareNotUsed),
		string(AuthenticationAppPolicyStatus_TamperResistantHardwareOutOfDate),
		string(AuthenticationAppPolicyStatus_TamperResistantHardwareUsed),
		string(AuthenticationAppPolicyStatus_Unknown),
	}
}

func (s *AuthenticationAppPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppPolicyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppPolicyStatus(input string) (*AuthenticationAppPolicyStatus, error) {
	vals := map[string]AuthenticationAppPolicyStatus{
		"appcontextnotshown":                AuthenticationAppPolicyStatus_AppContextNotShown,
		"appcontextoutofdate":               AuthenticationAppPolicyStatus_AppContextOutOfDate,
		"appcontextshown":                   AuthenticationAppPolicyStatus_AppContextShown,
		"applockdisabled":                   AuthenticationAppPolicyStatus_AppLockDisabled,
		"applockenabled":                    AuthenticationAppPolicyStatus_AppLockEnabled,
		"applockoutofdate":                  AuthenticationAppPolicyStatus_AppLockOutOfDate,
		"locationcontextnotshown":           AuthenticationAppPolicyStatus_LocationContextNotShown,
		"locationcontextoutofdate":          AuthenticationAppPolicyStatus_LocationContextOutOfDate,
		"locationcontextshown":              AuthenticationAppPolicyStatus_LocationContextShown,
		"numbermatchcorrectnumberentered":   AuthenticationAppPolicyStatus_NumberMatchCorrectNumberEntered,
		"numbermatchdeny":                   AuthenticationAppPolicyStatus_NumberMatchDeny,
		"numbermatchincorrectnumberentered": AuthenticationAppPolicyStatus_NumberMatchIncorrectNumberEntered,
		"numbermatchoutofdate":              AuthenticationAppPolicyStatus_NumberMatchOutOfDate,
		"tamperresistanthardwarenotused":    AuthenticationAppPolicyStatus_TamperResistantHardwareNotUsed,
		"tamperresistanthardwareoutofdate":  AuthenticationAppPolicyStatus_TamperResistantHardwareOutOfDate,
		"tamperresistanthardwareused":       AuthenticationAppPolicyStatus_TamperResistantHardwareUsed,
		"unknown":                           AuthenticationAppPolicyStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyStatus(input)
	return &out, nil
}
