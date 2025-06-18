package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationFailureReasonCode string

const (
	AuthenticationFailureReasonCode_BadRequest    AuthenticationFailureReasonCode = "badRequest"
	AuthenticationFailureReasonCode_ConfigError   AuthenticationFailureReasonCode = "configError"
	AuthenticationFailureReasonCode_Denied        AuthenticationFailureReasonCode = "denied"
	AuthenticationFailureReasonCode_Incomplete    AuthenticationFailureReasonCode = "incomplete"
	AuthenticationFailureReasonCode_Other         AuthenticationFailureReasonCode = "other"
	AuthenticationFailureReasonCode_SystemFailure AuthenticationFailureReasonCode = "systemFailure"
	AuthenticationFailureReasonCode_UserError     AuthenticationFailureReasonCode = "userError"
)

func PossibleValuesForAuthenticationFailureReasonCode() []string {
	return []string{
		string(AuthenticationFailureReasonCode_BadRequest),
		string(AuthenticationFailureReasonCode_ConfigError),
		string(AuthenticationFailureReasonCode_Denied),
		string(AuthenticationFailureReasonCode_Incomplete),
		string(AuthenticationFailureReasonCode_Other),
		string(AuthenticationFailureReasonCode_SystemFailure),
		string(AuthenticationFailureReasonCode_UserError),
	}
}

func (s *AuthenticationFailureReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationFailureReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationFailureReasonCode(input string) (*AuthenticationFailureReasonCode, error) {
	vals := map[string]AuthenticationFailureReasonCode{
		"badrequest":    AuthenticationFailureReasonCode_BadRequest,
		"configerror":   AuthenticationFailureReasonCode_ConfigError,
		"denied":        AuthenticationFailureReasonCode_Denied,
		"incomplete":    AuthenticationFailureReasonCode_Incomplete,
		"other":         AuthenticationFailureReasonCode_Other,
		"systemfailure": AuthenticationFailureReasonCode_SystemFailure,
		"usererror":     AuthenticationFailureReasonCode_UserError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationFailureReasonCode(input)
	return &out, nil
}
