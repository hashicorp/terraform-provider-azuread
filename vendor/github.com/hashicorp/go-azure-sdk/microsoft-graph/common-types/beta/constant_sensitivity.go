package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Sensitivity string

const (
	Sensitivity_Confidential Sensitivity = "confidential"
	Sensitivity_Normal       Sensitivity = "normal"
	Sensitivity_Personal     Sensitivity = "personal"
	Sensitivity_Private      Sensitivity = "private"
)

func PossibleValuesForSensitivity() []string {
	return []string{
		string(Sensitivity_Confidential),
		string(Sensitivity_Normal),
		string(Sensitivity_Personal),
		string(Sensitivity_Private),
	}
}

func (s *Sensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivity(input string) (*Sensitivity, error) {
	vals := map[string]Sensitivity{
		"confidential": Sensitivity_Confidential,
		"normal":       Sensitivity_Normal,
		"personal":     Sensitivity_Personal,
		"private":      Sensitivity_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Sensitivity(input)
	return &out, nil
}
