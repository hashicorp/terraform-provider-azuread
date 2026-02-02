package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingVariant string

const (
	SharingVariant_AddressBar             SharingVariant = "addressBar"
	SharingVariant_Embed                  SharingVariant = "embed"
	SharingVariant_None                   SharingVariant = "none"
	SharingVariant_PasswordProtected      SharingVariant = "passwordProtected"
	SharingVariant_RequiresAuthentication SharingVariant = "requiresAuthentication"
)

func PossibleValuesForSharingVariant() []string {
	return []string{
		string(SharingVariant_AddressBar),
		string(SharingVariant_Embed),
		string(SharingVariant_None),
		string(SharingVariant_PasswordProtected),
		string(SharingVariant_RequiresAuthentication),
	}
}

func (s *SharingVariant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingVariant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingVariant(input string) (*SharingVariant, error) {
	vals := map[string]SharingVariant{
		"addressbar":             SharingVariant_AddressBar,
		"embed":                  SharingVariant_Embed,
		"none":                   SharingVariant_None,
		"passwordprotected":      SharingVariant_PasswordProtected,
		"requiresauthentication": SharingVariant_RequiresAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingVariant(input)
	return &out, nil
}
