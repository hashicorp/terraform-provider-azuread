package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenProtectionStatus string

const (
	TokenProtectionStatus_Bound   TokenProtectionStatus = "bound"
	TokenProtectionStatus_None    TokenProtectionStatus = "none"
	TokenProtectionStatus_Unbound TokenProtectionStatus = "unbound"
)

func PossibleValuesForTokenProtectionStatus() []string {
	return []string{
		string(TokenProtectionStatus_Bound),
		string(TokenProtectionStatus_None),
		string(TokenProtectionStatus_Unbound),
	}
}

func (s *TokenProtectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTokenProtectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTokenProtectionStatus(input string) (*TokenProtectionStatus, error) {
	vals := map[string]TokenProtectionStatus{
		"bound":   TokenProtectionStatus_Bound,
		"none":    TokenProtectionStatus_None,
		"unbound": TokenProtectionStatus_Unbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TokenProtectionStatus(input)
	return &out, nil
}
