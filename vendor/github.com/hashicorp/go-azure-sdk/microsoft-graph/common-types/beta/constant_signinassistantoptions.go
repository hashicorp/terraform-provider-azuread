package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAssistantOptions string

const (
	SignInAssistantOptions_Disabled      SignInAssistantOptions = "disabled"
	SignInAssistantOptions_NotConfigured SignInAssistantOptions = "notConfigured"
)

func PossibleValuesForSignInAssistantOptions() []string {
	return []string{
		string(SignInAssistantOptions_Disabled),
		string(SignInAssistantOptions_NotConfigured),
	}
}

func (s *SignInAssistantOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInAssistantOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInAssistantOptions(input string) (*SignInAssistantOptions, error) {
	vals := map[string]SignInAssistantOptions{
		"disabled":      SignInAssistantOptions_Disabled,
		"notconfigured": SignInAssistantOptions_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAssistantOptions(input)
	return &out, nil
}
