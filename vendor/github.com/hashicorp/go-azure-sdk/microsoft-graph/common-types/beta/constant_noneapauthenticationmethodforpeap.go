package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NonEapAuthenticationMethodForPeap string

const (
	NonEapAuthenticationMethodForPeap_MicrosoftChapVersionTwo NonEapAuthenticationMethodForPeap = "microsoftChapVersionTwo"
	NonEapAuthenticationMethodForPeap_None                    NonEapAuthenticationMethodForPeap = "none"
)

func PossibleValuesForNonEapAuthenticationMethodForPeap() []string {
	return []string{
		string(NonEapAuthenticationMethodForPeap_MicrosoftChapVersionTwo),
		string(NonEapAuthenticationMethodForPeap_None),
	}
}

func (s *NonEapAuthenticationMethodForPeap) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNonEapAuthenticationMethodForPeap(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNonEapAuthenticationMethodForPeap(input string) (*NonEapAuthenticationMethodForPeap, error) {
	vals := map[string]NonEapAuthenticationMethodForPeap{
		"microsoftchapversiontwo": NonEapAuthenticationMethodForPeap_MicrosoftChapVersionTwo,
		"none":                    NonEapAuthenticationMethodForPeap_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NonEapAuthenticationMethodForPeap(input)
	return &out, nil
}
