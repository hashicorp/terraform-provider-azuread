package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenFormat string

const (
	TokenFormat_Jwt  TokenFormat = "jwt"
	TokenFormat_Saml TokenFormat = "saml"
)

func PossibleValuesForTokenFormat() []string {
	return []string{
		string(TokenFormat_Jwt),
		string(TokenFormat_Saml),
	}
}

func (s *TokenFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTokenFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTokenFormat(input string) (*TokenFormat, error) {
	vals := map[string]TokenFormat{
		"jwt":  TokenFormat_Jwt,
		"saml": TokenFormat_Saml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TokenFormat(input)
	return &out, nil
}
