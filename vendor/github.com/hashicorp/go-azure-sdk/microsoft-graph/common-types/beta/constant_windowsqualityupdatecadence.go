package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCadence string

const (
	WindowsQualityUpdateCadence_Monthly   WindowsQualityUpdateCadence = "monthly"
	WindowsQualityUpdateCadence_OutOfBand WindowsQualityUpdateCadence = "outOfBand"
)

func PossibleValuesForWindowsQualityUpdateCadence() []string {
	return []string{
		string(WindowsQualityUpdateCadence_Monthly),
		string(WindowsQualityUpdateCadence_OutOfBand),
	}
}

func (s *WindowsQualityUpdateCadence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdateCadence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdateCadence(input string) (*WindowsQualityUpdateCadence, error) {
	vals := map[string]WindowsQualityUpdateCadence{
		"monthly":   WindowsQualityUpdateCadence_Monthly,
		"outofband": WindowsQualityUpdateCadence_OutOfBand,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdateCadence(input)
	return &out, nil
}
