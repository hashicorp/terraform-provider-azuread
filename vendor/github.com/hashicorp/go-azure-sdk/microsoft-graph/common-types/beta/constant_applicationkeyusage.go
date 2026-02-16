package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationKeyUsage string

const (
	ApplicationKeyUsage_Sign   ApplicationKeyUsage = "sign"
	ApplicationKeyUsage_Verify ApplicationKeyUsage = "verify"
)

func PossibleValuesForApplicationKeyUsage() []string {
	return []string{
		string(ApplicationKeyUsage_Sign),
		string(ApplicationKeyUsage_Verify),
	}
}

func (s *ApplicationKeyUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationKeyUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationKeyUsage(input string) (*ApplicationKeyUsage, error) {
	vals := map[string]ApplicationKeyUsage{
		"sign":   ApplicationKeyUsage_Sign,
		"verify": ApplicationKeyUsage_Verify,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationKeyUsage(input)
	return &out, nil
}
