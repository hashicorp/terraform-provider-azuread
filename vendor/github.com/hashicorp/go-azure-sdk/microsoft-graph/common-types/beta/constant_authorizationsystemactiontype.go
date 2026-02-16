package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemActionType string

const (
	AuthorizationSystemActionType_Delete AuthorizationSystemActionType = "delete"
	AuthorizationSystemActionType_Read   AuthorizationSystemActionType = "read"
)

func PossibleValuesForAuthorizationSystemActionType() []string {
	return []string{
		string(AuthorizationSystemActionType_Delete),
		string(AuthorizationSystemActionType_Read),
	}
}

func (s *AuthorizationSystemActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemActionType(input string) (*AuthorizationSystemActionType, error) {
	vals := map[string]AuthorizationSystemActionType{
		"delete": AuthorizationSystemActionType_Delete,
		"read":   AuthorizationSystemActionType_Read,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemActionType(input)
	return &out, nil
}
