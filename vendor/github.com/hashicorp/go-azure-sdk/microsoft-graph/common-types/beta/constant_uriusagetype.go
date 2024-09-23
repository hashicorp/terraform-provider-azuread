package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UriUsageType string

const (
	UriUsageType_IdentifierUri UriUsageType = "identifierUri"
	UriUsageType_LoginUrl      UriUsageType = "loginUrl"
	UriUsageType_LogoutUrl     UriUsageType = "logoutUrl"
	UriUsageType_RedirectUri   UriUsageType = "redirectUri"
)

func PossibleValuesForUriUsageType() []string {
	return []string{
		string(UriUsageType_IdentifierUri),
		string(UriUsageType_LoginUrl),
		string(UriUsageType_LogoutUrl),
		string(UriUsageType_RedirectUri),
	}
}

func (s *UriUsageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUriUsageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUriUsageType(input string) (*UriUsageType, error) {
	vals := map[string]UriUsageType{
		"identifieruri": UriUsageType_IdentifierUri,
		"loginurl":      UriUsageType_LoginUrl,
		"logouturl":     UriUsageType_LogoutUrl,
		"redirecturi":   UriUsageType_RedirectUri,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UriUsageType(input)
	return &out, nil
}
