package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventType string

const (
	AuthenticationEventType_PageRenderStart    AuthenticationEventType = "pageRenderStart"
	AuthenticationEventType_TokenIssuanceStart AuthenticationEventType = "tokenIssuanceStart"
)

func PossibleValuesForAuthenticationEventType() []string {
	return []string{
		string(AuthenticationEventType_PageRenderStart),
		string(AuthenticationEventType_TokenIssuanceStart),
	}
}

func (s *AuthenticationEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationEventType(input string) (*AuthenticationEventType, error) {
	vals := map[string]AuthenticationEventType{
		"pagerenderstart":    AuthenticationEventType_PageRenderStart,
		"tokenissuancestart": AuthenticationEventType_TokenIssuanceStart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationEventType(input)
	return &out, nil
}
