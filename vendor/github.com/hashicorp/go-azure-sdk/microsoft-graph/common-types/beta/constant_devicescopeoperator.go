package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeOperator string

const (
	DeviceScopeOperator_Equals DeviceScopeOperator = "equals"
	DeviceScopeOperator_None   DeviceScopeOperator = "none"
)

func PossibleValuesForDeviceScopeOperator() []string {
	return []string{
		string(DeviceScopeOperator_Equals),
		string(DeviceScopeOperator_None),
	}
}

func (s *DeviceScopeOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceScopeOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceScopeOperator(input string) (*DeviceScopeOperator, error) {
	vals := map[string]DeviceScopeOperator{
		"equals": DeviceScopeOperator_Equals,
		"none":   DeviceScopeOperator_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceScopeOperator(input)
	return &out, nil
}
