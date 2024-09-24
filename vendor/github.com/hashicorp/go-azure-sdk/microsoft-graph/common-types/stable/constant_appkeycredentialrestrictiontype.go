package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppKeyCredentialRestrictionType string

const (
	AppKeyCredentialRestrictionType_AsymmetricKeyLifetime AppKeyCredentialRestrictionType = "asymmetricKeyLifetime"
)

func PossibleValuesForAppKeyCredentialRestrictionType() []string {
	return []string{
		string(AppKeyCredentialRestrictionType_AsymmetricKeyLifetime),
	}
}

func (s *AppKeyCredentialRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppKeyCredentialRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppKeyCredentialRestrictionType(input string) (*AppKeyCredentialRestrictionType, error) {
	vals := map[string]AppKeyCredentialRestrictionType{
		"asymmetrickeylifetime": AppKeyCredentialRestrictionType_AsymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppKeyCredentialRestrictionType(input)
	return &out, nil
}
