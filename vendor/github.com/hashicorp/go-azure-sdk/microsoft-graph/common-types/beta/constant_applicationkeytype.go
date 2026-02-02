package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationKeyType string

const (
	ApplicationKeyType_Certificate  ApplicationKeyType = "certificate"
	ApplicationKeyType_ClientSecret ApplicationKeyType = "clientSecret"
)

func PossibleValuesForApplicationKeyType() []string {
	return []string{
		string(ApplicationKeyType_Certificate),
		string(ApplicationKeyType_ClientSecret),
	}
}

func (s *ApplicationKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationKeyType(input string) (*ApplicationKeyType, error) {
	vals := map[string]ApplicationKeyType{
		"certificate":  ApplicationKeyType_Certificate,
		"clientsecret": ApplicationKeyType_ClientSecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationKeyType(input)
	return &out, nil
}
