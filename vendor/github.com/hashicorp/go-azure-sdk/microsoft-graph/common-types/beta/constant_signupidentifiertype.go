package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignUpIdentifierType string

const (
	SignUpIdentifierType_EmailAddress SignUpIdentifierType = "emailAddress"
)

func PossibleValuesForSignUpIdentifierType() []string {
	return []string{
		string(SignUpIdentifierType_EmailAddress),
	}
}

func (s *SignUpIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignUpIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignUpIdentifierType(input string) (*SignUpIdentifierType, error) {
	vals := map[string]SignUpIdentifierType{
		"emailaddress": SignUpIdentifierType_EmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignUpIdentifierType(input)
	return &out, nil
}
