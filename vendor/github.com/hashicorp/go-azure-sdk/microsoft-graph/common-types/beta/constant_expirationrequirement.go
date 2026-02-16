package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationRequirement string

const (
	ExpirationRequirement_AudienceTokenLifetimePolicy                       ExpirationRequirement = "audienceTokenLifetimePolicy"
	ExpirationRequirement_NgcMfa                                            ExpirationRequirement = "ngcMfa"
	ExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices ExpirationRequirement = "rememberMultifactorAuthenticationOnTrustedDevices"
	ExpirationRequirement_SignInFrequencyEveryTime                          ExpirationRequirement = "signInFrequencyEveryTime"
	ExpirationRequirement_SignInFrequencyPeriodicReauthentication           ExpirationRequirement = "signInFrequencyPeriodicReauthentication"
	ExpirationRequirement_TenantTokenLifetimePolicy                         ExpirationRequirement = "tenantTokenLifetimePolicy"
)

func PossibleValuesForExpirationRequirement() []string {
	return []string{
		string(ExpirationRequirement_AudienceTokenLifetimePolicy),
		string(ExpirationRequirement_NgcMfa),
		string(ExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices),
		string(ExpirationRequirement_SignInFrequencyEveryTime),
		string(ExpirationRequirement_SignInFrequencyPeriodicReauthentication),
		string(ExpirationRequirement_TenantTokenLifetimePolicy),
	}
}

func (s *ExpirationRequirement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExpirationRequirement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExpirationRequirement(input string) (*ExpirationRequirement, error) {
	vals := map[string]ExpirationRequirement{
		"audiencetokenlifetimepolicy": ExpirationRequirement_AudienceTokenLifetimePolicy,
		"ngcmfa":                      ExpirationRequirement_NgcMfa,
		"remembermultifactorauthenticationontrusteddevices": ExpirationRequirement_RememberMultifactorAuthenticationOnTrustedDevices,
		"signinfrequencyeverytime":                          ExpirationRequirement_SignInFrequencyEveryTime,
		"signinfrequencyperiodicreauthentication":           ExpirationRequirement_SignInFrequencyPeriodicReauthentication,
		"tenanttokenlifetimepolicy":                         ExpirationRequirement_TenantTokenLifetimePolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpirationRequirement(input)
	return &out, nil
}
