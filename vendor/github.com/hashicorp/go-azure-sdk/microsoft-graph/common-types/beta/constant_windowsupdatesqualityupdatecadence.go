package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCadence string

const (
	WindowsUpdatesQualityUpdateCadence_Monthly   WindowsUpdatesQualityUpdateCadence = "monthly"
	WindowsUpdatesQualityUpdateCadence_OutOfBand WindowsUpdatesQualityUpdateCadence = "outOfBand"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCadence() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCadence_Monthly),
		string(WindowsUpdatesQualityUpdateCadence_OutOfBand),
	}
}

func (s *WindowsUpdatesQualityUpdateCadence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesQualityUpdateCadence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesQualityUpdateCadence(input string) (*WindowsUpdatesQualityUpdateCadence, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCadence{
		"monthly":   WindowsUpdatesQualityUpdateCadence_Monthly,
		"outofband": WindowsUpdatesQualityUpdateCadence_OutOfBand,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCadence(input)
	return &out, nil
}
