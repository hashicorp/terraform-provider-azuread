package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsMinimumSessionSecurity string

const (
	LocalSecurityOptionsMinimumSessionSecurity_None                         LocalSecurityOptionsMinimumSessionSecurity = "none"
	LocalSecurityOptionsMinimumSessionSecurity_NtlmV2And128BitEncryption    LocalSecurityOptionsMinimumSessionSecurity = "ntlmV2And128BitEncryption"
	LocalSecurityOptionsMinimumSessionSecurity_Require128BitEncryption      LocalSecurityOptionsMinimumSessionSecurity = "require128BitEncryption"
	LocalSecurityOptionsMinimumSessionSecurity_RequireNtmlV2SessionSecurity LocalSecurityOptionsMinimumSessionSecurity = "requireNtmlV2SessionSecurity"
)

func PossibleValuesForLocalSecurityOptionsMinimumSessionSecurity() []string {
	return []string{
		string(LocalSecurityOptionsMinimumSessionSecurity_None),
		string(LocalSecurityOptionsMinimumSessionSecurity_NtlmV2And128BitEncryption),
		string(LocalSecurityOptionsMinimumSessionSecurity_Require128BitEncryption),
		string(LocalSecurityOptionsMinimumSessionSecurity_RequireNtmlV2SessionSecurity),
	}
}

func (s *LocalSecurityOptionsMinimumSessionSecurity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsMinimumSessionSecurity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsMinimumSessionSecurity(input string) (*LocalSecurityOptionsMinimumSessionSecurity, error) {
	vals := map[string]LocalSecurityOptionsMinimumSessionSecurity{
		"none":                         LocalSecurityOptionsMinimumSessionSecurity_None,
		"ntlmv2and128bitencryption":    LocalSecurityOptionsMinimumSessionSecurity_NtlmV2And128BitEncryption,
		"require128bitencryption":      LocalSecurityOptionsMinimumSessionSecurity_Require128BitEncryption,
		"requirentmlv2sessionsecurity": LocalSecurityOptionsMinimumSessionSecurity_RequireNtmlV2SessionSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsMinimumSessionSecurity(input)
	return &out, nil
}
