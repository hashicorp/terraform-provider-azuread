package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeParameter string

const (
	DeviceScopeParameter_None     DeviceScopeParameter = "none"
	DeviceScopeParameter_ScopeTag DeviceScopeParameter = "scopeTag"
)

func PossibleValuesForDeviceScopeParameter() []string {
	return []string{
		string(DeviceScopeParameter_None),
		string(DeviceScopeParameter_ScopeTag),
	}
}

func (s *DeviceScopeParameter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceScopeParameter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceScopeParameter(input string) (*DeviceScopeParameter, error) {
	vals := map[string]DeviceScopeParameter{
		"none":     DeviceScopeParameter_None,
		"scopetag": DeviceScopeParameter_ScopeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeParameter(input)
	return &out, nil
}
