package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsSModeConfiguration string

const (
	WindowsSModeConfiguration_Block         WindowsSModeConfiguration = "block"
	WindowsSModeConfiguration_NoRestriction WindowsSModeConfiguration = "noRestriction"
	WindowsSModeConfiguration_Unlock        WindowsSModeConfiguration = "unlock"
)

func PossibleValuesForWindowsSModeConfiguration() []string {
	return []string{
		string(WindowsSModeConfiguration_Block),
		string(WindowsSModeConfiguration_NoRestriction),
		string(WindowsSModeConfiguration_Unlock),
	}
}

func (s *WindowsSModeConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsSModeConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsSModeConfiguration(input string) (*WindowsSModeConfiguration, error) {
	vals := map[string]WindowsSModeConfiguration{
		"block":         WindowsSModeConfiguration_Block,
		"norestriction": WindowsSModeConfiguration_NoRestriction,
		"unlock":        WindowsSModeConfiguration_Unlock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsSModeConfiguration(input)
	return &out, nil
}
