package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PromptLoginBehavior string

const (
	PromptLoginBehavior_Disabled                               PromptLoginBehavior = "disabled"
	PromptLoginBehavior_NativeSupport                          PromptLoginBehavior = "nativeSupport"
	PromptLoginBehavior_TranslateToFreshPasswordAuthentication PromptLoginBehavior = "translateToFreshPasswordAuthentication"
)

func PossibleValuesForPromptLoginBehavior() []string {
	return []string{
		string(PromptLoginBehavior_Disabled),
		string(PromptLoginBehavior_NativeSupport),
		string(PromptLoginBehavior_TranslateToFreshPasswordAuthentication),
	}
}

func (s *PromptLoginBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePromptLoginBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePromptLoginBehavior(input string) (*PromptLoginBehavior, error) {
	vals := map[string]PromptLoginBehavior{
		"disabled":                               PromptLoginBehavior_Disabled,
		"nativesupport":                          PromptLoginBehavior_NativeSupport,
		"translatetofreshpasswordauthentication": PromptLoginBehavior_TranslateToFreshPasswordAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PromptLoginBehavior(input)
	return &out, nil
}
