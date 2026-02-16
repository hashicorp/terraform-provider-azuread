package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodClientAppName string

const (
	MicrosoftAuthenticatorAuthenticationMethodClientAppName_MicrosoftAuthenticator MicrosoftAuthenticatorAuthenticationMethodClientAppName = "microsoftAuthenticator"
	MicrosoftAuthenticatorAuthenticationMethodClientAppName_OutlookMobile          MicrosoftAuthenticatorAuthenticationMethodClientAppName = "outlookMobile"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodClientAppName() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppName_MicrosoftAuthenticator),
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppName_OutlookMobile),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMethodClientAppName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMethodClientAppName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMethodClientAppName(input string) (*MicrosoftAuthenticatorAuthenticationMethodClientAppName, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodClientAppName{
		"microsoftauthenticator": MicrosoftAuthenticatorAuthenticationMethodClientAppName_MicrosoftAuthenticator,
		"outlookmobile":          MicrosoftAuthenticatorAuthenticationMethodClientAppName_OutlookMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodClientAppName(input)
	return &out, nil
}
