package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalAudienceScope string

const (
	ExternalAudienceScope_All          ExternalAudienceScope = "all"
	ExternalAudienceScope_ContactsOnly ExternalAudienceScope = "contactsOnly"
	ExternalAudienceScope_None         ExternalAudienceScope = "none"
)

func PossibleValuesForExternalAudienceScope() []string {
	return []string{
		string(ExternalAudienceScope_All),
		string(ExternalAudienceScope_ContactsOnly),
		string(ExternalAudienceScope_None),
	}
}

func (s *ExternalAudienceScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalAudienceScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalAudienceScope(input string) (*ExternalAudienceScope, error) {
	vals := map[string]ExternalAudienceScope{
		"all":          ExternalAudienceScope_All,
		"contactsonly": ExternalAudienceScope_ContactsOnly,
		"none":         ExternalAudienceScope_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalAudienceScope(input)
	return &out, nil
}
