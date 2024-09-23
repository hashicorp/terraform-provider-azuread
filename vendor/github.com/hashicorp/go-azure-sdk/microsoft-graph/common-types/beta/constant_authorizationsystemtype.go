package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemType string

const (
	AuthorizationSystemType_Aws   AuthorizationSystemType = "aws"
	AuthorizationSystemType_Azure AuthorizationSystemType = "azure"
	AuthorizationSystemType_Gcp   AuthorizationSystemType = "gcp"
)

func PossibleValuesForAuthorizationSystemType() []string {
	return []string{
		string(AuthorizationSystemType_Aws),
		string(AuthorizationSystemType_Azure),
		string(AuthorizationSystemType_Gcp),
	}
}

func (s *AuthorizationSystemType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemType(input string) (*AuthorizationSystemType, error) {
	vals := map[string]AuthorizationSystemType{
		"aws":   AuthorizationSystemType_Aws,
		"azure": AuthorizationSystemType_Azure,
		"gcp":   AuthorizationSystemType_Gcp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemType(input)
	return &out, nil
}
