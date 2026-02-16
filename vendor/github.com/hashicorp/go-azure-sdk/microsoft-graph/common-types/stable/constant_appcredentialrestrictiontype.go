package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialRestrictionType string

const (
	AppCredentialRestrictionType_CustomPasswordAddition AppCredentialRestrictionType = "customPasswordAddition"
	AppCredentialRestrictionType_PasswordAddition       AppCredentialRestrictionType = "passwordAddition"
	AppCredentialRestrictionType_PasswordLifetime       AppCredentialRestrictionType = "passwordLifetime"
	AppCredentialRestrictionType_SymmetricKeyAddition   AppCredentialRestrictionType = "symmetricKeyAddition"
	AppCredentialRestrictionType_SymmetricKeyLifetime   AppCredentialRestrictionType = "symmetricKeyLifetime"
)

func PossibleValuesForAppCredentialRestrictionType() []string {
	return []string{
		string(AppCredentialRestrictionType_CustomPasswordAddition),
		string(AppCredentialRestrictionType_PasswordAddition),
		string(AppCredentialRestrictionType_PasswordLifetime),
		string(AppCredentialRestrictionType_SymmetricKeyAddition),
		string(AppCredentialRestrictionType_SymmetricKeyLifetime),
	}
}

func (s *AppCredentialRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppCredentialRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppCredentialRestrictionType(input string) (*AppCredentialRestrictionType, error) {
	vals := map[string]AppCredentialRestrictionType{
		"custompasswordaddition": AppCredentialRestrictionType_CustomPasswordAddition,
		"passwordaddition":       AppCredentialRestrictionType_PasswordAddition,
		"passwordlifetime":       AppCredentialRestrictionType_PasswordLifetime,
		"symmetrickeyaddition":   AppCredentialRestrictionType_SymmetricKeyAddition,
		"symmetrickeylifetime":   AppCredentialRestrictionType_SymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppCredentialRestrictionType(input)
	return &out, nil
}
