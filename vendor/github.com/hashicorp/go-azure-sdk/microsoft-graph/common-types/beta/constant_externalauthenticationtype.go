package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalAuthenticationType string

const (
	ExternalAuthenticationType_AadPreAuthentication ExternalAuthenticationType = "aadPreAuthentication"
	ExternalAuthenticationType_Passthru             ExternalAuthenticationType = "passthru"
)

func PossibleValuesForExternalAuthenticationType() []string {
	return []string{
		string(ExternalAuthenticationType_AadPreAuthentication),
		string(ExternalAuthenticationType_Passthru),
	}
}

func (s *ExternalAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalAuthenticationType(input string) (*ExternalAuthenticationType, error) {
	vals := map[string]ExternalAuthenticationType{
		"aadpreauthentication": ExternalAuthenticationType_AadPreAuthentication,
		"passthru":             ExternalAuthenticationType_Passthru,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalAuthenticationType(input)
	return &out, nil
}
