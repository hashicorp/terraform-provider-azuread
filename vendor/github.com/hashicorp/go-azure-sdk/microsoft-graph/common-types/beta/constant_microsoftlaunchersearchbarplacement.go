package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftLauncherSearchBarPlacement string

const (
	MicrosoftLauncherSearchBarPlacement_Bottom        MicrosoftLauncherSearchBarPlacement = "bottom"
	MicrosoftLauncherSearchBarPlacement_Hide          MicrosoftLauncherSearchBarPlacement = "hide"
	MicrosoftLauncherSearchBarPlacement_NotConfigured MicrosoftLauncherSearchBarPlacement = "notConfigured"
	MicrosoftLauncherSearchBarPlacement_Top           MicrosoftLauncherSearchBarPlacement = "top"
)

func PossibleValuesForMicrosoftLauncherSearchBarPlacement() []string {
	return []string{
		string(MicrosoftLauncherSearchBarPlacement_Bottom),
		string(MicrosoftLauncherSearchBarPlacement_Hide),
		string(MicrosoftLauncherSearchBarPlacement_NotConfigured),
		string(MicrosoftLauncherSearchBarPlacement_Top),
	}
}

func (s *MicrosoftLauncherSearchBarPlacement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftLauncherSearchBarPlacement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftLauncherSearchBarPlacement(input string) (*MicrosoftLauncherSearchBarPlacement, error) {
	vals := map[string]MicrosoftLauncherSearchBarPlacement{
		"bottom":        MicrosoftLauncherSearchBarPlacement_Bottom,
		"hide":          MicrosoftLauncherSearchBarPlacement_Hide,
		"notconfigured": MicrosoftLauncherSearchBarPlacement_NotConfigured,
		"top":           MicrosoftLauncherSearchBarPlacement_Top,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftLauncherSearchBarPlacement(input)
	return &out, nil
}
