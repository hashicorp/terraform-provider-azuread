package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodPlatform string

const (
	AuthenticationMethodPlatform_Android AuthenticationMethodPlatform = "android"
	AuthenticationMethodPlatform_IOS     AuthenticationMethodPlatform = "iOS"
	AuthenticationMethodPlatform_Linux   AuthenticationMethodPlatform = "linux"
	AuthenticationMethodPlatform_MacOS   AuthenticationMethodPlatform = "macOS"
	AuthenticationMethodPlatform_Unknown AuthenticationMethodPlatform = "unknown"
	AuthenticationMethodPlatform_Windows AuthenticationMethodPlatform = "windows"
)

func PossibleValuesForAuthenticationMethodPlatform() []string {
	return []string{
		string(AuthenticationMethodPlatform_Android),
		string(AuthenticationMethodPlatform_IOS),
		string(AuthenticationMethodPlatform_Linux),
		string(AuthenticationMethodPlatform_MacOS),
		string(AuthenticationMethodPlatform_Unknown),
		string(AuthenticationMethodPlatform_Windows),
	}
}

func (s *AuthenticationMethodPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodPlatform(input string) (*AuthenticationMethodPlatform, error) {
	vals := map[string]AuthenticationMethodPlatform{
		"android": AuthenticationMethodPlatform_Android,
		"ios":     AuthenticationMethodPlatform_IOS,
		"linux":   AuthenticationMethodPlatform_Linux,
		"macos":   AuthenticationMethodPlatform_MacOS,
		"unknown": AuthenticationMethodPlatform_Unknown,
		"windows": AuthenticationMethodPlatform_Windows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodPlatform(input)
	return &out, nil
}
