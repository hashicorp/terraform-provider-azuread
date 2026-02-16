package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsAdministratorElevationPromptBehaviorType string

const (
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_ElevateWithoutPrompting                LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "elevateWithoutPrompting"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_NotConfigured                          LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "notConfigured"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsent                       LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "promptForConsent"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentForNonWindowsBinaries  LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "promptForConsentForNonWindowsBinaries"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentOnTheSecureDesktop     LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "promptForConsentOnTheSecureDesktop"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentials                   LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "promptForCredentials"
	LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop LocalSecurityOptionsAdministratorElevationPromptBehaviorType = "promptForCredentialsOnTheSecureDesktop"
)

func PossibleValuesForLocalSecurityOptionsAdministratorElevationPromptBehaviorType() []string {
	return []string{
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_ElevateWithoutPrompting),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_NotConfigured),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsent),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentForNonWindowsBinaries),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentOnTheSecureDesktop),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentials),
		string(LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop),
	}
}

func (s *LocalSecurityOptionsAdministratorElevationPromptBehaviorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsAdministratorElevationPromptBehaviorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsAdministratorElevationPromptBehaviorType(input string) (*LocalSecurityOptionsAdministratorElevationPromptBehaviorType, error) {
	vals := map[string]LocalSecurityOptionsAdministratorElevationPromptBehaviorType{
		"elevatewithoutprompting":                LocalSecurityOptionsAdministratorElevationPromptBehaviorType_ElevateWithoutPrompting,
		"notconfigured":                          LocalSecurityOptionsAdministratorElevationPromptBehaviorType_NotConfigured,
		"promptforconsent":                       LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsent,
		"promptforconsentfornonwindowsbinaries":  LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentForNonWindowsBinaries,
		"promptforconsentonthesecuredesktop":     LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForConsentOnTheSecureDesktop,
		"promptforcredentials":                   LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentials,
		"promptforcredentialsonthesecuredesktop": LocalSecurityOptionsAdministratorElevationPromptBehaviorType_PromptForCredentialsOnTheSecureDesktop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsAdministratorElevationPromptBehaviorType(input)
	return &out, nil
}
