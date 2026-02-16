package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SingleSignOnMode string

const (
	SingleSignOnMode_AadHeaderBased     SingleSignOnMode = "aadHeaderBased"
	SingleSignOnMode_None               SingleSignOnMode = "none"
	SingleSignOnMode_OAuthToken         SingleSignOnMode = "oAuthToken"
	SingleSignOnMode_OnPremisesKerberos SingleSignOnMode = "onPremisesKerberos"
	SingleSignOnMode_PingHeaderBased    SingleSignOnMode = "pingHeaderBased"
	SingleSignOnMode_Saml               SingleSignOnMode = "saml"
)

func PossibleValuesForSingleSignOnMode() []string {
	return []string{
		string(SingleSignOnMode_AadHeaderBased),
		string(SingleSignOnMode_None),
		string(SingleSignOnMode_OAuthToken),
		string(SingleSignOnMode_OnPremisesKerberos),
		string(SingleSignOnMode_PingHeaderBased),
		string(SingleSignOnMode_Saml),
	}
}

func (s *SingleSignOnMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSingleSignOnMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSingleSignOnMode(input string) (*SingleSignOnMode, error) {
	vals := map[string]SingleSignOnMode{
		"aadheaderbased":     SingleSignOnMode_AadHeaderBased,
		"none":               SingleSignOnMode_None,
		"oauthtoken":         SingleSignOnMode_OAuthToken,
		"onpremiseskerberos": SingleSignOnMode_OnPremisesKerberos,
		"pingheaderbased":    SingleSignOnMode_PingHeaderBased,
		"saml":               SingleSignOnMode_Saml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SingleSignOnMode(input)
	return &out, nil
}
