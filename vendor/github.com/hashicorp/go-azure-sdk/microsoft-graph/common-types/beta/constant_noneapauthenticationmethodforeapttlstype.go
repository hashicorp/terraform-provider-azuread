package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NonEapAuthenticationMethodForEapTtlsType string

const (
	NonEapAuthenticationMethodForEapTtlsType_ChallengeHandshakeAuthenticationProtocol NonEapAuthenticationMethodForEapTtlsType = "challengeHandshakeAuthenticationProtocol"
	NonEapAuthenticationMethodForEapTtlsType_MicrosoftChap                            NonEapAuthenticationMethodForEapTtlsType = "microsoftChap"
	NonEapAuthenticationMethodForEapTtlsType_MicrosoftChapVersionTwo                  NonEapAuthenticationMethodForEapTtlsType = "microsoftChapVersionTwo"
	NonEapAuthenticationMethodForEapTtlsType_UnencryptedPassword                      NonEapAuthenticationMethodForEapTtlsType = "unencryptedPassword"
)

func PossibleValuesForNonEapAuthenticationMethodForEapTtlsType() []string {
	return []string{
		string(NonEapAuthenticationMethodForEapTtlsType_ChallengeHandshakeAuthenticationProtocol),
		string(NonEapAuthenticationMethodForEapTtlsType_MicrosoftChap),
		string(NonEapAuthenticationMethodForEapTtlsType_MicrosoftChapVersionTwo),
		string(NonEapAuthenticationMethodForEapTtlsType_UnencryptedPassword),
	}
}

func (s *NonEapAuthenticationMethodForEapTtlsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNonEapAuthenticationMethodForEapTtlsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNonEapAuthenticationMethodForEapTtlsType(input string) (*NonEapAuthenticationMethodForEapTtlsType, error) {
	vals := map[string]NonEapAuthenticationMethodForEapTtlsType{
		"challengehandshakeauthenticationprotocol": NonEapAuthenticationMethodForEapTtlsType_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           NonEapAuthenticationMethodForEapTtlsType_MicrosoftChap,
		"microsoftchapversiontwo": NonEapAuthenticationMethodForEapTtlsType_MicrosoftChapVersionTwo,
		"unencryptedpassword":     NonEapAuthenticationMethodForEapTtlsType_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NonEapAuthenticationMethodForEapTtlsType(input)
	return &out, nil
}
