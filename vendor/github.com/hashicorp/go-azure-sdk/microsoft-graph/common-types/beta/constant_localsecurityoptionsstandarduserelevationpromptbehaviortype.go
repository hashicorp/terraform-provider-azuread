package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsStandardUserElevationPromptBehaviorType string

const (
	LocalSecurityOptionsStandardUserElevationPromptBehaviorType_AutomaticallyDenyElevationRequests     LocalSecurityOptionsStandardUserElevationPromptBehaviorType = "automaticallyDenyElevationRequests"
	LocalSecurityOptionsStandardUserElevationPromptBehaviorType_NotConfigured                          LocalSecurityOptionsStandardUserElevationPromptBehaviorType = "notConfigured"
	LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentials                   LocalSecurityOptionsStandardUserElevationPromptBehaviorType = "promptForCredentials"
	LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop LocalSecurityOptionsStandardUserElevationPromptBehaviorType = "promptForCredentialsOnTheSecureDesktop"
)

func PossibleValuesForLocalSecurityOptionsStandardUserElevationPromptBehaviorType() []string {
	return []string{
		string(LocalSecurityOptionsStandardUserElevationPromptBehaviorType_AutomaticallyDenyElevationRequests),
		string(LocalSecurityOptionsStandardUserElevationPromptBehaviorType_NotConfigured),
		string(LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentials),
		string(LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop),
	}
}

func (s *LocalSecurityOptionsStandardUserElevationPromptBehaviorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsStandardUserElevationPromptBehaviorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsStandardUserElevationPromptBehaviorType(input string) (*LocalSecurityOptionsStandardUserElevationPromptBehaviorType, error) {
	vals := map[string]LocalSecurityOptionsStandardUserElevationPromptBehaviorType{
		"automaticallydenyelevationrequests":     LocalSecurityOptionsStandardUserElevationPromptBehaviorType_AutomaticallyDenyElevationRequests,
		"notconfigured":                          LocalSecurityOptionsStandardUserElevationPromptBehaviorType_NotConfigured,
		"promptforcredentials":                   LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentials,
		"promptforcredentialsonthesecuredesktop": LocalSecurityOptionsStandardUserElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsStandardUserElevationPromptBehaviorType(input)
	return &out, nil
}
