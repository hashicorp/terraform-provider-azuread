package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencyAuthenticationType string

const (
	SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication SignInFrequencyAuthenticationType = "primaryAndSecondaryAuthentication"
	SignInFrequencyAuthenticationType_SecondaryAuthentication           SignInFrequencyAuthenticationType = "secondaryAuthentication"
)

func PossibleValuesForSignInFrequencyAuthenticationType() []string {
	return []string{
		string(SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication),
		string(SignInFrequencyAuthenticationType_SecondaryAuthentication),
	}
}

func (s *SignInFrequencyAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInFrequencyAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInFrequencyAuthenticationType(input string) (*SignInFrequencyAuthenticationType, error) {
	vals := map[string]SignInFrequencyAuthenticationType{
		"primaryandsecondaryauthentication": SignInFrequencyAuthenticationType_PrimaryAndSecondaryAuthentication,
		"secondaryauthentication":           SignInFrequencyAuthenticationType_SecondaryAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencyAuthenticationType(input)
	return &out, nil
}
