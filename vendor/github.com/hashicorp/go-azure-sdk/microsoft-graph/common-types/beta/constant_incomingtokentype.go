package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncomingTokenType string

const (
	IncomingTokenType_None                IncomingTokenType = "none"
	IncomingTokenType_PrimaryRefreshToken IncomingTokenType = "primaryRefreshToken"
	IncomingTokenType_RemoteDesktopToken  IncomingTokenType = "remoteDesktopToken"
	IncomingTokenType_Saml11              IncomingTokenType = "saml11"
	IncomingTokenType_Saml20              IncomingTokenType = "saml20"
)

func PossibleValuesForIncomingTokenType() []string {
	return []string{
		string(IncomingTokenType_None),
		string(IncomingTokenType_PrimaryRefreshToken),
		string(IncomingTokenType_RemoteDesktopToken),
		string(IncomingTokenType_Saml11),
		string(IncomingTokenType_Saml20),
	}
}

func (s *IncomingTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncomingTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncomingTokenType(input string) (*IncomingTokenType, error) {
	vals := map[string]IncomingTokenType{
		"none":                IncomingTokenType_None,
		"primaryrefreshtoken": IncomingTokenType_PrimaryRefreshToken,
		"remotedesktoptoken":  IncomingTokenType_RemoteDesktopToken,
		"saml11":              IncomingTokenType_Saml11,
		"saml20":              IncomingTokenType_Saml20,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncomingTokenType(input)
	return &out, nil
}
