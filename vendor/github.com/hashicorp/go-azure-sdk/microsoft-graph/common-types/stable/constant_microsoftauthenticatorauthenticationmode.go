package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMode string

const (
	MicrosoftAuthenticatorAuthenticationMode_Any             MicrosoftAuthenticatorAuthenticationMode = "any"
	MicrosoftAuthenticatorAuthenticationMode_DeviceBasedPush MicrosoftAuthenticatorAuthenticationMode = "deviceBasedPush"
	MicrosoftAuthenticatorAuthenticationMode_Push            MicrosoftAuthenticatorAuthenticationMode = "push"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMode() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMode_Any),
		string(MicrosoftAuthenticatorAuthenticationMode_DeviceBasedPush),
		string(MicrosoftAuthenticatorAuthenticationMode_Push),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMode(input string) (*MicrosoftAuthenticatorAuthenticationMode, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMode{
		"any":             MicrosoftAuthenticatorAuthenticationMode_Any,
		"devicebasedpush": MicrosoftAuthenticatorAuthenticationMode_DeviceBasedPush,
		"push":            MicrosoftAuthenticatorAuthenticationMode_Push,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMode(input)
	return &out, nil
}
