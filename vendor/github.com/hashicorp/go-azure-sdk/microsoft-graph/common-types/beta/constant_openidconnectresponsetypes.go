package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectResponseTypes string

const (
	OpenIdConnectResponseTypes_Code    OpenIdConnectResponseTypes = "code"
	OpenIdConnectResponseTypes_Idtoken OpenIdConnectResponseTypes = "id_token"
	OpenIdConnectResponseTypes_Token   OpenIdConnectResponseTypes = "token"
)

func PossibleValuesForOpenIdConnectResponseTypes() []string {
	return []string{
		string(OpenIdConnectResponseTypes_Code),
		string(OpenIdConnectResponseTypes_Idtoken),
		string(OpenIdConnectResponseTypes_Token),
	}
}

func (s *OpenIdConnectResponseTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectResponseTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectResponseTypes(input string) (*OpenIdConnectResponseTypes, error) {
	vals := map[string]OpenIdConnectResponseTypes{
		"code":     OpenIdConnectResponseTypes_Code,
		"id_token": OpenIdConnectResponseTypes_Idtoken,
		"token":    OpenIdConnectResponseTypes_Token,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectResponseTypes(input)
	return &out, nil
}
