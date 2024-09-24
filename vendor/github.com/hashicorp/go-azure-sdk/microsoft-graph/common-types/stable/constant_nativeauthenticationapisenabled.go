package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NativeAuthenticationApisEnabled string

const (
	NativeAuthenticationApisEnabled_All  NativeAuthenticationApisEnabled = "all"
	NativeAuthenticationApisEnabled_None NativeAuthenticationApisEnabled = "none"
)

func PossibleValuesForNativeAuthenticationApisEnabled() []string {
	return []string{
		string(NativeAuthenticationApisEnabled_All),
		string(NativeAuthenticationApisEnabled_None),
	}
}

func (s *NativeAuthenticationApisEnabled) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNativeAuthenticationApisEnabled(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNativeAuthenticationApisEnabled(input string) (*NativeAuthenticationApisEnabled, error) {
	vals := map[string]NativeAuthenticationApisEnabled{
		"all":  NativeAuthenticationApisEnabled_All,
		"none": NativeAuthenticationApisEnabled_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NativeAuthenticationApisEnabled(input)
	return &out, nil
}
