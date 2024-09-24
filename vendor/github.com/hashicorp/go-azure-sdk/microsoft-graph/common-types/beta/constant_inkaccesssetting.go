package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InkAccessSetting string

const (
	InkAccessSetting_Disabled      InkAccessSetting = "disabled"
	InkAccessSetting_Enabled       InkAccessSetting = "enabled"
	InkAccessSetting_NotConfigured InkAccessSetting = "notConfigured"
)

func PossibleValuesForInkAccessSetting() []string {
	return []string{
		string(InkAccessSetting_Disabled),
		string(InkAccessSetting_Enabled),
		string(InkAccessSetting_NotConfigured),
	}
}

func (s *InkAccessSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInkAccessSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInkAccessSetting(input string) (*InkAccessSetting, error) {
	vals := map[string]InkAccessSetting{
		"disabled":      InkAccessSetting_Disabled,
		"enabled":       InkAccessSetting_Enabled,
		"notconfigured": InkAccessSetting_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InkAccessSetting(input)
	return &out, nil
}
