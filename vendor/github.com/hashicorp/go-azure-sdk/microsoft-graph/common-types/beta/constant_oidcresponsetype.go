package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OidcResponseType string

const (
	OidcResponseType_Code    OidcResponseType = "code"
	OidcResponseType_Idtoken OidcResponseType = "id_token"
	OidcResponseType_Token   OidcResponseType = "token"
)

func PossibleValuesForOidcResponseType() []string {
	return []string{
		string(OidcResponseType_Code),
		string(OidcResponseType_Idtoken),
		string(OidcResponseType_Token),
	}
}

func (s *OidcResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOidcResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOidcResponseType(input string) (*OidcResponseType, error) {
	vals := map[string]OidcResponseType{
		"code":     OidcResponseType_Code,
		"id_token": OidcResponseType_Idtoken,
		"token":    OidcResponseType_Token,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OidcResponseType(input)
	return &out, nil
}
