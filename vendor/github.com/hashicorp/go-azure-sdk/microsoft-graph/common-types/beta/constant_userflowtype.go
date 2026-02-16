package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFlowType string

const (
	UserFlowType_PasswordReset  UserFlowType = "passwordReset"
	UserFlowType_ProfileUpdate  UserFlowType = "profileUpdate"
	UserFlowType_ResourceOwner  UserFlowType = "resourceOwner"
	UserFlowType_SignIn         UserFlowType = "signIn"
	UserFlowType_SignUp         UserFlowType = "signUp"
	UserFlowType_SignUpOrSignIn UserFlowType = "signUpOrSignIn"
)

func PossibleValuesForUserFlowType() []string {
	return []string{
		string(UserFlowType_PasswordReset),
		string(UserFlowType_ProfileUpdate),
		string(UserFlowType_ResourceOwner),
		string(UserFlowType_SignIn),
		string(UserFlowType_SignUp),
		string(UserFlowType_SignUpOrSignIn),
	}
}

func (s *UserFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserFlowType(input string) (*UserFlowType, error) {
	vals := map[string]UserFlowType{
		"passwordreset":  UserFlowType_PasswordReset,
		"profileupdate":  UserFlowType_ProfileUpdate,
		"resourceowner":  UserFlowType_ResourceOwner,
		"signin":         UserFlowType_SignIn,
		"signup":         UserFlowType_SignUp,
		"signuporsignin": UserFlowType_SignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserFlowType(input)
	return &out, nil
}
