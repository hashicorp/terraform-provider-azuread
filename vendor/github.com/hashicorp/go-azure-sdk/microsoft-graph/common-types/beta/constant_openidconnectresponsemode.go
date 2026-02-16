package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectResponseMode string

const (
	OpenIdConnectResponseMode_Formpost OpenIdConnectResponseMode = "form_post"
	OpenIdConnectResponseMode_Query    OpenIdConnectResponseMode = "query"
)

func PossibleValuesForOpenIdConnectResponseMode() []string {
	return []string{
		string(OpenIdConnectResponseMode_Formpost),
		string(OpenIdConnectResponseMode_Query),
	}
}

func (s *OpenIdConnectResponseMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenIdConnectResponseMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenIdConnectResponseMode(input string) (*OpenIdConnectResponseMode, error) {
	vals := map[string]OpenIdConnectResponseMode{
		"form_post": OpenIdConnectResponseMode_Formpost,
		"query":     OpenIdConnectResponseMode_Query,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectResponseMode(input)
	return &out, nil
}
