package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSystemAccessMethods string

const (
	ExternalSystemAccessMethods_Direct       ExternalSystemAccessMethods = "direct"
	ExternalSystemAccessMethods_RoleChaining ExternalSystemAccessMethods = "roleChaining"
)

func PossibleValuesForExternalSystemAccessMethods() []string {
	return []string{
		string(ExternalSystemAccessMethods_Direct),
		string(ExternalSystemAccessMethods_RoleChaining),
	}
}

func (s *ExternalSystemAccessMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalSystemAccessMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalSystemAccessMethods(input string) (*ExternalSystemAccessMethods, error) {
	vals := map[string]ExternalSystemAccessMethods{
		"direct":       ExternalSystemAccessMethods_Direct,
		"rolechaining": ExternalSystemAccessMethods_RoleChaining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalSystemAccessMethods(input)
	return &out, nil
}
